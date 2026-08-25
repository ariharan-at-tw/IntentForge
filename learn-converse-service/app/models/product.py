from pydantic import BaseModel, Field


class Product(BaseModel):
    id: str
    name: str
    category: str
    price: float
    stock_quantity: int = Field(alias="stockQuantity")
    