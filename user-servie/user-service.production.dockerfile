FROM alpine:latest
WORKDIR /app
COPY userApp ./
CMD ./userApp