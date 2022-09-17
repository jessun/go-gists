package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTransStatMgr struct {
	*_BaseMgr
}

// AllVirtualTransStatMgr open func
func AllVirtualTransStatMgr(db *gorm.DB) *_AllVirtualTransStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransStatMgr) GetTableName() string {
	return "__all_virtual_trans_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransStatMgr) Reset() *_AllVirtualTransStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransStatMgr) Get() (result AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransStatMgr) Gets() (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTransStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithIncNum inc_num获取
func (obj *_AllVirtualTransStatMgr) WithIncNum(incNum int64) Option {
	return optionFunc(func(o *options) { o.query["inc_num"] = incNum })
}

// WithSessionID session_id获取
func (obj *_AllVirtualTransStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithProxyID proxy_id获取
func (obj *_AllVirtualTransStatMgr) WithProxyID(proxyID string) Option {
	return optionFunc(func(o *options) { o.query["proxy_id"] = proxyID })
}

// WithTransType trans_type获取
func (obj *_AllVirtualTransStatMgr) WithTransType(transType int64) Option {
	return optionFunc(func(o *options) { o.query["trans_type"] = transType })
}

// WithTransID trans_id获取
func (obj *_AllVirtualTransStatMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithIsExiting is_exiting获取
func (obj *_AllVirtualTransStatMgr) WithIsExiting(isExiting int64) Option {
	return optionFunc(func(o *options) { o.query["is_exiting"] = isExiting })
}

// WithIsReadonly is_readonly获取
func (obj *_AllVirtualTransStatMgr) WithIsReadonly(isReadonly int64) Option {
	return optionFunc(func(o *options) { o.query["is_readonly"] = isReadonly })
}

// WithIsDecided is_decided获取
func (obj *_AllVirtualTransStatMgr) WithIsDecided(isDecided int64) Option {
	return optionFunc(func(o *options) { o.query["is_decided"] = isDecided })
}

// WithTransMode trans_mode获取
func (obj *_AllVirtualTransStatMgr) WithTransMode(transMode string) Option {
	return optionFunc(func(o *options) { o.query["trans_mode"] = transMode })
}

// WithActiveMemstoreVersion active_memstore_version获取
func (obj *_AllVirtualTransStatMgr) WithActiveMemstoreVersion(activeMemstoreVersion string) Option {
	return optionFunc(func(o *options) { o.query["active_memstore_version"] = activeMemstoreVersion })
}

// WithPartition partition获取
func (obj *_AllVirtualTransStatMgr) WithPartition(partition string) Option {
	return optionFunc(func(o *options) { o.query["partition"] = partition })
}

// WithParticipants participants获取
func (obj *_AllVirtualTransStatMgr) WithParticipants(participants string) Option {
	return optionFunc(func(o *options) { o.query["participants"] = participants })
}

// WithAutocommit autocommit获取
func (obj *_AllVirtualTransStatMgr) WithAutocommit(autocommit int64) Option {
	return optionFunc(func(o *options) { o.query["autocommit"] = autocommit })
}

// WithTransConsistency trans_consistency获取
func (obj *_AllVirtualTransStatMgr) WithTransConsistency(transConsistency int64) Option {
	return optionFunc(func(o *options) { o.query["trans_consistency"] = transConsistency })
}

// WithCtxCreateTime ctx_create_time获取
func (obj *_AllVirtualTransStatMgr) WithCtxCreateTime(ctxCreateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ctx_create_time"] = ctxCreateTime })
}

// WithExpiredTime expired_time获取
func (obj *_AllVirtualTransStatMgr) WithExpiredTime(expiredTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expired_time"] = expiredTime })
}

// WithRefer refer获取
func (obj *_AllVirtualTransStatMgr) WithRefer(refer int64) Option {
	return optionFunc(func(o *options) { o.query["refer"] = refer })
}

// WithSQLNo sql_no获取
func (obj *_AllVirtualTransStatMgr) WithSQLNo(sqlNo int64) Option {
	return optionFunc(func(o *options) { o.query["sql_no"] = sqlNo })
}

// WithState state获取
func (obj *_AllVirtualTransStatMgr) WithState(state int64) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithPartTransAction part_trans_action获取
func (obj *_AllVirtualTransStatMgr) WithPartTransAction(partTransAction int64) Option {
	return optionFunc(func(o *options) { o.query["part_trans_action"] = partTransAction })
}

// WithLockForReadRetryCount lock_for_read_retry_count获取
func (obj *_AllVirtualTransStatMgr) WithLockForReadRetryCount(lockForReadRetryCount int64) Option {
	return optionFunc(func(o *options) { o.query["lock_for_read_retry_count"] = lockForReadRetryCount })
}

// WithCtxAddr ctx_addr获取
func (obj *_AllVirtualTransStatMgr) WithCtxAddr(ctxAddr int64) Option {
	return optionFunc(func(o *options) { o.query["ctx_addr"] = ctxAddr })
}

// WithPrevTransArr prev_trans_arr获取
func (obj *_AllVirtualTransStatMgr) WithPrevTransArr(prevTransArr string) Option {
	return optionFunc(func(o *options) { o.query["prev_trans_arr"] = prevTransArr })
}

// WithPrevTransCount prev_trans_count获取
func (obj *_AllVirtualTransStatMgr) WithPrevTransCount(prevTransCount int64) Option {
	return optionFunc(func(o *options) { o.query["prev_trans_count"] = prevTransCount })
}

// WithNextTransArr next_trans_arr获取
func (obj *_AllVirtualTransStatMgr) WithNextTransArr(nextTransArr string) Option {
	return optionFunc(func(o *options) { o.query["next_trans_arr"] = nextTransArr })
}

// WithNextTransCount next_trans_count获取
func (obj *_AllVirtualTransStatMgr) WithNextTransCount(nextTransCount int64) Option {
	return optionFunc(func(o *options) { o.query["next_trans_count"] = nextTransCount })
}

// WithPrevTransCommitCount prev_trans_commit_count获取
func (obj *_AllVirtualTransStatMgr) WithPrevTransCommitCount(prevTransCommitCount int64) Option {
	return optionFunc(func(o *options) { o.query["prev_trans_commit_count"] = prevTransCommitCount })
}

// WithCtxID ctx_id获取
func (obj *_AllVirtualTransStatMgr) WithCtxID(ctxID int64) Option {
	return optionFunc(func(o *options) { o.query["ctx_id"] = ctxID })
}

// WithPendingLogSize pending_log_size获取
func (obj *_AllVirtualTransStatMgr) WithPendingLogSize(pendingLogSize int64) Option {
	return optionFunc(func(o *options) { o.query["pending_log_size"] = pendingLogSize })
}

// WithFlushedLogSize flushed_log_size获取
func (obj *_AllVirtualTransStatMgr) WithFlushedLogSize(flushedLogSize int64) Option {
	return optionFunc(func(o *options) { o.query["flushed_log_size"] = flushedLogSize })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransStatMgr) GetByOption(opts ...Option) (result AllVirtualTransStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where(options.query)
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
func (obj *_AllVirtualTransStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTransStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromIncNum 通过inc_num获取内容
func (obj *_AllVirtualTransStatMgr) GetFromIncNum(incNum int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`inc_num` = ?", incNum).Find(&results).Error

	return
}

// GetBatchFromIncNum 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromIncNum(incNums []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`inc_num` IN (?)", incNums).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllVirtualTransStatMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromProxyID 通过proxy_id获取内容
func (obj *_AllVirtualTransStatMgr) GetFromProxyID(proxyID string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`proxy_id` = ?", proxyID).Find(&results).Error

	return
}

// GetBatchFromProxyID 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromProxyID(proxyIDs []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`proxy_id` IN (?)", proxyIDs).Find(&results).Error

	return
}

// GetFromTransType 通过trans_type获取内容
func (obj *_AllVirtualTransStatMgr) GetFromTransType(transType int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_type` = ?", transType).Find(&results).Error

	return
}

// GetBatchFromTransType 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromTransType(transTypes []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_type` IN (?)", transTypes).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualTransStatMgr) GetFromTransID(transID string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromTransID(transIDs []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromIsExiting 通过is_exiting获取内容
func (obj *_AllVirtualTransStatMgr) GetFromIsExiting(isExiting int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`is_exiting` = ?", isExiting).Find(&results).Error

	return
}

// GetBatchFromIsExiting 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromIsExiting(isExitings []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`is_exiting` IN (?)", isExitings).Find(&results).Error

	return
}

// GetFromIsReadonly 通过is_readonly获取内容
func (obj *_AllVirtualTransStatMgr) GetFromIsReadonly(isReadonly int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`is_readonly` = ?", isReadonly).Find(&results).Error

	return
}

// GetBatchFromIsReadonly 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromIsReadonly(isReadonlys []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`is_readonly` IN (?)", isReadonlys).Find(&results).Error

	return
}

// GetFromIsDecided 通过is_decided获取内容
func (obj *_AllVirtualTransStatMgr) GetFromIsDecided(isDecided int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`is_decided` = ?", isDecided).Find(&results).Error

	return
}

// GetBatchFromIsDecided 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromIsDecided(isDecideds []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`is_decided` IN (?)", isDecideds).Find(&results).Error

	return
}

// GetFromTransMode 通过trans_mode获取内容
func (obj *_AllVirtualTransStatMgr) GetFromTransMode(transMode string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_mode` = ?", transMode).Find(&results).Error

	return
}

// GetBatchFromTransMode 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromTransMode(transModes []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_mode` IN (?)", transModes).Find(&results).Error

	return
}

// GetFromActiveMemstoreVersion 通过active_memstore_version获取内容
func (obj *_AllVirtualTransStatMgr) GetFromActiveMemstoreVersion(activeMemstoreVersion string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`active_memstore_version` = ?", activeMemstoreVersion).Find(&results).Error

	return
}

// GetBatchFromActiveMemstoreVersion 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromActiveMemstoreVersion(activeMemstoreVersions []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`active_memstore_version` IN (?)", activeMemstoreVersions).Find(&results).Error

	return
}

// GetFromPartition 通过partition获取内容
func (obj *_AllVirtualTransStatMgr) GetFromPartition(partition string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`partition` = ?", partition).Find(&results).Error

	return
}

// GetBatchFromPartition 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromPartition(partitions []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`partition` IN (?)", partitions).Find(&results).Error

	return
}

