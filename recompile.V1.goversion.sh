#author:luosw
#create date:  2014-04-21
#modify date:2017-05-17

GOPATH=/home/didong/go
GOBIN=$GOPATH/bin
GOSRC=$GOPATH/src
BIN=didong-backend
PROJECT=didong

CURTIME=`date +%Y%m%d%H%M`

cd $GOSRC
if [ $? -eq 0 ]
then
 echo "[cd $GOSRC] is successful!!!"
else
 echo "[cd $GOSRC] is fail!!!"
 exit
fi

rm -rf $PROJECT*
if [ $? -eq 0 ]
then
 echo "[rm -rf $PROJECT*] is successful!!!"
else
 echo "[rm -rf $PROJECT*] is fail!!!"
 exit
fi

rz
if [ $? -eq 0 ]
then
 echo "[rz] is successful!!!"
else
 echo "[rz] is fail!!!"
 exit
fi

unzip $PROJECT.zip
if [ $? -eq 0 ]
then
 echo "[unzip $PROJECT.zip] is successful!!!"
else
 echo "[unzip $PROJECT.zip] is fail!!!"
 exit
fi

cd $GOSRC/$PROJECT
if [ $? -eq 0 ]
then
 echo "[cd $GOSRC/$PROJECT] is successful!!!"
else
 echo "[cd $GOSRC/$PROJECT] is fail!!!"
 exit
fi

go build -o ./$BIN
if [ $? -eq 0 ]
then
 echo "[go build -o ./$BIN] is successful!!!"
else
 echo "[go build -o ./$BIN] is fail!!!"
 exit
fi

cd $GOPATH
if [ $? -eq 0 ]
then
 echo "[cd $GOPATH] is successful!!!"
else
 echo "[cd $GOPATH] is fail!!!"
 exit
fi

./runtool.sh stop
if [ $? -eq 0 ]
then
 echo "[./runtool.sh stop] is successful!!!"
else
 echo "[./runtool.sh stop] is fail!!!"
 exit
fi

if [ -f $GOBIN/$BIN ];then
	mv $GOBIN/$BIN $GOBIN/$BIN.$CURTIME.bak
fi

cp $GOSRC/$PROJECT/$BIN $GOBIN/$BIN
if [ $? -eq 0 ]
then
 echo "[cp $GOSRC/$PROJECT/$BIN $GOBIN/$BIN] is successful!!!"
else
 echo "[cp $GOSRC/$PROJECT/$BIN $GOBIN/$BIN] is fail!!!"
 exit
fi

./runtool.sh start
if [ $? -eq 0 ]
then
 echo "[./runtool.sh start] is successful!!!"
else
 echo "[./runtool.sh start] is fail!!!"
 exit
fi

