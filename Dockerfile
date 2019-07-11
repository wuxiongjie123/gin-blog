
FROM scratch

WORKDIR $GOPATH/src/gin-blog
COPY . $GOPATH/src/gin-blog

EXPOSE 8000
CMD ["./gin-blog"]

