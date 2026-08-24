from app.models.classification import ClassificationResult


SUPPORTED_INTENTS = {
    "LIST_PRODUCTS",
}


def validate_classification(result: ClassificationResult) -> None:
    if result.intent not in SUPPORTED_INTENTS:
        raise ValueError(
            f"Unsupported intent: {result.intent}"
        )

    filters = result.filters

    if filters.min_price is not None and filters.min_price < 0:
        raise ValueError(
            "min_price cannot be negative"
        )

    if filters.max_price is not None and filters.max_price < 0:
        raise ValueError(
            "max_price cannot be negative"
        )

    if (
        filters.min_price is not None
        and filters.max_price is not None
        and filters.min_price > filters.max_price
    ):
        raise ValueError(
            "min_price cannot be greater than max_price"
        )