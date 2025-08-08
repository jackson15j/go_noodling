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


FROM build AS test
WORKDIR /build/reports/coverage
WORKDIR /src
RUN go test -v -coverprofile=/build/reports/coverage/coverage.out -cover ./...


FROM scratch AS app
COPY --from=build /app /app
# See: `use_openapi_generated_server/main.go::20`
EXPOSE 8087
CMD ["/app"]
