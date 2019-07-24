FROM scratch

EXPOSE 8080

ADD sine-service /sine-service

ENTRYPOINT ["/sine-service"]
