package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CdbObBackupBackupsetJobHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupBackupsetJobHistoryMgr open func
func CdbObBackupBackupsetJobHistoryMgr(db *gorm.DB) *_CdbObBackupBackupsetJobHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackupsetJobHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackupsetJobHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPSET_JOB_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPSET_JOB_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackupsetJobHistoryMgr) Reset() *_CdbObBackupBackupsetJobHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) Get() (result CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackupsetJobHistoryMgr) Gets() (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackupsetJobHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupSetID BACKUP_SET_ID获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_SET_ID"] = backupSetID })
}

// WithType TYPE获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithTenantName TENANT_NAME获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_NAME"] = tenantName })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithMaxBackupTimes MAX_BACKUP_TIMES获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithMaxBackupTimes(maxBackupTimes int64) Option {
	return optionFunc(func(o *options) { o.query["MAX_BACKUP_TIMES"] = maxBackupTimes })
}

// WithResult RESULT获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["RESULT"] = result })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupBackupsetJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackupsetJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackupsetJobHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackupsetJobHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where(options.query)
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
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过BACKUP_SET_ID获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromBackupSetID(backupSetID int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`BACKUP_SET_ID` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`BACKUP_SET_ID` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromType 通过TYPE获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromType(_type string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`TYPE` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromType(_types []string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`TYPE` IN (?)", _types).Find(&results).Error

	return
}

// GetFromTenantName 通过TENANT_NAME获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromTenantName(tenantName string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`TENANT_NAME` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromTenantName(tenantNames []string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`TENANT_NAME` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromBackupDest(backupDest string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromMaxBackupTimes 通过MAX_BACKUP_TIMES获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromMaxBackupTimes(maxBackupTimes int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`MAX_BACKUP_TIMES` = ?", maxBackupTimes).Find(&results).Error

	return
}

// GetBatchFromMaxBackupTimes 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromMaxBackupTimes(maxBackupTimess []int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`MAX_BACKUP_TIMES` IN (?)", maxBackupTimess).Find(&results).Error

	return
}

// GetFromResult 通过RESULT获取内容
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetFromResult(result int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`RESULT` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_CdbObBackupBackupsetJobHistoryMgr) GetBatchFromResult(results []int64) (results []*CdbObBackupBackupsetJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJobHistory{}).Where("`RESULT` IN (?)", results).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
