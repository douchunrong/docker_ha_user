FROM humbleadmin/docker_golang

#项目 USC-HA-User 配置
ENV USC-HA-User $GOPATH/src/USC-HA-User

RUN mkdir $USC-HA-User
ADD script/ObjectId/  $USC-HA-User/
ADD script/gopkg.in/mgo.v2 $GOPATH/src/gopkg.in/mgo.v2
ADD script/start.sh /usr/local/

RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee

#配置 bee
WORKDIR $GOPATH/src/github.com/beego/bee
RUN go install

#配置 ObjectId
WORKDIR $USC-HA-User/

RUN chmod 777 /usr/local/start.sh

EXPOSE 8080

CMD /usr/local/start.sh
