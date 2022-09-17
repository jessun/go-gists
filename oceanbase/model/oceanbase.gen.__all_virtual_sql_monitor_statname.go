package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualSQLMonitorStatnameMgr struct {
	*_BaseMgr
}

// AllVirtualSQLMonitorStatnameMgr open func
func AllVirtualSQLMonitorStatnameMgr(db *gorm.DB) *_AllVirtualSQLMonitorStatnameMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLMonitorStatnameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLMonitorStatnameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_monitor_statname"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetTableName() string {
	return "__all_virtual_sql_monitor_statname"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLMonitorStatnameMgr) Reset() *_AllVirtualSQLMonitorStatnameMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLMonitorStatnameMgr) Get() (result AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLMonitorStatnameMgr) Gets() (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLMonitorStatnameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_AllVirtualSQLMonitorStatnameMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithGroupID GROUP_ID获取
func (obj *_AllVirtualSQLMonitorStatnameMgr) WithGroupID(groupID int64) Option {
	return optionFunc(func(o *options) { o.query["GROUP_ID"] = groupID })
}

// WithName NAME获取
func (obj *_AllVirtualSQLMonitorStatnameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithDescription DESCRIPTION获取
func (obj *_AllVirtualSQLMonitorStatnameMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["DESCRIPTION"] = description })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetByOption(opts ...Option) (result AllVirtualSQLMonitorStatname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLMonitorStatname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLMonitorStatnameMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLMonitorStatname, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where(options.query)
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

// GetFromID 通过ID获取内容
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetFromID(id int64) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`ID` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetBatchFromID(ids []int64) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`ID` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupID 通过GROUP_ID获取内容
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetFromGroupID(groupID int64) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`GROUP_ID` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetBatchFromGroupID(groupIDs []int64) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`GROUP_ID` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromName 通过NAME获取内容
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetFromName(name string) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`NAME` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetBatchFromName(names []string) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`NAME` IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过DESCRIPTION获取内容
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetFromDescription(description string) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`DESCRIPTION` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找
func (obj *_AllVirtualSQLMonitorStatnameMgr) GetBatchFromDescription(descriptions []string) (results []*AllVirtualSQLMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLMonitorStatname{}).Where("`DESCRIPTION` IN (?)", descriptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
