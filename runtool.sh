#!/bin/bash
#*/1 * * * * /bin/bash -f /root/go/goproject/src/runtool.sh status >> /root/go/goproject/src/crontab.log 2>&1
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
		stop
		sleep 2
		start;;
	"stop")
		stop;;
	"restart")
		stop
		sleep 2
		start;;
	"status")
		if [ 0"$PID" = "0" ];then
			echo "$CURTIME program is not run..."
		fi
		;;
	*)
		echo $0 "$CURTIME start | stop | restart | status"
		exit;;
esac


#bee run -gendoc=true >> temp.log 2>&1 &
