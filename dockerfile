# Use official Golang image (specific version for consistency)
FROM golang:1.22.2-alpine

# Metadata (OCI-style)
LABEL org.opencontainers.image.title="ascii-art-web-dockerize"
LABEL org.opencontainers.image.description="Run the ASCII art web app anywhere using Docker solution"
LABEL org.opencontainers.image.version="1.0.0"
LABEL org.opencontainers.image.authors="Mohamed NOURI mednouri101@gmail.com"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://learn.zone01oujda.ma/git/bnomenja/ascii-art-web-dockerize.git"

# Set working directory inside the container
WORKDIR /ascii-art-web

# Copy everything from current directory into container
COPY . .

# Build the Go app
RUN go build -o ascii-art-web

# Expose port 8080 (the port our app listens on)
EXPOSE 8080

# Run the compiled program when container starts
CMD ["./ascii-art-web"]