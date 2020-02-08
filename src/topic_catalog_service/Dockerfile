FROM golang:1.13-alpine
WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o topic_catalog

FROM alpine
WORKDIR /app
COPY --from=0 /app/topic_catalog .
CMD [ "/app/topic_catalog" ]
