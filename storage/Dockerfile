FROM iron/base

ADD storageservice-linux-amd64 /

ENTRYPOINT ["./storageservice-linux-amd64", "--db_dialect=sqlite3", "--db_args=<ARGS>", "--amqp_url=<URL>"]