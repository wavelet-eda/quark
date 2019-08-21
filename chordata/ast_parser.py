import ast
import inspect
from textwrap import dedent

import astpretty


class VerilogConverterVisitor(ast.NodeVisitor):
    def __init__(self, _globals):
        self._globals = _globals

    def generic_visit(self, node):
        # print(f"**{ast.dump(node)}")
        # raise TypeError(f"Unsupported node {ast.dump(node)}")
        ast.NodeVisitor.generic_visit(self, node)

    def visit_Call(self, node):
        print(ast.dump(node))
        ast.NodeVisitor.generic_visit(self, node)

    def visit_Compare(self, node):
        print(ast.dump(node))
        ast.NodeVisitor.generic_visit(self, node)

    def visit_If(self, node):
        print(ast.dump(node))
        ast.NodeVisitor.generic_visit(self, node)

    def visit_BinOp(self, node):
        print(ast.dump(node))
        ast.NodeVisitor.generic_visit(self, node)

    def visit_Return(self, node):
        print(ast.dump(node))
        ast.NodeVisitor.generic_visit(self, node)

    def visit_Name(self, node):
        print(ast.dump(node))
        ast.NodeVisitor.generic_visit(self, node)


def parse_and_print(func):
    source = inspect.getsource(func)
    t = ast.parse(dedent(source))
    # print(ast.dump(t, include_attributes=True))
    stmts = t.body[0].body
    print(func.__globals__['something_else'])
    print(len(stmts))
    # astpretty.pprint(t)
    visitor = VerilogConverterVisitor(func.__globals__)
    [visitor.visit(stmt) for stmt in stmts]
