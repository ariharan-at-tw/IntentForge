import requests

from app.models.classification import ProductFilter
from app.models.product import Product


class ProductClient:

    def __init__(self, base_url: str):
        self.base_url = base_url.rstrip("/")

    def get_products(self, filters: ProductFilter) -> list[Product]:
        params = {}

        if filters.name:
            params["name"] = filters.name

        if filters.category:
            params["category"] = filters.category

        if filters.min_price is not None:
            params["min_price"] = filters.min_price

        if filters.max_price is not None:
            params["max_price"] = filters.max_price

        response = requests.get(
            f"{self.base_url}/products",
            params=params,
            timeout=5,
        )

        response.raise_for_status()

        return [
            Product.model_validate(product)
            for product in response.json()
        ]