// GetFromParticipants 通过participants获取内容
func (obj *_AllVirtualTransStatMgr) GetFromParticipants(participants string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`participants` = ?", participants).Find(&results).Error

	return
}

// GetBatchFromParticipants 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromParticipants(participantss []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`participants` IN (?)", participantss).Find(&results).Error

	return
}

// GetFromAutocommit 通过autocommit获取内容
func (obj *_AllVirtualTransStatMgr) GetFromAutocommit(autocommit int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`autocommit` = ?", autocommit).Find(&results).Error

	return
}

// GetBatchFromAutocommit 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromAutocommit(autocommits []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`autocommit` IN (?)", autocommits).Find(&results).Error

	return
}

// GetFromTransConsistency 通过trans_consistency获取内容
func (obj *_AllVirtualTransStatMgr) GetFromTransConsistency(transConsistency int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_consistency` = ?", transConsistency).Find(&results).Error

	return
}

// GetBatchFromTransConsistency 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromTransConsistency(transConsistencys []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`trans_consistency` IN (?)", transConsistencys).Find(&results).Error

	return
}

// GetFromCtxCreateTime 通过ctx_create_time获取内容
func (obj *_AllVirtualTransStatMgr) GetFromCtxCreateTime(ctxCreateTime time.Time) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`ctx_create_time` = ?", ctxCreateTime).Find(&results).Error

	return
}

// GetBatchFromCtxCreateTime 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromCtxCreateTime(ctxCreateTimes []time.Time) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`ctx_create_time` IN (?)", ctxCreateTimes).Find(&results).Error

	return
}

// GetFromExpiredTime 通过expired_time获取内容
func (obj *_AllVirtualTransStatMgr) GetFromExpiredTime(expiredTime time.Time) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`expired_time` = ?", expiredTime).Find(&results).Error

	return
}

// GetBatchFromExpiredTime 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromExpiredTime(expiredTimes []time.Time) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`expired_time` IN (?)", expiredTimes).Find(&results).Error

	return
}

// GetFromRefer 通过refer获取内容
func (obj *_AllVirtualTransStatMgr) GetFromRefer(refer int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`refer` = ?", refer).Find(&results).Error

	return
}

// GetBatchFromRefer 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromRefer(refers []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`refer` IN (?)", refers).Find(&results).Error

	return
}

// GetFromSQLNo 通过sql_no获取内容
func (obj *_AllVirtualTransStatMgr) GetFromSQLNo(sqlNo int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`sql_no` = ?", sqlNo).Find(&results).Error

	return
}

// GetBatchFromSQLNo 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromSQLNo(sqlNos []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`sql_no` IN (?)", sqlNos).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_AllVirtualTransStatMgr) GetFromState(state int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromState(states []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromPartTransAction 通过part_trans_action获取内容
func (obj *_AllVirtualTransStatMgr) GetFromPartTransAction(partTransAction int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`part_trans_action` = ?", partTransAction).Find(&results).Error

	return
}

// GetBatchFromPartTransAction 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromPartTransAction(partTransActions []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`part_trans_action` IN (?)", partTransActions).Find(&results).Error

	return
}

// GetFromLockForReadRetryCount 通过lock_for_read_retry_count获取内容
func (obj *_AllVirtualTransStatMgr) GetFromLockForReadRetryCount(lockForReadRetryCount int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`lock_for_read_retry_count` = ?", lockForReadRetryCount).Find(&results).Error

	return
}

// GetBatchFromLockForReadRetryCount 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromLockForReadRetryCount(lockForReadRetryCounts []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`lock_for_read_retry_count` IN (?)", lockForReadRetryCounts).Find(&results).Error

	return
}

// GetFromCtxAddr 通过ctx_addr获取内容
func (obj *_AllVirtualTransStatMgr) GetFromCtxAddr(ctxAddr int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`ctx_addr` = ?", ctxAddr).Find(&results).Error

	return
}

// GetBatchFromCtxAddr 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromCtxAddr(ctxAddrs []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`ctx_addr` IN (?)", ctxAddrs).Find(&results).Error

	return
}

// GetFromPrevTransArr 通过prev_trans_arr获取内容
func (obj *_AllVirtualTransStatMgr) GetFromPrevTransArr(prevTransArr string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`prev_trans_arr` = ?", prevTransArr).Find(&results).Error

	return
}

// GetBatchFromPrevTransArr 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromPrevTransArr(prevTransArrs []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`prev_trans_arr` IN (?)", prevTransArrs).Find(&results).Error

	return
}

// GetFromPrevTransCount 通过prev_trans_count获取内容
func (obj *_AllVirtualTransStatMgr) GetFromPrevTransCount(prevTransCount int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`prev_trans_count` = ?", prevTransCount).Find(&results).Error

	return
}

// GetBatchFromPrevTransCount 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromPrevTransCount(prevTransCounts []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`prev_trans_count` IN (?)", prevTransCounts).Find(&results).Error

	return
}

// GetFromNextTransArr 通过next_trans_arr获取内容
func (obj *_AllVirtualTransStatMgr) GetFromNextTransArr(nextTransArr string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`next_trans_arr` = ?", nextTransArr).Find(&results).Error

	return
}

// GetBatchFromNextTransArr 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromNextTransArr(nextTransArrs []string) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`next_trans_arr` IN (?)", nextTransArrs).Find(&results).Error

	return
}

// GetFromNextTransCount 通过next_trans_count获取内容
func (obj *_AllVirtualTransStatMgr) GetFromNextTransCount(nextTransCount int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`next_trans_count` = ?", nextTransCount).Find(&results).Error

	return
}

// GetBatchFromNextTransCount 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromNextTransCount(nextTransCounts []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`next_trans_count` IN (?)", nextTransCounts).Find(&results).Error

	return
}

// GetFromPrevTransCommitCount 通过prev_trans_commit_count获取内容
func (obj *_AllVirtualTransStatMgr) GetFromPrevTransCommitCount(prevTransCommitCount int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`prev_trans_commit_count` = ?", prevTransCommitCount).Find(&results).Error

	return
}

// GetBatchFromPrevTransCommitCount 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromPrevTransCommitCount(prevTransCommitCounts []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`prev_trans_commit_count` IN (?)", prevTransCommitCounts).Find(&results).Error

	return
}

// GetFromCtxID 通过ctx_id获取内容
func (obj *_AllVirtualTransStatMgr) GetFromCtxID(ctxID int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`ctx_id` = ?", ctxID).Find(&results).Error

	return
}

// GetBatchFromCtxID 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromCtxID(ctxIDs []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`ctx_id` IN (?)", ctxIDs).Find(&results).Error

	return
}

// GetFromPendingLogSize 通过pending_log_size获取内容
func (obj *_AllVirtualTransStatMgr) GetFromPendingLogSize(pendingLogSize int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`pending_log_size` = ?", pendingLogSize).Find(&results).Error

	return
}

// GetBatchFromPendingLogSize 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromPendingLogSize(pendingLogSizes []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`pending_log_size` IN (?)", pendingLogSizes).Find(&results).Error

	return
}

// GetFromFlushedLogSize 通过flushed_log_size获取内容
func (obj *_AllVirtualTransStatMgr) GetFromFlushedLogSize(flushedLogSize int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`flushed_log_size` = ?", flushedLogSize).Find(&results).Error

	return
}

// GetBatchFromFlushedLogSize 批量查找
func (obj *_AllVirtualTransStatMgr) GetBatchFromFlushedLogSize(flushedLogSizes []int64) (results []*AllVirtualTransStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransStat{}).Where("`flushed_log_size` IN (?)", flushedLogSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
