FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD go-oo /go-oo
ENTRYPOINT ["/go-oo"]

EXPOSE 3000
