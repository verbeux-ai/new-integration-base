FROM golang:1.24 as build


ARG GITHUB_TOKEN
RUN git config --global url."https://verbeuxdevops:${GITHUB_TOKEN}@github.com".insteadOf "https://github.com"
RUN go env -w GOPRIVATE=github.com/verbeux-ai/*

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 1

RUN apt update -y &&  apt-get install -y librdkafka-dev build-essential pkg-config git

RUN go build -a -installsuffix cgo -o app

FROM --platform=linux/amd64 debian

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

WORKDIR /app

COPY --from=build /app/app /app

ENTRYPOINT [ "/app/app" ]