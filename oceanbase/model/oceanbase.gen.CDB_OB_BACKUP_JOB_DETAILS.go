package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupJobDetailsMgr struct {
	*_BaseMgr
}

// CdbObBackupJobDetailsMgr open func
func CdbObBackupJobDetailsMgr(db *gorm.DB) *_CdbObBackupJobDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupJobDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupJobDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_JOB_DETAILS"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupJobDetailsMgr) GetTableName() string {
	return "CDB_OB_BACKUP_JOB_DETAILS"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupJobDetailsMgr) Reset() *_CdbObBackupJobDetailsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupJobDetailsMgr) Get() (result CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupJobDetailsMgr) Gets() (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupJobDetailsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupJobDetailsMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupJobDetailsMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithBsKey BS_KEY获取
func (obj *_CdbObBackupJobDetailsMgr) WithBsKey(bsKey int64) Option {
	return optionFunc(func(o *options) { o.query["BS_KEY"] = bsKey })
}

// WithBackupType BACKUP_TYPE获取
func (obj *_CdbObBackupJobDetailsMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_TYPE"] = backupType })
}

// WithEncryptionMode ENCRYPTION_MODE获取
func (obj *_CdbObBackupJobDetailsMgr) WithEncryptionMode(encryptionMode string) Option {
	return optionFunc(func(o *options) { o.query["ENCRYPTION_MODE"] = encryptionMode })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupJobDetailsMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_CdbObBackupJobDetailsMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithInputBytes INPUT_BYTES获取
func (obj *_CdbObBackupJobDetailsMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES"] = inputBytes })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupJobDetailsMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithOutputDeviceType OUTPUT_DEVICE_TYPE获取
func (obj *_CdbObBackupJobDetailsMgr) WithOutputDeviceType(outputDeviceType string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_DEVICE_TYPE"] = outputDeviceType })
}

// WithElapsedSecondes ELAPSED_SECONDES获取
func (obj *_CdbObBackupJobDetailsMgr) WithElapsedSecondes(elapsedSecondes float64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_SECONDES"] = elapsedSecondes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupJobDetailsMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithInputBytesPerSec INPUT_BYTES_PER_SEC获取
func (obj *_CdbObBackupJobDetailsMgr) WithInputBytesPerSec(inputBytesPerSec float64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES_PER_SEC"] = inputBytesPerSec })
}

// WithOutputBytesPerSec OUTPUT_BYTES_PER_SEC获取
func (obj *_CdbObBackupJobDetailsMgr) WithOutputBytesPerSec(outputBytesPerSec float64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_PER_SEC"] = outputBytesPerSec })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupJobDetailsMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithInputBytesDisplay INPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupJobDetailsMgr) WithInputBytesDisplay(inputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES_DISPLAY"] = inputBytesDisplay })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupJobDetailsMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// WithInputBytesPerSecDisplay INPUT_BYTES_PER_SEC_DISPLAY获取
func (obj *_CdbObBackupJobDetailsMgr) WithInputBytesPerSecDisplay(inputBytesPerSecDisplay string) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES_PER_SEC_DISPLAY"] = inputBytesPerSecDisplay })
}

// WithOutputBytesPerSecDisplay OUTPUT_BYTES_PER_SEC_DISPLAY获取
func (obj *_CdbObBackupJobDetailsMgr) WithOutputBytesPerSecDisplay(outputBytesPerSecDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_PER_SEC_DISPLAY"] = outputBytesPerSecDisplay })
}

// WithTimeTakenDisplay TIME_TAKEN_DISPLAY获取
func (obj *_CdbObBackupJobDetailsMgr) WithTimeTakenDisplay(timeTakenDisplay time.Time) Option {
	return optionFunc(func(o *options) { o.query["TIME_TAKEN_DISPLAY"] = timeTakenDisplay })
}

// WithStartReplayLogTs START_REPLAY_LOG_TS获取
func (obj *_CdbObBackupJobDetailsMgr) WithStartReplayLogTs(startReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["START_REPLAY_LOG_TS"] = startReplayLogTs })
}

// WithDate DATE获取
func (obj *_CdbObBackupJobDetailsMgr) WithDate(date int64) Option {
	return optionFunc(func(o *options) { o.query["DATE"] = date })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupJobDetailsMgr) GetByOption(opts ...Option) (result CdbObBackupJobDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupJobDetailsMgr) GetByOptions(opts ...Option) (results []*CdbObBackupJobDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupJobDetailsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupJobDetails, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where(options.query)
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
func (obj *_CdbObBackupJobDetailsMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromBsKey 通过BS_KEY获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromBsKey(bsKey int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`BS_KEY` = ?", bsKey).Find(&results).Error

	return
}

// GetBatchFromBsKey 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromBsKey(bsKeys []int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`BS_KEY` IN (?)", bsKeys).Find(&results).Error

	return
}

// GetFromBackupType 通过BACKUP_TYPE获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromBackupType(backupType string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`BACKUP_TYPE` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromBackupType(backupTypes []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`BACKUP_TYPE` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromEncryptionMode 通过ENCRYPTION_MODE获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromEncryptionMode(encryptionMode string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`ENCRYPTION_MODE` = ?", encryptionMode).Find(&results).Error

	return
}

