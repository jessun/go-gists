package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupSetDetailsMgr struct {
	*_BaseMgr
}

// CdbObBackupSetDetailsMgr open func
func CdbObBackupSetDetailsMgr(db *gorm.DB) *_CdbObBackupSetDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupSetDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupSetDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_SET_DETAILS"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupSetDetailsMgr) GetTableName() string {
	return "CDB_OB_BACKUP_SET_DETAILS"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupSetDetailsMgr) Reset() *_CdbObBackupSetDetailsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupSetDetailsMgr) Get() (result CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupSetDetailsMgr) Gets() (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupSetDetailsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupSetDetailsMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupSetDetailsMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupSetDetailsMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupSetDetailsMgr) WithCopyID(copyID string) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupSetDetailsMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithEncryptionMode ENCRYPTION_MODE获取
func (obj *_CdbObBackupSetDetailsMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["ENCRYPTION_MODE"] = encryptionMode })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupSetDetailsMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupSetDetailsMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithElapsedSecondes ELAPSED_SECONDES获取
func (obj *_CdbObBackupSetDetailsMgr) WithElapsedSecondes(elapsedSecondes float64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_SECONDES"] = elapsedSecondes })
}

// WithKeep KEEP获取
func (obj *_CdbObBackupSetDetailsMgr) WithKeep(keep string) Option {
	return optionFunc(func(o *options) { o.query["KEEP"] = keep })
}

// WithKeepUntil KEEP_UNTIL获取
func (obj *_CdbObBackupSetDetailsMgr) WithKeepUntil(keepUntil string) Option {
	return optionFunc(func(o *options) { o.query["KEEP_UNTIL"] = keepUntil })
}

// WithDeviceType DEVICE_TYPE获取
func (obj *_CdbObBackupSetDetailsMgr) WithDeviceType(deviceType string) Option {
	return optionFunc(func(o *options) { o.query["DEVICE_TYPE"] = deviceType })
}

// WithCompressed COMPRESSED获取
func (obj *_CdbObBackupSetDetailsMgr) WithCompressed(compressed string) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSED"] = compressed })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupSetDetailsMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithOutputRateBytes OUTPUT_RATE_BYTES获取
func (obj *_CdbObBackupSetDetailsMgr) WithOutputRateBytes(outputRateBytes float64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES"] = outputRateBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupSetDetailsMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetDetailsMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// WithOutputRateBytesDisplay OUTPUT_RATE_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetDetailsMgr) WithOutputRateBytesDisplay(outputRateBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES_DISPLAY"] = outputRateBytesDisplay })
}

// WithTimeTakenDisplay TIME_TAKEN_DISPLAY获取
func (obj *_CdbObBackupSetDetailsMgr) WithTimeTakenDisplay(timeTakenDisplay time.Time) Option {
	return optionFunc(func(o *options) { o.query["TIME_TAKEN_DISPLAY"] = timeTakenDisplay })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupSetDetailsMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupSetDetailsMgr) GetByOption(opts ...Option) (result CdbObBackupSetDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupSetDetailsMgr) GetByOptions(opts ...Option) (results []*CdbObBackupSetDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupSetDetailsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupSetDetails, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where(options.query)
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
func (obj *_CdbObBackupSetDetailsMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromCopyID(copyID string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromCopyID(copyIDs []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromBackupType(backupType string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过ENCRYPTION_MODE获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromEncryptionMode(encryptionMode string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`ENCRYPTION_MODE` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`ENCRYPTION_MODE` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromElapsedSecondes 通过ELAPSED_SECONDES获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromElapsedSecondes(elapsedSecondes float64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`ELAPSED_SECONDES` = ?", elapsedSecondes).Find(&results).Error

	return
}

// GetBatchFromElapsedSecondes 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromElapsedSecondes(elapsedSecondess []float64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`ELAPSED_SECONDES` IN (?)", elapsedSecondess).Find(&results).Error

	return
}

// GetFromKeep 通过KEEP获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromKeep(keep string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`KEEP` = ?", keep).Find(&results).Error

	return
}

// GetBatchFromKeep 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromKeep(keeps []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`KEEP` IN (?)", keeps).Find(&results).Error

	return
}

// GetFromKeepUntil 通过KEEP_UNTIL获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromKeepUntil(keepUntil string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`KEEP_UNTIL` = ?", keepUntil).Find(&results).Error

	return
}

// GetBatchFromKeepUntil 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromKeepUntil(keepUntils []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`KEEP_UNTIL` IN (?)", keepUntils).Find(&results).Error

	return
}

// GetFromDeviceType 通过DEVICE_TYPE获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromDeviceType(deviceType string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`DEVICE_TYPE` = ?", deviceType).Find(&results).Error

	return
}

// GetBatchFromDeviceType 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromDeviceType(deviceTypes []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`DEVICE_TYPE` IN (?)", deviceTypes).Find(&results).Error

	return
}

// GetFromCompressed 通过COMPRESSED获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromCompressed(compressed string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COMPRESSED` = ?", compressed).Find(&results).Error

	return
}

// GetBatchFromCompressed 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromCompressed(compresseds []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COMPRESSED` IN (?)", compresseds).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromOutputRateBytes 通过OUTPUT_RATE_BYTES获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromOutputRateBytes(outputRateBytes float64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_RATE_BYTES` = ?", outputRateBytes).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytes 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromOutputRateBytes(outputRateBytess []float64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_RATE_BYTES` IN (?)", outputRateBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputRateBytesDisplay 通过OUTPUT_RATE_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromOutputRateBytesDisplay(outputRateBytesDisplay string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` = ?", outputRateBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytesDisplay 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromOutputRateBytesDisplay(outputRateBytesDisplays []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` IN (?)", outputRateBytesDisplays).Find(&results).Error

	return
}

// GetFromTimeTakenDisplay 通过TIME_TAKEN_DISPLAY获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromTimeTakenDisplay(timeTakenDisplay time.Time) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`TIME_TAKEN_DISPLAY` = ?", timeTakenDisplay).Find(&results).Error

	return
}

// GetBatchFromTimeTakenDisplay 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromTimeTakenDisplay(timeTakenDisplays []time.Time) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`TIME_TAKEN_DISPLAY` IN (?)", timeTakenDisplays).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupSetDetailsMgr) GetFromStatus(status string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupSetDetailsMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupSetDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetDetails{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
