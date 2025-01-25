##
## Build
##
FROM golang:1.21-bullseye AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /metr ./Service/cmd/gometr/main.go

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /metr /metr
COPY --from=build /app/Service/configs/gometr.yaml /configs/gometr.yaml

EXPOSE 8000

ENTRYPOINT ["/metr"]