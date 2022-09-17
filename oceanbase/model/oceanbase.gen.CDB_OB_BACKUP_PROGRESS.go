package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupProgressMgr struct {
	*_BaseMgr
}

// CdbObBackupProgressMgr open func
func CdbObBackupProgressMgr(db *gorm.DB) *_CdbObBackupProgressMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupProgressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupProgressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_PROGRESS"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupProgressMgr) GetTableName() string {
	return "CDB_OB_BACKUP_PROGRESS"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupProgressMgr) Reset() *_CdbObBackupProgressMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupProgressMgr) Get() (result CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupProgressMgr) Gets() (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupProgressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupProgressMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupProgressMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupProgressMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupProgressMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithPartitionCount PARTITION_COUNT获取
func (obj *_CdbObBackupProgressMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_COUNT"] = partitionCount })
}

// WithMacroBlockCount MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupProgressMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["MACRO_BLOCK_COUNT"] = macroBlockCount })
}

// WithFinishPartitionCount FINISH_PARTITION_COUNT获取
func (obj *_CdbObBackupProgressMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PARTITION_COUNT"] = finishPartitionCount })
}

// WithFinishMacroBlockCount FINISH_MACRO_BLOCK_COUNT获取
func (obj *_CdbObBackupProgressMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_MACRO_BLOCK_COUNT"] = finishMacroBlockCount })
}

// WithInputBytes INPUT_BYTES获取
func (obj *_CdbObBackupProgressMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES"] = inputBytes })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupProgressMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupProgressMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupProgressMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupProgressMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupProgressMgr) GetByOption(opts ...Option) (result CdbObBackupProgress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupProgressMgr) GetByOptions(opts ...Option) (results []*CdbObBackupProgress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupProgressMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupProgress, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where(options.query)
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

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupProgressMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupProgressMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupProgressMgr) GetFromBackupType(backupType string) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupProgressMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPartitionCount 通过PARTITION_COUNT获取内容
func (obj *_CdbObBackupProgressMgr) GetFromPartitionCount(partitionCount int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`PARTITION_COUNT` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`PARTITION_COUNT` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupProgressMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`MACRO_BLOCK_COUNT` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`MACRO_BLOCK_COUNT` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过FINISH_PARTITION_COUNT获取内容
func (obj *_CdbObBackupProgressMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`FINISH_PARTITION_COUNT` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`FINISH_PARTITION_COUNT` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过FINISH_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObBackupProgressMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`FINISH_MACRO_BLOCK_COUNT` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`FINISH_MACRO_BLOCK_COUNT` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromInputBytes 通过INPUT_BYTES获取内容
func (obj *_CdbObBackupProgressMgr) GetFromInputBytes(inputBytes int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`INPUT_BYTES` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`INPUT_BYTES` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupProgressMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupProgressMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupProgressMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupProgressMgr) GetFromStatus(status string) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupProgressMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupProgress{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
