package reminder

import (
	"eatOnTime/util"
	"fmt"
	"io/ioutil"
	"time"
)

func Remind() {
	var cstZone = time.FixedZone("CST", 8*3600)
	for {
		currentHour := time.Now().In(cstZone).Hour()
		currentMin := time.Now().In(cstZone).Minute()
		if currentHour == 12 && (currentMin == 0 || currentMin == 5) {
			msg := fmt.Sprintf("该吃午饭啦！\n吃饭不积极，思想有问题! 都已经 %d:%d 了，让我看看是谁没站起来！", currentHour, currentMin)
			sendMsg(msg)
		}
		if currentHour == 18 && (currentMin == 30 || currentMin == 35 || currentMin == 40 || currentMin == 45) {
			msg := fmt.Sprintf("该吃晚饭啦！\n晚饭吃得早，下班下的好! 都已经 %d:%d 了，让我看看是谁还在加班！", currentHour, currentMin)
			sendMsg(msg)
		}
		time.Sleep(time.Minute)
	}
}

func sendMsg(msg string) {
	data := Data{
		Msg:  msg,
		To:   "3080188",
		Type: "4",
		Sub:  fmt.Sprintf("吃饭机器人"),
	}

	url := "http://tellus.corp.youdao.com/api/v2/send"

	header := make(map[string]string, 2)
	header["TOKEN"] = "eyJkZXB0SWQiOiAiRDc5OTcxMDAxMCIsICJkZXB0TmFtZSI6ICJcdTY3MDlcdTkwNTNcdTRlOGJcdTRlMWFcdTdmYTQtXHU2NzA5XHU5MDUzXHU1MTZjXHU1MTcxXHU2MjgwXHU2NzJmXHU2NTJmXHU2MzAxXHU5MGU4LVx1NTdmYVx1Nzg0MFx1NjdiNlx1Njc4NFx1N2VjNCIsICJwaG9uZSI6ICIxMzU4MTg2MDM5NyIsICJqb2JOdW1iZXIiOiAiQjI1MTgzIiwgImNvcnBNYWlsIjogImNoZW5sZWkxMkBjb3JwLm5ldGVhc2UuY29tIn0="
	resp, err := util.HTTPPost(url, data, nil, header)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

type Data struct {
	Msg  string
	To   string
	Type string
	Sub  string
}
