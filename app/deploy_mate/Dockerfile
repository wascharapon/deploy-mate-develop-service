FROM golang:1.21.6

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN mkdir -p /workspace
WORKDIR /workspace
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -trimpath -o .build/main -ldflags "-w -s" ./app/deploy_mate
FROM gcr.io/distroless/static

COPY --from=0 /workspace/.build/ /

ENTRYPOINT ["/main"]