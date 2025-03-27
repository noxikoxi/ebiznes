FROM ubuntu:latest

RUN apt-get update && apt-get install -y openjdk-21-jdk curl

# Skopiowanie aplikacji Play Framework (po sbt stage)
COPY src/target/universal/stage/ /app

EXPOSE 9000

# Domyślnie na porcie 9000
CMD ["./app/bin/play-scala-api"]