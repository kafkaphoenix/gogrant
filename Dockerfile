FROM golang:1.24.5-alpine3.22 AS builder

# environment variables
# CGO_ENABLED=0 disable cgo for static binary & portability
# GOOS and GOARCH set the target platform
# GO111MODULE=on enables Go modules
ENV CGO_ENABLED=0 \
    GOOS=linux    \
    GOARCH=amd64  \
    GO111MODULE=on

WORKDIR /app

# copy source code
COPY . .

# download dependencies with cache mount for better performance
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && \
    go mod verify

# -ldflags optimize binary size
# -trimpath remove absolute path from binary
RUN go build -ldflags="-s -w -extldflags '-static'" -trimpath -o app ./cmd/app/main.go

# final image
FROM alpine:3.21

LABEL maintainer="Javier Aguilera"

ARG USER=app
ARG UID=10001
ARG GID=10001
ENV HOME=/home/${USER}

RUN apk update && apk upgrade && \
    # ca-certificates is needed for SSL/TLS connections and libcap for capabilities
    apk add --no-cache ca-certificates=20250619-r0 libcap=2.71-r0 && \
    rm -rf /var/cache/apk/* && \
    # create non-root user and group
    addgroup -g ${GID} ${USER} && \
    adduser -D -g "" -h "${HOME}" -s "/sbin/nologin" -G ${USER} -u ${UID} ${USER} && \
    # setup app directory and permissions
    mkdir -p /app && \
    chown -R ${UID}:${GID} /app

WORKDIR /app

# copy binaries and config files
COPY --from=builder --chown=${UID}:${GID} /app/app .

# drop root privileges
USER ${USER}

CMD ["./app"]
