

class BitVector(object):

    def __init__(self, size=1):
        self._value = [bool()]*size

    @property
    def value(self):
        return self._val

    @value.setter
    def value(self, value):
        self._value = value

    def __get__(self, index):
        return self._value[index]


class Readable(object):
    pass


class Writeable(Readable):

    @property
    def foo(self):
        return self._foo

    @foo.setter
    def foo(self, value):
        self._foo = value
