#!/bin/bash
#59 23 * * * /bin/bash -f /root/go/goproject/src/runtool.sh stop >> /root/go/goproject/src/crontab.log 2>&1
#4 0 * * * /bin/bash -f /root/go/goproject/src/runtool.sh start >> /root/go/goproject/src/crontab.log 2>&1
GOSRC=/root/go/goproject/src/didong
BIN=didong
CURTIME=`date +[%Y-%m-%d:%H:%M]`

cd $GOSRC
if [ $# -ne 1 ]
	then
  	echo $CURTIME $0 "start | stop | restart | status"
 	exit
fi
PID=`ps -ef | grep $BIN | grep -v grep | awk '{print $2}'`
echo "$CURTIME current pid:" $PID
function start() {
	echo "$CURTIME start..."
	bee run -gendoc=true >> temp.log 2>&1 &
	sleep 3
	NEWPID=`ps -ef | grep $BIN | grep -v grep | awk '{print $2}'`
	echo "$CURTIME new pid:" $NEWPID
	echo "$CURTIME start succ..."
}
function stop() {
	echo "$CURTIME stop..."
	echo "$CURTIME kill pid:" $PID
	if [ $PID ];then
		kill -9 $PID
		echo "$CURTIME stop succ..."
	else
		echo "$CURTIME program is not run..."
	fi
}

case $1 in
	"start")
		echo "run start..."
		stop
		sleep 5
		start
		exit;;
	"stop")
		echo "run stop..."
		stop
		exit;;
	"restart")
		echo "run restart..."
		stop
		sleep 5
		start
		exit;;
	"status")
		echo "run status..."
		if [ 0"$PID" = "0" ];then
			echo "$CURTIME program is not run..."
		fi
		exit
		;;
	*)
		echo $0 "$CURTIME start | stop | restart | status"
		exit;;
esac


