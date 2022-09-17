package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPartitionSstableMergeInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionSstableMergeInfoMgr open func
func AllVirtualPartitionSstableMergeInfoMgr(db *gorm.DB) *_AllVirtualPartitionSstableMergeInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionSstableMergeInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionSstableMergeInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_sstable_merge_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetTableName() string {
	return "__all_virtual_partition_sstable_merge_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) Reset() *_AllVirtualPartitionSstableMergeInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) Get() (result AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) Gets() (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithVersion version获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithMergeType merge_type获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMergeType(mergeType string) Option {
	return optionFunc(func(o *options) { o.query["merge_type"] = mergeType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithTableType table_type获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithMajorTableID major_table_id获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMajorTableID(majorTableID int64) Option {
	return optionFunc(func(o *options) { o.query["major_table_id"] = majorTableID })
}

// WithMergeStartTime merge_start_time获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMergeStartTime(mergeStartTime int64) Option {
	return optionFunc(func(o *options) { o.query["merge_start_time"] = mergeStartTime })
}

// WithMergeFinishTime merge_finish_time获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMergeFinishTime(mergeFinishTime int64) Option {
	return optionFunc(func(o *options) { o.query["merge_finish_time"] = mergeFinishTime })
}

// WithMergeCostTime merge_cost_time获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMergeCostTime(mergeCostTime int64) Option {
	return optionFunc(func(o *options) { o.query["merge_cost_time"] = mergeCostTime })
}

// WithEstimateCostTime estimate_cost_time获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithEstimateCostTime(estimateCostTime int64) Option {
	return optionFunc(func(o *options) { o.query["estimate_cost_time"] = estimateCostTime })
}

// WithOccupySize occupy_size获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithOccupySize(occupySize int64) Option {
	return optionFunc(func(o *options) { o.query["occupy_size"] = occupySize })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithUseOldMacroBlockCount use_old_macro_block_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithUseOldMacroBlockCount(useOldMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["use_old_macro_block_count"] = useOldMacroBlockCount })
}

// WithBuildBloomfilterCount build_bloomfilter_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithBuildBloomfilterCount(buildBloomfilterCount int64) Option {
	return optionFunc(func(o *options) { o.query["build_bloomfilter_count"] = buildBloomfilterCount })
}

// WithTotalRowCount total_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithTotalRowCount(totalRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_row_count"] = totalRowCount })
}

// WithDeleteRowCount delete_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithDeleteRowCount(deleteRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["delete_row_count"] = deleteRowCount })
}

// WithInsertRowCount insert_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithInsertRowCount(insertRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["insert_row_count"] = insertRowCount })
}

// WithUpdateRowCount update_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithUpdateRowCount(updateRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["update_row_count"] = updateRowCount })
}

// WithBaseRowCount base_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithBaseRowCount(baseRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["base_row_count"] = baseRowCount })
}

// WithUseBaseRowCount use_base_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithUseBaseRowCount(useBaseRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["use_base_row_count"] = useBaseRowCount })
}

// WithMemtableRowCount memtable_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMemtableRowCount(memtableRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["memtable_row_count"] = memtableRowCount })
}

// WithPurgedRowCount purged_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithPurgedRowCount(purgedRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["purged_row_count"] = purgedRowCount })
}

// WithOutputRowCount output_row_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithOutputRowCount(outputRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["output_row_count"] = outputRowCount })
}

// WithMergeLevel merge_level获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMergeLevel(mergeLevel int64) Option {
	return optionFunc(func(o *options) { o.query["merge_level"] = mergeLevel })
}

// WithRewriteMacroOldMicroBlockCount rewrite_macro_old_micro_block_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithRewriteMacroOldMicroBlockCount(rewriteMacroOldMicroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["rewrite_macro_old_micro_block_count"] = rewriteMacroOldMicroBlockCount })
}

// WithRewriteMacroTotalMicroBlockCount rewrite_macro_total_micro_block_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithRewriteMacroTotalMicroBlockCount(rewriteMacroTotalMicroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["rewrite_macro_total_micro_block_count"] = rewriteMacroTotalMicroBlockCount })
}

// WithTotalChildTask total_child_task获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithTotalChildTask(totalChildTask int64) Option {
	return optionFunc(func(o *options) { o.query["total_child_task"] = totalChildTask })
}

// WithFinishChildTask finish_child_task获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithFinishChildTask(finishChildTask int64) Option {
	return optionFunc(func(o *options) { o.query["finish_child_task"] = finishChildTask })
}

// WithStepMergePercentage step_merge_percentage获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithStepMergePercentage(stepMergePercentage int64) Option {
	return optionFunc(func(o *options) { o.query["step_merge_percentage"] = stepMergePercentage })
}

// WithMergePercentage merge_percentage获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithMergePercentage(mergePercentage int64) Option {
	return optionFunc(func(o *options) { o.query["merge_percentage"] = mergePercentage })
}

// WithErrorCode error_code获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithErrorCode(errorCode int64) Option {
	return optionFunc(func(o *options) { o.query["error_code"] = errorCode })
}

