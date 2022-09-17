package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTimestampServiceMgr struct {
	*_BaseMgr
}

// AllVirtualTimestampServiceMgr open func
func AllVirtualTimestampServiceMgr(db *gorm.DB) *_AllVirtualTimestampServiceMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTimestampServiceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTimestampServiceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_timestamp_service"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTimestampServiceMgr) GetTableName() string {
	return "__all_virtual_timestamp_service"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTimestampServiceMgr) Reset() *_AllVirtualTimestampServiceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTimestampServiceMgr) Get() (result AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTimestampServiceMgr) Gets() (results []*AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTimestampServiceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTimestampServiceMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTsType ts_type获取
func (obj *_AllVirtualTimestampServiceMgr) WithTsType(tsType int64) Option {
	return optionFunc(func(o *options) { o.query["ts_type"] = tsType })
}

// WithTsValue ts_value获取
func (obj *_AllVirtualTimestampServiceMgr) WithTsValue(tsValue int64) Option {
	return optionFunc(func(o *options) { o.query["ts_value"] = tsValue })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTimestampServiceMgr) GetByOption(opts ...Option) (result AllVirtualTimestampService, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTimestampServiceMgr) GetByOptions(opts ...Option) (results []*AllVirtualTimestampService, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTimestampServiceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTimestampService, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualTimestampServiceMgr) GetFromTenantID(tenantID int64) (result AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTimestampServiceMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTsType 通过ts_type获取内容
func (obj *_AllVirtualTimestampServiceMgr) GetFromTsType(tsType int64) (results []*AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where("`ts_type` = ?", tsType).Find(&results).Error

	return
}

// GetBatchFromTsType 批量查找
func (obj *_AllVirtualTimestampServiceMgr) GetBatchFromTsType(tsTypes []int64) (results []*AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where("`ts_type` IN (?)", tsTypes).Find(&results).Error

	return
}

// GetFromTsValue 通过ts_value获取内容
func (obj *_AllVirtualTimestampServiceMgr) GetFromTsValue(tsValue int64) (results []*AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where("`ts_value` = ?", tsValue).Find(&results).Error

	return
}

// GetBatchFromTsValue 批量查找
func (obj *_AllVirtualTimestampServiceMgr) GetBatchFromTsValue(tsValues []int64) (results []*AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where("`ts_value` IN (?)", tsValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTimestampServiceMgr) FetchByPrimaryKey(tenantID int64) (result AllVirtualTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimestampService{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
