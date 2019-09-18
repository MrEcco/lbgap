FROM golang:1.13.0-alpine3.10 as builder

# Prepare buildement
RUN apk add --update --no-cache git \
   && mkdir -p /output

# Build
COPY . /go/src/lbgap
WORKDIR /go/src/lbgap
RUN go build -i -o /output/lbgap .

FROM alpine:3.10

COPY --from=builder /output/* /
ENTRYPOINT ["/lbgap"]
