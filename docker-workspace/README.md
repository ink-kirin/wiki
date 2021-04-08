# docker-workspace

```
├── config
│   ├── mysql
│   │   └── my.cnf
│   ├── nginx
│   │   ├── conf.d
│   │   │   ├── example.conf
│   │   │   └── test.conf
│   │   ├── Dockerfile
│   │   └── nginx.conf
│   ├── php
│   │   ├── Dockerfile
│   │   ├── php-fpm.conf
│   │   └── php.ini
│   └── redis
│       └── redis.conf
├── .env
├── src
├── data
├── log
└── docker-compose.yml
```

# 常用的Docker命令
```
# 1.显示所有容器
docker ps -a[包括未运行] -q[仅显示编号]

# 2.停止、重启、启动某一容器
docker stop|restart|start 容器id|容器名

# 3.停止、重启、启动所有容器
docker stop|restart|start $(docker ps -a -q)

# 4.获取容器ip
docker inspect 容器id

# 5.容器开机启动
docker update --restart=always $(docker ps -a -q)

# 6.删除容器[需要先停止运行]
docker rm 容器id|容器名

# 7.删除镜像[需要先停止且删除所有关联的容器]
docker rmi 镜像id

# 8.进入容器
docker exec -it 容器id|容器名 bash

# 9.搜索镜像
docker search 镜像关键字

# 10.下载镜像
docker pull 镜像名字:版本号

# 11.查看本机所有docker镜像
docker images

# 12.导出镜像
docker save -o 导出的镜像文件.tar 镜像名字:版本号

# 13.导入镜像
docker load -i 镜像文件.tar

# 14.从容器里面拷文件到宿主机
docker cp 容器名：要拷贝的文件在容器里面的路径   要拷贝到宿主机的相应路径
# 如：
docker cp myphp:/home/data/test/js/test.js /opt

# 15.从宿主机拷文件到容器里面
docker cp 要拷贝的文件路径 容器名：要拷贝到容器里面对应的路径
# 如：
docker cp /opt/test.js myphp:/home/data/test/js
```

# Docker-compose 指令
URL：https://deepzz.com/post/docker-compose-file.html
```
version：版本，当前版本定义为“3”即可，在文件docker-compose.yml的第一行必须指定
networks：定义网络连接，如果是在services下的networks则是指定容器的网络连接配置
services：各容器服务
container_name：指定容器名称
image：指定镜像，如果镜像在本地不存在，Compose 将会尝试拉取这个镜像
build：指定构建自定义镜像位置，对应的是docerfile文件。本文的PHP镜像采用的是自定义hwphp镜像
ports：暴露端口信息，使用宿主端口：容器端口 (HOST:CONTAINER) 格式
volumes：数据卷映射，数据卷所挂载路径设置，格式为宿主机路径:容器路径
restart：自动重启容器，容器挂掉之后自动重启机制
depends_on：指定容器启动顺序的依赖关系，此选项在 v3 版本中 使用 swarm 部署时将忽略该选项
environment：设置环境变量， environment 的值可以覆盖 env_file 的值 (等同于 docker run -e 的作用
extra_hosts: 添加主机名映射。类似 docker client --add-host。
networks: 配置容器连接的网络，引用顶级 networks 下的条目
```

---------------------------------------
Docker Url: http://hub.docker.com/_/php


### fpm alpine
URL：https://medium.com/@charliecc/%E7%94%A8docker%E5%BB%BA%E7%AB%8Bphp-nginx-%E9%96%8B%E7%99%BC%E7%92%B0%E5%A2%83-33c5f88edeb3
```
# 直接用 official image，可以省下很多步骤
FROM php:7.4.12-fpm-alpine

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 安裝 php extensions 的神器，請務必一試！
# https://github.com/mlocati/docker-php-extension-installer
COPY --from=mlocati/php-extension-installer /usr/bin/install-php-extensions /usr/bin/

RUN install-php-extensions redis mysqli xdebug

WORKDIR /var/www/
```

### fpm
Url地址：https://www.helloweba.net/server/627.html
```
FROM php:7.4-fpm

# 設定時區
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 更新安裝依賴包和PHP核心拓展
RUN apt-get update && apt-get install -y \
        --no-install-recommends libfreetype6-dev libjpeg62-turbo-dev libpng-dev curl \
        && rm -r /var/lib/apt/lists/* \
        && docker-php-ext-configure gd \
        && docker-php-ext-install -j$(nproc) gd opcache pdo_mysql gettext sockets

# 安裝 PECL 拓展，安裝Redis，swoole
RUN pecl install redis \
    && pecl install swoole \
    && docker-php-ext-enable redis swoole

# 安裝 Composer
ENV COMPOSER_HOME /root/composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
ENV PATH $COMPOSER_HOME/vendor/bin:$PATH

WORKDIR /data
```

### fpm
```
From php:7.4-fpm
RUN apt-get update && apt-get install -y \
    libfreetype6-dev \
    libjpeg62-turbo-dev \
    libpng-dev \
&& docker-php-ext-install -j$(nproc) iconv \
&& docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
&& docker-php-ext-install -j$(nproc) gd \
&& docker-php-ext-configure pdo_mysql \
&& docker-php-ext-install pdo_mysql \
&& pecl install redis-4.3.0 \
&& pecl install swoole \
&& docker-php-ext-enable redis swoole
```