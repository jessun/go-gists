package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualSQLWorkareaMemoryInfoMgr struct {
	*_BaseMgr
}

// AllVirtualSQLWorkareaMemoryInfoMgr open func
func AllVirtualSQLWorkareaMemoryInfoMgr(db *gorm.DB) *_AllVirtualSQLWorkareaMemoryInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLWorkareaMemoryInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLWorkareaMemoryInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_workarea_memory_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetTableName() string {
	return "__all_virtual_sql_workarea_memory_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) Reset() *_AllVirtualSQLWorkareaMemoryInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) Get() (result AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) Gets() (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithMaxWorkareaSize max_workarea_size获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithMaxWorkareaSize(maxWorkareaSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_workarea_size"] = maxWorkareaSize })
}

// WithWorkareaHoldSize workarea_hold_size获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithWorkareaHoldSize(workareaHoldSize int64) Option {
	return optionFunc(func(o *options) { o.query["workarea_hold_size"] = workareaHoldSize })
}

// WithMaxAutoWorkareaSize max_auto_workarea_size获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithMaxAutoWorkareaSize(maxAutoWorkareaSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_auto_workarea_size"] = maxAutoWorkareaSize })
}

// WithMemTarget mem_target获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithMemTarget(memTarget int64) Option {
	return optionFunc(func(o *options) { o.query["mem_target"] = memTarget })
}

// WithTotalMemUsed total_mem_used获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithTotalMemUsed(totalMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["total_mem_used"] = totalMemUsed })
}

// WithGlobalMemBound global_mem_bound获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithGlobalMemBound(globalMemBound int64) Option {
	return optionFunc(func(o *options) { o.query["global_mem_bound"] = globalMemBound })
}

// WithDriftSize drift_size获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithDriftSize(driftSize int64) Option {
	return optionFunc(func(o *options) { o.query["drift_size"] = driftSize })
}

// WithWorkareaCount workarea_count获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithWorkareaCount(workareaCount int64) Option {
	return optionFunc(func(o *options) { o.query["workarea_count"] = workareaCount })
}

// WithManualCalcCount manual_calc_count获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithManualCalcCount(manualCalcCount int64) Option {
	return optionFunc(func(o *options) { o.query["manual_calc_count"] = manualCalcCount })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetByOption(opts ...Option) (result AllVirtualSQLWorkareaMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLWorkareaMemoryInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromMaxWorkareaSize 通过max_workarea_size获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromMaxWorkareaSize(maxWorkareaSize int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`max_workarea_size` = ?", maxWorkareaSize).Find(&results).Error

	return
}

// GetBatchFromMaxWorkareaSize 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromMaxWorkareaSize(maxWorkareaSizes []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`max_workarea_size` IN (?)", maxWorkareaSizes).Find(&results).Error

	return
}

// GetFromWorkareaHoldSize 通过workarea_hold_size获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromWorkareaHoldSize(workareaHoldSize int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`workarea_hold_size` = ?", workareaHoldSize).Find(&results).Error

	return
}

// GetBatchFromWorkareaHoldSize 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromWorkareaHoldSize(workareaHoldSizes []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`workarea_hold_size` IN (?)", workareaHoldSizes).Find(&results).Error

	return
}

// GetFromMaxAutoWorkareaSize 通过max_auto_workarea_size获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromMaxAutoWorkareaSize(maxAutoWorkareaSize int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`max_auto_workarea_size` = ?", maxAutoWorkareaSize).Find(&results).Error

	return
}

// GetBatchFromMaxAutoWorkareaSize 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromMaxAutoWorkareaSize(maxAutoWorkareaSizes []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`max_auto_workarea_size` IN (?)", maxAutoWorkareaSizes).Find(&results).Error

	return
}

// GetFromMemTarget 通过mem_target获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromMemTarget(memTarget int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`mem_target` = ?", memTarget).Find(&results).Error

	return
}

// GetBatchFromMemTarget 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromMemTarget(memTargets []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`mem_target` IN (?)", memTargets).Find(&results).Error

	return
}

// GetFromTotalMemUsed 通过total_mem_used获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromTotalMemUsed(totalMemUsed int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`total_mem_used` = ?", totalMemUsed).Find(&results).Error

	return
}

// GetBatchFromTotalMemUsed 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromTotalMemUsed(totalMemUseds []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`total_mem_used` IN (?)", totalMemUseds).Find(&results).Error

	return
}

// GetFromGlobalMemBound 通过global_mem_bound获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromGlobalMemBound(globalMemBound int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`global_mem_bound` = ?", globalMemBound).Find(&results).Error

	return
}

// GetBatchFromGlobalMemBound 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromGlobalMemBound(globalMemBounds []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`global_mem_bound` IN (?)", globalMemBounds).Find(&results).Error

	return
}

// GetFromDriftSize 通过drift_size获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromDriftSize(driftSize int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`drift_size` = ?", driftSize).Find(&results).Error

	return
}

// GetBatchFromDriftSize 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromDriftSize(driftSizes []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`drift_size` IN (?)", driftSizes).Find(&results).Error

	return
}

// GetFromWorkareaCount 通过workarea_count获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromWorkareaCount(workareaCount int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`workarea_count` = ?", workareaCount).Find(&results).Error

	return
}

// GetBatchFromWorkareaCount 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromWorkareaCount(workareaCounts []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`workarea_count` IN (?)", workareaCounts).Find(&results).Error

	return
}

// GetFromManualCalcCount 通过manual_calc_count获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromManualCalcCount(manualCalcCount int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`manual_calc_count` = ?", manualCalcCount).Find(&results).Error

	return
}

// GetBatchFromManualCalcCount 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromManualCalcCount(manualCalcCounts []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`manual_calc_count` IN (?)", manualCalcCounts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSQLWorkareaMemoryInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaMemoryInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
