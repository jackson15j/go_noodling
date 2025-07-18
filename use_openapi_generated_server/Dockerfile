# https://hub.docker.com/_/golang
FROM golang:1.24-alpine AS base
WORKDIR /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .


FROM base AS build
RUN go generate ./...
RUN go build -v -o /usr/local/bin/app .


FROM build AS test
RUN go test -cover ./...


FROM scratch AS app
COPY --from=build /usr/local/bin/app /bin/
# See: `use_openapi_generated_server/main.go::20`
EXPOSE 8083
CMD ["/bin/app"]
