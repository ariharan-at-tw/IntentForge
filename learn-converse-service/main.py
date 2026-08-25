from app.classifier.product_classifier import ProductClassifier
from app.client.product_client import ProductClient
from app.llm.model import LLMModel
from app.service.product_query_service import ProductQueryService


def main():
    model = LLMModel()
    classifier = ProductClassifier(model)

    product_client = ProductClient(
        base_url="http://localhost:8080"
    )

    product_query_service = ProductQueryService(
        classifier=classifier,
        product_client=product_client,
    )

    user_input = "Show me laptops under 100000000"

    products = product_query_service.search(user_input)

    for product in products:
        print(product)


if __name__ == "__main__":
    main()