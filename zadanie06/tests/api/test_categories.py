import pytest
import requests
from utils.api_utils import get_base_url


def _get_endpoint_url():
    return f"{get_base_url()}/categories"


def test_get_categories():
    response = requests.get(_get_endpoint_url())
    assert response.status_code == 200
    products = response.json()
    assert isinstance(products, list)

@pytest.mark.parametrize("idx", [1, 2, 3])
def test_get_category_by_id_correct(idx):
    response = requests.get(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 200
    product = response.json()
    assert isinstance(product, dict)
    assert list(product.keys()) == ["id", "name"]


@pytest.mark.parametrize("idx", [999, 991, 990])
def test_get_category_by_id_incorrect_id(idx):
    response = requests.get(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 404
    assert response.json().get("error") == "Category not found"


@pytest.mark.parametrize("idx", ["text", "w", 15.523])
def test_get_category_by_id_incorrect_id_type(idx):
    response = requests.get(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 400
    assert response.json().get("error") == "Invalid ID"


def test_create_category_and_delete_correct():
    data = {"name": "test_category"}
    response = requests.post(f"{_get_endpoint_url()}", json=data)
    assert response.status_code == 201
    new_category_id = response.json()["id"]
    response = requests.delete(f"{_get_endpoint_url()}/{new_category_id}")
    assert response.status_code == 204


def test_crate_category_name_already_exists():
    data = {"name": "test_category"}
    response = requests.post(f"{_get_endpoint_url()}", json=data)
    assert response.status_code == 201
    new_category_id = response.json()["id"]

    response = requests.post(f"{_get_endpoint_url()}", json=data)
    assert response.status_code == 409

    response = requests.delete(f"{_get_endpoint_url()}/{new_category_id}")
    assert response.status_code == 204


@pytest.mark.parametrize("idx", [1312312, 323232, 434343])
def test_delete_category_invalid_id(idx):
    response = requests.delete(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 204


@pytest.mark.parametrize("idx", [13123.3232, "weird", "d"])
def test_delete_category_invalid_id_type(idx):
    response = requests.delete(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 400
    assert response.json().get("error") == "Invalid ID"


@pytest.mark.parametrize("idx", [1312312, 323232, 434343])
def test_edit_category_invalid_id(idx):
    data = {"name": "test_category"}
    response = requests.put(f"{_get_endpoint_url()}/{idx}", json=data)
    assert response.status_code == 404
    assert response.json().get("error") == "Category not found"


@pytest.mark.parametrize("idx", [13123.3232, "weird", "d"])
def test_edit_category_invalid_id_type(idx):
    data = {"name": "test_category"}
    response = requests.put(f"{_get_endpoint_url()}/{idx}", json=data)
    assert response.status_code == 400
    assert response.json().get("error") == "Invalid ID"


def test_edit_category_correct(idx=1, categoryName="New Premium Category"):
    old_category = requests.get(f"{_get_endpoint_url()}/{idx}").json().get("name")
    response = requests.put(f"{_get_endpoint_url()}/{idx}", json={"name": categoryName})
    assert response.status_code == 200
    assert response.json().get("name") == categoryName

    response = requests.put(f"{_get_endpoint_url()}/{idx}", json={"name": old_category})
    assert response.status_code == 200
    assert response.json().get("name") == old_category


def test_edit_category_invalid_data(idx=1):
    response = requests.put(f"{_get_endpoint_url()}/{idx}", json={"smth": ":))"})
    assert response.status_code == 400
