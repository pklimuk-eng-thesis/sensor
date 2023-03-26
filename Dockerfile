FROM golang:1.20.1-alpine AS builder

LABEL maintainer="Pavel Klimuk <pavelklimuk@outlook.com>"

WORKDIR /src 

COPY . .
RUN go mod download

RUN go build -o presenceSensor main.go


FROM scratch

WORKDIR /

COPY --from=builder /src/presenceSensor /presenceSensor

EXPOSE 8080

CMD ["/presenceSensor"]