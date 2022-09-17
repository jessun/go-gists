package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualTableColumnMgr struct {
	*_BaseMgr
}

// TenantVirtualTableColumnMgr open func
func TenantVirtualTableColumnMgr(db *gorm.DB) *_TenantVirtualTableColumnMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualTableColumnMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualTableColumnMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_table_column"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualTableColumnMgr) GetTableName() string {
	return "__tenant_virtual_table_column"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualTableColumnMgr) Reset() *_TenantVirtualTableColumnMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualTableColumnMgr) Get() (result TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualTableColumnMgr) Gets() (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualTableColumnMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_TenantVirtualTableColumnMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithField field获取
func (obj *_TenantVirtualTableColumnMgr) WithField(field string) Option {
	return optionFunc(func(o *options) { o.query["field"] = field })
}

// WithType type获取
func (obj *_TenantVirtualTableColumnMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCollation collation获取
func (obj *_TenantVirtualTableColumnMgr) WithCollation(collation string) Option {
	return optionFunc(func(o *options) { o.query["collation"] = collation })
}

// WithNull null获取
func (obj *_TenantVirtualTableColumnMgr) WithNull(null string) Option {
	return optionFunc(func(o *options) { o.query["null"] = null })
}

// WithKey key获取
func (obj *_TenantVirtualTableColumnMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithDefault default获取
func (obj *_TenantVirtualTableColumnMgr) WithDefault(_default string) Option {
	return optionFunc(func(o *options) { o.query["default"] = _default })
}

// WithExtra extra获取
func (obj *_TenantVirtualTableColumnMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithPrivileges privileges获取
func (obj *_TenantVirtualTableColumnMgr) WithPrivileges(privileges string) Option {
	return optionFunc(func(o *options) { o.query["privileges"] = privileges })
}

// WithComment comment获取
func (obj *_TenantVirtualTableColumnMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualTableColumnMgr) GetByOption(opts ...Option) (result TenantVirtualTableColumn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualTableColumnMgr) GetByOptions(opts ...Option) (results []*TenantVirtualTableColumn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualTableColumnMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualTableColumn, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where(options.query)
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
func (obj *_TenantVirtualTableColumnMgr) GetFromTableID(tableID int64) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromTableID(tableIDs []int64) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromField 通过field获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromField(field string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`field` = ?", field).Find(&results).Error

	return
}

// GetBatchFromField 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromField(fields []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`field` IN (?)", fields).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromType(_type string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromType(_types []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromCollation 通过collation获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromCollation(collation string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`collation` = ?", collation).Find(&results).Error

	return
}

// GetBatchFromCollation 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromCollation(collations []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`collation` IN (?)", collations).Find(&results).Error

	return
}

// GetFromNull 通过null获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromNull(null string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`null` = ?", null).Find(&results).Error

	return
}

// GetBatchFromNull 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromNull(nulls []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`null` IN (?)", nulls).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromKey(key string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`key` = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromKey(keys []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

// GetFromDefault 通过default获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromDefault(_default string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`default` = ?", _default).Find(&results).Error

	return
}

// GetBatchFromDefault 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromDefault(_defaults []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`default` IN (?)", _defaults).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromExtra(extra string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromExtra(extras []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromPrivileges 通过privileges获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromPrivileges(privileges string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`privileges` = ?", privileges).Find(&results).Error

	return
}

// GetBatchFromPrivileges 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromPrivileges(privilegess []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`privileges` IN (?)", privilegess).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_TenantVirtualTableColumnMgr) GetFromComment(comment string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_TenantVirtualTableColumnMgr) GetBatchFromComment(comments []string) (results []*TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualTableColumnMgr) FetchByPrimaryKey(tableID int64, field string) (result TenantVirtualTableColumn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableColumn{}).Where("`table_id` = ? AND `field` = ?", tableID, field).First(&result).Error

	return
}
