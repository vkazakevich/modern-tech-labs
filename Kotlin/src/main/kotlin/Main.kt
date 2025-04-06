import kotlinx.coroutines.*

suspend fun main(args: Array<String>) = coroutineScope {
    fillSeeds()

    val client = DiscordClient.create(
        token = Config.discordAuthToken, 
        channelId = Config.discordChannelID
    )

    val slackClient = SlackClient(
        appToken = Config.slackAppToken, 
        botToken = Config.slackBotToken
    )

    launch { client.run() }
    launch { slackClient.run() }

    val message = args.getOrNull(0) ?: "Hello everyone!"
    client.sendMessage(message)
}

fun fillSeeds() {
    categories.addAll(
        listOf(
            Category(name = "Mobile phones"),
            Category(name = "Game consoles"),
            Category(name = "Laptops")
        )
    )

    categories.forEach {
        val productsNames = mutableListOf<String>();

        if (it.name == "Mobile phones") {
            productsNames.addAll(listOf("iPhone", "Samsung", "Xiaomi"))
        }

        if (it.name == "Game consoles") {
            productsNames.addAll(listOf("PlayStation 5", "Xbox", "Nintendo Switch"))
        }

        if (it.name == "Laptops") {
            productsNames.addAll(listOf("Macbook", "HP", "Lenovo", "Acer"))
        }

        productsNames.forEach {
            name -> products.add(Product(name = name, category = it, price = (100..500).random()))
        }
    }
}