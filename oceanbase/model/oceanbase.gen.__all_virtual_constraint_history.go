package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualConstraintHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualConstraintHistoryMgr open func
func AllVirtualConstraintHistoryMgr(db *gorm.DB) *_AllVirtualConstraintHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualConstraintHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualConstraintHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_constraint_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualConstraintHistoryMgr) GetTableName() string {
	return "__all_virtual_constraint_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualConstraintHistoryMgr) Reset() *_AllVirtualConstraintHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualConstraintHistoryMgr) Get() (result AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualConstraintHistoryMgr) Gets() (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualConstraintHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualConstraintHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualConstraintHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithConstraintID constraint_id获取
func (obj *_AllVirtualConstraintHistoryMgr) WithConstraintID(constraintID int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_id"] = constraintID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualConstraintHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualConstraintHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualConstraintHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualConstraintHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithConstraintName constraint_name获取
func (obj *_AllVirtualConstraintHistoryMgr) WithConstraintName(constraintName string) Option {
	return optionFunc(func(o *options) { o.query["constraint_name"] = constraintName })
}

// WithCheckExpr check_expr获取
func (obj *_AllVirtualConstraintHistoryMgr) WithCheckExpr(checkExpr string) Option {
	return optionFunc(func(o *options) { o.query["check_expr"] = checkExpr })
}

// WithConstraintType constraint_type获取
func (obj *_AllVirtualConstraintHistoryMgr) WithConstraintType(constraintType int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_type"] = constraintType })
}

// WithRelyFlag rely_flag获取
func (obj *_AllVirtualConstraintHistoryMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// WithEnableFlag enable_flag获取
func (obj *_AllVirtualConstraintHistoryMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithValidateFlag validate_flag获取
func (obj *_AllVirtualConstraintHistoryMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualConstraintHistoryMgr) GetByOption(opts ...Option) (result AllVirtualConstraintHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualConstraintHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualConstraintHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualConstraintHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualConstraintHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where(options.query)
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
func (obj *_AllVirtualConstraintHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromTableID(tableID int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromConstraintID 通过constraint_id获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromConstraintID(constraintID int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`constraint_id` = ?", constraintID).Find(&results).Error

	return
}

// GetBatchFromConstraintID 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromConstraintID(constraintIDs []int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`constraint_id` IN (?)", constraintIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromConstraintName 通过constraint_name获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromConstraintName(constraintName string) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`constraint_name` = ?", constraintName).Find(&results).Error

	return
}

// GetBatchFromConstraintName 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromConstraintName(constraintNames []string) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`constraint_name` IN (?)", constraintNames).Find(&results).Error

	return
}

// GetFromCheckExpr 通过check_expr获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromCheckExpr(checkExpr string) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`check_expr` = ?", checkExpr).Find(&results).Error

	return
}

// GetBatchFromCheckExpr 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromCheckExpr(checkExprs []string) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`check_expr` IN (?)", checkExprs).Find(&results).Error

	return
}

// GetFromConstraintType 通过constraint_type获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromConstraintType(constraintType int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`constraint_type` = ?", constraintType).Find(&results).Error

	return
}

// GetBatchFromConstraintType 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromConstraintType(constraintTypes []int64) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`constraint_type` IN (?)", constraintTypes).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromRelyFlag(relyFlag int8) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromEnableFlag(enableFlag int8) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllVirtualConstraintHistoryMgr) GetFromValidateFlag(validateFlag int8) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllVirtualConstraintHistoryMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualConstraintHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, constraintID int64, schemaVersion int64) (result AllVirtualConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `constraint_id` = ? AND `schema_version` = ?", tenantID, tableID, constraintID, schemaVersion).First(&result).Error

	return
}
