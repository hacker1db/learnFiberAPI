# The base go-image
FROM golang:alpine
RUN apk add --update gcc musl-dev
# Create a directory for the app
RUN mkdir /app
# Copy all files from the current directory to the app directory
COPY . /app
# Set working directory
WORKDIR /app
# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o server . 
RUN addgroup -S nonrootgroup && adduser -S appuser -G nonrootgroup
# run as non root user
USER appuser
# Run the server executable
CMD [ "/app/server" ]
