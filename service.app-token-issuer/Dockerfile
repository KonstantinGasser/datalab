FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8006

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8006", \
                "-db-srv=mongodb://apptokenstorage:secure@192.168.0.177:27020", \
                "-app-srv=appservice:8003"]