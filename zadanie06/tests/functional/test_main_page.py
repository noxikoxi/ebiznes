from pages.shared import verify_page_title, get_elements_by_selector


def test_title(driver, base_url):
    driver.get(base_url)
    verify_page_title(driver, "Express")


def test_navigation_links_count(driver, base_url, expected_count=4):
    driver.get(base_url)
    links = get_elements_by_selector(driver, "div.middleContainer a")
    assert len(links) == expected_count, f"Oczekiwano {expected_count} linków, znaleziono {len(links)}"


def test_links_clickable(driver, base_url):
    driver.get(base_url)
    links = get_elements_by_selector(driver, "div.middleContainer a")
    for link in links:
        assert link.is_enabled(), f"Link '{link.text}' nie jest klikalny"

