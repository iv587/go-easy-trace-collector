FROM busybox
MAINTAINER lijunfei developerljf2020@163.com
WORKDIR /usr/app
COPY ./easycboot /usr/app/
COPY ./app.yml /usr/app/app.yml
ADD ./static /usr/app/static
EXPOSE 6666
EXPOSE 8081
ENTRYPOINT ["./easycboot"]