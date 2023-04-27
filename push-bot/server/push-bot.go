package main

import (
	"fmt"
	"github.com/robfig/cron"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

/*


添加算法细狗(push-bot)，每天三次攻击性催促打卡。
*/

const (
	reply1 = "嗨嗨嗨，今天你们打卡了吗？不要忘记写题哦.细狗们行不行？菜就多练👎👎👎"
	reply2 = "卧槽cnsm，怎么还有人没写题，快去打卡啊细狗们👎👎👎！！！"
	reply3 = "算法细狗快打卡啊，是不是细狗👎👎，不刷就滚👎"
)

func main() {
	var cstZone = time.FixedZone("CST", 8*3600)
	c := cron.New()
	_ = c.AddFunc("0 32 10,15,22 * * *", func() {
		switch time.Now().In(cstZone).Hour() {
		case 10:
			SendText(reply1)
		case 15:
			SendText(reply2)
		case 22:
			SendText(reply3)

		}

	}) ///每天的10,15,22:32攻击打卡

	c.Start()
	select {}

}

func SendText(text string) {
	j := fmt.Sprintf(`{
    "msg_type": "text",
    "content": {
        "text": "%v"
    }
}`, text)

	url := os.Getenv("push_bot_url")

	resp, err := http.Post(url, "application/json", strings.NewReader(j))
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	b, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		log.Println(err1)
		return
	}

	if !strings.Contains(string(b), "success") {
		log.Println("send failed")
	}

	return
}
