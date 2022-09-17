package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupBackupsetTaskHistoryMgr struct {
	*_BaseMgr
}

// AllBackupBackupsetTaskHistoryMgr open func
func AllBackupBackupsetTaskHistoryMgr(db *gorm.DB) *_AllBackupBackupsetTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupBackupsetTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupBackupsetTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_backupset_task_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetTableName() string {
	return "__all_backup_backupset_task_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupBackupsetTaskHistoryMgr) Reset() *_AllBackupBackupsetTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) Get() (result AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupBackupsetTaskHistoryMgr) Gets() (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupBackupsetTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithCopyID copy_id获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithBackupType backup_type获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPrevFullBackupSetID prev_full_backup_set_id获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithPrevFullBackupSetID(prevFullBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_full_backup_set_id"] = prevFullBackupSetID })
}

// WithPrevIncBackupSetID prev_inc_backup_set_id获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithPrevIncBackupSetID(prevIncBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_inc_backup_set_id"] = prevIncBackupSetID })
}

// WithPrevBackupDataVersion prev_backup_data_version获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithPrevBackupDataVersion(prevBackupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["prev_backup_data_version"] = prevBackupDataVersion })
}

// WithInputBytes input_bytes获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCompatible compatible获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithClusterID cluster_id获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithClusterVersion cluster_version获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithClusterVersionDisplay cluster_version_display获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithClusterVersionDisplay(clusterVersionDisplay string) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_display"] = clusterVersionDisplay })
}

// WithStatus status获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSrcBackupDest src_backup_dest获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithSrcBackupDest(srcBackupDest string) Option {
	return optionFunc(func(o *options) { o.query["src_backup_dest"] = srcBackupDest })
}

// WithDstBackupDest dst_backup_dest获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithDstBackupDest(dstBackupDest string) Option {
	return optionFunc(func(o *options) { o.query["dst_backup_dest"] = dstBackupDest })
}

// WithSrcDeviceType src_device_type获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithSrcDeviceType(srcDeviceType string) Option {
	return optionFunc(func(o *options) { o.query["src_device_type"] = srcDeviceType })
}

// WithDstDeviceType dst_device_type获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithDstDeviceType(dstDeviceType string) Option {
	return optionFunc(func(o *options) { o.query["dst_device_type"] = dstDeviceType })
}

// WithBackupDataVersion backup_data_version获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithBackupDataVersion(backupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_data_version"] = backupDataVersion })
}

// WithBackupSchemaVersion backup_schema_version获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithBackupSchemaVersion(backupSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_schema_version"] = backupSchemaVersion })
}

// WithTotalPgCount total_pg_count获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithTotalPgCount(totalPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_pg_count"] = totalPgCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithTotalPartitionCount total_partition_count获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_partition_count"] = totalPartitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithTotalMacroBlockCount total_macro_block_count获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_macro_block_count"] = totalMacroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithResult result获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithEncryptionMode encryption_mode获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["encryption_mode"] = encryptionMode })
}

// WithPasswd passwd获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithIsMarkDeleted is_mark_deleted获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithIsMarkDeleted(isMarkDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_mark_deleted"] = isMarkDeleted })
}

// WithStartReplayLogTs start_replay_log_ts获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithStartReplayLogTs(startReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["start_replay_log_ts"] = startReplayLogTs })
}

