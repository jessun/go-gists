package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupBackupLogArchiveStatusV2Mgr struct {
	*_BaseMgr
}

// AllBackupBackupLogArchiveStatusV2Mgr open func
func AllBackupBackupLogArchiveStatusV2Mgr(db *gorm.DB) *_AllBackupBackupLogArchiveStatusV2Mgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupBackupLogArchiveStatusV2Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupBackupLogArchiveStatusV2Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_backup_log_archive_status_v2"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetTableName() string {
	return "__all_backup_backup_log_archive_status_v2"
}

// Reset 重置gorm会话
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) Reset() *_AllBackupBackupLogArchiveStatusV2Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) Get() (result AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) Gets() (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithCopyID copy_id获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithLogArchiveRound log_archive_round获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_round"] = logArchiveRound })
}

// WithMinFirstTime min_first_time获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithMinFirstTime(minFirstTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["min_first_time"] = minFirstTime })
}

// WithMaxNextTime max_next_time获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithMaxNextTime(maxNextTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["max_next_time"] = maxNextTime })
}

// WithInputBytes input_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithDeletedInputBytes deleted_input_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithDeletedInputBytes(deletedInputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_input_bytes"] = deletedInputBytes })
}

// WithDeletedOutputBytes deleted_output_bytes获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithDeletedOutputBytes(deletedOutputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_output_bytes"] = deletedOutputBytes })
}

// WithPgCount pg_count获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithStatus status获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsMountFileCreated is_mount_file_created获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithIsMountFileCreated(isMountFileCreated int64) Option {
	return optionFunc(func(o *options) { o.query["is_mount_file_created"] = isMountFileCreated })
}

// WithCompatible compatible获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithStartPieceID start_piece_id获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithStartPieceID(startPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["start_piece_id"] = startPieceID })
}

// WithBackupPieceID backup_piece_id获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_piece_id"] = backupPieceID })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetByOption(opts ...Option) (result AllBackupBackupLogArchiveStatusV2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetByOptions(opts ...Option) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupBackupLogArchiveStatusV2, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where(options.query)
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
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromTenantID(tenantID int64) (result AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromCopyID(copyID int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromIncarnation(incarnation int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过log_archive_round获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`log_archive_round` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`log_archive_round` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过min_first_time获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromMinFirstTime(minFirstTime time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`min_first_time` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromMinFirstTime(minFirstTimes []time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`min_first_time` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过max_next_time获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromMaxNextTime(maxNextTime time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`max_next_time` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromMaxNextTime(maxNextTimes []time.Time) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`max_next_time` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromInputBytes(inputBytes int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromOutputBytes(outputBytes int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromDeletedInputBytes 通过deleted_input_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromDeletedInputBytes(deletedInputBytes int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`deleted_input_bytes` = ?", deletedInputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedInputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromDeletedInputBytes(deletedInputBytess []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`deleted_input_bytes` IN (?)", deletedInputBytess).Find(&results).Error

	return
}

// GetFromDeletedOutputBytes 通过deleted_output_bytes获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromDeletedOutputBytes(deletedOutputBytes int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`deleted_output_bytes` = ?", deletedOutputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedOutputBytes 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromDeletedOutputBytes(deletedOutputBytess []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`deleted_output_bytes` IN (?)", deletedOutputBytess).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromPgCount(pgCount int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromStatus(status string) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromStatus(statuss []string) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsMountFileCreated 通过is_mount_file_created获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromIsMountFileCreated(isMountFileCreated int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`is_mount_file_created` = ?", isMountFileCreated).Find(&results).Error

	return
}

// GetBatchFromIsMountFileCreated 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromIsMountFileCreated(isMountFileCreateds []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`is_mount_file_created` IN (?)", isMountFileCreateds).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromCompatible(compatible int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromCompatible(compatibles []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromStartPieceID 通过start_piece_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromStartPieceID(startPieceID int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`start_piece_id` = ?", startPieceID).Find(&results).Error

	return
}

// GetBatchFromStartPieceID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromStartPieceID(startPieceIDs []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`start_piece_id` IN (?)", startPieceIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过backup_piece_id获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromBackupPieceID(backupPieceID int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`backup_piece_id` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`backup_piece_id` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetFromBackupDest(backupDest string) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupBackupLogArchiveStatusV2Mgr) FetchByPrimaryKey(tenantID int64) (result AllBackupBackupLogArchiveStatusV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupLogArchiveStatusV2{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
