from enum import Enum

from pydantic import BaseModel


class ProductFilter(BaseModel):
    name: str | None = None
    category: str | None = None
    min_price: float | None = None
    max_price: float | None = None


class Intent(str, Enum):
    LIST_PRODUCTS = "LIST_PRODUCTS"
    UNSUPPORTED = "UNSUPPORTED"


class Classification(BaseModel):
    intent: Intent


class ListProductsClassification(Classification):
    intent: Intent = Intent.LIST_PRODUCTS
    filters: ProductFilter


class UnsupportedClassification(Classification):
    intent: Intent = Intent.UNSUPPORTED
    code: str


ClassificationResult = ListProductsClassification | UnsupportedClassification