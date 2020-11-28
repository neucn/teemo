package main

import (
	"errors"
	"fmt"
	"github.com/neucn/neugo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"time"
)

type config struct {
	Notifiers []notifierConfig
	Tasks     []taskConfig

	notifierMap map[string]notifier
}

// 根据配置生成notifier
func (c *config) GenerateNotifiers() error {
	if len(c.Notifiers) == 0 {
		return errors.New("未设置 Notifier")
	}
	c.notifierMap = map[string]notifier{}
	for _, nc := range c.Notifiers {
		nc.Type = strings.ToLower(nc.Type)
		// nf = notifier factory
		nf, ok := notifiers[nc.Type]
		if !ok {
			return errors.New("没有该类型的 Notifier: " + nc.Type)
		}
		// n = notifier
		n, err := nf(nc.Option)
		if err != nil {
			return err
		}
		c.notifierMap[nc.Name] = n
	}
	return nil
}

const gradeUrl = "http://219.216.96.4/eams/teach/grade/course/person!search.action?semesterId=0"

func (c *config) GenerateTasks() ([]*task, error) {
	if len(c.Tasks) == 0 {
		return nil, errors.New("未设置 Task")
	}
	var result []*task
	for _, u := range c.Tasks {
		if u.Interval < 0 {
			return nil, errors.New(fmt.Sprintf("时间间隔必须大于0: %d", u.Interval))
		}
		if u.Interval == 0 {
			u.Interval = 60
		}
		t := &task{
			Username:  u.Username,
			Password:  u.Password,
			Interval:  time.Duration(u.Interval) * time.Second,
			Notifiers: []notifier{},
		}
		// nn = notifier name
		for _, nn := range u.Use {
			n, ok := c.notifierMap[nn]
			if !ok {
				return nil, errors.New("引用了不存在的 Notifier: " + nn)
			}
			t.Notifiers = append(t.Notifiers, n)
		}

		if u.Webvpn {
			t.Platform = neugo.WebVPN
			t.URL = neugo.EncryptToWebVPN(gradeUrl)
		} else {
			t.Platform = neugo.CAS
			t.URL = gradeUrl
		}

		session := neugo.NewSession()
		session.Timeout = 10 * time.Second
		t.Session = session

		result = append(result, t)
	}
	return result, nil
}

type notifierConfig struct {
	Name, Type string
	Option     map[string]interface{}
}

type taskConfig struct {
	Username, Password string
	Webvpn             bool
	Use                []string
	Interval           int
}

func parseConfig(configPath string) ([]*task, error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	c := &config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	err = c.GenerateNotifiers()
	if err != nil {
		return nil, err
	}
	tasks, err := c.GenerateTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
