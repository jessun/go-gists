package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupSetExpiredMgr struct {
	*_BaseMgr
}

// CdbObBackupSetExpiredMgr open func
func CdbObBackupSetExpiredMgr(db *gorm.DB) *_CdbObBackupSetExpiredMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupSetExpiredMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupSetExpiredMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_SET_EXPIRED"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupSetExpiredMgr) GetTableName() string {
	return "CDB_OB_BACKUP_SET_EXPIRED"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupSetExpiredMgr) Reset() *_CdbObBackupSetExpiredMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupSetExpiredMgr) Get() (result CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupSetExpiredMgr) Gets() (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupSetExpiredMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupSetExpiredMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupSetExpiredMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupSetExpiredMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupSetExpiredMgr) WithCopyID(copyID string) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupSetExpiredMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithEncryptionMode ENCRYPTION_MODE获取
func (obj *_CdbObBackupSetExpiredMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["ENCRYPTION_MODE"] = encryptionMode })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupSetExpiredMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupSetExpiredMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithElapsedSecondes ELAPSED_SECONDES获取
func (obj *_CdbObBackupSetExpiredMgr) WithElapsedSecondes(elapsedSecondes float64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_SECONDES"] = elapsedSecondes })
}

// WithKeep KEEP获取
func (obj *_CdbObBackupSetExpiredMgr) WithKeep(keep string) Option {
	return optionFunc(func(o *options) { o.query["KEEP"] = keep })
}

// WithKeepUntil KEEP_UNTIL获取
func (obj *_CdbObBackupSetExpiredMgr) WithKeepUntil(keepUntil string) Option {
	return optionFunc(func(o *options) { o.query["KEEP_UNTIL"] = keepUntil })
}

// WithDeviceType DEVICE_TYPE获取
func (obj *_CdbObBackupSetExpiredMgr) WithDeviceType(deviceType string) Option {
	return optionFunc(func(o *options) { o.query["DEVICE_TYPE"] = deviceType })
}

// WithCompressed COMPRESSED获取
func (obj *_CdbObBackupSetExpiredMgr) WithCompressed(compressed string) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSED"] = compressed })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupSetExpiredMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithOutputRateBytes OUTPUT_RATE_BYTES获取
func (obj *_CdbObBackupSetExpiredMgr) WithOutputRateBytes(outputRateBytes float64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES"] = outputRateBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupSetExpiredMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetExpiredMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// WithOutputRateBytesDisplay OUTPUT_RATE_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetExpiredMgr) WithOutputRateBytesDisplay(outputRateBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES_DISPLAY"] = outputRateBytesDisplay })
}

// WithTimeTakenDisplay TIME_TAKEN_DISPLAY获取
func (obj *_CdbObBackupSetExpiredMgr) WithTimeTakenDisplay(timeTakenDisplay time.Time) Option {
	return optionFunc(func(o *options) { o.query["TIME_TAKEN_DISPLAY"] = timeTakenDisplay })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupSetExpiredMgr) GetByOption(opts ...Option) (result CdbObBackupSetExpired, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupSetExpiredMgr) GetByOptions(opts ...Option) (results []*CdbObBackupSetExpired, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupSetExpiredMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupSetExpired, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where(options.query)
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
func (obj *_CdbObBackupSetExpiredMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromCopyID(copyID string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromCopyID(copyIDs []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromBackupType(backupType string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过ENCRYPTION_MODE获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromEncryptionMode(encryptionMode string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`ENCRYPTION_MODE` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`ENCRYPTION_MODE` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromElapsedSecondes 通过ELAPSED_SECONDES获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromElapsedSecondes(elapsedSecondes float64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`ELAPSED_SECONDES` = ?", elapsedSecondes).Find(&results).Error

	return
}

// GetBatchFromElapsedSecondes 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromElapsedSecondes(elapsedSecondess []float64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`ELAPSED_SECONDES` IN (?)", elapsedSecondess).Find(&results).Error

	return
}

// GetFromKeep 通过KEEP获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromKeep(keep string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`KEEP` = ?", keep).Find(&results).Error

	return
}

// GetBatchFromKeep 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromKeep(keeps []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`KEEP` IN (?)", keeps).Find(&results).Error

	return
}

// GetFromKeepUntil 通过KEEP_UNTIL获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromKeepUntil(keepUntil string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`KEEP_UNTIL` = ?", keepUntil).Find(&results).Error

	return
}

// GetBatchFromKeepUntil 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromKeepUntil(keepUntils []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`KEEP_UNTIL` IN (?)", keepUntils).Find(&results).Error

	return
}

// GetFromDeviceType 通过DEVICE_TYPE获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromDeviceType(deviceType string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`DEVICE_TYPE` = ?", deviceType).Find(&results).Error

	return
}

// GetBatchFromDeviceType 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromDeviceType(deviceTypes []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`DEVICE_TYPE` IN (?)", deviceTypes).Find(&results).Error

	return
}

// GetFromCompressed 通过COMPRESSED获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromCompressed(compressed string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COMPRESSED` = ?", compressed).Find(&results).Error

	return
}

// GetBatchFromCompressed 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromCompressed(compresseds []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COMPRESSED` IN (?)", compresseds).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromOutputRateBytes 通过OUTPUT_RATE_BYTES获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromOutputRateBytes(outputRateBytes float64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_RATE_BYTES` = ?", outputRateBytes).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytes 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromOutputRateBytes(outputRateBytess []float64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_RATE_BYTES` IN (?)", outputRateBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputRateBytesDisplay 通过OUTPUT_RATE_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromOutputRateBytesDisplay(outputRateBytesDisplay string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` = ?", outputRateBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytesDisplay 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromOutputRateBytesDisplay(outputRateBytesDisplays []string) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` IN (?)", outputRateBytesDisplays).Find(&results).Error

	return
}

// GetFromTimeTakenDisplay 通过TIME_TAKEN_DISPLAY获取内容
func (obj *_CdbObBackupSetExpiredMgr) GetFromTimeTakenDisplay(timeTakenDisplay time.Time) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`TIME_TAKEN_DISPLAY` = ?", timeTakenDisplay).Find(&results).Error

	return
}

// GetBatchFromTimeTakenDisplay 批量查找
func (obj *_CdbObBackupSetExpiredMgr) GetBatchFromTimeTakenDisplay(timeTakenDisplays []time.Time) (results []*CdbObBackupSetExpired, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetExpired{}).Where("`TIME_TAKEN_DISPLAY` IN (?)", timeTakenDisplays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
