import ast
import inspect
from textwrap import dedent


def parse_and_print(func):
    source = inspect.getsource(func)
    t = ast.parse(dedent(source))
    # print(ast.dump(t, include_attributes=True))
    stmts = t.body[0].body
    print(len(stmts))
    [print(ast.dump(stmt)) for stmt in stmts]
