FROM ubuntu:latest

RUN apt-get update && apt-get install -y openjdk-21-jdk curl unzip zip ngrok

# Instalacja SDKMAN!
RUN curl -s "https://get.sdkman.io" | bash
ENV SDKMAN_DIR="/root/.sdkman"

RUN bash -c "source /root/.sdkman/bin/sdkman-init.sh && sdk install sbt"

WORKDIR /app
COPY . /app

EXPOSE 9000

CMD ["sbt", "run"]