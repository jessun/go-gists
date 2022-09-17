package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualRebalanceTenantStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceTenantStatMgr open func
func AllVirtualRebalanceTenantStatMgr(db *gorm.DB) *_AllVirtualRebalanceTenantStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceTenantStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceTenantStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_tenant_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceTenantStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_tenant_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceTenantStatMgr) Reset() *_AllVirtualRebalanceTenantStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceTenantStatMgr) Get() (result AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceTenantStatMgr) Gets() (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceTenantStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithCPUWeight cpu_weight获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithCPUWeight(cpuWeight float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_weight"] = cpuWeight })
}

// WithDiskWeight disk_weight获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithDiskWeight(diskWeight float64) Option {
	return optionFunc(func(o *options) { o.query["disk_weight"] = diskWeight })
}

// WithIopsWeight iops_weight获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithIopsWeight(iopsWeight float64) Option {
	return optionFunc(func(o *options) { o.query["iops_weight"] = iopsWeight })
}

// WithMemoryWeight memory_weight获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithMemoryWeight(memoryWeight float64) Option {
	return optionFunc(func(o *options) { o.query["memory_weight"] = memoryWeight })
}

// WithLoadImbalance load_imbalance获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithLoadImbalance(loadImbalance float64) Option {
	return optionFunc(func(o *options) { o.query["load_imbalance"] = loadImbalance })
}

// WithLoadAvg load_avg获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithLoadAvg(loadAvg float64) Option {
	return optionFunc(func(o *options) { o.query["load_avg"] = loadAvg })
}

// WithCPUImbalance cpu_imbalance获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithCPUImbalance(cpuImbalance float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_imbalance"] = cpuImbalance })
}

// WithCPUAvg cpu_avg获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithCPUAvg(cpuAvg float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_avg"] = cpuAvg })
}

// WithDiskImbalance disk_imbalance获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithDiskImbalance(diskImbalance float64) Option {
	return optionFunc(func(o *options) { o.query["disk_imbalance"] = diskImbalance })
}

// WithDiskAvg disk_avg获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithDiskAvg(diskAvg float64) Option {
	return optionFunc(func(o *options) { o.query["disk_avg"] = diskAvg })
}

// WithIopsImbalance iops_imbalance获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithIopsImbalance(iopsImbalance float64) Option {
	return optionFunc(func(o *options) { o.query["iops_imbalance"] = iopsImbalance })
}

// WithIopsAvg iops_avg获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithIopsAvg(iopsAvg float64) Option {
	return optionFunc(func(o *options) { o.query["iops_avg"] = iopsAvg })
}

// WithMemoryImbalance memory_imbalance获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithMemoryImbalance(memoryImbalance float64) Option {
	return optionFunc(func(o *options) { o.query["memory_imbalance"] = memoryImbalance })
}

// WithMemoryAvg memory_avg获取
func (obj *_AllVirtualRebalanceTenantStatMgr) WithMemoryAvg(memoryAvg float64) Option {
	return optionFunc(func(o *options) { o.query["memory_avg"] = memoryAvg })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceTenantStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceTenantStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceTenantStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceTenantStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceTenantStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceTenantStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where(options.query)
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
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromZone(zone string) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromCPUWeight 通过cpu_weight获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromCPUWeight(cpuWeight float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`cpu_weight` = ?", cpuWeight).Find(&results).Error

	return
}

// GetBatchFromCPUWeight 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromCPUWeight(cpuWeights []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`cpu_weight` IN (?)", cpuWeights).Find(&results).Error

	return
}

// GetFromDiskWeight 通过disk_weight获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromDiskWeight(diskWeight float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`disk_weight` = ?", diskWeight).Find(&results).Error

	return
}

// GetBatchFromDiskWeight 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromDiskWeight(diskWeights []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`disk_weight` IN (?)", diskWeights).Find(&results).Error

	return
}

// GetFromIopsWeight 通过iops_weight获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromIopsWeight(iopsWeight float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`iops_weight` = ?", iopsWeight).Find(&results).Error

	return
}

// GetBatchFromIopsWeight 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromIopsWeight(iopsWeights []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`iops_weight` IN (?)", iopsWeights).Find(&results).Error

	return
}

