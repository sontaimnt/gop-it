FROM ubuntu:latest
RUN apt update
RUN apt install golang git -y
RUN git clone https://github.com/sontaimnt/gop-it
WORKDIR /gop-it
RUN go build -v
