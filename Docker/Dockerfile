FROM ubuntu:24.04

RUN apt-get update \
    && apt-get install -y wget curl zip

RUN apt install software-properties-common -y

RUN add-apt-repository ppa:deadsnakes/ppa \
    && apt-get update \
    && apt-get install -y python3.10 python3.10-dev