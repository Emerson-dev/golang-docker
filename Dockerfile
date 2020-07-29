FROM golang:1.11 as builder

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix "static" golang-docker

FROM scratch

USER 1000

COPY --from=builder /go/bin/golang-docker /bin/golang-docker

ENTRYPOINT ["/bin/golang-docker"]