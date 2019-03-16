# docker-gin

***

环境:

Docker version 18.09.2


Dockerfile:

```
FROM scratch
ONBUILD RUN mkdir /app
ADD main /app/
WORKDIR /app
ENV PORT 8080
EXPOSE 8080
CMD ["/app/main"]
```

在终端运行指令```docker build -t docker-gin .```,之后运行指令```docker run -it -p 8080:8080 docker-gin```来开启服务


### 截图:

启动服务后在宿主机打开页面

<img src="https://github.com/GGG1235/docker-gin/blob/master/images/windows.png" width="375" alt="windows打开">

查看Windows局域网ip

<img src="https://github.com/GGG1235/docker-gin/blob/master/images/win-ip.png" width="375" alt="查看Windows局域网ip">

在mac上打开服务

<img src="https://github.com/GGG1235/docker-gin/blob/master/images/1.png" width="375" alt="在mac上打开服务">

可以看到windows下服务器端显示的信息,同一局域网下不同的机器访问同一个页面记录的ip地址,都是docker分配给宿主机的ip

<img src="https://github.com/GGG1235/docker-gin/blob/master/images/server.png" width="375" alt="在mac上打开服务">
