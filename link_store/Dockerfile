# https://hub.docker.com/_/golang
FROM golang:alpine AS base
WORKDIR /src
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .


FROM base AS build
RUN go build -v -o /app .


FROM scratch AS app
COPY --from=build /app /app
COPY index.html .
EXPOSE 8080
CMD ["/app"]
