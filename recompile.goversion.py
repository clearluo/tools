#!/usr/bin/python

#author:clearluo
#create date:  2017-05-27
#modify date:2017-05-27

import os
import sys
import datetime
import time

os.environ["GOPATH"] = "/home/didong/go"
os.environ["GOLOG"] = os.environ["GOPATH"] + "/logs"
os.environ["GOSRC"] = os.environ["GOPATH"] + "/src"
gobin = os.environ["GOPATH"] + "/bin"
os.environ["GOBIN"] = gobin
exebin = "didong-backend"
os.environ["BIN"]=exebin
os.environ["MYPROJECT"] = "didong"
os.environ["CURTIME"] = datetime.datetime.now().strftime('%Y%m%d%H%M')

if os.system('cd $GOSRC; rm -rf $MYPROJECT*'):
	exit()

if os.system('cd $GOSRC; rz'):
	exit()

if os.system('cd $GOSRC; unzip $MYPROJECT.zip'):
	exit()

if os.system('cd $GOSRC/$MYPROJECT; make build'):
	print('make build err')
	exit()

if os.system('./runtool.py stop'):
	exit()

if os.path.exists(gobin + "/" + exebin):
	if os.system('mv $GOBIN/$BIN $GOBIN/$BIN.$CURTIME.bak'):
		exit()

if os.system('cp $GOSRC/$MYPROJECT/$BIN $GOBIN/$BIN'):
	exit()

if os.system('./runtool.py start'):
	exit()

