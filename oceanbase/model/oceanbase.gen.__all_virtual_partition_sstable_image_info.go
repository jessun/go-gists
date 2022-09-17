package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualPartitionSstableImageInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionSstableImageInfoMgr open func
func AllVirtualPartitionSstableImageInfoMgr(db *gorm.DB) *_AllVirtualPartitionSstableImageInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionSstableImageInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionSstableImageInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_sstable_image_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetTableName() string {
	return "__all_virtual_partition_sstable_image_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionSstableImageInfoMgr) Reset() *_AllVirtualPartitionSstableImageInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) Get() (result AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionSstableImageInfoMgr) Gets() (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionSstableImageInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithZone zone获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithMajorVersion major_version获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithMajorVersion(majorVersion int64) Option {
	return optionFunc(func(o *options) { o.query["major_version"] = majorVersion })
}

// WithMinVersion min_version获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithMinVersion(minVersion int64) Option {
	return optionFunc(func(o *options) { o.query["min_version"] = minVersion })
}

// WithSsStoreCount ss_store_count获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithSsStoreCount(ssStoreCount int64) Option {
	return optionFunc(func(o *options) { o.query["ss_store_count"] = ssStoreCount })
}

// WithMergedSsStoreCount merged_ss_store_count获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithMergedSsStoreCount(mergedSsStoreCount int64) Option {
	return optionFunc(func(o *options) { o.query["merged_ss_store_count"] = mergedSsStoreCount })
}

// WithModifiedSsStoreCount modified_ss_store_count获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithModifiedSsStoreCount(modifiedSsStoreCount int64) Option {
	return optionFunc(func(o *options) { o.query["modified_ss_store_count"] = modifiedSsStoreCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithUseOldMacroBlockCount use_old_macro_block_count获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithUseOldMacroBlockCount(useOldMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["use_old_macro_block_count"] = useOldMacroBlockCount })
}

// WithMergeStartTime merge_start_time获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithMergeStartTime(mergeStartTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["merge_start_time"] = mergeStartTime })
}

// WithMergeFinishTime merge_finish_time获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithMergeFinishTime(mergeFinishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["merge_finish_time"] = mergeFinishTime })
}

// WithMergeProcess merge_process获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithMergeProcess(mergeProcess int64) Option {
	return optionFunc(func(o *options) { o.query["merge_process"] = mergeProcess })
}

// WithRewriteMacroOldMicroBlockCount rewrite_macro_old_micro_block_count获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithRewriteMacroOldMicroBlockCount(rewriteMacroOldMicroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["rewrite_macro_old_micro_block_count"] = rewriteMacroOldMicroBlockCount })
}

// WithRewriteMacroTotalMicroBlockCount rewrite_macro_total_micro_block_count获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) WithRewriteMacroTotalMicroBlockCount(rewriteMacroTotalMicroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["rewrite_macro_total_micro_block_count"] = rewriteMacroTotalMicroBlockCount })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetByOption(opts ...Option) (result AllVirtualPartitionSstableImageInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionSstableImageInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionSstableImageInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where(options.query)
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

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromZone(zone string) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromZone(zones []string) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromMajorVersion 通过major_version获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromMajorVersion(majorVersion int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`major_version` = ?", majorVersion).Find(&results).Error

	return
}

// GetBatchFromMajorVersion 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromMajorVersion(majorVersions []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`major_version` IN (?)", majorVersions).Find(&results).Error

	return
}

// GetFromMinVersion 通过min_version获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromMinVersion(minVersion int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`min_version` = ?", minVersion).Find(&results).Error

	return
}

// GetBatchFromMinVersion 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromMinVersion(minVersions []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`min_version` IN (?)", minVersions).Find(&results).Error

	return
}

// GetFromSsStoreCount 通过ss_store_count获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromSsStoreCount(ssStoreCount int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`ss_store_count` = ?", ssStoreCount).Find(&results).Error

	return
}

