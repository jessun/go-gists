package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualCharsetMgr struct {
	*_BaseMgr
}

// TenantVirtualCharsetMgr open func
func TenantVirtualCharsetMgr(db *gorm.DB) *_TenantVirtualCharsetMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualCharsetMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualCharsetMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_charset"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualCharsetMgr) GetTableName() string {
	return "__tenant_virtual_charset"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualCharsetMgr) Reset() *_TenantVirtualCharsetMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualCharsetMgr) Get() (result TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualCharsetMgr) Gets() (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualCharsetMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCharset charset获取
func (obj *_TenantVirtualCharsetMgr) WithCharset(charset string) Option {
	return optionFunc(func(o *options) { o.query["charset"] = charset })
}

// WithDescription description获取
func (obj *_TenantVirtualCharsetMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithDefaultCollation default_collation获取
func (obj *_TenantVirtualCharsetMgr) WithDefaultCollation(defaultCollation string) Option {
	return optionFunc(func(o *options) { o.query["default_collation"] = defaultCollation })
}

// WithMaxLength max_length获取
func (obj *_TenantVirtualCharsetMgr) WithMaxLength(maxLength int64) Option {
	return optionFunc(func(o *options) { o.query["max_length"] = maxLength })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualCharsetMgr) GetByOption(opts ...Option) (result TenantVirtualCharset, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualCharsetMgr) GetByOptions(opts ...Option) (results []*TenantVirtualCharset, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualCharsetMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualCharset, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where(options.query)
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

// GetFromCharset 通过charset获取内容
func (obj *_TenantVirtualCharsetMgr) GetFromCharset(charset string) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`charset` = ?", charset).Find(&results).Error

	return
}

// GetBatchFromCharset 批量查找
func (obj *_TenantVirtualCharsetMgr) GetBatchFromCharset(charsets []string) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`charset` IN (?)", charsets).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_TenantVirtualCharsetMgr) GetFromDescription(description string) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找
func (obj *_TenantVirtualCharsetMgr) GetBatchFromDescription(descriptions []string) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromDefaultCollation 通过default_collation获取内容
func (obj *_TenantVirtualCharsetMgr) GetFromDefaultCollation(defaultCollation string) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`default_collation` = ?", defaultCollation).Find(&results).Error

	return
}

// GetBatchFromDefaultCollation 批量查找
func (obj *_TenantVirtualCharsetMgr) GetBatchFromDefaultCollation(defaultCollations []string) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`default_collation` IN (?)", defaultCollations).Find(&results).Error

	return
}

// GetFromMaxLength 通过max_length获取内容
func (obj *_TenantVirtualCharsetMgr) GetFromMaxLength(maxLength int64) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`max_length` = ?", maxLength).Find(&results).Error

	return
}

// GetBatchFromMaxLength 批量查找
func (obj *_TenantVirtualCharsetMgr) GetBatchFromMaxLength(maxLengths []int64) (results []*TenantVirtualCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualCharset{}).Where("`max_length` IN (?)", maxLengths).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
