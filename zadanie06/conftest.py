import pytest
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

LOCAL = False


def pytest_addoption(parser):
    parser.addoption(
        "--base-url",
        action="store",
        default="http://127.0.0.1:9000",  # bazowy url
        help="Bazowy URL dla testów"
    )

@pytest.hookimpl(tryfirst=True, hookwrapper=True)
def pytest_runtest_makereport(item, call):
    outcome = yield
    rep = outcome.get_result()

    driver = item.funcargs.get("driver")

    # tylko po głównej części testu (nie setup/teardown)
    if driver and not LOCAL and rep.when == "call":
        status = "passed" if rep.passed else "failed"
        reason = "Test passed" if rep.passed else str(rep.longrepr)

        driver.execute_script(f'''
            browserstack_executor: {{
                "action": "setSessionStatus",
                "arguments": {{
                    "status": "{status}",
                    "reason": "{reason}"
                }}
            }}
        ''')


@pytest.fixture(scope="function")
def driver():
    if LOCAL:
        options = Options()
        options.add_experimental_option("prefs", {
            "credentials_enable_service": False,
            "profile.password_manager_enabled": False,
        })
        options.add_argument("--guest")
        options.add_argument("--headless")
        driver = webdriver.Chrome(options=options)
    else:
        options = Options()
        options.add_experimental_option("prefs", {
            "credentials_enable_service": False,
            "profile.password_manager_enabled": False,
        })
        options.add_argument("--guest")
        options.set_capability('sessionName', 'MaszynaNaDzien tests')
        driver = webdriver.Remote(
            command_executor='http://localhost:9000',
            options=options)
    yield driver
    driver.quit()


@pytest.fixture(scope="function")
def base_url(request):
    return request.config.getoption("--base-url")
