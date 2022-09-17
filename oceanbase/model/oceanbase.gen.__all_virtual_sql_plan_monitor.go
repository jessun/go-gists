package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSQLPlanMonitorMgr struct {
	*_BaseMgr
}

// AllVirtualSQLPlanMonitorMgr open func
func AllVirtualSQLPlanMonitorMgr(db *gorm.DB) *_AllVirtualSQLPlanMonitorMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLPlanMonitorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLPlanMonitorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_plan_monitor"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLPlanMonitorMgr) GetTableName() string {
	return "__all_virtual_sql_plan_monitor"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLPlanMonitorMgr) Reset() *_AllVirtualSQLPlanMonitorMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLPlanMonitorMgr) Get() (result AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLPlanMonitorMgr) Gets() (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLPlanMonitorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithRequestID REQUEST_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithRequestID(requestID int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST_ID"] = requestID })
}

// WithTraceID TRACE_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["TRACE_ID"] = traceID })
}

// WithFirstRefreshTime FIRST_REFRESH_TIME获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithFirstRefreshTime(firstRefreshTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["FIRST_REFRESH_TIME"] = firstRefreshTime })
}

// WithLastRefreshTime LAST_REFRESH_TIME获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithLastRefreshTime(lastRefreshTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_REFRESH_TIME"] = lastRefreshTime })
}

// WithFirstChangeTime FIRST_CHANGE_TIME获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithFirstChangeTime(firstChangeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["FIRST_CHANGE_TIME"] = firstChangeTime })
}

// WithLastChangeTime LAST_CHANGE_TIME获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithLastChangeTime(lastChangeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_CHANGE_TIME"] = lastChangeTime })
}

// WithOtherstat1ID OTHERSTAT_1_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat1ID(otherstat1ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_1_ID"] = otherstat1ID })
}

// WithOtherstat1Value OTHERSTAT_1_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat1Value(otherstat1Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_1_VALUE"] = otherstat1Value })
}

// WithOtherstat2ID OTHERSTAT_2_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat2ID(otherstat2ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_2_ID"] = otherstat2ID })
}

// WithOtherstat2Value OTHERSTAT_2_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat2Value(otherstat2Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_2_VALUE"] = otherstat2Value })
}

// WithOtherstat3ID OTHERSTAT_3_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat3ID(otherstat3ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_3_ID"] = otherstat3ID })
}

// WithOtherstat3Value OTHERSTAT_3_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat3Value(otherstat3Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_3_VALUE"] = otherstat3Value })
}

// WithOtherstat4ID OTHERSTAT_4_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat4ID(otherstat4ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_4_ID"] = otherstat4ID })
}

// WithOtherstat4Value OTHERSTAT_4_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat4Value(otherstat4Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_4_VALUE"] = otherstat4Value })
}

// WithOtherstat5ID OTHERSTAT_5_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat5ID(otherstat5ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_5_ID"] = otherstat5ID })
}

// WithOtherstat5Value OTHERSTAT_5_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat5Value(otherstat5Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_5_VALUE"] = otherstat5Value })
}

// WithOtherstat6ID OTHERSTAT_6_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat6ID(otherstat6ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_6_ID"] = otherstat6ID })
}

// WithOtherstat6Value OTHERSTAT_6_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat6Value(otherstat6Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_6_VALUE"] = otherstat6Value })
}

// WithOtherstat7ID OTHERSTAT_7_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat7ID(otherstat7ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_7_ID"] = otherstat7ID })
}

// WithOtherstat7Value OTHERSTAT_7_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat7Value(otherstat7Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_7_VALUE"] = otherstat7Value })
}

// WithOtherstat8ID OTHERSTAT_8_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat8ID(otherstat8ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_8_ID"] = otherstat8ID })
}

// WithOtherstat8Value OTHERSTAT_8_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat8Value(otherstat8Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_8_VALUE"] = otherstat8Value })
}

// WithOtherstat9ID OTHERSTAT_9_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat9ID(otherstat9ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_9_ID"] = otherstat9ID })
}

// WithOtherstat9Value OTHERSTAT_9_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat9Value(otherstat9Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_9_VALUE"] = otherstat9Value })
}

