FROM golang:1.17-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 7654
CMD ["./main"]
