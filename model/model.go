package model

import "strings"

type ZhipinData struct {
	Id                string `json:"id",gorm:"id"`
	EnterpriseName    string `json:"enterpriseName",gorm:"enterprise_name"`
	EnterpriseScale   string `json:"enterpriseScale",gorm:"enterprise_scale"`
	EnterpriseAddress string `json:"enterpriseAddress",gorm:"enterprise_address"`
	WorkYears         string `json:"workYears",gorm:"work_years"`
	SalaryRange       string `json:"salaryRange",gorm:"salary_range"`
	Category          string `json:"category",gorm:"category"`
	Education         string `json:"education",gorm:"education"`
	JobName           string `json:"jobName",gorm:"job_name"`
	JobDetails        string `json:"jobDetails",gorm:"job_details"`
	HrNickname        string `json:"hrNickname",gorm:"hr_nickname"`
	HrHeadImg         string `json:"hrHeadImg",gorm:"hr_head_img"`
	City              string `json:"city",gorm:"city"`
	JobTags           string `json:"jobTags",gorm:"job_tags"`
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
