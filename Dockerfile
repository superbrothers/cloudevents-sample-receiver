FROM golang:1.12

RUN mkdir /src
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM scratch

COPY --from=0 /src/out/cloudevents-sample-receiver /
ENTRYPOINT ["/cloudevents-sample-receiver"]
