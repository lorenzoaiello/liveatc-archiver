FROM golang:alpine as build
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY liveatc-archiver /go/bin/liveatc-archiver
ENV USER=docker
ENTRYPOINT [ "/go/bin/liveatc-archiver" ]