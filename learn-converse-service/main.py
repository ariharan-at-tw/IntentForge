from fastapi import FastAPI

from app.classifier.product_classifier import ProductClassifier
from app.client.product_client import ProductClient
from app.llm.model import LLMModel
from app.service.product_query_service import ProductQueryService


app = FastAPI(title="Learn Converse Service")


model = LLMModel()
classifier = ProductClassifier(model)

product_client = ProductClient(
    base_url="http://localhost:8080"
)

product_query_service = ProductQueryService(
    classifier=classifier,
    product_client=product_client,
)


@app.get("/health")
def health():
    return {"status": "ok"}