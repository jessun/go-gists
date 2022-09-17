package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualSessionVariableMgr struct {
	*_BaseMgr
}

// TenantVirtualSessionVariableMgr open func
func TenantVirtualSessionVariableMgr(db *gorm.DB) *_TenantVirtualSessionVariableMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualSessionVariableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualSessionVariableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_session_variable"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualSessionVariableMgr) GetTableName() string {
	return "__tenant_virtual_session_variable"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualSessionVariableMgr) Reset() *_TenantVirtualSessionVariableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualSessionVariableMgr) Get() (result TenantVirtualSessionVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualSessionVariableMgr) Gets() (results []*TenantVirtualSessionVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualSessionVariableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithVariableName variable_name获取
func (obj *_TenantVirtualSessionVariableMgr) WithVariableName(variableName string) Option {
	return optionFunc(func(o *options) { o.query["variable_name"] = variableName })
}

// WithValue value获取
func (obj *_TenantVirtualSessionVariableMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualSessionVariableMgr) GetByOption(opts ...Option) (result TenantVirtualSessionVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualSessionVariableMgr) GetByOptions(opts ...Option) (results []*TenantVirtualSessionVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualSessionVariableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualSessionVariable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Where(options.query)
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
func (obj *_TenantVirtualSessionVariableMgr) GetFromVariableName(variableName string) (results []*TenantVirtualSessionVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Where("`variable_name` = ?", variableName).Find(&results).Error

	return
}

// GetBatchFromVariableName 批量查找
func (obj *_TenantVirtualSessionVariableMgr) GetBatchFromVariableName(variableNames []string) (results []*TenantVirtualSessionVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Where("`variable_name` IN (?)", variableNames).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_TenantVirtualSessionVariableMgr) GetFromValue(value string) (results []*TenantVirtualSessionVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_TenantVirtualSessionVariableMgr) GetBatchFromValue(values []string) (results []*TenantVirtualSessionVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualSessionVariable{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
