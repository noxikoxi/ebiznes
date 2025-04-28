from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import TimeoutException


def verify_page_title(driver, expected_title):
    """Sprawdza, czy tytuł strony jest poprawny."""
    actual_title = driver.title
    assert actual_title == expected_title, f"Oczekiwano tytułu '{expected_title}', ale otrzymano '{actual_title}'"


def get_elements_by_selector(driver, selector):
    try:
        wait = WebDriverWait(driver, 10)
        return wait.until(EC.presence_of_all_elements_located((By.CSS_SELECTOR, selector)))
    except TimeoutException:
        raise AssertionError("Linki nawigacyjne nie zostały znalezione")


def get_email_password(driver):
    try:
        wait = WebDriverWait(driver, 10)
        email = wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, "input#email")))
        password = wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, "input#password")))
        return email, password
    except TimeoutException:
        raise AssertionError


def wait_and_get(driver, selector):
    try:
        wait = WebDriverWait(driver, 10)
        elem = wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, selector)))
        return elem
    except TimeoutException:
        raise AssertionError


def get_danger_text(driver):
    try:
        wait = WebDriverWait(driver, 10)
        return wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, "p.danger")))
    except TimeoutException:
        raise AssertionError


def main_page_link(driver, base_url, url):
    driver.get(url)
    link = wait_and_get(driver, "div.middleContainer a")
    link.click()
    WebDriverWait(driver, 0.5)
    assert driver.current_url == base_url + "/", f"URL nie pasuje. Oczekiwano: {base_url}, ale uzyskano: {driver.current_url}"


def login(driver, base_url, admin=False):
    driver.get(base_url + "/login")
    login_btn = wait_and_get(driver, "button.formButton")
    email, password = get_email_password(driver)
    email.clear()
    password.clear()
    if admin:
        email.send_keys("admin")
        password.send_keys("admin")
    else:
        email.send_keys("testowy@wp.pl")
        password.send_keys("test123")
    login_btn.click()


def nav_logged(driver, base_url):
    login(driver, base_url)
    nav_elements = get_elements_by_selector(driver, "nav ul li")
    assert len(nav_elements) == 4
    first_nav = nav_elements[0]
    second_nav = nav_elements[1]
    third_nav = nav_elements[2]
    fourth_nav = nav_elements[3]
    assert first_nav.text == "Moje Rezerwacje"
    assert second_nav.text == "Rezerwacje"
    assert third_nav.text == "Maszyny"
    assert fourth_nav.text == "Profil"
    for element in (first_nav, second_nav, third_nav, fourth_nav):
        assert element.is_enabled() is True


def nav_admin(driver, base_url):
    login(driver, base_url, admin=True)
    nav_elements = get_elements_by_selector(driver, "nav ul li")
    assert len(nav_elements) == 4
    first_nav = nav_elements[0]
    second_nav = nav_elements[1]
    third_nav = nav_elements[2]
    fourth_nav = nav_elements[3]
    assert first_nav.text == "Użytkownicy"
    assert second_nav.text == "Rezerwacje"
    assert third_nav.text == "Maszyny"
    assert fourth_nav.text == "Profil"
    for element in (first_nav, second_nav, third_nav, fourth_nav):
        assert element.is_enabled() is True


def nav_quest(driver, url):
    driver.get(url)
    nav_elements = get_elements_by_selector(driver, "nav ul li")
    assert len(nav_elements) == 2
    first_nav = nav_elements[0]
    second_nav = nav_elements[1]
    assert first_nav.text == "Rezerwacje"
    assert second_nav.text == "Maszyny"
    assert first_nav.is_enabled() is True
    assert second_nav.is_enabled() is True


def get_div_children(driver, parent_div_selector):
    WebDriverWait(driver, 5).until(EC.presence_of_element_located((By.CSS_SELECTOR, parent_div_selector)))
    parent_div = driver.find_element(By.CSS_SELECTOR, parent_div_selector)
    children = parent_div.find_elements(By.CSS_SELECTOR, "div")  # '>*' oznacza wszystkie bezpośrednie dzieci
    return children
