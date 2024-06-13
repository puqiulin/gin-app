# My gin app

> Just for study and fun 🤯

## Tech Stack

### Frontend

- [react](https://react.dev/)
- [next](https://nextjs.org/)
- [tailwind](https://tailwindcss.com/)

### Backend

- [gin](https://github.com/gin-gonic/gin)
- [postgres](https://www.postgresql.org/)
- [redis](https://github.com/redis/redis)
- [uptrace](https://uptrace.dev/)
- [grpc](https://grpc.io/)

### Devops

- [docker](https://www.docker.com/)
- [etcd](https://etcd.io/)
- [kafka](https://kafka.apache.org/)
- [grafana](https://grafana.com/)
- [loki](https://grafana.com/oss/loki/)

## Run

```shell
# run frontend
# http://127.0.0.1:3001
make run-frontend

# run backend
# http://127.0.0.1:9999
make run-backend

# run as docker
make run-docker-deps
make run-docker
```
