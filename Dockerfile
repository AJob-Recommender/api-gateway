# build stage
FROM golang:alpine AS build-env
WORKDIR src/baseapi

COPY go.mod go.sum ./

COPY . ./
RUN go build -o bin/baseapi -v ./cmd/api/

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/baseapi/bin /app/
COPY --from=build-env /go/src/baseapi/config /app/config
USER 1001
EXPOSE 8080
ENTRYPOINT ["./baseapi"]