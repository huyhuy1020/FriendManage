FROM golang:alphine as builder
ENV GO111MODULE = ON

#because the lastest alphine images does not have tool to support `bash`, `git`
#adding git, bash, openssh to image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh \

#Add name
LABEL maintainer = "Huy-Nguyen-Tuan"

WORKDIR /usr/src/app
COPY .env .
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#Run Go app
RUN go build -o main.

FROM alpine:latest
RUN apk--no-cache add ca-certificates

WORKDIR /root

#copy the Pre-build
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

#Expose port 8080 outside world
EXPOSE 8080

#command to execute
CMD["./main"]

