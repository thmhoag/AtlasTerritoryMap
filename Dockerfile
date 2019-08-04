FROM golang:1.12.6-stretch as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags '-s' -o output/bin/AtlasTerritoryMap

FROM alpine:3.9 as final

WORKDIR /app
COPY --from=builder /app/output/bin/AtlasTerritoryMap /usr/local/bin/AtlasTerritoryMap

VOLUME [ "/config" ]
EXPOSE 8881

ENTRYPOINT [ "AtlasTerritoryMap" ]
