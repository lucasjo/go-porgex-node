package te

import (
	"fmt"

	"github.com/lucasjo/go-porgex-node/db"
	"github.com/lucasjo/go-porgex-node/service"
)

func main() {

	//openshift node server test version
	db.Init()

	apps := service.GetServerApplication()

	if len(apps) > 0 {
		for _, a := range apps {
			fmt.Printf("app data : %v\n", a)
		}
	} else {
		fmt.Println("application 이 없음")
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

}
