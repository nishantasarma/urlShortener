FROM golang:1.17 AS BUILDER
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go 

FROM alpine:latest
WORKDIR /root/
COPY --from=BUILDER /app/main .
EXPOSE 8080
CMD ["./main"] 