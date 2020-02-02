FROM golang:1.13-alpine
WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o auth

FROM alpine
WORKDIR /app
COPY --from=0 /app/auth .
CMD [ "/app/auth" ]
EXPOSE 8080
