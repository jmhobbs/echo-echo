FROM golang AS gobuilder
WORKDIR /go/src/github.com/jmhobbs/echo-echo
ADD main.go /go/src/github.com/jmhobbs/echo-echo/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o echo-echo .

FROM scratch
COPY --from=gobuilder /go/src/github.com/jmhobbs/echo-echo/echo-echo /echo-echo
CMD ["/echo-echo"]

