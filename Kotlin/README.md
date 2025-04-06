# Task 3: Kotlin


## Progress

- 3.0 Create a client application in Kotlin using the Ktor framework that allows sending messages to the Discord platform
- 3.5 The application is able to receive user messages from the Discord platform directed to the application (bot)
- 4.0 Returns a list of categories upon a specific user request
- 4.5 Returns a list of products according to the requested category
- 5.0 The application will additionally support one of the platforms: Slack, Messenger, Webex

## Using

1. `cp .env.example .env`
2. Edit `.env`
3. `docker build -t kotlin .`
4. `docker run --rm -it -v "$(pwd)":/home/kotlin-app --env-file .env kotlin /bin/bash`