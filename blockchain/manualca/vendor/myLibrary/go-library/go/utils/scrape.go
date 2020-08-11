/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-17 09:14 
# @File : scraptch.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	keyUserId    = 389102
	detailPrefix = "https://tieba.baidu.com"
)

var (
	mainList []UserMainInfo
)
// key: 389102
// https://tieba.baidu.com/f?kw=%E5%92%8C%E5%B9%B3%E7%B2%BE%E8%8B%B1&ie=utf-8&pn=100

func scrape() {
	end := make([]Result, 0)
	baseUrl := "https://tieba.baidu.com/f?kw=%E5%92%8C%E5%B9%B3%E7%B2%BE%E8%8B%B1&ie=utf-8&pn="
	for i := 0; i < 10; i++ {
		url := baseUrl + strconv.Itoa(i*50)
		log.Println("开始爬取url=" + url + "的数据")
		res, err := http.Get(url)
		if nil != err {
			log.Println("打开失败,url=" + url)
			continue
		}
		if res.StatusCode != 200 {
			log.Printf("status code error: %d %s ,url=%s \n", res.StatusCode, res.Status, url)
			continue
		}
		resp := handlerResp(res)
		defer res.Body.Close()
		end = append(end, resp)
	}

	printEnd(end)
}

func printEnd(results []Result) {
	log.Println("hi ")
	for i := 0; i < len(results); i++ {
		log.Println(results[i].String())
	}
	log.Println("you know is an end")

}

type UserMainInfo struct {
	UserId int `json:"user_id"`
	Href   string
}

type UserIdInfo struct {
	UserId int `json:"user_id"`
}

type UserDetailInfo struct {
	UserIdInfo
	Detail string
	Href   string
}

type Result struct {
	MainList   []UserMainInfo
	DetailList []UserDetailInfo
}

func (end Result) String() string {
	str := strings.Builder{}
	str.WriteString("[ ")
	for i := 0; i < len(end.MainList); i++ {
		m := end.MainList[i]
		str.WriteString("href :" + m.Href + " =")
	}
	str.WriteString(" ]")

	str.WriteString(" { ")
	for i := 0; i < len(end.DetailList); i++ {
		detailInfo := end.DetailList[i]
		str.WriteString("href=" + detailInfo.Href + "  detail=" + detailInfo.Detail + "  ")
	}
	str.WriteString(" } ")

	return str.String()
}

func (this UserDetailInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(" [   id=" + strconv.Itoa(this.UserId) + "   ")
	builder.WriteString("detail=" + this.Detail + "  ] ")
	return builder.String()
}

func handlerResp(res *http.Response) Result {
	var result Result
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".j_thread_list").Each(func(i int, s *goquery.Selection) {
		v, b := s.Find(".tb_icon_author").Attr("data-field")
		if b {
			href, hrefExist := s.Find(".threadlist_title").
				Find("a").Attr("href")
			if hrefExist {
				var user UserIdInfo
				err := json.Unmarshal([]byte(v), &user)
				if nil != err {
					log.Println("解析成id失败")
				} else if user.UserId == keyUserId {
					main := UserMainInfo{}
					main.UserId = user.UserId
					main.Href = href
					log.Printf("抓到你了 href=%s,userId=%s \n", href, v)
					result.MainList = append(result.MainList, main)
				} else {
					// 抓取详情里的内容
					path := detailPrefix + href
					details, e := scrapeDetail(path)
					if nil != e {
						result.DetailList = append(result.DetailList, details...)
					}
				}
			}
		}
	})
	return result
}

func scrapeDetail(url string) ([]UserDetailInfo, error) {
	var (
		detail []UserDetailInfo
	)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println("开始爬取详情,详情url=" + url)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".p_postlist").
		Find(".l_post").Each(func(i int, area *goquery.Selection) {
		val, exists := area.Find(".d_author").Find(".d_name").Attr("data-field")
		var user UserDetailInfo
		var e error
		maybe := false
		if !exists {
		} else {
			e = json.Unmarshal([]byte(val), &user)
			text := area.Find(".d_post_content_main").Find(".d_post_content").Text()
			maybe = strings.Contains(text, "95")
			user.Detail = text
		}
		if nil != e {
			log.Println("发生错误:" + e.Error())
		} else if user.UserId == keyUserId {
			log.Println("抓到你了:", user.String())
			user.Href = url
			detail = append(detail, user)
		} else if maybe {
			log.Println("可能抓到你了:", user.String())
			user.Href = url
			detail = append(detail, user)
		}
	})
	return detail, err
}
