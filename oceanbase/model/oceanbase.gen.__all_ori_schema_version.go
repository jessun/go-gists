package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllOriSchemaVersionMgr struct {
	*_BaseMgr
}

// AllOriSchemaVersionMgr open func
func AllOriSchemaVersionMgr(db *gorm.DB) *_AllOriSchemaVersionMgr {
	if db == nil {
		panic(fmt.Errorf("AllOriSchemaVersionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllOriSchemaVersionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_ori_schema_version"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllOriSchemaVersionMgr) GetTableName() string {
	return "__all_ori_schema_version"
}

// Reset 重置gorm会话
func (obj *_AllOriSchemaVersionMgr) Reset() *_AllOriSchemaVersionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllOriSchemaVersionMgr) Get() (result AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllOriSchemaVersionMgr) Gets() (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllOriSchemaVersionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllOriSchemaVersionMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllOriSchemaVersionMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllOriSchemaVersionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllOriSchemaVersionMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithOriSchemaVersion ori_schema_version获取
func (obj *_AllOriSchemaVersionMgr) WithOriSchemaVersion(oriSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["ori_schema_version"] = oriSchemaVersion })
}

// WithBuildingSnapshot building_snapshot获取
func (obj *_AllOriSchemaVersionMgr) WithBuildingSnapshot(buildingSnapshot int64) Option {
	return optionFunc(func(o *options) { o.query["building_snapshot"] = buildingSnapshot })
}

// GetByOption 功能选项模式获取
func (obj *_AllOriSchemaVersionMgr) GetByOption(opts ...Option) (result AllOriSchemaVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllOriSchemaVersionMgr) GetByOptions(opts ...Option) (results []*AllOriSchemaVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllOriSchemaVersionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllOriSchemaVersion, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where(options.query)
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
func (obj *_AllOriSchemaVersionMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllOriSchemaVersionMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllOriSchemaVersionMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllOriSchemaVersionMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllOriSchemaVersionMgr) GetFromTenantID(tenantID int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllOriSchemaVersionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllOriSchemaVersionMgr) GetFromTableID(tableID int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllOriSchemaVersionMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromOriSchemaVersion 通过ori_schema_version获取内容
func (obj *_AllOriSchemaVersionMgr) GetFromOriSchemaVersion(oriSchemaVersion int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`ori_schema_version` = ?", oriSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromOriSchemaVersion 批量查找
func (obj *_AllOriSchemaVersionMgr) GetBatchFromOriSchemaVersion(oriSchemaVersions []int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`ori_schema_version` IN (?)", oriSchemaVersions).Find(&results).Error

	return
}

// GetFromBuildingSnapshot 通过building_snapshot获取内容
func (obj *_AllOriSchemaVersionMgr) GetFromBuildingSnapshot(buildingSnapshot int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`building_snapshot` = ?", buildingSnapshot).Find(&results).Error

	return
}

// GetBatchFromBuildingSnapshot 批量查找
func (obj *_AllOriSchemaVersionMgr) GetBatchFromBuildingSnapshot(buildingSnapshots []int64) (results []*AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`building_snapshot` IN (?)", buildingSnapshots).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllOriSchemaVersionMgr) FetchByPrimaryKey(tenantID int64, tableID int64) (result AllOriSchemaVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOriSchemaVersion{}).Where("`tenant_id` = ? AND `table_id` = ?", tenantID, tableID).First(&result).Error

	return
}
