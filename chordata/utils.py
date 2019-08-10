

class FSM(object):

    def __init__(self, update, initial_state, inputs, outputs):
        self._update = update
        self._initial_state = initial_state
        self._state_type = type(initial_state)
        self._inputs = inputs
        self._outputs = outputs

    def __repr__(self):
        return ""


class Stream(object):

    def __init__(self, size=1):
        self._size = 1

    def read():
        """Blocks until a value is pushed into stream."""
        pass

    def read_nb():
        """Return None if no data, else return result of read()."""
        pass

    def write():
        pass

    def write_nb():
        pass

    def __repr__(self):
        pass


class Streamable(object):

    def __init__(self):
        pass

    def __repr__(self):
        pass
