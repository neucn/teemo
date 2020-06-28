package main

import (
	"errors"
	"github.com/neucn/neugo"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type task struct {
	Username, Password string
	Interval           time.Duration
	Session            *http.Client
	Notifiers          []notifier

	// 请求的 url
	URL      string
	Platform neugo.Platform

	// 存储上一次获取的 GPA
	gpa float64

	// 重试次数
	retry byte
}

var gpaExp = regexp.MustCompile(`<div>总平均绩点：(.+?)</div>`)

func (t *task) getGPA() (float64, error) {
	resp, err := t.Session.Get(t.URL)
	if err != nil {
		return 0, err
	}
	body := readBody(resp)

	gpa := gpaExp.FindAllStringSubmatch(body, -1)
	if len(gpa) < 1 {
		return 0, errors.New("提取绩点失败")
	}
	return strconv.ParseFloat(gpa[0][1], 32)
}

func (t *task) handle() {
	newGPA, err := t.getGPA()
	if err != nil {
		log.Printf("%s\t获取失败: %s\n", t.Username, err.Error())
		if t.retry < 2 {
			t.retry += 1
			log.Printf("%s\t尝试重新登陆 第 %d 次\n", t.Username, t.retry)

			time.Sleep(3 * time.Second)

			_ = t.login()
			t.handle()
		}
		return
	}
	log.Printf("%s\t绩点: %.4f\n", t.Username, newGPA)
	if newGPA != t.gpa {
		if t.gpa == 0 {
			t.gpa = newGPA
			return
		}

		for _, n := range t.Notifiers {
			if err := n.Notify(t.Username, newGPA, t.gpa); err != nil {
				log.Printf("%s\t推送通知失败: %s\n", t.Username, err.Error())
			}
		}

		t.gpa = newGPA
	}
}

func (t *task) Start() error {
	err := t.login()
	if err != nil {
		log.Printf("%s\t登陆失败: %s\n", t.Username, err.Error())
		return err
	}
	log.Printf("%s\t登陆成功\n", t.Username)
	go func() {
		for {
			t.retry = 0
			t.handle()
			time.Sleep(t.Interval)
		}
	}()
	return nil
}

func (t *task) login() error {
	_, err := neugo.Use(t.Session).WithAuth(t.Username, t.Password).On(t.Platform).LoginService(t.URL)
	if err != nil {
		return err
	}
	return nil
}
