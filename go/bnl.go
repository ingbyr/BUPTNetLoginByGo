package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// BUPTNet
type BUPTNet struct {
	loginUrl  string
	logoutUrl string
}

var lines = map[string]string{
	"xyw": "",
	"lt":  "CUC-BRAS",
	"yd":  "CMCC-BRAS",
	"dx":  "CT-BRAS",
}

const version string = "0.0.1"

func (net *BUPTNet) login(username, password, line string) {
	lineOP, ok := lines[line]
	if !ok {
		fmt.Println("不可用线路： " + line)
		fmt.Printf("可用参数 xyw（校园网）、lt（联通）、yd（移动）、dx（电信）")
		return
	}
	params := url.Values{
		"user": {username},
		"pass": {password},
		"line": {lineOP},
	}

	resp, err := http.PostForm(net.loginUrl, params)
	if err != nil {
		fmt.Println("登陆网关失败")
	}
	defer resp.Body.Close()

	parseLoginOutput(resp.Body)
}

func (net *BUPTNet) logtout() {
	resp, err := http.Get(net.logoutUrl)
	if err != nil {
		fmt.Println("注销失败")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		println("成功注销北邮网关")
	}
}

func parseLoginOutput(body io.Reader) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		println("parse error")
	}
	doc.Find(".login-panel .notice-content").Each(func(i int, s *goquery.Selection) {
		scanner := bufio.NewScanner(strings.NewReader(s.Text()))
		for scanner.Scan() {
			fmt.Println(strings.TrimSpace(scanner.Text()))
		}
	})
}

func main() {
	net := BUPTNet{
		"http://ngw.bupt.edu.cn/login",
		"http://ngw.bupt.edu.cn/logout",
	}

	line := flag.String("l", "", "线路选择，可用参数 xyw（校园网）、lt（联通）、yd（移动）、dx（电信）")
	username := flag.String("u", "", "校园网账户名称")
	password := flag.String("p", "", "校园网账户密码")
	isLogout := flag.Bool("lo", false, "注销北邮校园网网关")
	showVersion := flag.Bool("v", false, "版本信息")

	flag.Parse()

	if len(*line) > 0 {
		net.login(*username, *password, *line)
	} else if *isLogout {
		net.logtout()
	} else if *showVersion {
		println("Ver " + version)
	} else {
		flag.Usage()
	}
}
