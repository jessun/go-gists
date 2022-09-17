package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualConstraintColumnHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualConstraintColumnHistoryMgr open func
func AllVirtualConstraintColumnHistoryMgr(db *gorm.DB) *_AllVirtualConstraintColumnHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualConstraintColumnHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualConstraintColumnHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_constraint_column_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetTableName() string {
	return "__all_virtual_constraint_column_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualConstraintColumnHistoryMgr) Reset() *_AllVirtualConstraintColumnHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) Get() (result AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualConstraintColumnHistoryMgr) Gets() (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualConstraintColumnHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithConstraintID constraint_id获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithConstraintID(constraintID int64) Option {
	return optionFunc(func(o *options) { o.query["constraint_id"] = constraintID })
}

// WithColumnID column_id获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetByOption(opts ...Option) (result AllVirtualConstraintColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualConstraintColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualConstraintColumnHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualConstraintColumnHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where(options.query)
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
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromTableID(tableID int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromConstraintID 通过constraint_id获取内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromConstraintID(constraintID int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`constraint_id` = ?", constraintID).Find(&results).Error

	return
}

// GetBatchFromConstraintID 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromConstraintID(constraintIDs []int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`constraint_id` IN (?)", constraintIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromColumnID(columnID int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualConstraintColumnHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualConstraintColumnHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, constraintID int64, columnID int64, schemaVersion int64) (result AllVirtualConstraintColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualConstraintColumnHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `constraint_id` = ? AND `column_id` = ? AND `schema_version` = ?", tenantID, tableID, constraintID, columnID, schemaVersion).First(&result).Error

	return
}
