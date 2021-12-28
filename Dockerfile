FROM golang:1.17.5-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy 
RUN go build -o main
CMD ["/app/main"]