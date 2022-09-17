package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllUnitLoadHistoryMgr struct {
	*_BaseMgr
}

// AllUnitLoadHistoryMgr open func
func AllUnitLoadHistoryMgr(db *gorm.DB) *_AllUnitLoadHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllUnitLoadHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllUnitLoadHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_unit_load_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllUnitLoadHistoryMgr) GetTableName() string {
	return "__all_unit_load_history"
}

// Reset 重置gorm会话
func (obj *_AllUnitLoadHistoryMgr) Reset() *_AllUnitLoadHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllUnitLoadHistoryMgr) Get() (result AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllUnitLoadHistoryMgr) Gets() (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllUnitLoadHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllUnitLoadHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithTenantID tenant_id获取
func (obj *_AllUnitLoadHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllUnitLoadHistoryMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrIP svr_ip获取
func (obj *_AllUnitLoadHistoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllUnitLoadHistoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithUnitID unit_id获取
func (obj *_AllUnitLoadHistoryMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithLoad load获取
func (obj *_AllUnitLoadHistoryMgr) WithLoad(load float64) Option {
	return optionFunc(func(o *options) { o.query["load"] = load })
}

// WithDiskUsageRate disk_usage_rate获取
func (obj *_AllUnitLoadHistoryMgr) WithDiskUsageRate(diskUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["disk_usage_rate"] = diskUsageRate })
}

// WithMemoryUsageRate memory_usage_rate获取
func (obj *_AllUnitLoadHistoryMgr) WithMemoryUsageRate(memoryUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["memory_usage_rate"] = memoryUsageRate })
}

// WithCPUUsageRate cpu_usage_rate获取
func (obj *_AllUnitLoadHistoryMgr) WithCPUUsageRate(cpuUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_usage_rate"] = cpuUsageRate })
}

// WithIopsUsageRate iops_usage_rate获取
func (obj *_AllUnitLoadHistoryMgr) WithIopsUsageRate(iopsUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["iops_usage_rate"] = iopsUsageRate })
}

// WithDiskWeight disk_weight获取
func (obj *_AllUnitLoadHistoryMgr) WithDiskWeight(diskWeight float64) Option {
	return optionFunc(func(o *options) { o.query["disk_weight"] = diskWeight })
}

// WithMemoryWeight memory_weight获取
func (obj *_AllUnitLoadHistoryMgr) WithMemoryWeight(memoryWeight float64) Option {
	return optionFunc(func(o *options) { o.query["memory_weight"] = memoryWeight })
}

// WithCPUWeight cpu_weight获取
func (obj *_AllUnitLoadHistoryMgr) WithCPUWeight(cpuWeight float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_weight"] = cpuWeight })
}

// WithIopsWeight iops_weight获取
func (obj *_AllUnitLoadHistoryMgr) WithIopsWeight(iopsWeight float64) Option {
	return optionFunc(func(o *options) { o.query["iops_weight"] = iopsWeight })
}

// WithRsSvrIP rs_svr_ip获取
func (obj *_AllUnitLoadHistoryMgr) WithRsSvrIP(rsSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_ip"] = rsSvrIP })
}

// WithRsSvrPort rs_svr_port获取
func (obj *_AllUnitLoadHistoryMgr) WithRsSvrPort(rsSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_port"] = rsSvrPort })
}

// GetByOption 功能选项模式获取
func (obj *_AllUnitLoadHistoryMgr) GetByOption(opts ...Option) (result AllUnitLoadHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllUnitLoadHistoryMgr) GetByOptions(opts ...Option) (results []*AllUnitLoadHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllUnitLoadHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllUnitLoadHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (result AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`gmt_create` = ?", gmtCreate).First(&result).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromZone(zone string) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromZone(zones []string) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromSvrIP(svrIP string) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromSvrPort(svrPort int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromUnitID(unitID int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromLoad 通过load获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromLoad(load float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`load` = ?", load).Find(&results).Error

	return
}

// GetBatchFromLoad 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromLoad(loads []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`load` IN (?)", loads).Find(&results).Error

	return
}

