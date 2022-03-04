# Start from golang base image
FROM golang:1.17.8 as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ENV GOSUMDB sum.golang.google.cn

WORKDIR /root/

# Get dependancies - will also be cached if we won't change mod/sum
COPY go.* ./
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git show -s --format=%h)'" -a -installsuffix cgo -o /go/bin/starter

######## Start a new stage from alpine:3.13 #######
FROM alpine:3.15

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories
RUN apk add ca-certificates tzdata

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY config.example.yaml config.yaml
COPY --from=builder /go/bin/starter .

EXPOSE 8080 9090

ENTRYPOINT ["./starter"]

CMD ["serve", "--config=./config.yaml"]
