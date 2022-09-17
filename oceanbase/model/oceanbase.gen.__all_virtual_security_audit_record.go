package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSecurityAuditRecordMgr struct {
	*_BaseMgr
}

// AllVirtualSecurityAuditRecordMgr open func
func AllVirtualSecurityAuditRecordMgr(db *gorm.DB) *_AllVirtualSecurityAuditRecordMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSecurityAuditRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSecurityAuditRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_security_audit_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSecurityAuditRecordMgr) GetTableName() string {
	return "__all_virtual_security_audit_record"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSecurityAuditRecordMgr) Reset() *_AllVirtualSecurityAuditRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSecurityAuditRecordMgr) Get() (result AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSecurityAuditRecordMgr) Gets() (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSecurityAuditRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithRecordTimestampUs record_timestamp_us获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithRecordTimestampUs(recordTimestampUs time.Time) Option {
	return optionFunc(func(o *options) { o.query["record_timestamp_us"] = recordTimestampUs })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithUserID user_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserName user_name获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithEffectiveUserID effective_user_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithEffectiveUserID(effectiveUserID uint64) Option {
	return optionFunc(func(o *options) { o.query["effective_user_id"] = effectiveUserID })
}

// WithEffectiveUserName effective_user_name获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithEffectiveUserName(effectiveUserName string) Option {
	return optionFunc(func(o *options) { o.query["effective_user_name"] = effectiveUserName })
}

// WithClientIP client_ip获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithClientIP(clientIP string) Option {
	return optionFunc(func(o *options) { o.query["client_ip"] = clientIP })
}

// WithUserClientIP user_client_ip获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithUserClientIP(userClientIP string) Option {
	return optionFunc(func(o *options) { o.query["user_client_ip"] = userClientIP })
}

// WithProxySessionID proxy_session_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithProxySessionID(proxySessionID uint64) Option {
	return optionFunc(func(o *options) { o.query["proxy_session_id"] = proxySessionID })
}

// WithSessionID session_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithSessionID(sessionID uint64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithEntryID entry_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithEntryID(entryID uint64) Option {
	return optionFunc(func(o *options) { o.query["entry_id"] = entryID })
}

// WithStatementID statement_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithStatementID(statementID uint64) Option {
	return optionFunc(func(o *options) { o.query["statement_id"] = statementID })
}

// WithTransID trans_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithCommitVersion commit_version获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithCommitVersion(commitVersion int64) Option {
	return optionFunc(func(o *options) { o.query["commit_version"] = commitVersion })
}

// WithTraceID trace_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// WithDbID db_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithDbID(dbID uint64) Option {
	return optionFunc(func(o *options) { o.query["db_id"] = dbID })
}

// WithCurDbID cur_db_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithCurDbID(curDbID uint64) Option {
	return optionFunc(func(o *options) { o.query["cur_db_id"] = curDbID })
}

// WithSQLTimestampUs sql_timestamp_us获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithSQLTimestampUs(sqlTimestampUs time.Time) Option {
	return optionFunc(func(o *options) { o.query["sql_timestamp_us"] = sqlTimestampUs })
}

// WithAuditID audit_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithAuditID(auditID uint64) Option {
	return optionFunc(func(o *options) { o.query["audit_id"] = auditID })
}

// WithAuditType audit_type获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithAuditType(auditType uint64) Option {
	return optionFunc(func(o *options) { o.query["audit_type"] = auditType })
}

// WithOperationType operation_type获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithOperationType(operationType uint64) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithActionID action_id获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithActionID(actionID uint64) Option {
	return optionFunc(func(o *options) { o.query["action_id"] = actionID })
}

// WithReturnCode return_code获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithReturnCode(returnCode int64) Option {
	return optionFunc(func(o *options) { o.query["return_code"] = returnCode })
}

