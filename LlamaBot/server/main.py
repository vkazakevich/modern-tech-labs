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

    start_prompt = random.choice(OPEN_TEMPLATES)
    end_prompt = random.choice(CLOSE_TEMPLATES)

    system_prompt = f"""
    AI: helpful, concise, friendly.
    - If user's first message is exactly `START_CHAT`: Reply only with "{start_prompt}".
    - If user signals conversation end (e.g., "thank you", "solved", "goodbye", "that's all"): Reply only with "{end_prompt}". This must be your entire response.
    - All other replies: Address the user's query or statement directly. If the user says "hi" or a similar greeting *after the conversation has started*, do NOT use "{start_prompt}" or any other formal greeting. Instead, you can briefly acknowledge their greeting (e.g., "Hi there!" or simply "Hello!") and then continue the ongoing conversation, or just proceed with the current topic without a separate acknowledgment if that feels more natural. Avoid repeated formal greetings.
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


@app.get("/templates")
def read_templates():
    return {
        "open": OPEN_TEMPLATES,
        "close": CLOSE_TEMPLATES
    }
