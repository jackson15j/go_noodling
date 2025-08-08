# https://hub.docker.com/_/golang
FROM golang:1.24-alpine AS base
WORKDIR /src
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .


FROM base AS build
RUN go generate ./...
RUN go build -v -o /app .

FROM golang:1.24-alpine AS build-debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /src
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go generate ./...
RUN go build -gcflags="all=-N -l" -v -o /app .
# RUN CGO_ENABLED=0 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv
# RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./app
EXPOSE 8087 40000
WORKDIR /
# https://github.com/go-delve/delve/blob/master/Documentation/faq.md#-how-do-i-use-delve-with-docker
CMD ["dlv", "exec", "--headless", "--continue", "--listen=:40000", "--accept-multiclient", "--api-version=2", "--log", "--log-output=debugger,dap,rpc", "/app"]


FROM build AS test
WORKDIR /build/reports/coverage
WORKDIR /src
RUN go test -v -coverprofile=/build/reports/coverage/coverage.out -cover ./...

FROM scratch AS app-debug
WORKDIR /
# https://blog.jetbrains.com/go/2020/05/06/debugging-a-go-application-inside-a-docker-container/
COPY --from=build-debug /go/bin/dlv /dlv
COPY --from=build-debug /app /app
EXPOSE 8087 40000
# CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "--log=true", "--log-output=debugger,debuglineerr,gdbwire,lldbout,rpc", "exec", "/app"]
ENTRYPOINT ["/dlv"]

FROM scratch AS app
COPY --from=build /app /app
# See: `use_openapi_generated_server/main.go::20`
EXPOSE 8087
CMD ["/app"]
