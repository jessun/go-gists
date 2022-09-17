package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualPxWorkerStatMgr struct {
	*_BaseMgr
}

// AllVirtualPxWorkerStatMgr open func
func AllVirtualPxWorkerStatMgr(db *gorm.DB) *_AllVirtualPxWorkerStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPxWorkerStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPxWorkerStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_px_worker_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPxWorkerStatMgr) GetTableName() string {
	return "__all_virtual_px_worker_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPxWorkerStatMgr) Reset() *_AllVirtualPxWorkerStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPxWorkerStatMgr) Get() (result AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPxWorkerStatMgr) Gets() (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPxWorkerStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSessionID session_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPxWorkerStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPxWorkerStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTraceID trace_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// WithQcID qc_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithQcID(qcID int64) Option {
	return optionFunc(func(o *options) { o.query["qc_id"] = qcID })
}

// WithSqcID sqc_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithSqcID(sqcID int64) Option {
	return optionFunc(func(o *options) { o.query["sqc_id"] = sqcID })
}

// WithWorkerID worker_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithWorkerID(workerID int64) Option {
	return optionFunc(func(o *options) { o.query["worker_id"] = workerID })
}

// WithDfoID dfo_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithDfoID(dfoID int64) Option {
	return optionFunc(func(o *options) { o.query["dfo_id"] = dfoID })
}

// WithStartTime start_time获取
func (obj *_AllVirtualPxWorkerStatMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithThreadID thread_id获取
func (obj *_AllVirtualPxWorkerStatMgr) WithThreadID(threadID int64) Option {
	return optionFunc(func(o *options) { o.query["thread_id"] = threadID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPxWorkerStatMgr) GetByOption(opts ...Option) (result AllVirtualPxWorkerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPxWorkerStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualPxWorkerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPxWorkerStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPxWorkerStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where(options.query)
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

// GetFromSessionID 通过session_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTraceID 通过trace_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromTraceID(traceID string) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`trace_id` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error

	return
}

// GetFromQcID 通过qc_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromQcID(qcID int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`qc_id` = ?", qcID).Find(&results).Error

	return
}

// GetBatchFromQcID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromQcID(qcIDs []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`qc_id` IN (?)", qcIDs).Find(&results).Error

	return
}

// GetFromSqcID 通过sqc_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromSqcID(sqcID int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`sqc_id` = ?", sqcID).Find(&results).Error

	return
}

// GetBatchFromSqcID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromSqcID(sqcIDs []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`sqc_id` IN (?)", sqcIDs).Find(&results).Error

	return
}

// GetFromWorkerID 通过worker_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromWorkerID(workerID int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`worker_id` = ?", workerID).Find(&results).Error

	return
}

// GetBatchFromWorkerID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromWorkerID(workerIDs []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`worker_id` IN (?)", workerIDs).Find(&results).Error

	return
}

// GetFromDfoID 通过dfo_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromDfoID(dfoID int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`dfo_id` = ?", dfoID).Find(&results).Error

	return
}

// GetBatchFromDfoID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromDfoID(dfoIDs []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`dfo_id` IN (?)", dfoIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromStartTime(startTime time.Time) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromThreadID 通过thread_id获取内容
func (obj *_AllVirtualPxWorkerStatMgr) GetFromThreadID(threadID int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`thread_id` = ?", threadID).Find(&results).Error

	return
}

// GetBatchFromThreadID 批量查找
func (obj *_AllVirtualPxWorkerStatMgr) GetBatchFromThreadID(threadIDs []int64) (results []*AllVirtualPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPxWorkerStat{}).Where("`thread_id` IN (?)", threadIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
