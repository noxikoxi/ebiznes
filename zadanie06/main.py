from selenium import webdriver
import pytest
from selenium.webdriver.chrome.options import Options


@pytest.fixture
def driver(request):
    headless = request.config.getoption("--headless")
    chrome_options = Options()
    if headless:
        chrome_options.add_argument("--headless")
    driver = webdriver.Chrome(options=chrome_options)
    yield driver
    driver.quit()


def test_example(driver):
    driver.get("https://example.com")
    assert "Example" in driver.title
