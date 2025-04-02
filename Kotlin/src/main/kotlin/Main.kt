import com.typesafe.config.Config
import com.typesafe.config.ConfigFactory
import io.ktor.http.*
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.Serializable

const val baseUrl = "https://discord.com/api/v10"

@Serializable
data class DiscordMessage(val content: String)

suspend fun main(args: Array<String>) {
    val config: Config = ConfigFactory.load()

    val discordAuthToken: String = config.getString("services.discord.token")
    val discordChannelID: Long = config.getLong("services.discord.channel_id")

    val message = args.getOrNull(0) ?: "Hello everyone!"

    val client = HttpClient(CIO) {
         install(ContentNegotiation) {
            json()
        }
    }

    val response: HttpResponse = client.post(baseUrl + "/channels/" + discordChannelID +"/messages") {
        contentType(ContentType.Application.Json)
        headers {
            append(HttpHeaders.Authorization, "Bot " + discordAuthToken)
        }
        setBody(DiscordMessage(content = message))
    }

    println(response.status)
    client.close()
}