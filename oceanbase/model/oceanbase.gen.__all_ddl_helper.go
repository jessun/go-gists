package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllDdlHelperMgr struct {
	*_BaseMgr
}

// AllDdlHelperMgr open func
func AllDdlHelperMgr(db *gorm.DB) *_AllDdlHelperMgr {
	if db == nil {
		panic(fmt.Errorf("AllDdlHelperMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDdlHelperMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_ddl_helper"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDdlHelperMgr) GetTableName() string {
	return "__all_ddl_helper"
}

// Reset 重置gorm会话
func (obj *_AllDdlHelperMgr) Reset() *_AllDdlHelperMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDdlHelperMgr) Get() (result AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDdlHelperMgr) Gets() (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDdlHelperMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDdlHelperMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDdlHelperMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDdlHelperMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllDdlHelperMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithSchemaID schema_id获取
func (obj *_AllDdlHelperMgr) WithSchemaID(schemaID int64) Option {
	return optionFunc(func(o *options) { o.query["schema_id"] = schemaID })
}

// WithDdlType ddl_type获取
func (obj *_AllDdlHelperMgr) WithDdlType(ddlType int64) Option {
	return optionFunc(func(o *options) { o.query["ddl_type"] = ddlType })
}

// GetByOption 功能选项模式获取
func (obj *_AllDdlHelperMgr) GetByOption(opts ...Option) (result AllDdlHelper, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDdlHelperMgr) GetByOptions(opts ...Option) (results []*AllDdlHelper, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDdlHelperMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDdlHelper, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where(options.query)
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
func (obj *_AllDdlHelperMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDdlHelperMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDdlHelperMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDdlHelperMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDdlHelperMgr) GetFromTenantID(tenantID int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDdlHelperMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllDdlHelperMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllDdlHelperMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromSchemaID 通过schema_id获取内容
func (obj *_AllDdlHelperMgr) GetFromSchemaID(schemaID int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`schema_id` = ?", schemaID).Find(&results).Error

	return
}

// GetBatchFromSchemaID 批量查找
func (obj *_AllDdlHelperMgr) GetBatchFromSchemaID(schemaIDs []int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`schema_id` IN (?)", schemaIDs).Find(&results).Error

	return
}

// GetFromDdlType 通过ddl_type获取内容
func (obj *_AllDdlHelperMgr) GetFromDdlType(ddlType int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`ddl_type` = ?", ddlType).Find(&results).Error

	return
}

// GetBatchFromDdlType 批量查找
func (obj *_AllDdlHelperMgr) GetBatchFromDdlType(ddlTypes []int64) (results []*AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`ddl_type` IN (?)", ddlTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDdlHelperMgr) FetchByPrimaryKey(tenantID int64, schemaVersion int64) (result AllDdlHelper, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlHelper{}).Where("`tenant_id` = ? AND `schema_version` = ?", tenantID, schemaVersion).First(&result).Error

	return
}
