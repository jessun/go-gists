package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CdbObBackupBackupArchivelogSummaryMgr struct {
	*_BaseMgr
}

// CdbObBackupBackupArchivelogSummaryMgr open func
func CdbObBackupBackupArchivelogSummaryMgr(db *gorm.DB) *_CdbObBackupBackupArchivelogSummaryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackupArchivelogSummaryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackupArchivelogSummaryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUP_ARCHIVELOG_SUMMARY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUP_ARCHIVELOG_SUMMARY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) Reset() *_CdbObBackupBackupArchivelogSummaryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) Get() (result CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) Gets() (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithLogArchiveRound LOG_ARCHIVE_ROUND获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["LOG_ARCHIVE_ROUND"] = logArchiveRound })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithMinFirstTime MIN_FIRST_TIME获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithMinFirstTime(minFirstTime string) Option {
	return optionFunc(func(o *options) { o.query["MIN_FIRST_TIME"] = minFirstTime })
}

// WithMaxNextTime MAX_NEXT_TIME获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithMaxNextTime(maxNextTime string) Option {
	return optionFunc(func(o *options) { o.query["MAX_NEXT_TIME"] = maxNextTime })
}

// WithInputBytes INPUT_BYTES获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES"] = inputBytes })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithInputBytesDisplay INPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithInputBytesDisplay(inputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES_DISPLAY"] = inputBytesDisplay })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetByOption(opts ...Option) (result CdbObBackupBackupArchivelogSummary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackupArchivelogSummary, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where(options.query)
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
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过LOG_ARCHIVE_ROUND获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`LOG_ARCHIVE_ROUND` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`LOG_ARCHIVE_ROUND` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromStatus(status string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过MIN_FIRST_TIME获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromMinFirstTime(minFirstTime string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`MIN_FIRST_TIME` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromMinFirstTime(minFirstTimes []string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`MIN_FIRST_TIME` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过MAX_NEXT_TIME获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromMaxNextTime(maxNextTime string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`MAX_NEXT_TIME` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromMaxNextTime(maxNextTimes []string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`MAX_NEXT_TIME` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过INPUT_BYTES获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromInputBytes(inputBytes int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`INPUT_BYTES` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`INPUT_BYTES` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromInputBytesDisplay 通过INPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromInputBytesDisplay(inputBytesDisplay string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`INPUT_BYTES_DISPLAY` = ?", inputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromInputBytesDisplay 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromInputBytesDisplay(inputBytesDisplays []string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`INPUT_BYTES_DISPLAY` IN (?)", inputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupBackupArchivelogSummaryMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupArchivelogSummary{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
