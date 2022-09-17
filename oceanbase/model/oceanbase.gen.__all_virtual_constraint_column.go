package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualConstraintColumnMgr struct {
	*_BaseMgr
}

// AllVirtualConstraintColumnMgr open func
func AllVirtualConstraintColumnMgr(db *gorm.DB) *_AllVirtualConstraintColumnMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualConstraintColumnMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualConstraintColumnMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_constraint_column"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualConstraintColumnMgr) GetTableName() string {
	return "__all_virtual_constraint_column"
}

// Reset 重置gorm会话
func (obj *_AllVirtualConstraintColumnMgr) Reset() *_AllVirtualConstraintColumnMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualConstraintColumnMgr) Get() (result AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualConstraintColumnMgr) Gets() (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualConstraintColumnMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualConstraintColumnMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualConstraintColumnMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithConstraintID constraint_id获取
func (obj *_AllVirtualConstraintColumnMgr) WithConstraintID(constraintID int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_id"] = constraintID })
}

// WithColumnID column_id获取
func (obj *_AllVirtualConstraintColumnMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualConstraintColumnMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualConstraintColumnMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualConstraintColumnMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualConstraintColumnMgr) GetByOption(opts ...Option) (result AllVirtualConstraintColumn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualConstraintColumnMgr) GetByOptions(opts ...Option) (results []*AllVirtualConstraintColumn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualConstraintColumnMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualConstraintColumn, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where(options.query)
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
func (obj *_AllVirtualConstraintColumnMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualConstraintColumnMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualConstraintColumnMgr) GetFromTableID(tableID int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualConstraintColumnMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromConstraintID 通过constraint_id获取内容
func (obj *_AllVirtualConstraintColumnMgr) GetFromConstraintID(constraintID int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`constraint_id` = ?", constraintID).Find(&results).Error

	return
}

// GetBatchFromConstraintID 批量查找
func (obj *_AllVirtualConstraintColumnMgr) GetBatchFromConstraintID(constraintIDs []int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`constraint_id` IN (?)", constraintIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualConstraintColumnMgr) GetFromColumnID(columnID int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualConstraintColumnMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualConstraintColumnMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualConstraintColumnMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualConstraintColumnMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualConstraintColumnMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualConstraintColumnMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualConstraintColumnMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualConstraintColumnMgr) FetchByPrimaryKey(tenantID int64, tableID int64, constraintID int64, columnID int64) (result AllVirtualConstraintColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumn{}).Where("`tenant_id` = ? AND `table_id` = ? AND `constraint_id` = ? AND `column_id` = ?", tenantID, tableID, constraintID, columnID).First(&result).Error

	return
}
