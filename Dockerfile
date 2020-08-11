FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GAILS_ENV=production

WORKDIR /gails

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o gails .

# 运行时所需要的小iamge
FROM debian:stretch-slim
ENV GAILS_ENV=production
COPY ./wait-for.sh /
COPY ./config /config
COPY ./app/views /app/views

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /gails/gails /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 需要运行的命令
ENTRYPOINT ["/gails", "config/app_prod.ini"]
