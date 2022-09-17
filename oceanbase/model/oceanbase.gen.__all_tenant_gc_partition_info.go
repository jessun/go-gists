package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantGcPartitionInfoMgr struct {
	*_BaseMgr
}

// AllTenantGcPartitionInfoMgr open func
func AllTenantGcPartitionInfoMgr(db *gorm.DB) *_AllTenantGcPartitionInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantGcPartitionInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantGcPartitionInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_gc_partition_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantGcPartitionInfoMgr) GetTableName() string {
	return "__all_tenant_gc_partition_info"
}

// Reset 重置gorm会话
func (obj *_AllTenantGcPartitionInfoMgr) Reset() *_AllTenantGcPartitionInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantGcPartitionInfoMgr) Get() (result AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantGcPartitionInfoMgr) Gets() (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantGcPartitionInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantGcPartitionInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantGcPartitionInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantGcPartitionInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllTenantGcPartitionInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllTenantGcPartitionInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantGcPartitionInfoMgr) GetByOption(opts ...Option) (result AllTenantGcPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantGcPartitionInfoMgr) GetByOptions(opts ...Option) (results []*AllTenantGcPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantGcPartitionInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantGcPartitionInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where(options.query)
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
func (obj *_AllTenantGcPartitionInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantGcPartitionInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantGcPartitionInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantGcPartitionInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantGcPartitionInfoMgr) GetFromTenantID(tenantID int64) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantGcPartitionInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTenantGcPartitionInfoMgr) GetFromTableID(tableID int64) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTenantGcPartitionInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllTenantGcPartitionInfoMgr) GetFromPartitionID(partitionID int64) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllTenantGcPartitionInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantGcPartitionInfoMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64) (result AllTenantGcPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantGcPartitionInfo{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, tableID, partitionID).First(&result).Error

	return
}
