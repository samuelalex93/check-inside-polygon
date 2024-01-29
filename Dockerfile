FROM golang:1.13.8-alpine3.11 as build

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /check-inside-polygon
FROM scratch

COPY --from=build /check-inside-polygon /check-inside-polygon

CMD [ "/check-inside-polygon" ]