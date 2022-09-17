package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualProxySysVariableMgr struct {
	*_BaseMgr
}

// AllVirtualProxySysVariableMgr open func
func AllVirtualProxySysVariableMgr(db *gorm.DB) *_AllVirtualProxySysVariableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProxySysVariableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProxySysVariableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_proxy_sys_variable"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProxySysVariableMgr) GetTableName() string {
	return "__all_virtual_proxy_sys_variable"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProxySysVariableMgr) Reset() *_AllVirtualProxySysVariableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProxySysVariableMgr) Get() (result AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProxySysVariableMgr) Gets() (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProxySysVariableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithName name获取
func (obj *_AllVirtualProxySysVariableMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualProxySysVariableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDataType data_type获取
func (obj *_AllVirtualProxySysVariableMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllVirtualProxySysVariableMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithFlags flags获取
func (obj *_AllVirtualProxySysVariableMgr) WithFlags(flags int64) Option {
	return optionFunc(func(o *options) { o.query["flags"] = flags })
}

// WithModifiedTime modified_time获取
func (obj *_AllVirtualProxySysVariableMgr) WithModifiedTime(modifiedTime int64) Option {
	return optionFunc(func(o *options) { o.query["modified_time"] = modifiedTime })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProxySysVariableMgr) GetByOption(opts ...Option) (result AllVirtualProxySysVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProxySysVariableMgr) GetByOptions(opts ...Option) (results []*AllVirtualProxySysVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProxySysVariableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProxySysVariable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where(options.query)
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
func (obj *_AllVirtualProxySysVariableMgr) GetFromName(name string) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualProxySysVariableMgr) GetBatchFromName(names []string) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualProxySysVariableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualProxySysVariableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualProxySysVariableMgr) GetFromDataType(dataType int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualProxySysVariableMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualProxySysVariableMgr) GetFromValue(value string) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualProxySysVariableMgr) GetBatchFromValue(values []string) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromFlags 通过flags获取内容
func (obj *_AllVirtualProxySysVariableMgr) GetFromFlags(flags int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`flags` = ?", flags).Find(&results).Error

	return
}

// GetBatchFromFlags 批量查找
func (obj *_AllVirtualProxySysVariableMgr) GetBatchFromFlags(flagss []int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`flags` IN (?)", flagss).Find(&results).Error

	return
}

// GetFromModifiedTime 通过modified_time获取内容
func (obj *_AllVirtualProxySysVariableMgr) GetFromModifiedTime(modifiedTime int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`modified_time` = ?", modifiedTime).Find(&results).Error

	return
}

// GetBatchFromModifiedTime 批量查找
func (obj *_AllVirtualProxySysVariableMgr) GetBatchFromModifiedTime(modifiedTimes []int64) (results []*AllVirtualProxySysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySysVariable{}).Where("`modified_time` IN (?)", modifiedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
