package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualForeignKeyColumnMgr struct {
	*_BaseMgr
}

// AllVirtualForeignKeyColumnMgr open func
func AllVirtualForeignKeyColumnMgr(db *gorm.DB) *_AllVirtualForeignKeyColumnMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualForeignKeyColumnMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualForeignKeyColumnMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_foreign_key_column"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualForeignKeyColumnMgr) GetTableName() string {
	return "__all_virtual_foreign_key_column"
}

// Reset 重置gorm会话
func (obj *_AllVirtualForeignKeyColumnMgr) Reset() *_AllVirtualForeignKeyColumnMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualForeignKeyColumnMgr) Get() (result AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualForeignKeyColumnMgr) Gets() (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualForeignKeyColumnMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualForeignKeyColumnMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithForeignKeyID foreign_key_id获取
func (obj *_AllVirtualForeignKeyColumnMgr) WithForeignKeyID(foreignKeyID int64) Option {
	return optionFunc(func(o *options) { o.query["foreign_key_id"] = foreignKeyID })
}

// WithChildColumnID child_column_id获取
func (obj *_AllVirtualForeignKeyColumnMgr) WithChildColumnID(childColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["child_column_id"] = childColumnID })
}

// WithParentColumnID parent_column_id获取
func (obj *_AllVirtualForeignKeyColumnMgr) WithParentColumnID(parentColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_column_id"] = parentColumnID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualForeignKeyColumnMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualForeignKeyColumnMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithPosition position获取
func (obj *_AllVirtualForeignKeyColumnMgr) WithPosition(position int64) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualForeignKeyColumnMgr) GetByOption(opts ...Option) (result AllVirtualForeignKeyColumn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualForeignKeyColumnMgr) GetByOptions(opts ...Option) (results []*AllVirtualForeignKeyColumn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualForeignKeyColumnMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualForeignKeyColumn, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where(options.query)
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
func (obj *_AllVirtualForeignKeyColumnMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualForeignKeyColumnMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromForeignKeyID 通过foreign_key_id获取内容
func (obj *_AllVirtualForeignKeyColumnMgr) GetFromForeignKeyID(foreignKeyID int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`foreign_key_id` = ?", foreignKeyID).Find(&results).Error

	return
}

// GetBatchFromForeignKeyID 批量查找
func (obj *_AllVirtualForeignKeyColumnMgr) GetBatchFromForeignKeyID(foreignKeyIDs []int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`foreign_key_id` IN (?)", foreignKeyIDs).Find(&results).Error

	return
}

// GetFromChildColumnID 通过child_column_id获取内容
func (obj *_AllVirtualForeignKeyColumnMgr) GetFromChildColumnID(childColumnID int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`child_column_id` = ?", childColumnID).Find(&results).Error

	return
}

// GetBatchFromChildColumnID 批量查找
func (obj *_AllVirtualForeignKeyColumnMgr) GetBatchFromChildColumnID(childColumnIDs []int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`child_column_id` IN (?)", childColumnIDs).Find(&results).Error

	return
}

// GetFromParentColumnID 通过parent_column_id获取内容
func (obj *_AllVirtualForeignKeyColumnMgr) GetFromParentColumnID(parentColumnID int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`parent_column_id` = ?", parentColumnID).Find(&results).Error

	return
}

// GetBatchFromParentColumnID 批量查找
func (obj *_AllVirtualForeignKeyColumnMgr) GetBatchFromParentColumnID(parentColumnIDs []int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`parent_column_id` IN (?)", parentColumnIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualForeignKeyColumnMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualForeignKeyColumnMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualForeignKeyColumnMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualForeignKeyColumnMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_AllVirtualForeignKeyColumnMgr) GetFromPosition(position int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`position` = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量查找
func (obj *_AllVirtualForeignKeyColumnMgr) GetBatchFromPosition(positions []int64) (results []*AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`position` IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualForeignKeyColumnMgr) FetchByPrimaryKey(tenantID int64, foreignKeyID int64, childColumnID int64, parentColumnID int64) (result AllVirtualForeignKeyColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualForeignKeyColumn{}).Where("`tenant_id` = ? AND `foreign_key_id` = ? AND `child_column_id` = ? AND `parent_column_id` = ?", tenantID, foreignKeyID, childColumnID, parentColumnID).First(&result).Error

	return
}
