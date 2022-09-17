package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualCollationMgr struct {
	*_BaseMgr
}

// TenantVirtualCollationMgr open func
func TenantVirtualCollationMgr(db *gorm.DB) *_TenantVirtualCollationMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualCollationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualCollationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_collation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualCollationMgr) GetTableName() string {
	return "__tenant_virtual_collation"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualCollationMgr) Reset() *_TenantVirtualCollationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualCollationMgr) Get() (result TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualCollationMgr) Gets() (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualCollationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCollation collation获取
func (obj *_TenantVirtualCollationMgr) WithCollation(collation string) Option {
	return optionFunc(func(o *options) { o.query["collation"] = collation })
}

// WithCharset charset获取
func (obj *_TenantVirtualCollationMgr) WithCharset(charset string) Option {
	return optionFunc(func(o *options) { o.query["charset"] = charset })
}

// WithID id获取
func (obj *_TenantVirtualCollationMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIsDefault is_default获取
func (obj *_TenantVirtualCollationMgr) WithIsDefault(isDefault string) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithIsCompiled is_compiled获取
func (obj *_TenantVirtualCollationMgr) WithIsCompiled(isCompiled string) Option {
	return optionFunc(func(o *options) { o.query["is_compiled"] = isCompiled })
}

// WithSortlen sortlen获取
func (obj *_TenantVirtualCollationMgr) WithSortlen(sortlen int64) Option {
	return optionFunc(func(o *options) { o.query["sortlen"] = sortlen })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualCollationMgr) GetByOption(opts ...Option) (result TenantVirtualCollation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualCollationMgr) GetByOptions(opts ...Option) (results []*TenantVirtualCollation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualCollationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualCollation, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where(options.query)
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

// GetFromCollation 通过collation获取内容
func (obj *_TenantVirtualCollationMgr) GetFromCollation(collation string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`collation` = ?", collation).Find(&results).Error

	return
}

// GetBatchFromCollation 批量查找
func (obj *_TenantVirtualCollationMgr) GetBatchFromCollation(collations []string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`collation` IN (?)", collations).Find(&results).Error

	return
}

// GetFromCharset 通过charset获取内容
func (obj *_TenantVirtualCollationMgr) GetFromCharset(charset string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`charset` = ?", charset).Find(&results).Error

	return
}

// GetBatchFromCharset 批量查找
func (obj *_TenantVirtualCollationMgr) GetBatchFromCharset(charsets []string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`charset` IN (?)", charsets).Find(&results).Error

	return
}

// GetFromID 通过id获取内容
func (obj *_TenantVirtualCollationMgr) GetFromID(id int64) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TenantVirtualCollationMgr) GetBatchFromID(ids []int64) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容
func (obj *_TenantVirtualCollationMgr) GetFromIsDefault(isDefault string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找
func (obj *_TenantVirtualCollationMgr) GetBatchFromIsDefault(isDefaults []string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromIsCompiled 通过is_compiled获取内容
func (obj *_TenantVirtualCollationMgr) GetFromIsCompiled(isCompiled string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`is_compiled` = ?", isCompiled).Find(&results).Error

	return
}

// GetBatchFromIsCompiled 批量查找
func (obj *_TenantVirtualCollationMgr) GetBatchFromIsCompiled(isCompileds []string) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`is_compiled` IN (?)", isCompileds).Find(&results).Error

	return
}

// GetFromSortlen 通过sortlen获取内容
func (obj *_TenantVirtualCollationMgr) GetFromSortlen(sortlen int64) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`sortlen` = ?", sortlen).Find(&results).Error

	return
}

// GetBatchFromSortlen 批量查找
func (obj *_TenantVirtualCollationMgr) GetBatchFromSortlen(sortlens []int64) (results []*TenantVirtualCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCollation{}).Where("`sortlen` IN (?)", sortlens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
