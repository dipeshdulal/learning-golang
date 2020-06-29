FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go

EXPOSE 5000

# Make binary executable inside container
RUN chmod 777 ./main

# When executing command with CMD ["main"], most of the environment variable wont be present
CMD ./main 
