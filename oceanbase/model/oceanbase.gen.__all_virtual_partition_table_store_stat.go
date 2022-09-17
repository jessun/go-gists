package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPartitionTableStoreStatMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionTableStoreStatMgr open func
func AllVirtualPartitionTableStoreStatMgr(db *gorm.DB) *_AllVirtualPartitionTableStoreStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionTableStoreStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionTableStoreStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_table_store_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetTableName() string {
	return "__all_virtual_partition_table_store_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionTableStoreStatMgr) Reset() *_AllVirtualPartitionTableStoreStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) Get() (result AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionTableStoreStatMgr) Gets() (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionTableStoreStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithRowCacheHitCount row_cache_hit_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithRowCacheHitCount(rowCacheHitCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_cache_hit_count"] = rowCacheHitCount })
}

// WithRowCacheMissCount row_cache_miss_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithRowCacheMissCount(rowCacheMissCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_cache_miss_count"] = rowCacheMissCount })
}

// WithRowCachePutCount row_cache_put_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithRowCachePutCount(rowCachePutCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_cache_put_count"] = rowCachePutCount })
}

// WithBfFilterCount bf_filter_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithBfFilterCount(bfFilterCount int64) Option {
	return optionFunc(func(o *options) { o.query["bf_filter_count"] = bfFilterCount })
}

// WithBfEmptyReadCount bf_empty_read_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithBfEmptyReadCount(bfEmptyReadCount int64) Option {
	return optionFunc(func(o *options) { o.query["bf_empty_read_count"] = bfEmptyReadCount })
}

// WithBfAccessCount bf_access_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithBfAccessCount(bfAccessCount int64) Option {
	return optionFunc(func(o *options) { o.query["bf_access_count"] = bfAccessCount })
}

// WithBlockCacheHitCount block_cache_hit_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithBlockCacheHitCount(blockCacheHitCount int64) Option {
	return optionFunc(func(o *options) { o.query["block_cache_hit_count"] = blockCacheHitCount })
}

// WithBlockCacheMissCount block_cache_miss_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithBlockCacheMissCount(blockCacheMissCount int64) Option {
	return optionFunc(func(o *options) { o.query["block_cache_miss_count"] = blockCacheMissCount })
}

// WithAccessRowCount access_row_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithAccessRowCount(accessRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["access_row_count"] = accessRowCount })
}

// WithOutputRowCount output_row_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithOutputRowCount(outputRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["output_row_count"] = outputRowCount })
}

// WithFuseRowCacheHitCount fuse_row_cache_hit_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithFuseRowCacheHitCount(fuseRowCacheHitCount int64) Option {
	return optionFunc(func(o *options) { o.query["fuse_row_cache_hit_count"] = fuseRowCacheHitCount })
}

// WithFuseRowCacheMissCount fuse_row_cache_miss_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithFuseRowCacheMissCount(fuseRowCacheMissCount int64) Option {
	return optionFunc(func(o *options) { o.query["fuse_row_cache_miss_count"] = fuseRowCacheMissCount })
}

// WithFuseRowCachePutCount fuse_row_cache_put_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithFuseRowCachePutCount(fuseRowCachePutCount int64) Option {
	return optionFunc(func(o *options) { o.query["fuse_row_cache_put_count"] = fuseRowCachePutCount })
}

// WithSingleGetCallCount single_get_call_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithSingleGetCallCount(singleGetCallCount int64) Option {
	return optionFunc(func(o *options) { o.query["single_get_call_count"] = singleGetCallCount })
}

// WithSingleGetOutputRowCount single_get_output_row_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithSingleGetOutputRowCount(singleGetOutputRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["single_get_output_row_count"] = singleGetOutputRowCount })
}

// WithMultiGetCallCount multi_get_call_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithMultiGetCallCount(multiGetCallCount int64) Option {
	return optionFunc(func(o *options) { o.query["multi_get_call_count"] = multiGetCallCount })
}

// WithMultiGetOutputRowCount multi_get_output_row_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithMultiGetOutputRowCount(multiGetOutputRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["multi_get_output_row_count"] = multiGetOutputRowCount })
}

