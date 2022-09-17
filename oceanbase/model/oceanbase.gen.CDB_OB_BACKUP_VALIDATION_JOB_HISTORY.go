package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CdbObBackupValidationJobHistoryMgr struct {
	*_BaseMgr
}

// CdbObBackupValidationJobHistoryMgr open func
func CdbObBackupValidationJobHistoryMgr(db *gorm.DB) *_CdbObBackupValidationJobHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupValidationJobHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupValidationJobHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_VALIDATION_JOB_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupValidationJobHistoryMgr) GetTableName() string {
	return "CDB_OB_BACKUP_VALIDATION_JOB_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupValidationJobHistoryMgr) Reset() *_CdbObBackupValidationJobHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupValidationJobHistoryMgr) Get() (result CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupValidationJobHistoryMgr) Gets() (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupValidationJobHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID JOB_ID获取
func (obj *_CdbObBackupValidationJobHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupValidationJobHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTenantName TENANT_NAME获取
func (obj *_CdbObBackupValidationJobHistoryMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_NAME"] = tenantName })
}

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupValidationJobHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithBackupSetID BACKUP_SET_ID获取
func (obj *_CdbObBackupValidationJobHistoryMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_SET_ID"] = backupSetID })
}

// WithProgressPercent PROGRESS_PERCENT获取
func (obj *_CdbObBackupValidationJobHistoryMgr) WithProgressPercent(progressPercent int64) Option {
	return optionFunc(func(o *options) { o.query["PROGRESS_PERCENT"] = progressPercent })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupValidationJobHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupValidationJobHistoryMgr) GetByOption(opts ...Option) (result CdbObBackupValidationJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupValidationJobHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObBackupValidationJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupValidationJobHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupValidationJobHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where(options.query)
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
func (obj *_CdbObBackupValidationJobHistoryMgr) GetFromJobID(jobID int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_CdbObBackupValidationJobHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupValidationJobHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupValidationJobHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过TENANT_NAME获取内容
func (obj *_CdbObBackupValidationJobHistoryMgr) GetFromTenantName(tenantName string) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`TENANT_NAME` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_CdbObBackupValidationJobHistoryMgr) GetBatchFromTenantName(tenantNames []string) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`TENANT_NAME` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupValidationJobHistoryMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupValidationJobHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过BACKUP_SET_ID获取内容
func (obj *_CdbObBackupValidationJobHistoryMgr) GetFromBackupSetID(backupSetID int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`BACKUP_SET_ID` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_CdbObBackupValidationJobHistoryMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`BACKUP_SET_ID` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromProgressPercent 通过PROGRESS_PERCENT获取内容
func (obj *_CdbObBackupValidationJobHistoryMgr) GetFromProgressPercent(progressPercent int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`PROGRESS_PERCENT` = ?", progressPercent).Find(&results).Error

	return
}

// GetBatchFromProgressPercent 批量查找
func (obj *_CdbObBackupValidationJobHistoryMgr) GetBatchFromProgressPercent(progressPercents []int64) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`PROGRESS_PERCENT` IN (?)", progressPercents).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupValidationJobHistoryMgr) GetFromStatus(status string) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupValidationJobHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupValidationJobHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
