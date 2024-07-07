# Use Alpine Linux as base image
FROM alpine:latest

# Install necessary packages (nmap and tcpdump)
RUN apk update && \
      apk add --no-cache \
      nmap \
      tcpdump \
      curl \
      mysql-client \
      openssl \
      mariadb-connector-c-dev

# Set the entrypoint to run bash
CMD ["ash"]
