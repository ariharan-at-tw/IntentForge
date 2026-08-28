from app.models.classification import (
    Classification,
    ListProductsClassification,
)


def validate_classification(result: Classification) -> None:
    if isinstance(result, ListProductsClassification):
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