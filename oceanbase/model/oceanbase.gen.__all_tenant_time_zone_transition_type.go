package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllTenantTimeZoneTransitionTypeMgr struct {
	*_BaseMgr
}

// AllTenantTimeZoneTransitionTypeMgr open func
func AllTenantTimeZoneTransitionTypeMgr(db *gorm.DB) *_AllTenantTimeZoneTransitionTypeMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantTimeZoneTransitionTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantTimeZoneTransitionTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_time_zone_transition_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetTableName() string {
	return "__all_tenant_time_zone_transition_type"
}

// Reset 重置gorm会话
func (obj *_AllTenantTimeZoneTransitionTypeMgr) Reset() *_AllTenantTimeZoneTransitionTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) Get() (result AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantTimeZoneTransitionTypeMgr) Gets() (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantTimeZoneTransitionTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTimeZoneID time_zone_id获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) WithTimeZoneID(timeZoneID int64) Option {
	return optionFunc(func(o *options) { o.query["time_zone_id"] = timeZoneID })
}

// WithTransitionTypeID transition_type_id获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) WithTransitionTypeID(transitionTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["transition_type_id"] = transitionTypeID })
}

// WithOffset offset获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) WithOffset(offset int64) Option {
	return optionFunc(func(o *options) { o.query["offset"] = offset })
}

// WithIsDst is_dst获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) WithIsDst(isDst int64) Option {
	return optionFunc(func(o *options) { o.query["is_dst"] = isDst })
}

// WithAbbreviation abbreviation获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) WithAbbreviation(abbreviation string) Option {
	return optionFunc(func(o *options) { o.query["abbreviation"] = abbreviation })
}

// WithVersion version获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetByOption(opts ...Option) (result AllTenantTimeZoneTransitionType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetByOptions(opts ...Option) (results []*AllTenantTimeZoneTransitionType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantTimeZoneTransitionTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantTimeZoneTransitionType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where(options.query)
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
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetFromTenantID(tenantID int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTimeZoneID 通过time_zone_id获取内容
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetFromTimeZoneID(timeZoneID int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`time_zone_id` = ?", timeZoneID).Find(&results).Error

	return
}

// GetBatchFromTimeZoneID 批量查找
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetBatchFromTimeZoneID(timeZoneIDs []int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`time_zone_id` IN (?)", timeZoneIDs).Find(&results).Error

	return
}

// GetFromTransitionTypeID 通过transition_type_id获取内容
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetFromTransitionTypeID(transitionTypeID int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`transition_type_id` = ?", transitionTypeID).Find(&results).Error

	return
}

// GetBatchFromTransitionTypeID 批量查找
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetBatchFromTransitionTypeID(transitionTypeIDs []int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`transition_type_id` IN (?)", transitionTypeIDs).Find(&results).Error

	return
}

// GetFromOffset 通过offset获取内容
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetFromOffset(offset int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`offset` = ?", offset).Find(&results).Error

	return
}

// GetBatchFromOffset 批量查找
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetBatchFromOffset(offsets []int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`offset` IN (?)", offsets).Find(&results).Error

	return
}

// GetFromIsDst 通过is_dst获取内容
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetFromIsDst(isDst int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`is_dst` = ?", isDst).Find(&results).Error

	return
}

// GetBatchFromIsDst 批量查找
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetBatchFromIsDst(isDsts []int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`is_dst` IN (?)", isDsts).Find(&results).Error

	return
}

// GetFromAbbreviation 通过abbreviation获取内容
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetFromAbbreviation(abbreviation string) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`abbreviation` = ?", abbreviation).Find(&results).Error

	return
}

// GetBatchFromAbbreviation 批量查找
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetBatchFromAbbreviation(abbreviations []string) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`abbreviation` IN (?)", abbreviations).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetFromVersion(version int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllTenantTimeZoneTransitionTypeMgr) GetBatchFromVersion(versions []int64) (results []*AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantTimeZoneTransitionTypeMgr) FetchByPrimaryKey(tenantID int64, timeZoneID int64, transitionTypeID int64) (result AllTenantTimeZoneTransitionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTimeZoneTransitionType{}).Where("`tenant_id` = ? AND `time_zone_id` = ? AND `transition_type_id` = ?", tenantID, timeZoneID, transitionTypeID).First(&result).Error

	return
}