// WithOtherstat10ID OTHERSTAT_10_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat10ID(otherstat10ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_10_ID"] = otherstat10ID })
}

// WithOtherstat10Value OTHERSTAT_10_VALUE获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOtherstat10Value(otherstat10Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_10_VALUE"] = otherstat10Value })
}

// WithThreadID THREAD_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithThreadID(threadID int64) Option {
	return optionFunc(func(o *options) { o.query["THREAD_ID"] = threadID })
}

// WithPlanOperation PLAN_OPERATION获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithPlanOperation(planOperation string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_OPERATION"] = planOperation })
}

// WithStarts STARTS获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithStarts(starts int64) Option {
	return optionFunc(func(o *options) { o.query["STARTS"] = starts })
}

// WithOutputRows OUTPUT_ROWS获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithOutputRows(outputRows int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_ROWS"] = outputRows })
}

// WithPlanLineID PLAN_LINE_ID获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithPlanLineID(planLineID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_LINE_ID"] = planLineID })
}

// WithPlanDepth PLAN_DEPTH获取
func (obj *_AllVirtualSQLPlanMonitorMgr) WithPlanDepth(planDepth int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_DEPTH"] = planDepth })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLPlanMonitorMgr) GetByOption(opts ...Option) (result AllVirtualSQLPlanMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLPlanMonitorMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLPlanMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLPlanMonitorMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLPlanMonitor, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where(options.query)
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

// GetFromSvrIP 通过SVR_IP获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过SVR_PORT获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRequestID 通过REQUEST_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromRequestID(requestID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`REQUEST_ID` = ?", requestID).Find(&results).Error

	return
}

// GetBatchFromRequestID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromRequestID(requestIDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`REQUEST_ID` IN (?)", requestIDs).Find(&results).Error

	return
}

// GetFromTraceID 通过TRACE_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromTraceID(traceID string) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`TRACE_ID` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`TRACE_ID` IN (?)", traceIDs).Find(&results).Error

	return
}

// GetFromFirstRefreshTime 通过FIRST_REFRESH_TIME获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromFirstRefreshTime(firstRefreshTime time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`FIRST_REFRESH_TIME` = ?", firstRefreshTime).Find(&results).Error

	return
}

// GetBatchFromFirstRefreshTime 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromFirstRefreshTime(firstRefreshTimes []time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`FIRST_REFRESH_TIME` IN (?)", firstRefreshTimes).Find(&results).Error

	return
}

// GetFromLastRefreshTime 通过LAST_REFRESH_TIME获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromLastRefreshTime(lastRefreshTime time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`LAST_REFRESH_TIME` = ?", lastRefreshTime).Find(&results).Error

	return
}

// GetBatchFromLastRefreshTime 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromLastRefreshTime(lastRefreshTimes []time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`LAST_REFRESH_TIME` IN (?)", lastRefreshTimes).Find(&results).Error

	return
}

// GetFromFirstChangeTime 通过FIRST_CHANGE_TIME获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromFirstChangeTime(firstChangeTime time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`FIRST_CHANGE_TIME` = ?", firstChangeTime).Find(&results).Error

	return
}

// GetBatchFromFirstChangeTime 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromFirstChangeTime(firstChangeTimes []time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`FIRST_CHANGE_TIME` IN (?)", firstChangeTimes).Find(&results).Error

	return
}

// GetFromLastChangeTime 通过LAST_CHANGE_TIME获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromLastChangeTime(lastChangeTime time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`LAST_CHANGE_TIME` = ?", lastChangeTime).Find(&results).Error

	return
}

// GetBatchFromLastChangeTime 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromLastChangeTime(lastChangeTimes []time.Time) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`LAST_CHANGE_TIME` IN (?)", lastChangeTimes).Find(&results).Error

	return
}

// GetFromOtherstat1ID 通过OTHERSTAT_1_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat1ID(otherstat1ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_1_ID` = ?", otherstat1ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat1ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat1ID(otherstat1IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_1_ID` IN (?)", otherstat1IDs).Find(&results).Error

	return
}

