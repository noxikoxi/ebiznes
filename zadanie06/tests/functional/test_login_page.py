from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

from utils.test_utils import verify_page_title, wait_and_get, get_email_password, get_danger_text, main_page_link, login


def get_curr_url(base_url):
    return base_url + "/login"


def test_login_without_credentials(driver, base_url):
    driver.get(get_curr_url(base_url))
    login_btn = driver.find_element(By.CSS_SELECTOR, "button.formButton")
    login_btn.click()
    assert driver.current_url == base_url + "/login"


def prepare_test(driver, base_url):
    driver.get(get_curr_url(base_url))
    login_btn = wait_and_get(driver, "button.formButton")
    email, password = get_email_password(driver)
    return login_btn, email, password


def test_login_empty(driver, base_url):
    login_btn, _, _ = prepare_test(driver, base_url)
    login_btn.click()
    assert driver.current_url == get_curr_url(base_url)


def test_login_bad_email(driver, base_url):
    login_btn, email, password = prepare_test(driver, base_url)
    assert password.get_attribute("type") == "password"
    email.send_keys("testowy")
    password.send_keys("testowy")
    login_btn.click()
    assert driver.current_url == get_curr_url(base_url)
    danger = get_danger_text(driver)
    assert danger.text == "Nieprawidłowy email"
    assert driver.current_url == get_curr_url(base_url)


def test_invalid_credentials(driver, base_url):
    login_btn, email, password = prepare_test(driver, base_url)
    email.send_keys("testowy@wp.pl")
    password.send_keys("testowy")
    login_btn.click()
    danger = get_danger_text(driver)
    assert danger.text == "Nieprawidłowe dane logowania"
    assert driver.current_url == get_curr_url(base_url)


def test_title(driver, base_url):
    driver.get(get_curr_url(base_url))
    verify_page_title(driver, "login")


def test_link(driver, base_url):
    driver.get(get_curr_url(base_url))
    link = wait_and_get(driver, "a.additionalInfo")
    link.click()
    WebDriverWait(driver, 10).until(EC.url_to_be(base_url + "/register"))
    assert driver.current_url == base_url + "/register"


def test_main_page_link(driver, base_url):
    main_page_link(driver, base_url, get_curr_url(base_url))


def test_correct_login(driver, base_url):
    driver.get(get_curr_url(base_url))
    login(driver, base_url)
    WebDriverWait(driver, 10).until(EC.url_to_be(base_url + "/users/7"))
    assert driver.current_url == base_url + "/users/7"
