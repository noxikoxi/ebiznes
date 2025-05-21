import streamlit as st
import requests
import random

API_URL = "http://localhost:8000/chat"

initial_greetings = [
    "Cześć! Jestem asystentem sklepu odzieżowego. Jak mogę Ci dziś pomóc w kwestii ubrań?",
    "Witaj! Jestem tu, by odpowiedzieć na Twoje pytania o nasz sklep i asortyment. Czym mogę służyć?",
    "Dzień dobry! Szukasz czegoś konkretnego? Chętnie pomogę w wyborze odzieży lub odpowiem na pytania o nasz sklep.",
    "Hej! Masz pytania o nasze produkty, rozmiary, czy promocje? Śmiało pytaj!",
    "W czym mogę pomóc? Pamiętaj, że odpowiadam tylko na pytania związane z naszym sklepem odzieżowym."
]

farewell_messages = [
    "Miłego dnia! Wróć, jeśli masz więcej pytań.",
    "Dziękuję za rozmowę! Mam nadzieję, że pomogłem.",
    "Do zobaczenia! Pamiętaj o naszych promocjach.",
    "Życzę udanych zakupów! W razie potrzeby służę pomocą.",
    "Dziękuję za odwiedzenie naszego chatu. Zapraszamy ponownie!"
]

# --- Interfejs Streamlit ---
st.set_page_config(page_title="Sklepowy Chatbot", page_icon="🛍️")
st.title("🛍️ Asystent Sklepu Odzieżowego")

if "greetings" not in st.session_state:
    st.session_state.greetings = random.choice(initial_greetings)
    st.session_state.chat_history = []

# Wyświetlenie przywitania
with st.chat_message("assistant"):
    st.write(st.session_state.greetings)
# Wyświetlanie historii czatu
for message in st.session_state.chat_history:
    if message["role"] == "user":
        with st.chat_message("user"):
            st.write(message["content"])
    elif message["role"] == "assistant":
        with st.chat_message("assistant"):
            st.write(message["content"])

user_input = st.chat_input("Napisz wiadomość...")

if user_input:
    for word in user_input.lower().split(" "):
        for finish_word in ["koniec", "pa", "do widzenia", "dziekuje za pomoc", "dzięki", "dzieki", "dziękuję", "narazie"]:
            if word.startswith(finish_word):
                with st.chat_message("user"):
                    st.write(user_input)
                with st.chat_message("assistant"):
                    st.write(random.choice(farewell_messages))
                # koniec rozmowy
                st.session_state.chat_history = []
                st.stop()

    st.session_state.chat_history.append({"role": "user", "content": user_input})

    # Wyświetl wiadomość użytkownika natychmiast
    with st.chat_message("user"):
        st.write(user_input)

    # Wyślij wiadomość i całą historię do API
    with st.spinner("Bot myśli..."):
        try:
            response = requests.post(
                API_URL,
                json={"user_message": user_input, "chat_history": st.session_state.chat_history}
            )
            response.raise_for_status()  # Wyrzuca wyjątek dla kodów błędów HTTP (4xx lub 5xx)
            data = response.json()
            # Wyświetl odpowiedź bota
            with st.chat_message("assistant"):
                st.write(data["response"])

            st.session_state.chat_history.append({"role": "assistant", "content": data["response"]})

        except requests.exceptions.ConnectionError:
            st.error("Błąd: Nie można połączyć się z serwisem chatbota. Upewnij się, że backend działa (FastAPI).")
        except requests.exceptions.HTTPError as e:
            st.error(f"Błąd HTTP: {e}")
        except Exception as e:
            st.error(f"Wystąpił nieoczekiwany błąd: {e}")
