package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllRestoreProgressMgr struct {
	*_BaseMgr
}

// AllRestoreProgressMgr open func
func AllRestoreProgressMgr(db *gorm.DB) *_AllRestoreProgressMgr {
	if db == nil {
		panic(fmt.Errorf("AllRestoreProgressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRestoreProgressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_restore_progress"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRestoreProgressMgr) GetTableName() string {
	return "__all_restore_progress"
}

// Reset 重置gorm会话
func (obj *_AllRestoreProgressMgr) Reset() *_AllRestoreProgressMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRestoreProgressMgr) Get() (result AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRestoreProgressMgr) Gets() (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRestoreProgressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRestoreProgressMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllRestoreProgressMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllRestoreProgressMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithExternalJobID external_job_id获取
func (obj *_AllRestoreProgressMgr) WithExternalJobID(externalJobID int64) Option {
	return optionFunc(func(o *options) { o.query["external_job_id"] = externalJobID })
}

// WithTenantID tenant_id获取
func (obj *_AllRestoreProgressMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取
func (obj *_AllRestoreProgressMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithStatus status获取
func (obj *_AllRestoreProgressMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithStartTime start_time获取
func (obj *_AllRestoreProgressMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithCompletionTime completion_time获取
func (obj *_AllRestoreProgressMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["completion_time"] = completionTime })
}

// WithPgCount pg_count获取
func (obj *_AllRestoreProgressMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllRestoreProgressMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithPartitionCount partition_count获取
func (obj *_AllRestoreProgressMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllRestoreProgressMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllRestoreProgressMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllRestoreProgressMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithRestoreStartTimestamp restore_start_timestamp获取
func (obj *_AllRestoreProgressMgr) WithRestoreStartTimestamp(restoreStartTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["restore_start_timestamp"] = restoreStartTimestamp })
}

// WithRestoreFinishTimestamp restore_finish_timestamp获取
func (obj *_AllRestoreProgressMgr) WithRestoreFinishTimestamp(restoreFinishTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["restore_finish_timestamp"] = restoreFinishTimestamp })
}

// WithRestoreCurrentTimestamp restore_current_timestamp获取
func (obj *_AllRestoreProgressMgr) WithRestoreCurrentTimestamp(restoreCurrentTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["restore_current_timestamp"] = restoreCurrentTimestamp })
}

// WithInfo info获取
func (obj *_AllRestoreProgressMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithBackupClusterID backup_cluster_id获取
func (obj *_AllRestoreProgressMgr) WithBackupClusterID(backupClusterID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_cluster_id"] = backupClusterID })
}

// WithBackupClusterName backup_cluster_name获取
func (obj *_AllRestoreProgressMgr) WithBackupClusterName(backupClusterName string) Option {
	return optionFunc(func(o *options) { o.query["backup_cluster_name"] = backupClusterName })
}

// WithBackupTenantID backup_tenant_id获取
func (obj *_AllRestoreProgressMgr) WithBackupTenantID(backupTenantID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_tenant_id"] = backupTenantID })
}

// WithBackupTenantName backup_tenant_name获取
func (obj *_AllRestoreProgressMgr) WithBackupTenantName(backupTenantName string) Option {
	return optionFunc(func(o *options) { o.query["backup_tenant_name"] = backupTenantName })
}

// WithWhiteList white_list获取
func (obj *_AllRestoreProgressMgr) WithWhiteList(whiteList string) Option {
	return optionFunc(func(o *options) { o.query["white_list"] = whiteList })
}

// WithBackupSetList backup_set_list获取
func (obj *_AllRestoreProgressMgr) WithBackupSetList(backupSetList string) Option {
	return optionFunc(func(o *options) { o.query["backup_set_list"] = backupSetList })
}

// WithBackupPieceList backup_piece_list获取
func (obj *_AllRestoreProgressMgr) WithBackupPieceList(backupPieceList string) Option {
	return optionFunc(func(o *options) { o.query["backup_piece_list"] = backupPieceList })
}

// GetByOption 功能选项模式获取
func (obj *_AllRestoreProgressMgr) GetByOption(opts ...Option) (result AllRestoreProgress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRestoreProgressMgr) GetByOptions(opts ...Option) (results []*AllRestoreProgress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRestoreProgressMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRestoreProgress, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where(options.query)
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
func (obj *_AllRestoreProgressMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllRestoreProgressMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllRestoreProgressMgr) GetFromJobID(jobID int64) (result AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromExternalJobID 通过external_job_id获取内容
func (obj *_AllRestoreProgressMgr) GetFromExternalJobID(externalJobID int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`external_job_id` = ?", externalJobID).Find(&results).Error

	return
}

// GetBatchFromExternalJobID 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromExternalJobID(externalJobIDs []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`external_job_id` IN (?)", externalJobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllRestoreProgressMgr) GetFromTenantID(tenantID int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllRestoreProgressMgr) GetFromTenantName(tenantName string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllRestoreProgressMgr) GetFromStatus(status string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromStatus(statuss []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllRestoreProgressMgr) GetFromStartTime(startTime time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过completion_time获取内容
func (obj *_AllRestoreProgressMgr) GetFromCompletionTime(completionTime time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`completion_time` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`completion_time` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllRestoreProgressMgr) GetFromPgCount(pgCount int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllRestoreProgressMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllRestoreProgressMgr) GetFromPartitionCount(partitionCount int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllRestoreProgressMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllRestoreProgressMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllRestoreProgressMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromRestoreStartTimestamp 通过restore_start_timestamp获取内容
func (obj *_AllRestoreProgressMgr) GetFromRestoreStartTimestamp(restoreStartTimestamp time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`restore_start_timestamp` = ?", restoreStartTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreStartTimestamp 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromRestoreStartTimestamp(restoreStartTimestamps []time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`restore_start_timestamp` IN (?)", restoreStartTimestamps).Find(&results).Error

	return
}

// GetFromRestoreFinishTimestamp 通过restore_finish_timestamp获取内容
func (obj *_AllRestoreProgressMgr) GetFromRestoreFinishTimestamp(restoreFinishTimestamp time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`restore_finish_timestamp` = ?", restoreFinishTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreFinishTimestamp 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromRestoreFinishTimestamp(restoreFinishTimestamps []time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`restore_finish_timestamp` IN (?)", restoreFinishTimestamps).Find(&results).Error

	return
}

// GetFromRestoreCurrentTimestamp 通过restore_current_timestamp获取内容
func (obj *_AllRestoreProgressMgr) GetFromRestoreCurrentTimestamp(restoreCurrentTimestamp time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`restore_current_timestamp` = ?", restoreCurrentTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreCurrentTimestamp 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromRestoreCurrentTimestamp(restoreCurrentTimestamps []time.Time) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`restore_current_timestamp` IN (?)", restoreCurrentTimestamps).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllRestoreProgressMgr) GetFromInfo(info string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromInfo(infos []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromBackupClusterID 通过backup_cluster_id获取内容
func (obj *_AllRestoreProgressMgr) GetFromBackupClusterID(backupClusterID int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_cluster_id` = ?", backupClusterID).Find(&results).Error

	return
}

// GetBatchFromBackupClusterID 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromBackupClusterID(backupClusterIDs []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_cluster_id` IN (?)", backupClusterIDs).Find(&results).Error

	return
}

// GetFromBackupClusterName 通过backup_cluster_name获取内容
func (obj *_AllRestoreProgressMgr) GetFromBackupClusterName(backupClusterName string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_cluster_name` = ?", backupClusterName).Find(&results).Error

	return
}

// GetBatchFromBackupClusterName 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromBackupClusterName(backupClusterNames []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_cluster_name` IN (?)", backupClusterNames).Find(&results).Error

	return
}

// GetFromBackupTenantID 通过backup_tenant_id获取内容
func (obj *_AllRestoreProgressMgr) GetFromBackupTenantID(backupTenantID int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_tenant_id` = ?", backupTenantID).Find(&results).Error

	return
}

// GetBatchFromBackupTenantID 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromBackupTenantID(backupTenantIDs []int64) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_tenant_id` IN (?)", backupTenantIDs).Find(&results).Error

	return
}

// GetFromBackupTenantName 通过backup_tenant_name获取内容
func (obj *_AllRestoreProgressMgr) GetFromBackupTenantName(backupTenantName string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_tenant_name` = ?", backupTenantName).Find(&results).Error

	return
}

// GetBatchFromBackupTenantName 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromBackupTenantName(backupTenantNames []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_tenant_name` IN (?)", backupTenantNames).Find(&results).Error

	return
}

// GetFromWhiteList 通过white_list获取内容
func (obj *_AllRestoreProgressMgr) GetFromWhiteList(whiteList string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`white_list` = ?", whiteList).Find(&results).Error

	return
}

// GetBatchFromWhiteList 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromWhiteList(whiteLists []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`white_list` IN (?)", whiteLists).Find(&results).Error

	return
}

// GetFromBackupSetList 通过backup_set_list获取内容
func (obj *_AllRestoreProgressMgr) GetFromBackupSetList(backupSetList string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_set_list` = ?", backupSetList).Find(&results).Error

	return
}

// GetBatchFromBackupSetList 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromBackupSetList(backupSetLists []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_set_list` IN (?)", backupSetLists).Find(&results).Error

	return
}

// GetFromBackupPieceList 通过backup_piece_list获取内容
func (obj *_AllRestoreProgressMgr) GetFromBackupPieceList(backupPieceList string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_piece_list` = ?", backupPieceList).Find(&results).Error

	return
}

// GetBatchFromBackupPieceList 批量查找
func (obj *_AllRestoreProgressMgr) GetBatchFromBackupPieceList(backupPieceLists []string) (results []*AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`backup_piece_list` IN (?)", backupPieceLists).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRestoreProgressMgr) FetchByPrimaryKey(jobID int64) (result AllRestoreProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreProgress{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}
