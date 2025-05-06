import pytest
import requests
from utils.api_utils import get_base_url


def _get_endpoint_url():
    return f"{get_base_url()}/products"


def test_get_products():
    response = requests.get(_get_endpoint_url())
    assert response.status_code == 200
    products = response.json()
    assert isinstance(products, list)


def test_get_products_invalid_min_price():
    response = requests.get(f"{get_base_url()}/products?minPrice=abc")
    assert response.status_code == 400


def test_get_products_min_price_correct():
    response = requests.get(f"{get_base_url()}/products?minPrice=100")
    assert response.status_code == 200
    assert len(response.json()) != 0


def test_get_products_filter_no_results():
    response = requests.get(f"{_get_endpoint_url()}?minPrice=999999&categoryID=999")
    assert response.status_code == 200
    assert response.json() is None


@pytest.mark.parametrize("idx", [1, 2, 3])
def test_get_product_by_id_correct(idx):
    response = requests.get(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 200
    product = response.json()
    assert isinstance(product, dict)
    assert list(product.keys()) == ["id", "name", "price", "category"]


@pytest.mark.parametrize("idx", [999, 991, 990])
def test_get_product_by_id_incorrect_id(idx):
    response = requests.get(f"{_get_endpoint_url()}/{idx}")
    print(response.json())
    assert response.status_code == 404
    assert response.json().get("error") == "Product not found"


@pytest.mark.parametrize("idx", ["text", "w", 15.523])
def test_get_product_by_id_incorrect_id_type(idx):
    response = requests.get(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 400
    assert response.json().get("error") == "Invalid ID"


def test_create_and_delete_product_correct():
    response = requests.post(f"{_get_endpoint_url()}", json={"name": "test", "price": 1999.0, "category_id": 1})
    assert response.status_code == 201
    idx = response.json().get("id")
    response = requests.delete(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 204


def test_create_product_incorrect_data():
    response = requests.post(f"{_get_endpoint_url()}", json={"xddd": "test", "323": 1999.0, "dasd": 1})
    assert response.status_code == 400


@pytest.mark.parametrize("idx", [999, 991, 990])
def test_delete_product_incorrect_idx(idx):
    response = requests.delete(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 404


@pytest.mark.parametrize("idx", ["text", "w", 15.523])
def test_delete_product_incorrect_idx_type(idx):
    response = requests.delete(f"{_get_endpoint_url()}/{idx}")
    assert response.status_code == 400


@pytest.mark.parametrize("idx", [1, 2, 3])
def test_update_product_incorrect_data(idx):
    response = requests.put(f"{_get_endpoint_url()}/{idx}", json={"xddd": "test", "323": 1999.0, "dasd": 1})
    assert response.status_code == 400


@pytest.mark.parametrize("idx", [999, 991, 990])
def test_update_product_incorrect_idx(idx):
    response = requests.put(f"{_get_endpoint_url()}/{idx}", json={"name": "test", "price": 1999.0, "category_id": 1})
    assert response.status_code == 404
