package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllTimeZoneMgr struct {
	*_BaseMgr
}

// AllTimeZoneMgr open func
func AllTimeZoneMgr(db *gorm.DB) *_AllTimeZoneMgr {
	if db == nil {
		panic(fmt.Errorf("AllTimeZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTimeZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_time_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTimeZoneMgr) GetTableName() string {
	return "__all_time_zone"
}

// Reset 重置gorm会话
func (obj *_AllTimeZoneMgr) Reset() *_AllTimeZoneMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTimeZoneMgr) Get() (result AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTimeZoneMgr) Gets() (results []*AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTimeZoneMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTimeZoneID Time_zone_id获取
func (obj *_AllTimeZoneMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["Time_zone_id"] = timeZoneID })
}

// WithUseLeapSeconds Use_leap_seconds获取
func (obj *_AllTimeZoneMgr) WithUseLeapSeconds(useLeapSeconds string) Option {
	return optionFunc(func(o *options) { o.query["Use_leap_seconds"] = useLeapSeconds })
}

// WithVersion Version获取
func (obj *_AllTimeZoneMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["Version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllTimeZoneMgr) GetByOption(opts ...Option) (result AllTimeZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTimeZoneMgr) GetByOptions(opts ...Option) (results []*AllTimeZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTimeZoneMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTimeZone, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where(options.query)
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

// GetFromTimeZoneID 通过Time_zone_id获取内容
func (obj *_AllTimeZoneMgr) GetFromTimeZoneID(timeZoneID int64) (result AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where("`Time_zone_id` = ?", timeZoneID).First(&result).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllTimeZoneMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where("`Time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromUseLeapSeconds 通过Use_leap_seconds获取内容
func (obj *_AllTimeZoneMgr) GetFromUseLeapSeconds(useLeapSeconds string) (results []*AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where("`Use_leap_seconds` = ?", useLeapSeconds).Find(&results).Error

	return
}

// GetBatchFromUseLeapSeconds 批量查找
func (obj *_AllTimeZoneMgr) GetBatchFromUseLeapSeconds(useLeapSecondss []string) (results []*AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where("`Use_leap_seconds` IN (?)", useLeapSecondss).Find(&results).Error

	return
}

// GetFromVersion 通过Version获取内容
func (obj *_AllTimeZoneMgr) GetFromVersion(version int64) (results []*AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where("`Version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllTimeZoneMgr) GetBatchFromVersion(versions []int64) (results []*AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where("`Version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTimeZoneMgr) FetchByPrimaryKey(timeZoneID int64) (result AllTimeZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZone{}).Where("`Time_zone_id` = ?", timeZoneID).First(&result).Error

	return
}
