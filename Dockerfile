FROM golang:1.12.3-alpine as build-env
WORKDIR /go/src/github.com/tuananh1998hust/go_api_todo
COPY . .
RUN chmod +x ./package.sh && \
    ./package.sh 

FROM alpine:3.10
WORKDIR /app
COPY --from=build-env /go/src/github.com/tuananh1998hust/go_api_todo/main ./
EXPOSE 8080
CMD ["./main"]