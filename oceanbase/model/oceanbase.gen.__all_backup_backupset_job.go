package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllBackupBackupsetJobMgr struct {
	*_BaseMgr
}

// AllBackupBackupsetJobMgr open func
func AllBackupBackupsetJobMgr(db *gorm.DB) *_AllBackupBackupsetJobMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupBackupsetJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupBackupsetJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_backupset_job"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupBackupsetJobMgr) GetTableName() string {
	return "__all_backup_backupset_job"
}

// Reset 重置gorm会话
func (obj *_AllBackupBackupsetJobMgr) Reset() *_AllBackupBackupsetJobMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupBackupsetJobMgr) Get() (result AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupBackupsetJobMgr) Gets() (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupBackupsetJobMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupBackupsetJobMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupBackupsetJobMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllBackupBackupsetJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupBackupsetJobMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupBackupsetJobMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllBackupBackupsetJobMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithCopyID copy_id获取
func (obj *_AllBackupBackupsetJobMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithBackupBackupsetType backup_backupset_type获取
func (obj *_AllBackupBackupsetJobMgr) WithBackupBackupsetType(backupBackupsetType string) Option {
	return optionFunc(func(o *options) { o.query["backup_backupset_type"] = backupBackupsetType })
}

// WithTenantName tenant_name获取
func (obj *_AllBackupBackupsetJobMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithStatus status获取
func (obj *_AllBackupBackupsetJobMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupBackupsetJobMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithMaxBackupTimes max_backup_times获取
func (obj *_AllBackupBackupsetJobMgr) WithMaxBackupTimes(maxBackupTimes int64) Option {
	return optionFunc(func(o *options) { o.query["max_backup_times"] = maxBackupTimes })
}

// WithResult result获取
func (obj *_AllBackupBackupsetJobMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithComment comment获取
func (obj *_AllBackupBackupsetJobMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupBackupsetJobMgr) GetByOption(opts ...Option) (result AllBackupBackupsetJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupBackupsetJobMgr) GetByOptions(opts ...Option) (results []*AllBackupBackupsetJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupBackupsetJobMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupBackupsetJob, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where(options.query)
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
func (obj *_AllBackupBackupsetJobMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromJobID(jobID int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromTenantID(tenantID int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromBackupSetID(backupSetID int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromCopyID(copyID int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupBackupsetType 通过backup_backupset_type获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromBackupBackupsetType(backupBackupsetType string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`backup_backupset_type` = ?", backupBackupsetType).Find(&results).Error

	return
}

// GetBatchFromBackupBackupsetType 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromBackupBackupsetType(backupBackupsetTypes []string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`backup_backupset_type` IN (?)", backupBackupsetTypes).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromTenantName(tenantName string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromStatus(status string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromBackupDest(backupDest string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromMaxBackupTimes 通过max_backup_times获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromMaxBackupTimes(maxBackupTimes int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`max_backup_times` = ?", maxBackupTimes).Find(&results).Error

	return
}

// GetBatchFromMaxBackupTimes 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromMaxBackupTimes(maxBackupTimess []int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`max_backup_times` IN (?)", maxBackupTimess).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromResult(result int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromResult(results []int64) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllBackupBackupsetJobMgr) GetFromComment(comment string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllBackupBackupsetJobMgr) GetBatchFromComment(comments []string) (results []*AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupBackupsetJobMgr) FetchByPrimaryKey(jobID int64, tenantID int64, incarnation int64, backupSetID int64, copyID int64) (result AllBackupBackupsetJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackupsetJob{}).Where("`job_id` = ? AND `tenant_id` = ? AND `incarnation` = ? AND `backup_set_id` = ? AND `copy_id` = ?", jobID, tenantID, incarnation, backupSetID, copyID).First(&result).Error

	return
}
