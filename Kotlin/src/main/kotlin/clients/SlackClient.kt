import com.slack.api.bolt.App;
import com.slack.api.bolt.AppConfig;
import com.slack.api.bolt.socket_mode.SocketModeApp;
import com.slack.api.model.event.AppMentionEvent;

class SlackClient(val appToken: String, val botToken: String) : Client() {
    override suspend fun run() {
        val appConfig: AppConfig = AppConfig.builder()
            .singleTeamBotToken(botToken)
            .build()

        val app: App = App(appConfig)
        val socketModeApp: SocketModeApp = SocketModeApp(appToken, app)

        app.command("/hi") { req, ctx -> 
         ctx.ack(":wave: Hi!")
        };

        app.command("/categories") { req, ctx ->
         ctx.ack(displayCategories())
        };

        app.command("/category") { req, ctx ->
         ctx.ack(displayProductsByCategory(req.payload.text))
        };

        socketModeApp.start()
    }
}