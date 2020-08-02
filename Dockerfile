FROM golang:1.14.6-alpine3.12 as builder
RUN mkdir /build 
COPY *.go /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]