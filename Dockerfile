FROM clmgr-base

RUN mkdir -p /go/src/myproj.com/clmgr-lrm

ENV GOPATH=/go/
ENV GOROOT=/usr/lib64/go/1.9/
ENV PATH=$PATH+:/usr/bin/go

# setting agents stuff
COPY ./test/etc/ /etc/
COPY ./test/resources/ /home/
COPY ./test/test_agents/ /opt/clmgr/agents/

WORKDIR /go/src/myproj.com/clmgr-lrm
ADD ./ /go/src/myproj.com/clmgr-lrm/

# coping config
COPY ./config/config.toml /opt/clmgr/config/config.toml

CMD ["./build/clmgr-lrm"]