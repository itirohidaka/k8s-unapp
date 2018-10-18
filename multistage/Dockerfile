FROM golang:alpine as build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM alpine
COPY --from=build /app/main /
WORKDIR /
EXPOSE 8080
CMD ["./main"]
