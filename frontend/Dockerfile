FROM golang:1.18.10-alpine3.17 as build

ENV GOPATH /go

ARG BACKEND_URL
ENV BACKEND_URL $BACKEND_URL

WORKDIR $GOPATH/src/app

COPY . .

RUN apk add --no-cache make

RUN go mod init github.com/liamchampton/tech-days/frontend

RUN go mod download

RUN go mod tidy

RUN go install github.com/gopherjs/gopherjs@v1.18.0-beta2

RUN make buildfrontend

RUN chmod +x frontend

FROM alpine:3.17

COPY --from=build go/src/app/layout /layout
COPY --from=build go/src/app/scripts /scripts
COPY --from=build go/src/app/data /data
COPY --from=build go/src/app/frontend /

EXPOSE 4321

ENTRYPOINT [ "./frontend" ]