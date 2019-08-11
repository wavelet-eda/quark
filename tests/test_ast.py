from textwrap import dedent
import inspect
import ast


from chordata.ast_parser import parse_and_print


def test_load_function():
    def something_else(a, b):
        return a
    def something(a, b):
        if a > b:
            return a
        a = something_else(a, b)
        return b - a
    parse_and_print(something)
    assert False
