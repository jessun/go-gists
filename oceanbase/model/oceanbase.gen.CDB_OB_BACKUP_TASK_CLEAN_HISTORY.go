package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupTaskCleanHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupTaskCleanHistoryMgr open func
func CdbObBackupTaskCleanHistoryMgr(db *gorm.DB) *_CdbObBackupTaskCleanHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupTaskCleanHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupTaskCleanHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_TASK_CLEAN_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_TASK_CLEAN_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupTaskCleanHistoryMgr) Reset() *_CdbObBackupTaskCleanHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) Get() (result CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupTaskCleanHistoryMgr) Gets() (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupTaskCleanHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithPartitionCount PARTITION_COUNT获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_COUNT"] = partitionCount })
}

// WithMacroBlockCount MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["MACRO_BLOCK_COUNT"] = macroBlockCount })
}

// WithFinishPartitionCount FINISH_PARTITION_COUNT获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PARTITION_COUNT"] = finishPartitionCount })
}

// WithFinishMacroBlockCount FINISH_MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_MACRO_BLOCK_COUNT"] = finishMacroBlockCount })
}

// WithInputBytes INPUT_BYTES获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES"] = inputBytes })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupTaskCleanHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupTaskCleanHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupTaskCleanHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupTaskCleanHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where(options.query)
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

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromBackupType(backupType string) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromPartitionCount 通过PARTITION_COUNT获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromPartitionCount(partitionCount int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`PARTITION_COUNT` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`PARTITION_COUNT` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`MACRO_BLOCK_COUNT` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`MACRO_BLOCK_COUNT` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过FINISH_PARTITION_COUNT获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`FINISH_PARTITION_COUNT` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`FINISH_PARTITION_COUNT` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过FINISH_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`FINISH_MACRO_BLOCK_COUNT` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`FINISH_MACRO_BLOCK_COUNT` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过INPUT_BYTES获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromInputBytes(inputBytes int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`INPUT_BYTES` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`INPUT_BYTES` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupTaskCleanHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupTaskCleanHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupTaskCleanHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
