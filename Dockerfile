FROM golang AS gobuilder
ENV GO111MODULE="on"
WORKDIR /go/src/github.com/jmhobbs/echo-echo
ADD go.* /go/src/github.com/jmhobbs/echo-echo/
RUN go mod download
ADD *.go /go/src/github.com/jmhobbs/echo-echo/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o echo-echo .

FROM scratch
COPY --from=gobuilder /go/src/github.com/jmhobbs/echo-echo/echo-echo /echo-echo
CMD ["/echo-echo"]

