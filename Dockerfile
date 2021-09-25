# syntax=docker/dockerfile:1

FROM golang:1.16 AS build

RUN apt-get update && apt-get install -y \
    dos2unix

WORKDIR /opt/app


COPY handlers handlers
COPY templates templates
COPY go.mod ./

# Install dependencies
RUN go mod download

# copy across all root level .go files
COPY *.go ./

RUN find . -type f -print0 | xargs -0 dos2unix --

# Build Image
RUN go build -o main

FROM golang:1.16 AS production

RUN useradd -m -d /opt/app --uid 1000 -s /bin/bash app

WORKDIR /opt/app

COPY --from=build /opt/app/main /opt/app/main

EXPOSE 8080

CMD ["/opt/app/main"]