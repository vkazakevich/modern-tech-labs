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

const val DISCORD_BASE_URL = "https://discord.com/api/v10"

class DiscordClient (val token: String, val channelId: Long) {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json()
        }
    }

    suspend fun run() {
        val kord = Kord(token)

        kord.on<MessageCreateEvent> {
            if (message.author?.isBot != false) return@on
            if (message.content != "!ping") return@on

            message.channel.createMessage("pong!")
        }

        kord.on<MessageCreateEvent> {
            if (message.author?.isBot != false) return@on
            if (message.content != "!categories") return@on

            message.channel.createMessage(showCategories())
        }

        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }

    fun showCategories() : String {
        return categories.map { it.name }.joinToString(separator = "\n\n")
    }

    suspend fun sendMessage(content: String) : String {
        val response: HttpResponse = client.post(DISCORD_BASE_URL + "/channels/${channelId}/messages") {
            contentType(ContentType.Application.Json)
            headers {
                append(HttpHeaders.Authorization, "Bot ${token}")
            }
            setBody(DiscordMessage(content = content))
        }

        return response.status.toString()
    }
}

@Serializable
data class DiscordMessage(val content: String)