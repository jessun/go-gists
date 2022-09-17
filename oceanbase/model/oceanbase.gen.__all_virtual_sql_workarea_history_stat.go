package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualSQLWorkareaHistoryStatMgr struct {
	*_BaseMgr
}

// AllVirtualSQLWorkareaHistoryStatMgr open func
func AllVirtualSQLWorkareaHistoryStatMgr(db *gorm.DB) *_AllVirtualSQLWorkareaHistoryStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLWorkareaHistoryStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLWorkareaHistoryStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_workarea_history_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetTableName() string {
	return "__all_virtual_sql_workarea_history_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) Reset() *_AllVirtualSQLWorkareaHistoryStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) Get() (result AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) Gets() (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPlanID plan_id获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithSQLID sql_id获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithOperationType operation_type获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithOperationType(operationType string) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithOperationID operation_id获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithOperationID(operationID int64) Option {
	return optionFunc(func(o *options) { o.query["operation_id"] = operationID })
}

// WithEstimatedOptimalSize estimated_optimal_size获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithEstimatedOptimalSize(estimatedOptimalSize int64) Option {
	return optionFunc(func(o *options) { o.query["estimated_optimal_size"] = estimatedOptimalSize })
}

// WithEstimatedOnepassSize estimated_onepass_size获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithEstimatedOnepassSize(estimatedOnepassSize int64) Option {
	return optionFunc(func(o *options) { o.query["estimated_onepass_size"] = estimatedOnepassSize })
}

// WithLastMemoryUsed last_memory_used获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithLastMemoryUsed(lastMemoryUsed int64) Option {
	return optionFunc(func(o *options) { o.query["last_memory_used"] = lastMemoryUsed })
}

// WithLastExecution last_execution获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithLastExecution(lastExecution string) Option {
	return optionFunc(func(o *options) { o.query["last_execution"] = lastExecution })
}

// WithLastDegree last_degree获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithLastDegree(lastDegree int64) Option {
	return optionFunc(func(o *options) { o.query["last_degree"] = lastDegree })
}

// WithTotalExecutions total_executions获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithTotalExecutions(totalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["total_executions"] = totalExecutions })
}

// WithOptimalExecutions optimal_executions获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithOptimalExecutions(optimalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["optimal_executions"] = optimalExecutions })
}

// WithOnepassExecutions onepass_executions获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithOnepassExecutions(onepassExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["onepass_executions"] = onepassExecutions })
}

// WithMultipassesExecutions multipasses_executions获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithMultipassesExecutions(multipassesExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["multipasses_executions"] = multipassesExecutions })
}

// WithActiveTime active_time获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithActiveTime(activeTime int64) Option {
	return optionFunc(func(o *options) { o.query["active_time"] = activeTime })
}

// WithMaxTempsegSize max_tempseg_size获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithMaxTempsegSize(maxTempsegSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_tempseg_size"] = maxTempsegSize })
}

// WithLastTempsegSize last_tempseg_size获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithLastTempsegSize(lastTempsegSize int64) Option {
	return optionFunc(func(o *options) { o.query["last_tempseg_size"] = lastTempsegSize })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPolicy policy获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) WithPolicy(policy string) Option {
	return optionFunc(func(o *options) { o.query["policy"] = policy })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetByOption(opts ...Option) (result AllVirtualSQLWorkareaHistoryStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLWorkareaHistoryStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where(options.query)
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
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPlanID 通过plan_id获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromPlanID(planID int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`plan_id` = ?", planID).Find(&results).Error

	return
}

// GetBatchFromPlanID 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromPlanID(planIDs []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromSQLID(sqlID string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromSQLID(sqlIDs []string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromOperationType(operationType string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromOperationType(operationTypes []string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromOperationID 通过operation_id获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromOperationID(operationID int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`operation_id` = ?", operationID).Find(&results).Error

	return
}

// GetBatchFromOperationID 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromOperationID(operationIDs []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`operation_id` IN (?)", operationIDs).Find(&results).Error

	return
}

// GetFromEstimatedOptimalSize 通过estimated_optimal_size获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromEstimatedOptimalSize(estimatedOptimalSize int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`estimated_optimal_size` = ?", estimatedOptimalSize).Find(&results).Error

	return
}

// GetBatchFromEstimatedOptimalSize 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromEstimatedOptimalSize(estimatedOptimalSizes []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`estimated_optimal_size` IN (?)", estimatedOptimalSizes).Find(&results).Error

	return
}

