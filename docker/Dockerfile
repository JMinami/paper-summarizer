# ビルドステージ
FROM golang:alpine AS build

WORKDIR /app
RUN apk add --no-cache git

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src/ .
RUN go build -o web_server ./cmd/prod/main.go

# 最終ステージ
FROM alpine

COPY ./secret/ /secret/
ARG GOOGLE_APPLICATION_CREDENTIALS
ARG LINE_MESSAGING_API_CHANNEL_ACCESS_TOKEN
ARG LINE_MESSAGING_API_CHANNEL_SECRET
ENV GOOGLE_APPLICATION_CREDENTIALS=/secret/google_application_credentials.json
ENV LINE_MESSAGING_API_CHANNEL_ACCESS_TOKEN=${LINE_MESSAGING_API_CHANNEL_ACCESS_TOKEN}
ENV LINE_MESSAGING_API_CHANNEL_SECRET=${LINE_MESSAGING_API_CHANNEL_SECRET}

RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /app/web_server /app/

# ユーザーとグループの追加
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app
USER appuser

EXPOSE 8080

CMD ["./web_server"]
