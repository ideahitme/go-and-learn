FROM debian:wheezy

RUN apt-get update 
RUN apt-get -y upgrade 
RUN apt-get install -y wget

RUN wget --no-check-certificate https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz && \
  tar -xvf go1.8.linux-amd64.tar.gz && \
  mv go /usr/local && \
  mkdir $HOME/go

ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

# clean up
RUN rm -rf go1.8.linux-amd64.tar.gz
RUN apt-get purge -y man  perl-modules vim-common vim-tiny && \
  apt-get clean autoclean && \
  apt-get autoremove -y && \
  rm -rf /var/lib/{apt,dpkg,cache,log}

