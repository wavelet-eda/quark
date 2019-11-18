import sys

from setuptools import setup
from setuptools.command.test import test as TestCommand


class PyTest(TestCommand):
    user_options = [('pytest-args=', 'a', "Arguments to pass into py.test")]

    def initialize_options(self):
        TestCommand.initialize_options(self)
        self.pytest_args = []

    def finalize_options(self):
        TestCommand.finalize_options(self)
        self.test_args = []
        self.test_suite = True

    def run_tests(self):
        import pytest

        errno = pytest.main(self.pytest_args)
        sys.exit(errno)


setup(
    name='wavelet',
    version='0.1',
    description='Hardware Design Language in Python',
    author='Will Werst',
    author_email='werstiv@gmail.com',
    packages=['wavelet'],
    scripts=[],
    install_requires=[
    ],
    tests_require=[
        'pytest',
    ],
    cmdclass={'test': PyTest},
)