// WithTableCount table_count获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) WithTableCount(tableCount int64) Option {
	return optionFunc(func(o *options) { o.query["table_count"] = tableCount })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetByOption(opts ...Option) (result AllVirtualPartitionSstableMergeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionSstableMergeInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where(options.query)
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

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromVersion(version string) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromVersion(versions []string) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromMergeType 通过merge_type获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMergeType(mergeType string) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_type` = ?", mergeType).Find(&results).Error

	return
}

// GetBatchFromMergeType 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMergeType(mergeTypes []string) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_type` IN (?)", mergeTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromTableType(tableType int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromTableType(tableTypes []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

// GetFromMajorTableID 通过major_table_id获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMajorTableID(majorTableID int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`major_table_id` = ?", majorTableID).Find(&results).Error

	return
}

// GetBatchFromMajorTableID 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMajorTableID(majorTableIDs []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`major_table_id` IN (?)", majorTableIDs).Find(&results).Error

	return
}

// GetFromMergeStartTime 通过merge_start_time获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMergeStartTime(mergeStartTime int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_start_time` = ?", mergeStartTime).Find(&results).Error

	return
}

// GetBatchFromMergeStartTime 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMergeStartTime(mergeStartTimes []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_start_time` IN (?)", mergeStartTimes).Find(&results).Error

	return
}

// GetFromMergeFinishTime 通过merge_finish_time获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMergeFinishTime(mergeFinishTime int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_finish_time` = ?", mergeFinishTime).Find(&results).Error

	return
}

// GetBatchFromMergeFinishTime 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMergeFinishTime(mergeFinishTimes []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_finish_time` IN (?)", mergeFinishTimes).Find(&results).Error

	return
}

// GetFromMergeCostTime 通过merge_cost_time获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMergeCostTime(mergeCostTime int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_cost_time` = ?", mergeCostTime).Find(&results).Error

	return
}

// GetBatchFromMergeCostTime 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMergeCostTime(mergeCostTimes []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_cost_time` IN (?)", mergeCostTimes).Find(&results).Error

	return
}

// GetFromEstimateCostTime 通过estimate_cost_time获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromEstimateCostTime(estimateCostTime int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`estimate_cost_time` = ?", estimateCostTime).Find(&results).Error

	return
}

// GetBatchFromEstimateCostTime 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromEstimateCostTime(estimateCostTimes []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`estimate_cost_time` IN (?)", estimateCostTimes).Find(&results).Error

	return
}

// GetFromOccupySize 通过occupy_size获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromOccupySize(occupySize int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`occupy_size` = ?", occupySize).Find(&results).Error

	return
}

// GetBatchFromOccupySize 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromOccupySize(occupySizes []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`occupy_size` IN (?)", occupySizes).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromUseOldMacroBlockCount 通过use_old_macro_block_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromUseOldMacroBlockCount(useOldMacroBlockCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`use_old_macro_block_count` = ?", useOldMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromUseOldMacroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromUseOldMacroBlockCount(useOldMacroBlockCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`use_old_macro_block_count` IN (?)", useOldMacroBlockCounts).Find(&results).Error

	return
}

// GetFromBuildBloomfilterCount 通过build_bloomfilter_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromBuildBloomfilterCount(buildBloomfilterCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`build_bloomfilter_count` = ?", buildBloomfilterCount).Find(&results).Error

	return
}

// GetBatchFromBuildBloomfilterCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromBuildBloomfilterCount(buildBloomfilterCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`build_bloomfilter_count` IN (?)", buildBloomfilterCounts).Find(&results).Error

	return
}

// GetFromTotalRowCount 通过total_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromTotalRowCount(totalRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`total_row_count` = ?", totalRowCount).Find(&results).Error

	return
}

// GetBatchFromTotalRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromTotalRowCount(totalRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`total_row_count` IN (?)", totalRowCounts).Find(&results).Error

	return
}

// GetFromDeleteRowCount 通过delete_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromDeleteRowCount(deleteRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`delete_row_count` = ?", deleteRowCount).Find(&results).Error

	return
}

// GetBatchFromDeleteRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromDeleteRowCount(deleteRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`delete_row_count` IN (?)", deleteRowCounts).Find(&results).Error

	return
}

// GetFromInsertRowCount 通过insert_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromInsertRowCount(insertRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`insert_row_count` = ?", insertRowCount).Find(&results).Error

	return
}

// GetBatchFromInsertRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromInsertRowCount(insertRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`insert_row_count` IN (?)", insertRowCounts).Find(&results).Error

	return
}

// GetFromUpdateRowCount 通过update_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromUpdateRowCount(updateRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`update_row_count` = ?", updateRowCount).Find(&results).Error

	return
}

// GetBatchFromUpdateRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromUpdateRowCount(updateRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`update_row_count` IN (?)", updateRowCounts).Find(&results).Error

	return
}

// GetFromBaseRowCount 通过base_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromBaseRowCount(baseRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`base_row_count` = ?", baseRowCount).Find(&results).Error

	return
}

// GetBatchFromBaseRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromBaseRowCount(baseRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`base_row_count` IN (?)", baseRowCounts).Find(&results).Error

	return
}

