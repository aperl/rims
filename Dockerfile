FROM golang:1.13 as base

RUN go get golang.org/x/lint/golint

FROM base as vendored
WORKDIR /src/

COPY go.mod .
COPY go.sum .
RUN go mod download

FROM vendored as dev

WORKDIR /src/
COPY . .

RUN go version
RUN golint -set_exit_status
RUN go vet
RUN go test
RUN go install ./...

WORKDIR /go/bin
EXPOSE 80
CMD ["rims", "80"]

FROM scratch AS release

COPY --from=dev /go/bin/rims .

EXPOSE 80
CMD ["rims", "80"]
