package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllForeignKeyColumnHistoryMgr struct {
	*_BaseMgr
}

// AllForeignKeyColumnHistoryMgr open func
func AllForeignKeyColumnHistoryMgr(db *gorm.DB) *_AllForeignKeyColumnHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllForeignKeyColumnHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllForeignKeyColumnHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_foreign_key_column_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllForeignKeyColumnHistoryMgr) GetTableName() string {
	return "__all_foreign_key_column_history"
}

// Reset 重置gorm会话
func (obj *_AllForeignKeyColumnHistoryMgr) Reset() *_AllForeignKeyColumnHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllForeignKeyColumnHistoryMgr) Get() (result AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllForeignKeyColumnHistoryMgr) Gets() (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllForeignKeyColumnHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithForeignKeyID foreign_key_id获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithForeignKeyID(foreignKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_id"] = foreignKeyID })
}

// WithChildColumnID child_column_id获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithChildColumnID(childColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["child_column_id"] = childColumnID })
}

// WithParentColumnID parent_column_id获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithParentColumnID(parentColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_column_id"] = parentColumnID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPosition position获取
func (obj *_AllForeignKeyColumnHistoryMgr) WithPosition(position int64) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
func (obj *_AllForeignKeyColumnHistoryMgr) GetByOption(opts ...Option) (result AllForeignKeyColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllForeignKeyColumnHistoryMgr) GetByOptions(opts ...Option) (results []*AllForeignKeyColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllForeignKeyColumnHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllForeignKeyColumnHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where(options.query)
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
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromForeignKeyID 通过foreign_key_id获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromForeignKeyID(foreignKeyID int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`foreign_key_id` = ?", foreignKeyID).Find(&results).Error

	return
}

// GetBatchFromForeignKeyID 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromForeignKeyID(foreignKeyIDs []int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`foreign_key_id` IN (?)", foreignKeyIDs).Find(&results).Error

	return
}

// GetFromChildColumnID 通过child_column_id获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromChildColumnID(childColumnID int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`child_column_id` = ?", childColumnID).Find(&results).Error

	return
}

// GetBatchFromChildColumnID 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromChildColumnID(childColumnIDs []int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`child_column_id` IN (?)", childColumnIDs).Find(&results).Error

	return
}

// GetFromParentColumnID 通过parent_column_id获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromParentColumnID(parentColumnID int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`parent_column_id` = ?", parentColumnID).Find(&results).Error

	return
}

// GetBatchFromParentColumnID 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromParentColumnID(parentColumnIDs []int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`parent_column_id` IN (?)", parentColumnIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_AllForeignKeyColumnHistoryMgr) GetFromPosition(position int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`position` = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量查找
func (obj *_AllForeignKeyColumnHistoryMgr) GetBatchFromPosition(positions []int64) (results []*AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`position` IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllForeignKeyColumnHistoryMgr) FetchByPrimaryKey(tenantID int64, foreignKeyID int64, childColumnID int64, parentColumnID int64, schemaVersion int64) (result AllForeignKeyColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllForeignKeyColumnHistory{}).Where("`tenant_id` = ? AND `foreign_key_id` = ? AND `child_column_id` = ? AND `parent_column_id` = ? AND `schema_version` = ?", tenantID, foreignKeyID, childColumnID, parentColumnID, schemaVersion).First(&result).Error

	return
}
