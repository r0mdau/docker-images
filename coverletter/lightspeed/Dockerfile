FROM golang:1.14 AS build
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -o /cover -a cover.go

FROM scratch
COPY --from=build /cover /cover
CMD ["/cover"]
