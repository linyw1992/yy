FROM golang:1.14.15-alpine3.13 AS build

COPY ./ /go/src/github.com/common-contracts
WORKDIR /go/src/github.com/common-contracts

# Build application
RUN go build -mod=vendor -o chaincode -v .

# Production ready image
# Pass the binary to the prod image
FROM alpine:3.13 as prod

COPY --from=build /go/src/github.com/common-contracts/chaincode /app/chaincode

USER 1000

WORKDIR /app
CMD ./chaincode