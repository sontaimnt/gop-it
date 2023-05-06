FROM alpine:latest
RUN apk add go git
RUN git clone https://github.com/sontaimnt/gop-it
RUN cd gop-it
RUN go build -v
