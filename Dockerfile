# Golag
# https://hub.docker.com/_/golang
FROM golang:1 as golang
RUN go get -u github.com/tockins/realize


# Nginx
# https://hub.docker.com/_/nginx
FROM nginx:1

WORKDIR /var/www/html

COPY --from=golang /go /go
COPY --from=golang /usr/local/go /usr/local/go
COPY ./app /var/www/html
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

CMD [ "bin/dev" ]
