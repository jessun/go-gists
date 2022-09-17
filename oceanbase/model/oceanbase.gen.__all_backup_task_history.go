package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupTaskHistoryMgr struct {
	*_BaseMgr
}

// AllBackupTaskHistoryMgr open func
func AllBackupTaskHistoryMgr(db *gorm.DB) *_AllBackupTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_task_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupTaskHistoryMgr) GetTableName() string {
	return "__all_backup_task_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupTaskHistoryMgr) Reset() *_AllBackupTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupTaskHistoryMgr) Get() (result AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupTaskHistoryMgr) Gets() (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupTaskHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupTaskHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupTaskHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllBackupTaskHistoryMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithBackupType backup_type获取
func (obj *_AllBackupTaskHistoryMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithDeviceType device_type获取
func (obj *_AllBackupTaskHistoryMgr) WithDeviceType(deviceType string) Option {
	return optionFunc(func(o *options) { o.query["device_type"] = deviceType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllBackupTaskHistoryMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPrevFullBackupSetID prev_full_backup_set_id获取
func (obj *_AllBackupTaskHistoryMgr) WithPrevFullBackupSetID(prevFullBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_full_backup_set_id"] = prevFullBackupSetID })
}

// WithPrevIncBackupSetID prev_inc_backup_set_id获取
func (obj *_AllBackupTaskHistoryMgr) WithPrevIncBackupSetID(prevIncBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_inc_backup_set_id"] = prevIncBackupSetID })
}

// WithPrevBackupDataVersion prev_backup_data_version获取
func (obj *_AllBackupTaskHistoryMgr) WithPrevBackupDataVersion(prevBackupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["prev_backup_data_version"] = prevBackupDataVersion })
}

// WithPgCount pg_count获取
func (obj *_AllBackupTaskHistoryMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllBackupTaskHistoryMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllBackupTaskHistoryMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllBackupTaskHistoryMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithInputBytes input_bytes获取
func (obj *_AllBackupTaskHistoryMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllBackupTaskHistoryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllBackupTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllBackupTaskHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCompatible compatible获取
func (obj *_AllBackupTaskHistoryMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithClusterVersion cluster_version获取
func (obj *_AllBackupTaskHistoryMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithStatus status获取
func (obj *_AllBackupTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithResult result获取
func (obj *_AllBackupTaskHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithClusterID cluster_id获取
func (obj *_AllBackupTaskHistoryMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupTaskHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithBackupDataVersion backup_data_version获取
func (obj *_AllBackupTaskHistoryMgr) WithBackupDataVersion(backupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_data_version"] = backupDataVersion })
}

// WithBackupSchemaVersion backup_schema_version获取
func (obj *_AllBackupTaskHistoryMgr) WithBackupSchemaVersion(backupSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_schema_version"] = backupSchemaVersion })
}

// WithClusterVersionDisplay cluster_version_display获取
func (obj *_AllBackupTaskHistoryMgr) WithClusterVersionDisplay(clusterVersionDisplay string) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_display"] = clusterVersionDisplay })
}

// WithPartitionCount partition_count获取
func (obj *_AllBackupTaskHistoryMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllBackupTaskHistoryMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithIsMarkDeleted is_mark_deleted获取
func (obj *_AllBackupTaskHistoryMgr) WithIsMarkDeleted(isMarkDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_mark_deleted"] = isMarkDeleted })
}

// WithEncryptionMode encryption_mode获取
func (obj *_AllBackupTaskHistoryMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["encryption_mode"] = encryptionMode })
}

// WithPasswd passwd获取
func (obj *_AllBackupTaskHistoryMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithStartReplayLogTs start_replay_log_ts获取
func (obj *_AllBackupTaskHistoryMgr) WithStartReplayLogTs(startReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["start_replay_log_ts"] = startReplayLogTs })
}

// WithDate date获取
func (obj *_AllBackupTaskHistoryMgr) WithDate(date int64) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupTaskHistoryMgr) GetByOption(opts ...Option) (result AllBackupTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupTaskHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where(options.query)
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
func (obj *_AllBackupTaskHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromBackupSetID(backupSetID int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromBackupType(backupType string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromDeviceType 通过device_type获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromDeviceType(deviceType string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`device_type` = ?", deviceType).Find(&results).Error

	return
}

// GetBatchFromDeviceType 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromDeviceType(deviceTypes []string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`device_type` IN (?)", deviceTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPrevFullBackupSetID 通过prev_full_backup_set_id获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromPrevFullBackupSetID(prevFullBackupSetID int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`prev_full_backup_set_id` = ?", prevFullBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevFullBackupSetID 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromPrevFullBackupSetID(prevFullBackupSetIDs []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`prev_full_backup_set_id` IN (?)", prevFullBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevIncBackupSetID 通过prev_inc_backup_set_id获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromPrevIncBackupSetID(prevIncBackupSetID int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`prev_inc_backup_set_id` = ?", prevIncBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevIncBackupSetID 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromPrevIncBackupSetID(prevIncBackupSetIDs []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`prev_inc_backup_set_id` IN (?)", prevIncBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevBackupDataVersion 通过prev_backup_data_version获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromPrevBackupDataVersion(prevBackupDataVersion int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`prev_backup_data_version` = ?", prevBackupDataVersion).Find(&results).Error

	return
}

// GetBatchFromPrevBackupDataVersion 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromPrevBackupDataVersion(prevBackupDataVersions []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`prev_backup_data_version` IN (?)", prevBackupDataVersions).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromPgCount(pgCount int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromInputBytes(inputBytes int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromOutputBytes(outputBytes int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromEndTime(endTime time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromCompatible(compatible int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromStatus(status string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromResult(result int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromResult(results []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromClusterID 通过cluster_id获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromClusterID(clusterID int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error

	return
}

// GetBatchFromClusterID 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromBackupDest(backupDest string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromBackupDataVersion 通过backup_data_version获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromBackupDataVersion(backupDataVersion int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_data_version` = ?", backupDataVersion).Find(&results).Error

	return
}

// GetBatchFromBackupDataVersion 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromBackupDataVersion(backupDataVersions []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_data_version` IN (?)", backupDataVersions).Find(&results).Error

	return
}

// GetFromBackupSchemaVersion 通过backup_schema_version获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromBackupSchemaVersion(backupSchemaVersion int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_schema_version` = ?", backupSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromBackupSchemaVersion 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromBackupSchemaVersion(backupSchemaVersions []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`backup_schema_version` IN (?)", backupSchemaVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDisplay 通过cluster_version_display获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromClusterVersionDisplay(clusterVersionDisplay string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`cluster_version_display` = ?", clusterVersionDisplay).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDisplay 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromClusterVersionDisplay(clusterVersionDisplays []string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`cluster_version_display` IN (?)", clusterVersionDisplays).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromPartitionCount(partitionCount int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromIsMarkDeleted 通过is_mark_deleted获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromIsMarkDeleted(isMarkDeleted int8) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`is_mark_deleted` = ?", isMarkDeleted).Find(&results).Error

	return
}

// GetBatchFromIsMarkDeleted 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromIsMarkDeleted(isMarkDeleteds []int8) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`is_mark_deleted` IN (?)", isMarkDeleteds).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过encryption_mode获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromEncryptionMode(encryptionMode string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`encryption_mode` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`encryption_mode` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromPasswd(passwd string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromPasswd(passwds []string) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromStartReplayLogTs 通过start_replay_log_ts获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromStartReplayLogTs(startReplayLogTs int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`start_replay_log_ts` = ?", startReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromStartReplayLogTs 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromStartReplayLogTs(startReplayLogTss []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`start_replay_log_ts` IN (?)", startReplayLogTss).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_AllBackupTaskHistoryMgr) GetFromDate(date int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_AllBackupTaskHistoryMgr) GetBatchFromDate(dates []int64) (results []*AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupTaskHistoryMgr) FetchByPrimaryKey(tenantID int64, incarnation int64, backupSetID int64) (result AllBackupTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupTaskHistory{}).Where("`tenant_id` = ? AND `incarnation` = ? AND `backup_set_id` = ?", tenantID, incarnation, backupSetID).First(&result).Error

	return
}
