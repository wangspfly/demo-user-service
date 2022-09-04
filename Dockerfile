FROM golang AS build

ENV GOPROXY=https://goproxy.cn
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR $GOPATH/src/demo-user-service
COPY . $GOPATH/src/demo-user-service
RUN go build -gcflags="all=-N -l" .

FROM golang
EXPOSE 8080 8081

WORKDIR $GOPATH/bin
COPY --from=build $GOPATH/bin/dlv $GOPATH/bin
COPY --from=build $GOPATH/src/demo-user-service/demo-user-service $GOPATH/bin

CMD ["/go/bin/dlv", "--listen=:8081", "--headless=true", "--api-version=2", "--continue", "--accept-multiclient", "exec", "./demo-user-service"]
