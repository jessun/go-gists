package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObTenantBackupValidationTaskMgr struct {
	*_BaseMgr
}

// CdbObTenantBackupValidationTaskMgr open func
func CdbObTenantBackupValidationTaskMgr(db *gorm.DB) *_CdbObTenantBackupValidationTaskMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObTenantBackupValidationTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObTenantBackupValidationTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_TENANT_BACKUP_VALIDATION_TASK"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObTenantBackupValidationTaskMgr) GetTableName() string {
	return "CDB_OB_TENANT_BACKUP_VALIDATION_TASK"
}

// Reset 重置gorm会话
func (obj *_CdbObTenantBackupValidationTaskMgr) Reset() *_CdbObTenantBackupValidationTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObTenantBackupValidationTaskMgr) Get() (result CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObTenantBackupValidationTaskMgr) Gets() (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObTenantBackupValidationTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTaskID TASK_ID获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["TASK_ID"] = taskID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupSetID BACKUP_SET_ID获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_SET_ID"] = backupSetID })
}

// WithStatus STATUS获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithStartTime START_TIME获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithTotalPgCount TOTAL_PG_COUNT获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithTotalPgCount(totalPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_PG_COUNT"] = totalPgCount })
}

// WithFinishPgCount FINISH_PG_COUNT获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PG_COUNT"] = finishPgCount })
}

// WithTotalPartitionCount TOTAL_PARTITION_COUNT获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_PARTITION_COUNT"] = totalPartitionCount })
}

// WithFinishPartitionCount FINISH_PARTITION_COUNT获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PARTITION_COUNT"] = finishPartitionCount })
}

// WithTotalMacroBlockCount TOTAL_MACRO_BLOCK_COUNT获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_MACRO_BLOCK_COUNT"] = totalMacroBlockCount })
}

// WithFinishMacroBlockCount FINISH_MACRO_BLOCK_COUNT获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_MACRO_BLOCK_COUNT"] = finishMacroBlockCount })
}

// WithLogSize LOG_SIZE获取
func (obj *_CdbObTenantBackupValidationTaskMgr) WithLogSize(logSize int64) Option {
	return optionFunc(func(o *options) { o.query["LOG_SIZE"] = logSize })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObTenantBackupValidationTaskMgr) GetByOption(opts ...Option) (result CdbObTenantBackupValidationTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObTenantBackupValidationTaskMgr) GetByOptions(opts ...Option) (results []*CdbObTenantBackupValidationTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObTenantBackupValidationTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObTenantBackupValidationTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where(options.query)
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

// GetFromJobID 通过JOB_ID获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromJobID(jobID int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过TASK_ID获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromTaskID(taskID int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TASK_ID` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromTaskID(taskIDs []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TASK_ID` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromTenantID(tenantID int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromIncarnation(incarnation int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过BACKUP_SET_ID获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromBackupSetID(backupSetID int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`BACKUP_SET_ID` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`BACKUP_SET_ID` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromStatus(status string) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromStatus(statuss []string) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromBackupDest(backupDest string) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromStartTime(startTime time.Time) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromEndTime(endTime time.Time) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTotalPgCount 通过TOTAL_PG_COUNT获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromTotalPgCount(totalPgCount int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TOTAL_PG_COUNT` = ?", totalPgCount).Find(&results).Error

	return
}

// GetBatchFromTotalPgCount 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromTotalPgCount(totalPgCounts []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TOTAL_PG_COUNT` IN (?)", totalPgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过FINISH_PG_COUNT获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromFinishPgCount(finishPgCount int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`FINISH_PG_COUNT` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`FINISH_PG_COUNT` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过TOTAL_PARTITION_COUNT获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TOTAL_PARTITION_COUNT` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TOTAL_PARTITION_COUNT` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过FINISH_PARTITION_COUNT获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`FINISH_PARTITION_COUNT` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`FINISH_PARTITION_COUNT` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过TOTAL_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TOTAL_MACRO_BLOCK_COUNT` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`TOTAL_MACRO_BLOCK_COUNT` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过FINISH_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`FINISH_MACRO_BLOCK_COUNT` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`FINISH_MACRO_BLOCK_COUNT` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromLogSize 通过LOG_SIZE获取内容
func (obj *_CdbObTenantBackupValidationTaskMgr) GetFromLogSize(logSize int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`LOG_SIZE` = ?", logSize).Find(&results).Error

	return
}

// GetBatchFromLogSize 批量查找
func (obj *_CdbObTenantBackupValidationTaskMgr) GetBatchFromLogSize(logSizes []int64) (results []*CdbObTenantBackupValidationTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObTenantBackupValidationTask{}).Where("`LOG_SIZE` IN (?)", logSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
