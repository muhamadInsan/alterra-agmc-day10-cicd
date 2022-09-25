FROM golang:1.19-alpine
WORKDIR /app
COPY ./tugas-crud-dinamis .
COPY . .
# RUN go mod download
# COPY *.go ./
RUN go build -o crud-dinamis
EXPOSE 8080
# CMD ./crud-dinamis