package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSQLMonitorMgr struct {
	*_BaseMgr
}

// AllVirtualSQLMonitorMgr open func
func AllVirtualSQLMonitorMgr(db *gorm.DB) *_AllVirtualSQLMonitorMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLMonitorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLMonitorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_monitor"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLMonitorMgr) GetTableName() string {
	return "__all_virtual_sql_monitor"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLMonitorMgr) Reset() *_AllVirtualSQLMonitorMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLMonitorMgr) Get() (result AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLMonitorMgr) Gets() (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLMonitorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSQLMonitorMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSQLMonitorMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSQLMonitorMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithRequestID request_id获取
func (obj *_AllVirtualSQLMonitorMgr) WithRequestID(requestID int64) Option {
	return optionFunc(func(o *options) { o.query["request_id"] = requestID })
}

// WithJobID job_id获取
func (obj *_AllVirtualSQLMonitorMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTaskID task_id获取
func (obj *_AllVirtualSQLMonitorMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithPlanID plan_id获取
func (obj *_AllVirtualSQLMonitorMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithSchedulerIP scheduler_ip获取
func (obj *_AllVirtualSQLMonitorMgr) WithSchedulerIP(schedulerIP string) Option {
	return optionFunc(func(o *options) { o.query["scheduler_ip"] = schedulerIP })
}

// WithSchedulerPort scheduler_port获取
func (obj *_AllVirtualSQLMonitorMgr) WithSchedulerPort(schedulerPort int64) Option {
	return optionFunc(func(o *options) { o.query["scheduler_port"] = schedulerPort })
}

// WithMonitorInfo monitor_info获取
func (obj *_AllVirtualSQLMonitorMgr) WithMonitorInfo(monitorInfo string) Option {
	return optionFunc(func(o *options) { o.query["monitor_info"] = monitorInfo })
}

// WithExtendInfo extend_info获取
func (obj *_AllVirtualSQLMonitorMgr) WithExtendInfo(extendInfo string) Option {
	return optionFunc(func(o *options) { o.query["extend_info"] = extendInfo })
}

// WithSQLExecStart sql_exec_start获取
func (obj *_AllVirtualSQLMonitorMgr) WithSQLExecStart(sqlExecStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["sql_exec_start"] = sqlExecStart })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLMonitorMgr) GetByOption(opts ...Option) (result AllVirtualSQLMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLMonitorMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLMonitorMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLMonitor, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where(options.query)
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
func (obj *_AllVirtualSQLMonitorMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromRequestID 通过request_id获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromRequestID(requestID int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`request_id` = ?", requestID).Find(&results).Error

	return
}

// GetBatchFromRequestID 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromRequestID(requestIDs []int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`request_id` IN (?)", requestIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromJobID(jobID int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromTaskID(taskID int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromPlanID 通过plan_id获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromPlanID(planID int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`plan_id` = ?", planID).Find(&results).Error

	return
}

// GetBatchFromPlanID 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromPlanID(planIDs []int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error

	return
}

// GetFromSchedulerIP 通过scheduler_ip获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromSchedulerIP(schedulerIP string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`scheduler_ip` = ?", schedulerIP).Find(&results).Error

	return
}

// GetBatchFromSchedulerIP 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromSchedulerIP(schedulerIPs []string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`scheduler_ip` IN (?)", schedulerIPs).Find(&results).Error

	return
}

// GetFromSchedulerPort 通过scheduler_port获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromSchedulerPort(schedulerPort int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`scheduler_port` = ?", schedulerPort).Find(&results).Error

	return
}

// GetBatchFromSchedulerPort 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromSchedulerPort(schedulerPorts []int64) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`scheduler_port` IN (?)", schedulerPorts).Find(&results).Error

	return
}

// GetFromMonitorInfo 通过monitor_info获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromMonitorInfo(monitorInfo string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`monitor_info` = ?", monitorInfo).Find(&results).Error

	return
}

// GetBatchFromMonitorInfo 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromMonitorInfo(monitorInfos []string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`monitor_info` IN (?)", monitorInfos).Find(&results).Error

	return
}

// GetFromExtendInfo 通过extend_info获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromExtendInfo(extendInfo string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`extend_info` = ?", extendInfo).Find(&results).Error

	return
}

// GetBatchFromExtendInfo 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromExtendInfo(extendInfos []string) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`extend_info` IN (?)", extendInfos).Find(&results).Error

	return
}

// GetFromSQLExecStart 通过sql_exec_start获取内容
func (obj *_AllVirtualSQLMonitorMgr) GetFromSQLExecStart(sqlExecStart time.Time) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`sql_exec_start` = ?", sqlExecStart).Find(&results).Error

	return
}

// GetBatchFromSQLExecStart 批量查找
func (obj *_AllVirtualSQLMonitorMgr) GetBatchFromSQLExecStart(sqlExecStarts []time.Time) (results []*AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`sql_exec_start` IN (?)", sqlExecStarts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSQLMonitorMgr) FetchByPrimaryKey(tenantID int64, svrIP string, svrPort int64, requestID int64, jobID int64, taskID int64) (result AllVirtualSQLMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitor{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `request_id` = ? AND `job_id` = ? AND `task_id` = ?", tenantID, svrIP, svrPort, requestID, jobID, taskID).First(&result).Error

	return
}
