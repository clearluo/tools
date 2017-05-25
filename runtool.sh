#!/bin/bash
#*/1 * * * * /bin/bash -f /home/didong/go/runtool.sh status >> /home/didong/go/crontab.log 2>&1
GOPATH=/home/didong/go
GOSRC=$GOPATH/src/didong
GOBIN=$GOPATH/bin
BIN=didong-backend
CURTIME=`date +[%Y-%m-%d:%H:%M]`

cd $GOPATH
if [ $# -ne 1 ]
	then
  	echo $CURTIME $0 "start | stop | restart | status | monitor"
 	exit
fi
PID=`ps -ef | grep $BIN | grep -v grep | awk '{print $2}'`
echo "$CURTIME current pid:" $PID
function start() {
	echo "$CURTIME start..."
	$GOBIN/$BIN >> temp.log 2>&1 &
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
		if [ 0"$PID" = "0" ];then
			start
		else
			echo "$CURTIME $BIN already run..."
			stop
			sleep 3
			start
		fi
		exit
		;;
	"stop")
		echo "run stop..."
		stop
		exit;;
	"restart")
		echo "run restart..."
		if [ 0"$PID" = "0" ];then
			start
		else
			stop
			sleep 3
			start
		fi
		exit
		;;
	"status")
		echo "run status..."
		if [ 0"$PID" = "0" ];then
			echo "$CURTIME program is not run..."
		fi
		exit
		;;
	"monitor")
		echo "$CURTIME run monitor..."
		if [ 0"$PID" = "0" ];then
			echo "$CURTIME monitor xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
			echo "$CURTIME program is not run..."
			start
		else
			echo "$CURTIME $BIN already run..."
		fi
		exit
		;;
	*)
		echo $0 "$CURTIME start | stop | restart | status | monitor"
		exit;;
esac


