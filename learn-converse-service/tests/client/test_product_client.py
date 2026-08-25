from unittest.mock import Mock, patch

from app.client.product_client import ProductClient
from app.models.classification import ProductFilter


def mock_response(products):
    response = Mock()
    response.json.return_value = products
    response.raise_for_status.return_value = None
    return response


@patch("app.client.product_client.requests.get")
def test_get_products_without_filters(mock_get):
    mock_get.return_value = mock_response([
        {
            "id": "P001",
            "name": "iPhone 17",
            "category": "smartphone",
            "price": 79999,
            "stockQuantity": 12,
        }
    ])

    client = ProductClient("http://localhost:8080")

    products = client.get_products(ProductFilter())

    mock_get.assert_called_once_with(
        "http://localhost:8080/products",
        params={},
        timeout=5,
    )

    assert len(products) == 1
    assert products[0].id == "P001"
    assert products[0].name == "iPhone 17"
    assert products[0].stock_quantity == 12


@patch("app.client.product_client.requests.get")
def test_get_products_with_filters(mock_get):
    mock_get.return_value = mock_response([
        {
            "id": "P001",
            "name": "iPhone 17",
            "category": "smartphone",
            "price": 79999,
            "stockQuantity": 12,
        }
    ])

    client = ProductClient("http://localhost:8080")

    filters = ProductFilter(
        name="Samsung",
        category="phone",
        min_price=20000,
        max_price=50000,
    )

    products = client.get_products(filters)

    mock_get.assert_called_once_with(
        "http://localhost:8080/products",
        params={
            "name": "Samsung",
            "category": "phone",
            "min_price": 20000,
            "max_price": 50000,
        },
        timeout=5,
    )

    assert len(products) == 1
    assert products[0].id == "P001"


@patch("app.client.product_client.requests.get")
def test_get_products_raises_for_http_error(mock_get):
    response = Mock()
    response.raise_for_status.side_effect = Exception("HTTP 500")
    mock_get.return_value = response

    client = ProductClient("http://localhost:8080")

    try:
        client.get_products(ProductFilter())
        assert False, "Expected exception"
    except Exception as error:
        assert str(error) == "HTTP 500"