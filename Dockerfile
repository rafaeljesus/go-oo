FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o /app/main --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
EXPOSE 8080
CMD ["/app/main"]
