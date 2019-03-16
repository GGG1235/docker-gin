FROM scratch
ONBUILD RUN mkdir /app
#ONBUILD RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
ADD main /app/
WORKDIR /app
ENV PORT 8080
EXPOSE 8080
CMD ["/app/main"]