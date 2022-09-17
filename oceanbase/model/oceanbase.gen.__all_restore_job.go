package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllRestoreJobMgr struct {
	*_BaseMgr
}

// AllRestoreJobMgr open func
func AllRestoreJobMgr(db *gorm.DB) *_AllRestoreJobMgr {
	if db == nil {
		panic(fmt.Errorf("AllRestoreJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRestoreJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_restore_job"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRestoreJobMgr) GetTableName() string {
	return "__all_restore_job"
}

// Reset 重置gorm会话
func (obj *_AllRestoreJobMgr) Reset() *_AllRestoreJobMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRestoreJobMgr) Get() (result AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRestoreJobMgr) Gets() (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRestoreJobMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRestoreJobMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllRestoreJobMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllRestoreJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTenantName tenant_name获取
func (obj *_AllRestoreJobMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithStartTime start_time获取
func (obj *_AllRestoreJobMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithBackupURI backup_uri获取
func (obj *_AllRestoreJobMgr) WithBackupURI(backupURI string) Option {
	return optionFunc(func(o *options) { o.query["backup_uri"] = backupURI })
}

// WithBackupEndTime backup_end_time获取
func (obj *_AllRestoreJobMgr) WithBackupEndTime(backupEndTime int64) Option {
	return optionFunc(func(o *options) { o.query["backup_end_time"] = backupEndTime })
}

// WithRecycleEndTime recycle_end_time获取
func (obj *_AllRestoreJobMgr) WithRecycleEndTime(recycleEndTime int64) Option {
	return optionFunc(func(o *options) { o.query["recycle_end_time"] = recycleEndTime })
}

// WithLevel level获取
func (obj *_AllRestoreJobMgr) WithLevel(level int64) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithStatus status获取
func (obj *_AllRestoreJobMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllRestoreJobMgr) GetByOption(opts ...Option) (result AllRestoreJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRestoreJobMgr) GetByOptions(opts ...Option) (results []*AllRestoreJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRestoreJobMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRestoreJob, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where(options.query)
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
func (obj *_AllRestoreJobMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllRestoreJobMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllRestoreJobMgr) GetFromJobID(jobID int64) (result AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllRestoreJobMgr) GetFromTenantName(tenantName string) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllRestoreJobMgr) GetFromStartTime(startTime int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromStartTime(startTimes []int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromBackupURI 通过backup_uri获取内容
func (obj *_AllRestoreJobMgr) GetFromBackupURI(backupURI string) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`backup_uri` = ?", backupURI).Find(&results).Error

	return
}

// GetBatchFromBackupURI 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromBackupURI(backupURIs []string) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`backup_uri` IN (?)", backupURIs).Find(&results).Error

	return
}

// GetFromBackupEndTime 通过backup_end_time获取内容
func (obj *_AllRestoreJobMgr) GetFromBackupEndTime(backupEndTime int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`backup_end_time` = ?", backupEndTime).Find(&results).Error

	return
}

// GetBatchFromBackupEndTime 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromBackupEndTime(backupEndTimes []int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`backup_end_time` IN (?)", backupEndTimes).Find(&results).Error

	return
}

// GetFromRecycleEndTime 通过recycle_end_time获取内容
func (obj *_AllRestoreJobMgr) GetFromRecycleEndTime(recycleEndTime int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`recycle_end_time` = ?", recycleEndTime).Find(&results).Error

	return
}

// GetBatchFromRecycleEndTime 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromRecycleEndTime(recycleEndTimes []int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`recycle_end_time` IN (?)", recycleEndTimes).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容
func (obj *_AllRestoreJobMgr) GetFromLevel(level int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromLevel(levels []int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllRestoreJobMgr) GetFromStatus(status int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllRestoreJobMgr) GetBatchFromStatus(statuss []int64) (results []*AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRestoreJobMgr) FetchByPrimaryKey(jobID int64) (result AllRestoreJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreJob{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}
