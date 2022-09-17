package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualRecyclebinMgr struct {
	*_BaseMgr
}

// AllVirtualRecyclebinMgr open func
func AllVirtualRecyclebinMgr(db *gorm.DB) *_AllVirtualRecyclebinMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRecyclebinMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRecyclebinMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_recyclebin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRecyclebinMgr) GetTableName() string {
	return "__all_virtual_recyclebin"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRecyclebinMgr) Reset() *_AllVirtualRecyclebinMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRecyclebinMgr) Get() (result AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRecyclebinMgr) Gets() (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRecyclebinMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRecyclebinMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjectName object_name获取
func (obj *_AllVirtualRecyclebinMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithType type获取
func (obj *_AllVirtualRecyclebinMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualRecyclebinMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualRecyclebinMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithTableID table_id获取
func (obj *_AllVirtualRecyclebinMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllVirtualRecyclebinMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithOriginalName original_name获取
func (obj *_AllVirtualRecyclebinMgr) WithOriginalName(originalName string) Option {
	return optionFunc(func(o *options) { o.query["original_name"] = originalName })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRecyclebinMgr) GetByOption(opts ...Option) (result AllVirtualRecyclebin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRecyclebinMgr) GetByOptions(opts ...Option) (results []*AllVirtualRecyclebin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRecyclebinMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRecyclebin, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where(options.query)
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
func (obj *_AllVirtualRecyclebinMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_AllVirtualRecyclebinMgr) GetFromObjectName(objectName string) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromObjectName(objectNames []string) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualRecyclebinMgr) GetFromType(_type int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromType(_types []int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualRecyclebinMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualRecyclebinMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualRecyclebinMgr) GetFromTableID(tableID int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllVirtualRecyclebinMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromOriginalName 通过original_name获取内容
func (obj *_AllVirtualRecyclebinMgr) GetFromOriginalName(originalName string) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`original_name` = ?", originalName).Find(&results).Error

	return
}

// GetBatchFromOriginalName 批量查找
func (obj *_AllVirtualRecyclebinMgr) GetBatchFromOriginalName(originalNames []string) (results []*AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`original_name` IN (?)", originalNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualRecyclebinMgr) FetchByPrimaryKey(tenantID int64, objectName string, _type int64) (result AllVirtualRecyclebin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRecyclebin{}).Where("`tenant_id` = ? AND `object_name` = ? AND `type` = ?", tenantID, objectName, _type).First(&result).Error

	return
}
