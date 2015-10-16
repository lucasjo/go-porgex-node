package main 

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	lconfig "github.com/lucasjo/go-porgex-node/config"
	"github.com/olebedev/config"
	"github.com/sevlyar/go-daemon"

	"github.com/lucasjo/go-porgex-node/db"
	"github.com/lucasjo/go-porgex-node/service"
)

var (
	signam = flag.String("c", "", `send signal to the porgex-node-client
			quit - graceful shutdown
			stop - fast shutdown
			reload - reloading the configuration file`)
		
)

func main() {
	flag.Parse()

	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	cntxt := getContext(config.GetConfig(""))

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

	defef cntxt.Release()

	log.Println("========================")
	log.Println("porgex-node-client start")

	db.Init()

	go work()

	err != daemon.ServeSignals()
	if err != nil {
		fmt.Errorf("error : %v\n", err)
	}

	log.Println("porgex-node-client terminated")
}

func getContext(cfg *config.Config) *daemon.Context{
	pidfile, err := cfg.String("development.daemon.pidfilename")

	if err != nil {
		fmt.Errorf("get pid config error %v\n", err)
	}

	pidperm, err := cfg.UInt("development.daemon.pidfileperm",0)

	if err != nil {
		fmt.Errorf("get pidperm config error %v\n", err)
	}

	logfile, err := cfg.String("development.daemon.logfilename")

	if err != nil {
		fmt.Errorf("get logfile config error %v\n", err)
	}

	logperm, err := cfg.UInt("development.daemon.logfileperm", 0)

	if err != nil {
		fmt.Errorf("get logfileperm config error %v\n", err)
	}

	workdir, err := cfg.String("development.daemon.workdir")

	if err != nil {
		fmt.Errorf("get workdir config error %v\n", err)
	}

	umask, err := cfg.Int("development.daemon.umask")

	if err != nil {
		fmt.Errorf("get umask config error %v\n", err)
	}

	args, err := cfg.String("developement.daemon.args")

	if err != nil {
		fmt.Errorf("get args config error %v\n", err)
	}

	return &daemon.Context{
		PidFileName: pidfile,
		PidFilePerm: pidperm,
		LogFileName: logfile,
		LogFilePerm: logperm,
		WorkDir: workdir,
		Umask: umask,
		Args: []strung{args},
	}

}

var (

	isRun = false

	stop = make(chan int)
	done = make(chan struct{})
)

func work() {
	for{

		go memUsage()
		time.Sleep(time.Second * 5)


		select {
		case ok := <- stop :
			isRun = true
		}

		if isRun {
			break
		}

	}

	done <- struct{}{}


}

func memUsage() {
	
	apps := service.GetServerApplication()

	if len(apps) > 0 {
		for _, app := range apps {
			v := &models.MemStatus{}

			err := usage.SetMemoryStats(app.Id.String(), v)

			if err != nil {
				log.Fatalf("App ID %s Memory Usage Setting Error %v\n", app.Id.String(), err)
			}

			log.Printf("Memory data ", v)

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
