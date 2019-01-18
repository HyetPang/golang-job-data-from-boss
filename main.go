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
	"net/http"
	"strconv"
	"strings"
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
	resp, err := http.Get(fmt.Sprintf(domain+"/job_detail/?query=%s&scity=101270100", query))
	if err != nil {
		panic(err)
	}
	defer func() {
		e := resp.Body.Close()
		if e != nil {
			panic(e)
		}
	}()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#main > div > div.job-list > ul > li > div > div.info-primary > h3 > a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if i == 0 {
			href, ok := s.Attr("href")
			if ok {
				getBossData(domain + href)
			} else {
				log.Fatal("第" + strconv.Itoa(i) + "不存在属性href")
			}
		}
	})

}

func getBossData(url string) {
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
		return
	}
	job.SalaryRange = document.Find("#main > div.job-banner > div > div > div.info-primary > div.name > span").Text()
	job.EnterpriseName = document.Find("#main > div.job-banner > div > div > div.info-company > h3 > a").Text()
	job.Category = document.Find("#main > div.job-banner > div > div > div.info-company > p > a").Text()
	var tags strings.Builder
	document.Find("#main > div.job-banner > div > div > div.info-primary > div.job-tags > span").Each(func(i int, selection *goquery.Selection) {
		tags.WriteString(selection.Text() + ",")
	})
	job.JobTags = tags.String()[:tags.Len()-1]
	job.HrHeadImg, _ = document.Find("#main > div.job-box > div > div.job-detail > div.detail-figure > img").Attr("src")
	job.HrNickname = document.Find("#main > div.job-box > div > div.job-detail > div.detail-op > h2").Text()
	job.JobDetails = document.Find("#main > div.job-box > div > div.job-detail > div.detail-content > div:nth-child(1) > div").Text()
	cityAndWorkYearsAndEducation := document.Find("#main > div.job-banner > div > div > div.info-primary > p")
	node := cityAndWorkYearsAndEducation.Nodes[0].FirstChild
	job.City = strings.Replace(node.Data, "城市：", "", -1)
	workYear := node.NextSibling.NextSibling.Data
	job.WorkYears = strings.Replace(workYear, "经验：", "", -1)
	education := node.NextSibling.NextSibling.NextSibling.NextSibling.Data
	job.Education = strings.Replace(education, "学历：", "", -1)
	job.EnterpriseAddress = document.Find("#main > div.job-box > div > div.job-detail > div.detail-content > div:nth-child(5) > div > div.location-address").Text()
	enterpriceInfoNode := document.Find("#main > div.job-banner > div > div > div.info-company > p")
	job.EnterpriseScale = enterpriceInfoNode.Nodes[0].FirstChild.NextSibling.NextSibling.Data
	job.JobName = document.Find("#main > div.job-banner > div > div > div.info-primary > div.name > h1").Text()
	job.TrimSpaceAndEnter()
	//rows, _ := mysqlDb.DB().Query("select uuid()")
	//for rows.Next() {
	//	rows.Scan(&job.Id)
	//}
	job.Id = uuid.Must(uuid.NewV4()).String()

	bytes, err := json.Marshal(job)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(bytes))
	mysqlDb.Create(&job)
}
