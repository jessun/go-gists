package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CdbObBackupValidationJobMgr struct {
	*_BaseMgr
}

// CdbObBackupValidationJobMgr open func
func CdbObBackupValidationJobMgr(db *gorm.DB) *_CdbObBackupValidationJobMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupValidationJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupValidationJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_VALIDATION_JOB"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupValidationJobMgr) GetTableName() string {
	return "CDB_OB_BACKUP_VALIDATION_JOB"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupValidationJobMgr) Reset() *_CdbObBackupValidationJobMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupValidationJobMgr) Get() (result CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupValidationJobMgr) Gets() (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupValidationJobMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupValidationJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupValidationJobMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTenantName TENANT_NAME获取
func (obj *_CdbObBackupValidationJobMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_NAME"] = tenantName })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupValidationJobMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupSetID BACKUP_SET_ID获取
func (obj *_CdbObBackupValidationJobMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_SET_ID"] = backupSetID })
}

// WithProgressPercent PROGRESS_PERCENT获取
func (obj *_CdbObBackupValidationJobMgr) WithProgressPercent(progressPercent int64) Option {
	return optionFunc(func(o *options) { o.query["PROGRESS_PERCENT"] = progressPercent })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupValidationJobMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupValidationJobMgr) GetByOption(opts ...Option) (result CdbObBackupValidationJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupValidationJobMgr) GetByOptions(opts ...Option) (results []*CdbObBackupValidationJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupValidationJobMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupValidationJob, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where(options.query)
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
func (obj *_CdbObBackupValidationJobMgr) GetFromJobID(jobID int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupValidationJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupValidationJobMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupValidationJobMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过TENANT_NAME获取内容
func (obj *_CdbObBackupValidationJobMgr) GetFromTenantName(tenantName string) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`TENANT_NAME` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_CdbObBackupValidationJobMgr) GetBatchFromTenantName(tenantNames []string) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`TENANT_NAME` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupValidationJobMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupValidationJobMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过BACKUP_SET_ID获取内容
func (obj *_CdbObBackupValidationJobMgr) GetFromBackupSetID(backupSetID int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`BACKUP_SET_ID` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_CdbObBackupValidationJobMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`BACKUP_SET_ID` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromProgressPercent 通过PROGRESS_PERCENT获取内容
func (obj *_CdbObBackupValidationJobMgr) GetFromProgressPercent(progressPercent int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`PROGRESS_PERCENT` = ?", progressPercent).Find(&results).Error

	return
}

// GetBatchFromProgressPercent 批量查找
func (obj *_CdbObBackupValidationJobMgr) GetBatchFromProgressPercent(progressPercents []int64) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`PROGRESS_PERCENT` IN (?)", progressPercents).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupValidationJobMgr) GetFromStatus(status string) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupValidationJobMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJob{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
