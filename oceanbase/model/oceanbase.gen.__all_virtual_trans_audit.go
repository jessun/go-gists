package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTransAuditMgr struct {
	*_BaseMgr
}

// AllVirtualTransAuditMgr open func
func AllVirtualTransAuditMgr(db *gorm.DB) *_AllVirtualTransAuditMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_audit"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransAuditMgr) GetTableName() string {
	return "__all_virtual_trans_audit"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransAuditMgr) Reset() *_AllVirtualTransAuditMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransAuditMgr) Get() (result AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransAuditMgr) Gets() (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTransAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransAuditMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransAuditMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTransID trans_id获取
func (obj *_AllVirtualTransAuditMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithPkey pkey获取
func (obj *_AllVirtualTransAuditMgr) WithPkey(pkey string) Option {
	return optionFunc(func(o *options) { o.query["pkey"] = pkey })
}

// WithSessionID session_id获取
func (obj *_AllVirtualTransAuditMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithProxyID proxy_id获取
func (obj *_AllVirtualTransAuditMgr) WithProxyID(proxyID string) Option {
	return optionFunc(func(o *options) { o.query["proxy_id"] = proxyID })
}

// WithTransType trans_type获取
func (obj *_AllVirtualTransAuditMgr) WithTransType(transType int64) Option {
	return optionFunc(func(o *options) { o.query["trans_type"] = transType })
}

// WithCtxCreateTime ctx_create_time获取
func (obj *_AllVirtualTransAuditMgr) WithCtxCreateTime(ctxCreateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ctx_create_time"] = ctxCreateTime })
}

// WithExpiredTime expired_time获取
func (obj *_AllVirtualTransAuditMgr) WithExpiredTime(expiredTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expired_time"] = expiredTime })
}

// WithTransParam trans_param获取
func (obj *_AllVirtualTransAuditMgr) WithTransParam(transParam string) Option {
	return optionFunc(func(o *options) { o.query["trans_param"] = transParam })
}

// WithTotalSQLNo total_sql_no获取
func (obj *_AllVirtualTransAuditMgr) WithTotalSQLNo(totalSQLNo int64) Option {
	return optionFunc(func(o *options) { o.query["total_sql_no"] = totalSQLNo })
}

// WithRefer refer获取
func (obj *_AllVirtualTransAuditMgr) WithRefer(refer int64) Option {
	return optionFunc(func(o *options) { o.query["refer"] = refer })
}

// WithPrevTransArr prev_trans_arr获取
func (obj *_AllVirtualTransAuditMgr) WithPrevTransArr(prevTransArr string) Option {
	return optionFunc(func(o *options) { o.query["prev_trans_arr"] = prevTransArr })
}

// WithNextTransArr next_trans_arr获取
func (obj *_AllVirtualTransAuditMgr) WithNextTransArr(nextTransArr string) Option {
	return optionFunc(func(o *options) { o.query["next_trans_arr"] = nextTransArr })
}

// WithCtxAddr ctx_addr获取
func (obj *_AllVirtualTransAuditMgr) WithCtxAddr(ctxAddr int64) Option {
	return optionFunc(func(o *options) { o.query["ctx_addr"] = ctxAddr })
}

// WithCtxType ctx_type获取
func (obj *_AllVirtualTransAuditMgr) WithCtxType(ctxType int64) Option {
	return optionFunc(func(o *options) { o.query["ctx_type"] = ctxType })
}

// WithTraceLog trace_log获取
func (obj *_AllVirtualTransAuditMgr) WithTraceLog(traceLog string) Option {
	return optionFunc(func(o *options) { o.query["trace_log"] = traceLog })
}

// WithStatus status获取
func (obj *_AllVirtualTransAuditMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithForReplay for_replay获取
func (obj *_AllVirtualTransAuditMgr) WithForReplay(forReplay int8) Option {
	return optionFunc(func(o *options) { o.query["for_replay"] = forReplay })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransAuditMgr) GetByOption(opts ...Option) (result AllVirtualTransAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransAuditMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransAuditMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransAudit, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where(options.query)
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
func (obj *_AllVirtualTransAuditMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromTransID(transID string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromTransID(transIDs []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromPkey 通过pkey获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromPkey(pkey string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`pkey` = ?", pkey).Find(&results).Error

	return
}

// GetBatchFromPkey 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromPkey(pkeys []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`pkey` IN (?)", pkeys).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromProxyID 通过proxy_id获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromProxyID(proxyID string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`proxy_id` = ?", proxyID).Find(&results).Error

	return
}

// GetBatchFromProxyID 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromProxyID(proxyIDs []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`proxy_id` IN (?)", proxyIDs).Find(&results).Error

	return
}

// GetFromTransType 通过trans_type获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromTransType(transType int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trans_type` = ?", transType).Find(&results).Error

	return
}

// GetBatchFromTransType 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromTransType(transTypes []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trans_type` IN (?)", transTypes).Find(&results).Error

	return
}

// GetFromCtxCreateTime 通过ctx_create_time获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromCtxCreateTime(ctxCreateTime time.Time) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`ctx_create_time` = ?", ctxCreateTime).Find(&results).Error

	return
}

// GetBatchFromCtxCreateTime 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromCtxCreateTime(ctxCreateTimes []time.Time) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`ctx_create_time` IN (?)", ctxCreateTimes).Find(&results).Error

	return
}

// GetFromExpiredTime 通过expired_time获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromExpiredTime(expiredTime time.Time) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`expired_time` = ?", expiredTime).Find(&results).Error

	return
}

// GetBatchFromExpiredTime 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromExpiredTime(expiredTimes []time.Time) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`expired_time` IN (?)", expiredTimes).Find(&results).Error

	return
}

