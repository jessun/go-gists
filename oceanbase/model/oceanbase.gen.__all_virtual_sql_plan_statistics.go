package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualSQLPlanStatisticsMgr struct {
	*_BaseMgr
}

// AllVirtualSQLPlanStatisticsMgr open func
func AllVirtualSQLPlanStatisticsMgr(db *gorm.DB) *_AllVirtualSQLPlanStatisticsMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLPlanStatisticsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLPlanStatisticsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_plan_statistics"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetTableName() string {
	return "__all_virtual_sql_plan_statistics"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLPlanStatisticsMgr) Reset() *_AllVirtualSQLPlanStatisticsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) Get() (result AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLPlanStatisticsMgr) Gets() (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLPlanStatisticsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPlanID plan_id获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithOperationID operation_id获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithOperationID(operationID int64) Option {
	return optionFunc(func(o *options) { o.query["operation_id"] = operationID })
}

// WithExecutions executions获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithExecutions(executions int64) Option {
	return optionFunc(func(o *options) { o.query["executions"] = executions })
}

// WithOutputRows output_rows获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithOutputRows(outputRows int64) Option {
	return optionFunc(func(o *options) { o.query["output_rows"] = outputRows })
}

// WithInputRows input_rows获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithInputRows(inputRows int64) Option {
	return optionFunc(func(o *options) { o.query["input_rows"] = inputRows })
}

// WithRescanTimes rescan_times获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithRescanTimes(rescanTimes int64) Option {
	return optionFunc(func(o *options) { o.query["rescan_times"] = rescanTimes })
}

// WithBufferGets buffer_gets获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithBufferGets(bufferGets int64) Option {
	return optionFunc(func(o *options) { o.query["buffer_gets"] = bufferGets })
}

// WithDiskReads disk_reads获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithDiskReads(diskReads int64) Option {
	return optionFunc(func(o *options) { o.query["disk_reads"] = diskReads })
}

// WithDiskWrites disk_writes获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithDiskWrites(diskWrites int64) Option {
	return optionFunc(func(o *options) { o.query["disk_writes"] = diskWrites })
}

// WithElapsedTime elapsed_time获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithElapsedTime(elapsedTime int64) Option {
	return optionFunc(func(o *options) { o.query["elapsed_time"] = elapsedTime })
}

// WithExtendInfo1 extend_info1获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithExtendInfo1(extendInfo1 string) Option {
	return optionFunc(func(o *options) { o.query["extend_info1"] = extendInfo1 })
}

// WithExtendInfo2 extend_info2获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) WithExtendInfo2(extendInfo2 string) Option {
	return optionFunc(func(o *options) { o.query["extend_info2"] = extendInfo2 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetByOption(opts ...Option) (result AllVirtualSQLPlanStatistics, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLPlanStatistics, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLPlanStatisticsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLPlanStatistics, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where(options.query)
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
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPlanID 通过plan_id获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromPlanID(planID int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`plan_id` = ?", planID).Find(&results).Error

	return
}

// GetBatchFromPlanID 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromPlanID(planIDs []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error

	return
}

// GetFromOperationID 通过operation_id获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromOperationID(operationID int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`operation_id` = ?", operationID).Find(&results).Error

	return
}

// GetBatchFromOperationID 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromOperationID(operationIDs []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`operation_id` IN (?)", operationIDs).Find(&results).Error

	return
}

// GetFromExecutions 通过executions获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromExecutions(executions int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`executions` = ?", executions).Find(&results).Error

	return
}

// GetBatchFromExecutions 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromExecutions(executionss []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`executions` IN (?)", executionss).Find(&results).Error

	return
}

// GetFromOutputRows 通过output_rows获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromOutputRows(outputRows int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`output_rows` = ?", outputRows).Find(&results).Error

	return
}

// GetBatchFromOutputRows 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromOutputRows(outputRowss []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`output_rows` IN (?)", outputRowss).Find(&results).Error

	return
}

// GetFromInputRows 通过input_rows获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromInputRows(inputRows int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`input_rows` = ?", inputRows).Find(&results).Error

	return
}

// GetBatchFromInputRows 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromInputRows(inputRowss []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`input_rows` IN (?)", inputRowss).Find(&results).Error

	return
}

// GetFromRescanTimes 通过rescan_times获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromRescanTimes(rescanTimes int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`rescan_times` = ?", rescanTimes).Find(&results).Error

	return
}

// GetBatchFromRescanTimes 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromRescanTimes(rescanTimess []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`rescan_times` IN (?)", rescanTimess).Find(&results).Error

	return
}

// GetFromBufferGets 通过buffer_gets获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromBufferGets(bufferGets int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`buffer_gets` = ?", bufferGets).Find(&results).Error

	return
}

// GetBatchFromBufferGets 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromBufferGets(bufferGetss []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`buffer_gets` IN (?)", bufferGetss).Find(&results).Error

	return
}

// GetFromDiskReads 通过disk_reads获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromDiskReads(diskReads int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`disk_reads` = ?", diskReads).Find(&results).Error

	return
}

// GetBatchFromDiskReads 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromDiskReads(diskReadss []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`disk_reads` IN (?)", diskReadss).Find(&results).Error

	return
}

// GetFromDiskWrites 通过disk_writes获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromDiskWrites(diskWrites int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`disk_writes` = ?", diskWrites).Find(&results).Error

	return
}

// GetBatchFromDiskWrites 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromDiskWrites(diskWritess []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`disk_writes` IN (?)", diskWritess).Find(&results).Error

	return
}

// GetFromElapsedTime 通过elapsed_time获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromElapsedTime(elapsedTime int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`elapsed_time` = ?", elapsedTime).Find(&results).Error

	return
}

// GetBatchFromElapsedTime 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromElapsedTime(elapsedTimes []int64) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`elapsed_time` IN (?)", elapsedTimes).Find(&results).Error

	return
}

// GetFromExtendInfo1 通过extend_info1获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromExtendInfo1(extendInfo1 string) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`extend_info1` = ?", extendInfo1).Find(&results).Error

	return
}

// GetBatchFromExtendInfo1 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromExtendInfo1(extendInfo1s []string) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`extend_info1` IN (?)", extendInfo1s).Find(&results).Error

	return
}

// GetFromExtendInfo2 通过extend_info2获取内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetFromExtendInfo2(extendInfo2 string) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`extend_info2` = ?", extendInfo2).Find(&results).Error

	return
}

// GetBatchFromExtendInfo2 批量查找
func (obj *_AllVirtualSQLPlanStatisticsMgr) GetBatchFromExtendInfo2(extendInfo2s []string) (results []*AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`extend_info2` IN (?)", extendInfo2s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSQLPlanStatisticsMgr) FetchByPrimaryKey(tenantID int64, svrIP string, svrPort int64, planID int64, operationID int64) (result AllVirtualSQLPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanStatistics{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `plan_id` = ? AND `operation_id` = ?", tenantID, svrIP, svrPort, planID, operationID).First(&result).Error

	return
}
