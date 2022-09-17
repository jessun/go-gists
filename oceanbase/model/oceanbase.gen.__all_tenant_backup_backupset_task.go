package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantBackupBackupsetTaskMgr struct {
	*_BaseMgr
}

// AllTenantBackupBackupsetTaskMgr open func
func AllTenantBackupBackupsetTaskMgr(db *gorm.DB) *_AllTenantBackupBackupsetTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantBackupBackupsetTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantBackupBackupsetTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_backup_backupset_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantBackupBackupsetTaskMgr) GetTableName() string {
	return "__all_tenant_backup_backupset_task"
}

// Reset 重置gorm会话
func (obj *_AllTenantBackupBackupsetTaskMgr) Reset() *_AllTenantBackupBackupsetTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantBackupBackupsetTaskMgr) Get() (result AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantBackupBackupsetTaskMgr) Gets() (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantBackupBackupsetTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithIncarnation incarnation获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithCopyID copy_id获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithBackupType backup_type获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPrevFullBackupSetID prev_full_backup_set_id获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithPrevFullBackupSetID(prevFullBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_full_backup_set_id"] = prevFullBackupSetID })
}

// WithPrevIncBackupSetID prev_inc_backup_set_id获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithPrevIncBackupSetID(prevIncBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_inc_backup_set_id"] = prevIncBackupSetID })
}

// WithPrevBackupDataVersion prev_backup_data_version获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithPrevBackupDataVersion(prevBackupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["prev_backup_data_version"] = prevBackupDataVersion })
}

// WithInputBytes input_bytes获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCompatible compatible获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithClusterID cluster_id获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithClusterVersion cluster_version获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithClusterVersionDisplay cluster_version_display获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithClusterVersionDisplay(clusterVersionDisplay string) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_display"] = clusterVersionDisplay })
}

// WithStatus status获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSrcBackupDest src_backup_dest获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithSrcBackupDest(srcBackupDest string) Option {
	return optionFunc(func(o *options) { o.query["src_backup_dest"] = srcBackupDest })
}

// WithDstBackupDest dst_backup_dest获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithDstBackupDest(dstBackupDest string) Option {
	return optionFunc(func(o *options) { o.query["dst_backup_dest"] = dstBackupDest })
}

// WithSrcDeviceType src_device_type获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithSrcDeviceType(srcDeviceType string) Option {
	return optionFunc(func(o *options) { o.query["src_device_type"] = srcDeviceType })
}

// WithDstDeviceType dst_device_type获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithDstDeviceType(dstDeviceType string) Option {
	return optionFunc(func(o *options) { o.query["dst_device_type"] = dstDeviceType })
}

// WithBackupDataVersion backup_data_version获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithBackupDataVersion(backupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_data_version"] = backupDataVersion })
}

// WithBackupSchemaVersion backup_schema_version获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithBackupSchemaVersion(backupSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_schema_version"] = backupSchemaVersion })
}

// WithTotalPgCount total_pg_count获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithTotalPgCount(totalPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_pg_count"] = totalPgCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithTotalPartitionCount total_partition_count获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_partition_count"] = totalPartitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithTotalMacroBlockCount total_macro_block_count获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_macro_block_count"] = totalMacroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithResult result获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithEncryptionMode encryption_mode获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["encryption_mode"] = encryptionMode })
}

// WithPasswd passwd获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithStartReplayLogTs start_replay_log_ts获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithStartReplayLogTs(startReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["start_replay_log_ts"] = startReplayLogTs })
}

