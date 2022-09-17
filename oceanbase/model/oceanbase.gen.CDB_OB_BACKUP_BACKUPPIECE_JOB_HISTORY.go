package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CdbObBackupBackuppieceJobHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupBackuppieceJobHistoryMgr open func
func CdbObBackupBackuppieceJobHistoryMgr(db *gorm.DB) *_CdbObBackupBackuppieceJobHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackuppieceJobHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackuppieceJobHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPPIECE_JOB_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPPIECE_JOB_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) Reset() *_CdbObBackupBackuppieceJobHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) Get() (result CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) Gets() (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupPieceID BACKUP_PIECE_ID获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_ID"] = backupPieceID })
}

// WithMaxBackupTimes MAX_BACKUP_TIMES获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithMaxBackupTimes(maxBackupTimes int64) Option {
	return optionFunc(func(o *options) { o.query["MAX_BACKUP_TIMES"] = maxBackupTimes })
}

// WithResult RESULT获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["RESULT"] = result })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithComment COMMENT获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["COMMENT"] = comment })
}

// WithType TYPE获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupBackuppieceJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackuppieceJobHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where(options.query)
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
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过BACKUP_PIECE_ID获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromBackupPieceID(backupPieceID int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`BACKUP_PIECE_ID` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`BACKUP_PIECE_ID` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromMaxBackupTimes 通过MAX_BACKUP_TIMES获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromMaxBackupTimes(maxBackupTimes int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`MAX_BACKUP_TIMES` = ?", maxBackupTimes).Find(&results).Error

	return
}

// GetBatchFromMaxBackupTimes 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromMaxBackupTimes(maxBackupTimess []int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`MAX_BACKUP_TIMES` IN (?)", maxBackupTimess).Find(&results).Error

	return
}

// GetFromResult 通过RESULT获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromResult(result int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`RESULT` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromResult(results []int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`RESULT` IN (?)", results).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromBackupDest(backupDest string) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromComment 通过COMMENT获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromComment(comment string) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`COMMENT` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromComment(comments []string) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`COMMENT` IN (?)", comments).Find(&results).Error

	return
}

// GetFromType 通过TYPE获取内容
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetFromType(_type int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`TYPE` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_CdbObBackupBackuppieceJobHistoryMgr) GetBatchFromType(_types []int64) (results []*CdbObBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJobHistory{}).Where("`TYPE` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
