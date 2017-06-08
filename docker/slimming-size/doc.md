## Slimming Docker images

### 1. Image layers:

**Read more in [AUFS](http://www.thegeekstuff.com/2013/05/linux-aufs/) section to understand how image layers work in docker**

`redis.tar.gz` is still in the image, because it is deleted in the next `RUN` command, hence only removed in the next image layer.
```Dockerfile
FROM debian:jessie

RUN apt-get update && \
 apt-get install -y gcc libc6-dev make curl && \
 curl -sSL "http://download.redis.io/releases/redis-3.0.5.tar.gz" -o redis.tar.gz

RUN rm -rf redis.tar.gz
```

Does not include redis.tar.gz in its image, because all occured within single `RUN` command, i.e. only single image layer is created 

```Dockerfile
FROM debian:jessie

RUN apt-get update && \
 apt-get install -y gcc libc6-dev make curl && \
 curl -sSL "http://download.redis.io/releases/redis-3.0.5.tar.gz" -o redis.tar.gz && \ 
 rm -rf redis.tar.gz
```


**Learnings:** Do everything within a single `RUN` if possible


### 2. Docker squash:

#### Install docker-squash

```bash
go get github.com/jwilder/docker-squash
```

For Mac OS X users, install GNU tar: 

```bash
brew install gnu-tar
export PATH=/usr/local/opt/gnu-tar/libexec/gnubin:$PATH
```

Dockerfile for setting up go environment in a debian environment: 

```bash
cd app 
docker build -t go-exp:unsquashed .
docker images | grep go-exp

#outputs: 
go-exp              unsquashed          5900655f9e2b        14 minutes ago      456 MB
```

Now squash it to include apt-get purge and archive clean-up: 
```bash
docker save go-exp:unsquashed > go-exp.tar
sudo docker-squash -i go-exp.tar -o squashed.tar -t squashed
cat squashed.tar| docker import - go-exp:squashed
```

Check new size: 
```
docker images | grep go-exp
#outputs: 
go-exp              squashed            2144699bb824        11 minutes ago      378 MB
go-exp              unsquashed          5900655f9e2b        16 minutes ago      456 MB
```

Articles on docker-squash:
http://jasonwilder.com/blog/2014/08/19/squashing-docker-images/

**Learnings:** Squash the docker image before deploying

### 3. Alpine:

Extremely lightweight linux images with preinstalled package-managers

**Learnings:** Use it as a base image!