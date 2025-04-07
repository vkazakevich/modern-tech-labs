enum class ClientPlatform { 
    Discord,
    Slack
}

class ClientFactory {
    companion object {
        suspend fun create(platform: ClientPlatform): Client {
            return when(platform) {
                ClientPlatform.Discord -> DiscordClient(Config.discordAuthToken)

                ClientPlatform.Slack -> SlackClient(
                    appToken = Config.slackAppToken, 
                    botToken = Config.slackBotToken
                )
            }
        }
    }
}