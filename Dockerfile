FROM golang:1.20-alpine3.17 as build
WORKDIR /app

COPY . .
RUN go build -o main main.go

FROM alpine:3.17 
WORKDIR /app

COPY --from=build /app/main .
COPY app.env .

EXPOSE 3000
CMD [ "/app/main" ]