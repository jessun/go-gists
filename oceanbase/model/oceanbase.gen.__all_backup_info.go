package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupInfoMgr struct {
	*_BaseMgr
}

// AllBackupInfoMgr open func
func AllBackupInfoMgr(db *gorm.DB) *_AllBackupInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupInfoMgr) GetTableName() string {
	return "__all_backup_info"
}

// Reset 重置gorm会话
func (obj *_AllBackupInfoMgr) Reset() *_AllBackupInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupInfoMgr) Get() (result AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupInfoMgr) Gets() (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithName name获取
func (obj *_AllBackupInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取
func (obj *_AllBackupInfoMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupInfoMgr) GetByOption(opts ...Option) (result AllBackupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupInfoMgr) GetByOptions(opts ...Option) (results []*AllBackupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where(options.query)
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
func (obj *_AllBackupInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllBackupInfoMgr) GetFromName(name string) (result AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllBackupInfoMgr) GetBatchFromName(names []string) (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllBackupInfoMgr) GetFromValue(value string) (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllBackupInfoMgr) GetBatchFromValue(values []string) (results []*AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupInfoMgr) FetchByPrimaryKey(name string) (result AllBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupInfo{}).Where("`name` = ?", name).First(&result).Error

	return
}
