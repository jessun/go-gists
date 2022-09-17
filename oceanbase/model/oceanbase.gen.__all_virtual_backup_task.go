package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualBackupTaskMgr struct {
	*_BaseMgr
}

// AllVirtualBackupTaskMgr open func
func AllVirtualBackupTaskMgr(db *gorm.DB) *_AllVirtualBackupTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualBackupTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualBackupTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_backup_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualBackupTaskMgr) GetTableName() string {
	return "__all_virtual_backup_task"
}

// Reset 重置gorm会话
func (obj *_AllVirtualBackupTaskMgr) Reset() *_AllVirtualBackupTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualBackupTaskMgr) Get() (result AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualBackupTaskMgr) Gets() (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualBackupTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualBackupTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualBackupTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllVirtualBackupTaskMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualBackupTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualBackupTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithBackupType backup_type获取
func (obj *_AllVirtualBackupTaskMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithDeviceType device_type获取
func (obj *_AllVirtualBackupTaskMgr) WithDeviceType(deviceType string) Option {
	return optionFunc(func(o *options) { o.query["device_type"] = deviceType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualBackupTaskMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPrevFullBackupSetID prev_full_backup_set_id获取
func (obj *_AllVirtualBackupTaskMgr) WithPrevFullBackupSetID(prevFullBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_full_backup_set_id"] = prevFullBackupSetID })
}

// WithPrevIncBackupSetID prev_inc_backup_set_id获取
func (obj *_AllVirtualBackupTaskMgr) WithPrevIncBackupSetID(prevIncBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_inc_backup_set_id"] = prevIncBackupSetID })
}

// WithPrevBackupDataVersion prev_backup_data_version获取
func (obj *_AllVirtualBackupTaskMgr) WithPrevBackupDataVersion(prevBackupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["prev_backup_data_version"] = prevBackupDataVersion })
}

// WithPgCount pg_count获取
func (obj *_AllVirtualBackupTaskMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllVirtualBackupTaskMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllVirtualBackupTaskMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllVirtualBackupTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithInputBytes input_bytes获取
func (obj *_AllVirtualBackupTaskMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllVirtualBackupTaskMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllVirtualBackupTaskMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllVirtualBackupTaskMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCompatible compatible获取
func (obj *_AllVirtualBackupTaskMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithClusterVersion cluster_version获取
func (obj *_AllVirtualBackupTaskMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithStatus status获取
func (obj *_AllVirtualBackupTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithResult result获取
func (obj *_AllVirtualBackupTaskMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithClusterID cluster_id获取
func (obj *_AllVirtualBackupTaskMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithBackupDest backup_dest获取
func (obj *_AllVirtualBackupTaskMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithBackupDataVersion backup_data_version获取
func (obj *_AllVirtualBackupTaskMgr) WithBackupDataVersion(backupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_data_version"] = backupDataVersion })
}

// WithBackupSchemaVersion backup_schema_version获取
func (obj *_AllVirtualBackupTaskMgr) WithBackupSchemaVersion(backupSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_schema_version"] = backupSchemaVersion })
}

// WithClusterVersionDisplay cluster_version_display获取
func (obj *_AllVirtualBackupTaskMgr) WithClusterVersionDisplay(clusterVersionDisplay string) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_display"] = clusterVersionDisplay })
}

// WithPartitionCount partition_count获取
func (obj *_AllVirtualBackupTaskMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllVirtualBackupTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithEncryptionMode encryption_mode获取
func (obj *_AllVirtualBackupTaskMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["encryption_mode"] = encryptionMode })
}

// WithPasswd passwd获取
func (obj *_AllVirtualBackupTaskMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithStartReplayLogTs start_replay_log_ts获取
func (obj *_AllVirtualBackupTaskMgr) WithStartReplayLogTs(startReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["start_replay_log_ts"] = startReplayLogTs })
}

// WithDate date获取
func (obj *_AllVirtualBackupTaskMgr) WithDate(date int64) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualBackupTaskMgr) GetByOption(opts ...Option) (result AllVirtualBackupTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualBackupTaskMgr) GetByOptions(opts ...Option) (results []*AllVirtualBackupTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualBackupTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualBackupTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromBackupSetID(backupSetID int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromBackupType(backupType string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromDeviceType 通过device_type获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromDeviceType(deviceType string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`device_type` = ?", deviceType).Find(&results).Error

	return
}

// GetBatchFromDeviceType 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromDeviceType(deviceTypes []string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`device_type` IN (?)", deviceTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPrevFullBackupSetID 通过prev_full_backup_set_id获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromPrevFullBackupSetID(prevFullBackupSetID int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`prev_full_backup_set_id` = ?", prevFullBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevFullBackupSetID 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromPrevFullBackupSetID(prevFullBackupSetIDs []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`prev_full_backup_set_id` IN (?)", prevFullBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevIncBackupSetID 通过prev_inc_backup_set_id获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromPrevIncBackupSetID(prevIncBackupSetID int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`prev_inc_backup_set_id` = ?", prevIncBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevIncBackupSetID 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromPrevIncBackupSetID(prevIncBackupSetIDs []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`prev_inc_backup_set_id` IN (?)", prevIncBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevBackupDataVersion 通过prev_backup_data_version获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromPrevBackupDataVersion(prevBackupDataVersion int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`prev_backup_data_version` = ?", prevBackupDataVersion).Find(&results).Error

	return
}

// GetBatchFromPrevBackupDataVersion 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromPrevBackupDataVersion(prevBackupDataVersions []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`prev_backup_data_version` IN (?)", prevBackupDataVersions).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromPgCount(pgCount int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromInputBytes(inputBytes int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromOutputBytes(outputBytes int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromStartTime(startTime time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromEndTime(endTime time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromCompatible(compatible int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromStatus(status string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromResult(result int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromResult(results []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromClusterID 通过cluster_id获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromClusterID(clusterID int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error

	return
}

// GetBatchFromClusterID 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromBackupDest(backupDest string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromBackupDataVersion 通过backup_data_version获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromBackupDataVersion(backupDataVersion int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_data_version` = ?", backupDataVersion).Find(&results).Error

	return
}

// GetBatchFromBackupDataVersion 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromBackupDataVersion(backupDataVersions []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_data_version` IN (?)", backupDataVersions).Find(&results).Error

	return
}

// GetFromBackupSchemaVersion 通过backup_schema_version获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromBackupSchemaVersion(backupSchemaVersion int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_schema_version` = ?", backupSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromBackupSchemaVersion 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromBackupSchemaVersion(backupSchemaVersions []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`backup_schema_version` IN (?)", backupSchemaVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDisplay 通过cluster_version_display获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromClusterVersionDisplay(clusterVersionDisplay string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`cluster_version_display` = ?", clusterVersionDisplay).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDisplay 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromClusterVersionDisplay(clusterVersionDisplays []string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`cluster_version_display` IN (?)", clusterVersionDisplays).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromPartitionCount(partitionCount int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过encryption_mode获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromEncryptionMode(encryptionMode string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`encryption_mode` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`encryption_mode` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromPasswd(passwd string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromPasswd(passwds []string) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromStartReplayLogTs 通过start_replay_log_ts获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromStartReplayLogTs(startReplayLogTs int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`start_replay_log_ts` = ?", startReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromStartReplayLogTs 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromStartReplayLogTs(startReplayLogTss []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`start_replay_log_ts` IN (?)", startReplayLogTss).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_AllVirtualBackupTaskMgr) GetFromDate(date int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_AllVirtualBackupTaskMgr) GetBatchFromDate(dates []int64) (results []*AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualBackupTaskMgr) FetchByPrimaryKey(tenantID int64, incarnation int64, backupSetID int64) (result AllVirtualBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupTask{}).Where("`tenant_id` = ? AND `incarnation` = ? AND `backup_set_id` = ?", tenantID, incarnation, backupSetID).First(&result).Error

	return
}
