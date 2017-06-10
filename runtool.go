//#!/usr/bin/env gorun
package main

/**
 * go get github.com/erning/gorun
 */
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Program struct {
	GoPath string
	GoPro  string
	GoBin  string
	GoLog  string
	Exe    string
	CurPid string
}

func init() {
	log.SetFlags(log.LstdFlags)
	time.Local = time.FixedZone("CST", 3600*8)
}

func main() {
	program := Program{
		GoPath: os.Getenv("GOPATH"),
		GoPro:  os.Getenv("GOPATH") + "/src/didong",
		GoBin:  os.Getenv("GOPATH") + "/bin",
		GoLog:  os.Getenv("GOPATH") + "/logs",
		Exe:    "didong-backend",
	}
	if len(os.Args) != 2 {
		printHelp()
		os.Exit(0)
	}
	program.CurPid = program.GetPid()
	switch os.Args[1] {
	case "start":
		if "" == program.CurPid {
			program.Start()
		} else {
			log.Println("The program already running, PID:", program.CurPid)
		}
	case "stop":
		if "" == program.CurPid {
			log.Println("The program is not running.")
		} else {
			program.Stop()
		}
	case "restart":
		if "" == program.CurPid {
			program.Start()
		} else {
			program.Stop()
			time.Sleep(time.Second * 3)
			program.Start()
		}
	case "status":
		if "" == program.CurPid {
			log.Println("The program is not running.")
		} else {
			log.Println("The current program PID:", program.CurPid)
		}
	case "monitor":
		if "" == program.CurPid {
			log.Println("Monitor program, program dead, start program.")
			program.Start()
		} else {
			log.Println("Monitor program, PID:", program.CurPid)
		}
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("./xx start | stop | restart | status | monitor")
}

func GetCurTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func TrimEnd(result []byte) string {
	if len(result) > 2 {
		result = append(result[0:0], result[0:len(result)-1]...)
	}
	return string(result)
}
func ExecShell(cmd string) (string, error) {
	ret, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Println(err)
		os.Exit(0)
		return "", err
	}
	return TrimEnd(ret), nil
}
func (p Program) GetPid() string {
	cmd := bytes.Buffer{}
	cmd.WriteString(`ps -ef | grep `)
	cmd.WriteString(p.Exe)
	cmd.WriteString(` | grep -v grep | awk '{print $2}'`)
	ret, _ := ExecShell(cmd.String())
	return ret
}

func (p *Program) Start() {
	if p.GetPid() != "" {
		log.Println("The program already running, PID:", p.CurPid)
		os.Exit(0)
	}
	cmd := p.GoBin + "/" + p.Exe + " >> " + p.GoLog + "/temp.log 2>&1 &"
	ExecShell(cmd)
	time.Sleep(time.Second * 3)
	p.CurPid = p.GetPid()
	if "" == p.CurPid {
		log.Println("The program start failure.")
	} else {
		log.Println("The program start success, PID:", p.CurPid)
	}
}
func (p Program) Stop() {
	ExecShell("kill -9 " + p.CurPid)
	log.Println("The program stop success, PID:", p.CurPid)
}
