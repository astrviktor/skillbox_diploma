# Собираем в go
FROM golang:1.17.6 as build

ENV BIN_FILE /opt/statuspage/statuspage-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/main.go

# На выходе тонкий образ
FROM alpine:3.9
#FROM debian:stable-20211220

LABEL ORGANIZATION="home"
LABEL SERVICE="statuspage"
LABEL MAINTAINERS="astrviktor@gmail.com"

ENV BIN_FILE "/opt/statuspage/statuspage-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

#ENV CONFIG_FILE /etc/statuspage/config.yaml
#COPY ./configs/config.yaml ${CONFIG_FILE}
COPY ./cmd/config.yaml /opt/statuspage
COPY ./cmd/data /opt/statuspage/data
COPY ./web /opt/statuspage/web

WORKDIR /opt/statuspage/

EXPOSE 9999
CMD ${BIN_FILE}