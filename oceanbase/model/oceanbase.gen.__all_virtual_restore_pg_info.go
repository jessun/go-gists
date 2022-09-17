package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualRestorePgInfoMgr struct {
	*_BaseMgr
}

// AllVirtualRestorePgInfoMgr open func
func AllVirtualRestorePgInfoMgr(db *gorm.DB) *_AllVirtualRestorePgInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRestorePgInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRestorePgInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_restore_pg_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRestorePgInfoMgr) GetTableName() string {
	return "__all_virtual_restore_pg_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRestorePgInfoMgr) Reset() *_AllVirtualRestorePgInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRestorePgInfoMgr) Get() (result AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRestorePgInfoMgr) Gets() (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRestorePgInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRestorePgInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualRestorePgInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualRestorePgInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualRestorePgInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualRestorePgInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllVirtualRestorePgInfoMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllVirtualRestorePgInfoMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithPartitionCount partition_count获取
func (obj *_AllVirtualRestorePgInfoMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllVirtualRestorePgInfoMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithRestoreInfo restore_info获取
func (obj *_AllVirtualRestorePgInfoMgr) WithRestoreInfo(restoreInfo string) Option {
	return optionFunc(func(o *options) { o.query["restore_info"] = restoreInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRestorePgInfoMgr) GetByOption(opts ...Option) (result AllVirtualRestorePgInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRestorePgInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualRestorePgInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRestorePgInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRestorePgInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where(options.query)
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
func (obj *_AllVirtualRestorePgInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromPartitionCount(partitionCount int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromRestoreInfo 通过restore_info获取内容
func (obj *_AllVirtualRestorePgInfoMgr) GetFromRestoreInfo(restoreInfo string) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`restore_info` = ?", restoreInfo).Find(&results).Error

	return
}

// GetBatchFromRestoreInfo 批量查找
func (obj *_AllVirtualRestorePgInfoMgr) GetBatchFromRestoreInfo(restoreInfos []string) (results []*AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`restore_info` IN (?)", restoreInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualRestorePgInfoMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64) (result AllVirtualRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRestorePgInfo{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, tableID, partitionID).First(&result).Error

	return
}
