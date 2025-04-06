import com.typesafe.config.Config
import com.typesafe.config.ConfigFactory

object Config {
    val config: Config = ConfigFactory.load()

    val discordAuthToken: String = config.getString("services.discord.token")
    val discordChannelID: Long = config.getLong("services.discord.channel_id")

    val slackAppToken: String = config.getString("services.slack.app_token")
    val slackBotToken: String = config.getString("services.slack.bot_token")
}