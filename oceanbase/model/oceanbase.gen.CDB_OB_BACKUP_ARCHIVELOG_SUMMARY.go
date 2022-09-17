package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CdbObBackupArchivelogSummaryMgr struct {
	*_BaseMgr
}

// CdbObBackupArchivelogSummaryMgr open func
func CdbObBackupArchivelogSummaryMgr(db *gorm.DB) *_CdbObBackupArchivelogSummaryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupArchivelogSummaryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupArchivelogSummaryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_ARCHIVELOG_SUMMARY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupArchivelogSummaryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_ARCHIVELOG_SUMMARY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupArchivelogSummaryMgr) Reset() *_CdbObBackupArchivelogSummaryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupArchivelogSummaryMgr) Get() (result CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupArchivelogSummaryMgr) Gets() (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupArchivelogSummaryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithLogArchiveRound LOG_ARCHIVE_ROUND获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["LOG_ARCHIVE_ROUND"] = logArchiveRound })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithStartPieceID START_PIECE_ID获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithStartPieceID(startPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["START_PIECE_ID"] = startPieceID })
}

// WithBackupPieceID BACKUP_PIECE_ID获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_ID"] = backupPieceID })
}

// WithMinFirstTime MIN_FIRST_TIME获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithMinFirstTime(minFirstTime string) Option {
	return optionFunc(func(o *options) { o.query["MIN_FIRST_TIME"] = minFirstTime })
}

// WithMaxNextTime MAX_NEXT_TIME获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithMaxNextTime(maxNextTime string) Option {
	return optionFunc(func(o *options) { o.query["MAX_NEXT_TIME"] = maxNextTime })
}

// WithInputBytes INPUT_BYTES获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES"] = inputBytes })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithInputBytesDisplay INPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithInputBytesDisplay(inputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES_DISPLAY"] = inputBytesDisplay })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupArchivelogSummaryMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupArchivelogSummaryMgr) GetByOption(opts ...Option) (result CdbObBackupArchivelogSummary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupArchivelogSummaryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupArchivelogSummary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupArchivelogSummaryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupArchivelogSummary, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where(options.query)
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
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过LOG_ARCHIVE_ROUND获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`LOG_ARCHIVE_ROUND` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`LOG_ARCHIVE_ROUND` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromStatus(status string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromStartPieceID 通过START_PIECE_ID获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromStartPieceID(startPieceID int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`START_PIECE_ID` = ?", startPieceID).Find(&results).Error

	return
}

// GetBatchFromStartPieceID 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromStartPieceID(startPieceIDs []int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`START_PIECE_ID` IN (?)", startPieceIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过BACKUP_PIECE_ID获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromBackupPieceID(backupPieceID int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`BACKUP_PIECE_ID` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`BACKUP_PIECE_ID` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过MIN_FIRST_TIME获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromMinFirstTime(minFirstTime string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`MIN_FIRST_TIME` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromMinFirstTime(minFirstTimes []string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`MIN_FIRST_TIME` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过MAX_NEXT_TIME获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromMaxNextTime(maxNextTime string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`MAX_NEXT_TIME` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromMaxNextTime(maxNextTimes []string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`MAX_NEXT_TIME` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过INPUT_BYTES获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromInputBytes(inputBytes int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`INPUT_BYTES` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`INPUT_BYTES` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromInputBytesDisplay 通过INPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromInputBytesDisplay(inputBytesDisplay string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`INPUT_BYTES_DISPLAY` = ?", inputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromInputBytesDisplay 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromInputBytesDisplay(inputBytesDisplays []string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`INPUT_BYTES_DISPLAY` IN (?)", inputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupArchivelogSummaryMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupArchivelogSummaryMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupArchivelogSummary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogSummary{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
