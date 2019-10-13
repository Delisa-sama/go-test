FROM iron/base
EXPOSE 9000

ADD newsservice-linux-amd64 /

ENTRYPOINT ["./newsservice-linux-amd64", "--port=9000", "--amqp_url=<URL>", "--storage_timeout=<SECONDS>"]