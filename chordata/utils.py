from copy import deepcopy


class FSM(object):

    def __init__(self, update, initial_state, inputs, outputs):
        self._update = update
        self._state = initial_state
        self._inputs = inputs
        self._outputs = outputs

    @property
    def state(self):
        return deepcopy(self._state)

    def __repr__(self):
        return ""


class Stream(object):

    def __init__(self, size=1):
        self._size = 1

    def read():
        """Blocks until a value is pushed into stream."""
        pass

    def read_nb():
        """Return None if no data, else return result of read()
            without affecting the state of stream.
            This is similar to a peek() method."""
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
