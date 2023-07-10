# syntax=docker/dockerfile:1.4

FROM python:3-slim

ARG version="main"

# install deps
RUN apt-get update \
    && apt-get install -y iproute2 git dumb-init curl unzip \
    && apt-get clean

RUN pip install --upgrade pip

# Add ExaBGP
RUN curl -O -L https://github.com/Exa-Networks/exabgp/archive/main.zip
RUN unzip main.zip

COPY exabgp.env exabgp-main/etc/exabgp

RUN useradd -r exa \
    && mkdir /etc/exabgp \
    && mkfifo /run/exabgp.in \
    && mkfifo /run/exabgp.out \
    && chown exa /run/exabgp.in \
    && chown exa /run/exabgp.out \
    && chmod 600 /run/exabgp.in \
    && chmod 600 /run/exabgp.out

ENV PATH /exabgp-main/bin:/exabgp-main/sbin:$PATH

EXPOSE 5555

CMD exabgp /etc/exabgp/exabgp.conf
