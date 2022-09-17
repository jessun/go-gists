package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupBackupsetTaskMgr struct {
	*_BaseMgr
}

// CdbObBackupBackupsetTaskMgr open func
func CdbObBackupBackupsetTaskMgr(db *gorm.DB) *_CdbObBackupBackupsetTaskMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackupsetTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackupsetTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPSET_TASK"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackupsetTaskMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPSET_TASK"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackupsetTaskMgr) Reset() *_CdbObBackupBackupsetTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackupsetTaskMgr) Get() (result CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackupsetTaskMgr) Gets() (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackupsetTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTotalPgCount TOTAL_PG_COUNT获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithTotalPgCount(totalPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_PG_COUNT"] = totalPgCount })
}

// WithFinishPgCount FINISH_PG_COUNT获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithFinishPgCount(finishPgCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PG_COUNT"] = finishPgCount })
}

// WithTotalPartitionCount TOTAL_PARTITION_COUNT获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_PARTITION_COUNT"] = totalPartitionCount })
}

// WithTotalMacroBlockCount TOTAL_MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_MACRO_BLOCK_COUNT"] = totalMacroBlockCount })
}

// WithFinishPartitionCount FINISH_PARTITION_COUNT获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PARTITION_COUNT"] = finishPartitionCount })
}

// WithFinishMacroBlockCount FINISH_MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_MACRO_BLOCK_COUNT"] = finishMacroBlockCount })
}

// WithInputBytes INPUT_BYTES获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES"] = inputBytes })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackupsetTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackupsetTaskMgr) GetByOption(opts ...Option) (result CdbObBackupBackupsetTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackupsetTaskMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackupsetTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackupsetTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackupsetTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where(options.query)
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
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromBackupType(backupType string) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTotalPgCount 通过TOTAL_PG_COUNT获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromTotalPgCount(totalPgCount int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TOTAL_PG_COUNT` = ?", totalPgCount).Find(&results).Error

	return
}

// GetBatchFromTotalPgCount 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromTotalPgCount(totalPgCounts []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TOTAL_PG_COUNT` IN (?)", totalPgCounts).Find(&results).Error

	return
}

// GetFromFinishPgCount 通过FINISH_PG_COUNT获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromFinishPgCount(finishPgCount int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`FINISH_PG_COUNT` = ?", finishPgCount).Find(&results).Error

	return
}

// GetBatchFromFinishPgCount 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromFinishPgCount(finishPgCounts []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`FINISH_PG_COUNT` IN (?)", finishPgCounts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过TOTAL_PARTITION_COUNT获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TOTAL_PARTITION_COUNT` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TOTAL_PARTITION_COUNT` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过TOTAL_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TOTAL_MACRO_BLOCK_COUNT` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`TOTAL_MACRO_BLOCK_COUNT` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过FINISH_PARTITION_COUNT获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`FINISH_PARTITION_COUNT` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`FINISH_PARTITION_COUNT` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过FINISH_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`FINISH_MACRO_BLOCK_COUNT` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`FINISH_MACRO_BLOCK_COUNT` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过INPUT_BYTES获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromInputBytes(inputBytes int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`INPUT_BYTES` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`INPUT_BYTES` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackupsetTaskMgr) GetFromStatus(status string) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackupsetTaskMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTask{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
