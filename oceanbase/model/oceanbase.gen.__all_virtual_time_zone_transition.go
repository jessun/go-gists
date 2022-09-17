package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualTimeZoneTransitionMgr struct {
	*_BaseMgr
}

// AllVirtualTimeZoneTransitionMgr open func
func AllVirtualTimeZoneTransitionMgr(db *gorm.DB) *_AllVirtualTimeZoneTransitionMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTimeZoneTransitionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTimeZoneTransitionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_time_zone_transition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTimeZoneTransitionMgr) GetTableName() string {
	return "__all_virtual_time_zone_transition"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTimeZoneTransitionMgr) Reset() *_AllVirtualTimeZoneTransitionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTimeZoneTransitionMgr) Get() (result AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTimeZoneTransitionMgr) Gets() (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTimeZoneTransitionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTimeZoneTransitionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTimeZoneID time_zone_id获取
func (obj *_AllVirtualTimeZoneTransitionMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["time_zone_id"] = timeZoneID })
}

// WithTransitionTime transition_time获取
func (obj *_AllVirtualTimeZoneTransitionMgr) WithTransitionTime(transitionTime int64) Option {
	return optionFunc(func(o *options) { o.query["transition_time"] = transitionTime })
}

// WithTransitionTypeID transition_type_id获取
func (obj *_AllVirtualTimeZoneTransitionMgr) WithTransitionTypeID(transitionTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["transition_type_id"] = transitionTypeID })
}

// WithVersion version获取
func (obj *_AllVirtualTimeZoneTransitionMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTimeZoneTransitionMgr) GetByOption(opts ...Option) (result AllVirtualTimeZoneTransition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTimeZoneTransitionMgr) GetByOptions(opts ...Option) (results []*AllVirtualTimeZoneTransition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTimeZoneTransitionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTimeZoneTransition, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where(options.query)
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
func (obj *_AllVirtualTimeZoneTransitionMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTimeZoneTransitionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTimeZoneID 通过time_zone_id获取内容
func (obj *_AllVirtualTimeZoneTransitionMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllVirtualTimeZoneTransitionMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromTransitionTime 通过transition_time获取内容
func (obj *_AllVirtualTimeZoneTransitionMgr) GetFromTransitionTime(transitionTime int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`transition_time` = ?", transitionTime).Find(&results).Error

	return
}

// GetBatchFromTransitionTime 批量查找
func (obj *_AllVirtualTimeZoneTransitionMgr) GetBatchFromTransitionTime(transitionTimes []int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`transition_time` IN (?)", transitionTimes).Find(&results).Error

	return
}

// GetFromTransitionTypeID 通过transition_type_id获取内容
func (obj *_AllVirtualTimeZoneTransitionMgr) GetFromTransitionTypeID(transitionTypeID int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`transition_type_id` = ?", transitionTypeID).Find(&results).Error

	return
}

// GetBatchFromTransitionTypeID 批量查找
func (obj *_AllVirtualTimeZoneTransitionMgr) GetBatchFromTransitionTypeID(transitionTypeIDs []int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`transition_type_id` IN (?)", transitionTypeIDs).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualTimeZoneTransitionMgr) GetFromVersion(version int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualTimeZoneTransitionMgr) GetBatchFromVersion(versions []int64) (results []*AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTimeZoneTransitionMgr) FetchByPrimaryKey(tenantID int64, timeZoneID int64, transitionTime int64) (result AllVirtualTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTimeZoneTransition{}).Where("`tenant_id` = ? AND `time_zone_id` = ? AND `transition_time` = ?", tenantID, timeZoneID, transitionTime).First(&result).Error

	return
}
