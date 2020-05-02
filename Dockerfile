FROM golang:1.14-alpine
LABEL maintainer="Dawud Ibrahim Ismail"
LABEL email="ismaildawud96@gmail.com"
RUN mkdir /app
ADD . /app
WORKDIR /app
# add gcc for building binaries
RUN apk add --no-cache gcc musl-dev
# Add this go mod download command to pull in any dependencies
RUN go mod download
# install swaggo to run swagger
RUN go get -u github.com/swaggo/swag/cmd/swag
# generate the api spec json/yml
RUN swag init
# Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]