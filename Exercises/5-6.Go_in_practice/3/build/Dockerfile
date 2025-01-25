##
## Build
##
FROM golang:1.21-bullseye AS build

WORKDIR /app

COPY go.* ./
# RUN go mod tidy
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /api ./Exercises/5-6.Go_in_practice/3/cmd/api/main.go
RUN CGO_ENABLED=0 go build -o /process ./Exercises/5-6.Go_in_practice/3/cmd/process/main.go
RUN CGO_ENABLED=0 go build -o /rabbitwaiter ./Exercises/5-6.Go_in_practice/3/utils/rabbitwaiter/main.go
RUN CGO_ENABLED=0 go build -o /natswaiter ./Exercises/5-6.Go_in_practice/3/utils/natswaiter/main.go

##
## Deploy
##
FROM alpine:3.19

WORKDIR /

COPY --from=build /api /api
COPY --from=build /process /process
COPY --from=build /rabbitwaiter /rabbitwaiter
COPY --from=build /natswaiter /natswaiter
COPY --from=build app/Exercises/5-6.Go_in_practice/3/utils/rabbitwaiter/wait.sh /wait.sh
COPY --from=build app/Exercises/5-6.Go_in_practice/3/utils/natswaiter/natswait.sh /natswait.sh

EXPOSE 8080
