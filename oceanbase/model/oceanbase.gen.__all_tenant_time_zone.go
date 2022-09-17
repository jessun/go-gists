package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllTenantTimeZoneMgr struct {
	*_BaseMgr
}

// AllTenantTimeZoneMgr open func
func AllTenantTimeZoneMgr(db *gorm.DB) *_AllTenantTimeZoneMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantTimeZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantTimeZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_time_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantTimeZoneMgr) GetTableName() string {
	return "__all_tenant_time_zone"
}

// Reset 重置gorm会话
func (obj *_AllTenantTimeZoneMgr) Reset() *_AllTenantTimeZoneMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantTimeZoneMgr) Get() (result AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantTimeZoneMgr) Gets() (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantTimeZoneMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllTenantTimeZoneMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTimeZoneID time_zone_id获取
func (obj *_AllTenantTimeZoneMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["time_zone_id"] = timeZoneID })
}

// WithUseLeapSeconds use_leap_seconds获取
func (obj *_AllTenantTimeZoneMgr) WithUseLeapSeconds(useLeapSeconds string) Option {
	return optionFunc(func(o *options) { o.query["use_leap_seconds"] = useLeapSeconds })
}

// WithVersion version获取
func (obj *_AllTenantTimeZoneMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantTimeZoneMgr) GetByOption(opts ...Option) (result AllTenantTimeZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantTimeZoneMgr) GetByOptions(opts ...Option) (results []*AllTenantTimeZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantTimeZoneMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantTimeZone, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where(options.query)
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
func (obj *_AllTenantTimeZoneMgr) GetFromTenantID(tenantID int64) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantTimeZoneMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTimeZoneID 通过time_zone_id获取内容
func (obj *_AllTenantTimeZoneMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllTenantTimeZoneMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromUseLeapSeconds 通过use_leap_seconds获取内容
func (obj *_AllTenantTimeZoneMgr) GetFromUseLeapSeconds(useLeapSeconds string) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`use_leap_seconds` = ?", useLeapSeconds).Find(&results).Error

	return
}

// GetBatchFromUseLeapSeconds 批量查找
func (obj *_AllTenantTimeZoneMgr) GetBatchFromUseLeapSeconds(useLeapSecondss []string) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`use_leap_seconds` IN (?)", useLeapSecondss).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllTenantTimeZoneMgr) GetFromVersion(version int64) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllTenantTimeZoneMgr) GetBatchFromVersion(versions []int64) (results []*AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantTimeZoneMgr) FetchByPrimaryKey(tenantID int64, timeZoneID int64) (result AllTenantTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZone{}).Where("`tenant_id` = ? AND `time_zone_id` = ?", tenantID, timeZoneID).First(&result).Error

	return
}
