import json
from unittest import result

from app.llm.model import LLMModel
from app.models.classification import (
    ClassificationResult,
    ProductFilter,
)

from app.classifier.validator import validate_classification

import json
from pathlib import Path


PROMPT_EXAMPLES_FILE = (
    Path(__file__).resolve().parents[2]
    / "evaluation"
    / "prompt_examples.json"
)


def load_prompt_examples():
    with open(PROMPT_EXAMPLES_FILE, "r") as file:
        return json.load(file)

def build_prompt_examples():
    examples = load_prompt_examples()

    result = []

    for example in examples:
        result.append(
            f"""
User: {example["query"]}

Expected:
{json.dumps(example["expected"], indent=2)}
"""
        )

    return "\n".join(result)


SYSTEM_PROMPT = """
You are an e-commerce query classifier.

Convert the user's query into JSON.

There is exactly one supported intent:
LIST_PRODUCTS

The JSON must have exactly this structure:

{
  "intent": "LIST_PRODUCTS",
  "filters": {
    "name": null,
    "category": null,
    "min_price": null,
    "max_price": null
  }
}

Rules:

1. intent must ALWAYS be "LIST_PRODUCTS".

2. category should contain the product category.
   Use singular category names.
   Examples:
   "laptops" -> "laptop"
   "smartphones" -> "smartphone"

3. name should contain a specific product name or brand.

4. Do not populate name when the user specifies only a category.

5. Price rules:
   "under X" means max_price = X.
   "below X" means max_price = X.
   "less than X" means max_price = X.
   "up to X" means max_price = X.

   "above X" means min_price = X.
   "over X" means min_price = X.
   "more than X" means min_price = X.

   "between X and Y" means:
   min_price = X
   max_price = Y

6. Use null when a filter is not specified.

7. Never use empty strings. Use null instead.

8. Extract product name and category separately.

If the query contains a brand followed by a product category,
put the brand in "name" and the product type in "category".

Example:
"Samsung phones" →
name = "Samsung"
category = "phone"

"Samsung mobiles" →
name = "Samsung"
category = "mobile"

9. Never infer numeric price values.

Only populate min_price or max_price when the user explicitly
provides a numeric price.

Words such as:
"cheap", "affordable", "expensive", "premium", "budget"
must NOT be converted into numeric prices.

Example:
"Show me affordable smartphones" →
min_price = null
max_price = null

10. Return ONLY valid JSON.
   Do not include markdown.
   Do not include explanations.

Examples:
"""

class ProductClassifier:

    def __init__(self, model: LLMModel):
        self.model = model

    def classify(self, user_input: str) -> ClassificationResult:
        prompt = SYSTEM_PROMPT + "\n" + build_prompt_examples()

        messages = [
            {
                "role": "system",
                "content": prompt,
            },
            {
                "role": "user",
                "content": user_input,
            },
        ]

        response = self.model.generate(messages)

        data = json.loads(response)

        result = ClassificationResult.model_validate(data)

        validate_classification(result)

        return result