package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type B struct {
	a int
	b int
}

type A struct {
	b1 *B
	b2 *B
}

func main() {

	company := getRandomCompanyMessage()
	str := escape(company.Info)
	fmt.Println("str", str)
	res := time.Now().Weekday().String()
	fmt.Println("res", res)

}

type Company struct {
	Name  string `json:"name"`
	Info  string `json:"info"`
	Url   string `json:"url"`
	Image string `json:"image"`
}

func escape(s string) string {
	j, _ := json.Marshal(s)
	n := len(j)
	return string(j[1 : n-1])
}
func getCompanyCard(company *Company) string {
	var imageKey string
	if company.Image != "" {
		imageKey = company.Image
	} else {
		imageKey = getRandomImage()
	}
	return fmt.Sprintf(`{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "alt": {
        "content": "",
        "tag": "plain_text"
      },
      "img_key": "%s",
      "tag": "img"
    },
    {
      "tag": "div",
      "text": {
        "content": "相关信息:%s",
        "tag": "lark_md"
      }
    },
 
    {
      "actions": [
        {
          "tag": "button",
          "text": {
            "content": "一键开润",
            "tag": "plain_text"
          },
          "type": "primary",
          "url": "%s"
        }
      ],
      "tag": "action"
    }
  ],
  "header": {
    "template": "red",
    "title": {
      "content": "🏃‍♀️🏃‍♀️🏃‍♀️ 今日跑路推荐: %s 🏃🏃🏃",
      "tag": "plain_text"
    }
  }
}`, imageKey, company.Info, company.Url, company.Name)
}

func getRandomCompanyMessage() *Company {
	companyList := []*Company{
		{
			Name: "PingCap",
			Info: "PingCap 😭 我的PingCap😭 你带我走吧 \n\n" +
				"PingCAP 成立于 2015 年，是一家企业级开源分布式数据库厂商，提供包括开源分布式数据库产品、解决方案与咨询、技术支持与培训认证服务，致力于为全球行业用户提供稳定高效、安全可靠、开放兼容的新型数据服务平台，解放企业生产力，加速企业数字化转型升级。在帮助企业释放增长空间的同时，也提供了一份具有高度可参考性的开源建设实践样本。\n\n" +
				"由 PingCAP 创立的分布式关系型数据库 TiDB，为企业关键业务打造，具备「分布式强一致事务、在线弹性水平扩展、故障自恢复的高可用、跨数据中心多活」等企业级核心特性，帮助企业最大化发挥数据价值，充分释放企业增长空间。",
			Image: "img_v2_4efdd859-8f93-4ed5-8214-89a4af3121dg",
			Url:   "https://pingcap.com/zh/join-us/",
		},
	}
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(companyList))
	return companyList[index]
}

func getRandomImage() string {
	imageList := []string{"img_v2_a5f42425-d2b4-4ba4-994f-305ea8d369bg",
		"img_v2_356dbc68-025b-452b-b7aa-1344d1fe1c8g",
		"img_v2_c5bfd0e0-9413-49e0-865f-c34a15613dfg",
		"img_v2_0151edcb-a456-4a85-899e-b3477d3e69eg"}
	str := "www.baidu.com"
	strings.HasPrefix(str, "")
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(imageList))
	return imageList[index]
}
