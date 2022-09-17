package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualConstraintMgr struct {
	*_BaseMgr
}

// AllVirtualConstraintMgr open func
func AllVirtualConstraintMgr(db *gorm.DB) *_AllVirtualConstraintMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualConstraintMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualConstraintMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_constraint"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualConstraintMgr) GetTableName() string {
	return "__all_virtual_constraint"
}

// Reset 重置gorm会话
func (obj *_AllVirtualConstraintMgr) Reset() *_AllVirtualConstraintMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualConstraintMgr) Get() (result AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualConstraintMgr) Gets() (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualConstraintMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualConstraintMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualConstraintMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithConstraintID constraint_id获取
func (obj *_AllVirtualConstraintMgr) WithConstraintID(constraintID int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_id"] = constraintID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualConstraintMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualConstraintMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithConstraintName constraint_name获取
func (obj *_AllVirtualConstraintMgr) WithConstraintName(constraintName string) Option {
	return optionFunc(func(o *options) { o.query["constraint_name"] = constraintName })
}

// WithCheckExpr check_expr获取
func (obj *_AllVirtualConstraintMgr) WithCheckExpr(checkExpr string) Option {
	return optionFunc(func(o *options) { o.query["check_expr"] = checkExpr })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualConstraintMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithConstraintType constraint_type获取
func (obj *_AllVirtualConstraintMgr) WithConstraintType(constraintType int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_type"] = constraintType })
}

// WithRelyFlag rely_flag获取
func (obj *_AllVirtualConstraintMgr) WithRelyFlag(relyFlag int8) Option {
	return optionFunc(func(o *options) { o.query["rely_flag"] = relyFlag })
}

// WithEnableFlag enable_flag获取
func (obj *_AllVirtualConstraintMgr) WithEnableFlag(enableFlag int8) Option {
	return optionFunc(func(o *options) { o.query["enable_flag"] = enableFlag })
}

// WithValidateFlag validate_flag获取
func (obj *_AllVirtualConstraintMgr) WithValidateFlag(validateFlag int8) Option {
	return optionFunc(func(o *options) { o.query["validate_flag"] = validateFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualConstraintMgr) GetByOption(opts ...Option) (result AllVirtualConstraint, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualConstraintMgr) GetByOptions(opts ...Option) (results []*AllVirtualConstraint, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualConstraintMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualConstraint, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where(options.query)
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
func (obj *_AllVirtualConstraintMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualConstraintMgr) GetFromTableID(tableID int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromConstraintID 通过constraint_id获取内容
func (obj *_AllVirtualConstraintMgr) GetFromConstraintID(constraintID int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`constraint_id` = ?", constraintID).Find(&results).Error

	return
}

// GetBatchFromConstraintID 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromConstraintID(constraintIDs []int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`constraint_id` IN (?)", constraintIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualConstraintMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualConstraintMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromConstraintName 通过constraint_name获取内容
func (obj *_AllVirtualConstraintMgr) GetFromConstraintName(constraintName string) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`constraint_name` = ?", constraintName).Find(&results).Error

	return
}

// GetBatchFromConstraintName 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromConstraintName(constraintNames []string) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`constraint_name` IN (?)", constraintNames).Find(&results).Error

	return
}

// GetFromCheckExpr 通过check_expr获取内容
func (obj *_AllVirtualConstraintMgr) GetFromCheckExpr(checkExpr string) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`check_expr` = ?", checkExpr).Find(&results).Error

	return
}

// GetBatchFromCheckExpr 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromCheckExpr(checkExprs []string) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`check_expr` IN (?)", checkExprs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualConstraintMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromConstraintType 通过constraint_type获取内容
func (obj *_AllVirtualConstraintMgr) GetFromConstraintType(constraintType int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`constraint_type` = ?", constraintType).Find(&results).Error

	return
}

// GetBatchFromConstraintType 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromConstraintType(constraintTypes []int64) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`constraint_type` IN (?)", constraintTypes).Find(&results).Error

	return
}

// GetFromRelyFlag 通过rely_flag获取内容
func (obj *_AllVirtualConstraintMgr) GetFromRelyFlag(relyFlag int8) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`rely_flag` = ?", relyFlag).Find(&results).Error

	return
}

// GetBatchFromRelyFlag 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromRelyFlag(relyFlags []int8) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`rely_flag` IN (?)", relyFlags).Find(&results).Error

	return
}

// GetFromEnableFlag 通过enable_flag获取内容
func (obj *_AllVirtualConstraintMgr) GetFromEnableFlag(enableFlag int8) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`enable_flag` = ?", enableFlag).Find(&results).Error

	return
}

// GetBatchFromEnableFlag 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromEnableFlag(enableFlags []int8) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`enable_flag` IN (?)", enableFlags).Find(&results).Error

	return
}

// GetFromValidateFlag 通过validate_flag获取内容
func (obj *_AllVirtualConstraintMgr) GetFromValidateFlag(validateFlag int8) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`validate_flag` = ?", validateFlag).Find(&results).Error

	return
}

// GetBatchFromValidateFlag 批量查找
func (obj *_AllVirtualConstraintMgr) GetBatchFromValidateFlag(validateFlags []int8) (results []*AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`validate_flag` IN (?)", validateFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualConstraintMgr) FetchByPrimaryKey(tenantID int64, tableID int64, constraintID int64) (result AllVirtualConstraint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraint{}).Where("`tenant_id` = ? AND `table_id` = ? AND `constraint_id` = ?", tenantID, tableID, constraintID).First(&result).Error

	return
}
