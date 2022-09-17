package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllRestoreHistoryMgr struct {
	*_BaseMgr
}

// AllRestoreHistoryMgr open func
func AllRestoreHistoryMgr(db *gorm.DB) *_AllRestoreHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllRestoreHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRestoreHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_restore_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRestoreHistoryMgr) GetTableName() string {
	return "__all_restore_history"
}

// Reset 重置gorm会话
func (obj *_AllRestoreHistoryMgr) Reset() *_AllRestoreHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRestoreHistoryMgr) Get() (result AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRestoreHistoryMgr) Gets() (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRestoreHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRestoreHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllRestoreHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllRestoreHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithExternalJobID external_job_id获取
func (obj *_AllRestoreHistoryMgr) WithExternalJobID(externalJobID int64) Option {
	return optionFunc(func(o *options) { o.query["external_job_id"] = externalJobID })
}

// WithTenantID tenant_id获取
func (obj *_AllRestoreHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取
func (obj *_AllRestoreHistoryMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithStatus status获取
func (obj *_AllRestoreHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithStartTime start_time获取
func (obj *_AllRestoreHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithCompletionTime completion_time获取
func (obj *_AllRestoreHistoryMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["completion_time"] = completionTime })
}

// WithPgCount pg_count获取
func (obj *_AllRestoreHistoryMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllRestoreHistoryMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithPartitionCount partition_count获取
func (obj *_AllRestoreHistoryMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllRestoreHistoryMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllRestoreHistoryMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllRestoreHistoryMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithRestoreStartTimestamp restore_start_timestamp获取
func (obj *_AllRestoreHistoryMgr) WithRestoreStartTimestamp(restoreStartTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["restore_start_timestamp"] = restoreStartTimestamp })
}

// WithRestoreFinishTimestamp restore_finish_timestamp获取
func (obj *_AllRestoreHistoryMgr) WithRestoreFinishTimestamp(restoreFinishTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["restore_finish_timestamp"] = restoreFinishTimestamp })
}

// WithRestoreCurrentTimestamp restore_current_timestamp获取
func (obj *_AllRestoreHistoryMgr) WithRestoreCurrentTimestamp(restoreCurrentTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["restore_current_timestamp"] = restoreCurrentTimestamp })
}

// WithRestoreDataVersion restore_data_version获取
func (obj *_AllRestoreHistoryMgr) WithRestoreDataVersion(restoreDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["restore_data_version"] = restoreDataVersion })
}

