package main

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/unbyte/beeep"
	"gopkg.in/gomail.v2"
	"time"
)

type notifier interface {
	Notify(username string, newGPA, oldGPA float64) error
}

type notifierFactory func(option map[string]interface{}) (notifier, error)

var notifiers = map[string]notifierFactory{
	"mail":  newMail,
	"toast": newToast,
}

type mail struct {
	Host              string
	Account, Password string
	Port              int
	Receiver          string
	dialer            *gomail.Dialer
}

func (m *mail) Notify(username string, newGPA, oldGPA float64) error {
	diff := newGPA - oldGPA
	var title, content string
	if diff > 0 {
		title = fmt.Sprintf("%s 绩点变高啦", username)
		content = fmt.Sprintf("绩点上升了\t%.4f\n当前绩点\t%.4f", diff, newGPA)
	} else {
		title = fmt.Sprintf("%s 绩点下降啦", username)
		content = fmt.Sprintf("绩点降低了\t%.4f\n当前绩点\t%.4f", -diff, newGPA)
	}
	content += `<br/>
<a target="_blank" href="http://219.216.96.4/eams/teach/grade/course/person!historyCourseGrade.action?projectType=MAJOR">
查看所有课程绩点</a>`

	return m.send(title, content)
}

func (m *mail) send(title, content string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.Account)
	msg.SetHeader("To", m.Receiver)
	msg.SetHeader("Subject", title)
	msg.SetBody("text/html", content)

	if err := m.dialer.DialAndSend(msg); err != nil {
		return errors.New("邮件发送失败: " + err.Error())
	}
	return nil
}

func newMail(option map[string]interface{}) (notifier, error) {
	m := &mail{}
	err := mapstructure.Decode(option, m)
	if err != nil {
		return nil, err
	}
	if len(m.Account) == 0 || len(m.Receiver) == 0 || m.Port == 0 {
		return nil, errors.New("邮箱配置有误")
	}
	m.dialer = gomail.NewDialer(m.Host, m.Port, m.Account, m.Password)
	if test, ok := option["test"].(bool); ok && test {
		if err := m.send("GPA 监听 邮件测试", time.Now().Format("2006-01-02 15:04:05")); err != nil {
			return nil, errors.New("测试未通过: " + err.Error())
		}
	}
	return m, nil
}

type toast struct{}

var (
	upPath, downPath = getImgPath()
)

func (t *toast) Notify(username string, newGPA, oldGPA float64) error {
	diff := newGPA - oldGPA
	if diff > 0 {
		return beeep.Notify("绩点变高啦", fmt.Sprintf("%s 绩点上升了\t%.4f\n当前绩点\t%.4f", username, diff, newGPA), upPath)
	}
	return beeep.Notify("绩点降低了", fmt.Sprintf("%s 绩点降低了\t%.4f\n当前绩点\t%.4f", username, -diff, newGPA), downPath)
}

func newToast(_ map[string]interface{}) (notifier, error) {
	return &toast{}, nil
}
