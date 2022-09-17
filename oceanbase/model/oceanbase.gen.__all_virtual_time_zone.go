package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTimeZoneMgr struct {
	*_BaseMgr
}

// AllVirtualTimeZoneMgr open func
func AllVirtualTimeZoneMgr(db *gorm.DB) *_AllVirtualTimeZoneMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTimeZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTimeZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_time_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTimeZoneMgr) GetTableName() string {
	return "__all_virtual_time_zone"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTimeZoneMgr) Reset() *_AllVirtualTimeZoneMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTimeZoneMgr) Get() (result AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTimeZoneMgr) Gets() (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTimeZoneMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTimeZoneMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTimeZoneID time_zone_id获取
func (obj *_AllVirtualTimeZoneMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["time_zone_id"] = timeZoneID })
}

// WithUseLeapSeconds use_leap_seconds获取
func (obj *_AllVirtualTimeZoneMgr) WithUseLeapSeconds(useLeapSeconds string) Option {
	return optionFunc(func(o *options) { o.query["use_leap_seconds"] = useLeapSeconds })
}

// WithVersion version获取
func (obj *_AllVirtualTimeZoneMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTimeZoneMgr) GetByOption(opts ...Option) (result AllVirtualTimeZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTimeZoneMgr) GetByOptions(opts ...Option) (results []*AllVirtualTimeZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTimeZoneMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTimeZone, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where(options.query)
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
func (obj *_AllVirtualTimeZoneMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTimeZoneMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTimeZoneID 通过time_zone_id获取内容
func (obj *_AllVirtualTimeZoneMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllVirtualTimeZoneMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromUseLeapSeconds 通过use_leap_seconds获取内容
func (obj *_AllVirtualTimeZoneMgr) GetFromUseLeapSeconds(useLeapSeconds string) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`use_leap_seconds` = ?", useLeapSeconds).Find(&results).Error

	return
}

// GetBatchFromUseLeapSeconds 批量查找
func (obj *_AllVirtualTimeZoneMgr) GetBatchFromUseLeapSeconds(useLeapSecondss []string) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`use_leap_seconds` IN (?)", useLeapSecondss).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualTimeZoneMgr) GetFromVersion(version int64) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualTimeZoneMgr) GetBatchFromVersion(versions []int64) (results []*AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTimeZoneMgr) FetchByPrimaryKey(tenantID int64, timeZoneID int64) (result AllVirtualTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZone{}).Where("`tenant_id` = ? AND `time_zone_id` = ?", tenantID, timeZoneID).First(&result).Error

	return
}
