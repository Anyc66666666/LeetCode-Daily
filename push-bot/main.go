package main

import (
	"context"
	"fmt"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
使用免费的腾讯云函数定时运行push-bot

添加算法细狗(push-bot)，每天三次攻击性催促打卡。部署在免费的腾讯云函数上面，定时运行。
*/

const (
	reply1 = "嗨嗨嗨，今天你们打卡了吗？不要忘记写题哦.细狗们行不行？菜就多练👎👎👎"
	reply2 = "卧槽cnsm，怎么还有人没写题，快去打卡啊细狗们👎👎👎！！！"
	reply3 = "算法细狗快打卡啊，是不是细狗👎👎，不刷就滚👎"
)

type DefineEvent struct {
	// test event define
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func hello(ctx context.Context, event DefineEvent) (string, error) {
	fmt.Println("key1:", event.Key1)
	fmt.Println("key2:", event.Key2)
	var cstZone = time.FixedZone("CST", 8*3600)
	switch time.Now().In(cstZone).Hour() {
	case 10:
		SendText(reply1)
	case 15:
		SendText(reply2)
	case 22:
		SendText(reply3)

	}

	fmt.Println(os.Getenv("url"))

	return fmt.Sprintf("Hello %s!", event.Key1), nil
	//0 22 10,15,22 * * * *

}

func main() {
	cloudfunction.Start(hello)

}

func SendText(text string) {
	j := fmt.Sprintf(`{
    "msg_type": "text",
    "content": {
        "text": "%v"
    }
}`, text)

	url := os.Getenv("url")

	resp, err := http.Post(url, "application/json", strings.NewReader(j))
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	b, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Println(err1)
		return
	}

	if !strings.Contains(string(b), "success") {
		log.Println("send failed")
	}

	return
}