// GetFromEstimatedOnepassSize 通过estimated_onepass_size获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromEstimatedOnepassSize(estimatedOnepassSize int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`estimated_onepass_size` = ?", estimatedOnepassSize).Find(&results).Error

	return
}

// GetBatchFromEstimatedOnepassSize 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromEstimatedOnepassSize(estimatedOnepassSizes []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`estimated_onepass_size` IN (?)", estimatedOnepassSizes).Find(&results).Error

	return
}

// GetFromLastMemoryUsed 通过last_memory_used获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromLastMemoryUsed(lastMemoryUsed int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_memory_used` = ?", lastMemoryUsed).Find(&results).Error

	return
}

// GetBatchFromLastMemoryUsed 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromLastMemoryUsed(lastMemoryUseds []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_memory_used` IN (?)", lastMemoryUseds).Find(&results).Error

	return
}

// GetFromLastExecution 通过last_execution获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromLastExecution(lastExecution string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_execution` = ?", lastExecution).Find(&results).Error

	return
}

// GetBatchFromLastExecution 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromLastExecution(lastExecutions []string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_execution` IN (?)", lastExecutions).Find(&results).Error

	return
}

// GetFromLastDegree 通过last_degree获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromLastDegree(lastDegree int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_degree` = ?", lastDegree).Find(&results).Error

	return
}

// GetBatchFromLastDegree 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromLastDegree(lastDegrees []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_degree` IN (?)", lastDegrees).Find(&results).Error

	return
}

// GetFromTotalExecutions 通过total_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromTotalExecutions(totalExecutions int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`total_executions` = ?", totalExecutions).Find(&results).Error

	return
}

// GetBatchFromTotalExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromTotalExecutions(totalExecutionss []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`total_executions` IN (?)", totalExecutionss).Find(&results).Error

	return
}

// GetFromOptimalExecutions 通过optimal_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromOptimalExecutions(optimalExecutions int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`optimal_executions` = ?", optimalExecutions).Find(&results).Error

	return
}

// GetBatchFromOptimalExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromOptimalExecutions(optimalExecutionss []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`optimal_executions` IN (?)", optimalExecutionss).Find(&results).Error

	return
}

// GetFromOnepassExecutions 通过onepass_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromOnepassExecutions(onepassExecutions int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`onepass_executions` = ?", onepassExecutions).Find(&results).Error

	return
}

// GetBatchFromOnepassExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromOnepassExecutions(onepassExecutionss []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`onepass_executions` IN (?)", onepassExecutionss).Find(&results).Error

	return
}

// GetFromMultipassesExecutions 通过multipasses_executions获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromMultipassesExecutions(multipassesExecutions int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`multipasses_executions` = ?", multipassesExecutions).Find(&results).Error

	return
}

// GetBatchFromMultipassesExecutions 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromMultipassesExecutions(multipassesExecutionss []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`multipasses_executions` IN (?)", multipassesExecutionss).Find(&results).Error

	return
}

// GetFromActiveTime 通过active_time获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromActiveTime(activeTime int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`active_time` = ?", activeTime).Find(&results).Error

	return
}

// GetBatchFromActiveTime 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromActiveTime(activeTimes []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`active_time` IN (?)", activeTimes).Find(&results).Error

	return
}

// GetFromMaxTempsegSize 通过max_tempseg_size获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromMaxTempsegSize(maxTempsegSize int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`max_tempseg_size` = ?", maxTempsegSize).Find(&results).Error

	return
}

// GetBatchFromMaxTempsegSize 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromMaxTempsegSize(maxTempsegSizes []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`max_tempseg_size` IN (?)", maxTempsegSizes).Find(&results).Error

	return
}

// GetFromLastTempsegSize 通过last_tempseg_size获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromLastTempsegSize(lastTempsegSize int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_tempseg_size` = ?", lastTempsegSize).Find(&results).Error

	return
}

// GetBatchFromLastTempsegSize 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromLastTempsegSize(lastTempsegSizes []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`last_tempseg_size` IN (?)", lastTempsegSizes).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPolicy 通过policy获取内容
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetFromPolicy(policy string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`policy` = ?", policy).Find(&results).Error

	return
}

// GetBatchFromPolicy 批量查找
func (obj *_AllVirtualSQLWorkareaHistoryStatMgr) GetBatchFromPolicy(policys []string) (results []*AllVirtualSQLWorkareaHistoryStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaHistoryStat{}).Where("`policy` IN (?)", policys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
