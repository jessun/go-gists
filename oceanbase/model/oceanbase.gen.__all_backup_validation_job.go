package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllBackupValidationJobMgr struct {
	*_BaseMgr
}

// AllBackupValidationJobMgr open func
func AllBackupValidationJobMgr(db *gorm.DB) *_AllBackupValidationJobMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupValidationJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupValidationJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_validation_job"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupValidationJobMgr) GetTableName() string {
	return "__all_backup_validation_job"
}

// Reset 重置gorm会话
func (obj *_AllBackupValidationJobMgr) Reset() *_AllBackupValidationJobMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupValidationJobMgr) Get() (result AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupValidationJobMgr) Gets() (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupValidationJobMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupValidationJobMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupValidationJobMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllBackupValidationJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupValidationJobMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupValidationJobMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllBackupValidationJobMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithTenantName tenant_name获取
func (obj *_AllBackupValidationJobMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithProgressPercent progress_percent获取
func (obj *_AllBackupValidationJobMgr) WithProgressPercent(progressPercent int64) Option {
	return optionFunc(func(o *options) { o.query["progress_percent"] = progressPercent })
}

// WithStatus status获取
func (obj *_AllBackupValidationJobMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupValidationJobMgr) GetByOption(opts ...Option) (result AllBackupValidationJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupValidationJobMgr) GetByOptions(opts ...Option) (results []*AllBackupValidationJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupValidationJobMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupValidationJob, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where(options.query)
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
func (obj *_AllBackupValidationJobMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupValidationJobMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupValidationJobMgr) GetFromJobID(jobID int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupValidationJobMgr) GetFromTenantID(tenantID int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupValidationJobMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllBackupValidationJobMgr) GetFromBackupSetID(backupSetID int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllBackupValidationJobMgr) GetFromTenantName(tenantName string) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromProgressPercent 通过progress_percent获取内容
func (obj *_AllBackupValidationJobMgr) GetFromProgressPercent(progressPercent int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`progress_percent` = ?", progressPercent).Find(&results).Error

	return
}

// GetBatchFromProgressPercent 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromProgressPercent(progressPercents []int64) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`progress_percent` IN (?)", progressPercents).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupValidationJobMgr) GetFromStatus(status string) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupValidationJobMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupValidationJobMgr) FetchByPrimaryKey(jobID int64, tenantID int64, incarnation int64, backupSetID int64) (result AllBackupValidationJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJob{}).Where("`job_id` = ? AND `tenant_id` = ? AND `incarnation` = ? AND `backup_set_id` = ?", jobID, tenantID, incarnation, backupSetID).First(&result).Error

	return
}
