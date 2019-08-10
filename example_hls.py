# Define state container
class MyState(object):

    def __init__(self):
        self.count = 0

    def update()

# Define input container
class MyInput:
    pass


# Define output container
class MyOutput:
    pass


# Define fsm
def update(state, fsm):
    value = fsm.read(0)
    
    value, valid = fsm.read_nb(0)

    fsm.write(0, MyOutput())
    state.count += 1
    return state


fsm0 = create_fsm(update1, MyState, (MyInput, MyInput), (MyOutput))

fsm1 = create_fsm(update2, MyState, (MyOutput), (MyFinalOutput))


fifo0 = create_fifo(MyOutput, 8)

fifo0.input = fsm0.outputs[0]
fsm1.inputs[0] = fifo0.output
