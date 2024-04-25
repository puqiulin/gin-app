# build frontend
FROM node:latest as frontend
ARG API_HOST
WORKDIR /web
COPY web/ ./
RUN yarn && yarn build

# build backend
FROM golang:bookworm AS backend
ARG API_HOST
WORKDIR /app
COPY . .
COPY --from=frontend /web/dist ./web/dist
ENV GOPROXY=https://goproxy.cn
ENV TZ=Asia/Shanghai
RUN go mod tidy
RUN go build -tags=jsoniter -o /appw
EXPOSE 8888 9999

WORKDIR /app/paopao-ce
COPY --from=backend /app/app .
COPY --from=backend /app/config.yaml.sample config.yaml

VOLUME ["/app/paopao-ce/custom"]
EXPOSE 8888 9999
HEALTHCHECK --interval=5s --timeout=5s  --retries=3  CMD ps -ef | grep app || exit 1
ENTRYPOINT ["/app/paopao"]
CMD ["serve"]