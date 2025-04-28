from pages.shared import verify_page_title, nav_logged, nav_quest


def get_curr_url(base_url):
    return base_url + "/reservations"


def test_title(driver, base_url):
    driver.get(get_curr_url(base_url))
    verify_page_title(driver, "Rezerwacje")


def test_nav_quest(driver, base_url):
    nav_quest(driver, get_curr_url(base_url))


def test_nav_logged(driver, base_url):
    nav_logged(driver, base_url, get_curr_url(base_url))