// WithIndexBackCallCount index_back_call_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithIndexBackCallCount(indexBackCallCount int64) Option {
	return optionFunc(func(o *options) { o.query["index_back_call_count"] = indexBackCallCount })
}

// WithIndexBackOutputRowCount index_back_output_row_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithIndexBackOutputRowCount(indexBackOutputRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["index_back_output_row_count"] = indexBackOutputRowCount })
}

// WithSingleScanCallCount single_scan_call_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithSingleScanCallCount(singleScanCallCount int64) Option {
	return optionFunc(func(o *options) { o.query["single_scan_call_count"] = singleScanCallCount })
}

// WithSingleScanOutputRowCount single_scan_output_row_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithSingleScanOutputRowCount(singleScanOutputRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["single_scan_output_row_count"] = singleScanOutputRowCount })
}

// WithMultiScanCallCount multi_scan_call_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithMultiScanCallCount(multiScanCallCount int64) Option {
	return optionFunc(func(o *options) { o.query["multi_scan_call_count"] = multiScanCallCount })
}

// WithMultiScanOutputRowCount multi_scan_output_row_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithMultiScanOutputRowCount(multiScanOutputRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["multi_scan_output_row_count"] = multiScanOutputRowCount })
}

// WithExistRowEffectReadCount exist_row_effect_read_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithExistRowEffectReadCount(existRowEffectReadCount int64) Option {
	return optionFunc(func(o *options) { o.query["exist_row_effect_read_count"] = existRowEffectReadCount })
}

// WithExistRowEmptyReadCount exist_row_empty_read_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithExistRowEmptyReadCount(existRowEmptyReadCount int64) Option {
	return optionFunc(func(o *options) { o.query["exist_row_empty_read_count"] = existRowEmptyReadCount })
}

// WithGetRowEffectReadCount get_row_effect_read_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithGetRowEffectReadCount(getRowEffectReadCount int64) Option {
	return optionFunc(func(o *options) { o.query["get_row_effect_read_count"] = getRowEffectReadCount })
}

// WithGetRowEmptyReadCount get_row_empty_read_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithGetRowEmptyReadCount(getRowEmptyReadCount int64) Option {
	return optionFunc(func(o *options) { o.query["get_row_empty_read_count"] = getRowEmptyReadCount })
}

// WithScanRowEffectReadCount scan_row_effect_read_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithScanRowEffectReadCount(scanRowEffectReadCount int64) Option {
	return optionFunc(func(o *options) { o.query["scan_row_effect_read_count"] = scanRowEffectReadCount })
}

// WithScanRowEmptyReadCount scan_row_empty_read_count获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithScanRowEmptyReadCount(scanRowEmptyReadCount int64) Option {
	return optionFunc(func(o *options) { o.query["scan_row_empty_read_count"] = scanRowEmptyReadCount })
}

// WithRowkeyPrefixAccessInfo rowkey_prefix_access_info获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) WithRowkeyPrefixAccessInfo(rowkeyPrefixAccessInfo string) Option {
	return optionFunc(func(o *options) { o.query["rowkey_prefix_access_info"] = rowkeyPrefixAccessInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetByOption(opts ...Option) (result AllVirtualPartitionTableStoreStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionTableStoreStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionTableStoreStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionTableStoreStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromRowCacheHitCount 通过row_cache_hit_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromRowCacheHitCount(rowCacheHitCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`row_cache_hit_count` = ?", rowCacheHitCount).Find(&results).Error

	return
}

// GetBatchFromRowCacheHitCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromRowCacheHitCount(rowCacheHitCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`row_cache_hit_count` IN (?)", rowCacheHitCounts).Find(&results).Error

	return
}

// GetFromRowCacheMissCount 通过row_cache_miss_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromRowCacheMissCount(rowCacheMissCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`row_cache_miss_count` = ?", rowCacheMissCount).Find(&results).Error

	return
}

// GetBatchFromRowCacheMissCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromRowCacheMissCount(rowCacheMissCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`row_cache_miss_count` IN (?)", rowCacheMissCounts).Find(&results).Error

	return
}

