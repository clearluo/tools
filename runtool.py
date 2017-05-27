#!/usr/bin/python
import os
import sys
import datetime
import time
#*/1 * * * * /usr/bin/python /home/didong/go/runtool.py monitor >> /home/didong/go/logs/crontab.log 2>&1
os.environ["GOPATH"] = "/home/didong/go"
os.environ["GOLOG"] = os.environ["GOPATH"] + "/logs"
os.environ["GOSRC"] = os.environ["GOPATH"] + "/src/didong"
os.environ["GOBIN"] = os.environ["GOPATH"] + "/bin"
os.environ["BIN"]="didong-backend"
curTime = "["+datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')+"] "
outPut = os.popen('ps -ef | grep $BIN | grep -v grep | awk \'{print $2}\'')
outPutStr = outPut.read()
curPid = 0
if outPutStr != "":
	curPid = int(outPutStr)
os.environ["CURPID"] = str(curPid)

if os.system('cd $GOPATH'):
	exit()

if len(sys.argv)!=2:
	print(curTime+sys.argv[0]+" start | stop | restart | status | monitor")
	exit()

def start():
	if os.system('$GOBIN/$BIN >> $GOLOG/temp.log 2>&1 &') == 0:
		time.sleep(3)
		newOutPut = os.popen('ps -ef | grep $BIN | grep -v grep | awk \'{print $2}\'')
		newOutPutStr = newOutPut.read()
		newCurPid = 0
		if newOutPutStr != "":
			newCurPid = int(newOutPutStr)
			print(curTime+"程序启动成功,进程ID:",newCurPid)
		else:
			print(curTime+"程序启动失败")
def stop():
	if os.system('kill -9 $CURPID') == 0:
		print(curTime+"程序停止成功,进程ID:",curPid)

cmd = sys.argv[1]
if cmd == "start":
	if curPid == 0:
		start()
	else:
		print(curTime+"程序已经运行,进程ID:",curPid)
elif cmd == "stop":
	if curPid == 0:
		print(curTime+"程序没有运行")
	else:
		stop()
elif cmd == "restart":
	if curPid == 0:
		start()
	else:
		stop()
		time.sleep(3)
		start()
elif cmd == "status":
	if curPid == 0:
		print(curTime+"程序没有运行")
	else:
		print(curTime+"当前程序进程ID:",curPid)
elif cmd == "monitor":
	if curPid == 0:
		print(curTime+"监控程序:程序已死,启动程序")
		start()
	else:
		print(curTime+"监控程序:进程ID:",curPid)
else:
	print(curTime+sys.argv[0]+" start | stop | restart | status | monitor")
	exit()



