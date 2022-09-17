package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantRestorePgInfoMgr struct {
	*_BaseMgr
}

// AllTenantRestorePgInfoMgr open func
func AllTenantRestorePgInfoMgr(db *gorm.DB) *_AllTenantRestorePgInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantRestorePgInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantRestorePgInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_restore_pg_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantRestorePgInfoMgr) GetTableName() string {
	return "__all_tenant_restore_pg_info"
}

// Reset 重置gorm会话
func (obj *_AllTenantRestorePgInfoMgr) Reset() *_AllTenantRestorePgInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantRestorePgInfoMgr) Get() (result AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantRestorePgInfoMgr) Gets() (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantRestorePgInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantRestorePgInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantRestorePgInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTableID table_id获取
func (obj *_AllTenantRestorePgInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllTenantRestorePgInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllTenantRestorePgInfoMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllTenantRestorePgInfoMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithPartitionCount partition_count获取
func (obj *_AllTenantRestorePgInfoMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllTenantRestorePgInfoMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithRestoreInfo restore_info获取
func (obj *_AllTenantRestorePgInfoMgr) WithRestoreInfo(restoreInfo string) Option {
	return optionFunc(func(o *options) { o.query["restore_info"] = restoreInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantRestorePgInfoMgr) GetByOption(opts ...Option) (result AllTenantRestorePgInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantRestorePgInfoMgr) GetByOptions(opts ...Option) (results []*AllTenantRestorePgInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantRestorePgInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantRestorePgInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where(options.query)
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
func (obj *_AllTenantRestorePgInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromTableID(tableID int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromPartitionID(partitionID int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromPartitionCount(partitionCount int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromRestoreInfo 通过restore_info获取内容
func (obj *_AllTenantRestorePgInfoMgr) GetFromRestoreInfo(restoreInfo string) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`restore_info` = ?", restoreInfo).Find(&results).Error

	return
}

// GetBatchFromRestoreInfo 批量查找
func (obj *_AllTenantRestorePgInfoMgr) GetBatchFromRestoreInfo(restoreInfos []string) (results []*AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`restore_info` IN (?)", restoreInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantRestorePgInfoMgr) FetchByPrimaryKey(tableID int64, partitionID int64) (result AllTenantRestorePgInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantRestorePgInfo{}).Where("`table_id` = ? AND `partition_id` = ?", tableID, partitionID).First(&result).Error

	return
}
