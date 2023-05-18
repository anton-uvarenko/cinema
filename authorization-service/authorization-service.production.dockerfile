FROM alpine:latest
WORKDIR /app
COPY authApp ./
COPY ./templates ./templates
CMD ./authApp