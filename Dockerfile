FROM golang:alpine as golang
WORKDIR /usr/local/go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build src/hostinfo/hostinfo.go

FROM scratch
COPY --from=golang /usr/local/go/hostinfo /
CMD ["/hostinfo"]
