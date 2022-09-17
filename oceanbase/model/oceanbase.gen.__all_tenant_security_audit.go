package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantSecurityAuditMgr struct {
	*_BaseMgr
}

// AllTenantSecurityAuditMgr open func
func AllTenantSecurityAuditMgr(db *gorm.DB) *_AllTenantSecurityAuditMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantSecurityAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantSecurityAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_security_audit"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantSecurityAuditMgr) GetTableName() string {
	return "__all_tenant_security_audit"
}

// Reset 重置gorm会话
func (obj *_AllTenantSecurityAuditMgr) Reset() *_AllTenantSecurityAuditMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantSecurityAuditMgr) Get() (result AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantSecurityAuditMgr) Gets() (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantSecurityAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantSecurityAuditMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantSecurityAuditMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantSecurityAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithAuditID audit_id获取
func (obj *_AllTenantSecurityAuditMgr) WithAuditID(auditID int64) Option {
	return optionFunc(func(o *options) { o.query["audit_id"] = auditID })
}

// WithAuditType audit_type获取
func (obj *_AllTenantSecurityAuditMgr) WithAuditType(auditType uint64) Option {
	return optionFunc(func(o *options) { o.query["audit_type"] = auditType })
}

// WithOwnerID owner_id获取
func (obj *_AllTenantSecurityAuditMgr) WithOwnerID(ownerID uint64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithOperationType operation_type获取
func (obj *_AllTenantSecurityAuditMgr) WithOperationType(operationType uint64) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithInSuccess in_success获取
func (obj *_AllTenantSecurityAuditMgr) WithInSuccess(inSuccess uint64) Option {
	return optionFunc(func(o *options) { o.query["in_success"] = inSuccess })
}

// WithInFailure in_failure获取
func (obj *_AllTenantSecurityAuditMgr) WithInFailure(inFailure uint64) Option {
	return optionFunc(func(o *options) { o.query["in_failure"] = inFailure })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantSecurityAuditMgr) GetByOption(opts ...Option) (result AllTenantSecurityAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantSecurityAuditMgr) GetByOptions(opts ...Option) (results []*AllTenantSecurityAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantSecurityAuditMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantSecurityAudit, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromTenantID(tenantID int64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromAuditID 通过audit_id获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromAuditID(auditID int64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`audit_id` = ?", auditID).Find(&results).Error

	return
}

// GetBatchFromAuditID 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromAuditID(auditIDs []int64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`audit_id` IN (?)", auditIDs).Find(&results).Error

	return
}

// GetFromAuditType 通过audit_type获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromAuditType(auditType uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`audit_type` = ?", auditType).Find(&results).Error

	return
}

// GetBatchFromAuditType 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromAuditType(auditTypes []uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`audit_type` IN (?)", auditTypes).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromOwnerID(ownerID uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromOwnerID(ownerIDs []uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromOperationType(operationType uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromOperationType(operationTypes []uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromInSuccess 通过in_success获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromInSuccess(inSuccess uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`in_success` = ?", inSuccess).Find(&results).Error

	return
}

// GetBatchFromInSuccess 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromInSuccess(inSuccesss []uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`in_success` IN (?)", inSuccesss).Find(&results).Error

	return
}

// GetFromInFailure 通过in_failure获取内容
func (obj *_AllTenantSecurityAuditMgr) GetFromInFailure(inFailure uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`in_failure` = ?", inFailure).Find(&results).Error

	return
}

// GetBatchFromInFailure 批量查找
func (obj *_AllTenantSecurityAuditMgr) GetBatchFromInFailure(inFailures []uint64) (results []*AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`in_failure` IN (?)", inFailures).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantSecurityAuditMgr) FetchByPrimaryKey(tenantID int64, auditID int64) (result AllTenantSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAudit{}).Where("`tenant_id` = ? AND `audit_id` = ?", tenantID, auditID).First(&result).Error

	return
}
