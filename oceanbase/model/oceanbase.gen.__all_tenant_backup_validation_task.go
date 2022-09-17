package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantBackupValidationTaskMgr struct {
	*_BaseMgr
}

// AllTenantBackupValidationTaskMgr open func
func AllTenantBackupValidationTaskMgr(db *gorm.DB) *_AllTenantBackupValidationTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantBackupValidationTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantBackupValidationTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_backup_validation_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantBackupValidationTaskMgr) GetTableName() string {
	return "__all_tenant_backup_validation_task"
}

// Reset 重置gorm会话
func (obj *_AllTenantBackupValidationTaskMgr) Reset() *_AllTenantBackupValidationTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantBackupValidationTaskMgr) Get() (result AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantBackupValidationTaskMgr) Gets() (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantBackupValidationTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantBackupValidationTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantBackupValidationTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantBackupValidationTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllTenantBackupValidationTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTaskID task_id获取
func (obj *_AllTenantBackupValidationTaskMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithIncarnation incarnation获取
func (obj *_AllTenantBackupValidationTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllTenantBackupValidationTaskMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithStatus status获取
func (obj *_AllTenantBackupValidationTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBackupDest backup_dest获取
func (obj *_AllTenantBackupValidationTaskMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithStartTime start_time获取
func (obj *_AllTenantBackupValidationTaskMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllTenantBackupValidationTaskMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithTotalPgCount total_pg_count获取
func (obj *_AllTenantBackupValidationTaskMgr) WithTotalPgCount(totalPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_pg_count"] = totalPgCount })
}

// WithFinishPgCount finish_pg_count获取
func (obj *_AllTenantBackupValidationTaskMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_pg_count"] = finishPgCount })
}

// WithTotalPartitionCount total_partition_count获取
func (obj *_AllTenantBackupValidationTaskMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_partition_count"] = totalPartitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllTenantBackupValidationTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithTotalMacroBlockCount total_macro_block_count获取
func (obj *_AllTenantBackupValidationTaskMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_macro_block_count"] = totalMacroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllTenantBackupValidationTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithLogSize log_size获取
func (obj *_AllTenantBackupValidationTaskMgr) WithLogSize(logSize int64) Option {
	return optionFunc(func(o *options) { o.query["log_size"] = logSize })
}

// WithResult result获取
func (obj *_AllTenantBackupValidationTaskMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithComment comment获取
func (obj *_AllTenantBackupValidationTaskMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantBackupValidationTaskMgr) GetByOption(opts ...Option) (result AllTenantBackupValidationTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantBackupValidationTaskMgr) GetByOptions(opts ...Option) (results []*AllTenantBackupValidationTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantBackupValidationTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantBackupValidationTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where(options.query)
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
func (obj *_AllTenantBackupValidationTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromTenantID(tenantID int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromJobID(jobID int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromTaskID(taskID int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromIncarnation(incarnation int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromBackupSetID(backupSetID int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromStatus(status string) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromBackupDest(backupDest string) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromStartTime(startTime time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromEndTime(endTime time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTotalPgCount 通过total_pg_count获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromTotalPgCount(totalPgCount int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`total_pg_count` = ?", totalPgCount).Find(&results).Error

	return
}

// GetBatchFromTotalPgCount 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromTotalPgCount(totalPgCounts []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`total_pg_count` IN (?)", totalPgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过finish_pg_count获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromFinishPgCount(finishPgCount int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`finish_pg_count` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`finish_pg_count` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过total_partition_count获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`total_partition_count` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`total_partition_count` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过total_macro_block_count获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`total_macro_block_count` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`total_macro_block_count` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromLogSize 通过log_size获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromLogSize(logSize int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`log_size` = ?", logSize).Find(&results).Error

	return
}

// GetBatchFromLogSize 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromLogSize(logSizes []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`log_size` IN (?)", logSizes).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromResult(result int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromResult(results []int64) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllTenantBackupValidationTaskMgr) GetFromComment(comment string) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllTenantBackupValidationTaskMgr) GetBatchFromComment(comments []string) (results []*AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantBackupValidationTaskMgr) FetchByPrimaryKey(tenantID int64, jobID int64, taskID int64, incarnation int64, backupSetID int64) (result AllTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupValidationTask{}).Where("`tenant_id` = ? AND `job_id` = ? AND `task_id` = ? AND `incarnation` = ? AND `backup_set_id` = ?", tenantID, jobID, taskID, incarnation, backupSetID).First(&result).Error

	return
}
