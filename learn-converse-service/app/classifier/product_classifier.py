import json
from unittest import result
from urllib import response

from app.llm.model import LLMModel
from app.models.classification import ClassificationResult

from app.classifier.validator import validate_classification

import json
from pathlib import Path

from pydantic import TypeAdapter

classification_adapter = TypeAdapter(ClassificationResult)


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

Convert the user's query into exactly one valid JSON classification.

SUPPORTED INTENTS

1. LIST_PRODUCTS
2. UNSUPPORTED

LIST_PRODUCTS

Use LIST_PRODUCTS only when the complete request can be represented
using these filters:

- name
- category
- min_price
- max_price

Return exactly:

{
  "intent": "LIST_PRODUCTS",
  "filters": {
    "name": null,
    "category": null,
    "min_price": null,
    "max_price": null
  }
}

UNSUPPORTED

Use UNSUPPORTED when the user asks for anything that cannot be
represented using name, category, min_price, or max_price.

Examples:
- stock quantity
- availability
- rating
- sorting
- shipping
- discounts

Never ignore or silently drop an unsupported requirement.

Return exactly:

{
  "intent": "UNSUPPORTED",
  "code": "UNSUPPORTED_QUERY"
}

Do not include "filters" for UNSUPPORTED.

FILTER RULES

1. Category

Extract the product category and use a singular form.

Examples:
"laptops" -> "laptop"
"smartphones" -> "smartphone"

2. Name

Extract a specific product name or brand.

Do not populate name when the user specifies only a category.

If a brand and category are both present, keep them separate.

Example:
"Samsung phones" ->
name = "Samsung"
category = "phone"

3. Price

Only populate a price field when the user explicitly provides
a numeric price.

Lower-price conditions:

"under X"
"below X"
"less than X"
"up to X"

-> max_price = X
-> min_price = null

Higher-price conditions:

"above X"
"over X"
"more than X"
"at least X"

-> min_price = X
-> max_price = null

Range:

"between X and Y"

-> min_price = X
-> max_price = Y

IMPORTANT:
Preserve the direction of the user's price condition.
Do not swap min_price and max_price.

Do not infer numeric prices from words such as:
cheap, affordable, expensive, premium, budget.

4. Missing filters

Use null when a filter is not specified.

Never use empty strings.

OUTPUT

Return ONLY valid JSON.
Do not return markdown.
Do not return explanations.
"""


class ProductClassifier:

    def __init__(self, model: LLMModel):
        self.model = model

    def classify(self, user_input: str) -> ClassificationResult:
        prompt = (
            SYSTEM_PROMPT
            + "\n\n"
            + "Here are examples of the expected classification behavior:\n"
            + build_prompt_examples()
        )

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

        print("LLM RESPONSE:")
        print(response)

        data = json.loads(response)

        result = classification_adapter.validate_python(data)

        validate_classification(result)

        return result