FROM golang:1.19-alpine AS builder
MAINTAINER EAS Barbosa <easbarba@outlook.com>
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux GOARCH=amd64
RUN go build -race -ldflags "-extldflags '-static'" -o ./apito .

# FINAL STAGE
FROM golang:1.19-alpine
COPY --from=builder /app/apito /opt/apito
EXPOSE 8080
ENTRYPOINT [ "/opt/apito" ]
