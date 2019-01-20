create table zhipin_data
(
	id varchar(64) not null comment '主键
',
	enterprise_name varchar(100) null comment '企业名',
	enterprise_scale varchar(100) null comment '企业规模',
	enterprise_address varchar(255) null comment '公司地址',
	work_years varchar(10) null comment '工作年限',
	salary_range varchar(10) null comment '薪水范围',
	category varchar(10) null comment '类别',
	education varchar(10) null comment '学历',
	job_name varchar(100) null comment '职位名字',
	job_details text null comment '职位描述',
	hr_nickname varchar(25) null comment 'hr昵称',
	hr_head_img varchar(255) null comment 'hr头像',
	city varchar(10) null comment '城市',
	job_tags varchar(100) null comment '标签'
)
comment '获取的boss招聘数据';

