package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantBackupBackupLogArchiveStatusMgr struct {
	*_BaseMgr
}

// AllTenantBackupBackupLogArchiveStatusMgr open func
func AllTenantBackupBackupLogArchiveStatusMgr(db *gorm.DB) *_AllTenantBackupBackupLogArchiveStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantBackupBackupLogArchiveStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantBackupBackupLogArchiveStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_backup_backup_log_archive_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetTableName() string {
	return "__all_tenant_backup_backup_log_archive_status"
}

// Reset 重置gorm会话
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) Reset() *_AllTenantBackupBackupLogArchiveStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) Get() (result AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) Gets() (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithLogArchiveRound log_archive_round获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_round"] = logArchiveRound })
}

// WithCopyID copy_id获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithMinFirstTime min_first_time获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithMinFirstTime(minFirstTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["min_first_time"] = minFirstTime })
}

// WithMaxNextTime max_next_time获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithMaxNextTime(maxNextTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["max_next_time"] = maxNextTime })
}

// WithInputBytes input_bytes获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithDeletedInputBytes deleted_input_bytes获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithDeletedInputBytes(deletedInputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_input_bytes"] = deletedInputBytes })
}

// WithDeletedOutputBytes deleted_output_bytes获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithDeletedOutputBytes(deletedOutputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_output_bytes"] = deletedOutputBytes })
}

// WithPgCount pg_count获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithStatus status获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetByOption(opts ...Option) (result AllTenantBackupBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetByOptions(opts ...Option) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantBackupBackupLogArchiveStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where(options.query)
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
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromTenantID(tenantID int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromIncarnation(incarnation int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过log_archive_round获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`log_archive_round` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`log_archive_round` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromCopyID(copyID int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过min_first_time获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromMinFirstTime(minFirstTime time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`min_first_time` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromMinFirstTime(minFirstTimes []time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`min_first_time` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过max_next_time获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromMaxNextTime(maxNextTime time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`max_next_time` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromMaxNextTime(maxNextTimes []time.Time) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`max_next_time` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromInputBytes(inputBytes int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromOutputBytes(outputBytes int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromDeletedInputBytes 通过deleted_input_bytes获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromDeletedInputBytes(deletedInputBytes int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`deleted_input_bytes` = ?", deletedInputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedInputBytes 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromDeletedInputBytes(deletedInputBytess []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`deleted_input_bytes` IN (?)", deletedInputBytess).Find(&results).Error

	return
}

// GetFromDeletedOutputBytes 通过deleted_output_bytes获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromDeletedOutputBytes(deletedOutputBytes int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`deleted_output_bytes` = ?", deletedOutputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedOutputBytes 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromDeletedOutputBytes(deletedOutputBytess []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`deleted_output_bytes` IN (?)", deletedOutputBytess).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromPgCount(pgCount int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetFromStatus(status string) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantBackupBackupLogArchiveStatusMgr) FetchByPrimaryKey(tenantID int64, incarnation int64, logArchiveRound int64, copyID int64) (result AllTenantBackupBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupBackupLogArchiveStatus{}).Where("`tenant_id` = ? AND `incarnation` = ? AND `log_archive_round` = ? AND `copy_id` = ?", tenantID, incarnation, logArchiveRound, copyID).First(&result).Error

	return
}
