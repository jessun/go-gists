package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualUpgradeInspectionMgr struct {
	*_BaseMgr
}

// AllVirtualUpgradeInspectionMgr open func
func AllVirtualUpgradeInspectionMgr(db *gorm.DB) *_AllVirtualUpgradeInspectionMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualUpgradeInspectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualUpgradeInspectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_upgrade_inspection"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualUpgradeInspectionMgr) GetTableName() string {
	return "__all_virtual_upgrade_inspection"
}

// Reset 重置gorm会话
func (obj *_AllVirtualUpgradeInspectionMgr) Reset() *_AllVirtualUpgradeInspectionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualUpgradeInspectionMgr) Get() (result AllVirtualUpgradeInspection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualUpgradeInspectionMgr) Gets() (results []*AllVirtualUpgradeInspection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualUpgradeInspectionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithName name获取
func (obj *_AllVirtualUpgradeInspectionMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithInfo info获取
func (obj *_AllVirtualUpgradeInspectionMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualUpgradeInspectionMgr) GetByOption(opts ...Option) (result AllVirtualUpgradeInspection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualUpgradeInspectionMgr) GetByOptions(opts ...Option) (results []*AllVirtualUpgradeInspection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualUpgradeInspectionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualUpgradeInspection, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Where(options.query)
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

// GetFromName 通过name获取内容
func (obj *_AllVirtualUpgradeInspectionMgr) GetFromName(name string) (results []*AllVirtualUpgradeInspection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualUpgradeInspectionMgr) GetBatchFromName(names []string) (results []*AllVirtualUpgradeInspection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualUpgradeInspectionMgr) GetFromInfo(info string) (results []*AllVirtualUpgradeInspection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualUpgradeInspectionMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualUpgradeInspection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUpgradeInspection{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
