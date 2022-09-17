package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllKvTTLTaskMgr struct {
	*_BaseMgr
}

// AllKvTTLTaskMgr open func
func AllKvTTLTaskMgr(db *gorm.DB) *_AllKvTTLTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllKvTTLTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllKvTTLTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_kv_ttl_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllKvTTLTaskMgr) GetTableName() string {
	return "__all_kv_ttl_task"
}

// Reset 重置gorm会话
func (obj *_AllKvTTLTaskMgr) Reset() *_AllKvTTLTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllKvTTLTaskMgr) Get() (result AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllKvTTLTaskMgr) Gets() (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllKvTTLTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllKvTTLTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllKvTTLTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllKvTTLTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTaskID task_id获取
func (obj *_AllKvTTLTaskMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithTableID table_id获取
func (obj *_AllKvTTLTaskMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllKvTTLTaskMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithTaskStartTime task_start_time获取
func (obj *_AllKvTTLTaskMgr) WithTaskStartTime(taskStartTime int64) Option {
	return optionFunc(func(o *options) { o.query["task_start_time"] = taskStartTime })
}

// WithTaskUpdateTime task_update_time获取
func (obj *_AllKvTTLTaskMgr) WithTaskUpdateTime(taskUpdateTime int64) Option {
	return optionFunc(func(o *options) { o.query["task_update_time"] = taskUpdateTime })
}

// WithTriggerType trigger_type获取
func (obj *_AllKvTTLTaskMgr) WithTriggerType(triggerType int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_type"] = triggerType })
}

// WithStatus status获取
func (obj *_AllKvTTLTaskMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTTLDelCnt ttl_del_cnt获取
func (obj *_AllKvTTLTaskMgr) WithTTLDelCnt(ttlDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["ttl_del_cnt"] = ttlDelCnt })
}

// WithMaxVersionDelCnt max_version_del_cnt获取
func (obj *_AllKvTTLTaskMgr) WithMaxVersionDelCnt(maxVersionDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["max_version_del_cnt"] = maxVersionDelCnt })
}

// WithScanCnt scan_cnt获取
func (obj *_AllKvTTLTaskMgr) WithScanCnt(scanCnt int64) Option {
	return optionFunc(func(o *options) { o.query["scan_cnt"] = scanCnt })
}

// WithRowKey row_key获取
func (obj *_AllKvTTLTaskMgr) WithRowKey(rowKey []byte) Option {
	return optionFunc(func(o *options) { o.query["row_key"] = rowKey })
}

// WithRetCode ret_code获取
func (obj *_AllKvTTLTaskMgr) WithRetCode(retCode string) Option {
	return optionFunc(func(o *options) { o.query["ret_code"] = retCode })
}

// GetByOption 功能选项模式获取
func (obj *_AllKvTTLTaskMgr) GetByOption(opts ...Option) (result AllKvTTLTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllKvTTLTaskMgr) GetByOptions(opts ...Option) (results []*AllKvTTLTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllKvTTLTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllKvTTLTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where(options.query)
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
func (obj *_AllKvTTLTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllKvTTLTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllKvTTLTaskMgr) GetFromTenantID(tenantID int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllKvTTLTaskMgr) GetFromTaskID(taskID int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllKvTTLTaskMgr) GetFromTableID(tableID int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllKvTTLTaskMgr) GetFromPartitionID(partitionID int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromTaskStartTime 通过task_start_time获取内容
func (obj *_AllKvTTLTaskMgr) GetFromTaskStartTime(taskStartTime int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`task_start_time` = ?", taskStartTime).Find(&results).Error

	return
}

// GetBatchFromTaskStartTime 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromTaskStartTime(taskStartTimes []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`task_start_time` IN (?)", taskStartTimes).Find(&results).Error

	return
}

// GetFromTaskUpdateTime 通过task_update_time获取内容
func (obj *_AllKvTTLTaskMgr) GetFromTaskUpdateTime(taskUpdateTime int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`task_update_time` = ?", taskUpdateTime).Find(&results).Error

	return
}

// GetBatchFromTaskUpdateTime 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromTaskUpdateTime(taskUpdateTimes []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`task_update_time` IN (?)", taskUpdateTimes).Find(&results).Error

	return
}

// GetFromTriggerType 通过trigger_type获取内容
func (obj *_AllKvTTLTaskMgr) GetFromTriggerType(triggerType int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`trigger_type` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromTriggerType(triggerTypes []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`trigger_type` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllKvTTLTaskMgr) GetFromStatus(status int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromStatus(statuss []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTTLDelCnt 通过ttl_del_cnt获取内容
func (obj *_AllKvTTLTaskMgr) GetFromTTLDelCnt(ttlDelCnt int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`ttl_del_cnt` = ?", ttlDelCnt).Find(&results).Error

	return
}

// GetBatchFromTTLDelCnt 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromTTLDelCnt(ttlDelCnts []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`ttl_del_cnt` IN (?)", ttlDelCnts).Find(&results).Error

	return
}

// GetFromMaxVersionDelCnt 通过max_version_del_cnt获取内容
func (obj *_AllKvTTLTaskMgr) GetFromMaxVersionDelCnt(maxVersionDelCnt int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`max_version_del_cnt` = ?", maxVersionDelCnt).Find(&results).Error

	return
}

// GetBatchFromMaxVersionDelCnt 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromMaxVersionDelCnt(maxVersionDelCnts []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`max_version_del_cnt` IN (?)", maxVersionDelCnts).Find(&results).Error

	return
}

// GetFromScanCnt 通过scan_cnt获取内容
func (obj *_AllKvTTLTaskMgr) GetFromScanCnt(scanCnt int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`scan_cnt` = ?", scanCnt).Find(&results).Error

	return
}

// GetBatchFromScanCnt 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromScanCnt(scanCnts []int64) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`scan_cnt` IN (?)", scanCnts).Find(&results).Error

	return
}

// GetFromRowKey 通过row_key获取内容
func (obj *_AllKvTTLTaskMgr) GetFromRowKey(rowKey []byte) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`row_key` = ?", rowKey).Find(&results).Error

	return
}

// GetBatchFromRowKey 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromRowKey(rowKeys [][]byte) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`row_key` IN (?)", rowKeys).Find(&results).Error

	return
}

// GetFromRetCode 通过ret_code获取内容
func (obj *_AllKvTTLTaskMgr) GetFromRetCode(retCode string) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`ret_code` = ?", retCode).Find(&results).Error

	return
}

// GetBatchFromRetCode 批量查找
func (obj *_AllKvTTLTaskMgr) GetBatchFromRetCode(retCodes []string) (results []*AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`ret_code` IN (?)", retCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllKvTTLTaskMgr) FetchByPrimaryKey(tenantID int64, taskID int64, tableID int64, partitionID int64) (result AllKvTTLTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTask{}).Where("`tenant_id` = ? AND `task_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, taskID, tableID, partitionID).First(&result).Error

	return
}
