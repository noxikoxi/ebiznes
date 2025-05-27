package org.example

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.websocket.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.websocket.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import java.io.FileInputStream
import java.util.*

import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.*
import kotlin.random.Random

import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.serialization.kotlinx.json.json

const val CHAT_ENDPOINT = "http://localhost:8000/chat"
// Maksymalna liczba par wiadomości (użytkownik + bot) do przechowywania
const val MAX_CHAT_HISTORY_PAIRS = 10

@Serializable
data class ChatMessage(val role: String, val content: String)

@Serializable
data class ChatRequest(
    val user_message: String,
    val chat_history: List<ChatMessage> = emptyList()
)

@Serializable
data class ChatResponse(
    val response: String
)


suspend fun main() {
    val appProps = Properties()
    withContext(Dispatchers.IO) {
        appProps.load(FileInputStream(".env"))
    }
    val token = appProps.getProperty("TOKEN")

    val client = HttpClient(CIO){
        install(WebSockets)
        install(ContentNegotiation){
            json()
        }
    }
    val url = getGatewayUrl(client)
    var resume_gateway_url: String? = null
    var session_id: String? = null
    var lastSequence: Int? = null
    var identified = false
    var running = true
    var botId: String? = null

    // Odpowiedzi na żadania
    val categories = arrayOf("Damskie", "Męskie", "Dziecięce", "Akcesoria", "Promocje")

    val products = mapOf(
        "damskie" to listOf("Sukienki", "Spódnice", "Bluzki", "Spodnie damskie", "Swetry damskie", "Kurtki damskie", "Płaszcze damskie"),
        "męskie" to listOf("Koszule", "T-shirty męskie", "Spodnie męskie", "Bluzy męskie", "Swetry męskie", "Marynarki", "Kurtki męskie"),
        "dziecięce" to listOf("Body niemowlęce", "Spodnie dziecięce", "Sukienki dziecięce", "Kurtki dziecięce", "Piżamy dziecięce", "Zestawy dla dzieci"),
        "akcesoria" to listOf("Torebki", "Paski", "Szaliki", "Czapki", "Biżuteria", "Okulary przeciwsłoneczne", "Skarpetki"),
        "promocje" to listOf("Wyprzedaż letnia", "Black Friday", "Kolekcja Outlet", "Darmowa dostawa", "Kody rabatowe")
    )

    // ZARZĄDZANIE HISTORIĄ CZATU DLA UŻYTKOWNIKÓW
    val userChatHistories = mutableMapOf<String, MutableList<ChatMessage>>()

    while (running) {

        try {
            val gatewayUrl = resume_gateway_url ?: url
            println("Łączę się z: $gatewayUrl")

            client.webSocket(url) {
                var lastAckReceived = true
                // GUILD_MESSAGES + DIRECT_MESSAGES + GUILDS + MESSAGE_CONTENT
                val identifyPayload = """
                    {
                        "op": 2,
                        "d": {
                            "token": "$token",
                            "intents": 37377,
                            "properties": {
                                "os": "linux",
                                "browser": "ktor",
                                "device": "ktor"
                            }
                        }
                    }
                """.trimIndent()

                val resumePayload = """
                    {
                        "op": 6,
                        "d": {
                            "token": "$token",
                            "session_id": "$session_id",
                            "seq": ${lastSequence ?: "null"}
                        }
                    }
                """.trimIndent()


                for (frame in incoming) {
                    if (frame is Frame.Text) {
                        var text = frame.readText()
//                        println("Otrzymano: $text")
                        val json = Json.parseToJsonElement(text).jsonObject
                        when (json["op"]?.jsonPrimitive?.int) {
                            10 -> { // Hello event
                                val interval = json["d"]?.jsonObject?.get("heartbeat_interval")?.jsonPrimitive?.long
                                if (interval != null) {
                                    val jitter = Random.nextDouble(0.0, 1.0)
                                    val initialDelay = (interval * jitter).toLong()
                                    println("Heartbeat interval: $interval, initial delay: $initialDelay")

                                    launch {
                                        delay(initialDelay)
                                        println("Zaczynam wysyłanie heartbeat")
                                        while (true) {
                                            if (!lastAckReceived) {
                                                println("Brak ACK, zamykam połączenie")
                                                close(
                                                    CloseReason(
                                                        4000,
                                                        "No Heartbeat ACK"
                                                    )
                                                ) // Niestandardowy kod zamknięcia
                                                return@launch
                                            }
                                            val heartbeatPayload = "{\"op\": 1, \"d\": ${lastSequence ?: "null"}}"
//                                            println("Wysyłam heartbeat: $heartbeatPayload")
                                            send(Frame.Text(heartbeatPayload))
                                            lastAckReceived = false

                                            // identyfikacja po rozpoczęciu wysyłania heartbeat
                                            if (!identified && session_id != null && resume_gateway_url != null) {
                                                send(Frame.Text(resumePayload))
                                                println("Wysłano Resume")
                                            } else if (!identified) {
                                                send(Frame.Text(identifyPayload))
                                                println("Wysłano Identify")
                                                identified = true
                                            }

                                            delay(interval)
                                        }

                                    }

                                }
                            }

                            11 -> { // Heartbeat ACK
//                                println("Odebrałem ACK")
                                lastAckReceived = true
                            }

                            9 -> { // Invalid Session
                                println("Błąd: Invalid Session")
                            }

                            0 -> { // Dispatch (np. Ready)
                                lastSequence = json["s"]?.jsonPrimitive?.int
                                println("Dispatch, sekwencja: $lastSequence")
                                when (json["t"]?.jsonPrimitive?.content) {
                                    "READY" -> {
                                        println("Otrzymano READY - sesja rozpoczęta!")
                                        session_id = json["d"]?.jsonObject?.get("session_id").toString()
                                        resume_gateway_url = json["d"]?.jsonObject?.get("resume_gateway_url").toString()
                                        botId = json["d"]?.jsonObject?.get("user")?.jsonObject?.get("id")?.jsonPrimitive?.content
                                        println("Session ID: $session_id, Resume URL: $resume_gateway_url, Bot ID: $botId")
                                        identified = true
                                    }
                                    "RESUMED" -> {
                                        println("Otrzymano RESUMED - sesja wznowiona!")
                                    }
                                    "MESSAGE_CREATE" -> {
                                        val content = json["d"]?.jsonObject?.get("content")?.jsonPrimitive?.content
                                        val channelId = json["d"]?.jsonObject?.get("channel_id")?.jsonPrimitive?.content
                                        val author = json["d"]?.jsonObject?.get("author")?.jsonObject?.get("username")?.jsonPrimitive?.content
                                        val authorId = json["d"]?.jsonObject?.get("author")?.jsonObject?.get("id")?.jsonPrimitive?.content
                                        val guildId = json["d"]?.jsonObject?.get("guild_id")?.jsonPrimitive?.content
                                        println("Wiadomość od $author: $content | DM(${guildId==null})")

                                        if (content != null && channelId != null && authorId != null) {
                                            val mention = "<@$botId>"
                                            val userMention = "<@$authorId>"
                                            when {
                                                (content == "!hello") -> {
                                                    sendMessage(client, token, channelId, "$userMention, Cześć!")
                                                }

                                                (content == "!kategorie") -> {
                                                        sendMessage(
                                                            client,
                                                            token,
                                                            channelId,
                                                            "$userMention ${categories.joinToString(" | ", "[", "]")}"
                                                        )
                                                }

                                                content.startsWith("!produkty") -> {
                                                    val commandList = content.trim().split(" ")
                                                    val category = commandList.getOrNull(1)?.lowercase() ?: "brak"
                                                    val response = products[category]?.joinToString(" | ", "[", "]") ?: "Nie znaleziono kategorii"
                                                        sendMessage(
                                                            client,
                                                            token,
                                                            channelId,
                                                            "$userMention Produkty w $category: $response"
                                                        )
                                                }
                                                // Chat
                                                content.startsWith("!chat") -> {
                                                    val chatMessageContent = content.removePrefix("!chat").trim()
                                                    if (chatMessageContent.isBlank()) {
                                                        sendMessage(client, token, channelId, "$userMention Proszę podać wiadomość po `chat` po mentionie.")
                                                    }
                                                    handleChatCommand(client, token, channelId, authorId, chatMessageContent, userChatHistories)
                                                }

                                                content.startsWith(mention) -> {
                                                    val command = content.removePrefix(mention).trim().split(" ").filter{ it.isNotEmpty() }
                                                    val arg = command.getOrNull(1) ?: ""
                                                    if (command.isNotEmpty()) {
                                                        when {
                                                            // Przywitanie
                                                            command[0] == "hello" -> sendMessage(
                                                                client,
                                                                token,
                                                                channelId,
                                                                "$userMention, Cześć!"
                                                            )
                                                            // Kategorie
                                                            command[0] == "kategorie" -> sendMessage(
                                                                client,
                                                                token,
                                                                channelId,
                                                                "$userMention ${
                                                                    categories.joinToString(
                                                                        " | ",
                                                                        "[",
                                                                        "]"
                                                                    )
                                                                }"
                                                            )
                                                            // Produkty
                                                            command[0] == "produkty" -> {
                                                                val response =
                                                                    products[arg]?.joinToString(" | ", "[", "]")
                                                                        ?: "Nie znaleziono kategorii"
                                                                sendMessage(
                                                                    client,
                                                                    token,
                                                                    channelId,
                                                                    "$userMention Produkty w ${arg}: $response"
                                                                )
                                                            }
                                                            // Chat
                                                            command[0] == "chat" -> {
                                                                if (arg.isBlank()) {
                                                                    sendMessage(client, token, channelId, "$userMention Proszę podać wiadomość po `chat` po mentionie.")
                                                                }
                                                                handleChatCommand(client, token, channelId, authorId, arg, userChatHistories)
                                                            }
                                                            else -> sendMessage(
                                                                client,
                                                                token,
                                                                channelId,
                                                                "$userMention Dozwolone komendy to: [!hello | !kategorie | !produkty {kateogria} | !chat {wiadomość}]"
                                                            )
                                                        }
                                                    }else{
                                                        sendMessage(
                                                            client,
                                                            token,
                                                            channelId,
                                                            "$userMention Dozwolone komendy to: [!hello | !kategorie | !produkty {kateogria}]"
                                                        )
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }

                            else -> { // Inne zdarzenia
                                lastSequence = json["s"]?.jsonPrimitive?.int
                                println("Zaktualizowano sekwencję: $lastSequence")
                            }
                        }
                    }
                }

            }
        } catch (e: Exception){
            println("Błąd połączenia: ${e.message}")
            delay(2000) // Przed ponowną próba połączenia

            if (e.message?.contains("Connection refused") == true) {
                println("Krytyczny błąd, kończę program")
                running = false
            }
        }
    }

    client.close()
}

suspend fun getGatewayUrl(client: HttpClient): String {
    val response: HttpResponse = client.get("https://discord.com/api/v10/gateway")
    val jsonElement = Json.parseToJsonElement(response.bodyAsText())
    return jsonElement.jsonObject["url"]?.jsonPrimitive?.content ?: throw IllegalStateException("Brak URL w odpowiedzi")
}

suspend fun handleChatCommand(
    client: HttpClient,
    token: String,
    channelId: String,
    authorId: String,
    messageContent: String,
    userChatHistories: MutableMap<String, MutableList<ChatMessage>>,
) {
    val userMention = "<@$authorId>"
    val currentChatHistory = userChatHistories.getOrPut(authorId) { mutableListOf() }

    // Dodaj wiadomość użytkownika do historii
    val userMsg = ChatMessage("user", messageContent)
    currentChatHistory.add(userMsg)

    // Obcinanie historii, jeśli jest za długa
    while (currentChatHistory.size > MAX_CHAT_HISTORY_PAIRS * 2) {
        // Usuń najstarsze widomości, ale w historii zawsze użytkownika musi być pierwsza
        currentChatHistory.removeAt(0)
        currentChatHistory.removeAt(0)
    }
    // Utwórz obiekt ChatRequest
    val chatRequest = ChatRequest(
        user_message = userMsg.content,
        chat_history = currentChatHistory
    )
    try {
        // Wyślij żądanie do serwisu FastAPI
        val chatResponse = client.post(CHAT_ENDPOINT) {
            contentType(ContentType.Application.Json)
            setBody(chatRequest)
        }.body<ChatResponse>()

        // Dodaj odpowiedź bota do historii
        val botResponseMsg = ChatMessage("assistant", chatResponse.response)
        currentChatHistory.add(botResponseMsg)

        // Wyślij odpowiedź bota na Discorda
        sendMessage(client, token, channelId, "$userMention ${chatResponse.response}")
    } catch (e: Exception) {
        println("Błąd podczas komunikacji z serwisem czatu: ${e.message}")
        e.printStackTrace() // To pomoże w debugowaniu
        sendMessage(client, token, channelId, "$userMention Przepraszam, coś poszło nie tak z usługą czatu.")
    }
}

suspend fun sendMessage(client: HttpClient, token: String, channelId: String, message: String) {
    try {
        val response = client.post("https://discord.com/api/v10/channels/$channelId/messages") {
            header("Authorization", "Bot $token")
            header("Content-Type", "application/json")
            setBody(Json.encodeToString(mapOf("content" to message)))
        }
        println("Wysłano odpowiedź: $message")
    } catch (e: Exception) {
        println("Błąd wysyłania wiadomości: ${e.message}")
    }
}