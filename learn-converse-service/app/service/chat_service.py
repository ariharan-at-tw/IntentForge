from app.classifier.product_classifier import ProductClassifier
from app.service.product_query_service import ProductQueryService
from app.models.classification import (
    ListProductsClassification,
    UnsupportedClassification,
)

MESSAGES = {
    "UNSUPPORTED_QUERY": (
        "This query isn't currently supported.\n\n"
        "You can currently search products by:\n"
        "• Name\n"
        "• Category\n"
        "• Price"
    )
}

class ChatService:

    def __init__(
        self,
        classifier: ProductClassifier,
        product_query_service: ProductQueryService,
    ):
        self.classifier = classifier
        self.product_query_service = product_query_service

    def process(self, user_input: str):
        classification = self.classifier.classify(user_input)

        if isinstance(classification, UnsupportedClassification):
            return {
                "code": MESSAGES[classification.code]
            }

        if isinstance(classification, ListProductsClassification):
            products = self.product_query_service.search(
                classification.filters
            )

            return {
                "products": products
            }

        raise ValueError(
            f"Unsupported classification: {classification}"
        )