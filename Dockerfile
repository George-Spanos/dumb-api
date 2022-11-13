FROM golang:alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY static static
RUN ls
RUN go build
EXPOSE 8080
CMD [ "./learn-go-web" ]