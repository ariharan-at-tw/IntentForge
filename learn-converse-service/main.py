from app.classifier.product_classifier import ProductClassifier
from app.llm.model import LLMModel


def main():
    model = LLMModel()
    classifier = ProductClassifier(model)

    user_input = "Show me laptops under 80000"

    result = classifier.classify(user_input)

    print(result)


if __name__ == "__main__":
    main()