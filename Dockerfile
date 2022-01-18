# Build stage
FROM golang:1.17.5 AS build-env

ARG GOARCH="amd64"

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOARCH=${GOARCH}

WORKDIR /app/kaznet-status
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o kaznet-status
RUN chmod +x kaznet-status

# Release stage
FROM alpine:latest AS release
WORKDIR /app/
COPY --from=build-env /app/kaznet-status/kaznet-status .
COPY --from=build-env /app/kaznet-status/templates templates
RUN mkdir database
COPY --from=build-env /app/kaznet-status/database/GeoLite2-City.mmdb /app/database/GeoLite2-City.mmdb
ENV WORKDIR "/app/"
ENV PATH "${WORKDIR}:${PATH}"
EXPOSE 2112

CMD ["kaznet-status"]
