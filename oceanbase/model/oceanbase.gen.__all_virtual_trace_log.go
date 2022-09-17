package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTraceLogMgr struct {
	*_BaseMgr
}

// AllVirtualTraceLogMgr open func
func AllVirtualTraceLogMgr(db *gorm.DB) *_AllVirtualTraceLogMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTraceLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTraceLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trace_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTraceLogMgr) GetTableName() string {
	return "__all_virtual_trace_log"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTraceLogMgr) Reset() *_AllVirtualTraceLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTraceLogMgr) Get() (result AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTraceLogMgr) Gets() (results []*AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTraceLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTitle title获取
func (obj *_AllVirtualTraceLogMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithKeyValue key_value获取
func (obj *_AllVirtualTraceLogMgr) WithKeyValue(keyValue string) Option {
	return optionFunc(func(o *options) { o.query["key_value"] = keyValue })
}

// WithTime time获取
func (obj *_AllVirtualTraceLogMgr) WithTime(time string) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTraceLogMgr) GetByOption(opts ...Option) (result AllVirtualTraceLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTraceLogMgr) GetByOptions(opts ...Option) (results []*AllVirtualTraceLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTraceLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTraceLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where(options.query)
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

// GetFromTitle 通过title获取内容
func (obj *_AllVirtualTraceLogMgr) GetFromTitle(title string) (results []*AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找
func (obj *_AllVirtualTraceLogMgr) GetBatchFromTitle(titles []string) (results []*AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromKeyValue 通过key_value获取内容
func (obj *_AllVirtualTraceLogMgr) GetFromKeyValue(keyValue string) (results []*AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where("`key_value` = ?", keyValue).Find(&results).Error

	return
}

// GetBatchFromKeyValue 批量查找
func (obj *_AllVirtualTraceLogMgr) GetBatchFromKeyValue(keyValues []string) (results []*AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where("`key_value` IN (?)", keyValues).Find(&results).Error

	return
}

// GetFromTime 通过time获取内容
func (obj *_AllVirtualTraceLogMgr) GetFromTime(time string) (results []*AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where("`time` = ?", time).Find(&results).Error

	return
}

// GetBatchFromTime 批量查找
func (obj *_AllVirtualTraceLogMgr) GetBatchFromTime(times []string) (results []*AllVirtualTraceLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTraceLog{}).Where("`time` IN (?)", times).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
