package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllCharsetMgr struct {
	*_BaseMgr
}

// AllCharsetMgr open func
func AllCharsetMgr(db *gorm.DB) *_AllCharsetMgr {
	if db == nil {
		panic(fmt.Errorf("AllCharsetMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllCharsetMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_charset"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllCharsetMgr) GetTableName() string {
	return "__all_charset"
}

// Reset 重置gorm会话
func (obj *_AllCharsetMgr) Reset() *_AllCharsetMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllCharsetMgr) Get() (result AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllCharsetMgr) Gets() (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllCharsetMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCharset charset获取
func (obj *_AllCharsetMgr) WithCharset(charset string) Option {
	return optionFunc(func(o *options) { o.query["charset"] = charset })
}

// WithDescription description获取
func (obj *_AllCharsetMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithDefaultCollation default_collation获取
func (obj *_AllCharsetMgr) WithDefaultCollation(defaultCollation string) Option {
	return optionFunc(func(o *options) { o.query["default_collation"] = defaultCollation })
}

// WithMaxLength max_length获取
func (obj *_AllCharsetMgr) WithMaxLength(maxLength int64) Option {
	return optionFunc(func(o *options) { o.query["max_length"] = maxLength })
}

// GetByOption 功能选项模式获取
func (obj *_AllCharsetMgr) GetByOption(opts ...Option) (result AllCharset, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllCharsetMgr) GetByOptions(opts ...Option) (results []*AllCharset, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllCharsetMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllCharset, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where(options.query)
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
func (obj *_AllCharsetMgr) GetFromCharset(charset string) (result AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`charset` = ?", charset).First(&result).Error

	return
}

// GetBatchFromCharset 批量查找
func (obj *_AllCharsetMgr) GetBatchFromCharset(charsets []string) (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`charset` IN (?)", charsets).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_AllCharsetMgr) GetFromDescription(description string) (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找
func (obj *_AllCharsetMgr) GetBatchFromDescription(descriptions []string) (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromDefaultCollation 通过default_collation获取内容
func (obj *_AllCharsetMgr) GetFromDefaultCollation(defaultCollation string) (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`default_collation` = ?", defaultCollation).Find(&results).Error

	return
}

// GetBatchFromDefaultCollation 批量查找
func (obj *_AllCharsetMgr) GetBatchFromDefaultCollation(defaultCollations []string) (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`default_collation` IN (?)", defaultCollations).Find(&results).Error

	return
}

// GetFromMaxLength 通过max_length获取内容
func (obj *_AllCharsetMgr) GetFromMaxLength(maxLength int64) (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`max_length` = ?", maxLength).Find(&results).Error

	return
}

// GetBatchFromMaxLength 批量查找
func (obj *_AllCharsetMgr) GetBatchFromMaxLength(maxLengths []int64) (results []*AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`max_length` IN (?)", maxLengths).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllCharsetMgr) FetchByPrimaryKey(charset string) (result AllCharset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCharset{}).Where("`charset` = ?", charset).First(&result).Error

	return
}