// GetFromUseBaseRowCount 通过use_base_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromUseBaseRowCount(useBaseRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`use_base_row_count` = ?", useBaseRowCount).Find(&results).Error

	return
}

// GetBatchFromUseBaseRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromUseBaseRowCount(useBaseRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`use_base_row_count` IN (?)", useBaseRowCounts).Find(&results).Error

	return
}

// GetFromMemtableRowCount 通过memtable_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMemtableRowCount(memtableRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`memtable_row_count` = ?", memtableRowCount).Find(&results).Error

	return
}

// GetBatchFromMemtableRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMemtableRowCount(memtableRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`memtable_row_count` IN (?)", memtableRowCounts).Find(&results).Error

	return
}

// GetFromPurgedRowCount 通过purged_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromPurgedRowCount(purgedRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`purged_row_count` = ?", purgedRowCount).Find(&results).Error

	return
}

// GetBatchFromPurgedRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromPurgedRowCount(purgedRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`purged_row_count` IN (?)", purgedRowCounts).Find(&results).Error

	return
}

// GetFromOutputRowCount 通过output_row_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromOutputRowCount(outputRowCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`output_row_count` = ?", outputRowCount).Find(&results).Error

	return
}

// GetBatchFromOutputRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromOutputRowCount(outputRowCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`output_row_count` IN (?)", outputRowCounts).Find(&results).Error

	return
}

// GetFromMergeLevel 通过merge_level获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMergeLevel(mergeLevel int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_level` = ?", mergeLevel).Find(&results).Error

	return
}

// GetBatchFromMergeLevel 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMergeLevel(mergeLevels []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_level` IN (?)", mergeLevels).Find(&results).Error

	return
}

// GetFromRewriteMacroOldMicroBlockCount 通过rewrite_macro_old_micro_block_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromRewriteMacroOldMicroBlockCount(rewriteMacroOldMicroBlockCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`rewrite_macro_old_micro_block_count` = ?", rewriteMacroOldMicroBlockCount).Find(&results).Error

	return
}

// GetBatchFromRewriteMacroOldMicroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromRewriteMacroOldMicroBlockCount(rewriteMacroOldMicroBlockCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`rewrite_macro_old_micro_block_count` IN (?)", rewriteMacroOldMicroBlockCounts).Find(&results).Error

	return
}

// GetFromRewriteMacroTotalMicroBlockCount 通过rewrite_macro_total_micro_block_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromRewriteMacroTotalMicroBlockCount(rewriteMacroTotalMicroBlockCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`rewrite_macro_total_micro_block_count` = ?", rewriteMacroTotalMicroBlockCount).Find(&results).Error

	return
}

// GetBatchFromRewriteMacroTotalMicroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromRewriteMacroTotalMicroBlockCount(rewriteMacroTotalMicroBlockCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`rewrite_macro_total_micro_block_count` IN (?)", rewriteMacroTotalMicroBlockCounts).Find(&results).Error

	return
}

// GetFromTotalChildTask 通过total_child_task获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromTotalChildTask(totalChildTask int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`total_child_task` = ?", totalChildTask).Find(&results).Error

	return
}

// GetBatchFromTotalChildTask 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromTotalChildTask(totalChildTasks []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`total_child_task` IN (?)", totalChildTasks).Find(&results).Error

	return
}

// GetFromFinishChildTask 通过finish_child_task获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromFinishChildTask(finishChildTask int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`finish_child_task` = ?", finishChildTask).Find(&results).Error

	return
}

// GetBatchFromFinishChildTask 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromFinishChildTask(finishChildTasks []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`finish_child_task` IN (?)", finishChildTasks).Find(&results).Error

	return
}

// GetFromStepMergePercentage 通过step_merge_percentage获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromStepMergePercentage(stepMergePercentage int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`step_merge_percentage` = ?", stepMergePercentage).Find(&results).Error

	return
}

// GetBatchFromStepMergePercentage 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromStepMergePercentage(stepMergePercentages []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`step_merge_percentage` IN (?)", stepMergePercentages).Find(&results).Error

	return
}

// GetFromMergePercentage 通过merge_percentage获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromMergePercentage(mergePercentage int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_percentage` = ?", mergePercentage).Find(&results).Error

	return
}

// GetBatchFromMergePercentage 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromMergePercentage(mergePercentages []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`merge_percentage` IN (?)", mergePercentages).Find(&results).Error

	return
}

// GetFromErrorCode 通过error_code获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromErrorCode(errorCode int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`error_code` = ?", errorCode).Find(&results).Error

	return
}

// GetBatchFromErrorCode 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromErrorCode(errorCodes []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`error_code` IN (?)", errorCodes).Find(&results).Error

	return
}

// GetFromTableCount 通过table_count获取内容
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetFromTableCount(tableCount int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`table_count` = ?", tableCount).Find(&results).Error

	return
}

// GetBatchFromTableCount 批量查找
func (obj *_AllVirtualPartitionSstableMergeInfoMgr) GetBatchFromTableCount(tableCounts []int64) (results []*AllVirtualPartitionSstableMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMergeInfo{}).Where("`table_count` IN (?)", tableCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
