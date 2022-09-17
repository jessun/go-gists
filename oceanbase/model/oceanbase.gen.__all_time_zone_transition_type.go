package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllTimeZoneTransitionTypeMgr struct {
	*_BaseMgr
}

// AllTimeZoneTransitionTypeMgr open func
func AllTimeZoneTransitionTypeMgr(db *gorm.DB) *_AllTimeZoneTransitionTypeMgr {
	if db == nil {
		panic(fmt.Errorf("AllTimeZoneTransitionTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTimeZoneTransitionTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_time_zone_transition_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTimeZoneTransitionTypeMgr) GetTableName() string {
	return "__all_time_zone_transition_type"
}

// Reset 重置gorm会话
func (obj *_AllTimeZoneTransitionTypeMgr) Reset() *_AllTimeZoneTransitionTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTimeZoneTransitionTypeMgr) Get() (result AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTimeZoneTransitionTypeMgr) Gets() (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTimeZoneTransitionTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTimeZoneID Time_zone_id获取
func (obj *_AllTimeZoneTransitionTypeMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["Time_zone_id"] = timeZoneID })
}

// WithTransitionTypeID Transition_type_id获取
func (obj *_AllTimeZoneTransitionTypeMgr) WithTransitionTypeID(transitionTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["Transition_type_id"] = transitionTypeID })
}

// WithOffset Offset获取
func (obj *_AllTimeZoneTransitionTypeMgr) WithOffset(offset int64) Option {
	return optionFunc(func(o *options) { o.query["Offset"] = offset })
}

// WithIsDst Is_DST获取
func (obj *_AllTimeZoneTransitionTypeMgr) WithIsDst(isDst int64) Option {
	return optionFunc(func(o *options) { o.query["Is_DST"] = isDst })
}

// WithAbbreviation Abbreviation获取
func (obj *_AllTimeZoneTransitionTypeMgr) WithAbbreviation(abbreviation string) Option {
	return optionFunc(func(o *options) { o.query["Abbreviation"] = abbreviation })
}

// WithVersion Version获取
func (obj *_AllTimeZoneTransitionTypeMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["Version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllTimeZoneTransitionTypeMgr) GetByOption(opts ...Option) (result AllTimeZoneTransitionType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTimeZoneTransitionTypeMgr) GetByOptions(opts ...Option) (results []*AllTimeZoneTransitionType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTimeZoneTransitionTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTimeZoneTransitionType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where(options.query)
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
func (obj *_AllTimeZoneTransitionTypeMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllTimeZoneTransitionTypeMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromTransitionTypeID 通过Transition_type_id获取内容
func (obj *_AllTimeZoneTransitionTypeMgr) GetFromTransitionTypeID(transitionTypeID int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Transition_type_id` = ?", transitionTypeID).Find(&results).Error

	return
}

// GetBatchFromTransitionTypeID 批量查找
func (obj *_AllTimeZoneTransitionTypeMgr) GetBatchFromTransitionTypeID(transitionTypeIDs []int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Transition_type_id` IN (?)", transitionTypeIDs).Find(&results).Error

	return
}

// GetFromOffset 通过Offset获取内容
func (obj *_AllTimeZoneTransitionTypeMgr) GetFromOffset(offset int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Offset` = ?", offset).Find(&results).Error

	return
}

// GetBatchFromOffset 批量查找
func (obj *_AllTimeZoneTransitionTypeMgr) GetBatchFromOffset(offsets []int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Offset` IN (?)", offsets).Find(&results).Error

	return
}

// GetFromIsDst 通过Is_DST获取内容
func (obj *_AllTimeZoneTransitionTypeMgr) GetFromIsDst(isDst int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Is_DST` = ?", isDst).Find(&results).Error

	return
}

// GetBatchFromIsDst 批量查找
func (obj *_AllTimeZoneTransitionTypeMgr) GetBatchFromIsDst(isDsts []int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Is_DST` IN (?)", isDsts).Find(&results).Error

	return
}

// GetFromAbbreviation 通过Abbreviation获取内容
func (obj *_AllTimeZoneTransitionTypeMgr) GetFromAbbreviation(abbreviation string) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Abbreviation` = ?", abbreviation).Find(&results).Error

	return
}

// GetBatchFromAbbreviation 批量查找
func (obj *_AllTimeZoneTransitionTypeMgr) GetBatchFromAbbreviation(abbreviations []string) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Abbreviation` IN (?)", abbreviations).Find(&results).Error

	return
}

// GetFromVersion 通过Version获取内容
func (obj *_AllTimeZoneTransitionTypeMgr) GetFromVersion(version int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllTimeZoneTransitionTypeMgr) GetBatchFromVersion(versions []int64) (results []*AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTimeZoneTransitionTypeMgr) FetchByPrimaryKey(timeZoneID int64, transitionTypeID int64) (result AllTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTimeZoneTransitionType{}).Where("`Time_zone_id` = ? AND `Transition_type_id` = ?", timeZoneID, transitionTypeID).First(&result).Error

	return
}
