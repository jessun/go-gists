package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CdbObBackupBackupsetJobMgr struct {
	*_BaseMgr
}

// CdbObBackupBackupsetJobMgr open func
func CdbObBackupBackupsetJobMgr(db *gorm.DB) *_CdbObBackupBackupsetJobMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupBackupsetJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupBackupsetJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_BACKUPSET_JOB"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupBackupsetJobMgr) GetTableName() string {
	return "CDB_OB_BACKUP_BACKUPSET_JOB"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupBackupsetJobMgr) Reset() *_CdbObBackupBackupsetJobMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupBackupsetJobMgr) Get() (result CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupBackupsetJobMgr) Gets() (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupBackupsetJobMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupBackupsetJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupBackupsetJobMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupBackupsetJobMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupSetID BACKUP_SET_ID获取
func (obj *_CdbObBackupBackupsetJobMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_SET_ID"] = backupSetID })
}

// WithType TYPE获取
func (obj *_CdbObBackupBackupsetJobMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithTenantName TENANT_NAME获取
func (obj *_CdbObBackupBackupsetJobMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_NAME"] = tenantName })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupBackupsetJobMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithBackupDest BACKUP_DEST获取
func (obj *_CdbObBackupBackupsetJobMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_DEST"] = backupDest })
}

// WithMaxBackupTimes MAX_BACKUP_TIMES获取
func (obj *_CdbObBackupBackupsetJobMgr) WithMaxBackupTimes(maxBackupTimes int64) Option {
	return optionFunc(func(o *options) { o.query["MAX_BACKUP_TIMES"] = maxBackupTimes })
}

// WithResult RESULT获取
func (obj *_CdbObBackupBackupsetJobMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["RESULT"] = result })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupBackupsetJobMgr) GetByOption(opts ...Option) (result CdbObBackupBackupsetJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupBackupsetJobMgr) GetByOptions(opts ...Option) (results []*CdbObBackupBackupsetJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupBackupsetJobMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupBackupsetJob, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where(options.query)
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
func (obj *_CdbObBackupBackupsetJobMgr) GetFromJobID(jobID int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过BACKUP_SET_ID获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromBackupSetID(backupSetID int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`BACKUP_SET_ID` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`BACKUP_SET_ID` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromType 通过TYPE获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromType(_type string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`TYPE` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromType(_types []string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`TYPE` IN (?)", _types).Find(&results).Error

	return
}

// GetFromTenantName 通过TENANT_NAME获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromTenantName(tenantName string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`TENANT_NAME` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromTenantName(tenantNames []string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`TENANT_NAME` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromStatus(status string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过BACKUP_DEST获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromBackupDest(backupDest string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`BACKUP_DEST` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromBackupDest(backupDests []string) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`BACKUP_DEST` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromMaxBackupTimes 通过MAX_BACKUP_TIMES获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromMaxBackupTimes(maxBackupTimes int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`MAX_BACKUP_TIMES` = ?", maxBackupTimes).Find(&results).Error

	return
}

// GetBatchFromMaxBackupTimes 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromMaxBackupTimes(maxBackupTimess []int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`MAX_BACKUP_TIMES` IN (?)", maxBackupTimess).Find(&results).Error

	return
}

// GetFromResult 通过RESULT获取内容
func (obj *_CdbObBackupBackupsetJobMgr) GetFromResult(result int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`RESULT` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_CdbObBackupBackupsetJobMgr) GetBatchFromResult(results []int64) (results []*CdbObBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupBackupsetJob{}).Where("`RESULT` IN (?)", results).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