// GetFromMemoryWeight 通过memory_weight获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromMemoryWeight(memoryWeight float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`memory_weight` = ?", memoryWeight).Find(&results).Error

	return
}

// GetBatchFromMemoryWeight 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromMemoryWeight(memoryWeights []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`memory_weight` IN (?)", memoryWeights).Find(&results).Error

	return
}

// GetFromLoadImbalance 通过load_imbalance获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromLoadImbalance(loadImbalance float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`load_imbalance` = ?", loadImbalance).Find(&results).Error

	return
}

// GetBatchFromLoadImbalance 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromLoadImbalance(loadImbalances []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`load_imbalance` IN (?)", loadImbalances).Find(&results).Error

	return
}

// GetFromLoadAvg 通过load_avg获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromLoadAvg(loadAvg float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`load_avg` = ?", loadAvg).Find(&results).Error

	return
}

// GetBatchFromLoadAvg 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromLoadAvg(loadAvgs []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`load_avg` IN (?)", loadAvgs).Find(&results).Error

	return
}

// GetFromCPUImbalance 通过cpu_imbalance获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromCPUImbalance(cpuImbalance float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`cpu_imbalance` = ?", cpuImbalance).Find(&results).Error

	return
}

// GetBatchFromCPUImbalance 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromCPUImbalance(cpuImbalances []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`cpu_imbalance` IN (?)", cpuImbalances).Find(&results).Error

	return
}

// GetFromCPUAvg 通过cpu_avg获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromCPUAvg(cpuAvg float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`cpu_avg` = ?", cpuAvg).Find(&results).Error

	return
}

// GetBatchFromCPUAvg 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromCPUAvg(cpuAvgs []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`cpu_avg` IN (?)", cpuAvgs).Find(&results).Error

	return
}

// GetFromDiskImbalance 通过disk_imbalance获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromDiskImbalance(diskImbalance float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`disk_imbalance` = ?", diskImbalance).Find(&results).Error

	return
}

// GetBatchFromDiskImbalance 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromDiskImbalance(diskImbalances []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`disk_imbalance` IN (?)", diskImbalances).Find(&results).Error

	return
}

// GetFromDiskAvg 通过disk_avg获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromDiskAvg(diskAvg float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`disk_avg` = ?", diskAvg).Find(&results).Error

	return
}

// GetBatchFromDiskAvg 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromDiskAvg(diskAvgs []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`disk_avg` IN (?)", diskAvgs).Find(&results).Error

	return
}

// GetFromIopsImbalance 通过iops_imbalance获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromIopsImbalance(iopsImbalance float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`iops_imbalance` = ?", iopsImbalance).Find(&results).Error

	return
}

// GetBatchFromIopsImbalance 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromIopsImbalance(iopsImbalances []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`iops_imbalance` IN (?)", iopsImbalances).Find(&results).Error

	return
}

// GetFromIopsAvg 通过iops_avg获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromIopsAvg(iopsAvg float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`iops_avg` = ?", iopsAvg).Find(&results).Error

	return
}

// GetBatchFromIopsAvg 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromIopsAvg(iopsAvgs []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`iops_avg` IN (?)", iopsAvgs).Find(&results).Error

	return
}

// GetFromMemoryImbalance 通过memory_imbalance获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromMemoryImbalance(memoryImbalance float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`memory_imbalance` = ?", memoryImbalance).Find(&results).Error

	return
}

// GetBatchFromMemoryImbalance 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromMemoryImbalance(memoryImbalances []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`memory_imbalance` IN (?)", memoryImbalances).Find(&results).Error

	return
}

// GetFromMemoryAvg 通过memory_avg获取内容
func (obj *_AllVirtualRebalanceTenantStatMgr) GetFromMemoryAvg(memoryAvg float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`memory_avg` = ?", memoryAvg).Find(&results).Error

	return
}

// GetBatchFromMemoryAvg 批量查找
func (obj *_AllVirtualRebalanceTenantStatMgr) GetBatchFromMemoryAvg(memoryAvgs []float64) (results []*AllVirtualRebalanceTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTenantStat{}).Where("`memory_avg` IN (?)", memoryAvgs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
