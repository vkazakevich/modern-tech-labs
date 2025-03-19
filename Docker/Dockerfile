FROM ubuntu:24.04

RUN apt-get update \
    && apt-get install -y wget curl unzip zip sqlite3

RUN apt install software-properties-common -y

RUN add-apt-repository ppa:deadsnakes/ppa \
    && apt-get update \
    && apt-get install -y python3.10 python3.10-dev

RUN apt-get install -y openjdk-8-jdk

SHELL ["/bin/bash", "-c"]

ENV HOME=/root

RUN curl -s "https://get.sdkman.io" | bash
RUN source "$HOME/.sdkman/bin/sdkman-init.sh" \
    && sdk install kotlin \
    && sdk install gradle

ENV PATH=$HOME/.sdkman/candidates/kotlin/current/bin:$PATH
ENV PATH=$HOME/.sdkman/candidates/gradle/current/bin:$PATH

COPY kotlin-app ./home/kotlin-app
WORKDIR /home/kotlin-app

CMD ["gradle", "run"]