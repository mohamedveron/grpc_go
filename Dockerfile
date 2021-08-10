FROM golang:1.15-alpine

RUN apk add --no-cache git
RUN apk add --update make
RUN mkdir -p /app

# Move to working directory /app
WORKDIR /app

# Copy the code into the container
COPY . .

RUN go mod download

RUN make build

# Export necessary port
EXPOSE 9090
# Command to run when starting the container
CMD ["bin/app"]