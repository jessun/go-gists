package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualPgBackupTaskMgr struct {
	*_BaseMgr
}

// AllVirtualPgBackupTaskMgr open func
func AllVirtualPgBackupTaskMgr(db *gorm.DB) *_AllVirtualPgBackupTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPgBackupTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPgBackupTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_pg_backup_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPgBackupTaskMgr) GetTableName() string {
	return "__all_virtual_pg_backup_task"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPgBackupTaskMgr) Reset() *_AllVirtualPgBackupTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPgBackupTaskMgr) Get() (result AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPgBackupTaskMgr) Gets() (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPgBackupTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPgBackupTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPgBackupTaskMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPgBackupTaskMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualPgBackupTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllVirtualPgBackupTaskMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualPgBackupTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualPgBackupTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithBackupType backup_type获取
func (obj *_AllVirtualPgBackupTaskMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualPgBackupTaskMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithPartitionCount partition_count获取
func (obj *_AllVirtualPgBackupTaskMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithMacroBlockCount macro_block_count获取
func (obj *_AllVirtualPgBackupTaskMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_count"] = macroBlockCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllVirtualPgBackupTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllVirtualPgBackupTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithInputBytes input_bytes获取
func (obj *_AllVirtualPgBackupTaskMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllVirtualPgBackupTaskMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithStartTime start_time获取
func (obj *_AllVirtualPgBackupTaskMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllVirtualPgBackupTaskMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithRetryCount retry_count获取
func (obj *_AllVirtualPgBackupTaskMgr) WithRetryCount(retryCount int64) Option {
	return optionFunc(func(o *options) { o.query["retry_count"] = retryCount })
}

// WithReplicaRole replica_role获取
func (obj *_AllVirtualPgBackupTaskMgr) WithReplicaRole(replicaRole int64) Option {
	return optionFunc(func(o *options) { o.query["replica_role"] = replicaRole })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualPgBackupTaskMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPgBackupTaskMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPgBackupTaskMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithStatus status获取
func (obj *_AllVirtualPgBackupTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTaskID task_id获取
func (obj *_AllVirtualPgBackupTaskMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithResult result获取
func (obj *_AllVirtualPgBackupTaskMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithTraceID trace_id获取
func (obj *_AllVirtualPgBackupTaskMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPgBackupTaskMgr) GetByOption(opts ...Option) (result AllVirtualPgBackupTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPgBackupTaskMgr) GetByOptions(opts ...Option) (results []*AllVirtualPgBackupTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPgBackupTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPgBackupTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where(options.query)
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
func (obj *_AllVirtualPgBackupTaskMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromTableID(tableID int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromBackupSetID(backupSetID int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromBackupType 通过backup_type获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromBackupType(backupType string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromBackupType(backupTypes []string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromPartitionCount(partitionCount int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过macro_block_count获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`macro_block_count` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`macro_block_count` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromInputBytes(inputBytes int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromOutputBytes(outputBytes int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromStartTime(startTime time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromEndTime(endTime time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromRetryCount 通过retry_count获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromRetryCount(retryCount int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`retry_count` = ?", retryCount).Find(&results).Error

	return
}

// GetBatchFromRetryCount 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromRetryCount(retryCounts []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`retry_count` IN (?)", retryCounts).Find(&results).Error

	return
}

// GetFromReplicaRole 通过replica_role获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromReplicaRole(replicaRole int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`replica_role` = ?", replicaRole).Find(&results).Error

	return
}

// GetBatchFromReplicaRole 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromReplicaRole(replicaRoles []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`replica_role` IN (?)", replicaRoles).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromStatus(status string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromTaskID(taskID int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromResult(result int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromResult(results []int64) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromTraceID 通过trace_id获取内容
func (obj *_AllVirtualPgBackupTaskMgr) GetFromTraceID(traceID string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`trace_id` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualPgBackupTaskMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPgBackupTaskMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, incarnation int64, backupSetID int64) (result AllVirtualPgBackupTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupTask{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `incarnation` = ? AND `backup_set_id` = ?", tenantID, tableID, partitionID, incarnation, backupSetID).First(&result).Error

	return
}
