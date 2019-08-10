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


