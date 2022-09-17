package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantBackupLogArchiveStatusMgr struct {
	*_BaseMgr
}

// AllTenantBackupLogArchiveStatusMgr open func
func AllTenantBackupLogArchiveStatusMgr(db *gorm.DB) *_AllTenantBackupLogArchiveStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantBackupLogArchiveStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantBackupLogArchiveStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_backup_log_archive_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetTableName() string {
	return "__all_tenant_backup_log_archive_status"
}

// Reset 重置gorm会话
func (obj *_AllTenantBackupLogArchiveStatusMgr) Reset() *_AllTenantBackupLogArchiveStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) Get() (result AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantBackupLogArchiveStatusMgr) Gets() (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantBackupLogArchiveStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithLogArchiveRound log_archive_round获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_round"] = logArchiveRound })
}

// WithMinFirstTime min_first_time获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithMinFirstTime(minFirstTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["min_first_time"] = minFirstTime })
}

// WithMaxNextTime max_next_time获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithMaxNextTime(maxNextTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["max_next_time"] = maxNextTime })
}

// WithInputBytes input_bytes获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithInputBytes(inputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["input_bytes"] = inputBytes })
}

// WithOutputBytes output_bytes获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithOutputBytes(outputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["output_bytes"] = outputBytes })
}

// WithDeletedInputBytes deleted_input_bytes获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithDeletedInputBytes(deletedInputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_input_bytes"] = deletedInputBytes })
}

// WithDeletedOutputBytes deleted_output_bytes获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithDeletedOutputBytes(deletedOutputBytes int64) Option {
	return optionFunc(func(o *options) { o.query["deleted_output_bytes"] = deletedOutputBytes })
}

// WithPgCount pg_count获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithStatus status获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsMountFileCreated is_mount_file_created获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithIsMountFileCreated(isMountFileCreated int64) Option {
	return optionFunc(func(o *options) { o.query["is_mount_file_created"] = isMountFileCreated })
}

// WithCompatible compatible获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetByOption(opts ...Option) (result AllTenantBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetByOptions(opts ...Option) (results []*AllTenantBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantBackupLogArchiveStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantBackupLogArchiveStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where(options.query)
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
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromTenantID(tenantID int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromIncarnation(incarnation int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过log_archive_round获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`log_archive_round` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`log_archive_round` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过min_first_time获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromMinFirstTime(minFirstTime time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`min_first_time` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromMinFirstTime(minFirstTimes []time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`min_first_time` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过max_next_time获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromMaxNextTime(maxNextTime time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`max_next_time` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromMaxNextTime(maxNextTimes []time.Time) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`max_next_time` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromInputBytes 通过input_bytes获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromInputBytes(inputBytes int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`input_bytes` = ?", inputBytes).Find(&results).Error

	return
}

// GetBatchFromInputBytes 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromInputBytes(inputBytess []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`input_bytes` IN (?)", inputBytess).Find(&results).Error

	return
}

// GetFromOutputBytes 通过output_bytes获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromOutputBytes(outputBytes int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`output_bytes` = ?", outputBytes).Find(&results).Error

	return
}

// GetBatchFromOutputBytes 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromOutputBytes(outputBytess []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`output_bytes` IN (?)", outputBytess).Find(&results).Error

	return
}

// GetFromDeletedInputBytes 通过deleted_input_bytes获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromDeletedInputBytes(deletedInputBytes int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`deleted_input_bytes` = ?", deletedInputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedInputBytes 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromDeletedInputBytes(deletedInputBytess []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`deleted_input_bytes` IN (?)", deletedInputBytess).Find(&results).Error

	return
}

// GetFromDeletedOutputBytes 通过deleted_output_bytes获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromDeletedOutputBytes(deletedOutputBytes int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`deleted_output_bytes` = ?", deletedOutputBytes).Find(&results).Error

	return
}

// GetBatchFromDeletedOutputBytes 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromDeletedOutputBytes(deletedOutputBytess []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`deleted_output_bytes` IN (?)", deletedOutputBytess).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromPgCount(pgCount int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromStatus(status string) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsMountFileCreated 通过is_mount_file_created获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromIsMountFileCreated(isMountFileCreated int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`is_mount_file_created` = ?", isMountFileCreated).Find(&results).Error

	return
}

// GetBatchFromIsMountFileCreated 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromIsMountFileCreated(isMountFileCreateds []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`is_mount_file_created` IN (?)", isMountFileCreateds).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetFromCompatible(compatible int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllTenantBackupLogArchiveStatusMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantBackupLogArchiveStatusMgr) FetchByPrimaryKey(tenantID int64, incarnation int64, logArchiveRound int64) (result AllTenantBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupLogArchiveStatus{}).Where("`tenant_id` = ? AND `incarnation` = ? AND `log_archive_round` = ?", tenantID, incarnation, logArchiveRound).First(&result).Error

	return
}