// GetBatchFromEncryptionMode 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromEncryptionMode(encryptionModes []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`ENCRYPTION_MODE` IN (?)", encryptionModes).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromEndTime(endTime time.Time) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过INPUT_BYTES获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromInputBytes(inputBytes int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromOutputDeviceType 通过OUTPUT_DEVICE_TYPE获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromOutputDeviceType(outputDeviceType string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_DEVICE_TYPE` = ?", outputDeviceType).Find(&results).Error

	return
}

// GetBatchFromOutputDeviceType 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromOutputDeviceType(outputDeviceTypes []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_DEVICE_TYPE` IN (?)", outputDeviceTypes).Find(&results).Error

	return
}

// GetFromElapsedSecondes 通过ELAPSED_SECONDES获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromElapsedSecondes(elapsedSecondes float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`ELAPSED_SECONDES` = ?", elapsedSecondes).Find(&results).Error

	return
}

// GetBatchFromElapsedSecondes 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromElapsedSecondes(elapsedSecondess []float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`ELAPSED_SECONDES` IN (?)", elapsedSecondess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromInputBytesPerSec 通过INPUT_BYTES_PER_SEC获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromInputBytesPerSec(inputBytesPerSec float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES_PER_SEC` = ?", inputBytesPerSec).Find(&results).Error

	return
}

// GetBatchFromInputBytesPerSec 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromInputBytesPerSec(inputBytesPerSecs []float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES_PER_SEC` IN (?)", inputBytesPerSecs).Find(&results).Error

	return
}

// GetFromOutputBytesPerSec 通过OUTPUT_BYTES_PER_SEC获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromOutputBytesPerSec(outputBytesPerSec float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES_PER_SEC` = ?", outputBytesPerSec).Find(&results).Error

	return
}

// GetBatchFromOutputBytesPerSec 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromOutputBytesPerSec(outputBytesPerSecs []float64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES_PER_SEC` IN (?)", outputBytesPerSecs).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromStatus(status string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromInputBytesDisplay 通过INPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromInputBytesDisplay(inputBytesDisplay string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES_DISPLAY` = ?", inputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromInputBytesDisplay 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromInputBytesDisplay(inputBytesDisplays []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES_DISPLAY` IN (?)", inputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

// GetFromInputBytesPerSecDisplay 通过INPUT_BYTES_PER_SEC_DISPLAY获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromInputBytesPerSecDisplay(inputBytesPerSecDisplay string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES_PER_SEC_DISPLAY` = ?", inputBytesPerSecDisplay).Find(&results).Error

	return
}

// GetBatchFromInputBytesPerSecDisplay 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromInputBytesPerSecDisplay(inputBytesPerSecDisplays []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`INPUT_BYTES_PER_SEC_DISPLAY` IN (?)", inputBytesPerSecDisplays).Find(&results).Error

	return
}

// GetFromOutputBytesPerSecDisplay 通过OUTPUT_BYTES_PER_SEC_DISPLAY获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromOutputBytesPerSecDisplay(outputBytesPerSecDisplay string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES_PER_SEC_DISPLAY` = ?", outputBytesPerSecDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesPerSecDisplay 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromOutputBytesPerSecDisplay(outputBytesPerSecDisplays []string) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`OUTPUT_BYTES_PER_SEC_DISPLAY` IN (?)", outputBytesPerSecDisplays).Find(&results).Error

	return
}

// GetFromTimeTakenDisplay 通过TIME_TAKEN_DISPLAY获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromTimeTakenDisplay(timeTakenDisplay time.Time) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`TIME_TAKEN_DISPLAY` = ?", timeTakenDisplay).Find(&results).Error

	return
}

// GetBatchFromTimeTakenDisplay 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromTimeTakenDisplay(timeTakenDisplays []time.Time) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`TIME_TAKEN_DISPLAY` IN (?)", timeTakenDisplays).Find(&results).Error

	return
}

// GetFromStartReplayLogTs 通过START_REPLAY_LOG_TS获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromStartReplayLogTs(startReplayLogTs int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`START_REPLAY_LOG_TS` = ?", startReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromStartReplayLogTs 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromStartReplayLogTs(startReplayLogTss []int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`START_REPLAY_LOG_TS` IN (?)", startReplayLogTss).Find(&results).Error

	return
}

// GetFromDate 通过DATE获取内容
func (obj *_CdbObBackupJobDetailsMgr) GetFromDate(date int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`DATE` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_CdbObBackupJobDetailsMgr) GetBatchFromDate(dates []int64) (results []*CdbObBackupJobDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupJobDetails{}).Where("`DATE` IN (?)", dates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
