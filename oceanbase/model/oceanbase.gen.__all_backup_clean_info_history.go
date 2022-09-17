package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupCleanInfoHistoryMgr struct {
	*_BaseMgr
}

// AllBackupCleanInfoHistoryMgr open func
func AllBackupCleanInfoHistoryMgr(db *gorm.DB) *_AllBackupCleanInfoHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupCleanInfoHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupCleanInfoHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_clean_info_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupCleanInfoHistoryMgr) GetTableName() string {
	return "__all_backup_clean_info_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupCleanInfoHistoryMgr) Reset() *_AllBackupCleanInfoHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupCleanInfoHistoryMgr) Get() (result AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupCleanInfoHistoryMgr) Gets() (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupCleanInfoHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithStartTime start_time获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithType type获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithStatus status获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithParameter parameter获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithParameter(parameter string) Option {
	return optionFunc(func(o *options) { o.query["parameter"] = parameter })
}

// WithErrorMsg error_msg获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithErrorMsg(errorMsg string) Option {
	return optionFunc(func(o *options) { o.query["error_msg"] = errorMsg })
}

// WithComment comment获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithClogGcSnapshot clog_gc_snapshot获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithClogGcSnapshot(clogGcSnapshot int64) Option {
	return optionFunc(func(o *options) { o.query["clog_gc_snapshot"] = clogGcSnapshot })
}

// WithResult result获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithCopyID copy_id获取
func (obj *_AllBackupCleanInfoHistoryMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupCleanInfoHistoryMgr) GetByOption(opts ...Option) (result AllBackupCleanInfoHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupCleanInfoHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupCleanInfoHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupCleanInfoHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupCleanInfoHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where(options.query)
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
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromJobID(jobID int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromStartTime(startTime time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromEndTime(endTime time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromType(_type string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromType(_types []string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromStatus(status string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromParameter 通过parameter获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromParameter(parameter string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`parameter` = ?", parameter).Find(&results).Error

	return
}

// GetBatchFromParameter 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromParameter(parameters []string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`parameter` IN (?)", parameters).Find(&results).Error

	return
}

// GetFromErrorMsg 通过error_msg获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromErrorMsg(errorMsg string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`error_msg` = ?", errorMsg).Find(&results).Error

	return
}

// GetBatchFromErrorMsg 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromErrorMsg(errorMsgs []string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`error_msg` IN (?)", errorMsgs).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromComment(comment string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromComment(comments []string) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromClogGcSnapshot 通过clog_gc_snapshot获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromClogGcSnapshot(clogGcSnapshot int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`clog_gc_snapshot` = ?", clogGcSnapshot).Find(&results).Error

	return
}

// GetBatchFromClogGcSnapshot 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromClogGcSnapshot(clogGcSnapshots []int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`clog_gc_snapshot` IN (?)", clogGcSnapshots).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromResult(result int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromResult(results []int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupCleanInfoHistoryMgr) GetFromCopyID(copyID int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupCleanInfoHistoryMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupCleanInfoHistoryMgr) FetchByPrimaryKey(tenantID int64, jobID int64) (result AllBackupCleanInfoHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupCleanInfoHistory{}).Where("`tenant_id` = ? AND `job_id` = ?", tenantID, jobID).First(&result).Error

	return
}
