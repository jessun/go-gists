package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CdbObBackupBackuppieceTaskMgr struct {
	*_BaseMgr
}

// CdbObBackupBackuppieceTaskMgr open func
func CdbObBackupBackuppieceTaskMgr(db *gorm.DB) *_CdbObBackupBackuppieceTaskMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackuppieceTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackuppieceTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPPIECE_TASK"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackuppieceTaskMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPPIECE_TASK"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackuppieceTaskMgr) Reset() *_CdbObBackupBackuppieceTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackuppieceTaskMgr) Get() (result CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackuppieceTaskMgr) Gets() (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackuppieceTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithRoundID ROUND_ID获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithRoundID(roundID int64) Option {
	return optionFunc(func(o *options) { o.query["ROUND_ID"] = roundID })
}

// WithBackupPieceID BACKUP_PIECE_ID获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_ID"] = backupPieceID })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithStartTime START_TIME获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithResult RESULT获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["RESULT"] = result })
}

// WithComment COMMENT获取
func (obj *_CdbObBackupBackuppieceTaskMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["COMMENT"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackuppieceTaskMgr) GetByOption(opts ...Option) (result CdbObBackupBackuppieceTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackuppieceTaskMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackuppieceTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackuppieceTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackuppieceTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where(options.query)
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
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoundID 通过ROUND_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromRoundID(roundID int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`ROUND_ID` = ?", roundID).Find(&results).Error

	return
}

// GetBatchFromRoundID 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromRoundID(roundIDs []int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`ROUND_ID` IN (?)", roundIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过BACKUP_PIECE_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromBackupPieceID(backupPieceID int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`BACKUP_PIECE_ID` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`BACKUP_PIECE_ID` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromStartTime(startTime time.Time) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromEndTime(endTime time.Time) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromStatus(status string) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromBackupDest(backupDest string) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromResult 通过RESULT获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromResult(result int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`RESULT` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromResult(results []int64) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`RESULT` IN (?)", results).Find(&results).Error

	return
}

// GetFromComment 通过COMMENT获取内容
func (obj *_CdbObBackupBackuppieceTaskMgr) GetFromComment(comment string) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`COMMENT` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_CdbObBackupBackuppieceTaskMgr) GetBatchFromComment(comments []string) (results []*CdbObBackupBackuppieceTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceTask{}).Where("`COMMENT` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
