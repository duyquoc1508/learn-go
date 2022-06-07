FROM golang:alpine

WORKDIR /app

COPY . .

CMD ["./go-demo.exe"]
