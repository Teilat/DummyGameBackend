FROM golang:1.18-alpine

WORKDIR /app

COPY config.yaml ./
COPY .build/linux/* ./

EXPOSE 8080
CMD ["./dummygamebackend"]