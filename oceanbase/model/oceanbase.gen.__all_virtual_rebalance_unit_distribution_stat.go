package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualRebalanceUnitDistributionStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceUnitDistributionStatMgr open func
func AllVirtualRebalanceUnitDistributionStatMgr(db *gorm.DB) *_AllVirtualRebalanceUnitDistributionStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceUnitDistributionStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceUnitDistributionStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_unit_distribution_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_unit_distribution_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) Reset() *_AllVirtualRebalanceUnitDistributionStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) Get() (result AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) Gets() (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUnitID unit_id获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithZone zone获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithMaxCPU max_cpu获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMaxCPU(maxCPU float64) Option {
	return optionFunc(func(o *options) { o.query["max_cpu"] = maxCPU })
}

// WithMinCPU min_cpu获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMinCPU(minCPU float64) Option {
	return optionFunc(func(o *options) { o.query["min_cpu"] = minCPU })
}

// WithMaxMemory max_memory获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMaxMemory(maxMemory int64) Option {
	return optionFunc(func(o *options) { o.query["max_memory"] = maxMemory })
}

// WithMinMemory min_memory获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMinMemory(minMemory int64) Option {
	return optionFunc(func(o *options) { o.query["min_memory"] = minMemory })
}

// WithMaxIops max_iops获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMaxIops(maxIops int64) Option {
	return optionFunc(func(o *options) { o.query["max_iops"] = maxIops })
}

// WithMinIops min_iops获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMinIops(minIops int64) Option {
	return optionFunc(func(o *options) { o.query["min_iops"] = minIops })
}

// WithMaxDiskSize max_disk_size获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMaxDiskSize(maxDiskSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_disk_size"] = maxDiskSize })
}

// WithMaxSessionNum max_session_num获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) WithMaxSessionNum(maxSessionNum int64) Option {
	return optionFunc(func(o *options) { o.query["max_session_num"] = maxSessionNum })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceUnitDistributionStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceUnitDistributionStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where(options.query)
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

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromUnitID(unitID int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromZone(zone string) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromMaxCPU 通过max_cpu获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMaxCPU(maxCPU float64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_cpu` = ?", maxCPU).Find(&results).Error

	return
}

// GetBatchFromMaxCPU 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMaxCPU(maxCPUs []float64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_cpu` IN (?)", maxCPUs).Find(&results).Error

	return
}

// GetFromMinCPU 通过min_cpu获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMinCPU(minCPU float64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`min_cpu` = ?", minCPU).Find(&results).Error

	return
}

// GetBatchFromMinCPU 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMinCPU(minCPUs []float64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`min_cpu` IN (?)", minCPUs).Find(&results).Error

	return
}

// GetFromMaxMemory 通过max_memory获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMaxMemory(maxMemory int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_memory` = ?", maxMemory).Find(&results).Error

	return
}

// GetBatchFromMaxMemory 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMaxMemory(maxMemorys []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_memory` IN (?)", maxMemorys).Find(&results).Error

	return
}

// GetFromMinMemory 通过min_memory获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMinMemory(minMemory int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`min_memory` = ?", minMemory).Find(&results).Error

	return
}

// GetBatchFromMinMemory 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMinMemory(minMemorys []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`min_memory` IN (?)", minMemorys).Find(&results).Error

	return
}

// GetFromMaxIops 通过max_iops获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMaxIops(maxIops int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_iops` = ?", maxIops).Find(&results).Error

	return
}

// GetBatchFromMaxIops 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMaxIops(maxIopss []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_iops` IN (?)", maxIopss).Find(&results).Error

	return
}

// GetFromMinIops 通过min_iops获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMinIops(minIops int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`min_iops` = ?", minIops).Find(&results).Error

	return
}

// GetBatchFromMinIops 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMinIops(minIopss []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`min_iops` IN (?)", minIopss).Find(&results).Error

	return
}

// GetFromMaxDiskSize 通过max_disk_size获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMaxDiskSize(maxDiskSize int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_disk_size` = ?", maxDiskSize).Find(&results).Error

	return
}

// GetBatchFromMaxDiskSize 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMaxDiskSize(maxDiskSizes []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_disk_size` IN (?)", maxDiskSizes).Find(&results).Error

	return
}

// GetFromMaxSessionNum 通过max_session_num获取内容
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetFromMaxSessionNum(maxSessionNum int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_session_num` = ?", maxSessionNum).Find(&results).Error

	return
}

// GetBatchFromMaxSessionNum 批量查找
func (obj *_AllVirtualRebalanceUnitDistributionStatMgr) GetBatchFromMaxSessionNum(maxSessionNums []int64) (results []*AllVirtualRebalanceUnitDistributionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceUnitDistributionStat{}).Where("`max_session_num` IN (?)", maxSessionNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
