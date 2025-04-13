FROM ubuntu:24.04

RUN apt-get update && apt-get install -y \
    wget \
    curl \
    unzip \
    zip \
    openjdk-21-jdk

SHELL ["/bin/bash", "-c"]

RUN curl -s "https://get.sdkman.io" | bash
RUN source "$HOME/.sdkman/bin/sdkman-init.sh" \
    && sdk install kotlin \
    && sdk install gradle

ENV PATH=/root/.sdkman/candidates/kotlin/current/bin:$PATH
ENV PATH=/root/.sdkman/candidates/gradle/current/bin:$PATH

WORKDIR /home/kotlin-app

COPY . .

CMD [ "gradle", "run" ]