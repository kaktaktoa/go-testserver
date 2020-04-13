FROM golang:latest

RUN mkdir ./app
ADD . ./app
WORkDIR ./app
RUN go build -o main
CMD ["/app/main"]