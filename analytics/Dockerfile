FROM golang:1.20

WORKDIR /wrk

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /analyzer

CMD [ "/analyzer" ]
