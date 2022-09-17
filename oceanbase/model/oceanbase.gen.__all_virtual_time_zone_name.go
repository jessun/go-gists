package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTimeZoneNameMgr struct {
	*_BaseMgr
}

// AllVirtualTimeZoneNameMgr open func
func AllVirtualTimeZoneNameMgr(db *gorm.DB) *_AllVirtualTimeZoneNameMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTimeZoneNameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTimeZoneNameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_time_zone_name"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTimeZoneNameMgr) GetTableName() string {
	return "__all_virtual_time_zone_name"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTimeZoneNameMgr) Reset() *_AllVirtualTimeZoneNameMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTimeZoneNameMgr) Get() (result AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTimeZoneNameMgr) Gets() (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTimeZoneNameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTimeZoneNameMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithName name获取
func (obj *_AllVirtualTimeZoneNameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTimeZoneID time_zone_id获取
func (obj *_AllVirtualTimeZoneNameMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["time_zone_id"] = timeZoneID })
}

// WithVersion version获取
func (obj *_AllVirtualTimeZoneNameMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTimeZoneNameMgr) GetByOption(opts ...Option) (result AllVirtualTimeZoneName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTimeZoneNameMgr) GetByOptions(opts ...Option) (results []*AllVirtualTimeZoneName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTimeZoneNameMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTimeZoneName, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where(options.query)
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
func (obj *_AllVirtualTimeZoneNameMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTimeZoneNameMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualTimeZoneNameMgr) GetFromName(name string) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualTimeZoneNameMgr) GetBatchFromName(names []string) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTimeZoneID 通过time_zone_id获取内容
func (obj *_AllVirtualTimeZoneNameMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllVirtualTimeZoneNameMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualTimeZoneNameMgr) GetFromVersion(version int64) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualTimeZoneNameMgr) GetBatchFromVersion(versions []int64) (results []*AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTimeZoneNameMgr) FetchByPrimaryKey(tenantID int64, name string) (result AllVirtualTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneName{}).Where("`tenant_id` = ? AND `name` = ?", tenantID, name).First(&result).Error

	return
}
