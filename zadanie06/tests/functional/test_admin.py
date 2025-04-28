from selenium.webdriver.common.alert import Alert
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

from utils.test_utils import login, nav_admin, wait_and_get, get_danger_text


def get_create_user_inputs(driver):
    email = wait_and_get(driver, "input#email")
    password = wait_and_get(driver, "input#password")
    return email, password


def get_user_table_row(driver, row):
    table = WebDriverWait(driver, 10).until(EC.presence_of_element_located((By.TAG_NAME, "tbody")))
    rows = table.find_elements(By.TAG_NAME, "tr")
    return rows[row]


def test_navigation(driver, base_url):
    nav_admin(driver, base_url)


def test_login(driver, base_url):
    login(driver, base_url, admin=True)
    driver.get(base_url + "/admin/users")
    WebDriverWait(driver, 3).until(EC.url_contains("admin"))
    assert driver.current_url == base_url + "/admin/users"


def test_users(driver, base_url):
    login(driver, base_url, admin=True)
    WebDriverWait(driver, 10).until(EC.url_contains("users"))
    driver.get(base_url + "/admin/users")
    WebDriverWait(driver, 10).until(EC.url_to_be(base_url + "/admin/users"))
    info = wait_and_get(driver, "p.info.text-lg")
    assert info.text == "Użytkownicy"
    add_user = wait_and_get(driver, "a.button-like")
    add_user.click()
    WebDriverWait(driver, 5).until(EC.url_contains("create"))
    assert driver.current_url == base_url + "/admin/users/create"


def prepare_create_test(driver, base_url):
    login(driver, base_url, admin=True)
    driver.get(base_url + "/admin/users/create")
    WebDriverWait(driver, 3).until(EC.url_contains("create"))
    btn = wait_and_get(driver, "button.formButton")
    email, password = get_create_user_inputs(driver)
    return btn, email, password


def test_create_users_empty(driver, base_url):
    btn, _, _ = prepare_create_test(driver, base_url)
    btn.click()
    dangers = WebDriverWait(driver, 5).until(EC.presence_of_all_elements_located((By.CSS_SELECTOR, "p.danger")))
    assert len(dangers) == 2
    assert dangers[0].text == "Nieprawidłowy email"
    assert dangers[1].text == "Hasło musi mieć przynajmniej 6 znaków"


def test_create_user_short_password(driver, base_url, email_text="correct@wp.pl", password_text="123"):
    btn, email, password = prepare_create_test(driver, base_url)
    email.send_keys(email_text)
    password.send_keys(password_text)
    btn.click()
    danger = get_danger_text(driver)
    assert danger.text == "Hasło musi mieć przynajmniej 6 znaków"


def test_create_user_correct(driver, base_url, email_text="correct@wp.pl", password_text="correct1234"):
    btn, email, password = prepare_create_test(driver, base_url)
    email.send_keys(email_text)
    password.send_keys(password_text)
    btn.click()
    WebDriverWait(driver, 5).until(EC.url_to_be(base_url + "/admin/users"))

    assert driver.current_url == base_url + "/admin/users"
    last_row = get_user_table_row(driver, -1)
    email = last_row.find_elements(By.TAG_NAME, "td")[3]
    assert email.text == email_text


def test_delete_user(driver, base_url, email_text="correct@wp.pl"):
    login(driver, base_url, admin=True)
    driver.get(base_url + "/admin/users")
    WebDriverWait(driver, 5).until(EC.url_contains("admin"))
    last_row = get_user_table_row(driver, -1)
    trash = last_row.find_element(By.TAG_NAME, "button")
    trash.click()
    WebDriverWait(driver, 5).until(EC.alert_is_present())
    alert = Alert(driver)
    alert.accept()

    last_row = get_user_table_row(driver, -1)
    email = last_row.find_elements(By.TAG_NAME, "td")[3]
    assert email.text != email_text
