import kotlinx.coroutines.*

suspend fun main(args: Array<String>) = coroutineScope {
    val client = DiscordClient(
        token = Config.discordAuthToken, 
        channelId = Config.discordChannelID
    )

    launch { client.run() }

    val message = args.getOrNull(0) ?: "Hello everyone!"
    val res: String = client.sendMessage(message)
}
