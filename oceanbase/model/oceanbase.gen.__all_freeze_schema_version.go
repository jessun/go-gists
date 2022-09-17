package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllFreezeSchemaVersionMgr struct {
	*_BaseMgr
}

// AllFreezeSchemaVersionMgr open func
func AllFreezeSchemaVersionMgr(db *gorm.DB) *_AllFreezeSchemaVersionMgr {
	if db == nil {
		panic(fmt.Errorf("AllFreezeSchemaVersionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllFreezeSchemaVersionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_freeze_schema_version"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllFreezeSchemaVersionMgr) GetTableName() string {
	return "__all_freeze_schema_version"
}

// Reset 重置gorm会话
func (obj *_AllFreezeSchemaVersionMgr) Reset() *_AllFreezeSchemaVersionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllFreezeSchemaVersionMgr) Get() (result AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllFreezeSchemaVersionMgr) Gets() (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllFreezeSchemaVersionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllFreezeSchemaVersionMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllFreezeSchemaVersionMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithFrozenVersion frozen_version获取
func (obj *_AllFreezeSchemaVersionMgr) WithFrozenVersion(frozenVersion int64) Option {
	return optionFunc(func(o *options) { o.query["frozen_version"] = frozenVersion })
}

// WithTenantID tenant_id获取
func (obj *_AllFreezeSchemaVersionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllFreezeSchemaVersionMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllFreezeSchemaVersionMgr) GetByOption(opts ...Option) (result AllFreezeSchemaVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllFreezeSchemaVersionMgr) GetByOptions(opts ...Option) (results []*AllFreezeSchemaVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllFreezeSchemaVersionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllFreezeSchemaVersion, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where(options.query)
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
func (obj *_AllFreezeSchemaVersionMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllFreezeSchemaVersionMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllFreezeSchemaVersionMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllFreezeSchemaVersionMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromFrozenVersion 通过frozen_version获取内容
func (obj *_AllFreezeSchemaVersionMgr) GetFromFrozenVersion(frozenVersion int64) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`frozen_version` = ?", frozenVersion).Find(&results).Error

	return
}

// GetBatchFromFrozenVersion 批量查找
func (obj *_AllFreezeSchemaVersionMgr) GetBatchFromFrozenVersion(frozenVersions []int64) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`frozen_version` IN (?)", frozenVersions).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllFreezeSchemaVersionMgr) GetFromTenantID(tenantID int64) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllFreezeSchemaVersionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllFreezeSchemaVersionMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllFreezeSchemaVersionMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllFreezeSchemaVersionMgr) FetchByPrimaryKey(frozenVersion int64, tenantID int64) (result AllFreezeSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFreezeSchemaVersion{}).Where("`frozen_version` = ? AND `tenant_id` = ?", frozenVersion, tenantID).First(&result).Error

	return
}
