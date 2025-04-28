from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

from utils.test_utils import verify_page_title, wait_and_get, get_elements_by_selector, login, nav_logged, nav_quest, \
    get_div_children


def get_curr_url(base_url):
    return base_url + "/machines"


def test_title(driver, base_url):
    driver.get(get_curr_url(base_url))
    verify_page_title(driver, "Maszyny")


def test_nav_quest(driver, base_url):
    nav_quest(driver, get_curr_url(base_url))


def test_nav_logged(driver, base_url):
    nav_logged(driver, base_url)


def test_content_quest(driver, base_url):
    driver.get(get_curr_url(base_url))
    items = get_div_children(driver, "div.machine-item")
    assert len(items) == 3


def test_content_logged(driver, base_url):
    login(driver, base_url)
    driver.get(get_curr_url(base_url))
    WebDriverWait(driver, 5).until(EC.url_contains("machines"))
    children = get_div_children(driver, "div.machine-item")
    assert len(children) == 4
    reservation = driver.find_element(By.CSS_SELECTOR, "div.machine-reserve a")
    assert reservation
