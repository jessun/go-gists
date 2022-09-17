package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllDummyMgr struct {
	*_BaseMgr
}

// AllDummyMgr open func
func AllDummyMgr(db *gorm.DB) *_AllDummyMgr {
	if db == nil {
		panic(fmt.Errorf("AllDummyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDummyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_dummy"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDummyMgr) GetTableName() string {
	return "__all_dummy"
}

// Reset 重置gorm会话
func (obj *_AllDummyMgr) Reset() *_AllDummyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDummyMgr) Get() (result AllDummy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDummy{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDummyMgr) Gets() (results []*AllDummy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDummyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithKey key获取
func (obj *_AllDummyMgr) WithKey(key int64) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// GetByOption 功能选项模式获取
func (obj *_AllDummyMgr) GetByOption(opts ...Option) (result AllDummy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDummyMgr) GetByOptions(opts ...Option) (results []*AllDummy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDummyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDummy, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Where(options.query)
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

// GetFromKey 通过key获取内容
func (obj *_AllDummyMgr) GetFromKey(key int64) (result AllDummy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Where("`key` = ?", key).First(&result).Error

	return
}

// GetBatchFromKey 批量查找
func (obj *_AllDummyMgr) GetBatchFromKey(keys []int64) (results []*AllDummy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDummyMgr) FetchByPrimaryKey(key int64) (result AllDummy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDummy{}).Where("`key` = ?", key).First(&result).Error

	return
}
