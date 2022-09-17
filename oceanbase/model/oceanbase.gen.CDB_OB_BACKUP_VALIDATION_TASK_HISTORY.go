package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CdbObBackupValidationTaskHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupValidationTaskHistoryMgr open func
func CdbObBackupValidationTaskHistoryMgr(db *gorm.DB) *_CdbObBackupValidationTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupValidationTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupValidationTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_VALIDATION_TASK_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_VALIDATION_TASK_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupValidationTaskHistoryMgr) Reset() *_CdbObBackupValidationTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) Get() (result CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupValidationTaskHistoryMgr) Gets() (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupValidationTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTaskID TASK_ID获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["TASK_ID"] = taskID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupSetID BACKUP_SET_ID获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_SET_ID"] = backupSetID })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithTotalPgCount TOTAL_PG_COUNT获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithTotalPgCount(totalPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_PG_COUNT"] = totalPgCount })
}

// WithFinishPgCount FINISH_PG_COUNT获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PG_COUNT"] = finishPgCount })
}

// WithTotalPartitionCount TOTAL_PARTITION_COUNT获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_PARTITION_COUNT"] = totalPartitionCount })
}

// WithFinishPartitionCount FINISH_PARTITION_COUNT获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PARTITION_COUNT"] = finishPartitionCount })
}

// WithTotalMacroBlockCount TOTAL_MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_MACRO_BLOCK_COUNT"] = totalMacroBlockCount })
}

// WithFinishMacroBlockCount FINISH_MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_MACRO_BLOCK_COUNT"] = finishMacroBlockCount })
}

// WithLogSize LOG_SIZE获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) WithLogSize(logSize int64) Option {
	return optionFunc(func(o *options) { o.query["LOG_SIZE"] = logSize })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupValidationTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupValidationTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupValidationTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupValidationTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where(options.query)
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
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromJobID(jobID int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过TASK_ID获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromTaskID(taskID int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TASK_ID` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromTaskID(taskIDs []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TASK_ID` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过BACKUP_SET_ID获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromBackupSetID(backupSetID int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`BACKUP_SET_ID` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`BACKUP_SET_ID` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromBackupDest(backupDest string) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromEndTime(endTime time.Time) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTotalPgCount 通过TOTAL_PG_COUNT获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromTotalPgCount(totalPgCount int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TOTAL_PG_COUNT` = ?", totalPgCount).Find(&results).Error

	return
}

// GetBatchFromTotalPgCount 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromTotalPgCount(totalPgCounts []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TOTAL_PG_COUNT` IN (?)", totalPgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过FINISH_PG_COUNT获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromFinishPgCount(finishPgCount int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`FINISH_PG_COUNT` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`FINISH_PG_COUNT` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过TOTAL_PARTITION_COUNT获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TOTAL_PARTITION_COUNT` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TOTAL_PARTITION_COUNT` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过FINISH_PARTITION_COUNT获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`FINISH_PARTITION_COUNT` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`FINISH_PARTITION_COUNT` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过TOTAL_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TOTAL_MACRO_BLOCK_COUNT` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`TOTAL_MACRO_BLOCK_COUNT` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过FINISH_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`FINISH_MACRO_BLOCK_COUNT` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`FINISH_MACRO_BLOCK_COUNT` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromLogSize 通过LOG_SIZE获取内容
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetFromLogSize(logSize int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`LOG_SIZE` = ?", logSize).Find(&results).Error

	return
}

// GetBatchFromLogSize 批量查找
func (obj *_CdbObBackupValidationTaskHistoryMgr) GetBatchFromLogSize(logSizes []int64) (results []*CdbObBackupValidationTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationTaskHistory{}).Where("`LOG_SIZE` IN (?)", logSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