// GetFromOtherstat1Value 通过OTHERSTAT_1_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat1Value(otherstat1Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_1_VALUE` = ?", otherstat1Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat1Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat1Value(otherstat1Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_1_VALUE` IN (?)", otherstat1Values).Find(&results).Error

	return
}

// GetFromOtherstat2ID 通过OTHERSTAT_2_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat2ID(otherstat2ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_2_ID` = ?", otherstat2ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat2ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat2ID(otherstat2IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_2_ID` IN (?)", otherstat2IDs).Find(&results).Error

	return
}

// GetFromOtherstat2Value 通过OTHERSTAT_2_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat2Value(otherstat2Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_2_VALUE` = ?", otherstat2Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat2Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat2Value(otherstat2Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_2_VALUE` IN (?)", otherstat2Values).Find(&results).Error

	return
}

// GetFromOtherstat3ID 通过OTHERSTAT_3_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat3ID(otherstat3ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_3_ID` = ?", otherstat3ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat3ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat3ID(otherstat3IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_3_ID` IN (?)", otherstat3IDs).Find(&results).Error

	return
}

// GetFromOtherstat3Value 通过OTHERSTAT_3_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat3Value(otherstat3Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_3_VALUE` = ?", otherstat3Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat3Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat3Value(otherstat3Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_3_VALUE` IN (?)", otherstat3Values).Find(&results).Error

	return
}

// GetFromOtherstat4ID 通过OTHERSTAT_4_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat4ID(otherstat4ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_4_ID` = ?", otherstat4ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat4ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat4ID(otherstat4IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_4_ID` IN (?)", otherstat4IDs).Find(&results).Error

	return
}

// GetFromOtherstat4Value 通过OTHERSTAT_4_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat4Value(otherstat4Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_4_VALUE` = ?", otherstat4Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat4Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat4Value(otherstat4Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_4_VALUE` IN (?)", otherstat4Values).Find(&results).Error

	return
}

// GetFromOtherstat5ID 通过OTHERSTAT_5_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat5ID(otherstat5ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_5_ID` = ?", otherstat5ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat5ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat5ID(otherstat5IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_5_ID` IN (?)", otherstat5IDs).Find(&results).Error

	return
}

// GetFromOtherstat5Value 通过OTHERSTAT_5_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat5Value(otherstat5Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_5_VALUE` = ?", otherstat5Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat5Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat5Value(otherstat5Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_5_VALUE` IN (?)", otherstat5Values).Find(&results).Error

	return
}

// GetFromOtherstat6ID 通过OTHERSTAT_6_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat6ID(otherstat6ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_6_ID` = ?", otherstat6ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat6ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat6ID(otherstat6IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_6_ID` IN (?)", otherstat6IDs).Find(&results).Error

	return
}

// GetFromOtherstat6Value 通过OTHERSTAT_6_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat6Value(otherstat6Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_6_VALUE` = ?", otherstat6Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat6Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat6Value(otherstat6Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_6_VALUE` IN (?)", otherstat6Values).Find(&results).Error

	return
}

// GetFromOtherstat7ID 通过OTHERSTAT_7_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat7ID(otherstat7ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_7_ID` = ?", otherstat7ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat7ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat7ID(otherstat7IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_7_ID` IN (?)", otherstat7IDs).Find(&results).Error

	return
}

// GetFromOtherstat7Value 通过OTHERSTAT_7_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat7Value(otherstat7Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_7_VALUE` = ?", otherstat7Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat7Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat7Value(otherstat7Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_7_VALUE` IN (?)", otherstat7Values).Find(&results).Error

	return
}

// GetFromOtherstat8ID 通过OTHERSTAT_8_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat8ID(otherstat8ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_8_ID` = ?", otherstat8ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat8ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat8ID(otherstat8IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_8_ID` IN (?)", otherstat8IDs).Find(&results).Error

	return
}

// GetFromOtherstat8Value 通过OTHERSTAT_8_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat8Value(otherstat8Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_8_VALUE` = ?", otherstat8Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat8Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat8Value(otherstat8Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_8_VALUE` IN (?)", otherstat8Values).Find(&results).Error

	return
}

