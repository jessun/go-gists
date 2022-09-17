package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllConstraintMgr struct {
	*_BaseMgr
}

// AllConstraintMgr open func
func AllConstraintMgr(db *gorm.DB) *_AllConstraintMgr {
	if db == nil {
		panic(fmt.Errorf("AllConstraintMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllConstraintMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_constraint"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllConstraintMgr) GetTableName() string {
	return "__all_constraint"
}

// Reset 重置gorm会话
func (obj *_AllConstraintMgr) Reset() *_AllConstraintMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllConstraintMgr) Get() (result AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllConstraintMgr) Gets() (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllConstraintMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllConstraintMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllConstraintMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllConstraintMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllConstraintMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithConstraintID constraint_id获取
func (obj *_AllConstraintMgr) WithConstraintID(constraintID int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_id"] = constraintID })
}

// WithConstraintName constraint_name获取
func (obj *_AllConstraintMgr) WithConstraintName(constraintName string) Option {
	return optionFunc(func(o *options) { o.query["constraint_name"] = constraintName })
}

// WithCheckExpr check_expr获取
func (obj *_AllConstraintMgr) WithCheckExpr(checkExpr string) Option {
	return optionFunc(func(o *options) { o.query["check_expr"] = checkExpr })
}

// WithSchemaVersion schema_version获取
func (obj *_AllConstraintMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithConstraintType constraint_type获取
func (obj *_AllConstraintMgr) WithConstraintType(constraintType int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_type"] = constraintType })
}

// WithRelyFlag rely_flag获取
func (obj *_AllConstraintMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// WithEnableFlag enable_flag获取
func (obj *_AllConstraintMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithValidateFlag validate_flag获取
func (obj *_AllConstraintMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllConstraintMgr) GetByOption(opts ...Option) (result AllConstraint, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllConstraintMgr) GetByOptions(opts ...Option) (results []*AllConstraint, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllConstraintMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllConstraint, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where(options.query)
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
func (obj *_AllConstraintMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllConstraintMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllConstraintMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllConstraintMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllConstraintMgr) GetFromTenantID(tenantID int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllConstraintMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllConstraintMgr) GetFromTableID(tableID int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllConstraintMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromConstraintID 通过constraint_id获取内容
func (obj *_AllConstraintMgr) GetFromConstraintID(constraintID int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`constraint_id` = ?", constraintID).Find(&results).Error

	return
}

// GetBatchFromConstraintID 批量查找
func (obj *_AllConstraintMgr) GetBatchFromConstraintID(constraintIDs []int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`constraint_id` IN (?)", constraintIDs).Find(&results).Error

	return
}

// GetFromConstraintName 通过constraint_name获取内容
func (obj *_AllConstraintMgr) GetFromConstraintName(constraintName string) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`constraint_name` = ?", constraintName).Find(&results).Error

	return
}

// GetBatchFromConstraintName 批量查找
func (obj *_AllConstraintMgr) GetBatchFromConstraintName(constraintNames []string) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`constraint_name` IN (?)", constraintNames).Find(&results).Error

	return
}

// GetFromCheckExpr 通过check_expr获取内容
func (obj *_AllConstraintMgr) GetFromCheckExpr(checkExpr string) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`check_expr` = ?", checkExpr).Find(&results).Error

	return
}

// GetBatchFromCheckExpr 批量查找
func (obj *_AllConstraintMgr) GetBatchFromCheckExpr(checkExprs []string) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`check_expr` IN (?)", checkExprs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllConstraintMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllConstraintMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromConstraintType 通过constraint_type获取内容
func (obj *_AllConstraintMgr) GetFromConstraintType(constraintType int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`constraint_type` = ?", constraintType).Find(&results).Error

	return
}

// GetBatchFromConstraintType 批量查找
func (obj *_AllConstraintMgr) GetBatchFromConstraintType(constraintTypes []int64) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`constraint_type` IN (?)", constraintTypes).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllConstraintMgr) GetFromRelyFlag(relyFlag int8) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllConstraintMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllConstraintMgr) GetFromEnableFlag(enableFlag int8) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllConstraintMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllConstraintMgr) GetFromValidateFlag(validateFlag int8) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllConstraintMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllConstraintMgr) FetchByPrimaryKey(tenantID int64, tableID int64, constraintID int64) (result AllConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllConstraint{}).Where("`tenant_id` = ? AND `table_id` = ? AND `constraint_id` = ?", tenantID, tableID, constraintID).First(&result).Error

	return
}
