package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantBackupInfoMgr struct {
	*_BaseMgr
}

// AllTenantBackupInfoMgr open func
func AllTenantBackupInfoMgr(db *gorm.DB) *_AllTenantBackupInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantBackupInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantBackupInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_backup_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantBackupInfoMgr) GetTableName() string {
	return "__all_tenant_backup_info"
}

// Reset 重置gorm会话
func (obj *_AllTenantBackupInfoMgr) Reset() *_AllTenantBackupInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantBackupInfoMgr) Get() (result AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantBackupInfoMgr) Gets() (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantBackupInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantBackupInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantBackupInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantBackupInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithName name获取
func (obj *_AllTenantBackupInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取
func (obj *_AllTenantBackupInfoMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantBackupInfoMgr) GetByOption(opts ...Option) (result AllTenantBackupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantBackupInfoMgr) GetByOptions(opts ...Option) (results []*AllTenantBackupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantBackupInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantBackupInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where(options.query)
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
func (obj *_AllTenantBackupInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantBackupInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantBackupInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantBackupInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantBackupInfoMgr) GetFromTenantID(tenantID int64) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantBackupInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllTenantBackupInfoMgr) GetFromName(name string) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllTenantBackupInfoMgr) GetBatchFromName(names []string) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllTenantBackupInfoMgr) GetFromValue(value string) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllTenantBackupInfoMgr) GetBatchFromValue(values []string) (results []*AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantBackupInfoMgr) FetchByPrimaryKey(tenantID int64, name string) (result AllTenantBackupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupInfo{}).Where("`tenant_id` = ? AND `name` = ?", tenantID, name).First(&result).Error

	return
}
