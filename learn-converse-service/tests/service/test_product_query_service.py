from unittest.mock import Mock

import pytest

from app.models.classification import ClassificationResult, ProductFilter
from app.models.product import Product
from app.service.product_query_service import ProductQueryService


def test_search_classifies_query_and_fetches_products():
    classifier = Mock()
    product_client = Mock()

    classification = ClassificationResult(
        intent="LIST_PRODUCTS",
        filters=ProductFilter(
            category="laptop",
            max_price=60000,
        ),
    )

    expected_products = [
        Product(
            id="P002",
            name="MacBook Air",
            category="laptop",
            price=59999,
            stockQuantity=5,
        )
    ]

    classifier.classify.return_value = classification
    product_client.get_products.return_value = expected_products

    service = ProductQueryService(
        classifier=classifier,
        product_client=product_client,
    )

    result = service.search("Show me laptops under 60000")

    classifier.classify.assert_called_once_with(
        "Show me laptops under 60000"
    )

    product_client.get_products.assert_called_once_with(
        classification.filters
    )

    assert result == expected_products


def test_search_rejects_unsupported_intent():
    classifier = Mock()
    product_client = Mock()

    classifier.classify.return_value = ClassificationResult(
        intent="GET_PRODUCT",
        filters=ProductFilter(),
    )

    service = ProductQueryService(
        classifier=classifier,
        product_client=product_client,
    )

    with pytest.raises(ValueError, match="Unsupported intent: GET_PRODUCT"):
        service.search("Show me product P001")

    product_client.get_products.assert_not_called()