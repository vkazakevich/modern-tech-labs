package clients

import utils.Config

object ClientFactory {
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