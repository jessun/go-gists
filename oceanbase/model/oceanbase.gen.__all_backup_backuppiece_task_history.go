package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupBackuppieceTaskHistoryMgr struct {
	*_BaseMgr
}

// AllBackupBackuppieceTaskHistoryMgr open func
func AllBackupBackuppieceTaskHistoryMgr(db *gorm.DB) *_AllBackupBackuppieceTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupBackuppieceTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupBackuppieceTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_backuppiece_task_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetTableName() string {
	return "__all_backup_backuppiece_task_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupBackuppieceTaskHistoryMgr) Reset() *_AllBackupBackuppieceTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) Get() (result AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupBackuppieceTaskHistoryMgr) Gets() (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupBackuppieceTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRoundID round_id获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithRoundID(roundID int64) Option {
	return optionFunc(func(o *options) { o.query["round_id"] = roundID })
}

// WithBackupPieceID backup_piece_id获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_piece_id"] = backupPieceID })
}

// WithCopyID copy_id获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithStartTime start_time获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithStatus status获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithResult result获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithComment comment获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetByOption(opts ...Option) (result AllBackupBackuppieceTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupBackuppieceTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupBackuppieceTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupBackuppieceTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where(options.query)
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
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromJobID(jobID int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoundID 通过round_id获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromRoundID(roundID int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`round_id` = ?", roundID).Find(&results).Error

	return
}

// GetBatchFromRoundID 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromRoundID(roundIDs []int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`round_id` IN (?)", roundIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过backup_piece_id获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromBackupPieceID(backupPieceID int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`backup_piece_id` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`backup_piece_id` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromCopyID(copyID int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromEndTime(endTime time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromStatus(status string) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromBackupDest(backupDest string) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromResult(result int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromResult(results []int64) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetFromComment(comment string) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllBackupBackuppieceTaskHistoryMgr) GetBatchFromComment(comments []string) (results []*AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupBackuppieceTaskHistoryMgr) FetchByPrimaryKey(jobID int64, incarnation int64, tenantID int64, roundID int64, backupPieceID int64, copyID int64) (result AllBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceTaskHistory{}).Where("`job_id` = ? AND `incarnation` = ? AND `tenant_id` = ? AND `round_id` = ? AND `backup_piece_id` = ? AND `copy_id` = ?", jobID, incarnation, tenantID, roundID, backupPieceID, copyID).First(&result).Error

	return
}
