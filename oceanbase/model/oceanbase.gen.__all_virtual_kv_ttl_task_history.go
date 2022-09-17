package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualKvTTLTaskHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualKvTTLTaskHistoryMgr open func
func AllVirtualKvTTLTaskHistoryMgr(db *gorm.DB) *_AllVirtualKvTTLTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualKvTTLTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualKvTTLTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_kv_ttl_task_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetTableName() string {
	return "__all_virtual_kv_ttl_task_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualKvTTLTaskHistoryMgr) Reset() *_AllVirtualKvTTLTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) Get() (result AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualKvTTLTaskHistoryMgr) Gets() (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualKvTTLTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTaskID task_id获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithTableID table_id获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTaskStartTime task_start_time获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithTaskStartTime(taskStartTime int64) Option {
	return optionFunc(func(o *options) { o.query["task_start_time"] = taskStartTime })
}

// WithTaskUpdateTime task_update_time获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithTaskUpdateTime(taskUpdateTime int64) Option {
	return optionFunc(func(o *options) { o.query["task_update_time"] = taskUpdateTime })
}

// WithTriggerType trigger_type获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithTriggerType(triggerType int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_type"] = triggerType })
}

// WithStatus status获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTTLDelCnt ttl_del_cnt获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithTTLDelCnt(ttlDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["ttl_del_cnt"] = ttlDelCnt })
}

// WithMaxVersionDelCnt max_version_del_cnt获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithMaxVersionDelCnt(maxVersionDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["max_version_del_cnt"] = maxVersionDelCnt })
}

// WithScanCnt scan_cnt获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithScanCnt(scanCnt int64) Option {
	return optionFunc(func(o *options) { o.query["scan_cnt"] = scanCnt })
}

// WithRowKey row_key获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithRowKey(rowKey []byte) Option {
	return optionFunc(func(o *options) { o.query["row_key"] = rowKey })
}

// WithRetCode ret_code获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) WithRetCode(retCode string) Option {
	return optionFunc(func(o *options) { o.query["ret_code"] = retCode })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetByOption(opts ...Option) (result AllVirtualKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualKvTTLTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualKvTTLTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where(options.query)
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
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromTaskID(taskID int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromTableID(tableID int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTaskStartTime 通过task_start_time获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromTaskStartTime(taskStartTime int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`task_start_time` = ?", taskStartTime).Find(&results).Error

	return
}

// GetBatchFromTaskStartTime 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromTaskStartTime(taskStartTimes []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`task_start_time` IN (?)", taskStartTimes).Find(&results).Error

	return
}

// GetFromTaskUpdateTime 通过task_update_time获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromTaskUpdateTime(taskUpdateTime int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`task_update_time` = ?", taskUpdateTime).Find(&results).Error

	return
}

// GetBatchFromTaskUpdateTime 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromTaskUpdateTime(taskUpdateTimes []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`task_update_time` IN (?)", taskUpdateTimes).Find(&results).Error

	return
}

// GetFromTriggerType 通过trigger_type获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromTriggerType(triggerType int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`trigger_type` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromTriggerType(triggerTypes []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`trigger_type` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromStatus(status int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTTLDelCnt 通过ttl_del_cnt获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromTTLDelCnt(ttlDelCnt int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`ttl_del_cnt` = ?", ttlDelCnt).Find(&results).Error

	return
}

// GetBatchFromTTLDelCnt 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromTTLDelCnt(ttlDelCnts []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`ttl_del_cnt` IN (?)", ttlDelCnts).Find(&results).Error

	return
}

// GetFromMaxVersionDelCnt 通过max_version_del_cnt获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromMaxVersionDelCnt(maxVersionDelCnt int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`max_version_del_cnt` = ?", maxVersionDelCnt).Find(&results).Error

	return
}

// GetBatchFromMaxVersionDelCnt 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromMaxVersionDelCnt(maxVersionDelCnts []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`max_version_del_cnt` IN (?)", maxVersionDelCnts).Find(&results).Error

	return
}

// GetFromScanCnt 通过scan_cnt获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromScanCnt(scanCnt int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`scan_cnt` = ?", scanCnt).Find(&results).Error

	return
}

// GetBatchFromScanCnt 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromScanCnt(scanCnts []int64) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`scan_cnt` IN (?)", scanCnts).Find(&results).Error

	return
}

// GetFromRowKey 通过row_key获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromRowKey(rowKey []byte) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`row_key` = ?", rowKey).Find(&results).Error

	return
}

// GetBatchFromRowKey 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromRowKey(rowKeys [][]byte) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`row_key` IN (?)", rowKeys).Find(&results).Error

	return
}

// GetFromRetCode 通过ret_code获取内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetFromRetCode(retCode string) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`ret_code` = ?", retCode).Find(&results).Error

	return
}

// GetBatchFromRetCode 批量查找
func (obj *_AllVirtualKvTTLTaskHistoryMgr) GetBatchFromRetCode(retCodes []string) (results []*AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`ret_code` IN (?)", retCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualKvTTLTaskHistoryMgr) FetchByPrimaryKey(tenantID int64, taskID int64, tableID int64, partitionID int64) (result AllVirtualKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualKvTTLTaskHistory{}).Where("`tenant_id` = ? AND `task_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, taskID, tableID, partitionID).First(&result).Error

	return
}
