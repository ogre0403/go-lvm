##############################################
# This Dockerfile is created for the CI test #
##############################################

FROM golang:1.15.7-alpine3.13 as build

RUN apk update ; \
apk add --no-cache libblockdev-dev lvm2 gcc make libc-dev

RUN mkdir /go-lvm
WORKDIR /go-lvm

COPY . .
#RUN if [ ! -d "/iscsi-provisioner/vendor" ]; then  go mod vendor; fi

RUN go build cmd/main.go



FROM alpine:3.13

RUN apk update ; \
apk add --no-cache libblockdev-dev lvm2

COPY --from=build /go-lvm/main /
CMD ["/main"]