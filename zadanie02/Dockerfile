FROM ubuntu:latest

RUN apt-get update && apt-get install -y openjdk-21-jdk curl

# Instalacja ngrok
RUN curl -sSL https://ngrok-agent.s3.amazonaws.com/ngrok.asc \
  | tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null \
  && echo "deb https://ngrok-agent.s3.amazonaws.com buster main" \
  | tee /etc/apt/sources.list.d/ngrok.list \
  && apt update \
  && apt install ngrok

# Skopiowanie aplikacji Play Framework (po sbt stage)
COPY src/target/universal/stage/ /app

# Kopiowanie skrypty uruchamiającego aplikacje
COPY start.sh /home

EXPOSE 9000

CMD ["/home/start.sh"]