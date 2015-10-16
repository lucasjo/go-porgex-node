package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/olebedev/config"
	"github.com/sevlyar/go-daemon"

	lconfig "github.com/lucasjo/go-porgex-node/config"
	"github.com/lucasjo/go-porgex-node/db"
	"github.com/lucasjo/go-porgex-node/service"
)

var (
	signal = flag.String("c", "", `send signal to the porgex-node-client
			quit - graceful shutdown
			stop - fast shutdown
			reload - reloading the configuration file`)
)

func main() {
	flag.Parse()

	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	cntxt := getContext(lconfig.GetConfig(""))

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()

		if err != nil {
			log.Fatalln("Unable send signal to the porgex-node-client", err)
		}
		daemon.SendCommands(d)
	}

	child, err := cntxt.Reborn()

	if err != nil {
		log.Fatalln(err)
	}

	if child != nil {
		return
	}

	defer cntxt.Release()

	log.Println("========================")
	log.Println("porgex-node-client start")

	db.Init()

	go work()

	err = daemon.ServeSignals()

	if err != nil {
		fmt.Errorf("error : %v\n", err)
	}

	log.Println("porgex-node-client terminated")
}

func getContext(cfg *config.Config) *daemon.Context {
	pidfile, err := cfg.String("development.daemon.pidfilename")

	if err != nil {
		log.Fatalf("get pid config error %v\n", err)
	}

	logfile, err := cfg.String("development.daemon.logfilename")

	if err != nil {
		log.Fatalf("get logfile config error %v\n", err)
	}

	workdir, err := cfg.String("development.daemon.workdir")

	if err != nil {
		log.Fatalf("get workdir config error %v\n", err)
	}

	umask, err := cfg.Int("development.daemon.umask")

	if err != nil {
		log.Fatalf("get umask config error %v\n", err)
	}

	arg, err := cfg.String("development.daemon.args")

	args := []string{arg}

	if err != nil {
		log.Fatalf("get args config error %v\n", err)
	}

	return &daemon.Context{
		PidFileName: pidfile,
		PidFilePerm: 0644,
		LogFileName: logfile,
		LogFilePerm: 0640,
		WorkDir:     workdir,
		Umask:       umask,
		Args:        args,
	}

}

var (
	isRun = false

	stop = make(chan int)
	done = make(chan struct{})
)

func work() {
	for {
		log.Println("aaaaaa")
		go memUsage()
		time.Sleep(time.Second * 5)

		select {
		case ok := <-stop:
			if ok == 0 {
				isRun = true
			}

		}

		if isRun {
			break
		}

	}

	done <- struct{}{}

}

func memUsage() {

	apps := service.GetServerApplication()

	log.Printf("app count : %v\n", len(apps))

	if len(apps) > 0 {
		for _, app := range apps {

			log.Println(app.ID.String())

			/*
				v := &models.MemStats{}

				err := usage.SetMemoryStats(app.ID.String(), v)

				if err != nil {
					log.Fatalf("App ID %s Memory Usage Setting Error %v\n", app.ID.String(), err)
				}

				log.Printf("Memory data ", v)
			*/

		}
	}

}

func termHandler(sig os.Signal) error {
	log.Println("terminating....")

	stop <- 0

	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

func reloadHandler(sig os.Signal) error {
	log.Println("configuration reload")

	return nil
}

/* 이것은 최종적으로 client service 에 반영 되어야 한다
conn, err := net.Dial("tcp", "127.0.0.1:3001")

if err != nil {
	fmt.Errorf("err : %v\n", err)
	os.Exit(1)
}
str := &models.MemStats{
	Id:            bson.NewObjectId(),
	AppId:         "5000130384e12",
	Max_usage:     801010,
	Limit_usage:   801010,
	Current_usage: 77733,
	Create_at:     time.Now(),
}

d, e := json.Marshal(str)
fmt.Printf("str : %v\n", string(d))
hostname, _ := os.Hostname()

req := &models.Request{
	Service:  "memory",
	Fromhost: hostname,
	Data:     d,
}

b, e := json.Marshal(req)

if e != nil {
	os.Exit(1)
}

_, err = conn.Write(b)

conn.Close()
*/
