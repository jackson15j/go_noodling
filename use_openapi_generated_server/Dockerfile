# https://hub.docker.com/_/golang
FROM golang:1.20-alpine as base
WORKDIR /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .


FROM base AS build
RUN go generate ./...
RUN go build -v -o /usr/local/bin/app .


FROM build as test
RUN go test -cover ./...


FROM scratch as app
COPY --from=build /usr/local/bin/app /bin/
# See: `use_openapi_generated_server/main.go::20`
EXPOSE 8083
CMD ["/bin/app"]
