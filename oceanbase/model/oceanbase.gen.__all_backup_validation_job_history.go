package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllBackupValidationJobHistoryMgr struct {
	*_BaseMgr
}

// AllBackupValidationJobHistoryMgr open func
func AllBackupValidationJobHistoryMgr(db *gorm.DB) *_AllBackupValidationJobHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupValidationJobHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupValidationJobHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_validation_job_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupValidationJobHistoryMgr) GetTableName() string {
	return "__all_backup_validation_job_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupValidationJobHistoryMgr) Reset() *_AllBackupValidationJobHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupValidationJobHistoryMgr) Get() (result AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupValidationJobHistoryMgr) Gets() (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupValidationJobHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupValidationJobHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupValidationJobHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllBackupValidationJobHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupValidationJobHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupValidationJobHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllBackupValidationJobHistoryMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithTenantName tenant_name获取
func (obj *_AllBackupValidationJobHistoryMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithProgressPercent progress_percent获取
func (obj *_AllBackupValidationJobHistoryMgr) WithProgressPercent(progressPercent int64) Option {
	return optionFunc(func(o *options) { o.query["progress_percent"] = progressPercent })
}

// WithStatus status获取
func (obj *_AllBackupValidationJobHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupValidationJobHistoryMgr) GetByOption(opts ...Option) (result AllBackupValidationJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupValidationJobHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupValidationJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupValidationJobHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupValidationJobHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where(options.query)
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
func (obj *_AllBackupValidationJobHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromJobID(jobID int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromBackupSetID(backupSetID int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromTenantName(tenantName string) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromProgressPercent 通过progress_percent获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromProgressPercent(progressPercent int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`progress_percent` = ?", progressPercent).Find(&results).Error

	return
}

// GetBatchFromProgressPercent 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromProgressPercent(progressPercents []int64) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`progress_percent` IN (?)", progressPercents).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupValidationJobHistoryMgr) GetFromStatus(status string) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupValidationJobHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupValidationJobHistoryMgr) FetchByPrimaryKey(jobID int64, tenantID int64, incarnation int64, backupSetID int64) (result AllBackupValidationJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupValidationJobHistory{}).Where("`job_id` = ? AND `tenant_id` = ? AND `incarnation` = ? AND `backup_set_id` = ?", jobID, tenantID, incarnation, backupSetID).First(&result).Error

	return
}
