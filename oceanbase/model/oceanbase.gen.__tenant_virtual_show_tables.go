package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualShowTablesMgr struct {
	*_BaseMgr
}

// TenantVirtualShowTablesMgr open func
func TenantVirtualShowTablesMgr(db *gorm.DB) *_TenantVirtualShowTablesMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualShowTablesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualShowTablesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_show_tables"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualShowTablesMgr) GetTableName() string {
	return "__tenant_virtual_show_tables"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualShowTablesMgr) Reset() *_TenantVirtualShowTablesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualShowTablesMgr) Get() (result TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualShowTablesMgr) Gets() (results []*TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualShowTablesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDatabaseID database_id获取
func (obj *_TenantVirtualShowTablesMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithTableName table_name获取
func (obj *_TenantVirtualShowTablesMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithTableType table_type获取
func (obj *_TenantVirtualShowTablesMgr) WithTableType(tableType string) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualShowTablesMgr) GetByOption(opts ...Option) (result TenantVirtualShowTables, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualShowTablesMgr) GetByOptions(opts ...Option) (results []*TenantVirtualShowTables, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualShowTablesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualShowTables, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where(options.query)
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

// GetFromDatabaseID 通过database_id获取内容
func (obj *_TenantVirtualShowTablesMgr) GetFromDatabaseID(databaseID int64) (results []*TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_TenantVirtualShowTablesMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_TenantVirtualShowTablesMgr) GetFromTableName(tableName string) (results []*TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_TenantVirtualShowTablesMgr) GetBatchFromTableName(tableNames []string) (results []*TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_TenantVirtualShowTablesMgr) GetFromTableType(tableType string) (results []*TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_TenantVirtualShowTablesMgr) GetBatchFromTableType(tableTypes []string) (results []*TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualShowTablesMgr) FetchByPrimaryKey(databaseID int64, tableName string) (result TenantVirtualShowTables, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowTables{}).Where("`database_id` = ? AND `table_name` = ?", databaseID, tableName).First(&result).Error

	return
}
