FROM golang:1.19

RUN mkdir -p /go/src/github.com/Dionizio8/go-task  
ADD . /go/src/github.com/Dionizio8/go-task/
WORKDIR /go/src/github.com/Dionizio8/go-task
RUN go build -o worker cmd/worker/main.go

CMD ["/go/src/github.com/Dionizio8/go-task/worker"]