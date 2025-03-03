FROM golang:1.24-alpine3.21 AS build
WORKDIR /src

COPY . .
RUN go mod download -x && CGO_ENABLED=0 go build -ldflags="-s -w" -o differ .

FROM alpine:3.21.3 

RUN addgroup -g 1000 app && \
    adduser -u 1000 -h /app -G app -S app
WORKDIR /app
USER app

COPY --from=build /src/differ .

CMD ["./differ"] 
