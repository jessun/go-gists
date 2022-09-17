package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupSetObsoleteMgr struct {
	*_BaseMgr
}

// CdbObBackupSetObsoleteMgr open func
func CdbObBackupSetObsoleteMgr(db *gorm.DB) *_CdbObBackupSetObsoleteMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupSetObsoleteMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupSetObsoleteMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_SET_OBSOLETE"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupSetObsoleteMgr) GetTableName() string {
	return "CDB_OB_BACKUP_SET_OBSOLETE"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupSetObsoleteMgr) Reset() *_CdbObBackupSetObsoleteMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupSetObsoleteMgr) Get() (result CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupSetObsoleteMgr) Gets() (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupSetObsoleteMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupSetObsoleteMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupSetObsoleteMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupSetObsoleteMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupSetObsoleteMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithEncryptionMode ENCRYPTION_MODE获取
func (obj *_CdbObBackupSetObsoleteMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["ENCRYPTION_MODE"] = encryptionMode })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupSetObsoleteMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithCompletionTime COMPLETION_TIME获取
func (obj *_CdbObBackupSetObsoleteMgr) WithCompletionTime(completionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["COMPLETION_TIME"] = completionTime })
}

// WithElapsedSecondes ELAPSED_SECONDES获取
func (obj *_CdbObBackupSetObsoleteMgr) WithElapsedSecondes(elapsedSecondes float64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_SECONDES"] = elapsedSecondes })
}

// WithKeep KEEP获取
func (obj *_CdbObBackupSetObsoleteMgr) WithKeep(keep string) Option {
	return optionFunc(func(o *options) { o.query["KEEP"] = keep })
}

// WithKeepUntil KEEP_UNTIL获取
func (obj *_CdbObBackupSetObsoleteMgr) WithKeepUntil(keepUntil string) Option {
	return optionFunc(func(o *options) { o.query["KEEP_UNTIL"] = keepUntil })
}

// WithDeviceType DEVICE_TYPE获取
func (obj *_CdbObBackupSetObsoleteMgr) WithDeviceType(deviceType string) Option {
	return optionFunc(func(o *options) { o.query["DEVICE_TYPE"] = deviceType })
}

// WithCompressed COMPRESSED获取
func (obj *_CdbObBackupSetObsoleteMgr) WithCompressed(compressed string) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSED"] = compressed })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupSetObsoleteMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithOutputRateBytes OUTPUT_RATE_BYTES获取
func (obj *_CdbObBackupSetObsoleteMgr) WithOutputRateBytes(outputRateBytes float64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES"] = outputRateBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupSetObsoleteMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetObsoleteMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// WithOutputRateBytesDisplay OUTPUT_RATE_BYTES_DISPLAY获取
func (obj *_CdbObBackupSetObsoleteMgr) WithOutputRateBytesDisplay(outputRateBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_RATE_BYTES_DISPLAY"] = outputRateBytesDisplay })
}

// WithTimeTakenDisplay TIME_TAKEN_DISPLAY获取
func (obj *_CdbObBackupSetObsoleteMgr) WithTimeTakenDisplay(timeTakenDisplay time.Time) Option {
	return optionFunc(func(o *options) { o.query["TIME_TAKEN_DISPLAY"] = timeTakenDisplay })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupSetObsoleteMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupSetObsoleteMgr) GetByOption(opts ...Option) (result CdbObBackupSetObsolete, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupSetObsoleteMgr) GetByOptions(opts ...Option) (results []*CdbObBackupSetObsolete, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupSetObsoleteMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupSetObsolete, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where(options.query)
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
func (obj *_CdbObBackupSetObsoleteMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromBackupType(backupType string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过ENCRYPTION_MODE获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromEncryptionMode(encryptionMode string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`ENCRYPTION_MODE` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`ENCRYPTION_MODE` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromCompletionTime 通过COMPLETION_TIME获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromCompletionTime(completionTime time.Time) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`COMPLETION_TIME` = ?", completionTime).Find(&results).Error

	return
}

// GetBatchFromCompletionTime 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromCompletionTime(completionTimes []time.Time) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`COMPLETION_TIME` IN (?)", completionTimes).Find(&results).Error

	return
}

// GetFromElapsedSecondes 通过ELAPSED_SECONDES获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromElapsedSecondes(elapsedSecondes float64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`ELAPSED_SECONDES` = ?", elapsedSecondes).Find(&results).Error

	return
}

// GetBatchFromElapsedSecondes 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromElapsedSecondes(elapsedSecondess []float64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`ELAPSED_SECONDES` IN (?)", elapsedSecondess).Find(&results).Error

	return
}

// GetFromKeep 通过KEEP获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromKeep(keep string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`KEEP` = ?", keep).Find(&results).Error

	return
}

// GetBatchFromKeep 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromKeep(keeps []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`KEEP` IN (?)", keeps).Find(&results).Error

	return
}

// GetFromKeepUntil 通过KEEP_UNTIL获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromKeepUntil(keepUntil string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`KEEP_UNTIL` = ?", keepUntil).Find(&results).Error

	return
}

// GetBatchFromKeepUntil 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromKeepUntil(keepUntils []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`KEEP_UNTIL` IN (?)", keepUntils).Find(&results).Error

	return
}

// GetFromDeviceType 通过DEVICE_TYPE获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromDeviceType(deviceType string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`DEVICE_TYPE` = ?", deviceType).Find(&results).Error

	return
}

// GetBatchFromDeviceType 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromDeviceType(deviceTypes []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`DEVICE_TYPE` IN (?)", deviceTypes).Find(&results).Error

	return
}

// GetFromCompressed 通过COMPRESSED获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromCompressed(compressed string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`COMPRESSED` = ?", compressed).Find(&results).Error

	return
}

// GetBatchFromCompressed 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromCompressed(compresseds []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`COMPRESSED` IN (?)", compresseds).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromOutputRateBytes 通过OUTPUT_RATE_BYTES获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromOutputRateBytes(outputRateBytes float64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_RATE_BYTES` = ?", outputRateBytes).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytes 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromOutputRateBytes(outputRateBytess []float64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_RATE_BYTES` IN (?)", outputRateBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputRateBytesDisplay 通过OUTPUT_RATE_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromOutputRateBytesDisplay(outputRateBytesDisplay string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` = ?", outputRateBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputRateBytesDisplay 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromOutputRateBytesDisplay(outputRateBytesDisplays []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`OUTPUT_RATE_BYTES_DISPLAY` IN (?)", outputRateBytesDisplays).Find(&results).Error

	return
}

// GetFromTimeTakenDisplay 通过TIME_TAKEN_DISPLAY获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromTimeTakenDisplay(timeTakenDisplay time.Time) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`TIME_TAKEN_DISPLAY` = ?", timeTakenDisplay).Find(&results).Error

	return
}

// GetBatchFromTimeTakenDisplay 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromTimeTakenDisplay(timeTakenDisplays []time.Time) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`TIME_TAKEN_DISPLAY` IN (?)", timeTakenDisplays).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupSetObsoleteMgr) GetFromStatus(status string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupSetObsoleteMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupSetObsolete, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupSetObsolete{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