// WithObjOwnerName obj_owner_name获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithObjOwnerName(objOwnerName string) Option {
	return optionFunc(func(o *options) { o.query["obj_owner_name"] = objOwnerName })
}

// WithObjName obj_name获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithObjName(objName string) Option {
	return optionFunc(func(o *options) { o.query["obj_name"] = objName })
}

// WithNewObjOwnerName new_obj_owner_name获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithNewObjOwnerName(newObjOwnerName string) Option {
	return optionFunc(func(o *options) { o.query["new_obj_owner_name"] = newObjOwnerName })
}

// WithNewObjName new_obj_name获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithNewObjName(newObjName string) Option {
	return optionFunc(func(o *options) { o.query["new_obj_name"] = newObjName })
}

// WithAuthPrivileges auth_privileges获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithAuthPrivileges(authPrivileges string) Option {
	return optionFunc(func(o *options) { o.query["auth_privileges"] = authPrivileges })
}

// WithAuthGrantee auth_grantee获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithAuthGrantee(authGrantee string) Option {
	return optionFunc(func(o *options) { o.query["auth_grantee"] = authGrantee })
}

// WithLogoffLogicalRead logoff_logical_read获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffLogicalRead(logoffLogicalRead uint64) Option {
	return optionFunc(func(o *options) { o.query["logoff_logical_read"] = logoffLogicalRead })
}

// WithLogoffPhysicalRead logoff_physical_read获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffPhysicalRead(logoffPhysicalRead uint64) Option {
	return optionFunc(func(o *options) { o.query["logoff_physical_read"] = logoffPhysicalRead })
}

// WithLogoffLogicalWrite logoff_logical_write获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffLogicalWrite(logoffLogicalWrite uint64) Option {
	return optionFunc(func(o *options) { o.query["logoff_logical_write"] = logoffLogicalWrite })
}

// WithLogoffLockCount logoff_lock_count获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffLockCount(logoffLockCount uint64) Option {
	return optionFunc(func(o *options) { o.query["logoff_lock_count"] = logoffLockCount })
}

// WithLogoffDeadLock logoff_dead_lock获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffDeadLock(logoffDeadLock string) Option {
	return optionFunc(func(o *options) { o.query["logoff_dead_lock"] = logoffDeadLock })
}

// WithLogoffCPUTimeUs logoff_cpu_time_us获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffCPUTimeUs(logoffCPUTimeUs uint64) Option {
	return optionFunc(func(o *options) { o.query["logoff_cpu_time_us"] = logoffCPUTimeUs })
}

// WithLogoffExecTimeUs logoff_exec_time_us获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffExecTimeUs(logoffExecTimeUs uint64) Option {
	return optionFunc(func(o *options) { o.query["logoff_exec_time_us"] = logoffExecTimeUs })
}

// WithLogoffAliveTimeUs logoff_alive_time_us获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithLogoffAliveTimeUs(logoffAliveTimeUs uint64) Option {
	return optionFunc(func(o *options) { o.query["logoff_alive_time_us"] = logoffAliveTimeUs })
}

// WithCommentText comment_text获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithCommentText(commentText string) Option {
	return optionFunc(func(o *options) { o.query["comment_text"] = commentText })
}

// WithSQLBind sql_bind获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithSQLBind(sqlBind string) Option {
	return optionFunc(func(o *options) { o.query["sql_bind"] = sqlBind })
}

// WithSQLText sql_text获取
func (obj *_AllVirtualSecurityAuditRecordMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSecurityAuditRecordMgr) GetByOption(opts ...Option) (result AllVirtualSecurityAuditRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSecurityAuditRecordMgr) GetByOptions(opts ...Option) (results []*AllVirtualSecurityAuditRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSecurityAuditRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSecurityAuditRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where(options.query)
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
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromRecordTimestampUs 通过record_timestamp_us获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromRecordTimestampUs(recordTimestampUs time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`record_timestamp_us` = ?", recordTimestampUs).Find(&results).Error

	return
}

// GetBatchFromRecordTimestampUs 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromRecordTimestampUs(recordTimestampUss []time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`record_timestamp_us` IN (?)", recordTimestampUss).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromUserID(userID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromUserID(userIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromUserName(userName string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromEffectiveUserID 通过effective_user_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromEffectiveUserID(effectiveUserID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`effective_user_id` = ?", effectiveUserID).Find(&results).Error

	return
}

