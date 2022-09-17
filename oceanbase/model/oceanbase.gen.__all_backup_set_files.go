package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllBackupSetFilesMgr struct {
	*_BaseMgr
}

// AllBackupSetFilesMgr open func
func AllBackupSetFilesMgr(db *gorm.DB) *_AllBackupSetFilesMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupSetFilesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupSetFilesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_set_files"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupSetFilesMgr) GetTableName() string {
	return "__all_backup_set_files"
}

// Reset 重置gorm会话
func (obj *_AllBackupSetFilesMgr) Reset() *_AllBackupSetFilesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupSetFilesMgr) Get() (result AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupSetFilesMgr) Gets() (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupSetFilesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupSetFilesMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupSetFilesMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupSetFilesMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupSetFilesMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllBackupSetFilesMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithCopyID copy_id获取
func (obj *_AllBackupSetFilesMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithBackupType backup_type获取
func (obj *_AllBackupSetFilesMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllBackupSetFilesMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPrevFullBackupSetID prev_full_backup_set_id获取
func (obj *_AllBackupSetFilesMgr) WithPrevFullBackupSetID(prevFullBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_full_backup_set_id"] = prevFullBackupSetID })
}

// WithPrevIncBackupSetID prev_inc_backup_set_id获取
func (obj *_AllBackupSetFilesMgr) WithPrevIncBackupSetID(prevIncBackupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_inc_backup_set_id"] = prevIncBackupSetID })
}

// WithPrevBackupDataVersion prev_backup_data_version获取
func (obj *_AllBackupSetFilesMgr) WithPrevBackupDataVersion(prevBackupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["prev_backup_data_version"] = prevBackupDataVersion })
}

// WithPgCount pg_count获取
func (obj *_AllBackupSetFilesMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllBackupSetFilesMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllBackupSetFilesMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllBackupSetFilesMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithInputBytes input_bytes获取
func (obj *_AllBackupSetFilesMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllBackupSetFilesMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllBackupSetFilesMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllBackupSetFilesMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCompatible compatible获取
func (obj *_AllBackupSetFilesMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithClusterVersion cluster_version获取
func (obj *_AllBackupSetFilesMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithStatus status获取
func (obj *_AllBackupSetFilesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithResult result获取
func (obj *_AllBackupSetFilesMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithClusterID cluster_id获取
func (obj *_AllBackupSetFilesMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithBackupDataVersion backup_data_version获取
func (obj *_AllBackupSetFilesMgr) WithBackupDataVersion(backupDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_data_version"] = backupDataVersion })
}

// WithBackupSchemaVersion backup_schema_version获取
func (obj *_AllBackupSetFilesMgr) WithBackupSchemaVersion(backupSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["backup_schema_version"] = backupSchemaVersion })
}

// WithClusterVersionDisplay cluster_version_display获取
func (obj *_AllBackupSetFilesMgr) WithClusterVersionDisplay(clusterVersionDisplay string) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_display"] = clusterVersionDisplay })
}

// WithPartitionCount partition_count获取
func (obj *_AllBackupSetFilesMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllBackupSetFilesMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithEncryptionMode encryption_mode获取
func (obj *_AllBackupSetFilesMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["encryption_mode"] = encryptionMode })
}

// WithPasswd passwd获取
func (obj *_AllBackupSetFilesMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithFileStatus file_status获取
func (obj *_AllBackupSetFilesMgr) WithFileStatus(fileStatus string) Option {
	return optionFunc(func(o *options) { o.query["file_status"] = fileStatus })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupSetFilesMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithStartReplayLogTs start_replay_log_ts获取
func (obj *_AllBackupSetFilesMgr) WithStartReplayLogTs(startReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["start_replay_log_ts"] = startReplayLogTs })
}

// WithDate date获取
func (obj *_AllBackupSetFilesMgr) WithDate(date int64) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupSetFilesMgr) GetByOption(opts ...Option) (result AllBackupSetFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupSetFilesMgr) GetByOptions(opts ...Option) (results []*AllBackupSetFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupSetFilesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupSetFiles, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where(options.query)
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
func (obj *_AllBackupSetFilesMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupSetFilesMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupSetFilesMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupSetFilesMgr) GetFromTenantID(tenantID int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllBackupSetFilesMgr) GetFromBackupSetID(backupSetID int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupSetFilesMgr) GetFromCopyID(copyID int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllBackupSetFilesMgr) GetFromBackupType(backupType string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllBackupSetFilesMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPrevFullBackupSetID 通过prev_full_backup_set_id获取内容
func (obj *_AllBackupSetFilesMgr) GetFromPrevFullBackupSetID(prevFullBackupSetID int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`prev_full_backup_set_id` = ?", prevFullBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevFullBackupSetID 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromPrevFullBackupSetID(prevFullBackupSetIDs []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`prev_full_backup_set_id` IN (?)", prevFullBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevIncBackupSetID 通过prev_inc_backup_set_id获取内容
func (obj *_AllBackupSetFilesMgr) GetFromPrevIncBackupSetID(prevIncBackupSetID int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`prev_inc_backup_set_id` = ?", prevIncBackupSetID).Find(&results).Error

	return
}

// GetBatchFromPrevIncBackupSetID 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromPrevIncBackupSetID(prevIncBackupSetIDs []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`prev_inc_backup_set_id` IN (?)", prevIncBackupSetIDs).Find(&results).Error

	return
}

// GetFromPrevBackupDataVersion 通过prev_backup_data_version获取内容
func (obj *_AllBackupSetFilesMgr) GetFromPrevBackupDataVersion(prevBackupDataVersion int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`prev_backup_data_version` = ?", prevBackupDataVersion).Find(&results).Error

	return
}

// GetBatchFromPrevBackupDataVersion 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromPrevBackupDataVersion(prevBackupDataVersions []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`prev_backup_data_version` IN (?)", prevBackupDataVersions).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllBackupSetFilesMgr) GetFromPgCount(pgCount int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllBackupSetFilesMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllBackupSetFilesMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllBackupSetFilesMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllBackupSetFilesMgr) GetFromInputBytes(inputBytes int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllBackupSetFilesMgr) GetFromOutputBytes(outputBytes int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllBackupSetFilesMgr) GetFromStartTime(startTime int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromStartTime(startTimes []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllBackupSetFilesMgr) GetFromEndTime(endTime int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromEndTime(endTimes []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllBackupSetFilesMgr) GetFromCompatible(compatible int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllBackupSetFilesMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupSetFilesMgr) GetFromStatus(status string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupSetFilesMgr) GetFromResult(result int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromResult(results []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromClusterID 通过cluster_id获取内容
func (obj *_AllBackupSetFilesMgr) GetFromClusterID(clusterID int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error

	return
}

// GetBatchFromClusterID 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error

	return
}

// GetFromBackupDataVersion 通过backup_data_version获取内容
func (obj *_AllBackupSetFilesMgr) GetFromBackupDataVersion(backupDataVersion int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_data_version` = ?", backupDataVersion).Find(&results).Error

	return
}

// GetBatchFromBackupDataVersion 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromBackupDataVersion(backupDataVersions []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_data_version` IN (?)", backupDataVersions).Find(&results).Error

	return
}

// GetFromBackupSchemaVersion 通过backup_schema_version获取内容
func (obj *_AllBackupSetFilesMgr) GetFromBackupSchemaVersion(backupSchemaVersion int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_schema_version` = ?", backupSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromBackupSchemaVersion 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromBackupSchemaVersion(backupSchemaVersions []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_schema_version` IN (?)", backupSchemaVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDisplay 通过cluster_version_display获取内容
func (obj *_AllBackupSetFilesMgr) GetFromClusterVersionDisplay(clusterVersionDisplay string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`cluster_version_display` = ?", clusterVersionDisplay).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDisplay 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromClusterVersionDisplay(clusterVersionDisplays []string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`cluster_version_display` IN (?)", clusterVersionDisplays).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllBackupSetFilesMgr) GetFromPartitionCount(partitionCount int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllBackupSetFilesMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过encryption_mode获取内容
func (obj *_AllBackupSetFilesMgr) GetFromEncryptionMode(encryptionMode string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`encryption_mode` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`encryption_mode` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllBackupSetFilesMgr) GetFromPasswd(passwd string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromPasswd(passwds []string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromFileStatus 通过file_status获取内容
func (obj *_AllBackupSetFilesMgr) GetFromFileStatus(fileStatus string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`file_status` = ?", fileStatus).Find(&results).Error

	return
}

// GetBatchFromFileStatus 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromFileStatus(fileStatuss []string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`file_status` IN (?)", fileStatuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupSetFilesMgr) GetFromBackupDest(backupDest string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromStartReplayLogTs 通过start_replay_log_ts获取内容
func (obj *_AllBackupSetFilesMgr) GetFromStartReplayLogTs(startReplayLogTs int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`start_replay_log_ts` = ?", startReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromStartReplayLogTs 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromStartReplayLogTs(startReplayLogTss []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`start_replay_log_ts` IN (?)", startReplayLogTss).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_AllBackupSetFilesMgr) GetFromDate(date int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_AllBackupSetFilesMgr) GetBatchFromDate(dates []int64) (results []*AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupSetFilesMgr) FetchByPrimaryKey(incarnation int64, tenantID int64, backupSetID int64, copyID int64) (result AllBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupSetFiles{}).Where("`incarnation` = ? AND `tenant_id` = ? AND `backup_set_id` = ? AND `copy_id` = ?", incarnation, tenantID, backupSetID, copyID).First(&result).Error

	return
}
