# Ebiznes

### Zadanie 01 Docker

:white_check_mark: 3.0 obraz ubuntu z Pythonem 3.10

:white_check_mark: 3.5 obraz ubuntu 24.04 z Javą 8 i Kotlinem

:white_check_mark: 4.0 najnowszy Gradle z JDBC SQLite (build.gradle)

:white_check_mark: 4.5 program typu Hello World uruchamiany przez gradle

:white_check_mark: 5.0 konfiguracja docker-compose

Obraz -> [dockerhub](https://hub.docker.com/r/noxikoxi/my-java-app)

Link do commita -> https://github.com/noxikoxi/ebiznes/commit/f2ea32575d91c781623d77f38e84f89201dac896

Kod -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie01

Demo -> https://github.com/noxikoxi/ebiznes/blob/main/demos/zadanie01.zip


### Zadanie 02 Scala + Play

:white_check_mark: 3.0 kontroler do Produktów

:white_check_mark: 3.5 endpointy zgodne z CRUD do Produktów

:white_check_mark: 4.0 kontrolery do Kategorii i Koszyka + CRUD 

:white_check_mark: 4.5 obraz dockera z aplikacja + skrypt uruchamiający via ngrok

:white_check_mark: 5.0 konfiguracja CORS dla dwóch hostów dla metod CRUD

Link do commita -> https://github.com/noxikoxi/ebiznes/commit/08908c405dd7f8b9997bfab6e0ee59edb6f7533e

Kod -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie02

Demo -> https://github.com/noxikoxi/ebiznes/blob/main/demos/zadanie02.zip

### Zadanie 03 Kotlin + Ktor + Discord Bot

:white_check_mark: 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord.

:white_check_mark: 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota).

:white_check_mark: 4.0 Zwróci listę kategorii na określone żądanie użytkownika.

:white_check_mark: 4.5 Zwróci listę produktów wg żądanej kategorii.

:x: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger, Webex

Link do commita -> https://github.com/noxikoxi/ebiznes/commit/ce4ceba7498863f20dfe50db682dc9eae234841c

Kod -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie03

Demo -> https://github.com/noxikoxi/ebiznes/blob/main/demos/zadanie03-kotlin.zip

### Zadanie 04 Go + Echo + GORM

:white_check_mark: 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD

:white_check_mark: 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)

:white_check_mark: 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint

:white_check_mark: 4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem

:white_check_mark: 5.0 pogrupować zapytania w gorm’owe scope'y

Link do commita -> https://github.com/noxikoxi/ebiznes/commit/0fc59dfb8fa68fa932d472c97722809c7ddb43d4

Kod -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie04

Demo -> https://github.com/noxikoxi/ebiznes/blob/main/demos/zadanie04-go.zip

### Zadanie 05 Frontend, React + Go

:white_check_mark: 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach z aplikacji serwerowej.

:white_check_mark: 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing.

:white_check_mark: 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za pomocą React hook.

:white_check_mark: 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz kliencką na dockerze via docker-compose.

:white_check_mark: 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS.

Link do commita -> https://github.com/noxikoxi/ebiznes/commit/94dda5219c9774e65ae9f7ca29085efbdc991d5f

Kod -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie05

Demo -> https://github.com/noxikoxi/ebiznes/blob/main/demos/zadanie05-frontend.zip

### Zadanie 06 Selenium, testy

:white_check_mark: 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium(Kotlin, Python, Java, JS, Go, Scala)

:white_check_mark: 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50 asercji

:white_check_mark: 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego projektu z minimum 50 asercjami

:white_check_mark: 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z minimum jednym scenariuszem negatywnym per endpoint

:white_check_mark: 5.0 Należy uruchomić testy funkcjonalne na Browserstacku

Link do readme odnośnie zadania -> https://github.com/noxikoxi/ebiznes/blob/main/zadanie06/readme.md

Kod -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie06


### Zadanie 07 Sonar

:white_check_mark: 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w hookach gita

:white_check_mark: 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod aplikacji serwerowej)

:white_check_mark: 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod aplikacji serwerowej)

:white_check_mark: 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa w kodzie w Sonarze (kod aplikacji serwerowej)

:white_check_mark: 5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie aplikacji klienckiej

Link do repozytorium analizowanego przez Sonar -> https://github.com/noxikoxi/products

Link do commita -> https://github.com/noxikoxi/products/commit/31f2e0867d63cc4ea0e2d1f2a97596af14ae4933

Kod hooka, zrzuty ekrany pokazujące jego działanie i badge w readme -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie07

### Zadanie 08 OAuth2

:white_check_mark: 3.0 logowanie przez aplikację serwerową (bez Oauth2)

:white_check_mark: 3.5 rejestracja przez aplikację serwerową (bez Oauth2)

:white_check_mark: 4.0 logowanie via Google OAuth2

:white_check_mark: 4.5 logowanie via Facebook lub Github OAuth2

:white_check_mark: 5.0 zapisywanie danych logowania OAuth2 po stronie serwera

Link do commita -> https://github.com/noxikoxi/ebiznes/66ac7cc17a4573ed5f37e87b80dbc74570de3f5d

Kod -> https://github.com/noxikoxi/ebiznes/tree/main/zadanie08

Demo -> https://github.com/noxikoxi/ebiznes/blob/main/demos/zadanie08-Oauth2.zip