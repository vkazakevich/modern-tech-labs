import kotlinx.coroutines.*

suspend fun main(args: Array<String>) = coroutineScope {
    val client = DiscordClient(
        token = Config.discordAuthToken, 
        channelId = Config.discordChannelID
    )

    categories.addAll(
        listOf(
            Category(name = "Category 1"),
            Category(name = "Category 2"),
            Category(name = "Category 3")
        )
    )

    launch { client.run() }

    val message = args.getOrNull(0) ?: "Hello everyone!"
    val res: String = client.sendMessage(message)
}
