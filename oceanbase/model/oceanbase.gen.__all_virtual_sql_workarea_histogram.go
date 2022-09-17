package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualSQLWorkareaHistogramMgr struct {
	*_BaseMgr
}

// AllVirtualSQLWorkareaHistogramMgr open func
func AllVirtualSQLWorkareaHistogramMgr(db *gorm.DB) *_AllVirtualSQLWorkareaHistogramMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLWorkareaHistogramMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLWorkareaHistogramMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_workarea_histogram"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetTableName() string {
	return "__all_virtual_sql_workarea_histogram"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLWorkareaHistogramMgr) Reset() *_AllVirtualSQLWorkareaHistogramMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) Get() (result AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLWorkareaHistogramMgr) Gets() (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLWorkareaHistogramMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithLowOptimalSize low_optimal_size获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithLowOptimalSize(lowOptimalSize int64) Option {
	return optionFunc(func(o *options) { o.query["low_optimal_size"] = lowOptimalSize })
}

// WithHighOptimalSize high_optimal_size获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithHighOptimalSize(highOptimalSize int64) Option {
	return optionFunc(func(o *options) { o.query["high_optimal_size"] = highOptimalSize })
}

// WithOptimalExecutions optimal_executions获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithOptimalExecutions(optimalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["optimal_executions"] = optimalExecutions })
}

// WithOnepassExecutions onepass_executions获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithOnepassExecutions(onepassExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["onepass_executions"] = onepassExecutions })
}

// WithMultipassesExecutions multipasses_executions获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithMultipassesExecutions(multipassesExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["multipasses_executions"] = multipassesExecutions })
}

// WithTotalExecutions total_executions获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithTotalExecutions(totalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["total_executions"] = totalExecutions })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetByOption(opts ...Option) (result AllVirtualSQLWorkareaHistogram, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLWorkareaHistogramMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLWorkareaHistogram, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where(options.query)
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
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromLowOptimalSize 通过low_optimal_size获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromLowOptimalSize(lowOptimalSize int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`low_optimal_size` = ?", lowOptimalSize).Find(&results).Error

	return
}

// GetBatchFromLowOptimalSize 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromLowOptimalSize(lowOptimalSizes []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`low_optimal_size` IN (?)", lowOptimalSizes).Find(&results).Error

	return
}

// GetFromHighOptimalSize 通过high_optimal_size获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromHighOptimalSize(highOptimalSize int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`high_optimal_size` = ?", highOptimalSize).Find(&results).Error

	return
}

// GetBatchFromHighOptimalSize 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromHighOptimalSize(highOptimalSizes []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`high_optimal_size` IN (?)", highOptimalSizes).Find(&results).Error

	return
}

// GetFromOptimalExecutions 通过optimal_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromOptimalExecutions(optimalExecutions int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`optimal_executions` = ?", optimalExecutions).Find(&results).Error

	return
}

// GetBatchFromOptimalExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromOptimalExecutions(optimalExecutionss []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`optimal_executions` IN (?)", optimalExecutionss).Find(&results).Error

	return
}

// GetFromOnepassExecutions 通过onepass_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromOnepassExecutions(onepassExecutions int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`onepass_executions` = ?", onepassExecutions).Find(&results).Error

	return
}

// GetBatchFromOnepassExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromOnepassExecutions(onepassExecutionss []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`onepass_executions` IN (?)", onepassExecutionss).Find(&results).Error

	return
}

// GetFromMultipassesExecutions 通过multipasses_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromMultipassesExecutions(multipassesExecutions int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`multipasses_executions` = ?", multipassesExecutions).Find(&results).Error

	return
}

// GetBatchFromMultipassesExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromMultipassesExecutions(multipassesExecutionss []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`multipasses_executions` IN (?)", multipassesExecutionss).Find(&results).Error

	return
}

// GetFromTotalExecutions 通过total_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromTotalExecutions(totalExecutions int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`total_executions` = ?", totalExecutions).Find(&results).Error

	return
}

// GetBatchFromTotalExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromTotalExecutions(totalExecutionss []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`total_executions` IN (?)", totalExecutionss).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSQLWorkareaHistogramMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistogram{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
