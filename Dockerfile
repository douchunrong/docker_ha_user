FROM humbleadmin/docker_golang

#项目 USC-HA-User 配置
ENV USC_HA_User_PATH $GOPATH/src/USC-HA-User
RUN echo $USC_HA_User_PATH
RUN mkdir $USC_HA_User_PATH

ADD script/USC-HA-User/  $USC_HA_User_PATH/
#ADD script/gopkg.in/mgo.v2 $GOPATH/src/gopkg.in/mgo.v2
ADD script/start.sh /usr/local/

RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get gopkg.in/mgo.v2
#配置 bee
WORKDIR $GOPATH/src/github.com/beego/bee
RUN go install

#配置 HAUser
WORKDIR $USC_HA_User_PATH/

RUN chmod 777 /usr/local/start.sh

EXPOSE 8080

CMD /usr/local/start.sh
