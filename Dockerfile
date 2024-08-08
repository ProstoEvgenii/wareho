FROM golang:alpine as BUILD
WORKDIR /app
COPY . .
RUN go build -o warehouse

FROM alpine
COPY --from=BUILD /app/warehouse .
CMD ["./warehouse"]
