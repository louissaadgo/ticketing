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
FROM alpine
WORKDIR /
COPY --from=builder ./app/auth ./
CMD [ "./auth" ]