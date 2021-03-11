FROM golang:1.16
RUN apt update
RUN apt -y install htop strace nano
COPY hello.go ./
CMD bash
