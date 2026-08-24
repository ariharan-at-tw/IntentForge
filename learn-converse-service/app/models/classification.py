from pydantic import BaseModel


class ProductFilter(BaseModel):
    name: str | None = None
    category: str | None = None
    min_price: float | None = None
    max_price: float | None = None


class ClassificationResult(BaseModel):
    intent: str
    filters: ProductFilter