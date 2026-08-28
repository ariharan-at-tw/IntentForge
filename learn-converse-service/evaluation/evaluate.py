import argparse
import json

from app.classifier.product_classifier import ProductClassifier
from app.llm.model import LLMModel


DATASET_FILE = "evaluation/product_queries.json"


def is_match(result, expected):
    return result.model_dump() == expected

def parse_arguments():
    parser = argparse.ArgumentParser(
        description="Evaluate the product classifier."
    )

    parser.add_argument(
        "--verbose",
        action="store_true",
        help="Print every query, expected result, and actual result.",
    )

    parser.add_argument(
        "--mismatches",
        action="store_true",
        help="Print only queries where the prediction does not match the expected result.",
    )

    return parser.parse_args()


def main():
    args = parse_arguments()

    with open(DATASET_FILE, "r") as file:
        dataset = json.load(file)

    model = LLMModel()
    classifier = ProductClassifier(model)

    total = 0
    passed = 0
    failed = 0
    errors = 0

    for item in dataset:
        query = item["query"]
        expected = item["expected"]

        total += 1

        try:
            result = classifier.classify(query)

            if is_match(result, expected):
                passed += 1
                status = "PASS"
            else:
                failed += 1
                status = "FAIL"

        except Exception as error:
            errors += 1
            status = "ERROR"
            result = None
            error_message = str(error)

        should_print = (
            args.verbose
            or (args.mismatches and status != "PASS")
        )

        if should_print:
            print("=" * 80)
            print(f"Query:    {query}")
            print(f"Expected: {expected}")

            if result is not None:
                print(f"Actual:   {result}")
            else:
                print(f"Error:    {error_message}")

            print(f"Result:   {status}")

    print("=" * 80)
    print("Evaluation Summary")
    print("=" * 80)
    print(f"Total:    {total}")
    print(f"Passed:   {passed}")
    print(f"Failed:   {failed}")
    print(f"Errors:   {errors}")

    if total > 0:
        print(f"Accuracy: {passed / total * 100:.1f}%")


if __name__ == "__main__":
    main()

