FROM debian:stable

EXPOSE 179

RUN apt-get update && apt-get install -y \
    autoconf \
	bison \
	build-essential \
	curl \
	flex \
	libreadline-dev \
	libncurses5-dev \
	m4 \
	unzip

RUN curl -O -L https://github.com/CZ-NIC/bird/archive/master.zip
RUN unzip master.zip

# This directory must be mounted as a local volume with '-v `pwd`/bird:/etc/bird:rw' docker's command line option.
# The host's file at `pwd`/bird/bird.conf is used as the configuration file for BIRD.
RUN mkdir /etc/bird

RUN cd bird-master && \
    autoreconf && \
	./configure && \
	make && \
    make install

CMD bird -c /etc/bird/bird.conf -d
