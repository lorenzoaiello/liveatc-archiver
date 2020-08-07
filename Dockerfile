FROM scratch
COPY liveatc-archiver /go/bin/liveatc-archiver
ENV USER=docker
ENTRYPOINT [ "/go/bin/liveatc-archiver" ]