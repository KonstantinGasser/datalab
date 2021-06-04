FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8003

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8003", \
                "-permissions-srv=service-user-authentication:8002", \
                "-config-srv=service-app-configuration:8005", \
                "-apptoken-srv=service-app-tokenissuer:8006", \
                "-db-srv=mongodb://appadminstorage:secure@192.168.0.177:27018"]