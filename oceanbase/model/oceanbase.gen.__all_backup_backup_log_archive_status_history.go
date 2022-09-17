package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupBackupLogArchiveStatusHistoryMgr struct {
	*_BaseMgr
}

// AllBackupBackupLogArchiveStatusHistoryMgr open func
func AllBackupBackupLogArchiveStatusHistoryMgr(db *gorm.DB) *_AllBackupBackupLogArchiveStatusHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupBackupLogArchiveStatusHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupBackupLogArchiveStatusHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_backup_log_archive_status_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetTableName() string {
	return "__all_backup_backup_log_archive_status_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) Reset() *_AllBackupBackupLogArchiveStatusHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) Get() (result AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) Gets() (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithLogArchiveRound log_archive_round获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_round"] = logArchiveRound })
}

// WithCopyID copy_id获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithMinFirstTime min_first_time获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithMinFirstTime(minFirstTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["min_first_time"] = minFirstTime })
}

// WithMaxNextTime max_next_time获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithMaxNextTime(maxNextTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["max_next_time"] = maxNextTime })
}

// WithInputBytes input_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithDeletedInputBytes deleted_input_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithDeletedInputBytes(deletedInputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_input_bytes"] = deletedInputBytes })
}

// WithDeletedOutputBytes deleted_output_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithDeletedOutputBytes(deletedOutputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_output_bytes"] = deletedOutputBytes })
}

// WithPgCount pg_count获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithIsMarkDeleted is_mark_deleted获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithIsMarkDeleted(isMarkDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_mark_deleted"] = isMarkDeleted })
}

// WithCompatible compatible获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithStartPieceID start_piece_id获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithStartPieceID(startPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["start_piece_id"] = startPieceID })
}

// WithBackupPieceID backup_piece_id获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_piece_id"] = backupPieceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetByOption(opts ...Option) (result AllBackupBackupLogArchiveStatusHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupBackupLogArchiveStatusHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过log_archive_round获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`log_archive_round` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`log_archive_round` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromCopyID(copyID int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过min_first_time获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromMinFirstTime(minFirstTime time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`min_first_time` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromMinFirstTime(minFirstTimes []time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`min_first_time` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过max_next_time获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromMaxNextTime(maxNextTime time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`max_next_time` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromMaxNextTime(maxNextTimes []time.Time) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`max_next_time` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromInputBytes(inputBytes int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromOutputBytes(outputBytes int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromDeletedInputBytes 通过deleted_input_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromDeletedInputBytes(deletedInputBytes int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`deleted_input_bytes` = ?", deletedInputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedInputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromDeletedInputBytes(deletedInputBytess []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`deleted_input_bytes` IN (?)", deletedInputBytess).Find(&results).Error

	return
}

// GetFromDeletedOutputBytes 通过deleted_output_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromDeletedOutputBytes(deletedOutputBytes int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`deleted_output_bytes` = ?", deletedOutputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedOutputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromDeletedOutputBytes(deletedOutputBytess []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`deleted_output_bytes` IN (?)", deletedOutputBytess).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromPgCount(pgCount int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromBackupDest(backupDest string) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromIsMarkDeleted 通过is_mark_deleted获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromIsMarkDeleted(isMarkDeleted int8) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`is_mark_deleted` = ?", isMarkDeleted).Find(&results).Error

	return
}

// GetBatchFromIsMarkDeleted 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromIsMarkDeleted(isMarkDeleteds []int8) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`is_mark_deleted` IN (?)", isMarkDeleteds).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromCompatible(compatible int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromStartPieceID 通过start_piece_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromStartPieceID(startPieceID int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`start_piece_id` = ?", startPieceID).Find(&results).Error

	return
}

// GetBatchFromStartPieceID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromStartPieceID(startPieceIDs []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`start_piece_id` IN (?)", startPieceIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过backup_piece_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetFromBackupPieceID(backupPieceID int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`backup_piece_id` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`backup_piece_id` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupBackupLogArchiveStatusHistoryMgr) FetchByPrimaryKey(tenantID int64, incarnation int64, logArchiveRound int64, copyID int64) (result AllBackupBackupLogArchiveStatusHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusHistory{}).Where("`tenant_id` = ? AND `incarnation` = ? AND `log_archive_round` = ? AND `copy_id` = ?", tenantID, incarnation, logArchiveRound, copyID).First(&result).Error

	return
}
