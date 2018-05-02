FROM clmgr-base

RUN mkdir -p /go/src/myproj.com/clmgr-lrm

ENV GOPATH=/go/
ENV GOROOT=/usr/lib64/go/1.9/
ENV PATH=$PATH+:/usr/bin/go

WORKDIR /go/src/myproj.com/clmgr-lrm
ADD ./ /go/src/myproj.com/clmgr-lrm/

CMD ["make"]