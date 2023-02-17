FROM busybox
MAINTAINER lijunfei developerljf2020@163.com
WORKDIR /usr/app
COPY ./easycboot /usr/app/
COPY ./app.yml /usr/app/app.yml
ADD ./static /usr/app/static
ADD ./zoneinfo.tar.gz /usr/share/
RUN  cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
EXPOSE 6666
EXPOSE 8081
ENTRYPOINT ["./easycboot"]