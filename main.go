package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin/json"
	"github.com/hyetpang/golang-job-data-from-boss/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const domain = "https://www.zhipin.com"

var mysqlDb *gorm.DB

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/boss_data?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	mysqlDb = db
	// 获取信息
	getData("Golang")
	defer mysqlDb.Close()
}

func getData(query string) {
	url := fmt.Sprintf(domain+"/c101270100/?query=%s&scity=101270100&page=1", query)
	var n = 1000
	var loopN = 1
	for {
		sleepTime := rand.Intn(n)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		var detailsN = n
		doc.Find("#main > div > div.job-list > ul > li > div > div.info-primary > h3 > a").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			if loopN < 3 {
				detailsN += 100
			} else {
				detailsN += 1300
			}
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			href, ok := s.Attr("href")
			if ok {
				// 换ip
				notExec := getBossData(domain + href)
				if !notExec {
					panic("被封了,延迟是:" + strconv.Itoa(detailsN))
				}
			} else {
				log.Fatal("第" + strconv.Itoa(i) + "不存在属性href")
			}
			sleepTime = rand.Intn(detailsN)
		})
		if loopN > 3 {
			n += 1500
		} else {
			n += 900

		}
		// 关闭流
		e := resp.Body.Close()
		if e != nil {
			panic(e)
		}
		href, ok := doc.Find("#main > div > div.job-list > div.page > a.next").Attr("href")
		if !ok {
			// 不存在
			fmt.Println("已经不存在下一页了!")
			break
		}
		if href != "javascript:;" {
			url = domain + href
		} else {
			break
		}
		loopN++
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
}

func getBossData(url string) bool {
	var job model.ZhipinData
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		e := resp.Body.Close()
		if e != nil {
			panic(e)
		}
	}()
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return true
	}
	_, has := document.Find("#wrap > div > form").Attr("action")
	if has {
		return false
	}
	job.SalaryRange = document.Find("#main > div.job-banner > div > div > div.info-primary > div.name > span").Text()
	job.EnterpriseName = document.Find("#main > div.job-banner > div > div > div.info-company > h3 > a").Text()
	job.Category = document.Find("#main > div.job-banner > div > div > div.info-company > p > a").Text()
	var tags strings.Builder
	document.Find("#main > div.job-banner > div > div > div.info-primary > div.job-tags > span").Each(func(i int, selection *goquery.Selection) {
		tags.WriteString(selection.Text() + ",")
	})
	if len(tags.String()) > 0 {
		job.JobTags = tags.String()[:tags.Len()-1]
	}
	job.HrHeadImg, _ = document.Find("#main > div.job-box > div > div.job-detail > div.detail-figure > img").Attr("src")
	job.HrNickname = document.Find("#main > div.job-box > div > div.job-detail > div.detail-op > h2").Text()
	job.JobDetails = document.Find("#main > div.job-box > div > div.job-detail > div.detail-content > div:nth-child(1) > div").Text()
	cityAndWorkYearsAndEducation := document.Find("#main > div.job-banner > div > div > div.info-primary > p")
	if len(cityAndWorkYearsAndEducation.Nodes) > 0 {
		node := cityAndWorkYearsAndEducation.Nodes[0].FirstChild
		job.City = strings.Replace(node.Data, "城市：", "", -1)
		workYear := node.NextSibling.NextSibling.Data
		job.WorkYears = strings.Replace(workYear, "经验：", "", -1)
		education := node.NextSibling.NextSibling.NextSibling.NextSibling.Data
		job.Education = strings.Replace(education, "学历：", "", -1)
	}
	job.EnterpriseAddress = document.Find("#main > div.job-box > div > div.job-detail > div.detail-content > div:nth-child(5) > div > div.location-address").Text()
	enterpriceInfoNode := document.Find("#main > div.job-banner > div > div > div.info-company > p")
	if len(enterpriceInfoNode.Nodes) > 0 {
		job.EnterpriseScale = enterpriceInfoNode.Nodes[0].FirstChild.NextSibling.NextSibling.Data
	}
	job.JobName = document.Find("#main > div.job-banner > div > div > div.info-primary > div.name > h1").Text()
	job.TrimSpaceAndEnter()
	job.Id = uuid.Must(uuid.NewV4()).String()

	bytes, err := json.Marshal(job)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println(string(bytes))
	mysqlDb.Create(&job)
	return true
}
