package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualGlobalVariableMgr struct {
	*_BaseMgr
}

// TenantVirtualGlobalVariableMgr open func
func TenantVirtualGlobalVariableMgr(db *gorm.DB) *_TenantVirtualGlobalVariableMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualGlobalVariableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualGlobalVariableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_global_variable"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualGlobalVariableMgr) GetTableName() string {
	return "__tenant_virtual_global_variable"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualGlobalVariableMgr) Reset() *_TenantVirtualGlobalVariableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualGlobalVariableMgr) Get() (result TenantVirtualGlobalVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualGlobalVariableMgr) Gets() (results []*TenantVirtualGlobalVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualGlobalVariableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithVariableName variable_name获取
func (obj *_TenantVirtualGlobalVariableMgr) WithVariableName(variableName string) Option {
	return optionFunc(func(o *options) { o.query["variable_name"] = variableName })
}

// WithValue value获取
func (obj *_TenantVirtualGlobalVariableMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualGlobalVariableMgr) GetByOption(opts ...Option) (result TenantVirtualGlobalVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualGlobalVariableMgr) GetByOptions(opts ...Option) (results []*TenantVirtualGlobalVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualGlobalVariableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualGlobalVariable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Where(options.query)
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

// GetFromVariableName 通过variable_name获取内容
func (obj *_TenantVirtualGlobalVariableMgr) GetFromVariableName(variableName string) (results []*TenantVirtualGlobalVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Where("`variable_name` = ?", variableName).Find(&results).Error

	return
}

// GetBatchFromVariableName 批量查找
func (obj *_TenantVirtualGlobalVariableMgr) GetBatchFromVariableName(variableNames []string) (results []*TenantVirtualGlobalVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Where("`variable_name` IN (?)", variableNames).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_TenantVirtualGlobalVariableMgr) GetFromValue(value string) (results []*TenantVirtualGlobalVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_TenantVirtualGlobalVariableMgr) GetBatchFromValue(values []string) (results []*TenantVirtualGlobalVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualGlobalVariable{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
