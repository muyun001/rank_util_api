FROM golang:1.13 as builder

ARG GOPROXY
ENV GORPOXY ${GOPROXY}

ADD . /builder

WORKDIR /builder

RUN go build main.go

FROM golang:1.13

COPY --from=builder /builder/main /app/rank-util-api

WORKDIR /app

CMD ["./rank-util-api"]

EXPOSE 8080