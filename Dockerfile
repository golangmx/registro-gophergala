MAINTAINER Arturo Vergara <hello@arturovm.me>
FROM scratch

COPY bin/registro-linux-amd64 /registro
COPY static/dist /static
EXPOSE 8080
CMD ["/registro"]
