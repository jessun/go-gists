package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllIndexBuildStatMgr struct {
	*_BaseMgr
}

// AllIndexBuildStatMgr open func
func AllIndexBuildStatMgr(db *gorm.DB) *_AllIndexBuildStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllIndexBuildStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllIndexBuildStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_index_build_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllIndexBuildStatMgr) GetTableName() string {
	return "__all_index_build_stat"
}

// Reset 重置gorm会话
func (obj *_AllIndexBuildStatMgr) Reset() *_AllIndexBuildStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllIndexBuildStatMgr) Get() (result AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllIndexBuildStatMgr) Gets() (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllIndexBuildStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllIndexBuildStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllIndexBuildStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllIndexBuildStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDataTableID data_table_id获取
func (obj *_AllIndexBuildStatMgr) WithDataTableID(dataTableID int64) Option {
	return optionFunc(func(o *options) { o.query["data_table_id"] = dataTableID })
}

// WithIndexTableID index_table_id获取
func (obj *_AllIndexBuildStatMgr) WithIndexTableID(indexTableID int64) Option {
	return optionFunc(func(o *options) { o.query["index_table_id"] = indexTableID })
}

// WithStatus status获取
func (obj *_AllIndexBuildStatMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSnapshot snapshot获取
func (obj *_AllIndexBuildStatMgr) WithSnapshot(snapshot int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot"] = snapshot })
}

// WithSchemaVersion schema_version获取
func (obj *_AllIndexBuildStatMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllIndexBuildStatMgr) GetByOption(opts ...Option) (result AllIndexBuildStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllIndexBuildStatMgr) GetByOptions(opts ...Option) (results []*AllIndexBuildStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllIndexBuildStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllIndexBuildStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where(options.query)
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
func (obj *_AllIndexBuildStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllIndexBuildStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllIndexBuildStatMgr) GetFromTenantID(tenantID int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDataTableID 通过data_table_id获取内容
func (obj *_AllIndexBuildStatMgr) GetFromDataTableID(dataTableID int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}

// GetBatchFromDataTableID 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromDataTableID(dataTableIDs []int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`data_table_id` IN (?)", dataTableIDs).Find(&results).Error

	return
}

// GetFromIndexTableID 通过index_table_id获取内容
func (obj *_AllIndexBuildStatMgr) GetFromIndexTableID(indexTableID int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`index_table_id` = ?", indexTableID).Find(&results).Error

	return
}

// GetBatchFromIndexTableID 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromIndexTableID(indexTableIDs []int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`index_table_id` IN (?)", indexTableIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllIndexBuildStatMgr) GetFromStatus(status int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromStatus(statuss []int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSnapshot 通过snapshot获取内容
func (obj *_AllIndexBuildStatMgr) GetFromSnapshot(snapshot int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`snapshot` = ?", snapshot).Find(&results).Error

	return
}

// GetBatchFromSnapshot 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromSnapshot(snapshots []int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`snapshot` IN (?)", snapshots).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllIndexBuildStatMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllIndexBuildStatMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllIndexBuildStatMgr) FetchByPrimaryKey(tenantID int64, dataTableID int64, indexTableID int64) (result AllIndexBuildStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexBuildStat{}).Where("`tenant_id` = ? AND `data_table_id` = ? AND `index_table_id` = ?", tenantID, dataTableID, indexTableID).First(&result).Error

	return
}
