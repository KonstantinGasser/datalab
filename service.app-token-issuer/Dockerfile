FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8006

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8006", \
                "-db-srv=mongodb://apptokenissuer:secure@192.168.0.179:27020"]