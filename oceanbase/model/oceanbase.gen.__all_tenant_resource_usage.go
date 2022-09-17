package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantResourceUsageMgr struct {
	*_BaseMgr
}

// AllTenantResourceUsageMgr open func
func AllTenantResourceUsageMgr(db *gorm.DB) *_AllTenantResourceUsageMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantResourceUsageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantResourceUsageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_resource_usage"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantResourceUsageMgr) GetTableName() string {
	return "__all_tenant_resource_usage"
}

// Reset 重置gorm会话
func (obj *_AllTenantResourceUsageMgr) Reset() *_AllTenantResourceUsageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantResourceUsageMgr) Get() (result AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantResourceUsageMgr) Gets() (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantResourceUsageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllTenantResourceUsageMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllTenantResourceUsageMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllTenantResourceUsageMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithReportTime report_time获取
func (obj *_AllTenantResourceUsageMgr) WithReportTime(reportTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["report_time"] = reportTime })
}

// WithCPUQuota cpu_quota获取
func (obj *_AllTenantResourceUsageMgr) WithCPUQuota(cpuQuota float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_quota"] = cpuQuota })
}

// WithCPUQuotaUsed cpu_quota_used获取
func (obj *_AllTenantResourceUsageMgr) WithCPUQuotaUsed(cpuQuotaUsed float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_quota_used"] = cpuQuotaUsed })
}

// WithMemQuota mem_quota获取
func (obj *_AllTenantResourceUsageMgr) WithMemQuota(memQuota float64) Option {
	return optionFunc(func(o *options) { o.query["mem_quota"] = memQuota })
}

// WithMemQuotaUsed mem_quota_used获取
func (obj *_AllTenantResourceUsageMgr) WithMemQuotaUsed(memQuotaUsed float64) Option {
	return optionFunc(func(o *options) { o.query["mem_quota_used"] = memQuotaUsed })
}

// WithDiskQuota disk_quota获取
func (obj *_AllTenantResourceUsageMgr) WithDiskQuota(diskQuota float64) Option {
	return optionFunc(func(o *options) { o.query["disk_quota"] = diskQuota })
}

// WithDiskQuotaUsed disk_quota_used获取
func (obj *_AllTenantResourceUsageMgr) WithDiskQuotaUsed(diskQuotaUsed float64) Option {
	return optionFunc(func(o *options) { o.query["disk_quota_used"] = diskQuotaUsed })
}

// WithIopsQuota iops_quota获取
func (obj *_AllTenantResourceUsageMgr) WithIopsQuota(iopsQuota float64) Option {
	return optionFunc(func(o *options) { o.query["iops_quota"] = iopsQuota })
}

// WithIopsQuotaUsed iops_quota_used获取
func (obj *_AllTenantResourceUsageMgr) WithIopsQuotaUsed(iopsQuotaUsed float64) Option {
	return optionFunc(func(o *options) { o.query["iops_quota_used"] = iopsQuotaUsed })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantResourceUsageMgr) GetByOption(opts ...Option) (result AllTenantResourceUsage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantResourceUsageMgr) GetByOptions(opts ...Option) (results []*AllTenantResourceUsage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantResourceUsageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantResourceUsage, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromTenantID(tenantID int64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromSvrIP(svrIP string) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromSvrPort(svrPort int64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromReportTime 通过report_time获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromReportTime(reportTime time.Time) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`report_time` = ?", reportTime).Find(&results).Error

	return
}

// GetBatchFromReportTime 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromReportTime(reportTimes []time.Time) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`report_time` IN (?)", reportTimes).Find(&results).Error

	return
}

// GetFromCPUQuota 通过cpu_quota获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromCPUQuota(cpuQuota float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`cpu_quota` = ?", cpuQuota).Find(&results).Error

	return
}

// GetBatchFromCPUQuota 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromCPUQuota(cpuQuotas []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`cpu_quota` IN (?)", cpuQuotas).Find(&results).Error

	return
}

// GetFromCPUQuotaUsed 通过cpu_quota_used获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromCPUQuotaUsed(cpuQuotaUsed float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`cpu_quota_used` = ?", cpuQuotaUsed).Find(&results).Error

	return
}

// GetBatchFromCPUQuotaUsed 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromCPUQuotaUsed(cpuQuotaUseds []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`cpu_quota_used` IN (?)", cpuQuotaUseds).Find(&results).Error

	return
}

// GetFromMemQuota 通过mem_quota获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromMemQuota(memQuota float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`mem_quota` = ?", memQuota).Find(&results).Error

	return
}

// GetBatchFromMemQuota 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromMemQuota(memQuotas []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`mem_quota` IN (?)", memQuotas).Find(&results).Error

	return
}

// GetFromMemQuotaUsed 通过mem_quota_used获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromMemQuotaUsed(memQuotaUsed float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`mem_quota_used` = ?", memQuotaUsed).Find(&results).Error

	return
}

// GetBatchFromMemQuotaUsed 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromMemQuotaUsed(memQuotaUseds []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`mem_quota_used` IN (?)", memQuotaUseds).Find(&results).Error

	return
}

// GetFromDiskQuota 通过disk_quota获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromDiskQuota(diskQuota float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`disk_quota` = ?", diskQuota).Find(&results).Error

	return
}

// GetBatchFromDiskQuota 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromDiskQuota(diskQuotas []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`disk_quota` IN (?)", diskQuotas).Find(&results).Error

	return
}

// GetFromDiskQuotaUsed 通过disk_quota_used获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromDiskQuotaUsed(diskQuotaUsed float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`disk_quota_used` = ?", diskQuotaUsed).Find(&results).Error

	return
}

// GetBatchFromDiskQuotaUsed 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromDiskQuotaUsed(diskQuotaUseds []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`disk_quota_used` IN (?)", diskQuotaUseds).Find(&results).Error

	return
}

// GetFromIopsQuota 通过iops_quota获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromIopsQuota(iopsQuota float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`iops_quota` = ?", iopsQuota).Find(&results).Error

	return
}

// GetBatchFromIopsQuota 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromIopsQuota(iopsQuotas []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`iops_quota` IN (?)", iopsQuotas).Find(&results).Error

	return
}

// GetFromIopsQuotaUsed 通过iops_quota_used获取内容
func (obj *_AllTenantResourceUsageMgr) GetFromIopsQuotaUsed(iopsQuotaUsed float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`iops_quota_used` = ?", iopsQuotaUsed).Find(&results).Error

	return
}

// GetBatchFromIopsQuotaUsed 批量查找
func (obj *_AllTenantResourceUsageMgr) GetBatchFromIopsQuotaUsed(iopsQuotaUseds []float64) (results []*AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`iops_quota_used` IN (?)", iopsQuotaUseds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantResourceUsageMgr) FetchByPrimaryKey(tenantID int64, svrIP string, svrPort int64) (result AllTenantResourceUsage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantResourceUsage{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, svrIP, svrPort).First(&result).Error

	return
}
