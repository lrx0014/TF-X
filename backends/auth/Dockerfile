FROM golang:1.24.2 AS builder

COPY . /src
WORKDIR /src

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /app/configs

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./auth", "-conf", "/app/conf"]
