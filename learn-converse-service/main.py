from fastapi import FastAPI
from fastapi.responses import FileResponse

from app.models.chat import ChatRequest
from app.classifier.product_classifier import ProductClassifier
from app.client.product_client import ProductClient
from app.llm.model import LLMModel
from app.service.product_query_service import ProductQueryService
from app.service.chat_service import ChatService

app = FastAPI(title="Learn Converse Service")


model = LLMModel()
classifier = ProductClassifier(model)

product_client = ProductClient(
    base_url="http://localhost:8080"
)

product_query_service = ProductQueryService(
    product_client=product_client,
)

chat_service = ChatService(
    classifier=classifier,
    product_query_service=product_query_service,
)


@app.get("/health")
def health():
    return {"status": "ok"}

@app.post("/chat")
def chat(request: ChatRequest):
    print(f"Received chat request: {request.message}")
    return chat_service.process(request.message)

@app.get("/")
def index():
    return FileResponse("static/index.html")

if __name__ == "__main__":
    model = LLMModel()
    classifier = ProductClassifier(model)
    print(classifier.classify("list the products above 10000"))