// GetFromOtherstat9ID 通过OTHERSTAT_9_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat9ID(otherstat9ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_9_ID` = ?", otherstat9ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat9ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat9ID(otherstat9IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_9_ID` IN (?)", otherstat9IDs).Find(&results).Error

	return
}

// GetFromOtherstat9Value 通过OTHERSTAT_9_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat9Value(otherstat9Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_9_VALUE` = ?", otherstat9Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat9Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat9Value(otherstat9Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_9_VALUE` IN (?)", otherstat9Values).Find(&results).Error

	return
}

// GetFromOtherstat10ID 通过OTHERSTAT_10_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat10ID(otherstat10ID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_10_ID` = ?", otherstat10ID).Find(&results).Error

	return
}

// GetBatchFromOtherstat10ID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat10ID(otherstat10IDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_10_ID` IN (?)", otherstat10IDs).Find(&results).Error

	return
}

// GetFromOtherstat10Value 通过OTHERSTAT_10_VALUE获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOtherstat10Value(otherstat10Value int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_10_VALUE` = ?", otherstat10Value).Find(&results).Error

	return
}

// GetBatchFromOtherstat10Value 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOtherstat10Value(otherstat10Values []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OTHERSTAT_10_VALUE` IN (?)", otherstat10Values).Find(&results).Error

	return
}

// GetFromThreadID 通过THREAD_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromThreadID(threadID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`THREAD_ID` = ?", threadID).Find(&results).Error

	return
}

// GetBatchFromThreadID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromThreadID(threadIDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`THREAD_ID` IN (?)", threadIDs).Find(&results).Error

	return
}

// GetFromPlanOperation 通过PLAN_OPERATION获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromPlanOperation(planOperation string) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`PLAN_OPERATION` = ?", planOperation).Find(&results).Error

	return
}

// GetBatchFromPlanOperation 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromPlanOperation(planOperations []string) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`PLAN_OPERATION` IN (?)", planOperations).Find(&results).Error

	return
}

// GetFromStarts 通过STARTS获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromStarts(starts int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`STARTS` = ?", starts).Find(&results).Error

	return
}

// GetBatchFromStarts 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromStarts(startss []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`STARTS` IN (?)", startss).Find(&results).Error

	return
}

// GetFromOutputRows 通过OUTPUT_ROWS获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromOutputRows(outputRows int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OUTPUT_ROWS` = ?", outputRows).Find(&results).Error

	return
}

// GetBatchFromOutputRows 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromOutputRows(outputRowss []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`OUTPUT_ROWS` IN (?)", outputRowss).Find(&results).Error

	return
}

// GetFromPlanLineID 通过PLAN_LINE_ID获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromPlanLineID(planLineID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`PLAN_LINE_ID` = ?", planLineID).Find(&results).Error

	return
}

// GetBatchFromPlanLineID 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromPlanLineID(planLineIDs []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`PLAN_LINE_ID` IN (?)", planLineIDs).Find(&results).Error

	return
}

// GetFromPlanDepth 通过PLAN_DEPTH获取内容
func (obj *_AllVirtualSQLPlanMonitorMgr) GetFromPlanDepth(planDepth int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`PLAN_DEPTH` = ?", planDepth).Find(&results).Error

	return
}

// GetBatchFromPlanDepth 批量查找
func (obj *_AllVirtualSQLPlanMonitorMgr) GetBatchFromPlanDepth(planDepths []int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`PLAN_DEPTH` IN (?)", planDepths).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSQLPlanMonitorMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, requestID int64) (result AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`SVR_IP` = ? AND `SVR_PORT` = ? AND `TENANT_ID` = ? AND `REQUEST_ID` = ?", svrIP, svrPort, tenantID, requestID).First(&result).Error

	return
}

// FetchIndexByAllVirtualSQLPlanMonitorI1  获取多个内容
func (obj *_AllVirtualSQLPlanMonitorMgr) FetchIndexByAllVirtualSQLPlanMonitorI1(tenantID int64, requestID int64) (results []*AllVirtualSQLPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLPlanMonitor{}).Where("`TENANT_ID` = ? AND `REQUEST_ID` = ?", tenantID, requestID).Find(&results).Error

	return
}
