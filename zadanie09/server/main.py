from transformers import AutoTokenizer, AutoModelForCausalLM, BitsAndBytesConfig
import torch
from fastapi import FastAPI
from pydantic import BaseModel

# Model
model_id = "google/gemma-3-4b-it"
tokenizer = AutoTokenizer.from_pretrained(model_id)
model = AutoModelForCausalLM.from_pretrained(
    model_id,
    device_map="cuda",
    torch_dtype=torch.bfloat16,
    quantization_config=BitsAndBytesConfig(
        load_in_4bit=True,
    )
).eval()

app = FastAPI()


class ChatRequest(BaseModel):
    user_message: str  # Wiadomość od użytkownika (wymagana)
    chat_history: list = []  # Opcjonalna historia rozmowy, domyślnie pusta lista


# --- Endpointy API ---

# Testowy
@app.get("/")
async def read_root():
    return {"message": "FastAPI Chatbot service is running!"}


@app.post("/chat")
async def chat_endpoint(request: ChatRequest):
    """
    Endpoint do obsługi czatu z modelem LLM.
    Przyjmuje wiadomość użytkownika i historię rozmowy,
    zwraca odpowiedź modelu i zaktualizowaną historię.
    """

    user_message = request.user_message
    chat = request.chat_history  # Historia rozmowy od frontendu

    conversation_for_model = [{
        "role": "system",
        "content": "Jesteś asystentem klienta w sklepie odzieżowym. Odpowiadaj wyłącznie na pytania dotyczące "
                   "produktów (ubrania, akcesoria), dostępności, cen, rozmiarów, zwrotów, promocji i ogólnych "
                   "informacji o sklepie. Nie odpowiadaj na pytania niezwiązane ze sklepem odzieżowym. Jeśli "
                   "pytanie wykracza poza zakres sklepu, grzecznie odmów. Twoja odpowiedzi muszą być wyłącznie miłe i "
                   "z szacunkiem dla klienta."
    }]

    conversation_for_model.extend(chat)
    prompt = tokenizer.apply_chat_template(conversation_for_model, tokenize=False, add_generation_prompt=True)
    inputs = tokenizer.encode(prompt, add_special_tokens=False, return_tensors="pt").to(model.device)
    # Generowanie odpowiedzi przez model LLM
    outputs = model.generate(input_ids=inputs, max_new_tokens=150, temperature=0.7, top_p=0.9, do_sample=True)
    response_text = tokenizer.decode(outputs[0][len(inputs[0]):], skip_special_tokens=True).strip()

    return {"response": response_text}
