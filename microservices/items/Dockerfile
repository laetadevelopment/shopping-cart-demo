FROM golang:1.14-stretch as builder
WORKDIR $GOPATH/src/items
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM golang:alpine
COPY --from=builder /app ./
ENTRYPOINT ["./app"]