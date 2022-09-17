package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CdbObBackupArchivelogMgr struct {
	*_BaseMgr
}

// CdbObBackupArchivelogMgr open func
func CdbObBackupArchivelogMgr(db *gorm.DB) *_CdbObBackupArchivelogMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupArchivelogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupArchivelogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_ARCHIVELOG"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupArchivelogMgr) GetTableName() string {
	return "CDB_OB_BACKUP_ARCHIVELOG"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupArchivelogMgr) Reset() *_CdbObBackupArchivelogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupArchivelogMgr) Get() (result CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupArchivelogMgr) Gets() (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupArchivelogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupArchivelogMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithLogArchiveRound LOG_ARCHIVE_ROUND获取
func (obj *_CdbObBackupArchivelogMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["LOG_ARCHIVE_ROUND"] = logArchiveRound })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupArchivelogMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupArchivelogMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithStartPieceID START_PIECE_ID获取
func (obj *_CdbObBackupArchivelogMgr) WithStartPieceID(startPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["START_PIECE_ID"] = startPieceID })
}

// WithBackupPieceID BACKUP_PIECE_ID获取
func (obj *_CdbObBackupArchivelogMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_ID"] = backupPieceID })
}

// WithMinFirstTime MIN_FIRST_TIME获取
func (obj *_CdbObBackupArchivelogMgr) WithMinFirstTime(minFirstTime string) Option {
	return optionFunc(func(o *options) { o.query["MIN_FIRST_TIME"] = minFirstTime })
}

// WithMaxNextTime MAX_NEXT_TIME获取
func (obj *_CdbObBackupArchivelogMgr) WithMaxNextTime(maxNextTime string) Option {
	return optionFunc(func(o *options) { o.query["MAX_NEXT_TIME"] = maxNextTime })
}

// WithInputBytes INPUT_BYTES获取
func (obj *_CdbObBackupArchivelogMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES"] = inputBytes })
}

// WithOutputBytes OUTPUT_BYTES获取
func (obj *_CdbObBackupArchivelogMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES"] = outputBytes })
}

// WithCompressionRatio COMPRESSION_RATIO获取
func (obj *_CdbObBackupArchivelogMgr) WithCompressionRatio(compressionRatio float64) Option {
	return optionFunc(func(o *options) { o.query["COMPRESSION_RATIO"] = compressionRatio })
}

// WithInputBytesDisplay INPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupArchivelogMgr) WithInputBytesDisplay(inputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["INPUT_BYTES_DISPLAY"] = inputBytesDisplay })
}

// WithOutputBytesDisplay OUTPUT_BYTES_DISPLAY获取
func (obj *_CdbObBackupArchivelogMgr) WithOutputBytesDisplay(outputBytesDisplay string) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_BYTES_DISPLAY"] = outputBytesDisplay })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupArchivelogMgr) GetByOption(opts ...Option) (result CdbObBackupArchivelog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupArchivelogMgr) GetByOptions(opts ...Option) (results []*CdbObBackupArchivelog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupArchivelogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupArchivelog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where(options.query)
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
func (obj *_CdbObBackupArchivelogMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过LOG_ARCHIVE_ROUND获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`LOG_ARCHIVE_ROUND` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`LOG_ARCHIVE_ROUND` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromStatus(status string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromStartPieceID 通过START_PIECE_ID获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromStartPieceID(startPieceID int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`START_PIECE_ID` = ?", startPieceID).Find(&results).Error

	return
}

// GetBatchFromStartPieceID 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromStartPieceID(startPieceIDs []int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`START_PIECE_ID` IN (?)", startPieceIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过BACKUP_PIECE_ID获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromBackupPieceID(backupPieceID int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`BACKUP_PIECE_ID` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`BACKUP_PIECE_ID` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过MIN_FIRST_TIME获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromMinFirstTime(minFirstTime string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`MIN_FIRST_TIME` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromMinFirstTime(minFirstTimes []string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`MIN_FIRST_TIME` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过MAX_NEXT_TIME获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromMaxNextTime(maxNextTime string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`MAX_NEXT_TIME` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromMaxNextTime(maxNextTimes []string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`MAX_NEXT_TIME` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过INPUT_BYTES获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromInputBytes(inputBytes int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`INPUT_BYTES` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`INPUT_BYTES` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过OUTPUT_BYTES获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromOutputBytes(outputBytes int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`OUTPUT_BYTES` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`OUTPUT_BYTES` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromCompressionRatio 通过COMPRESSION_RATIO获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromCompressionRatio(compressionRatio float64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`COMPRESSION_RATIO` = ?", compressionRatio).Find(&results).Error

	return
}

// GetBatchFromCompressionRatio 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromCompressionRatio(compressionRatios []float64) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`COMPRESSION_RATIO` IN (?)", compressionRatios).Find(&results).Error

	return
}

// GetFromInputBytesDisplay 通过INPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromInputBytesDisplay(inputBytesDisplay string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`INPUT_BYTES_DISPLAY` = ?", inputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromInputBytesDisplay 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromInputBytesDisplay(inputBytesDisplays []string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`INPUT_BYTES_DISPLAY` IN (?)", inputBytesDisplays).Find(&results).Error

	return
}

// GetFromOutputBytesDisplay 通过OUTPUT_BYTES_DISPLAY获取内容
func (obj *_CdbObBackupArchivelogMgr) GetFromOutputBytesDisplay(outputBytesDisplay string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`OUTPUT_BYTES_DISPLAY` = ?", outputBytesDisplay).Find(&results).Error

	return
}

// GetBatchFromOutputBytesDisplay 批量查找
func (obj *_CdbObBackupArchivelogMgr) GetBatchFromOutputBytesDisplay(outputBytesDisplays []string) (results []*CdbObBackupArchivelog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelog{}).Where("`OUTPUT_BYTES_DISPLAY` IN (?)", outputBytesDisplays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
