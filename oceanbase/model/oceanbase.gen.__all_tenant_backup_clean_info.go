package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantBackupCleanInfoMgr struct {
	*_BaseMgr
}

// AllTenantBackupCleanInfoMgr open func
func AllTenantBackupCleanInfoMgr(db *gorm.DB) *_AllTenantBackupCleanInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantBackupCleanInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantBackupCleanInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_backup_clean_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantBackupCleanInfoMgr) GetTableName() string {
	return "__all_tenant_backup_clean_info"
}

// Reset 重置gorm会话
func (obj *_AllTenantBackupCleanInfoMgr) Reset() *_AllTenantBackupCleanInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantBackupCleanInfoMgr) Get() (result AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantBackupCleanInfoMgr) Gets() (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantBackupCleanInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantBackupCleanInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantBackupCleanInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantBackupCleanInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllTenantBackupCleanInfoMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithStartTime start_time获取
func (obj *_AllTenantBackupCleanInfoMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllTenantBackupCleanInfoMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithIncarnation incarnation获取
func (obj *_AllTenantBackupCleanInfoMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithType type获取
func (obj *_AllTenantBackupCleanInfoMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithStatus status获取
func (obj *_AllTenantBackupCleanInfoMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithParameter parameter获取
func (obj *_AllTenantBackupCleanInfoMgr) WithParameter(parameter string) Option {
	return optionFunc(func(o *options) { o.query["parameter"] = parameter })
}

// WithErrorMsg error_msg获取
func (obj *_AllTenantBackupCleanInfoMgr) WithErrorMsg(errorMsg string) Option {
	return optionFunc(func(o *options) { o.query["error_msg"] = errorMsg })
}

// WithComment comment获取
func (obj *_AllTenantBackupCleanInfoMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithClogGcSnapshot clog_gc_snapshot获取
func (obj *_AllTenantBackupCleanInfoMgr) WithClogGcSnapshot(clogGcSnapshot int64) Option {
	return optionFunc(func(o *options) { o.query["clog_gc_snapshot"] = clogGcSnapshot })
}

// WithResult result获取
func (obj *_AllTenantBackupCleanInfoMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithCopyID copy_id获取
func (obj *_AllTenantBackupCleanInfoMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantBackupCleanInfoMgr) GetByOption(opts ...Option) (result AllTenantBackupCleanInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantBackupCleanInfoMgr) GetByOptions(opts ...Option) (results []*AllTenantBackupCleanInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantBackupCleanInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantBackupCleanInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where(options.query)
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
func (obj *_AllTenantBackupCleanInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromTenantID(tenantID int64) (result AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromJobID(jobID int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromStartTime(startTime time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromEndTime(endTime time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromIncarnation(incarnation int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromType(_type string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromType(_types []string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromStatus(status string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromParameter 通过parameter获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromParameter(parameter string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`parameter` = ?", parameter).Find(&results).Error

	return
}

// GetBatchFromParameter 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromParameter(parameters []string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`parameter` IN (?)", parameters).Find(&results).Error

	return
}

// GetFromErrorMsg 通过error_msg获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromErrorMsg(errorMsg string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`error_msg` = ?", errorMsg).Find(&results).Error

	return
}

// GetBatchFromErrorMsg 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromErrorMsg(errorMsgs []string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`error_msg` IN (?)", errorMsgs).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromComment(comment string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromComment(comments []string) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromClogGcSnapshot 通过clog_gc_snapshot获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromClogGcSnapshot(clogGcSnapshot int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`clog_gc_snapshot` = ?", clogGcSnapshot).Find(&results).Error

	return
}

// GetBatchFromClogGcSnapshot 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromClogGcSnapshot(clogGcSnapshots []int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`clog_gc_snapshot` IN (?)", clogGcSnapshots).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromResult(result int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromResult(results []int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllTenantBackupCleanInfoMgr) GetFromCopyID(copyID int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllTenantBackupCleanInfoMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantBackupCleanInfoMgr) FetchByPrimaryKey(tenantID int64) (result AllTenantBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantBackupCleanInfo{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
