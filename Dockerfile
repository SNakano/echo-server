FROM golang:latest AS build

COPY . /src
WORKDIR /src
ENV CGO_ENABLED 0
RUN go build -o echo-server

FROM scratch
COPY --from=build /src/echo-server /
ENTRYPOINT [ "/echo-server" ]
