FROM golang:latest

WORKDIR /app

# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]
