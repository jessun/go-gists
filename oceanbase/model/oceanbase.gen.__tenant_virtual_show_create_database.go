package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualShowCreateDatabaseMgr struct {
	*_BaseMgr
}

// TenantVirtualShowCreateDatabaseMgr open func
func TenantVirtualShowCreateDatabaseMgr(db *gorm.DB) *_TenantVirtualShowCreateDatabaseMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualShowCreateDatabaseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualShowCreateDatabaseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_show_create_database"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetTableName() string {
	return "__tenant_virtual_show_create_database"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualShowCreateDatabaseMgr) Reset() *_TenantVirtualShowCreateDatabaseMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualShowCreateDatabaseMgr) Get() (result TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualShowCreateDatabaseMgr) Gets() (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualShowCreateDatabaseMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDatabaseID database_id获取
func (obj *_TenantVirtualShowCreateDatabaseMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithDatabaseName database_name获取
func (obj *_TenantVirtualShowCreateDatabaseMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithCreateDatabase create_database获取
func (obj *_TenantVirtualShowCreateDatabaseMgr) WithCreateDatabase(createDatabase string) Option {
	return optionFunc(func(o *options) { o.query["create_database"] = createDatabase })
}

// WithCreateDatabaseWithIfNotExists create_database_with_if_not_exists获取
func (obj *_TenantVirtualShowCreateDatabaseMgr) WithCreateDatabaseWithIfNotExists(createDatabaseWithIfNotExists string) Option {
	return optionFunc(func(o *options) { o.query["create_database_with_if_not_exists"] = createDatabaseWithIfNotExists })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetByOption(opts ...Option) (result TenantVirtualShowCreateDatabase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetByOptions(opts ...Option) (results []*TenantVirtualShowCreateDatabase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualShowCreateDatabaseMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualShowCreateDatabase, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where(options.query)
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
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetFromDatabaseID(databaseID int64) (result TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`database_id` = ?", databaseID).First(&result).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetFromDatabaseName(databaseName string) (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromCreateDatabase 通过create_database获取内容
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetFromCreateDatabase(createDatabase string) (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`create_database` = ?", createDatabase).Find(&results).Error

	return
}

// GetBatchFromCreateDatabase 批量查找
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetBatchFromCreateDatabase(createDatabases []string) (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`create_database` IN (?)", createDatabases).Find(&results).Error

	return
}

// GetFromCreateDatabaseWithIfNotExists 通过create_database_with_if_not_exists获取内容
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetFromCreateDatabaseWithIfNotExists(createDatabaseWithIfNotExists string) (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`create_database_with_if_not_exists` = ?", createDatabaseWithIfNotExists).Find(&results).Error

	return
}

// GetBatchFromCreateDatabaseWithIfNotExists 批量查找
func (obj *_TenantVirtualShowCreateDatabaseMgr) GetBatchFromCreateDatabaseWithIfNotExists(createDatabaseWithIfNotExistss []string) (results []*TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`create_database_with_if_not_exists` IN (?)", createDatabaseWithIfNotExistss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualShowCreateDatabaseMgr) FetchByPrimaryKey(databaseID int64) (result TenantVirtualShowCreateDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateDatabase{}).Where("`database_id` = ?", databaseID).First(&result).Error

	return
}