// GetFromTransParam 通过trans_param获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromTransParam(transParam string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trans_param` = ?", transParam).Find(&results).Error

	return
}

// GetBatchFromTransParam 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromTransParam(transParams []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trans_param` IN (?)", transParams).Find(&results).Error

	return
}

// GetFromTotalSQLNo 通过total_sql_no获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromTotalSQLNo(totalSQLNo int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`total_sql_no` = ?", totalSQLNo).Find(&results).Error

	return
}

// GetBatchFromTotalSQLNo 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromTotalSQLNo(totalSQLNos []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`total_sql_no` IN (?)", totalSQLNos).Find(&results).Error

	return
}

// GetFromRefer 通过refer获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromRefer(refer int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`refer` = ?", refer).Find(&results).Error

	return
}

// GetBatchFromRefer 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromRefer(refers []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`refer` IN (?)", refers).Find(&results).Error

	return
}

// GetFromPrevTransArr 通过prev_trans_arr获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromPrevTransArr(prevTransArr string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`prev_trans_arr` = ?", prevTransArr).Find(&results).Error

	return
}

// GetBatchFromPrevTransArr 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromPrevTransArr(prevTransArrs []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`prev_trans_arr` IN (?)", prevTransArrs).Find(&results).Error

	return
}

// GetFromNextTransArr 通过next_trans_arr获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromNextTransArr(nextTransArr string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`next_trans_arr` = ?", nextTransArr).Find(&results).Error

	return
}

// GetBatchFromNextTransArr 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromNextTransArr(nextTransArrs []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`next_trans_arr` IN (?)", nextTransArrs).Find(&results).Error

	return
}

// GetFromCtxAddr 通过ctx_addr获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromCtxAddr(ctxAddr int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`ctx_addr` = ?", ctxAddr).Find(&results).Error

	return
}

// GetBatchFromCtxAddr 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromCtxAddr(ctxAddrs []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`ctx_addr` IN (?)", ctxAddrs).Find(&results).Error

	return
}

// GetFromCtxType 通过ctx_type获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromCtxType(ctxType int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`ctx_type` = ?", ctxType).Find(&results).Error

	return
}

// GetBatchFromCtxType 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromCtxType(ctxTypes []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`ctx_type` IN (?)", ctxTypes).Find(&results).Error

	return
}

// GetFromTraceLog 通过trace_log获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromTraceLog(traceLog string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trace_log` = ?", traceLog).Find(&results).Error

	return
}

// GetBatchFromTraceLog 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromTraceLog(traceLogs []string) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`trace_log` IN (?)", traceLogs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromStatus(status int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromForReplay 通过for_replay获取内容
func (obj *_AllVirtualTransAuditMgr) GetFromForReplay(forReplay int8) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`for_replay` = ?", forReplay).Find(&results).Error

	return
}

// GetBatchFromForReplay 批量查找
func (obj *_AllVirtualTransAuditMgr) GetBatchFromForReplay(forReplays []int8) (results []*AllVirtualTransAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransAudit{}).Where("`for_replay` IN (?)", forReplays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
