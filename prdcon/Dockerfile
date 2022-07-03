FROM golang:1.18-alpine3.16

WORKDIR /prdcon

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

COPY ./pb ./pb

COPY ./taskman/main.go ./

RUN go get prdcon
RUN go install
RUN go build -o /taskman


EXPOSE 50051

CMD [ "/taskman" ]