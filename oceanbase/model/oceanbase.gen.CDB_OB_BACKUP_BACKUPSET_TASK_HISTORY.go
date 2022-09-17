package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CdbObBackupBackupsetTaskHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupBackupsetTaskHistoryMgr open func
func CdbObBackupBackupsetTaskHistoryMgr(db *gorm.DB) *_CdbObBackupBackupsetTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackupsetTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackupsetTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPSET_TASK_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPSET_TASK_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) Reset() *_CdbObBackupBackupsetTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) Get() (result CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) Gets() (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithEncryptionMode ENCRYPTION_MODE获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["ENCRYPTION_MODE"] = encryptionMode })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithElapsedSecondes ELAPSED_SECONDES获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithElapsedSecondes(elapsedSecondes float64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_SECONDES"] = elapsedSecondes })
}

// WithKeep KEEP获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithKeep(keep string) Option {
	return optionFunc(func(o *options) { o.query["KEEP"] = keep })
}

// WithKeepUntil KEEP_UNTIL获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithKeepUntil(keepUntil string) Option {
	return optionFunc(func(o *options) { o.query["KEEP_UNTIL"] = keepUntil })
}

// WithSrcDeviceType SRC_DEVICE_TYPE获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithSrcDeviceType(srcDeviceType string) Option {
	return optionFunc(func(o *options) { o.query["SRC_DEVICE_TYPE"] = srcDeviceType })
}

// WithDstDeviceType DST_DEVICE_TYPE获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithDstDeviceType(dstDeviceType string) Option {
	return optionFunc(func(o *options) { o.query["DST_DEVICE_TYPE"] = dstDeviceType })
}

// WithCompressed COMPRESSED获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithCompressed(compressed string) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSED"] = compressed })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithOutputRateBytes OUTPUT_RATE_BYTES获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithOutputRateBytes(outputRateBytes float64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES"] = outputRateBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// WithOutputRateBytesDisplay OUTPUT_RATE_BYTES_DISPLAY获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithOutputRateBytesDisplay(outputRateBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES_DISPLAY"] = outputRateBytesDisplay })
}

// WithTimeTakenDisplay TIME_TAKEN_DISPLAY获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithTimeTakenDisplay(timeTakenDisplay time.Time) Option {
	return optionFunc(func(o *options) { o.query["TIME_TAKEN_DISPLAY"] = timeTakenDisplay })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupBackupsetTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackupsetTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where(options.query)
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
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromBackupType(backupType string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过ENCRYPTION_MODE获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromEncryptionMode(encryptionMode string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`ENCRYPTION_MODE` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`ENCRYPTION_MODE` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromElapsedSecondes 通过ELAPSED_SECONDES获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromElapsedSecondes(elapsedSecondes float64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`ELAPSED_SECONDES` = ?", elapsedSecondes).Find(&results).Error

	return
}

// GetBatchFromElapsedSecondes 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromElapsedSecondes(elapsedSecondess []float64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`ELAPSED_SECONDES` IN (?)", elapsedSecondess).Find(&results).Error

	return
}

// GetFromKeep 通过KEEP获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromKeep(keep string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`KEEP` = ?", keep).Find(&results).Error

	return
}

// GetBatchFromKeep 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromKeep(keeps []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`KEEP` IN (?)", keeps).Find(&results).Error

	return
}

// GetFromKeepUntil 通过KEEP_UNTIL获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromKeepUntil(keepUntil string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`KEEP_UNTIL` = ?", keepUntil).Find(&results).Error

	return
}

// GetBatchFromKeepUntil 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromKeepUntil(keepUntils []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`KEEP_UNTIL` IN (?)", keepUntils).Find(&results).Error

	return
}

// GetFromSrcDeviceType 通过SRC_DEVICE_TYPE获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromSrcDeviceType(srcDeviceType string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`SRC_DEVICE_TYPE` = ?", srcDeviceType).Find(&results).Error

	return
}

// GetBatchFromSrcDeviceType 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromSrcDeviceType(srcDeviceTypes []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`SRC_DEVICE_TYPE` IN (?)", srcDeviceTypes).Find(&results).Error

	return
}

// GetFromDstDeviceType 通过DST_DEVICE_TYPE获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromDstDeviceType(dstDeviceType string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`DST_DEVICE_TYPE` = ?", dstDeviceType).Find(&results).Error

	return
}

// GetBatchFromDstDeviceType 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromDstDeviceType(dstDeviceTypes []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`DST_DEVICE_TYPE` IN (?)", dstDeviceTypes).Find(&results).Error

	return
}

// GetFromCompressed 通过COMPRESSED获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromCompressed(compressed string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COMPRESSED` = ?", compressed).Find(&results).Error

	return
}

// GetBatchFromCompressed 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromCompressed(compresseds []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COMPRESSED` IN (?)", compresseds).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromOutputRateBytes 通过OUTPUT_RATE_BYTES获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromOutputRateBytes(outputRateBytes float64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_RATE_BYTES` = ?", outputRateBytes).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytes 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromOutputRateBytes(outputRateBytess []float64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_RATE_BYTES` IN (?)", outputRateBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputRateBytesDisplay 通过OUTPUT_RATE_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromOutputRateBytesDisplay(outputRateBytesDisplay string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` = ?", outputRateBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytesDisplay 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromOutputRateBytesDisplay(outputRateBytesDisplays []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` IN (?)", outputRateBytesDisplays).Find(&results).Error

	return
}

// GetFromTimeTakenDisplay 通过TIME_TAKEN_DISPLAY获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromTimeTakenDisplay(timeTakenDisplay time.Time) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`TIME_TAKEN_DISPLAY` = ?", timeTakenDisplay).Find(&results).Error

	return
}

// GetBatchFromTimeTakenDisplay 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromTimeTakenDisplay(timeTakenDisplays []time.Time) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`TIME_TAKEN_DISPLAY` IN (?)", timeTakenDisplays).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackupsetTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackupsetTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetTaskHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
