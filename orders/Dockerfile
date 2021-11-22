#Starting base image
FROM golang:1.17.2-alpine3.14 as builder

#Setting up working directory
WORKDIR /app

#Copying and downloading dependencies
COPY ./go.mod ./go.sum ./
RUN go mod download

#Copying everything else
COPY ./ ./

#Creating go executable
RUN go build

#Creating final image
#Starting base image
FROM alpine

#Setting up working directory
WORKDIR /

#Copying go executable from previous image
COPY --from=builder ./app/orders ./

#Setting up starting command
CMD [ "./orders" ]