FROM golang:1-alpine

RUN apk add --no-cache git librsvg ffmpeg

COPY . /src
WORKDIR /src

RUN go get -d -v github.com/ajstarks/svgo/float
RUN go install -v github.com/ajstarks/svgo/float

CMD go build -o goflakes main.go