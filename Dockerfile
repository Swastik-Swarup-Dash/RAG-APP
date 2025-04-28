
FROM golang:1.21-alpine AS builder


RUN apk add --no-cache git ca-certificates


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o rag-app ./cmd/server/main.go


FROM alpine:latest


RUN apk --no-cache add ca-certificates


WORKDIR /root/


COPY --from=builder /app/rag-app .




EXPOSE 8080


 ENV GEMINI_API_KEY="AIzaSyBn1g_Hg0pXc-OqNvSPoPO7iQHs30ukE_c"
 ENV DB_CONN="postgres://postgres:ishan1234@localhost:5432/APIDATABASE"

CMD ["./rag-app"]
