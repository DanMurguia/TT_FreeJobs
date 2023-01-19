FROM golang:latest
COPY . /go/src/github.com/DanMurguia/TT_FreeJobs
WORKDIR /go/src/github.com/DanMurguia/TT_FreeJobs
RUN go get -d -v
RUN go build -o TT_FreeJobs
EXPOSE 8080
CMD ["./TT_FreeJobs"]