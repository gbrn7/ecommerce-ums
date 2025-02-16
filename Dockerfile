FROM golang:1.23

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o ecommerce-ums

RUN chmod +x ecommerce-ums

EXPOSE 9000

CMD [ "./ecommerce-ums" ]