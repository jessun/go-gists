package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObBackupBackuppieceTaskHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupBackuppieceTaskHistoryMgr open func
func CdbObBackupBackuppieceTaskHistoryMgr(db *gorm.DB) *_CdbObBackupBackuppieceTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackuppieceTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackuppieceTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPPIECE_TASK_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPPIECE_TASK_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) Reset() *_CdbObBackupBackuppieceTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) Get() (result CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) Gets() (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithRoundID ROUND_ID获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithRoundID(roundID int64) Option {
	return optionFunc(func(o *options) { o.query["ROUND_ID"] = roundID })
}

// WithBackupPieceID BACKUP_PIECE_ID获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_ID"] = backupPieceID })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithResult RESULT获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["RESULT"] = result })
}

// WithComment COMMENT获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["COMMENT"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupBackuppieceTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackuppieceTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where(options.query)
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

// GetFromJobID 通过JOB_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoundID 通过ROUND_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromRoundID(roundID int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`ROUND_ID` = ?", roundID).Find(&results).Error

	return
}

// GetBatchFromRoundID 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromRoundID(roundIDs []int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`ROUND_ID` IN (?)", roundIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过BACKUP_PIECE_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromBackupPieceID(backupPieceID int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`BACKUP_PIECE_ID` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`BACKUP_PIECE_ID` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromEndTime(endTime time.Time) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromBackupDest(backupDest string) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromResult 通过RESULT获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromResult(result int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`RESULT` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromResult(results []int64) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`RESULT` IN (?)", results).Find(&results).Error

	return
}

// GetFromComment 通过COMMENT获取内容
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetFromComment(comment string) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`COMMENT` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_CdbObBackupBackuppieceTaskHistoryMgr) GetBatchFromComment(comments []string) (results []*CdbObBackupBackuppieceTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTaskHistory{}).Where("`COMMENT` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
