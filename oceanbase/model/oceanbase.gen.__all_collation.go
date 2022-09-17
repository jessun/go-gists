package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllCollationMgr struct {
	*_BaseMgr
}

// AllCollationMgr open func
func AllCollationMgr(db *gorm.DB) *_AllCollationMgr {
	if db == nil {
		panic(fmt.Errorf("AllCollationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllCollationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_collation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllCollationMgr) GetTableName() string {
	return "__all_collation"
}

// Reset 重置gorm会话
func (obj *_AllCollationMgr) Reset() *_AllCollationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllCollationMgr) Get() (result AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllCollationMgr) Gets() (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllCollationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCollation collation获取
func (obj *_AllCollationMgr) WithCollation(collation string) Option {
	return optionFunc(func(o *options) { o.query["collation"] = collation })
}

// WithCharset charset获取
func (obj *_AllCollationMgr) WithCharset(charset string) Option {
	return optionFunc(func(o *options) { o.query["charset"] = charset })
}

// WithID id获取
func (obj *_AllCollationMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIsDefault is_default获取
func (obj *_AllCollationMgr) WithIsDefault(isDefault string) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithIsCompiled is_compiled获取
func (obj *_AllCollationMgr) WithIsCompiled(isCompiled string) Option {
	return optionFunc(func(o *options) { o.query["is_compiled"] = isCompiled })
}

// WithSortlen sortlen获取
func (obj *_AllCollationMgr) WithSortlen(sortlen int64) Option {
	return optionFunc(func(o *options) { o.query["sortlen"] = sortlen })
}

// GetByOption 功能选项模式获取
func (obj *_AllCollationMgr) GetByOption(opts ...Option) (result AllCollation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllCollationMgr) GetByOptions(opts ...Option) (results []*AllCollation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllCollationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllCollation, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where(options.query)
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
func (obj *_AllCollationMgr) GetFromCollation(collation string) (result AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`collation` = ?", collation).First(&result).Error

	return
}

// GetBatchFromCollation 批量查找
func (obj *_AllCollationMgr) GetBatchFromCollation(collations []string) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`collation` IN (?)", collations).Find(&results).Error

	return
}

// GetFromCharset 通过charset获取内容
func (obj *_AllCollationMgr) GetFromCharset(charset string) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`charset` = ?", charset).Find(&results).Error

	return
}

// GetBatchFromCharset 批量查找
func (obj *_AllCollationMgr) GetBatchFromCharset(charsets []string) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`charset` IN (?)", charsets).Find(&results).Error

	return
}

// GetFromID 通过id获取内容
func (obj *_AllCollationMgr) GetFromID(id int64) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AllCollationMgr) GetBatchFromID(ids []int64) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容
func (obj *_AllCollationMgr) GetFromIsDefault(isDefault string) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找
func (obj *_AllCollationMgr) GetBatchFromIsDefault(isDefaults []string) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromIsCompiled 通过is_compiled获取内容
func (obj *_AllCollationMgr) GetFromIsCompiled(isCompiled string) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`is_compiled` = ?", isCompiled).Find(&results).Error

	return
}

// GetBatchFromIsCompiled 批量查找
func (obj *_AllCollationMgr) GetBatchFromIsCompiled(isCompileds []string) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`is_compiled` IN (?)", isCompileds).Find(&results).Error

	return
}

// GetFromSortlen 通过sortlen获取内容
func (obj *_AllCollationMgr) GetFromSortlen(sortlen int64) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`sortlen` = ?", sortlen).Find(&results).Error

	return
}

// GetBatchFromSortlen 批量查找
func (obj *_AllCollationMgr) GetBatchFromSortlen(sortlens []int64) (results []*AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`sortlen` IN (?)", sortlens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllCollationMgr) FetchByPrimaryKey(collation string) (result AllCollation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollation{}).Where("`collation` = ?", collation).First(&result).Error

	return
}
