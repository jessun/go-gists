package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObRestoreHistoryMgr struct {
	*_BaseMgr
}

// CdbObRestoreHistoryMgr open func
func CdbObRestoreHistoryMgr(db *gorm.DB) *_CdbObRestoreHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObRestoreHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObRestoreHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_RESTORE_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObRestoreHistoryMgr) GetTableName() string {
	return "CDB_OB_RESTORE_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObRestoreHistoryMgr) Reset() *_CdbObRestoreHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObRestoreHistoryMgr) Get() (result CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObRestoreHistoryMgr) Gets() (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObRestoreHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObRestoreHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithExternalJobID EXTERNAL_JOB_ID获取
func (obj *_CdbObRestoreHistoryMgr) WithExternalJobID(externalJobID int64) Option {
	return optionFunc(func(o *options) { o.query["EXTERNAL_JOB_ID"] = externalJobID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObRestoreHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTenantName TENANT_NAME获取
func (obj *_CdbObRestoreHistoryMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_NAME"] = tenantName })
}

// WithBackupTenantID BACKUP_TENANT_ID获取
func (obj *_CdbObRestoreHistoryMgr) WithBackupTenantID(backupTenantID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TENANT_ID"] = backupTenantID })
}

// WithBackupTenantName BACKUP_TENANT_NAME获取
func (obj *_CdbObRestoreHistoryMgr) WithBackupTenantName(backupTenantName string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TENANT_NAME"] = backupTenantName })
}

// WithBackupClusterID BACKUP_CLUSTER_ID获取
func (obj *_CdbObRestoreHistoryMgr) WithBackupClusterID(backupClusterID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_CLUSTER_ID"] = backupClusterID })
}

// WithBackupClusterName BACKUP_CLUSTER_NAME获取
func (obj *_CdbObRestoreHistoryMgr) WithBackupClusterName(backupClusterName string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_CLUSTER_NAME"] = backupClusterName })
}

// WithWhiteList WHITE_LIST获取
func (obj *_CdbObRestoreHistoryMgr) WithWhiteList(whiteList string) Option {
	return optionFunc(func(o *options) { o.query["WHITE_LIST"] = whiteList })
}

// WithStatus STATUS获取
func (obj *_CdbObRestoreHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithStartTime START_TIME获取
func (obj *_CdbObRestoreHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObRestoreHistoryMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithPartitionCount PARTITION_COUNT获取
func (obj *_CdbObRestoreHistoryMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_COUNT"] = partitionCount })
}

// WithMacroBlockCount MACRO_BLOCK_COUNT获取
func (obj *_CdbObRestoreHistoryMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["MACRO_BLOCK_COUNT"] = macroBlockCount })
}

// WithFinishPartitionCount FINISH_PARTITION_COUNT获取
func (obj *_CdbObRestoreHistoryMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_PARTITION_COUNT"] = finishPartitionCount })
}

// WithFinishMacroBlockCount FINISH_MACRO_BLOCK_COUNT获取
func (obj *_CdbObRestoreHistoryMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["FINISH_MACRO_BLOCK_COUNT"] = finishMacroBlockCount })
}

// WithRestoreStartTimestamp RESTORE_START_TIMESTAMP获取
func (obj *_CdbObRestoreHistoryMgr) WithRestoreStartTimestamp(restoreStartTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["RESTORE_START_TIMESTAMP"] = restoreStartTimestamp })
}

// WithRestoreFinishTimestamp RESTORE_FINISH_TIMESTAMP获取
func (obj *_CdbObRestoreHistoryMgr) WithRestoreFinishTimestamp(restoreFinishTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["RESTORE_FINISH_TIMESTAMP"] = restoreFinishTimestamp })
}

// WithRestoreCurrentTimestamp RESTORE_CURRENT_TIMESTAMP获取
func (obj *_CdbObRestoreHistoryMgr) WithRestoreCurrentTimestamp(restoreCurrentTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["RESTORE_CURRENT_TIMESTAMP"] = restoreCurrentTimestamp })
}

// WithBackupSetList BACKUP_SET_LIST获取
func (obj *_CdbObRestoreHistoryMgr) WithBackupSetList(backupSetList string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_SET_LIST"] = backupSetList })
}

// WithBackupPieceList BACKUP_PIECE_LIST获取
func (obj *_CdbObRestoreHistoryMgr) WithBackupPieceList(backupPieceList string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_LIST"] = backupPieceList })
}

// WithInfo INFO获取
func (obj *_CdbObRestoreHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["INFO"] = info })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObRestoreHistoryMgr) GetByOption(opts ...Option) (result CdbObRestoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObRestoreHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObRestoreHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObRestoreHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObRestoreHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where(options.query)
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
func (obj *_CdbObRestoreHistoryMgr) GetFromJobID(jobID int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromExternalJobID 通过EXTERNAL_JOB_ID获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromExternalJobID(externalJobID int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`EXTERNAL_JOB_ID` = ?", externalJobID).Find(&results).Error

	return
}

// GetBatchFromExternalJobID 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromExternalJobID(externalJobIDs []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`EXTERNAL_JOB_ID` IN (?)", externalJobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过TENANT_NAME获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromTenantName(tenantName string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`TENANT_NAME` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromTenantName(tenantNames []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`TENANT_NAME` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromBackupTenantID 通过BACKUP_TENANT_ID获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromBackupTenantID(backupTenantID int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_TENANT_ID` = ?", backupTenantID).Find(&results).Error

	return
}

// GetBatchFromBackupTenantID 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromBackupTenantID(backupTenantIDs []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_TENANT_ID` IN (?)", backupTenantIDs).Find(&results).Error

	return
}

