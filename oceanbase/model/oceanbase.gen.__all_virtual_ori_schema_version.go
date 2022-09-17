package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualOriSchemaVersionMgr struct {
	*_BaseMgr
}

// AllVirtualOriSchemaVersionMgr open func
func AllVirtualOriSchemaVersionMgr(db *gorm.DB) *_AllVirtualOriSchemaVersionMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualOriSchemaVersionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualOriSchemaVersionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_ori_schema_version"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualOriSchemaVersionMgr) GetTableName() string {
	return "__all_virtual_ori_schema_version"
}

// Reset 重置gorm会话
func (obj *_AllVirtualOriSchemaVersionMgr) Reset() *_AllVirtualOriSchemaVersionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualOriSchemaVersionMgr) Get() (result AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualOriSchemaVersionMgr) Gets() (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualOriSchemaVersionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualOriSchemaVersionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualOriSchemaVersionMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualOriSchemaVersionMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualOriSchemaVersionMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithOriSchemaVersion ori_schema_version获取
func (obj *_AllVirtualOriSchemaVersionMgr) WithOriSchemaVersion(oriSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["ori_schema_version"] = oriSchemaVersion })
}

// WithBuildingSnapshot building_snapshot获取
func (obj *_AllVirtualOriSchemaVersionMgr) WithBuildingSnapshot(buildingSnapshot int64) Option {
	return optionFunc(func(o *options) { o.query["building_snapshot"] = buildingSnapshot })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualOriSchemaVersionMgr) GetByOption(opts ...Option) (result AllVirtualOriSchemaVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualOriSchemaVersionMgr) GetByOptions(opts ...Option) (results []*AllVirtualOriSchemaVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualOriSchemaVersionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualOriSchemaVersion, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where(options.query)
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
func (obj *_AllVirtualOriSchemaVersionMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualOriSchemaVersionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualOriSchemaVersionMgr) GetFromTableID(tableID int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualOriSchemaVersionMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualOriSchemaVersionMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualOriSchemaVersionMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualOriSchemaVersionMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualOriSchemaVersionMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromOriSchemaVersion 通过ori_schema_version获取内容
func (obj *_AllVirtualOriSchemaVersionMgr) GetFromOriSchemaVersion(oriSchemaVersion int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`ori_schema_version` = ?", oriSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromOriSchemaVersion 批量查找
func (obj *_AllVirtualOriSchemaVersionMgr) GetBatchFromOriSchemaVersion(oriSchemaVersions []int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`ori_schema_version` IN (?)", oriSchemaVersions).Find(&results).Error

	return
}

// GetFromBuildingSnapshot 通过building_snapshot获取内容
func (obj *_AllVirtualOriSchemaVersionMgr) GetFromBuildingSnapshot(buildingSnapshot int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`building_snapshot` = ?", buildingSnapshot).Find(&results).Error

	return
}

// GetBatchFromBuildingSnapshot 批量查找
func (obj *_AllVirtualOriSchemaVersionMgr) GetBatchFromBuildingSnapshot(buildingSnapshots []int64) (results []*AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`building_snapshot` IN (?)", buildingSnapshots).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualOriSchemaVersionMgr) FetchByPrimaryKey(tenantID int64, tableID int64) (result AllVirtualOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOriSchemaVersion{}).Where("`tenant_id` = ? AND `table_id` = ?", tenantID, tableID).First(&result).Error

	return
}
