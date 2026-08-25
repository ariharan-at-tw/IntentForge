from app.classifier.product_classifier import ProductClassifier
from app.client.product_client import ProductClient
from app.models.product import Product


class ProductQueryService:

    def __init__(
        self,
        classifier: ProductClassifier,
        product_client: ProductClient,
    ):
        self.classifier = classifier
        self.product_client = product_client

    def search(self, query: str) -> list[Product]:
        classification = self.classifier.classify(query)

        if classification.intent != "LIST_PRODUCTS":
            raise ValueError(
                f"Unsupported intent: {classification.intent}"
            )

        return self.product_client.get_products(
            classification.filters
        )