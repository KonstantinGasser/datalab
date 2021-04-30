FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8080

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8080", \ 
                "-user-srv=userservice:8001", \  
                "-app-srv=appservice:8003", \
                "-token-srv=tokenservice:8002"]
