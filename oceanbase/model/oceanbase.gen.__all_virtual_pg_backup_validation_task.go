package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualPgBackupValidationTaskMgr struct {
	*_BaseMgr
}

// AllVirtualPgBackupValidationTaskMgr open func
func AllVirtualPgBackupValidationTaskMgr(db *gorm.DB) *_AllVirtualPgBackupValidationTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPgBackupValidationTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPgBackupValidationTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_pg_backup_validation_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetTableName() string {
	return "__all_virtual_pg_backup_validation_task"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPgBackupValidationTaskMgr) Reset() *_AllVirtualPgBackupValidationTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) Get() (result AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPgBackupValidationTaskMgr) Gets() (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPgBackupValidationTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTaskID task_id获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithArchiveRound archive_round获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithArchiveRound(archiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["archive_round"] = archiveRound })
}

// WithStatus status获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTraceID trace_id获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTotalPartitionCount total_partition_count获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_partition_count"] = totalPartitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithTotalMacroBlockCount total_macro_block_count获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_macro_block_count"] = totalMacroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithLogInfo log_info获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithLogInfo(logInfo string) Option {
	return optionFunc(func(o *options) { o.query["log_info"] = logInfo })
}

// WithLogSize log_size获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithLogSize(logSize int64) Option {
	return optionFunc(func(o *options) { o.query["log_size"] = logSize })
}

// WithResult result获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithComment comment获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetByOption(opts ...Option) (result AllVirtualPgBackupValidationTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetByOptions(opts ...Option) (results []*AllVirtualPgBackupValidationTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPgBackupValidationTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPgBackupValidationTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where(options.query)
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
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromJobID(jobID int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromTaskID(taskID int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromBackupSetID(backupSetID int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromTableID(tableID int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromArchiveRound 通过archive_round获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromArchiveRound(archiveRound int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`archive_round` = ?", archiveRound).Find(&results).Error

	return
}

// GetBatchFromArchiveRound 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromArchiveRound(archiveRounds []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`archive_round` IN (?)", archiveRounds).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromStatus(status string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTraceID 通过trace_id获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromTraceID(traceID string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`trace_id` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过total_partition_count获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`total_partition_count` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`total_partition_count` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过total_macro_block_count获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`total_macro_block_count` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`total_macro_block_count` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromLogInfo 通过log_info获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromLogInfo(logInfo string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`log_info` = ?", logInfo).Find(&results).Error

	return
}

// GetBatchFromLogInfo 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromLogInfo(logInfos []string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`log_info` IN (?)", logInfos).Find(&results).Error

	return
}

// GetFromLogSize 通过log_size获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromLogSize(logSize int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`log_size` = ?", logSize).Find(&results).Error

	return
}

// GetBatchFromLogSize 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromLogSize(logSizes []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`log_size` IN (?)", logSizes).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromResult(result int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromResult(results []int64) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetFromComment(comment string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualPgBackupValidationTaskMgr) GetBatchFromComment(comments []string) (results []*AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPgBackupValidationTaskMgr) FetchByPrimaryKey(tenantID int64, jobID int64, taskID int64, incarnation int64, backupSetID int64, tableID int64, partitionID int64) (result AllVirtualPgBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupValidationTask{}).Where("`tenant_id` = ? AND `job_id` = ? AND `task_id` = ? AND `incarnation` = ? AND `backup_set_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, jobID, taskID, incarnation, backupSetID, tableID, partitionID).First(&result).Error

	return
}
