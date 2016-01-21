FROM scratch

MAINTAINER Arturo Vergara <hello@arturovm.me>

COPY bin/registro-linux-amd64 /registro
COPY static/dist /static
EXPOSE 8080
CMD ["/registro"]
