package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllTempTableMgr struct {
	*_BaseMgr
}

// AllTempTableMgr open func
func AllTempTableMgr(db *gorm.DB) *_AllTempTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllTempTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTempTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_temp_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTempTableMgr) GetTableName() string {
	return "__all_temp_table"
}

// Reset 重置gorm会话
func (obj *_AllTempTableMgr) Reset() *_AllTempTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTempTableMgr) Get() (result AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTempTableMgr) Gets() (results []*AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTempTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllTempTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllTempTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithCreateHost create_host获取
func (obj *_AllTempTableMgr) WithCreateHost(createHost string) Option {
	return optionFunc(func(o *options) { o.query["create_host"] = createHost })
}

// GetByOption 功能选项模式获取
func (obj *_AllTempTableMgr) GetByOption(opts ...Option) (result AllTempTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTempTableMgr) GetByOptions(opts ...Option) (results []*AllTempTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTempTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTempTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where(options.query)
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
func (obj *_AllTempTableMgr) GetFromTenantID(tenantID int64) (results []*AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTempTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTempTableMgr) GetFromTableID(tableID int64) (results []*AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTempTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromCreateHost 通过create_host获取内容
func (obj *_AllTempTableMgr) GetFromCreateHost(createHost string) (results []*AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where("`create_host` = ?", createHost).Find(&results).Error

	return
}

// GetBatchFromCreateHost 批量查找
func (obj *_AllTempTableMgr) GetBatchFromCreateHost(createHosts []string) (results []*AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where("`create_host` IN (?)", createHosts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTempTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64) (result AllTempTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTempTable{}).Where("`tenant_id` = ? AND `table_id` = ?", tenantID, tableID).First(&result).Error

	return
}