// GetFromBackupTenantName 通过BACKUP_TENANT_NAME获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromBackupTenantName(backupTenantName string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_TENANT_NAME` = ?", backupTenantName).Find(&results).Error

	return
}

// GetBatchFromBackupTenantName 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromBackupTenantName(backupTenantNames []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_TENANT_NAME` IN (?)", backupTenantNames).Find(&results).Error

	return
}

// GetFromBackupClusterID 通过BACKUP_CLUSTER_ID获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromBackupClusterID(backupClusterID int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_CLUSTER_ID` = ?", backupClusterID).Find(&results).Error

	return
}

// GetBatchFromBackupClusterID 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromBackupClusterID(backupClusterIDs []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_CLUSTER_ID` IN (?)", backupClusterIDs).Find(&results).Error

	return
}

// GetFromBackupClusterName 通过BACKUP_CLUSTER_NAME获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromBackupClusterName(backupClusterName string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_CLUSTER_NAME` = ?", backupClusterName).Find(&results).Error

	return
}

// GetBatchFromBackupClusterName 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromBackupClusterName(backupClusterNames []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_CLUSTER_NAME` IN (?)", backupClusterNames).Find(&results).Error

	return
}

// GetFromWhiteList 通过WHITE_LIST获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromWhiteList(whiteList string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`WHITE_LIST` = ?", whiteList).Find(&results).Error

	return
}

// GetBatchFromWhiteList 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromWhiteList(whiteLists []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`WHITE_LIST` IN (?)", whiteLists).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromStatus(status string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromPartitionCount 通过PARTITION_COUNT获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromPartitionCount(partitionCount int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`PARTITION_COUNT` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`PARTITION_COUNT` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromMacroBlockCount 通过MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`MACRO_BLOCK_COUNT` = ?", macroBlockCount).Find(&results).Error

	return
}

// GetBatchFromMacroBlockCount 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`MACRO_BLOCK_COUNT` IN (?)", macroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过FINISH_PARTITION_COUNT获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`FINISH_PARTITION_COUNT` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`FINISH_PARTITION_COUNT` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过FINISH_MACRO_BLOCK_COUNT获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`FINISH_MACRO_BLOCK_COUNT` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`FINISH_MACRO_BLOCK_COUNT` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromRestoreStartTimestamp 通过RESTORE_START_TIMESTAMP获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromRestoreStartTimestamp(restoreStartTimestamp time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`RESTORE_START_TIMESTAMP` = ?", restoreStartTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreStartTimestamp 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromRestoreStartTimestamp(restoreStartTimestamps []time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`RESTORE_START_TIMESTAMP` IN (?)", restoreStartTimestamps).Find(&results).Error

	return
}

// GetFromRestoreFinishTimestamp 通过RESTORE_FINISH_TIMESTAMP获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromRestoreFinishTimestamp(restoreFinishTimestamp time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`RESTORE_FINISH_TIMESTAMP` = ?", restoreFinishTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreFinishTimestamp 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromRestoreFinishTimestamp(restoreFinishTimestamps []time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`RESTORE_FINISH_TIMESTAMP` IN (?)", restoreFinishTimestamps).Find(&results).Error

	return
}

// GetFromRestoreCurrentTimestamp 通过RESTORE_CURRENT_TIMESTAMP获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromRestoreCurrentTimestamp(restoreCurrentTimestamp time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`RESTORE_CURRENT_TIMESTAMP` = ?", restoreCurrentTimestamp).Find(&results).Error

	return
}

// GetBatchFromRestoreCurrentTimestamp 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromRestoreCurrentTimestamp(restoreCurrentTimestamps []time.Time) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`RESTORE_CURRENT_TIMESTAMP` IN (?)", restoreCurrentTimestamps).Find(&results).Error

	return
}

// GetFromBackupSetList 通过BACKUP_SET_LIST获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromBackupSetList(backupSetList string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_SET_LIST` = ?", backupSetList).Find(&results).Error

	return
}

// GetBatchFromBackupSetList 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromBackupSetList(backupSetLists []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_SET_LIST` IN (?)", backupSetLists).Find(&results).Error

	return
}

// GetFromBackupPieceList 通过BACKUP_PIECE_LIST获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromBackupPieceList(backupPieceList string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_PIECE_LIST` = ?", backupPieceList).Find(&results).Error

	return
}

// GetBatchFromBackupPieceList 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromBackupPieceList(backupPieceLists []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`BACKUP_PIECE_LIST` IN (?)", backupPieceLists).Find(&results).Error

	return
}

// GetFromInfo 通过INFO获取内容
func (obj *_CdbObRestoreHistoryMgr) GetFromInfo(info string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`INFO` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_CdbObRestoreHistoryMgr) GetBatchFromInfo(infos []string) (results []*CdbObRestoreHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObRestoreHistory{}).Where("`INFO` IN (?)", infos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
