# Task 3: Kotlin


## Progress

- 3.0 Create a client application in Kotlin using the Ktor framework that allows sending messages to the Discord platform [Commit `95cfaf8`](https://github.com/vkazakevich/ebiznes/commit/95cfaf83fce5f6123a60b4fa97a5ef5d4aad2900)
- 3.5 The application is able to receive user messages from the Discord platform directed to the application (bot) [Commit `05f79ed`](https://github.com/vkazakevich/ebiznes/commit/05f79edd5b9d76ec748e87909a66cc2b3332ea93)
- 4.0 Returns a list of categories upon a specific user request [Commit `1d5e8a5`](https://github.com/vkazakevich/ebiznes/commit/1d5e8a5286368b8a857493194d6debb6dfc2f5d9)
- 4.5 Returns a list of products according to the requested category [Commit `898c95d`](https://github.com/vkazakevich/ebiznes/commit/898c95db5dbd7f7eec561d96714a75511b67d2a5)
- 5.0 The application will additionally support one of the platforms: Slack, Messenger, Webex [Commit `39bcc29`](https://github.com/vkazakevich/ebiznes/commit/39bcc29aa963e467f1519f02aaa0a3364a707a57)

## Using

1. `cp .env.example .env`
2. Edit `.env`
3. `docker build -t kotlin .`
4. `docker run --rm -it -v "$(pwd)":/home/kotlin-app --env-file .env kotlin`