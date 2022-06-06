FROM golang:1.16-alpine
WORKDIR /app
COPY . .
RUN go build -o ds-output .
EXPOSE 9090
ENTRYPOINT [ "./ds-output" ]
