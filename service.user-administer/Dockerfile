FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8001

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8001", \
                "-db-srv=mongodb://userstorage:secure@userstorage:27017"]   