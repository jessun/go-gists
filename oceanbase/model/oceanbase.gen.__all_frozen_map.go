package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllFrozenMapMgr struct {
	*_BaseMgr
}

// AllFrozenMapMgr open func
func AllFrozenMapMgr(db *gorm.DB) *_AllFrozenMapMgr {
	if db == nil {
		panic(fmt.Errorf("AllFrozenMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllFrozenMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_frozen_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllFrozenMapMgr) GetTableName() string {
	return "__all_frozen_map"
}

// Reset 重置gorm会话
func (obj *_AllFrozenMapMgr) Reset() *_AllFrozenMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllFrozenMapMgr) Get() (result AllFrozenMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllFrozenMapMgr) Gets() (results []*AllFrozenMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllFrozenMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithFrozenVersion frozen_version获取
func (obj *_AllFrozenMapMgr) WithFrozenVersion(frozenVersion int64) Option {
	return optionFunc(func(o *options) { o.query["frozen_version"] = frozenVersion })
}

// WithFrozenTimestamp frozen_timestamp获取
func (obj *_AllFrozenMapMgr) WithFrozenTimestamp(frozenTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["frozen_timestamp"] = frozenTimestamp })
}

// GetByOption 功能选项模式获取
func (obj *_AllFrozenMapMgr) GetByOption(opts ...Option) (result AllFrozenMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllFrozenMapMgr) GetByOptions(opts ...Option) (results []*AllFrozenMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllFrozenMapMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllFrozenMap, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where(options.query)
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

// GetFromFrozenVersion 通过frozen_version获取内容
func (obj *_AllFrozenMapMgr) GetFromFrozenVersion(frozenVersion int64) (results []*AllFrozenMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where("`frozen_version` = ?", frozenVersion).Find(&results).Error

	return
}

// GetBatchFromFrozenVersion 批量查找
func (obj *_AllFrozenMapMgr) GetBatchFromFrozenVersion(frozenVersions []int64) (results []*AllFrozenMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where("`frozen_version` IN (?)", frozenVersions).Find(&results).Error

	return
}

// GetFromFrozenTimestamp 通过frozen_timestamp获取内容
func (obj *_AllFrozenMapMgr) GetFromFrozenTimestamp(frozenTimestamp int64) (results []*AllFrozenMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where("`frozen_timestamp` = ?", frozenTimestamp).Find(&results).Error

	return
}

// GetBatchFromFrozenTimestamp 批量查找
func (obj *_AllFrozenMapMgr) GetBatchFromFrozenTimestamp(frozenTimestamps []int64) (results []*AllFrozenMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where("`frozen_timestamp` IN (?)", frozenTimestamps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllFrozenMapMgr) FetchByPrimaryKey(frozenVersion int64, frozenTimestamp int64) (result AllFrozenMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFrozenMap{}).Where("`frozen_version` = ? AND `frozen_timestamp` = ?", frozenVersion, frozenTimestamp).First(&result).Error

	return
}
