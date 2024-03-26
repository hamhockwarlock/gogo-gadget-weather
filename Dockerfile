FROM golang:1.22.1-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gogo-gadget-weather

ENV PORT=4242
EXPOSE 4242

CMD ["/gogo-gadget-weather"]
