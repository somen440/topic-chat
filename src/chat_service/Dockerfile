FROM golang:1.13-alpine
WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o chat

FROM alpine
WORKDIR /app
COPY --from=0 /app/chat .
CMD [ "/app/chat" ]
