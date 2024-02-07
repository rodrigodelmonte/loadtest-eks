FROM golang:1.21-alpine@sha256:dd8888bb7f1b0b05e1e625aa29483f50f38a9b64073a4db00b04076cec52b71c as builder
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /bin/loadtest

# Build final image containing only the loadtest binary
FROM scratch
COPY --from=builder /bin/loadtest /bin/loadtest
ENV FREQUENCY=10
ENV DURATION=10
ENV METHOD=GET
ENV PROMETHEUS_ADDR=0.0.0.0:8880
EXPOSE 8880
ENTRYPOINT ["/bin/loadtest"]