# build frontend
FROM node:latest as frontend
WORKDIR /web

COPY web/package.json .
RUN npm install

COPY web/ .
RUN npm run build

# build backend
FROM golang:bookworm AS backend
ENV TZ=Asia/Shanghai CGO_ENABLED=0 GOOS=linux GO111MODULE=on  GOPROXY=https://goproxy.cn
WORKDIR /app

COPY . .
RUN --mount=type=cache,target=/root/go/pkg/mod go mod download && go mod verify
RUN go build -o app

# build all
FROM alpine:latest as finall
WORKDIR /app
ENV NODE_ENV production

RUN apk --no-cache add ca-certificates

COPY --from=frontend /web/.next /app/web/.next
COPY --from=frontend /web/public /app/web/public
ENV NEXT_PUBLIC_DIR=/app/web/.next

COPY --from=backend /app/app /app/app

EXPOSE 8888 9999
CMD ["/app/app"]
