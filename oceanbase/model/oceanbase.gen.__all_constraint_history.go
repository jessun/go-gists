package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllConstraintHistoryMgr struct {
	*_BaseMgr
}

// AllConstraintHistoryMgr open func
func AllConstraintHistoryMgr(db *gorm.DB) *_AllConstraintHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllConstraintHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllConstraintHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_constraint_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllConstraintHistoryMgr) GetTableName() string {
	return "__all_constraint_history"
}

// Reset 重置gorm会话
func (obj *_AllConstraintHistoryMgr) Reset() *_AllConstraintHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllConstraintHistoryMgr) Get() (result AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllConstraintHistoryMgr) Gets() (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllConstraintHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllConstraintHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllConstraintHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllConstraintHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllConstraintHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithConstraintID constraint_id获取
func (obj *_AllConstraintHistoryMgr) WithConstraintID(constraintID int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_id"] = constraintID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllConstraintHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllConstraintHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithConstraintName constraint_name获取
func (obj *_AllConstraintHistoryMgr) WithConstraintName(constraintName string) Option {
	return optionFunc(func(o *options) { o.query["constraint_name"] = constraintName })
}

// WithCheckExpr check_expr获取
func (obj *_AllConstraintHistoryMgr) WithCheckExpr(checkExpr string) Option {
	return optionFunc(func(o *options) { o.query["check_expr"] = checkExpr })
}

// WithConstraintType constraint_type获取
func (obj *_AllConstraintHistoryMgr) WithConstraintType(constraintType int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_type"] = constraintType })
}

// WithRelyFlag rely_flag获取
func (obj *_AllConstraintHistoryMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// WithEnableFlag enable_flag获取
func (obj *_AllConstraintHistoryMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithValidateFlag validate_flag获取
func (obj *_AllConstraintHistoryMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllConstraintHistoryMgr) GetByOption(opts ...Option) (result AllConstraintHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllConstraintHistoryMgr) GetByOptions(opts ...Option) (results []*AllConstraintHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllConstraintHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllConstraintHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where(options.query)
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
func (obj *_AllConstraintHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllConstraintHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllConstraintHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllConstraintHistoryMgr) GetFromTableID(tableID int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromConstraintID 通过constraint_id获取内容
func (obj *_AllConstraintHistoryMgr) GetFromConstraintID(constraintID int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`constraint_id` = ?", constraintID).Find(&results).Error

	return
}

// GetBatchFromConstraintID 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromConstraintID(constraintIDs []int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`constraint_id` IN (?)", constraintIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllConstraintHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllConstraintHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromConstraintName 通过constraint_name获取内容
func (obj *_AllConstraintHistoryMgr) GetFromConstraintName(constraintName string) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`constraint_name` = ?", constraintName).Find(&results).Error

	return
}

// GetBatchFromConstraintName 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromConstraintName(constraintNames []string) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`constraint_name` IN (?)", constraintNames).Find(&results).Error

	return
}

// GetFromCheckExpr 通过check_expr获取内容
func (obj *_AllConstraintHistoryMgr) GetFromCheckExpr(checkExpr string) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`check_expr` = ?", checkExpr).Find(&results).Error

	return
}

// GetBatchFromCheckExpr 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromCheckExpr(checkExprs []string) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`check_expr` IN (?)", checkExprs).Find(&results).Error

	return
}

// GetFromConstraintType 通过constraint_type获取内容
func (obj *_AllConstraintHistoryMgr) GetFromConstraintType(constraintType int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`constraint_type` = ?", constraintType).Find(&results).Error

	return
}

// GetBatchFromConstraintType 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromConstraintType(constraintTypes []int64) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`constraint_type` IN (?)", constraintTypes).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllConstraintHistoryMgr) GetFromRelyFlag(relyFlag int8) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllConstraintHistoryMgr) GetFromEnableFlag(enableFlag int8) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllConstraintHistoryMgr) GetFromValidateFlag(validateFlag int8) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllConstraintHistoryMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllConstraintHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, constraintID int64, schemaVersion int64) (result AllConstraintHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraintHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `constraint_id` = ? AND `schema_version` = ?", tenantID, tableID, constraintID, schemaVersion).First(&result).Error

	return
}
