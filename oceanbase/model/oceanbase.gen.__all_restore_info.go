package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllRestoreInfoMgr struct {
	*_BaseMgr
}

// AllRestoreInfoMgr open func
func AllRestoreInfoMgr(db *gorm.DB) *_AllRestoreInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllRestoreInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRestoreInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_restore_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRestoreInfoMgr) GetTableName() string {
	return "__all_restore_info"
}

// Reset 重置gorm会话
func (obj *_AllRestoreInfoMgr) Reset() *_AllRestoreInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRestoreInfoMgr) Get() (result AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRestoreInfoMgr) Gets() (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRestoreInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRestoreInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllRestoreInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllRestoreInfoMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithName name获取
func (obj *_AllRestoreInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取
func (obj *_AllRestoreInfoMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_AllRestoreInfoMgr) GetByOption(opts ...Option) (result AllRestoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRestoreInfoMgr) GetByOptions(opts ...Option) (results []*AllRestoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRestoreInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRestoreInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where(options.query)
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
func (obj *_AllRestoreInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRestoreInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllRestoreInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllRestoreInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllRestoreInfoMgr) GetFromJobID(jobID int64) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllRestoreInfoMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllRestoreInfoMgr) GetFromName(name string) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllRestoreInfoMgr) GetBatchFromName(names []string) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllRestoreInfoMgr) GetFromValue(value string) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllRestoreInfoMgr) GetBatchFromValue(values []string) (results []*AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRestoreInfoMgr) FetchByPrimaryKey(jobID int64, name string) (result AllRestoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreInfo{}).Where("`job_id` = ? AND `name` = ?", jobID, name).First(&result).Error

	return
}
