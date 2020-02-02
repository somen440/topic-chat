FROM golang:1.13-alpine
WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o topic_catalog

FROM alpine
WORKDIR /app
COPY --from=0 /app/topic_catalog .
ADD topic_catalog.yaml topic_catalog.yaml
CMD [ "/app/topic_catalog" ]
