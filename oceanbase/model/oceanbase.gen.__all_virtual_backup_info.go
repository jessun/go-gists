package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualBackupInfoMgr struct {
	*_BaseMgr
}

// AllVirtualBackupInfoMgr open func
func AllVirtualBackupInfoMgr(db *gorm.DB) *_AllVirtualBackupInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualBackupInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualBackupInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_backup_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualBackupInfoMgr) GetTableName() string {
	return "__all_virtual_backup_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualBackupInfoMgr) Reset() *_AllVirtualBackupInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualBackupInfoMgr) Get() (result AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualBackupInfoMgr) Gets() (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualBackupInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualBackupInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithName name获取
func (obj *_AllVirtualBackupInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualBackupInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualBackupInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithValue value获取
func (obj *_AllVirtualBackupInfoMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualBackupInfoMgr) GetByOption(opts ...Option) (result AllVirtualBackupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualBackupInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualBackupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualBackupInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualBackupInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where(options.query)
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
func (obj *_AllVirtualBackupInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualBackupInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualBackupInfoMgr) GetFromName(name string) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualBackupInfoMgr) GetBatchFromName(names []string) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualBackupInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualBackupInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualBackupInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualBackupInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualBackupInfoMgr) GetFromValue(value string) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualBackupInfoMgr) GetBatchFromValue(values []string) (results []*AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualBackupInfoMgr) FetchByPrimaryKey(tenantID int64, name string) (result AllVirtualBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupInfo{}).Where("`tenant_id` = ? AND `name` = ?", tenantID, name).First(&result).Error

	return
}
