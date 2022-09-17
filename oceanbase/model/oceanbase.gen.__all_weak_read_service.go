package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllWeakReadServiceMgr struct {
	*_BaseMgr
}

// AllWeakReadServiceMgr open func
func AllWeakReadServiceMgr(db *gorm.DB) *_AllWeakReadServiceMgr {
	if db == nil {
		panic(fmt.Errorf("AllWeakReadServiceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllWeakReadServiceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_weak_read_service"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllWeakReadServiceMgr) GetTableName() string {
	return "__all_weak_read_service"
}

// Reset 重置gorm会话
func (obj *_AllWeakReadServiceMgr) Reset() *_AllWeakReadServiceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllWeakReadServiceMgr) Get() (result AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllWeakReadServiceMgr) Gets() (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllWeakReadServiceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllWeakReadServiceMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllWeakReadServiceMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithLevelID level_id获取
func (obj *_AllWeakReadServiceMgr) WithLevelID(levelID int64) Option {
	return optionFunc(func(o *options) { o.query["level_id"] = levelID })
}

// WithLevelValue level_value获取
func (obj *_AllWeakReadServiceMgr) WithLevelValue(levelValue string) Option {
	return optionFunc(func(o *options) { o.query["level_value"] = levelValue })
}

// WithLevelName level_name获取
func (obj *_AllWeakReadServiceMgr) WithLevelName(levelName string) Option {
	return optionFunc(func(o *options) { o.query["level_name"] = levelName })
}

// WithMinVersion min_version获取
func (obj *_AllWeakReadServiceMgr) WithMinVersion(minVersion int64) Option {
	return optionFunc(func(o *options) { o.query["min_version"] = minVersion })
}

// WithMaxVersion max_version获取
func (obj *_AllWeakReadServiceMgr) WithMaxVersion(maxVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_version"] = maxVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllWeakReadServiceMgr) GetByOption(opts ...Option) (result AllWeakReadService, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllWeakReadServiceMgr) GetByOptions(opts ...Option) (results []*AllWeakReadService, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllWeakReadServiceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllWeakReadService, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllWeakReadServiceMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllWeakReadServiceMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllWeakReadServiceMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllWeakReadServiceMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromLevelID 通过level_id获取内容
func (obj *_AllWeakReadServiceMgr) GetFromLevelID(levelID int64) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`level_id` = ?", levelID).Find(&results).Error

	return
}

// GetBatchFromLevelID 批量查找
func (obj *_AllWeakReadServiceMgr) GetBatchFromLevelID(levelIDs []int64) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`level_id` IN (?)", levelIDs).Find(&results).Error

	return
}

// GetFromLevelValue 通过level_value获取内容
func (obj *_AllWeakReadServiceMgr) GetFromLevelValue(levelValue string) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`level_value` = ?", levelValue).Find(&results).Error

	return
}

// GetBatchFromLevelValue 批量查找
func (obj *_AllWeakReadServiceMgr) GetBatchFromLevelValue(levelValues []string) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`level_value` IN (?)", levelValues).Find(&results).Error

	return
}

// GetFromLevelName 通过level_name获取内容
func (obj *_AllWeakReadServiceMgr) GetFromLevelName(levelName string) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`level_name` = ?", levelName).Find(&results).Error

	return
}

// GetBatchFromLevelName 批量查找
func (obj *_AllWeakReadServiceMgr) GetBatchFromLevelName(levelNames []string) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`level_name` IN (?)", levelNames).Find(&results).Error

	return
}

// GetFromMinVersion 通过min_version获取内容
func (obj *_AllWeakReadServiceMgr) GetFromMinVersion(minVersion int64) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`min_version` = ?", minVersion).Find(&results).Error

	return
}

// GetBatchFromMinVersion 批量查找
func (obj *_AllWeakReadServiceMgr) GetBatchFromMinVersion(minVersions []int64) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`min_version` IN (?)", minVersions).Find(&results).Error

	return
}

// GetFromMaxVersion 通过max_version获取内容
func (obj *_AllWeakReadServiceMgr) GetFromMaxVersion(maxVersion int64) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`max_version` = ?", maxVersion).Find(&results).Error

	return
}

// GetBatchFromMaxVersion 批量查找
func (obj *_AllWeakReadServiceMgr) GetBatchFromMaxVersion(maxVersions []int64) (results []*AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`max_version` IN (?)", maxVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllWeakReadServiceMgr) FetchByPrimaryKey(levelID int64, levelValue string) (result AllWeakReadService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllWeakReadService{}).Where("`level_id` = ? AND `level_value` = ?", levelID, levelValue).First(&result).Error

	return
}
