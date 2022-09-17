package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllTimeZoneTransitionMgr struct {
	*_BaseMgr
}

// AllTimeZoneTransitionMgr open func
func AllTimeZoneTransitionMgr(db *gorm.DB) *_AllTimeZoneTransitionMgr {
	if db == nil {
		panic(fmt.Errorf("AllTimeZoneTransitionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTimeZoneTransitionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_time_zone_transition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTimeZoneTransitionMgr) GetTableName() string {
	return "__all_time_zone_transition"
}

// Reset 重置gorm会话
func (obj *_AllTimeZoneTransitionMgr) Reset() *_AllTimeZoneTransitionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTimeZoneTransitionMgr) Get() (result AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTimeZoneTransitionMgr) Gets() (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTimeZoneTransitionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTimeZoneID Time_zone_id获取
func (obj *_AllTimeZoneTransitionMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["Time_zone_id"] = timeZoneID })
}

// WithTransitionTime Transition_time获取
func (obj *_AllTimeZoneTransitionMgr) WithTransitionTime(transitionTime int64) Option {
	return optionFunc(func(o *options) { o.query["Transition_time"] = transitionTime })
}

// WithTransitionTypeID Transition_type_id获取
func (obj *_AllTimeZoneTransitionMgr) WithTransitionTypeID(transitionTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["Transition_type_id"] = transitionTypeID })
}

// WithVersion Version获取
func (obj *_AllTimeZoneTransitionMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["Version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllTimeZoneTransitionMgr) GetByOption(opts ...Option) (result AllTimeZoneTransition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTimeZoneTransitionMgr) GetByOptions(opts ...Option) (results []*AllTimeZoneTransition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTimeZoneTransitionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTimeZoneTransition, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where(options.query)
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
func (obj *_AllTimeZoneTransitionMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllTimeZoneTransitionMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromTransitionTime 通过Transition_time获取内容
func (obj *_AllTimeZoneTransitionMgr) GetFromTransitionTime(transitionTime int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Transition_time` = ?", transitionTime).Find(&results).Error

	return
}

// GetBatchFromTransitionTime 批量查找
func (obj *_AllTimeZoneTransitionMgr) GetBatchFromTransitionTime(transitionTimes []int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Transition_time` IN (?)", transitionTimes).Find(&results).Error

	return
}

// GetFromTransitionTypeID 通过Transition_type_id获取内容
func (obj *_AllTimeZoneTransitionMgr) GetFromTransitionTypeID(transitionTypeID int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Transition_type_id` = ?", transitionTypeID).Find(&results).Error

	return
}

// GetBatchFromTransitionTypeID 批量查找
func (obj *_AllTimeZoneTransitionMgr) GetBatchFromTransitionTypeID(transitionTypeIDs []int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Transition_type_id` IN (?)", transitionTypeIDs).Find(&results).Error

	return
}

// GetFromVersion 通过Version获取内容
func (obj *_AllTimeZoneTransitionMgr) GetFromVersion(version int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllTimeZoneTransitionMgr) GetBatchFromVersion(versions []int64) (results []*AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTimeZoneTransitionMgr) FetchByPrimaryKey(timeZoneID int64, transitionTime int64) (result AllTimeZoneTransition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransition{}).Where("`Time_zone_id` = ? AND `Transition_time` = ?", timeZoneID, transitionTime).First(&result).Error

	return
}