// GetBatchFromEffectiveUserID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromEffectiveUserID(effectiveUserIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`effective_user_id` IN (?)", effectiveUserIDs).Find(&results).Error

	return
}

// GetFromEffectiveUserName 通过effective_user_name获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromEffectiveUserName(effectiveUserName string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`effective_user_name` = ?", effectiveUserName).Find(&results).Error

	return
}

// GetBatchFromEffectiveUserName 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromEffectiveUserName(effectiveUserNames []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`effective_user_name` IN (?)", effectiveUserNames).Find(&results).Error

	return
}

// GetFromClientIP 通过client_ip获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromClientIP(clientIP string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`client_ip` = ?", clientIP).Find(&results).Error

	return
}

// GetBatchFromClientIP 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromClientIP(clientIPs []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`client_ip` IN (?)", clientIPs).Find(&results).Error

	return
}

// GetFromUserClientIP 通过user_client_ip获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromUserClientIP(userClientIP string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`user_client_ip` = ?", userClientIP).Find(&results).Error

	return
}

// GetBatchFromUserClientIP 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromUserClientIP(userClientIPs []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`user_client_ip` IN (?)", userClientIPs).Find(&results).Error

	return
}

// GetFromProxySessionID 通过proxy_session_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromProxySessionID(proxySessionID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`proxy_session_id` = ?", proxySessionID).Find(&results).Error

	return
}

// GetBatchFromProxySessionID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromProxySessionID(proxySessionIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`proxy_session_id` IN (?)", proxySessionIDs).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromSessionID(sessionID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromSessionID(sessionIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromEntryID 通过entry_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromEntryID(entryID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`entry_id` = ?", entryID).Find(&results).Error

	return
}

// GetBatchFromEntryID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromEntryID(entryIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`entry_id` IN (?)", entryIDs).Find(&results).Error

	return
}

// GetFromStatementID 通过statement_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromStatementID(statementID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`statement_id` = ?", statementID).Find(&results).Error

	return
}

// GetBatchFromStatementID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromStatementID(statementIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`statement_id` IN (?)", statementIDs).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromTransID(transID string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromTransID(transIDs []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromCommitVersion 通过commit_version获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromCommitVersion(commitVersion int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`commit_version` = ?", commitVersion).Find(&results).Error

	return
}

// GetBatchFromCommitVersion 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromCommitVersion(commitVersions []int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`commit_version` IN (?)", commitVersions).Find(&results).Error

	return
}

// GetFromTraceID 通过trace_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromTraceID(traceID string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`trace_id` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error

	return
}

// GetFromDbID 通过db_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromDbID(dbID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`db_id` = ?", dbID).Find(&results).Error

	return
}

// GetBatchFromDbID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromDbID(dbIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`db_id` IN (?)", dbIDs).Find(&results).Error

	return
}

// GetFromCurDbID 通过cur_db_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromCurDbID(curDbID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`cur_db_id` = ?", curDbID).Find(&results).Error

	return
}

// GetBatchFromCurDbID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromCurDbID(curDbIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`cur_db_id` IN (?)", curDbIDs).Find(&results).Error

	return
}

// GetFromSQLTimestampUs 通过sql_timestamp_us获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromSQLTimestampUs(sqlTimestampUs time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`sql_timestamp_us` = ?", sqlTimestampUs).Find(&results).Error

	return
}

