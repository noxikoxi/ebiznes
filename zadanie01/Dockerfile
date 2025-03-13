FROM ubuntu:24.04

# software-properties-common -> dla pythona 3.10, bo domyślnie ubuntu 24.04 ma pythona 3.12
RUN apt-get update && apt-get install -y software-properties-common

RUN add-apt-repository ppa:deadsnakes/ppa -y

# unzip zip -> dla sdkman
RUN apt-get update && apt-get install -y python3.10 openjdk-8-jdk curl unzip zip

# Instalacja SDKMAN! -> do kotlina i gradle
RUN curl -s "https://get.sdkman.io" | bash
ENV SDKMAN_DIR="/root/.sdkman"

# SDKMAN! wymaga podania source
RUN bash -c "source /root/.sdkman/bin/sdkman-init.sh && sdk install gradle"
RUN bash -c "source /root/.sdkman/bin/sdkman-init.sh && sdk install kotlin"

ENV PATH="${SDKMAN_DIR}/candidates/gradle/current/bin:${SDKMAN_DIR}/candidates/kotlin/current/bin:${PATH}"

# python 3.10 jest teraz domyślny
RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.10 2
RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.12 1

WORKDIR /app

COPY build.gradle /app
COPY Main.java /app/src/main/java/

RUN gradle build --no-daemon --quiet

CMD ["gradle", "run", "--no-daemon", "--quiet"]