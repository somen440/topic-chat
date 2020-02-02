FROM golang:1.13-alpine
WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o frontend

FROM alpine
WORKDIR /app
COPY --from=0 /app/frontend .
ADD templates templates
CMD [ "/app/frontend" ]
EXPOSE 8080