// GetBatchFromSQLTimestampUs 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromSQLTimestampUs(sqlTimestampUss []time.Time) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`sql_timestamp_us` IN (?)", sqlTimestampUss).Find(&results).Error

	return
}

// GetFromAuditID 通过audit_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromAuditID(auditID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`audit_id` = ?", auditID).Find(&results).Error

	return
}

// GetBatchFromAuditID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromAuditID(auditIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`audit_id` IN (?)", auditIDs).Find(&results).Error

	return
}

// GetFromAuditType 通过audit_type获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromAuditType(auditType uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`audit_type` = ?", auditType).Find(&results).Error

	return
}

// GetBatchFromAuditType 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromAuditType(auditTypes []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`audit_type` IN (?)", auditTypes).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromOperationType(operationType uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromOperationType(operationTypes []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromActionID 通过action_id获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromActionID(actionID uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`action_id` = ?", actionID).Find(&results).Error

	return
}

// GetBatchFromActionID 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromActionID(actionIDs []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`action_id` IN (?)", actionIDs).Find(&results).Error

	return
}

// GetFromReturnCode 通过return_code获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromReturnCode(returnCode int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`return_code` = ?", returnCode).Find(&results).Error

	return
}

// GetBatchFromReturnCode 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromReturnCode(returnCodes []int64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`return_code` IN (?)", returnCodes).Find(&results).Error

	return
}

// GetFromObjOwnerName 通过obj_owner_name获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromObjOwnerName(objOwnerName string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`obj_owner_name` = ?", objOwnerName).Find(&results).Error

	return
}

// GetBatchFromObjOwnerName 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromObjOwnerName(objOwnerNames []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`obj_owner_name` IN (?)", objOwnerNames).Find(&results).Error

	return
}

// GetFromObjName 通过obj_name获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromObjName(objName string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`obj_name` = ?", objName).Find(&results).Error

	return
}

// GetBatchFromObjName 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromObjName(objNames []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`obj_name` IN (?)", objNames).Find(&results).Error

	return
}

// GetFromNewObjOwnerName 通过new_obj_owner_name获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromNewObjOwnerName(newObjOwnerName string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`new_obj_owner_name` = ?", newObjOwnerName).Find(&results).Error

	return
}

// GetBatchFromNewObjOwnerName 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromNewObjOwnerName(newObjOwnerNames []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`new_obj_owner_name` IN (?)", newObjOwnerNames).Find(&results).Error

	return
}

// GetFromNewObjName 通过new_obj_name获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromNewObjName(newObjName string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`new_obj_name` = ?", newObjName).Find(&results).Error

	return
}

// GetBatchFromNewObjName 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromNewObjName(newObjNames []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`new_obj_name` IN (?)", newObjNames).Find(&results).Error

	return
}

// GetFromAuthPrivileges 通过auth_privileges获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromAuthPrivileges(authPrivileges string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`auth_privileges` = ?", authPrivileges).Find(&results).Error

	return
}

// GetBatchFromAuthPrivileges 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromAuthPrivileges(authPrivilegess []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`auth_privileges` IN (?)", authPrivilegess).Find(&results).Error

	return
}

// GetFromAuthGrantee 通过auth_grantee获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromAuthGrantee(authGrantee string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`auth_grantee` = ?", authGrantee).Find(&results).Error

	return
}

// GetBatchFromAuthGrantee 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromAuthGrantee(authGrantees []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`auth_grantee` IN (?)", authGrantees).Find(&results).Error

	return
}

// GetFromLogoffLogicalRead 通过logoff_logical_read获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffLogicalRead(logoffLogicalRead uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_logical_read` = ?", logoffLogicalRead).Find(&results).Error

	return
}

// GetBatchFromLogoffLogicalRead 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffLogicalRead(logoffLogicalReads []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_logical_read` IN (?)", logoffLogicalReads).Find(&results).Error

	return
}

// GetFromLogoffPhysicalRead 通过logoff_physical_read获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffPhysicalRead(logoffPhysicalRead uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_physical_read` = ?", logoffPhysicalRead).Find(&results).Error

	return
}

