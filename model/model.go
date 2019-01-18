package model

import "strings"

type ZhipinData struct {
	Category          string `json:"category",gorm:"Column:category"`
	City              string `json:"city",gorm:"Column:city"`
	Education         string `json:"education",gorm:"Column:education"`
	EnterpriseAddress string `json:"enterpriseAddress",gorm:"Column:enterprise_address"`
	EnterpriseName    string `json:"enterpriseName",gorm:"Column:enterprise_name"`
	EnterpriseScale   string `json:"enterpriseScale",gorm:"Column:enterprise_scale"`
	HrHeadImg         string `json:"hrHeadImg",gorm:"Column:hr_headImg"`
	HrNickname        string `json:"hrNickname",gorm:"Column:hr_nickname"`
	Id                string `json:"id",gorm:"Column:id"`
	JobDetails        string `json:"jobDetails",gorm:"Column:job_details"`
	JobName           string `json:"jobName",gorm:"Column:job_name"`
	JobTags           string `json:"jobTags",gorm:"Column:job_tags"`
	SalaryRange       string `json:"salaryRange",gorm:"Column:salary_range"`
	WorkYears         string `json:"workYears",gorm:"Column:work_years"`
}

func (ZhipinData) TableName() string {
	return "zhipin_data"
}

// 去掉空格，换行符等
func (z *ZhipinData) TrimSpaceAndEnter() {
	z.Category = strings.Trim(strings.TrimSpace(z.Category), "\r")
	z.City = strings.Trim(strings.TrimSpace(z.City), "\r")
	z.Education = strings.Trim(strings.TrimSpace(z.Education), "\r")
	z.EnterpriseAddress = strings.Trim(strings.TrimSpace(z.EnterpriseAddress), "\r")
	z.EnterpriseName = strings.Trim(strings.TrimSpace(z.EnterpriseName), "\r")
	z.EnterpriseScale = strings.Trim(strings.TrimSpace(z.EnterpriseScale), "\r")
	z.HrHeadImg = strings.Trim(strings.TrimSpace(z.HrHeadImg), "\r")
	z.HrNickname = strings.Trim(strings.TrimSpace(z.HrNickname), "\r")
	z.Id = strings.Trim(strings.TrimSpace(z.Id), "\r")
	z.JobDetails = strings.Trim(strings.TrimSpace(z.JobDetails), "\r")
	z.JobName = strings.Trim(strings.TrimSpace(z.JobName), "\r")
	z.JobTags = strings.Trim(strings.TrimSpace(z.JobTags), "\r")
	z.SalaryRange = strings.Trim(strings.TrimSpace(z.SalaryRange), "\r")
	z.WorkYears = strings.Trim(strings.TrimSpace(z.WorkYears), "\r")

}
