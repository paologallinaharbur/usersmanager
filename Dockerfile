
FROM golang:1.15 as builder
COPY . /go/src/github.com/paologallinaharbur/usersmanager
ENV GOPATH=/go
ENV GO111MODULE=on
WORKDIR /go/src/github.com/paologallinaharbur/usersmanager
RUN CGO_ENABLED=0 GOOS=linux go build -o /userManager /go/src/github.com/paologallinaharbur/usersmanager/cmd/user-manager-server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /userManager .
COPY --from=0 /go/src/github.com/paologallinaharbur/usersmanager/swagger-ui ./swagger-ui
EXPOSE 35307/tcp

CMD ["/userManager", "--port", "35307", "--host", "0.0.0.0"]