package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualFreezeInfoMgr struct {
	*_BaseMgr
}

// AllVirtualFreezeInfoMgr open func
func AllVirtualFreezeInfoMgr(db *gorm.DB) *_AllVirtualFreezeInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualFreezeInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualFreezeInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_freeze_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualFreezeInfoMgr) GetTableName() string {
	return "__all_virtual_freeze_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualFreezeInfoMgr) Reset() *_AllVirtualFreezeInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualFreezeInfoMgr) Get() (result AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualFreezeInfoMgr) Gets() (results []*AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualFreezeInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithFrozenVersion frozen_version获取
func (obj *_AllVirtualFreezeInfoMgr) WithFrozenVersion(frozenVersion int64) Option {
	return optionFunc(func(o *options) { o.query["frozen_version"] = frozenVersion })
}

// WithFrozenTimestamp frozen_timestamp获取
func (obj *_AllVirtualFreezeInfoMgr) WithFrozenTimestamp(frozenTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["frozen_timestamp"] = frozenTimestamp })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualFreezeInfoMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualFreezeInfoMgr) GetByOption(opts ...Option) (result AllVirtualFreezeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualFreezeInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualFreezeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualFreezeInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualFreezeInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where(options.query)
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
func (obj *_AllVirtualFreezeInfoMgr) GetFromFrozenVersion(frozenVersion int64) (result AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where("`frozen_version` = ?", frozenVersion).First(&result).Error

	return
}

// GetBatchFromFrozenVersion 批量查找
func (obj *_AllVirtualFreezeInfoMgr) GetBatchFromFrozenVersion(frozenVersions []int64) (results []*AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where("`frozen_version` IN (?)", frozenVersions).Find(&results).Error

	return
}

// GetFromFrozenTimestamp 通过frozen_timestamp获取内容
func (obj *_AllVirtualFreezeInfoMgr) GetFromFrozenTimestamp(frozenTimestamp int64) (results []*AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where("`frozen_timestamp` = ?", frozenTimestamp).Find(&results).Error

	return
}

// GetBatchFromFrozenTimestamp 批量查找
func (obj *_AllVirtualFreezeInfoMgr) GetBatchFromFrozenTimestamp(frozenTimestamps []int64) (results []*AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where("`frozen_timestamp` IN (?)", frozenTimestamps).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualFreezeInfoMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualFreezeInfoMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualFreezeInfoMgr) FetchByPrimaryKey(frozenVersion int64) (result AllVirtualFreezeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFreezeInfo{}).Where("`frozen_version` = ?", frozenVersion).First(&result).Error

	return
}
