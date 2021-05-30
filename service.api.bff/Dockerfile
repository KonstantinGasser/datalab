FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8080

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8080", \ 
                "-user-srv=service-user-administer:8001", \  
                "-app-srv=service-app-administer:8003", \
                "-apptoken-srv=service-app-tokenissuer:8006", \
                "-config-srv=service-app-configuration:8005", \
                "-token-srv=service-user-authentication:8002", \
                "-notify-srv=service-notification-live:8008"]