// GetFromRowCachePutCount 通过row_cache_put_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromRowCachePutCount(rowCachePutCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`row_cache_put_count` = ?", rowCachePutCount).Find(&results).Error

	return
}

// GetBatchFromRowCachePutCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromRowCachePutCount(rowCachePutCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`row_cache_put_count` IN (?)", rowCachePutCounts).Find(&results).Error

	return
}

// GetFromBfFilterCount 通过bf_filter_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromBfFilterCount(bfFilterCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`bf_filter_count` = ?", bfFilterCount).Find(&results).Error

	return
}

// GetBatchFromBfFilterCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromBfFilterCount(bfFilterCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`bf_filter_count` IN (?)", bfFilterCounts).Find(&results).Error

	return
}

// GetFromBfEmptyReadCount 通过bf_empty_read_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromBfEmptyReadCount(bfEmptyReadCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`bf_empty_read_count` = ?", bfEmptyReadCount).Find(&results).Error

	return
}

// GetBatchFromBfEmptyReadCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromBfEmptyReadCount(bfEmptyReadCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`bf_empty_read_count` IN (?)", bfEmptyReadCounts).Find(&results).Error

	return
}

// GetFromBfAccessCount 通过bf_access_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromBfAccessCount(bfAccessCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`bf_access_count` = ?", bfAccessCount).Find(&results).Error

	return
}

// GetBatchFromBfAccessCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromBfAccessCount(bfAccessCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`bf_access_count` IN (?)", bfAccessCounts).Find(&results).Error

	return
}

// GetFromBlockCacheHitCount 通过block_cache_hit_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromBlockCacheHitCount(blockCacheHitCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`block_cache_hit_count` = ?", blockCacheHitCount).Find(&results).Error

	return
}

// GetBatchFromBlockCacheHitCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromBlockCacheHitCount(blockCacheHitCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`block_cache_hit_count` IN (?)", blockCacheHitCounts).Find(&results).Error

	return
}

// GetFromBlockCacheMissCount 通过block_cache_miss_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromBlockCacheMissCount(blockCacheMissCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`block_cache_miss_count` = ?", blockCacheMissCount).Find(&results).Error

	return
}

// GetBatchFromBlockCacheMissCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromBlockCacheMissCount(blockCacheMissCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`block_cache_miss_count` IN (?)", blockCacheMissCounts).Find(&results).Error

	return
}

// GetFromAccessRowCount 通过access_row_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromAccessRowCount(accessRowCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`access_row_count` = ?", accessRowCount).Find(&results).Error

	return
}

// GetBatchFromAccessRowCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromAccessRowCount(accessRowCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`access_row_count` IN (?)", accessRowCounts).Find(&results).Error

	return
}

// GetFromOutputRowCount 通过output_row_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromOutputRowCount(outputRowCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`output_row_count` = ?", outputRowCount).Find(&results).Error

	return
}

// GetBatchFromOutputRowCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromOutputRowCount(outputRowCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`output_row_count` IN (?)", outputRowCounts).Find(&results).Error

	return
}

// GetFromFuseRowCacheHitCount 通过fuse_row_cache_hit_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromFuseRowCacheHitCount(fuseRowCacheHitCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`fuse_row_cache_hit_count` = ?", fuseRowCacheHitCount).Find(&results).Error

	return
}

// GetBatchFromFuseRowCacheHitCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromFuseRowCacheHitCount(fuseRowCacheHitCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`fuse_row_cache_hit_count` IN (?)", fuseRowCacheHitCounts).Find(&results).Error

	return
}

// GetFromFuseRowCacheMissCount 通过fuse_row_cache_miss_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromFuseRowCacheMissCount(fuseRowCacheMissCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`fuse_row_cache_miss_count` = ?", fuseRowCacheMissCount).Find(&results).Error

	return
}

// GetBatchFromFuseRowCacheMissCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromFuseRowCacheMissCount(fuseRowCacheMissCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`fuse_row_cache_miss_count` IN (?)", fuseRowCacheMissCounts).Find(&results).Error

	return
}

