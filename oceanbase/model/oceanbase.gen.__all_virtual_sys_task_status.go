package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSysTaskStatusMgr struct {
	*_BaseMgr
}

// AllVirtualSysTaskStatusMgr open func
func AllVirtualSysTaskStatusMgr(db *gorm.DB) *_AllVirtualSysTaskStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysTaskStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysTaskStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sys_task_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysTaskStatusMgr) GetTableName() string {
	return "__all_virtual_sys_task_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysTaskStatusMgr) Reset() *_AllVirtualSysTaskStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSysTaskStatusMgr) Get() (result AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysTaskStatusMgr) Gets() (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysTaskStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithStartTime start_time获取
func (obj *_AllVirtualSysTaskStatusMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithTaskType task_type获取
func (obj *_AllVirtualSysTaskStatusMgr) WithTaskType(taskType string) Option {
	return optionFunc(func(o *options) { o.query["task_type"] = taskType })
}

// WithTaskID task_id获取
func (obj *_AllVirtualSysTaskStatusMgr) WithTaskID(taskID string) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSysTaskStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSysTaskStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualSysTaskStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithComment comment获取
func (obj *_AllVirtualSysTaskStatusMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithIsCancel is_cancel获取
func (obj *_AllVirtualSysTaskStatusMgr) WithIsCancel(isCancel int64) Option {
	return optionFunc(func(o *options) { o.query["is_cancel"] = isCancel })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysTaskStatusMgr) GetByOption(opts ...Option) (result AllVirtualSysTaskStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysTaskStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysTaskStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSysTaskStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysTaskStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where(options.query)
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

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromStartTime(startTime time.Time) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromTaskType 通过task_type获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromTaskType(taskType string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`task_type` = ?", taskType).Find(&results).Error

	return
}

// GetBatchFromTaskType 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromTaskType(taskTypes []string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`task_type` IN (?)", taskTypes).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromTaskID(taskID string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromTaskID(taskIDs []string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromComment(comment string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromComment(comments []string) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromIsCancel 通过is_cancel获取内容
func (obj *_AllVirtualSysTaskStatusMgr) GetFromIsCancel(isCancel int64) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`is_cancel` = ?", isCancel).Find(&results).Error

	return
}

// GetBatchFromIsCancel 批量查找
func (obj *_AllVirtualSysTaskStatusMgr) GetBatchFromIsCancel(isCancels []int64) (results []*AllVirtualSysTaskStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysTaskStatus{}).Where("`is_cancel` IN (?)", isCancels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
