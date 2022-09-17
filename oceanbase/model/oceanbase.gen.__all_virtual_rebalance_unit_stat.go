package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualRebalanceUnitStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceUnitStatMgr open func
func AllVirtualRebalanceUnitStatMgr(db *gorm.DB) *_AllVirtualRebalanceUnitStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceUnitStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceUnitStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_unit_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceUnitStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_unit_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceUnitStatMgr) Reset() *_AllVirtualRebalanceUnitStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceUnitStatMgr) Get() (result AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceUnitStatMgr) Gets() (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceUnitStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithLoad load获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithLoad(load float64) Option {
	return optionFunc(func(o *options) { o.query["load"] = load })
}

// WithCPUUsageRate cpu_usage_rate获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithCPUUsageRate(cpuUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_usage_rate"] = cpuUsageRate })
}

// WithDiskUsageRate disk_usage_rate获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithDiskUsageRate(diskUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["disk_usage_rate"] = diskUsageRate })
}

// WithIopsUsageRate iops_usage_rate获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithIopsUsageRate(iopsUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["iops_usage_rate"] = iopsUsageRate })
}

// WithMemoryUsageRate memory_usage_rate获取
func (obj *_AllVirtualRebalanceUnitStatMgr) WithMemoryUsageRate(memoryUsageRate float64) Option {
	return optionFunc(func(o *options) { o.query["memory_usage_rate"] = memoryUsageRate })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceUnitStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceUnitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceUnitStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceUnitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceUnitStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceUnitStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where(options.query)
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
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromZone(zone string) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromUnitID(unitID int64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromLoad 通过load获取内容
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromLoad(load float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`load` = ?", load).Find(&results).Error

	return
}

// GetBatchFromLoad 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromLoad(loads []float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`load` IN (?)", loads).Find(&results).Error

	return
}

// GetFromCPUUsageRate 通过cpu_usage_rate获取内容
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromCPUUsageRate(cpuUsageRate float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`cpu_usage_rate` = ?", cpuUsageRate).Find(&results).Error

	return
}

// GetBatchFromCPUUsageRate 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromCPUUsageRate(cpuUsageRates []float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`cpu_usage_rate` IN (?)", cpuUsageRates).Find(&results).Error

	return
}

// GetFromDiskUsageRate 通过disk_usage_rate获取内容
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromDiskUsageRate(diskUsageRate float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`disk_usage_rate` = ?", diskUsageRate).Find(&results).Error

	return
}

// GetBatchFromDiskUsageRate 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromDiskUsageRate(diskUsageRates []float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`disk_usage_rate` IN (?)", diskUsageRates).Find(&results).Error

	return
}

// GetFromIopsUsageRate 通过iops_usage_rate获取内容
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromIopsUsageRate(iopsUsageRate float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`iops_usage_rate` = ?", iopsUsageRate).Find(&results).Error

	return
}

// GetBatchFromIopsUsageRate 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromIopsUsageRate(iopsUsageRates []float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`iops_usage_rate` IN (?)", iopsUsageRates).Find(&results).Error

	return
}

// GetFromMemoryUsageRate 通过memory_usage_rate获取内容
func (obj *_AllVirtualRebalanceUnitStatMgr) GetFromMemoryUsageRate(memoryUsageRate float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`memory_usage_rate` = ?", memoryUsageRate).Find(&results).Error

	return
}

// GetBatchFromMemoryUsageRate 批量查找
func (obj *_AllVirtualRebalanceUnitStatMgr) GetBatchFromMemoryUsageRate(memoryUsageRates []float64) (results []*AllVirtualRebalanceUnitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitStat{}).Where("`memory_usage_rate` IN (?)", memoryUsageRates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
