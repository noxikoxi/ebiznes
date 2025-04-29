import pytest
import requests

from utils.api_utils import get_base_url


def _get_endpoint_url():
    return f'{get_base_url()}/carts'


def test_get_carts_correct():
    response = requests.get(_get_endpoint_url())
    assert response.status_code == 200
    carts = response.json()
    assert isinstance(carts, list)
    assert list(carts[0].keys()) == ["id", "cart_items", "total"]


@pytest.mark.parametrize("idx", [3, 6])
def test_get_cart_by_idx_correct(idx):
    response = requests.get(f'{_get_endpoint_url()}/{idx}')
    assert response.status_code == 200
    carts = response.json()
    assert isinstance(carts, dict)
    assert list(carts.keys()) == ["id", "cart_items", "total"]


@pytest.mark.parametrize("idx", [999, 9991, 9992])
def test_get_cart_by_idx_incorrect_id(idx):
    response = requests.get(f'{_get_endpoint_url()}/{idx}')
    assert response.status_code == 404


@pytest.mark.parametrize("idx", ["strr", "d", 1312.3123])
def test_get_cart_by_idx_incorrect_id_type(idx):
    response = requests.get(f'{_get_endpoint_url()}/{idx}')
    assert response.status_code == 400


def test_create_and_delete_cart_correct():
    response = requests.post(f'{_get_endpoint_url()}', json={})
    idx = response.json().get('id')
    assert idx is not None
    assert response.status_code == 201

    response = requests.delete(f'{_get_endpoint_url()}/{idx}')
    assert response.status_code == 204


@pytest.mark.parametrize("idx", [999, 9991, 9992])
def test_delete_cart_incorrect_idx(idx):
    response = requests.delete(f'{_get_endpoint_url()}/{idx}')
    assert response.status_code == 204


@pytest.mark.parametrize("idx", ["strr", "d", 1312.3123])
def test_delete_cart_incorrect_idx_type(idx):
    response = requests.delete(f'{_get_endpoint_url()}/{idx}')
    assert response.status_code == 400
    assert response.json().get("error") == "Invalid ID"


def test_add_cart_item_and_delete_correct():
    response = requests.post(f'{_get_endpoint_url()}')
    new_cart_id = response.json().get('id')
    assert response.status_code == 201

    response = requests.post(f"{_get_endpoint_url()}/{new_cart_id}/1/99")
    assert response.status_code == 200
    cartItem = response.json()
    assert cartItem.get('cart_id') is not None
    assert cartItem.get('product_id') == 1
    assert cartItem.get('quantity') == 99

    response = requests.delete(f"{_get_endpoint_url()}/{new_cart_id}/1")
    assert response.status_code == 204

    response = requests.delete(f'{_get_endpoint_url()}/{new_cart_id}')
    assert response.status_code == 204


def test_add_cart_item_and_delete_incorrect_product_idx():
    response = requests.post(f'{_get_endpoint_url()}')
    new_cart_id = response.json().get('id')
    assert response.status_code == 201

    response = requests.post(f"{_get_endpoint_url()}/{new_cart_id}/9999/99")
    assert response.status_code == 404
    assert response.json().get('error') == "Product not found"

    response = requests.delete(f'{_get_endpoint_url()}/{new_cart_id}')
    assert response.status_code == 204


def test_update_cart_item_incorrect_cart_idx():
    response = requests.put(f"{_get_endpoint_url()}/9999/9999/99")
    assert response.status_code == 404
    assert response.json().get('error') == "Cart not found"


def test_update_cart_item_incorrect_cart_idx_type():
    response = requests.put(f"{_get_endpoint_url()}/fwfwf/9999/99")
    assert response.status_code == 400
    assert response.json().get('error') == "cart_id, product_id and quantity have to be uint type"


def test_update_cart_item_correct():
    response = requests.post(f'{_get_endpoint_url()}')
    new_cart_id = response.json().get('id')
    assert response.status_code == 201

    response = requests.post(f"{_get_endpoint_url()}/{new_cart_id}/1/55")
    assert response.status_code == 200
    assert response.json().get('quantity') == 55

    response = requests.put(f"{_get_endpoint_url()}/{new_cart_id}/1/25")
    assert response.status_code == 200
    assert response.json().get('quantity') == 25

    response = requests.delete(f'{_get_endpoint_url()}/{new_cart_id}')
    assert response.status_code == 204


def test_delete_cart_item_correct():
    response = requests.post(f'{_get_endpoint_url()}')
    new_cart_id = response.json().get('id')
    assert response.status_code == 201

    response = requests.post(f"{_get_endpoint_url()}/{new_cart_id}/1/55")
    assert response.status_code == 200
    assert response.json().get('quantity') == 55

    response = requests.delete(f"{_get_endpoint_url()}/{new_cart_id}/1")
    assert response.status_code == 204

    response = requests.delete(f'{_get_endpoint_url()}/{new_cart_id}')
    assert response.status_code == 204


def test_delete_cart_item_incorrect_idx():
    response = requests.post(f'{_get_endpoint_url()}')
    new_cart_id = response.json().get('id')
    assert response.status_code == 201

    response = requests.post(f"{_get_endpoint_url()}/{new_cart_id}/1/55")
    assert response.status_code == 200
    assert response.json().get('quantity') == 55

    response = requests.delete(f"{_get_endpoint_url()}/{new_cart_id}/2")
    assert response.status_code == 404

    response = requests.delete(f"{_get_endpoint_url()}/{new_cart_id}/1")
    assert response.status_code == 204

    response = requests.delete(f'{_get_endpoint_url()}/{new_cart_id}')
    assert response.status_code == 204
