FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8002

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8005"]