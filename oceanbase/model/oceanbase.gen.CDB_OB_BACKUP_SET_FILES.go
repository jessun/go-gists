package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CdbObBackupSetFilesMgr struct {
	*_BaseMgr
}

// CdbObBackupSetFilesMgr open func
func CdbObBackupSetFilesMgr(db *gorm.DB) *_CdbObBackupSetFilesMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupSetFilesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupSetFilesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_SET_FILES"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupSetFilesMgr) GetTableName() string {
	return "CDB_OB_BACKUP_SET_FILES"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupSetFilesMgr) Reset() *_CdbObBackupSetFilesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupSetFilesMgr) Get() (result CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupSetFilesMgr) Gets() (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupSetFilesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupSetFilesMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupSetFilesMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupSetFilesMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupSetFilesMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupSetFilesMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithEncryptionMode ENCRYPTION_MODE获取
func (obj *_CdbObBackupSetFilesMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["ENCRYPTION_MODE"] = encryptionMode })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupSetFilesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithFileStatus FILE_STATUS获取
func (obj *_CdbObBackupSetFilesMgr) WithFileStatus(fileStatus string) Option {
	return optionFunc(func(o *options) { o.query["FILE_STATUS"] = fileStatus })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupSetFilesMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupSetFilesMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithElapsedSecondes ELAPSED_SECONDES获取
func (obj *_CdbObBackupSetFilesMgr) WithElapsedSecondes(elapsedSecondes float64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_SECONDES"] = elapsedSecondes })
}

// WithKeep KEEP获取
func (obj *_CdbObBackupSetFilesMgr) WithKeep(keep string) Option {
	return optionFunc(func(o *options) { o.query["KEEP"] = keep })
}

// WithKeepUntil KEEP_UNTIL获取
func (obj *_CdbObBackupSetFilesMgr) WithKeepUntil(keepUntil string) Option {
	return optionFunc(func(o *options) { o.query["KEEP_UNTIL"] = keepUntil })
}

// WithCompressed COMPRESSED获取
func (obj *_CdbObBackupSetFilesMgr) WithCompressed(compressed string) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSED"] = compressed })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupSetFilesMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithOutputRateBytes OUTPUT_RATE_BYTES获取
func (obj *_CdbObBackupSetFilesMgr) WithOutputRateBytes(outputRateBytes float64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES"] = outputRateBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupSetFilesMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetFilesMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// WithOutputRateBytesDisplay OUTPUT_RATE_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetFilesMgr) WithOutputRateBytesDisplay(outputRateBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES_DISPLAY"] = outputRateBytesDisplay })
}

// WithTimeTakenDisplay TIME_TAKEN_DISPLAY获取
func (obj *_CdbObBackupSetFilesMgr) WithTimeTakenDisplay(timeTakenDisplay time.Time) Option {
	return optionFunc(func(o *options) { o.query["TIME_TAKEN_DISPLAY"] = timeTakenDisplay })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupSetFilesMgr) GetByOption(opts ...Option) (result CdbObBackupSetFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupSetFilesMgr) GetByOptions(opts ...Option) (results []*CdbObBackupSetFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupSetFilesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupSetFiles, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where(options.query)
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
func (obj *_CdbObBackupSetFilesMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromBackupType(backupType string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过ENCRYPTION_MODE获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromEncryptionMode(encryptionMode string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`ENCRYPTION_MODE` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`ENCRYPTION_MODE` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromStatus(status string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromFileStatus 通过FILE_STATUS获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromFileStatus(fileStatus string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`FILE_STATUS` = ?", fileStatus).Find(&results).Error

	return
}

// GetBatchFromFileStatus 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromFileStatus(fileStatuss []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`FILE_STATUS` IN (?)", fileStatuss).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromElapsedSecondes 通过ELAPSED_SECONDES获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromElapsedSecondes(elapsedSecondes float64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`ELAPSED_SECONDES` = ?", elapsedSecondes).Find(&results).Error

	return
}

// GetBatchFromElapsedSecondes 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromElapsedSecondes(elapsedSecondess []float64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`ELAPSED_SECONDES` IN (?)", elapsedSecondess).Find(&results).Error

	return
}

// GetFromKeep 通过KEEP获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromKeep(keep string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`KEEP` = ?", keep).Find(&results).Error

	return
}

// GetBatchFromKeep 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromKeep(keeps []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`KEEP` IN (?)", keeps).Find(&results).Error

	return
}

// GetFromKeepUntil 通过KEEP_UNTIL获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromKeepUntil(keepUntil string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`KEEP_UNTIL` = ?", keepUntil).Find(&results).Error

	return
}

// GetBatchFromKeepUntil 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromKeepUntil(keepUntils []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`KEEP_UNTIL` IN (?)", keepUntils).Find(&results).Error

	return
}

// GetFromCompressed 通过COMPRESSED获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromCompressed(compressed string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COMPRESSED` = ?", compressed).Find(&results).Error

	return
}

// GetBatchFromCompressed 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromCompressed(compresseds []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COMPRESSED` IN (?)", compresseds).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromOutputRateBytes 通过OUTPUT_RATE_BYTES获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromOutputRateBytes(outputRateBytes float64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_RATE_BYTES` = ?", outputRateBytes).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytes 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromOutputRateBytes(outputRateBytess []float64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_RATE_BYTES` IN (?)", outputRateBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputRateBytesDisplay 通过OUTPUT_RATE_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromOutputRateBytesDisplay(outputRateBytesDisplay string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` = ?", outputRateBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytesDisplay 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromOutputRateBytesDisplay(outputRateBytesDisplays []string) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` IN (?)", outputRateBytesDisplays).Find(&results).Error

	return
}

// GetFromTimeTakenDisplay 通过TIME_TAKEN_DISPLAY获取内容
func (obj *_CdbObBackupSetFilesMgr) GetFromTimeTakenDisplay(timeTakenDisplay time.Time) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`TIME_TAKEN_DISPLAY` = ?", timeTakenDisplay).Find(&results).Error

	return
}

// GetBatchFromTimeTakenDisplay 批量查找
func (obj *_CdbObBackupSetFilesMgr) GetBatchFromTimeTakenDisplay(timeTakenDisplays []time.Time) (results []*CdbObBackupSetFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetFiles{}).Where("`TIME_TAKEN_DISPLAY` IN (?)", timeTakenDisplays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
