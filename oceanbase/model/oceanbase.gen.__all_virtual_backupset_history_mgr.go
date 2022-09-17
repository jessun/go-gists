package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualBackupsetHistoryMgrMgr struct {
	*_BaseMgr
}

// AllVirtualBackupsetHistoryMgrMgr open func
func AllVirtualBackupsetHistoryMgrMgr(db *gorm.DB) *_AllVirtualBackupsetHistoryMgrMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualBackupsetHistoryMgrMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualBackupsetHistoryMgrMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_backupset_history_mgr"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetTableName() string {
	return "__all_virtual_backupset_history_mgr"
}

// Reset 重置gorm会话
func (obj *_AllVirtualBackupsetHistoryMgrMgr) Reset() *_AllVirtualBackupsetHistoryMgrMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) Get() (result AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualBackupsetHistoryMgrMgr) Gets() (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualBackupsetHistoryMgrMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithBackupType backup_type获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithDeviceType device_type获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithDeviceType(deviceType string) Option {
	return optionFunc(func(o *options) { o.query["device_type"] = deviceType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPrevFullBackupSetID prev_full_backup_set_id获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithPrevFullBackupSetID(prevFullBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_full_backup_set_id"] = prevFullBackupSetID })
}

// WithPrevIncBackupSetID prev_inc_backup_set_id获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithPrevIncBackupSetID(prevIncBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_inc_backup_set_id"] = prevIncBackupSetID })
}

// WithPrevBackupDataVersion prev_backup_data_version获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithPrevBackupDataVersion(prevBackupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["prev_backup_data_version"] = prevBackupDataVersion })
}

// WithPgCount pg_count获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithInputBytes input_bytes获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCompatible compatible获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithClusterVersion cluster_version获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithStatus status获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithResult result获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithClusterID cluster_id获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithBackupDest backup_dest获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithBackupDataVersion backup_data_version获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithBackupDataVersion(backupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_data_version"] = backupDataVersion })
}

// WithBackupSchemaVersion backup_schema_version获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithBackupSchemaVersion(backupSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_schema_version"] = backupSchemaVersion })
}

// WithClusterVersionDisplay cluster_version_display获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithClusterVersionDisplay(clusterVersionDisplay string) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_display"] = clusterVersionDisplay })
}

// WithPartitionCount partition_count获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithIsMarkDeleted is_mark_deleted获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithIsMarkDeleted(isMarkDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_mark_deleted"] = isMarkDeleted })
}

// WithEncryptionMode encryption_mode获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["encryption_mode"] = encryptionMode })
}

// WithPasswd passwd获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithBackupRecoveryWindow backup_recovery_window获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) WithBackupRecoveryWindow(backupRecoveryWindow int64) Option {
	return optionFunc(func(o *options) { o.query["backup_recovery_window"] = backupRecoveryWindow })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetByOption(opts ...Option) (result AllVirtualBackupsetHistoryMgr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetByOptions(opts ...Option) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualBackupsetHistoryMgrMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualBackupsetHistoryMgr, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where(options.query)
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
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromBackupSetID(backupSetID int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromBackupType(backupType string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromDeviceType 通过device_type获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromDeviceType(deviceType string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`device_type` = ?", deviceType).Find(&results).Error

	return
}

// GetBatchFromDeviceType 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromDeviceType(deviceTypes []string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`device_type` IN (?)", deviceTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPrevFullBackupSetID 通过prev_full_backup_set_id获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromPrevFullBackupSetID(prevFullBackupSetID int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`prev_full_backup_set_id` = ?", prevFullBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevFullBackupSetID 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromPrevFullBackupSetID(prevFullBackupSetIDs []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`prev_full_backup_set_id` IN (?)", prevFullBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevIncBackupSetID 通过prev_inc_backup_set_id获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromPrevIncBackupSetID(prevIncBackupSetID int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`prev_inc_backup_set_id` = ?", prevIncBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevIncBackupSetID 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromPrevIncBackupSetID(prevIncBackupSetIDs []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`prev_inc_backup_set_id` IN (?)", prevIncBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevBackupDataVersion 通过prev_backup_data_version获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromPrevBackupDataVersion(prevBackupDataVersion int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`prev_backup_data_version` = ?", prevBackupDataVersion).Find(&results).Error

	return
}

// GetBatchFromPrevBackupDataVersion 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromPrevBackupDataVersion(prevBackupDataVersions []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`prev_backup_data_version` IN (?)", prevBackupDataVersions).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromPgCount(pgCount int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromInputBytes(inputBytes int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromOutputBytes(outputBytes int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromStartTime(startTime time.Time) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromEndTime(endTime time.Time) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromCompatible(compatible int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromStatus(status string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromResult(result int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromResult(results []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromClusterID 通过cluster_id获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromClusterID(clusterID int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error

	return
}

// GetBatchFromClusterID 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromBackupDest(backupDest string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromBackupDataVersion 通过backup_data_version获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromBackupDataVersion(backupDataVersion int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_data_version` = ?", backupDataVersion).Find(&results).Error

	return
}

// GetBatchFromBackupDataVersion 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromBackupDataVersion(backupDataVersions []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_data_version` IN (?)", backupDataVersions).Find(&results).Error

	return
}

// GetFromBackupSchemaVersion 通过backup_schema_version获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromBackupSchemaVersion(backupSchemaVersion int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_schema_version` = ?", backupSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromBackupSchemaVersion 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromBackupSchemaVersion(backupSchemaVersions []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_schema_version` IN (?)", backupSchemaVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDisplay 通过cluster_version_display获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromClusterVersionDisplay(clusterVersionDisplay string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`cluster_version_display` = ?", clusterVersionDisplay).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDisplay 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromClusterVersionDisplay(clusterVersionDisplays []string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`cluster_version_display` IN (?)", clusterVersionDisplays).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromPartitionCount(partitionCount int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromIsMarkDeleted 通过is_mark_deleted获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromIsMarkDeleted(isMarkDeleted int8) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`is_mark_deleted` = ?", isMarkDeleted).Find(&results).Error

	return
}

// GetBatchFromIsMarkDeleted 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromIsMarkDeleted(isMarkDeleteds []int8) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`is_mark_deleted` IN (?)", isMarkDeleteds).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过encryption_mode获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromEncryptionMode(encryptionMode string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`encryption_mode` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`encryption_mode` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromPasswd(passwd string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromPasswd(passwds []string) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromBackupRecoveryWindow 通过backup_recovery_window获取内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetFromBackupRecoveryWindow(backupRecoveryWindow int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_recovery_window` = ?", backupRecoveryWindow).Find(&results).Error

	return
}

// GetBatchFromBackupRecoveryWindow 批量查找
func (obj *_AllVirtualBackupsetHistoryMgrMgr) GetBatchFromBackupRecoveryWindow(backupRecoveryWindows []int64) (results []*AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`backup_recovery_window` IN (?)", backupRecoveryWindows).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualBackupsetHistoryMgrMgr) FetchByPrimaryKey(tenantID int64, incarnation int64, backupSetID int64) (result AllVirtualBackupsetHistoryMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupsetHistoryMgr{}).Where("`tenant_id` = ? AND `incarnation` = ? AND `backup_set_id` = ?", tenantID, incarnation, backupSetID).First(&result).Error

	return
}
