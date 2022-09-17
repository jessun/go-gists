package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualBackupCleanInfoMgr struct {
	*_BaseMgr
}

// AllVirtualBackupCleanInfoMgr open func
func AllVirtualBackupCleanInfoMgr(db *gorm.DB) *_AllVirtualBackupCleanInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualBackupCleanInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualBackupCleanInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_backup_clean_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualBackupCleanInfoMgr) GetTableName() string {
	return "__all_virtual_backup_clean_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualBackupCleanInfoMgr) Reset() *_AllVirtualBackupCleanInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualBackupCleanInfoMgr) Get() (result AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualBackupCleanInfoMgr) Gets() (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualBackupCleanInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithStartTime start_time获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithType type获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithStatus status获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithParameter parameter获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithParameter(parameter int64) Option {
	return optionFunc(func(o *options) { o.query["parameter"] = parameter })
}

// WithErrorMsg error_msg获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithErrorMsg(errorMsg string) Option {
	return optionFunc(func(o *options) { o.query["error_msg"] = errorMsg })
}

// WithComment comment获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithClogGcSnapshot clog_gc_snapshot获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithClogGcSnapshot(clogGcSnapshot int64) Option {
	return optionFunc(func(o *options) { o.query["clog_gc_snapshot"] = clogGcSnapshot })
}

// WithResult result获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithCopyID copy_id获取
func (obj *_AllVirtualBackupCleanInfoMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualBackupCleanInfoMgr) GetByOption(opts ...Option) (result AllVirtualBackupCleanInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualBackupCleanInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualBackupCleanInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualBackupCleanInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualBackupCleanInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromTenantID(tenantID int64) (result AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromJobID(jobID int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromStartTime(startTime int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromStartTime(startTimes []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromEndTime(endTime int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromEndTime(endTimes []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromType(_type string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromType(_types []string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromStatus(status string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromParameter 通过parameter获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromParameter(parameter int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`parameter` = ?", parameter).Find(&results).Error

	return
}

// GetBatchFromParameter 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromParameter(parameters []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`parameter` IN (?)", parameters).Find(&results).Error

	return
}

// GetFromErrorMsg 通过error_msg获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromErrorMsg(errorMsg string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`error_msg` = ?", errorMsg).Find(&results).Error

	return
}

// GetBatchFromErrorMsg 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromErrorMsg(errorMsgs []string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`error_msg` IN (?)", errorMsgs).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromComment(comment string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromComment(comments []string) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromClogGcSnapshot 通过clog_gc_snapshot获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromClogGcSnapshot(clogGcSnapshot int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`clog_gc_snapshot` = ?", clogGcSnapshot).Find(&results).Error

	return
}

// GetBatchFromClogGcSnapshot 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromClogGcSnapshot(clogGcSnapshots []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`clog_gc_snapshot` IN (?)", clogGcSnapshots).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromResult(result int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromResult(results []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllVirtualBackupCleanInfoMgr) GetFromCopyID(copyID int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllVirtualBackupCleanInfoMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualBackupCleanInfoMgr) FetchByPrimaryKey(tenantID int64) (result AllVirtualBackupCleanInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBackupCleanInfo{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
