FROM golang:1.14-alpine as builder

WORKDIR /go/src
COPY ./src/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o src
RUN ls

FROM scratch
COPY --from=builder  /go/src .

CMD [ "./src" ]