// GetFromFuseRowCachePutCount 通过fuse_row_cache_put_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromFuseRowCachePutCount(fuseRowCachePutCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`fuse_row_cache_put_count` = ?", fuseRowCachePutCount).Find(&results).Error

	return
}

// GetBatchFromFuseRowCachePutCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromFuseRowCachePutCount(fuseRowCachePutCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`fuse_row_cache_put_count` IN (?)", fuseRowCachePutCounts).Find(&results).Error

	return
}

// GetFromSingleGetCallCount 通过single_get_call_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromSingleGetCallCount(singleGetCallCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_get_call_count` = ?", singleGetCallCount).Find(&results).Error

	return
}

// GetBatchFromSingleGetCallCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromSingleGetCallCount(singleGetCallCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_get_call_count` IN (?)", singleGetCallCounts).Find(&results).Error

	return
}

// GetFromSingleGetOutputRowCount 通过single_get_output_row_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromSingleGetOutputRowCount(singleGetOutputRowCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_get_output_row_count` = ?", singleGetOutputRowCount).Find(&results).Error

	return
}

// GetBatchFromSingleGetOutputRowCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromSingleGetOutputRowCount(singleGetOutputRowCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_get_output_row_count` IN (?)", singleGetOutputRowCounts).Find(&results).Error

	return
}

// GetFromMultiGetCallCount 通过multi_get_call_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromMultiGetCallCount(multiGetCallCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_get_call_count` = ?", multiGetCallCount).Find(&results).Error

	return
}

// GetBatchFromMultiGetCallCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromMultiGetCallCount(multiGetCallCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_get_call_count` IN (?)", multiGetCallCounts).Find(&results).Error

	return
}

// GetFromMultiGetOutputRowCount 通过multi_get_output_row_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromMultiGetOutputRowCount(multiGetOutputRowCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_get_output_row_count` = ?", multiGetOutputRowCount).Find(&results).Error

	return
}

// GetBatchFromMultiGetOutputRowCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromMultiGetOutputRowCount(multiGetOutputRowCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_get_output_row_count` IN (?)", multiGetOutputRowCounts).Find(&results).Error

	return
}

// GetFromIndexBackCallCount 通过index_back_call_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromIndexBackCallCount(indexBackCallCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`index_back_call_count` = ?", indexBackCallCount).Find(&results).Error

	return
}

// GetBatchFromIndexBackCallCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromIndexBackCallCount(indexBackCallCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`index_back_call_count` IN (?)", indexBackCallCounts).Find(&results).Error

	return
}

// GetFromIndexBackOutputRowCount 通过index_back_output_row_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromIndexBackOutputRowCount(indexBackOutputRowCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`index_back_output_row_count` = ?", indexBackOutputRowCount).Find(&results).Error

	return
}

// GetBatchFromIndexBackOutputRowCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromIndexBackOutputRowCount(indexBackOutputRowCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`index_back_output_row_count` IN (?)", indexBackOutputRowCounts).Find(&results).Error

	return
}

// GetFromSingleScanCallCount 通过single_scan_call_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromSingleScanCallCount(singleScanCallCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_scan_call_count` = ?", singleScanCallCount).Find(&results).Error

	return
}

// GetBatchFromSingleScanCallCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromSingleScanCallCount(singleScanCallCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_scan_call_count` IN (?)", singleScanCallCounts).Find(&results).Error

	return
}

// GetFromSingleScanOutputRowCount 通过single_scan_output_row_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromSingleScanOutputRowCount(singleScanOutputRowCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_scan_output_row_count` = ?", singleScanOutputRowCount).Find(&results).Error

	return
}

// GetBatchFromSingleScanOutputRowCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromSingleScanOutputRowCount(singleScanOutputRowCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`single_scan_output_row_count` IN (?)", singleScanOutputRowCounts).Find(&results).Error

	return
}

// GetFromMultiScanCallCount 通过multi_scan_call_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromMultiScanCallCount(multiScanCallCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_scan_call_count` = ?", multiScanCallCount).Find(&results).Error

	return
}

// GetBatchFromMultiScanCallCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromMultiScanCallCount(multiScanCallCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_scan_call_count` IN (?)", multiScanCallCounts).Find(&results).Error

	return
}

