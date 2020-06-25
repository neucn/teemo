package main

import (
	"flag"
	"fmt"
	"github.com/neucn/neugo"
	"github.com/unbyte/beeep"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var (
	u, p string
	f    int
	v    bool

	gpa, reqUrl string
	client      *http.Client

	upPath, downPath string
)

func init() {
	flag.StringVar(&u, "u", "", "学号")
	flag.StringVar(&p, "p", "", "密码")
	flag.IntVar(&f, "f", 60, "频率 单位秒")
	flag.BoolVar(&v, "v", false, "使用webvpn")
	flag.Usage = usage

	upPath, downPath = getImgPath()
}

func main() {
	flag.Parse()
	beeep.AppID = "Teemo"

	go func() {
		// 登陆
		client = login()
		fmt.Println("登陆成功")
		for {
			handle()
			time.Sleep(time.Duration(f) * time.Second)
		}
	}()

	select {}
}

func login() *http.Client {
	session := neugo.NewSession()
	session.Timeout = 10 * time.Second

	var platform neugo.Platform
	if v {
		platform = neugo.WebVPN
		reqUrl = "https://219-216-96-4.webvpn.neu.edu.cn/eams/teach/grade/course/person!search.action?semesterId=0"
	} else {
		platform = neugo.CAS
		reqUrl = "http://219.216.96.4/eams/teach/grade/course/person!search.action?semesterId=0"
	}
	_, err := neugo.Use(session).WithAuth(u, p).On(platform).LoginService(reqUrl)
	if err != nil {
		fmt.Println("登陆失败: ", err.Error())
		os.Exit(1)
	}
	return session
}

var gpaExp = regexp.MustCompile(`<div>总平均绩点：(.+?)</div>`)

func getGPA() string {
	resp, err := client.Get(reqUrl)
	if err != nil {
		fmt.Println("网络错误")
		return "获取失败"
	}
	body := readBody(resp)

	gpa := gpaExp.FindAllStringSubmatch(body, -1)
	if len(gpa) < 1 {
		return "获取失败"
	}
	return gpa[0][1]
}

func handle() {
	newGPA := getGPA()
	fmt.Printf("%10s\t绩点: %s\n", time.Now().Format("2006-01-02 15:04:05"), newGPA)
	if newGPA != gpa {
		if newGPA == "获取失败" {
			return
		}

		if len(gpa) < 1 {
			gpa = newGPA
			return
		}

		n, _ := strconv.ParseFloat(newGPA, 32)
		g, _ := strconv.ParseFloat(gpa, 32)
		diff := n - g

		var err error
		if diff > 0 {
			err = beeep.Notify("绩点变高啦", fmt.Sprintf("绩点上升了\t%.4f\n当前绩点\t%s", diff, newGPA), upPath)
		} else {
			err = beeep.Notify("绩点降低了", fmt.Sprintf("绩点降低了\t%.4f\n当前绩点\t%s", diff, newGPA), downPath)
		}

		if err != nil {
			fmt.Println("推送提示失败")
		}

		gpa = newGPA
	}
}

func readBody(resp *http.Response) (body string) {
	res, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return string(res)
}

func usage() {
	fmt.Println(`
teemo -u 学号 -p 密码
    如 teemo -u 2018xxxx -p abcdefg
teemo -u 学号 -p 密码 -f 监控频率(单位秒)
    如 teemo -u 2018xxxx -p abcdefg -f 60
teemo -u 学号 -p 密码 -v 使用webvpn
    如 teemo -u 2018xxxx -p abcdefg -v
若不指定f，默认60`)
}

func getImgPath() (upPath, downPath string) {
	p, _ := os.Executable()
	absPath, _ := filepath.Abs(p)
	dir := filepath.Dir(absPath)
	return path.Join(dir, "img", "up.png"), path.Join(dir, "img", "down.png")
}
