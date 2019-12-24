FROM golang AS build-env

MAINTAINER yanorei32

COPY ./src /work/

WORKDIR /work/

RUN go test \
		-bench . \
		-benchtime 5s \
		-cpu 1

