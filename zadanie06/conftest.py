import pytest


def pytest_addoption(parser):
    parser.addoption("--headless", action="store_true", help="Run browser in headless mode")
