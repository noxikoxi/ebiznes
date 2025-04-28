from pages.shared import verify_page_title, wait_and_get, get_email_password, get_danger_text, main_page_link


def get_curr_url(base_url):
    return base_url + "/register"


def test_title(driver, base_url):
    driver.get(get_curr_url(base_url))
    verify_page_title(driver, "register")


def test_link(driver, base_url):
    driver.get(get_curr_url(base_url))
    link = wait_and_get(driver, "a.additionalInfo")
    link.click()
    assert driver.current_url == base_url + "/login"


def test_register_bad_cases(driver, base_url):
    driver.get(get_curr_url(base_url))
    register_btn = wait_and_get(driver, "button.formButton")
    email, password = get_email_password(driver)
    assert password.get_attribute("type") == "password"

    # no credentials
    register_btn.click()
    assert driver.current_url == get_curr_url(base_url)

    email.send_keys("testowy")
    register_btn.click()
    assert driver.current_url == get_curr_url(base_url)

    email, password = get_email_password(driver)
    email.send_keys("testowy")
    password.send_keys("testowy")
    register_btn.click()
    assert driver.current_url == get_curr_url(base_url)
    danger = get_danger_text(driver)
    assert danger.text == "Nieprawidłowy email"
    assert driver.current_url == get_curr_url(base_url)

    register_btn = wait_and_get(driver, "button.formButton")
    email, password = get_email_password(driver)
    email.send_keys("testowy@wp.pl")
    password.send_keys("test")
    register_btn.click()
    danger = get_danger_text(driver)
    assert danger.text == "Hasło musi mieć przynajmniej 6 znaków"
    assert driver.current_url == get_curr_url(base_url)

    register_btn = wait_and_get(driver, "button.formButton")
    email, password = get_email_password(driver)
    email.send_keys("testowy@wp.pl")
    password.send_keys("testowy")
    register_btn.click()
    danger = get_danger_text(driver)
    assert danger.text == "Użytkownik o takim adresie email już istnieje"
    assert driver.current_url == get_curr_url(base_url)


def test_main_page_link(driver, base_url):
    main_page_link(driver, base_url, get_curr_url(base_url))


# def test_register_correct(driver, base_url):
#     driver.get(get_curr_url(base_url))
#     register_btn = wait_and_get(driver, "button")
#     login, password = get_email_password(driver)
#
#     login.send_keys("testowy123@wp.pl")
#     password.send_keys("testowyPassword")
#     register_btn.click()
#
#     assert driver.current_url == base_url + "/login"
