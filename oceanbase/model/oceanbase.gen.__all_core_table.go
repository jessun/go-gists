package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllCoreTableMgr struct {
	*_BaseMgr
}

// AllCoreTableMgr open func
func AllCoreTableMgr(db *gorm.DB) *_AllCoreTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllCoreTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllCoreTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_core_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllCoreTableMgr) GetTableName() string {
	return "__all_core_table"
}

// Reset 重置gorm会话
func (obj *_AllCoreTableMgr) Reset() *_AllCoreTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllCoreTableMgr) Get() (result AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllCoreTableMgr) Gets() (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllCoreTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllCoreTableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllCoreTableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTableName table_name获取
func (obj *_AllCoreTableMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithRowID row_id获取
func (obj *_AllCoreTableMgr) WithRowID(rowID int64) Option {
	return optionFunc(func(o *options) { o.query["row_id"] = rowID })
}

// WithColumnName column_name获取
func (obj *_AllCoreTableMgr) WithColumnName(columnName string) Option {
	return optionFunc(func(o *options) { o.query["column_name"] = columnName })
}

// WithColumnValue column_value获取
func (obj *_AllCoreTableMgr) WithColumnValue(columnValue string) Option {
	return optionFunc(func(o *options) { o.query["column_value"] = columnValue })
}

// GetByOption 功能选项模式获取
func (obj *_AllCoreTableMgr) GetByOption(opts ...Option) (result AllCoreTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllCoreTableMgr) GetByOptions(opts ...Option) (results []*AllCoreTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllCoreTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllCoreTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where(options.query)
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
func (obj *_AllCoreTableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllCoreTableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllCoreTableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllCoreTableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllCoreTableMgr) GetFromTableName(tableName string) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllCoreTableMgr) GetBatchFromTableName(tableNames []string) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromRowID 通过row_id获取内容
func (obj *_AllCoreTableMgr) GetFromRowID(rowID int64) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`row_id` = ?", rowID).Find(&results).Error

	return
}

// GetBatchFromRowID 批量查找
func (obj *_AllCoreTableMgr) GetBatchFromRowID(rowIDs []int64) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`row_id` IN (?)", rowIDs).Find(&results).Error

	return
}

// GetFromColumnName 通过column_name获取内容
func (obj *_AllCoreTableMgr) GetFromColumnName(columnName string) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`column_name` = ?", columnName).Find(&results).Error

	return
}

// GetBatchFromColumnName 批量查找
func (obj *_AllCoreTableMgr) GetBatchFromColumnName(columnNames []string) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`column_name` IN (?)", columnNames).Find(&results).Error

	return
}

// GetFromColumnValue 通过column_value获取内容
func (obj *_AllCoreTableMgr) GetFromColumnValue(columnValue string) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`column_value` = ?", columnValue).Find(&results).Error

	return
}

// GetBatchFromColumnValue 批量查找
func (obj *_AllCoreTableMgr) GetBatchFromColumnValue(columnValues []string) (results []*AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`column_value` IN (?)", columnValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllCoreTableMgr) FetchByPrimaryKey(tableName string, rowID int64, columnName string) (result AllCoreTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCoreTable{}).Where("`table_name` = ? AND `row_id` = ? AND `column_name` = ?", tableName, rowID, columnName).First(&result).Error

	return
}
