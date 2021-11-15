FROM golang:alpine as builder
WORKDIR /
copy . ./
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go build -o app main.go

FROM alpine
WORKDIR /
COPY --from=builder /app /app
ENTRYPOINT /app
EXPOSE 8080