package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualShowCreateTableMgr struct {
	*_BaseMgr
}

// TenantVirtualShowCreateTableMgr open func
func TenantVirtualShowCreateTableMgr(db *gorm.DB) *_TenantVirtualShowCreateTableMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualShowCreateTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualShowCreateTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_show_create_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualShowCreateTableMgr) GetTableName() string {
	return "__tenant_virtual_show_create_table"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualShowCreateTableMgr) Reset() *_TenantVirtualShowCreateTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualShowCreateTableMgr) Get() (result TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualShowCreateTableMgr) Gets() (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualShowCreateTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_TenantVirtualShowCreateTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTableName table_name获取
func (obj *_TenantVirtualShowCreateTableMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithCreateTable create_table获取
func (obj *_TenantVirtualShowCreateTableMgr) WithCreateTable(createTable string) Option {
	return optionFunc(func(o *options) { o.query["create_table"] = createTable })
}

// WithCharacterSetClient character_set_client获取
func (obj *_TenantVirtualShowCreateTableMgr) WithCharacterSetClient(characterSetClient string) Option {
	return optionFunc(func(o *options) { o.query["character_set_client"] = characterSetClient })
}

// WithCollationConnection collation_connection获取
func (obj *_TenantVirtualShowCreateTableMgr) WithCollationConnection(collationConnection string) Option {
	return optionFunc(func(o *options) { o.query["collation_connection"] = collationConnection })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualShowCreateTableMgr) GetByOption(opts ...Option) (result TenantVirtualShowCreateTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualShowCreateTableMgr) GetByOptions(opts ...Option) (results []*TenantVirtualShowCreateTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualShowCreateTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualShowCreateTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where(options.query)
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

// GetFromTableID 通过table_id获取内容
func (obj *_TenantVirtualShowCreateTableMgr) GetFromTableID(tableID int64) (result TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`table_id` = ?", tableID).First(&result).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_TenantVirtualShowCreateTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_TenantVirtualShowCreateTableMgr) GetFromTableName(tableName string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_TenantVirtualShowCreateTableMgr) GetBatchFromTableName(tableNames []string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromCreateTable 通过create_table获取内容
func (obj *_TenantVirtualShowCreateTableMgr) GetFromCreateTable(createTable string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`create_table` = ?", createTable).Find(&results).Error

	return
}

// GetBatchFromCreateTable 批量查找
func (obj *_TenantVirtualShowCreateTableMgr) GetBatchFromCreateTable(createTables []string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`create_table` IN (?)", createTables).Find(&results).Error

	return
}

// GetFromCharacterSetClient 通过character_set_client获取内容
func (obj *_TenantVirtualShowCreateTableMgr) GetFromCharacterSetClient(characterSetClient string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`character_set_client` = ?", characterSetClient).Find(&results).Error

	return
}

// GetBatchFromCharacterSetClient 批量查找
func (obj *_TenantVirtualShowCreateTableMgr) GetBatchFromCharacterSetClient(characterSetClients []string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`character_set_client` IN (?)", characterSetClients).Find(&results).Error

	return
}

// GetFromCollationConnection 通过collation_connection获取内容
func (obj *_TenantVirtualShowCreateTableMgr) GetFromCollationConnection(collationConnection string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`collation_connection` = ?", collationConnection).Find(&results).Error

	return
}

// GetBatchFromCollationConnection 批量查找
func (obj *_TenantVirtualShowCreateTableMgr) GetBatchFromCollationConnection(collationConnections []string) (results []*TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`collation_connection` IN (?)", collationConnections).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualShowCreateTableMgr) FetchByPrimaryKey(tableID int64) (result TenantVirtualShowCreateTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTable{}).Where("`table_id` = ?", tableID).First(&result).Error

	return
}