// WithBackupDest backup_dest获取
func (obj *_AllRestoreHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithRestoreOption restore_option获取
func (obj *_AllRestoreHistoryMgr) WithRestoreOption(restoreOption string) Option {
	return optionFunc(func(o *options) { o.query["restore_option"] = restoreOption })
}

// WithInfo info获取
func (obj *_AllRestoreHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithBackupClusterID backup_cluster_id获取
func (obj *_AllRestoreHistoryMgr) WithBackupClusterID(backupClusterID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_cluster_id"] = backupClusterID })
}

// WithBackupClusterName backup_cluster_name获取
func (obj *_AllRestoreHistoryMgr) WithBackupClusterName(backupClusterName string) Option {
	return optionFunc(func(o *options) { o.query["backup_cluster_name"] = backupClusterName })
}

// WithBackupTenantID backup_tenant_id获取
func (obj *_AllRestoreHistoryMgr) WithBackupTenantID(backupTenantID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_tenant_id"] = backupTenantID })
}

// WithBackupTenantName backup_tenant_name获取
func (obj *_AllRestoreHistoryMgr) WithBackupTenantName(backupTenantName string) Option {
	return optionFunc(func(o *options) { o.query["backup_tenant_name"] = backupTenantName })
}

// WithWhiteList white_list获取
func (obj *_AllRestoreHistoryMgr) WithWhiteList(whiteList string) Option {
	return optionFunc(func(o *options) { o.query["white_list"] = whiteList })
}

// WithBackupSetList backup_set_list获取
func (obj *_AllRestoreHistoryMgr) WithBackupSetList(backupSetList string) Option {
	return optionFunc(func(o *options) { o.query["backup_set_list"] = backupSetList })
}

// WithBackupPieceList backup_piece_list获取
func (obj *_AllRestoreHistoryMgr) WithBackupPieceList(backupPieceList string) Option {
	return optionFunc(func(o *options) { o.query["backup_piece_list"] = backupPieceList })
}

// GetByOption 功能选项模式获取
func (obj *_AllRestoreHistoryMgr) GetByOption(opts ...Option) (result AllRestoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRestoreHistoryMgr) GetByOptions(opts ...Option) (results []*AllRestoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRestoreHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRestoreHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where(options.query)
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
func (obj *_AllRestoreHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllRestoreHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllRestoreHistoryMgr) GetFromJobID(jobID int64) (result AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromExternalJobID 通过external_job_id获取内容
func (obj *_AllRestoreHistoryMgr) GetFromExternalJobID(externalJobID int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`external_job_id` = ?", externalJobID).Find(&results).Error

	return
}

// GetBatchFromExternalJobID 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromExternalJobID(externalJobIDs []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`external_job_id` IN (?)", externalJobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllRestoreHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllRestoreHistoryMgr) GetFromTenantName(tenantName string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllRestoreHistoryMgr) GetFromStatus(status string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllRestoreHistoryMgr) GetFromStartTime(startTime time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过completion_time获取内容
func (obj *_AllRestoreHistoryMgr) GetFromCompletionTime(completionTime time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`completion_time` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`completion_time` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllRestoreHistoryMgr) GetFromPgCount(pgCount int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllRestoreHistoryMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllRestoreHistoryMgr) GetFromPartitionCount(partitionCount int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllRestoreHistoryMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllRestoreHistoryMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllRestoreHistoryMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromRestoreStartTimestamp 通过restore_start_timestamp获取内容
func (obj *_AllRestoreHistoryMgr) GetFromRestoreStartTimestamp(restoreStartTimestamp time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_start_timestamp` = ?", restoreStartTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreStartTimestamp 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromRestoreStartTimestamp(restoreStartTimestamps []time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_start_timestamp` IN (?)", restoreStartTimestamps).Find(&results).Error

	return
}

// GetFromRestoreFinishTimestamp 通过restore_finish_timestamp获取内容
func (obj *_AllRestoreHistoryMgr) GetFromRestoreFinishTimestamp(restoreFinishTimestamp time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_finish_timestamp` = ?", restoreFinishTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreFinishTimestamp 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromRestoreFinishTimestamp(restoreFinishTimestamps []time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_finish_timestamp` IN (?)", restoreFinishTimestamps).Find(&results).Error

	return
}

// GetFromRestoreCurrentTimestamp 通过restore_current_timestamp获取内容
func (obj *_AllRestoreHistoryMgr) GetFromRestoreCurrentTimestamp(restoreCurrentTimestamp time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_current_timestamp` = ?", restoreCurrentTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreCurrentTimestamp 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromRestoreCurrentTimestamp(restoreCurrentTimestamps []time.Time) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_current_timestamp` IN (?)", restoreCurrentTimestamps).Find(&results).Error

	return
}

// GetFromRestoreDataVersion 通过restore_data_version获取内容
func (obj *_AllRestoreHistoryMgr) GetFromRestoreDataVersion(restoreDataVersion int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_data_version` = ?", restoreDataVersion).Find(&results).Error

	return
}

// GetBatchFromRestoreDataVersion 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromRestoreDataVersion(restoreDataVersions []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_data_version` IN (?)", restoreDataVersions).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllRestoreHistoryMgr) GetFromBackupDest(backupDest string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromRestoreOption 通过restore_option获取内容
func (obj *_AllRestoreHistoryMgr) GetFromRestoreOption(restoreOption string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_option` = ?", restoreOption).Find(&results).Error

	return
}

// GetBatchFromRestoreOption 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromRestoreOption(restoreOptions []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`restore_option` IN (?)", restoreOptions).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllRestoreHistoryMgr) GetFromInfo(info string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromBackupClusterID 通过backup_cluster_id获取内容
func (obj *_AllRestoreHistoryMgr) GetFromBackupClusterID(backupClusterID int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_cluster_id` = ?", backupClusterID).Find(&results).Error

	return
}

// GetBatchFromBackupClusterID 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromBackupClusterID(backupClusterIDs []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_cluster_id` IN (?)", backupClusterIDs).Find(&results).Error

	return
}

// GetFromBackupClusterName 通过backup_cluster_name获取内容
func (obj *_AllRestoreHistoryMgr) GetFromBackupClusterName(backupClusterName string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_cluster_name` = ?", backupClusterName).Find(&results).Error

	return
}

// GetBatchFromBackupClusterName 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromBackupClusterName(backupClusterNames []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_cluster_name` IN (?)", backupClusterNames).Find(&results).Error

	return
}

// GetFromBackupTenantID 通过backup_tenant_id获取内容
func (obj *_AllRestoreHistoryMgr) GetFromBackupTenantID(backupTenantID int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_tenant_id` = ?", backupTenantID).Find(&results).Error

	return
}

// GetBatchFromBackupTenantID 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromBackupTenantID(backupTenantIDs []int64) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_tenant_id` IN (?)", backupTenantIDs).Find(&results).Error

	return
}

// GetFromBackupTenantName 通过backup_tenant_name获取内容
func (obj *_AllRestoreHistoryMgr) GetFromBackupTenantName(backupTenantName string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_tenant_name` = ?", backupTenantName).Find(&results).Error

	return
}

// GetBatchFromBackupTenantName 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromBackupTenantName(backupTenantNames []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_tenant_name` IN (?)", backupTenantNames).Find(&results).Error

	return
}

// GetFromWhiteList 通过white_list获取内容
func (obj *_AllRestoreHistoryMgr) GetFromWhiteList(whiteList string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`white_list` = ?", whiteList).Find(&results).Error

	return
}

// GetBatchFromWhiteList 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromWhiteList(whiteLists []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`white_list` IN (?)", whiteLists).Find(&results).Error

	return
}

// GetFromBackupSetList 通过backup_set_list获取内容
func (obj *_AllRestoreHistoryMgr) GetFromBackupSetList(backupSetList string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_set_list` = ?", backupSetList).Find(&results).Error

	return
}

// GetBatchFromBackupSetList 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromBackupSetList(backupSetLists []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_set_list` IN (?)", backupSetLists).Find(&results).Error

	return
}

// GetFromBackupPieceList 通过backup_piece_list获取内容
func (obj *_AllRestoreHistoryMgr) GetFromBackupPieceList(backupPieceList string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_piece_list` = ?", backupPieceList).Find(&results).Error

	return
}

// GetBatchFromBackupPieceList 批量查找
func (obj *_AllRestoreHistoryMgr) GetBatchFromBackupPieceList(backupPieceLists []string) (results []*AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`backup_piece_list` IN (?)", backupPieceLists).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRestoreHistoryMgr) FetchByPrimaryKey(jobID int64) (result AllRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreHistory{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}
