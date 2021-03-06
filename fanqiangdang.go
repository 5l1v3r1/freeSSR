package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var templet string = "# SSR 账号每日更新 \n> 数据来源: [翻墙党](https://fanqiangdang.com/) \n----------------------------------------------\n## 更新日期：%s \n***食用方法：复制下面的节点到SSR客户端去重添加即可***\n\n %s"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("fanqiangdang.com"),
	)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		comments := e.ChildText("#postmessage_379")
		reg := regexp.MustCompile(`ssr(.*?)[\n]`)
		ssrAddr := reg.FindAllString(comments, -1)
		fmt.Print(ssrAddr)
		ssrAddr2Str := strings.Replace(strings.Trim(fmt.Sprint(ssrAddr), "[]"), " ", "\n", -1)
		file, error := os.OpenFile("C:\\Users\\Administrator\\Desktop\\Bean\\Go\\src\\github.com\\BeanWei\\SSR\\README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
		defer file.Close()
		if error != nil {
			fmt.Println(error)
		}
		t := time.Now().Format("2006-01-02 15:04:05")
		strText := fmt.Sprintf(templet, t, ssrAddr2Str)
		file.WriteString(strText)
		// fmt.Print(">>> Git同步···\n")
		// autogit("git add C:\\Users\\Administrator\\Desktop\\Bean\\Go\\src\\github.com\\BeanWei\\SSR\\README.md")
		// autogit("git commit -m \"ssr节点分享，每日更新\"")
		// autogit("git remote add origin git@github.com:BeanWei/freeSSR.git")
		// autogit("git pull origin master")
		// autogit("git push origin master")
		// fmt.Print("\n>>> Done")
	})

	c.Visit("https://fanqiangdang.com/forum.php?mod=viewthread&tid=84&page=1&extra=#pid379")
}

// func autogit(strCmd string) {
// 	stout, err := exec.Command("cmd", "/C", strCmd).Output()
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	fmt.Print(string(stout))
// }
