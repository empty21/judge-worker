FROM ubuntu:20.04

MAINTAINER Thieu Quan Ngoc
WORKDIR /judge
RUN dpkg --add-architecture i386
RUN apt-get update
RUN apt-get install curl -y
RUN curl -sL https://deb.nodesource.com/setup_14.x | /bin/bash -
RUN apt-get update
RUN apt-get install -y openjdk-8-jdk
RUN apt-get install -y python3
RUN apt-get install -y build-essential
RUN apt-get install -y fpc
RUN apt-get install -y nodejs
RUN apt-get install -y time
