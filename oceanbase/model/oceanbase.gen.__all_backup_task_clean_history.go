package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllBackupTaskCleanHistoryMgr struct {
	*_BaseMgr
}

// AllBackupTaskCleanHistoryMgr open func
func AllBackupTaskCleanHistoryMgr(db *gorm.DB) *_AllBackupTaskCleanHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupTaskCleanHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupTaskCleanHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_task_clean_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupTaskCleanHistoryMgr) GetTableName() string {
	return "__all_backup_task_clean_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupTaskCleanHistoryMgr) Reset() *_AllBackupTaskCleanHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupTaskCleanHistoryMgr) Get() (result AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupTaskCleanHistoryMgr) Gets() (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupTaskCleanHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithJobID job_id获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithBackupType backup_type获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithDeviceType device_type获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithDeviceType(deviceType string) Option {
	return optionFunc(func(o *options) { o.query["device_type"] = deviceType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPrevFullBackupSetID prev_full_backup_set_id获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithPrevFullBackupSetID(prevFullBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_full_backup_set_id"] = prevFullBackupSetID })
}

// WithPrevIncBackupSetID prev_inc_backup_set_id获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithPrevIncBackupSetID(prevIncBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_inc_backup_set_id"] = prevIncBackupSetID })
}

// WithPrevBackupDataVersion prev_backup_data_version获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithPrevBackupDataVersion(prevBackupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["prev_backup_data_version"] = prevBackupDataVersion })
}

// WithPgCount pg_count获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithInputBytes input_bytes获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCompatible compatible获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithClusterVersion cluster_version获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithStatus status获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithResult result获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithClusterID cluster_id获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithBackupDataVersion backup_data_version获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithBackupDataVersion(backupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_data_version"] = backupDataVersion })
}

// WithBackupSchemaVersion backup_schema_version获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithBackupSchemaVersion(backupSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_schema_version"] = backupSchemaVersion })
}

// WithClusterVersionDisplay cluster_version_display获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithClusterVersionDisplay(clusterVersionDisplay string) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_display"] = clusterVersionDisplay })
}

// WithPartitionCount partition_count获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithCopyID copy_id获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithStartReplayLogTs start_replay_log_ts获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithStartReplayLogTs(startReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["start_replay_log_ts"] = startReplayLogTs })
}

// WithDate date获取
func (obj *_AllBackupTaskCleanHistoryMgr) WithDate(date int64) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupTaskCleanHistoryMgr) GetByOption(opts ...Option) (result AllBackupTaskCleanHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupTaskCleanHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupTaskCleanHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupTaskCleanHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupTaskCleanHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where(options.query)
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
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromBackupSetID(backupSetID int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromJobID(jobID int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromBackupType(backupType string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromDeviceType 通过device_type获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromDeviceType(deviceType string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`device_type` = ?", deviceType).Find(&results).Error

	return
}

// GetBatchFromDeviceType 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromDeviceType(deviceTypes []string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`device_type` IN (?)", deviceTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPrevFullBackupSetID 通过prev_full_backup_set_id获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromPrevFullBackupSetID(prevFullBackupSetID int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`prev_full_backup_set_id` = ?", prevFullBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevFullBackupSetID 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromPrevFullBackupSetID(prevFullBackupSetIDs []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`prev_full_backup_set_id` IN (?)", prevFullBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevIncBackupSetID 通过prev_inc_backup_set_id获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromPrevIncBackupSetID(prevIncBackupSetID int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`prev_inc_backup_set_id` = ?", prevIncBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevIncBackupSetID 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromPrevIncBackupSetID(prevIncBackupSetIDs []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`prev_inc_backup_set_id` IN (?)", prevIncBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevBackupDataVersion 通过prev_backup_data_version获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromPrevBackupDataVersion(prevBackupDataVersion int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`prev_backup_data_version` = ?", prevBackupDataVersion).Find(&results).Error

	return
}

// GetBatchFromPrevBackupDataVersion 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromPrevBackupDataVersion(prevBackupDataVersions []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`prev_backup_data_version` IN (?)", prevBackupDataVersions).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromPgCount(pgCount int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromInputBytes(inputBytes int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromOutputBytes(outputBytes int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromStartTime(startTime time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromEndTime(endTime time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromCompatible(compatible int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromStatus(status string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromResult(result int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromResult(results []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromClusterID 通过cluster_id获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromClusterID(clusterID int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error

	return
}

// GetBatchFromClusterID 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromBackupDest(backupDest string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromBackupDataVersion 通过backup_data_version获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromBackupDataVersion(backupDataVersion int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_data_version` = ?", backupDataVersion).Find(&results).Error

	return
}

// GetBatchFromBackupDataVersion 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromBackupDataVersion(backupDataVersions []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_data_version` IN (?)", backupDataVersions).Find(&results).Error

	return
}

// GetFromBackupSchemaVersion 通过backup_schema_version获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromBackupSchemaVersion(backupSchemaVersion int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_schema_version` = ?", backupSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromBackupSchemaVersion 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromBackupSchemaVersion(backupSchemaVersions []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`backup_schema_version` IN (?)", backupSchemaVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDisplay 通过cluster_version_display获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromClusterVersionDisplay(clusterVersionDisplay string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`cluster_version_display` = ?", clusterVersionDisplay).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDisplay 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromClusterVersionDisplay(clusterVersionDisplays []string) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`cluster_version_display` IN (?)", clusterVersionDisplays).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromPartitionCount(partitionCount int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromCopyID(copyID int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromStartReplayLogTs 通过start_replay_log_ts获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromStartReplayLogTs(startReplayLogTs int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`start_replay_log_ts` = ?", startReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromStartReplayLogTs 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromStartReplayLogTs(startReplayLogTss []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`start_replay_log_ts` IN (?)", startReplayLogTss).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_AllBackupTaskCleanHistoryMgr) GetFromDate(date int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_AllBackupTaskCleanHistoryMgr) GetBatchFromDate(dates []int64) (results []*AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupTaskCleanHistoryMgr) FetchByPrimaryKey(tenantID int64, incarnation int64, backupSetID int64, jobID int64) (result AllBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskCleanHistory{}).Where("`tenant_id` = ? AND `incarnation` = ? AND `backup_set_id` = ? AND `job_id` = ?", tenantID, incarnation, backupSetID, jobID).First(&result).Error

	return
}
