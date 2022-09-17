package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CdbObBackupBackuppieceJobMgr struct {
	*_BaseMgr
}

// CdbObBackupBackuppieceJobMgr open func
func CdbObBackupBackuppieceJobMgr(db *gorm.DB) *_CdbObBackupBackuppieceJobMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackuppieceJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackuppieceJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPPIECE_JOB"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackuppieceJobMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPPIECE_JOB"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackuppieceJobMgr) Reset() *_CdbObBackupBackuppieceJobMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackuppieceJobMgr) Get() (result CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackuppieceJobMgr) Gets() (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackuppieceJobMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupPieceID BACKUP_PIECE_ID获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_ID"] = backupPieceID })
}

// WithMaxBackupTimes MAX_BACKUP_TIMES获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithMaxBackupTimes(maxBackupTimes int64) Option {
	return optionFunc(func(o *options) { o.query["MAX_BACKUP_TIMES"] = maxBackupTimes })
}

// WithResult RESULT获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["RESULT"] = result })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithComment COMMENT获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["COMMENT"] = comment })
}

// WithType TYPE获取
func (obj *_CdbObBackupBackuppieceJobMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackuppieceJobMgr) GetByOption(opts ...Option) (result CdbObBackupBackuppieceJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackuppieceJobMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackuppieceJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackuppieceJobMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackuppieceJob, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where(options.query)
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
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过BACKUP_PIECE_ID获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromBackupPieceID(backupPieceID int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`BACKUP_PIECE_ID` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`BACKUP_PIECE_ID` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromMaxBackupTimes 通过MAX_BACKUP_TIMES获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromMaxBackupTimes(maxBackupTimes int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`MAX_BACKUP_TIMES` = ?", maxBackupTimes).Find(&results).Error

	return
}

// GetBatchFromMaxBackupTimes 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromMaxBackupTimes(maxBackupTimess []int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`MAX_BACKUP_TIMES` IN (?)", maxBackupTimess).Find(&results).Error

	return
}

// GetFromResult 通过RESULT获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromResult(result int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`RESULT` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromResult(results []int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`RESULT` IN (?)", results).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromStatus(status string) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromBackupDest(backupDest string) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromComment 通过COMMENT获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromComment(comment string) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`COMMENT` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromComment(comments []string) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`COMMENT` IN (?)", comments).Find(&results).Error

	return
}

// GetFromType 通过TYPE获取内容
func (obj *_CdbObBackupBackuppieceJobMgr) GetFromType(_type int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`TYPE` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_CdbObBackupBackuppieceJobMgr) GetBatchFromType(_types []int64) (results []*CdbObBackupBackuppieceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackuppieceJob{}).Where("`TYPE` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