// GetFromDiskUsageRate 通过disk_usage_rate获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromDiskUsageRate(diskUsageRate float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`disk_usage_rate` = ?", diskUsageRate).Find(&results).Error

	return
}

// GetBatchFromDiskUsageRate 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromDiskUsageRate(diskUsageRates []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`disk_usage_rate` IN (?)", diskUsageRates).Find(&results).Error

	return
}

// GetFromMemoryUsageRate 通过memory_usage_rate获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromMemoryUsageRate(memoryUsageRate float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`memory_usage_rate` = ?", memoryUsageRate).Find(&results).Error

	return
}

// GetBatchFromMemoryUsageRate 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromMemoryUsageRate(memoryUsageRates []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`memory_usage_rate` IN (?)", memoryUsageRates).Find(&results).Error

	return
}

// GetFromCPUUsageRate 通过cpu_usage_rate获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromCPUUsageRate(cpuUsageRate float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`cpu_usage_rate` = ?", cpuUsageRate).Find(&results).Error

	return
}

// GetBatchFromCPUUsageRate 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromCPUUsageRate(cpuUsageRates []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`cpu_usage_rate` IN (?)", cpuUsageRates).Find(&results).Error

	return
}

// GetFromIopsUsageRate 通过iops_usage_rate获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromIopsUsageRate(iopsUsageRate float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`iops_usage_rate` = ?", iopsUsageRate).Find(&results).Error

	return
}

// GetBatchFromIopsUsageRate 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromIopsUsageRate(iopsUsageRates []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`iops_usage_rate` IN (?)", iopsUsageRates).Find(&results).Error

	return
}

// GetFromDiskWeight 通过disk_weight获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromDiskWeight(diskWeight float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`disk_weight` = ?", diskWeight).Find(&results).Error

	return
}

// GetBatchFromDiskWeight 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromDiskWeight(diskWeights []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`disk_weight` IN (?)", diskWeights).Find(&results).Error

	return
}

// GetFromMemoryWeight 通过memory_weight获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromMemoryWeight(memoryWeight float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`memory_weight` = ?", memoryWeight).Find(&results).Error

	return
}

// GetBatchFromMemoryWeight 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromMemoryWeight(memoryWeights []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`memory_weight` IN (?)", memoryWeights).Find(&results).Error

	return
}

// GetFromCPUWeight 通过cpu_weight获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromCPUWeight(cpuWeight float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`cpu_weight` = ?", cpuWeight).Find(&results).Error

	return
}

// GetBatchFromCPUWeight 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromCPUWeight(cpuWeights []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`cpu_weight` IN (?)", cpuWeights).Find(&results).Error

	return
}

// GetFromIopsWeight 通过iops_weight获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromIopsWeight(iopsWeight float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`iops_weight` = ?", iopsWeight).Find(&results).Error

	return
}

// GetBatchFromIopsWeight 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromIopsWeight(iopsWeights []float64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`iops_weight` IN (?)", iopsWeights).Find(&results).Error

	return
}

// GetFromRsSvrIP 通过rs_svr_ip获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromRsSvrIP(rsSvrIP string) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`rs_svr_ip` = ?", rsSvrIP).Find(&results).Error

	return
}

// GetBatchFromRsSvrIP 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromRsSvrIP(rsSvrIPs []string) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`rs_svr_ip` IN (?)", rsSvrIPs).Find(&results).Error

	return
}

// GetFromRsSvrPort 通过rs_svr_port获取内容
func (obj *_AllUnitLoadHistoryMgr) GetFromRsSvrPort(rsSvrPort int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`rs_svr_port` = ?", rsSvrPort).Find(&results).Error

	return
}

// GetBatchFromRsSvrPort 批量查找
func (obj *_AllUnitLoadHistoryMgr) GetBatchFromRsSvrPort(rsSvrPorts []int64) (results []*AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`rs_svr_port` IN (?)", rsSvrPorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllUnitLoadHistoryMgr) FetchByPrimaryKey(gmtCreate time.Time) (result AllUnitLoadHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitLoadHistory{}).Where("`gmt_create` = ?", gmtCreate).First(&result).Error

	return
}
