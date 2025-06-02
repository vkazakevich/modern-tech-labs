from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from ollama import chat, Client, ChatResponse
from pydantic import BaseModel
import random

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=[
        "http://localhost:5173"
    ],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

OPEN_TEMPLATES = (
    "Hi! How can I help?",
    "Support here. How can I assist?",
    "Hello! Need help with an order or product?",
    "Welcome to support! What's your question?",
    "How can I help you today?"
)

CLOSE_TEMPLATES = (
    "Glad I could help! Have a great day.",
    "Thanks for contacting support! We're here if you need us again in the future.",
    "Happy to have assisted you. Take care!",
    "Hope that resolved everything for you. Goodbye for now!",
    "Thanks for your patience. Have a good one!"
)


class Message(BaseModel):
    content: str


@app.post("/chat")
def send_message(message: Message):
    client = Client(host='http://host.docker.internal:11434')

    open_message = random.choice(OPEN_TEMPLATES)
    end_message = random.choice(CLOSE_TEMPLATES)

    if message.content.strip() == "[START_CHAT]":
        return {
            "content": open_message
        }

    system_prompt = f"""
    AI for online clothing store.
    - User signals end (e.g., "thank you", "solved", "goodbye", "that's all"): Reply only with "{end_message}". (Entire response).
    - Other replies: Answer user's query directly. If user says "hi" or a similar greeting standalone mid-conversation, do NOT greet back. Continue with the current topic or await their actual query. You should not send any greetings at any time.
    """

    response: ChatResponse = client.chat(model='llama3.2', messages=[
        {
            'role': 'system',
            'content': system_prompt,
        },
        {
            'role': 'user',
            'content': message.content,
        },
    ])

    return {
        "content": response.message.content
    }
