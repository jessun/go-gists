package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantSecurityAuditHistoryMgr struct {
	*_BaseMgr
}

// AllTenantSecurityAuditHistoryMgr open func
func AllTenantSecurityAuditHistoryMgr(db *gorm.DB) *_AllTenantSecurityAuditHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantSecurityAuditHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantSecurityAuditHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_security_audit_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantSecurityAuditHistoryMgr) GetTableName() string {
	return "__all_tenant_security_audit_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantSecurityAuditHistoryMgr) Reset() *_AllTenantSecurityAuditHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantSecurityAuditHistoryMgr) Get() (result AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantSecurityAuditHistoryMgr) Gets() (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantSecurityAuditHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithAuditID audit_id获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithAuditID(auditID int64) Option {
	return optionFunc(func(o *options) { o.query["audit_id"] = auditID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithAuditType audit_type获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithAuditType(auditType uint64) Option {
	return optionFunc(func(o *options) { o.query["audit_type"] = auditType })
}

// WithOwnerID owner_id获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithOwnerID(ownerID uint64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithOperationType operation_type获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithOperationType(operationType uint64) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithInSuccess in_success获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithInSuccess(inSuccess uint64) Option {
	return optionFunc(func(o *options) { o.query["in_success"] = inSuccess })
}

// WithInFailure in_failure获取
func (obj *_AllTenantSecurityAuditHistoryMgr) WithInFailure(inFailure uint64) Option {
	return optionFunc(func(o *options) { o.query["in_failure"] = inFailure })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantSecurityAuditHistoryMgr) GetByOption(opts ...Option) (result AllTenantSecurityAuditHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantSecurityAuditHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantSecurityAuditHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantSecurityAuditHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantSecurityAuditHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where(options.query)
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
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromAuditID 通过audit_id获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromAuditID(auditID int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`audit_id` = ?", auditID).Find(&results).Error

	return
}

// GetBatchFromAuditID 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromAuditID(auditIDs []int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`audit_id` IN (?)", auditIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromAuditType 通过audit_type获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromAuditType(auditType uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`audit_type` = ?", auditType).Find(&results).Error

	return
}

// GetBatchFromAuditType 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromAuditType(auditTypes []uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`audit_type` IN (?)", auditTypes).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromOwnerID(ownerID uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromOwnerID(ownerIDs []uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromOperationType(operationType uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromOperationType(operationTypes []uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromInSuccess 通过in_success获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromInSuccess(inSuccess uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`in_success` = ?", inSuccess).Find(&results).Error

	return
}

// GetBatchFromInSuccess 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromInSuccess(inSuccesss []uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`in_success` IN (?)", inSuccesss).Find(&results).Error

	return
}

// GetFromInFailure 通过in_failure获取内容
func (obj *_AllTenantSecurityAuditHistoryMgr) GetFromInFailure(inFailure uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`in_failure` = ?", inFailure).Find(&results).Error

	return
}

// GetBatchFromInFailure 批量查找
func (obj *_AllTenantSecurityAuditHistoryMgr) GetBatchFromInFailure(inFailures []uint64) (results []*AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`in_failure` IN (?)", inFailures).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantSecurityAuditHistoryMgr) FetchByPrimaryKey(tenantID int64, auditID int64, schemaVersion int64) (result AllTenantSecurityAuditHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSecurityAuditHistory{}).Where("`tenant_id` = ? AND `audit_id` = ? AND `schema_version` = ?", tenantID, auditID, schemaVersion).First(&result).Error

	return
}
