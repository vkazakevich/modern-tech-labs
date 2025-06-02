from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from ollama import chat, Client, ChatResponse
from pydantic import BaseModel

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
class Question(BaseModel):
    content: str

@app.post("/send_question")
def send_question(question: Question):
    client = Client(host='http://host.docker.internal:11434')

    response: ChatResponse = client.chat(model='llama3.2', messages=[
        {
            'role': 'user',
            'content': question.content,
        },
    ])

    return {
        "content": response.message.content
    }
