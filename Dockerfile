FROM golang:1.17-alpine

WORKDIR /home

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN mkdir views

COPY *.go ./

COPY views ./views

#RUN go build -o website
RUN GOOS=linux GOARCH=amd64 go build -o website
RUN chmod +x website

EXPOSE 3000

CMD [ "./website" ]

