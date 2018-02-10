FROM golang:1.8.0-alpine as build
COPY hello.go src/
RUN go build src/hello.go

FROM scratch
COPY --from=build /go/hello .
CMD ["./hello"]
