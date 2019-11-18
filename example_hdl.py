from wavelet.utils import (
    FSM,
    Stream,
    Streamable,
)

from wavelet.types import (
    BitVector,
)


class MultiplyCommand(Streamable):

    def __init__(self, a, b):
        self._a = a
        self._b = b

    def __repr__(self):
        return f"a = {self._a} b = {self._b}"


class MultiplyCommandID(Streamable):

    def __init__(self, a, b, ident):
        super().__init__(a, b)
        self._ident = ident

    def __repr__(self):
        return f"ident = {self._ident} " + super().__repr__()


class MultiplyResult(Streamable):

    def __init__(self, c):
        self._c = c


class MultiplyResultID(MultiplyResult):

    def __init__(self, c, ident):
        super().__init__(c)
        self._ident = ident

    def __repr__(self):
        return f"ident = {self._ident} " + super().__repr__()


class OutputSwitchState(Streamable):

    def __init__(self):
        self.ident = BitVector(size=5)  # Compile time how to know what size?
        self.ident.value = 0


def output_switch_update(fsm):
    next_state = fsm.state
    for o in fsm.inputs:
        res = o.read_nb()
        if res is not None and res.ident == fsm.state.ident:
            next_state.ident += 1
            fsm.outputs[0].write(o.read())
    return next_state


class InputSwitchState(Streamable):

    def __init__(self):
        self.ident = BitVector(size=5)  # Compile time how to know what size?
        self.ident.value = 0


def input_switch_update(fsm):
    next_state = fsm.state
    req = fsm.inputs[0].read()
    req_with_id = MultiplyCommandID(req.a, req.b, fsm.state.ident)
    for o in fsm.outputs:
        res = o.write_nb(req_with_id)
        if res:
            next_state.ident.value += 1
            break
    return next_state


class MultiplierState(Streamable):

    def __init__(self, operand):
        self.index = BitVector()
        self.operand = operand
        self.accumulator = BitVector()


def multiplier_update(fsm):
    next_state = fsm.state
    data_input = fsm.inputs[0].read_nb()
    if data_input is not None:
        if data_input.b[fsm.state.index]:
            next_state.accumulator += fsm.state.operand
        next_state.operand <<= 1
        next_state.index += 1
        if fsm.state.index == len(fsm.state.operand) - 1:
            fsm.inputs[0].read()
            data_output = MultiplyResultID(
                next_state.accumulator, data_input.ident)
            fsm.outputs[0].write(data_output)
            next_state = MultiplierState()
    return next_state


MULTIPLIER_COUNT = 5


input_fsm = FSM(
    update=input_switch_update,
    initial_state=InputSwitchState(),
    inputs=[MultiplyCommand],
    outputs=[MultiplyCommandID]*MULTIPLIER_COUNT)


multiplier_fsm = FSM(
        update=multiplier_update,
        initial_state=MultiplierState(),
        inputs=[MultiplyCommandID],
        outputs=[MultiplyResultID])


output_fsm = FSM(
    update=output_switch_update,
    initial_state=OutputSwitchState(),
    inputs=[MultiplyResultID]*MULTIPLIER_COUNT,
    outputs=[MultiplyResult])


input_stream = [MultiplyCommand()]
multiplier_inputs = input_fsm(inputs=input_stream)
multiplier_outputs = [multiplier_fsm(inputs=i) for i in multiplier_inputs]
output_stream = output_fsm(inputs=multiplier_outputs)

output_stream.synthesize()
### Synthesizes the full graph needed to generate output stream
