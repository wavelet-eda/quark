from chordata.utils import (
    FSM,
    Stream,
    Streamable,
)

from chordata.types import (
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


output_fsm = FSM(
    update=output_switch_update,
    initial_state=OutputSwitchState(),
    inputs=[MultiplyResultID]*5,
    outputs=[MultiplyResult])