// WithDate date获取
func (obj *_AllTenantBackupBackupsetTaskMgr) WithDate(date int64) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantBackupBackupsetTaskMgr) GetByOption(opts ...Option) (result AllTenantBackupBackupsetTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantBackupBackupsetTaskMgr) GetByOptions(opts ...Option) (results []*AllTenantBackupBackupsetTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantBackupBackupsetTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantBackupBackupsetTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where(options.query)
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
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromTenantID(tenantID int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromJobID(jobID int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromIncarnation(incarnation int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromBackupSetID(backupSetID int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromCopyID(copyID int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromBackupType(backupType string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPrevFullBackupSetID 通过prev_full_backup_set_id获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromPrevFullBackupSetID(prevFullBackupSetID int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`prev_full_backup_set_id` = ?", prevFullBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevFullBackupSetID 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromPrevFullBackupSetID(prevFullBackupSetIDs []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`prev_full_backup_set_id` IN (?)", prevFullBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevIncBackupSetID 通过prev_inc_backup_set_id获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromPrevIncBackupSetID(prevIncBackupSetID int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`prev_inc_backup_set_id` = ?", prevIncBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevIncBackupSetID 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromPrevIncBackupSetID(prevIncBackupSetIDs []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`prev_inc_backup_set_id` IN (?)", prevIncBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevBackupDataVersion 通过prev_backup_data_version获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromPrevBackupDataVersion(prevBackupDataVersion int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`prev_backup_data_version` = ?", prevBackupDataVersion).Find(&results).Error

	return
}

// GetBatchFromPrevBackupDataVersion 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromPrevBackupDataVersion(prevBackupDataVersions []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`prev_backup_data_version` IN (?)", prevBackupDataVersions).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromInputBytes(inputBytes int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromOutputBytes(outputBytes int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromStartTime(startTime time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromEndTime(endTime time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromCompatible(compatible int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromClusterID 通过cluster_id获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromClusterID(clusterID int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error

	return
}

// GetBatchFromClusterID 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDisplay 通过cluster_version_display获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromClusterVersionDisplay(clusterVersionDisplay string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`cluster_version_display` = ?", clusterVersionDisplay).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDisplay 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromClusterVersionDisplay(clusterVersionDisplays []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`cluster_version_display` IN (?)", clusterVersionDisplays).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromStatus(status string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSrcBackupDest 通过src_backup_dest获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromSrcBackupDest(srcBackupDest string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`src_backup_dest` = ?", srcBackupDest).Find(&results).Error

	return
}

// GetBatchFromSrcBackupDest 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromSrcBackupDest(srcBackupDests []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`src_backup_dest` IN (?)", srcBackupDests).Find(&results).Error

	return
}

// GetFromDstBackupDest 通过dst_backup_dest获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromDstBackupDest(dstBackupDest string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`dst_backup_dest` = ?", dstBackupDest).Find(&results).Error

	return
}

// GetBatchFromDstBackupDest 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromDstBackupDest(dstBackupDests []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`dst_backup_dest` IN (?)", dstBackupDests).Find(&results).Error

	return
}

// GetFromSrcDeviceType 通过src_device_type获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromSrcDeviceType(srcDeviceType string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`src_device_type` = ?", srcDeviceType).Find(&results).Error

	return
}

// GetBatchFromSrcDeviceType 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromSrcDeviceType(srcDeviceTypes []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`src_device_type` IN (?)", srcDeviceTypes).Find(&results).Error

	return
}

// GetFromDstDeviceType 通过dst_device_type获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromDstDeviceType(dstDeviceType string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`dst_device_type` = ?", dstDeviceType).Find(&results).Error

	return
}

// GetBatchFromDstDeviceType 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromDstDeviceType(dstDeviceTypes []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`dst_device_type` IN (?)", dstDeviceTypes).Find(&results).Error

	return
}

// GetFromBackupDataVersion 通过backup_data_version获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromBackupDataVersion(backupDataVersion int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_data_version` = ?", backupDataVersion).Find(&results).Error

	return
}

// GetBatchFromBackupDataVersion 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromBackupDataVersion(backupDataVersions []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_data_version` IN (?)", backupDataVersions).Find(&results).Error

	return
}

// GetFromBackupSchemaVersion 通过backup_schema_version获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromBackupSchemaVersion(backupSchemaVersion int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_schema_version` = ?", backupSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromBackupSchemaVersion 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromBackupSchemaVersion(backupSchemaVersions []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`backup_schema_version` IN (?)", backupSchemaVersions).Find(&results).Error

	return
}

// GetFromTotalPgCount 通过total_pg_count获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromTotalPgCount(totalPgCount int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`total_pg_count` = ?", totalPgCount).Find(&results).Error

	return
}

// GetBatchFromTotalPgCount 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromTotalPgCount(totalPgCounts []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`total_pg_count` IN (?)", totalPgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过total_partition_count获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`total_partition_count` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`total_partition_count` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过total_macro_block_count获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`total_macro_block_count` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`total_macro_block_count` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromResult(result int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromResult(results []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过encryption_mode获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromEncryptionMode(encryptionMode string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`encryption_mode` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`encryption_mode` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromPasswd(passwd string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromPasswd(passwds []string) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromStartReplayLogTs 通过start_replay_log_ts获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromStartReplayLogTs(startReplayLogTs int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`start_replay_log_ts` = ?", startReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromStartReplayLogTs 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromStartReplayLogTs(startReplayLogTss []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`start_replay_log_ts` IN (?)", startReplayLogTss).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_AllTenantBackupBackupsetTaskMgr) GetFromDate(date int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_AllTenantBackupBackupsetTaskMgr) GetBatchFromDate(dates []int64) (results []*AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantBackupBackupsetTaskMgr) FetchByPrimaryKey(tenantID int64, jobID int64, incarnation int64, backupSetID int64, copyID int64) (result AllTenantBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupsetTask{}).Where("`tenant_id` = ? AND `job_id` = ? AND `incarnation` = ? AND `backup_set_id` = ? AND `copy_id` = ?", tenantID, jobID, incarnation, backupSetID, copyID).First(&result).Error

	return
}