// GetBatchFromLogoffPhysicalRead 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffPhysicalRead(logoffPhysicalReads []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_physical_read` IN (?)", logoffPhysicalReads).Find(&results).Error

	return
}

// GetFromLogoffLogicalWrite 通过logoff_logical_write获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffLogicalWrite(logoffLogicalWrite uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_logical_write` = ?", logoffLogicalWrite).Find(&results).Error

	return
}

// GetBatchFromLogoffLogicalWrite 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffLogicalWrite(logoffLogicalWrites []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_logical_write` IN (?)", logoffLogicalWrites).Find(&results).Error

	return
}

// GetFromLogoffLockCount 通过logoff_lock_count获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffLockCount(logoffLockCount uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_lock_count` = ?", logoffLockCount).Find(&results).Error

	return
}

// GetBatchFromLogoffLockCount 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffLockCount(logoffLockCounts []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_lock_count` IN (?)", logoffLockCounts).Find(&results).Error

	return
}

// GetFromLogoffDeadLock 通过logoff_dead_lock获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffDeadLock(logoffDeadLock string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_dead_lock` = ?", logoffDeadLock).Find(&results).Error

	return
}

// GetBatchFromLogoffDeadLock 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffDeadLock(logoffDeadLocks []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_dead_lock` IN (?)", logoffDeadLocks).Find(&results).Error

	return
}

// GetFromLogoffCPUTimeUs 通过logoff_cpu_time_us获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffCPUTimeUs(logoffCPUTimeUs uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_cpu_time_us` = ?", logoffCPUTimeUs).Find(&results).Error

	return
}

// GetBatchFromLogoffCPUTimeUs 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffCPUTimeUs(logoffCPUTimeUss []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_cpu_time_us` IN (?)", logoffCPUTimeUss).Find(&results).Error

	return
}

// GetFromLogoffExecTimeUs 通过logoff_exec_time_us获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffExecTimeUs(logoffExecTimeUs uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_exec_time_us` = ?", logoffExecTimeUs).Find(&results).Error

	return
}

// GetBatchFromLogoffExecTimeUs 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffExecTimeUs(logoffExecTimeUss []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_exec_time_us` IN (?)", logoffExecTimeUss).Find(&results).Error

	return
}

// GetFromLogoffAliveTimeUs 通过logoff_alive_time_us获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromLogoffAliveTimeUs(logoffAliveTimeUs uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_alive_time_us` = ?", logoffAliveTimeUs).Find(&results).Error

	return
}

// GetBatchFromLogoffAliveTimeUs 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromLogoffAliveTimeUs(logoffAliveTimeUss []uint64) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`logoff_alive_time_us` IN (?)", logoffAliveTimeUss).Find(&results).Error

	return
}

// GetFromCommentText 通过comment_text获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromCommentText(commentText string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`comment_text` = ?", commentText).Find(&results).Error

	return
}

// GetBatchFromCommentText 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromCommentText(commentTexts []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`comment_text` IN (?)", commentTexts).Find(&results).Error

	return
}

// GetFromSQLBind 通过sql_bind获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromSQLBind(sqlBind string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`sql_bind` = ?", sqlBind).Find(&results).Error

	return
}

// GetBatchFromSQLBind 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromSQLBind(sqlBinds []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`sql_bind` IN (?)", sqlBinds).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllVirtualSecurityAuditRecordMgr) GetFromSQLText(sqlText string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllVirtualSecurityAuditRecordMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSecurityAuditRecordMgr) FetchByPrimaryKey(tenantID int64, svrIP string, svrPort int64, recordTimestampUs time.Time) (result AllVirtualSecurityAuditRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAuditRecord{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `record_timestamp_us` = ?", tenantID, svrIP, svrPort, recordTimestampUs).First(&result).Error

	return
}
