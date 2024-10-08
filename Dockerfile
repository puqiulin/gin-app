# This may be caused by the glibc compatibility library of Alpine.
# So just use Debian to save your life.

# build frontend
FROM node:latest AS frontend
WORKDIR /app

COPY web/package.json web/package-loock.lock* ./
RUN npm install --frozen-lockfile

COPY web/. .

RUN npm run build

# build backend
FROM golang:alpine AS backend
WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on

COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download && \
    go mod verify

COPY . .
RUN go build -ldflags="-s -w" -o app .

# build ultimate
FROM alpine:latest as ultimate
ENV TZ=Asia/Shanghai NODE_ENV=production
WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk update &&  \
#    https://stackoverflow.com/questions/66963068/docker-alpine-executable-binary-not-found-even-if-in-path/66974607#66974607
#    apk add gcompat &&  \
    apk add --no-cache ca-certificates tzdata npm && \
    npm install -g serve


COPY --from=frontend /app/.next ./.next
COPY --from=frontend /app/public ./public

USER nextjs

COPY --from=backend /app/app .

COPY .env .

EXPOSE 3001 8888 9999
CMD ["sh", "-c", "./app & npm start"]