// GetBatchFromSsStoreCount 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromSsStoreCount(ssStoreCounts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`ss_store_count` IN (?)", ssStoreCounts).Find(&results).Error

	return
}

// GetFromMergedSsStoreCount 通过merged_ss_store_count获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromMergedSsStoreCount(mergedSsStoreCount int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merged_ss_store_count` = ?", mergedSsStoreCount).Find(&results).Error

	return
}

// GetBatchFromMergedSsStoreCount 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromMergedSsStoreCount(mergedSsStoreCounts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merged_ss_store_count` IN (?)", mergedSsStoreCounts).Find(&results).Error

	return
}

// GetFromModifiedSsStoreCount 通过modified_ss_store_count获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromModifiedSsStoreCount(modifiedSsStoreCount int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`modified_ss_store_count` = ?", modifiedSsStoreCount).Find(&results).Error

	return
}

// GetBatchFromModifiedSsStoreCount 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromModifiedSsStoreCount(modifiedSsStoreCounts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`modified_ss_store_count` IN (?)", modifiedSsStoreCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromUseOldMacroBlockCount 通过use_old_macro_block_count获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromUseOldMacroBlockCount(useOldMacroBlockCount int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`use_old_macro_block_count` = ?", useOldMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromUseOldMacroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromUseOldMacroBlockCount(useOldMacroBlockCounts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`use_old_macro_block_count` IN (?)", useOldMacroBlockCounts).Find(&results).Error

	return
}

// GetFromMergeStartTime 通过merge_start_time获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromMergeStartTime(mergeStartTime time.Time) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merge_start_time` = ?", mergeStartTime).Find(&results).Error

	return
}

// GetBatchFromMergeStartTime 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromMergeStartTime(mergeStartTimes []time.Time) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merge_start_time` IN (?)", mergeStartTimes).Find(&results).Error

	return
}

// GetFromMergeFinishTime 通过merge_finish_time获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromMergeFinishTime(mergeFinishTime time.Time) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merge_finish_time` = ?", mergeFinishTime).Find(&results).Error

	return
}

// GetBatchFromMergeFinishTime 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromMergeFinishTime(mergeFinishTimes []time.Time) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merge_finish_time` IN (?)", mergeFinishTimes).Find(&results).Error

	return
}

// GetFromMergeProcess 通过merge_process获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromMergeProcess(mergeProcess int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merge_process` = ?", mergeProcess).Find(&results).Error

	return
}

// GetBatchFromMergeProcess 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromMergeProcess(mergeProcesss []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`merge_process` IN (?)", mergeProcesss).Find(&results).Error

	return
}

// GetFromRewriteMacroOldMicroBlockCount 通过rewrite_macro_old_micro_block_count获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromRewriteMacroOldMicroBlockCount(rewriteMacroOldMicroBlockCount int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`rewrite_macro_old_micro_block_count` = ?", rewriteMacroOldMicroBlockCount).Find(&results).Error

	return
}

// GetBatchFromRewriteMacroOldMicroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromRewriteMacroOldMicroBlockCount(rewriteMacroOldMicroBlockCounts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`rewrite_macro_old_micro_block_count` IN (?)", rewriteMacroOldMicroBlockCounts).Find(&results).Error

	return
}

// GetFromRewriteMacroTotalMicroBlockCount 通过rewrite_macro_total_micro_block_count获取内容
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetFromRewriteMacroTotalMicroBlockCount(rewriteMacroTotalMicroBlockCount int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`rewrite_macro_total_micro_block_count` = ?", rewriteMacroTotalMicroBlockCount).Find(&results).Error

	return
}

// GetBatchFromRewriteMacroTotalMicroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableImageInfoMgr) GetBatchFromRewriteMacroTotalMicroBlockCount(rewriteMacroTotalMicroBlockCounts []int64) (results []*AllVirtualPartitionSstableImageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableImageInfo{}).Where("`rewrite_macro_total_micro_block_count` IN (?)", rewriteMacroTotalMicroBlockCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
