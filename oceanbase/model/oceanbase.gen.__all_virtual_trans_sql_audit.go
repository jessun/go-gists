package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualTransSQLAuditMgr struct {
	*_BaseMgr
}

// AllVirtualTransSQLAuditMgr open func
func AllVirtualTransSQLAuditMgr(db *gorm.DB) *_AllVirtualTransSQLAuditMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransSQLAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransSQLAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_sql_audit"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransSQLAuditMgr) GetTableName() string {
	return "__all_virtual_trans_sql_audit"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransSQLAuditMgr) Reset() *_AllVirtualTransSQLAuditMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransSQLAuditMgr) Get() (result AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransSQLAuditMgr) Gets() (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransSQLAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTransSQLAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransSQLAuditMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransSQLAuditMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTransID trans_id获取
func (obj *_AllVirtualTransSQLAuditMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithPkey pkey获取
func (obj *_AllVirtualTransSQLAuditMgr) WithPkey(pkey string) Option {
	return optionFunc(func(o *options) { o.query["pkey"] = pkey })
}

// WithSQLNo sql_no获取
func (obj *_AllVirtualTransSQLAuditMgr) WithSQLNo(sqlNo int64) Option {
	return optionFunc(func(o *options) { o.query["sql_no"] = sqlNo })
}

// WithTraceID trace_id获取
func (obj *_AllVirtualTransSQLAuditMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// WithPhyPlanType phy_plan_type获取
func (obj *_AllVirtualTransSQLAuditMgr) WithPhyPlanType(phyPlanType int64) Option {
	return optionFunc(func(o *options) { o.query["phy_plan_type"] = phyPlanType })
}

// WithProxyReceiveUs proxy_receive_us获取
func (obj *_AllVirtualTransSQLAuditMgr) WithProxyReceiveUs(proxyReceiveUs int64) Option {
	return optionFunc(func(o *options) { o.query["proxy_receive_us"] = proxyReceiveUs })
}

// WithServerReceiveUs server_receive_us获取
func (obj *_AllVirtualTransSQLAuditMgr) WithServerReceiveUs(serverReceiveUs int64) Option {
	return optionFunc(func(o *options) { o.query["server_receive_us"] = serverReceiveUs })
}

// WithTransReceiveUs trans_receive_us获取
func (obj *_AllVirtualTransSQLAuditMgr) WithTransReceiveUs(transReceiveUs int64) Option {
	return optionFunc(func(o *options) { o.query["trans_receive_us"] = transReceiveUs })
}

// WithTransExecuteUs trans_execute_us获取
func (obj *_AllVirtualTransSQLAuditMgr) WithTransExecuteUs(transExecuteUs int64) Option {
	return optionFunc(func(o *options) { o.query["trans_execute_us"] = transExecuteUs })
}

// WithLockForReadCnt lock_for_read_cnt获取
func (obj *_AllVirtualTransSQLAuditMgr) WithLockForReadCnt(lockForReadCnt int64) Option {
	return optionFunc(func(o *options) { o.query["lock_for_read_cnt"] = lockForReadCnt })
}

// WithCtxAddr ctx_addr获取
func (obj *_AllVirtualTransSQLAuditMgr) WithCtxAddr(ctxAddr int64) Option {
	return optionFunc(func(o *options) { o.query["ctx_addr"] = ctxAddr })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransSQLAuditMgr) GetByOption(opts ...Option) (result AllVirtualTransSQLAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransSQLAuditMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransSQLAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransSQLAuditMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransSQLAudit, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where(options.query)
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
func (obj *_AllVirtualTransSQLAuditMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromTransID(transID string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromTransID(transIDs []string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromPkey 通过pkey获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromPkey(pkey string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`pkey` = ?", pkey).Find(&results).Error

	return
}

// GetBatchFromPkey 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromPkey(pkeys []string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`pkey` IN (?)", pkeys).Find(&results).Error

	return
}

// GetFromSQLNo 通过sql_no获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromSQLNo(sqlNo int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`sql_no` = ?", sqlNo).Find(&results).Error

	return
}

// GetBatchFromSQLNo 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromSQLNo(sqlNos []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`sql_no` IN (?)", sqlNos).Find(&results).Error

	return
}

// GetFromTraceID 通过trace_id获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromTraceID(traceID string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trace_id` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error

	return
}

// GetFromPhyPlanType 通过phy_plan_type获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromPhyPlanType(phyPlanType int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`phy_plan_type` = ?", phyPlanType).Find(&results).Error

	return
}

// GetBatchFromPhyPlanType 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromPhyPlanType(phyPlanTypes []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`phy_plan_type` IN (?)", phyPlanTypes).Find(&results).Error

	return
}

// GetFromProxyReceiveUs 通过proxy_receive_us获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromProxyReceiveUs(proxyReceiveUs int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`proxy_receive_us` = ?", proxyReceiveUs).Find(&results).Error

	return
}

// GetBatchFromProxyReceiveUs 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromProxyReceiveUs(proxyReceiveUss []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`proxy_receive_us` IN (?)", proxyReceiveUss).Find(&results).Error

	return
}

// GetFromServerReceiveUs 通过server_receive_us获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromServerReceiveUs(serverReceiveUs int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`server_receive_us` = ?", serverReceiveUs).Find(&results).Error

	return
}

// GetBatchFromServerReceiveUs 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromServerReceiveUs(serverReceiveUss []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`server_receive_us` IN (?)", serverReceiveUss).Find(&results).Error

	return
}

// GetFromTransReceiveUs 通过trans_receive_us获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromTransReceiveUs(transReceiveUs int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trans_receive_us` = ?", transReceiveUs).Find(&results).Error

	return
}

// GetBatchFromTransReceiveUs 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromTransReceiveUs(transReceiveUss []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trans_receive_us` IN (?)", transReceiveUss).Find(&results).Error

	return
}

// GetFromTransExecuteUs 通过trans_execute_us获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromTransExecuteUs(transExecuteUs int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trans_execute_us` = ?", transExecuteUs).Find(&results).Error

	return
}

// GetBatchFromTransExecuteUs 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromTransExecuteUs(transExecuteUss []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`trans_execute_us` IN (?)", transExecuteUss).Find(&results).Error

	return
}

// GetFromLockForReadCnt 通过lock_for_read_cnt获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromLockForReadCnt(lockForReadCnt int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`lock_for_read_cnt` = ?", lockForReadCnt).Find(&results).Error

	return
}

// GetBatchFromLockForReadCnt 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromLockForReadCnt(lockForReadCnts []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`lock_for_read_cnt` IN (?)", lockForReadCnts).Find(&results).Error

	return
}

// GetFromCtxAddr 通过ctx_addr获取内容
func (obj *_AllVirtualTransSQLAuditMgr) GetFromCtxAddr(ctxAddr int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`ctx_addr` = ?", ctxAddr).Find(&results).Error

	return
}

// GetBatchFromCtxAddr 批量查找
func (obj *_AllVirtualTransSQLAuditMgr) GetBatchFromCtxAddr(ctxAddrs []int64) (results []*AllVirtualTransSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransSQLAudit{}).Where("`ctx_addr` IN (?)", ctxAddrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