// GetFromMultiScanOutputRowCount 通过multi_scan_output_row_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromMultiScanOutputRowCount(multiScanOutputRowCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_scan_output_row_count` = ?", multiScanOutputRowCount).Find(&results).Error

	return
}

// GetBatchFromMultiScanOutputRowCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromMultiScanOutputRowCount(multiScanOutputRowCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`multi_scan_output_row_count` IN (?)", multiScanOutputRowCounts).Find(&results).Error

	return
}

// GetFromExistRowEffectReadCount 通过exist_row_effect_read_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromExistRowEffectReadCount(existRowEffectReadCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`exist_row_effect_read_count` = ?", existRowEffectReadCount).Find(&results).Error

	return
}

// GetBatchFromExistRowEffectReadCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromExistRowEffectReadCount(existRowEffectReadCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`exist_row_effect_read_count` IN (?)", existRowEffectReadCounts).Find(&results).Error

	return
}

// GetFromExistRowEmptyReadCount 通过exist_row_empty_read_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromExistRowEmptyReadCount(existRowEmptyReadCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`exist_row_empty_read_count` = ?", existRowEmptyReadCount).Find(&results).Error

	return
}

// GetBatchFromExistRowEmptyReadCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromExistRowEmptyReadCount(existRowEmptyReadCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`exist_row_empty_read_count` IN (?)", existRowEmptyReadCounts).Find(&results).Error

	return
}

// GetFromGetRowEffectReadCount 通过get_row_effect_read_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromGetRowEffectReadCount(getRowEffectReadCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`get_row_effect_read_count` = ?", getRowEffectReadCount).Find(&results).Error

	return
}

// GetBatchFromGetRowEffectReadCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromGetRowEffectReadCount(getRowEffectReadCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`get_row_effect_read_count` IN (?)", getRowEffectReadCounts).Find(&results).Error

	return
}

// GetFromGetRowEmptyReadCount 通过get_row_empty_read_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromGetRowEmptyReadCount(getRowEmptyReadCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`get_row_empty_read_count` = ?", getRowEmptyReadCount).Find(&results).Error

	return
}

// GetBatchFromGetRowEmptyReadCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromGetRowEmptyReadCount(getRowEmptyReadCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`get_row_empty_read_count` IN (?)", getRowEmptyReadCounts).Find(&results).Error

	return
}

// GetFromScanRowEffectReadCount 通过scan_row_effect_read_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromScanRowEffectReadCount(scanRowEffectReadCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`scan_row_effect_read_count` = ?", scanRowEffectReadCount).Find(&results).Error

	return
}

// GetBatchFromScanRowEffectReadCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromScanRowEffectReadCount(scanRowEffectReadCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`scan_row_effect_read_count` IN (?)", scanRowEffectReadCounts).Find(&results).Error

	return
}

// GetFromScanRowEmptyReadCount 通过scan_row_empty_read_count获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromScanRowEmptyReadCount(scanRowEmptyReadCount int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`scan_row_empty_read_count` = ?", scanRowEmptyReadCount).Find(&results).Error

	return
}

// GetBatchFromScanRowEmptyReadCount 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromScanRowEmptyReadCount(scanRowEmptyReadCounts []int64) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`scan_row_empty_read_count` IN (?)", scanRowEmptyReadCounts).Find(&results).Error

	return
}

// GetFromRowkeyPrefixAccessInfo 通过rowkey_prefix_access_info获取内容
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetFromRowkeyPrefixAccessInfo(rowkeyPrefixAccessInfo string) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`rowkey_prefix_access_info` = ?", rowkeyPrefixAccessInfo).Find(&results).Error

	return
}

// GetBatchFromRowkeyPrefixAccessInfo 批量查找
func (obj *_AllVirtualPartitionTableStoreStatMgr) GetBatchFromRowkeyPrefixAccessInfo(rowkeyPrefixAccessInfos []string) (results []*AllVirtualPartitionTableStoreStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTableStoreStat{}).Where("`rowkey_prefix_access_info` IN (?)", rowkeyPrefixAccessInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
