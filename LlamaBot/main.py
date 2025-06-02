from fastapi import FastAPI
from ollama import chat, Client, ChatResponse

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.post("/chat")
def request_chat(content: str):
    client = Client(host='http://host.docker.internal:11434')

    response: ChatResponse = client.chat(model='llama3.2', messages=[
        {
            'role': 'user',
            'content': content,
        },
    ])

    return response.message.content
