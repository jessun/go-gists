package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualShowCreateTablegroupMgr struct {
	*_BaseMgr
}

// TenantVirtualShowCreateTablegroupMgr open func
func TenantVirtualShowCreateTablegroupMgr(db *gorm.DB) *_TenantVirtualShowCreateTablegroupMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualShowCreateTablegroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualShowCreateTablegroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_show_create_tablegroup"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetTableName() string {
	return "__tenant_virtual_show_create_tablegroup"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualShowCreateTablegroupMgr) Reset() *_TenantVirtualShowCreateTablegroupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualShowCreateTablegroupMgr) Get() (result TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualShowCreateTablegroupMgr) Gets() (results []*TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualShowCreateTablegroupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTablegroupID tablegroup_id获取
func (obj *_TenantVirtualShowCreateTablegroupMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithTablegroupName tablegroup_name获取
func (obj *_TenantVirtualShowCreateTablegroupMgr) WithTablegroupName(tablegroupName string) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_name"] = tablegroupName })
}

// WithCreateTablegroup create_tablegroup获取
func (obj *_TenantVirtualShowCreateTablegroupMgr) WithCreateTablegroup(createTablegroup string) Option {
	return optionFunc(func(o *options) { o.query["create_tablegroup"] = createTablegroup })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetByOption(opts ...Option) (result TenantVirtualShowCreateTablegroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetByOptions(opts ...Option) (results []*TenantVirtualShowCreateTablegroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualShowCreateTablegroupMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualShowCreateTablegroup, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where(options.query)
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

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetFromTablegroupID(tablegroupID int64) (result TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where("`tablegroup_id` = ?", tablegroupID).First(&result).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromTablegroupName 通过tablegroup_name获取内容
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetFromTablegroupName(tablegroupName string) (results []*TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where("`tablegroup_name` = ?", tablegroupName).Find(&results).Error

	return
}

// GetBatchFromTablegroupName 批量查找
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetBatchFromTablegroupName(tablegroupNames []string) (results []*TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where("`tablegroup_name` IN (?)", tablegroupNames).Find(&results).Error

	return
}

// GetFromCreateTablegroup 通过create_tablegroup获取内容
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetFromCreateTablegroup(createTablegroup string) (results []*TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where("`create_tablegroup` = ?", createTablegroup).Find(&results).Error

	return
}

// GetBatchFromCreateTablegroup 批量查找
func (obj *_TenantVirtualShowCreateTablegroupMgr) GetBatchFromCreateTablegroup(createTablegroups []string) (results []*TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where("`create_tablegroup` IN (?)", createTablegroups).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualShowCreateTablegroupMgr) FetchByPrimaryKey(tablegroupID int64) (result TenantVirtualShowCreateTablegroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateTablegroup{}).Where("`tablegroup_id` = ?", tablegroupID).First(&result).Error

	return
}
