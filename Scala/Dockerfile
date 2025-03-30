FROM ubuntu:24.04

RUN apt-get update \
    && apt-get install -y \
    wget \
    curl \
    unzip \
    zip \ 
    sqlite3 \ 
    default-jdk

# scala, sbt
RUN curl -fL https://github.com/VirtusLab/coursier-m1/releases/latest/download/cs-aarch64-pc-linux.gz | gzip -d > cs \
    && chmod +x cs \
    && ./cs setup -y

ENV PATH="$PATH:/root/.local/share/coursier/bin"

WORKDIR /home/scala-app

COPY . .

EXPOSE 9000

CMD ["sbt", "run"]