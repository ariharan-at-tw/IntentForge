from app.client.product_client import ProductClient
from app.models.classification import ProductFilter


class ProductQueryService:

    def __init__(self, product_client: ProductClient):
        self.product_client = product_client

    def search(self, filters: ProductFilter):
        return self.product_client.get_products(filters)