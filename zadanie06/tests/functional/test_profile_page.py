from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from utils.test_utils import login, verify_page_title, wait_and_get, get_danger_text


def get_profile_inputs(driver):
    name = wait_and_get(driver, "input#given_name")
    surname = wait_and_get(driver, "input#surname")
    phone = wait_and_get(driver, "input#phone")
    address = wait_and_get(driver, "input#address")
    email = wait_and_get(driver, "input#email")
    return name, surname, phone, address, email


def test_title_and_logout(driver, base_url):
    login(driver, base_url)
    logout_btn = wait_and_get(driver, "a.logout-button")
    verify_page_title(driver, "Profil")
    logout_btn.click()
    WebDriverWait(driver, 10).until(
        EC.url_to_be(base_url + "/login")
    )
    assert driver.current_url == base_url + "/login"


def test_profile_invalid(driver, base_url):
    login(driver, base_url)
    name, surname, phone, address, email = get_profile_inputs(driver)
    assert name and surname and phone and address and email
    assert not email.is_enabled()
    for elem in (name, surname, phone, address, email):
        assert elem.is_displayed()
    phone.clear()
    phone.send_keys("555")
    save = wait_and_get(driver, "button.formButton")
    save.click()
    danger = get_danger_text(driver)
    assert danger.is_displayed() and danger.text == "Numer telefonu musi składać się z dokładnie 9 cyfr."


def test_profile_correct(driver, base_url):
    login(driver, base_url)
    name, surname, phone, address, email = get_profile_inputs(driver)
    save = wait_and_get(driver, "button.formButton")
    for elem in (name, surname, phone, address, email):
        if elem.is_enabled():
            elem.clear()
    phone.send_keys("555111222")
    name.send_keys("NAME")
    surname.send_keys("SURNAME")
    address.send_keys("ADDRESS")
    save.click()
    WebDriverWait(driver, 10).until(
        lambda d: get_profile_inputs(d)[0].get_attribute("value") == "NAME"
    )
    name, surname, phone, address, email = get_profile_inputs(driver)
    assert name.get_attribute("value") == "NAME"
    assert surname.get_attribute("value") == "SURNAME"
    assert phone.get_attribute("value") == "555111222"
    assert address.get_attribute("value") == "ADDRESS"


