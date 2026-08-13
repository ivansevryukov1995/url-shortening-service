# build a tiny docker image
FROM alpine:latest

WORKDIR /app

COPY /ussApp /app/ussApp
RUN chmod +x /app/ussApp

CMD [ "/app/ussApp" ]