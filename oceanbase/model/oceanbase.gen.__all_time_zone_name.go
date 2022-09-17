package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllTimeZoneNameMgr struct {
	*_BaseMgr
}

// AllTimeZoneNameMgr open func
func AllTimeZoneNameMgr(db *gorm.DB) *_AllTimeZoneNameMgr {
	if db == nil {
		panic(fmt.Errorf("AllTimeZoneNameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTimeZoneNameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_time_zone_name"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTimeZoneNameMgr) GetTableName() string {
	return "__all_time_zone_name"
}

// Reset 重置gorm会话
func (obj *_AllTimeZoneNameMgr) Reset() *_AllTimeZoneNameMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTimeZoneNameMgr) Get() (result AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTimeZoneNameMgr) Gets() (results []*AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTimeZoneNameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithName Name获取
func (obj *_AllTimeZoneNameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithTimeZoneID Time_zone_id获取
func (obj *_AllTimeZoneNameMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["Time_zone_id"] = timeZoneID })
}

// WithVersion Version获取
func (obj *_AllTimeZoneNameMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["Version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllTimeZoneNameMgr) GetByOption(opts ...Option) (result AllTimeZoneName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTimeZoneNameMgr) GetByOptions(opts ...Option) (results []*AllTimeZoneName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTimeZoneNameMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTimeZoneName, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where(options.query)
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

// GetFromName 通过Name获取内容
func (obj *_AllTimeZoneNameMgr) GetFromName(name string) (result AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where("`Name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllTimeZoneNameMgr) GetBatchFromName(names []string) (results []*AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where("`Name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTimeZoneID 通过Time_zone_id获取内容
func (obj *_AllTimeZoneNameMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where("`Time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllTimeZoneNameMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where("`Time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromVersion 通过Version获取内容
func (obj *_AllTimeZoneNameMgr) GetFromVersion(version int64) (results []*AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where("`Version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllTimeZoneNameMgr) GetBatchFromVersion(versions []int64) (results []*AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where("`Version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTimeZoneNameMgr) FetchByPrimaryKey(name string) (result AllTimeZoneName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneName{}).Where("`Name` = ?", name).First(&result).Error

	return
}
