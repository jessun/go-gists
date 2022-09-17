package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSecurityAuditMgr struct {
	*_BaseMgr
}

// AllVirtualSecurityAuditMgr open func
func AllVirtualSecurityAuditMgr(db *gorm.DB) *_AllVirtualSecurityAuditMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSecurityAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSecurityAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_security_audit"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSecurityAuditMgr) GetTableName() string {
	return "__all_virtual_security_audit"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSecurityAuditMgr) Reset() *_AllVirtualSecurityAuditMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSecurityAuditMgr) Get() (result AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSecurityAuditMgr) Gets() (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSecurityAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSecurityAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithAuditID audit_id获取
func (obj *_AllVirtualSecurityAuditMgr) WithAuditID(auditID int64) Option {
	return optionFunc(func(o *options) { o.query["audit_id"] = auditID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSecurityAuditMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSecurityAuditMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithAuditType audit_type获取
func (obj *_AllVirtualSecurityAuditMgr) WithAuditType(auditType uint64) Option {
	return optionFunc(func(o *options) { o.query["audit_type"] = auditType })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualSecurityAuditMgr) WithOwnerID(ownerID uint64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithOperationType operation_type获取
func (obj *_AllVirtualSecurityAuditMgr) WithOperationType(operationType uint64) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithInSuccess in_success获取
func (obj *_AllVirtualSecurityAuditMgr) WithInSuccess(inSuccess uint64) Option {
	return optionFunc(func(o *options) { o.query["in_success"] = inSuccess })
}

// WithInFailure in_failure获取
func (obj *_AllVirtualSecurityAuditMgr) WithInFailure(inFailure uint64) Option {
	return optionFunc(func(o *options) { o.query["in_failure"] = inFailure })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSecurityAuditMgr) GetByOption(opts ...Option) (result AllVirtualSecurityAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSecurityAuditMgr) GetByOptions(opts ...Option) (results []*AllVirtualSecurityAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSecurityAuditMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSecurityAudit, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where(options.query)
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
func (obj *_AllVirtualSecurityAuditMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromAuditID 通过audit_id获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromAuditID(auditID int64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`audit_id` = ?", auditID).Find(&results).Error

	return
}

// GetBatchFromAuditID 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromAuditID(auditIDs []int64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`audit_id` IN (?)", auditIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromAuditType 通过audit_type获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromAuditType(auditType uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`audit_type` = ?", auditType).Find(&results).Error

	return
}

// GetBatchFromAuditType 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromAuditType(auditTypes []uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`audit_type` IN (?)", auditTypes).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromOwnerID(ownerID uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromOwnerID(ownerIDs []uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromOperationType(operationType uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromOperationType(operationTypes []uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromInSuccess 通过in_success获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromInSuccess(inSuccess uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`in_success` = ?", inSuccess).Find(&results).Error

	return
}

// GetBatchFromInSuccess 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromInSuccess(inSuccesss []uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`in_success` IN (?)", inSuccesss).Find(&results).Error

	return
}

// GetFromInFailure 通过in_failure获取内容
func (obj *_AllVirtualSecurityAuditMgr) GetFromInFailure(inFailure uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`in_failure` = ?", inFailure).Find(&results).Error

	return
}

// GetBatchFromInFailure 批量查找
func (obj *_AllVirtualSecurityAuditMgr) GetBatchFromInFailure(inFailures []uint64) (results []*AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`in_failure` IN (?)", inFailures).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSecurityAuditMgr) FetchByPrimaryKey(tenantID int64, auditID int64) (result AllVirtualSecurityAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSecurityAudit{}).Where("`tenant_id` = ? AND `audit_id` = ?", tenantID, auditID).First(&result).Error

	return
}
