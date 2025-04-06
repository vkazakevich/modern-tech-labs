import io.ktor.http.*
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.Serializable
import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import dev.kord.core.behavior.channel.MessageChannelBehavior
import dev.kord.common.entity.Snowflake

class DiscordClient (val kord: Kord, val channel: MessageChannelBehavior) {
    companion object Factory {
        suspend fun create(token: String, channelId: Long) : DiscordClient {
            val kord = Kord(token)
            val channelSnowflakeId = Snowflake(channelId)

            return DiscordClient(
                Kord(token), 
                MessageChannelBehavior(channelSnowflakeId, kord)
            )
        }
    }

    suspend fun run() {
        kord.on<MessageCreateEvent> {
            if (message.author?.isBot != false) return@on

            val args = message.content.split(" ", limit=2).toList()

            when (args.get(0)) {
                "!ping" -> message.channel.createMessage("pong!")
                "!categories" -> message.channel.createMessage(getCategories())
                "!category" -> {
                    val category: Category? = categories.firstOrNull { it.name == args.get(1) }

                    if (category == null) {
                        message.channel.createMessage("Incorrect category!")
                        return@on
                    }

                    message.channel.createMessage(getProductsByCategory(category))
                }
            }
        }

        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }

    fun getCategories() : String {
        return categories.map { it.name }.joinToString(separator = "\n")
    }

    fun getProductsByCategory(category: Category) : String {
        return products
            .filter{ it.category == category }
            .map { "${it.name} (\$${it.price})" }
            .joinToString(separator = "\n")
    }

    suspend fun sendMessage(content: String) {
        channel.createMessage(content)
    }
}

@Serializable
data class DiscordMessage(val content: String)