// WithDate date获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) WithDate(date int64) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetByOption(opts ...Option) (result AllBackupBackupsetTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupBackupsetTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupBackupsetTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupBackupsetTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where(options.query)
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
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromJobID(jobID int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromBackupSetID(backupSetID int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromCopyID(copyID int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromBackupType(backupType string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPrevFullBackupSetID 通过prev_full_backup_set_id获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromPrevFullBackupSetID(prevFullBackupSetID int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`prev_full_backup_set_id` = ?", prevFullBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevFullBackupSetID 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromPrevFullBackupSetID(prevFullBackupSetIDs []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`prev_full_backup_set_id` IN (?)", prevFullBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevIncBackupSetID 通过prev_inc_backup_set_id获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromPrevIncBackupSetID(prevIncBackupSetID int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`prev_inc_backup_set_id` = ?", prevIncBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevIncBackupSetID 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromPrevIncBackupSetID(prevIncBackupSetIDs []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`prev_inc_backup_set_id` IN (?)", prevIncBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevBackupDataVersion 通过prev_backup_data_version获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromPrevBackupDataVersion(prevBackupDataVersion int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`prev_backup_data_version` = ?", prevBackupDataVersion).Find(&results).Error

	return
}

// GetBatchFromPrevBackupDataVersion 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromPrevBackupDataVersion(prevBackupDataVersions []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`prev_backup_data_version` IN (?)", prevBackupDataVersions).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromInputBytes(inputBytes int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromOutputBytes(outputBytes int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromEndTime(endTime time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromCompatible(compatible int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromClusterID 通过cluster_id获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromClusterID(clusterID int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error

	return
}

// GetBatchFromClusterID 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDisplay 通过cluster_version_display获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromClusterVersionDisplay(clusterVersionDisplay string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`cluster_version_display` = ?", clusterVersionDisplay).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDisplay 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromClusterVersionDisplay(clusterVersionDisplays []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`cluster_version_display` IN (?)", clusterVersionDisplays).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromStatus(status string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSrcBackupDest 通过src_backup_dest获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromSrcBackupDest(srcBackupDest string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`src_backup_dest` = ?", srcBackupDest).Find(&results).Error

	return
}

// GetBatchFromSrcBackupDest 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromSrcBackupDest(srcBackupDests []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`src_backup_dest` IN (?)", srcBackupDests).Find(&results).Error

	return
}

// GetFromDstBackupDest 通过dst_backup_dest获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromDstBackupDest(dstBackupDest string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`dst_backup_dest` = ?", dstBackupDest).Find(&results).Error

	return
}

// GetBatchFromDstBackupDest 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromDstBackupDest(dstBackupDests []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`dst_backup_dest` IN (?)", dstBackupDests).Find(&results).Error

	return
}

// GetFromSrcDeviceType 通过src_device_type获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromSrcDeviceType(srcDeviceType string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`src_device_type` = ?", srcDeviceType).Find(&results).Error

	return
}

// GetBatchFromSrcDeviceType 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromSrcDeviceType(srcDeviceTypes []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`src_device_type` IN (?)", srcDeviceTypes).Find(&results).Error

	return
}

// GetFromDstDeviceType 通过dst_device_type获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromDstDeviceType(dstDeviceType string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`dst_device_type` = ?", dstDeviceType).Find(&results).Error

	return
}

// GetBatchFromDstDeviceType 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromDstDeviceType(dstDeviceTypes []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`dst_device_type` IN (?)", dstDeviceTypes).Find(&results).Error

	return
}

// GetFromBackupDataVersion 通过backup_data_version获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromBackupDataVersion(backupDataVersion int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_data_version` = ?", backupDataVersion).Find(&results).Error

	return
}

// GetBatchFromBackupDataVersion 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromBackupDataVersion(backupDataVersions []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_data_version` IN (?)", backupDataVersions).Find(&results).Error

	return
}

// GetFromBackupSchemaVersion 通过backup_schema_version获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromBackupSchemaVersion(backupSchemaVersion int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_schema_version` = ?", backupSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromBackupSchemaVersion 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromBackupSchemaVersion(backupSchemaVersions []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`backup_schema_version` IN (?)", backupSchemaVersions).Find(&results).Error

	return
}

// GetFromTotalPgCount 通过total_pg_count获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromTotalPgCount(totalPgCount int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`total_pg_count` = ?", totalPgCount).Find(&results).Error

	return
}

// GetBatchFromTotalPgCount 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromTotalPgCount(totalPgCounts []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`total_pg_count` IN (?)", totalPgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过total_partition_count获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`total_partition_count` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`total_partition_count` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过total_macro_block_count获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`total_macro_block_count` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`total_macro_block_count` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromResult(result int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromResult(results []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过encryption_mode获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromEncryptionMode(encryptionMode string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`encryption_mode` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`encryption_mode` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromPasswd(passwd string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromPasswd(passwds []string) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromIsMarkDeleted 通过is_mark_deleted获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromIsMarkDeleted(isMarkDeleted int8) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`is_mark_deleted` = ?", isMarkDeleted).Find(&results).Error

	return
}

// GetBatchFromIsMarkDeleted 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromIsMarkDeleted(isMarkDeleteds []int8) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`is_mark_deleted` IN (?)", isMarkDeleteds).Find(&results).Error

	return
}

// GetFromStartReplayLogTs 通过start_replay_log_ts获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromStartReplayLogTs(startReplayLogTs int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`start_replay_log_ts` = ?", startReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromStartReplayLogTs 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromStartReplayLogTs(startReplayLogTss []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`start_replay_log_ts` IN (?)", startReplayLogTss).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetFromDate(date int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_AllBackupBackupsetTaskHistoryMgr) GetBatchFromDate(dates []int64) (results []*AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupBackupsetTaskHistoryMgr) FetchByPrimaryKey(tenantID int64, jobID int64, incarnation int64, backupSetID int64, copyID int64) (result AllBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetTaskHistory{}).Where("`tenant_id` = ? AND `job_id` = ? AND `incarnation` = ? AND `backup_set_id` = ? AND `copy_id` = ?", tenantID, jobID, incarnation, backupSetID, copyID).First(&result).Error

	return
}
