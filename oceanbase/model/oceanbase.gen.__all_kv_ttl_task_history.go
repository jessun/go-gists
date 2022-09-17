package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllKvTTLTaskHistoryMgr struct {
	*_BaseMgr
}

// AllKvTTLTaskHistoryMgr open func
func AllKvTTLTaskHistoryMgr(db *gorm.DB) *_AllKvTTLTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllKvTTLTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllKvTTLTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_kv_ttl_task_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllKvTTLTaskHistoryMgr) GetTableName() string {
	return "__all_kv_ttl_task_history"
}

// Reset 重置gorm会话
func (obj *_AllKvTTLTaskHistoryMgr) Reset() *_AllKvTTLTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllKvTTLTaskHistoryMgr) Get() (result AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllKvTTLTaskHistoryMgr) Gets() (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllKvTTLTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllKvTTLTaskHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllKvTTLTaskHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllKvTTLTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTaskID task_id获取
func (obj *_AllKvTTLTaskHistoryMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithTableID table_id获取
func (obj *_AllKvTTLTaskHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllKvTTLTaskHistoryMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithTaskStartTime task_start_time获取
func (obj *_AllKvTTLTaskHistoryMgr) WithTaskStartTime(taskStartTime int64) Option {
	return optionFunc(func(o *options) { o.query["task_start_time"] = taskStartTime })
}

// WithTaskUpdateTime task_update_time获取
func (obj *_AllKvTTLTaskHistoryMgr) WithTaskUpdateTime(taskUpdateTime int64) Option {
	return optionFunc(func(o *options) { o.query["task_update_time"] = taskUpdateTime })
}

// WithTriggerType trigger_type获取
func (obj *_AllKvTTLTaskHistoryMgr) WithTriggerType(triggerType int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_type"] = triggerType })
}

// WithStatus status获取
func (obj *_AllKvTTLTaskHistoryMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTTLDelCnt ttl_del_cnt获取
func (obj *_AllKvTTLTaskHistoryMgr) WithTTLDelCnt(ttlDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["ttl_del_cnt"] = ttlDelCnt })
}

// WithMaxVersionDelCnt max_version_del_cnt获取
func (obj *_AllKvTTLTaskHistoryMgr) WithMaxVersionDelCnt(maxVersionDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["max_version_del_cnt"] = maxVersionDelCnt })
}

// WithScanCnt scan_cnt获取
func (obj *_AllKvTTLTaskHistoryMgr) WithScanCnt(scanCnt int64) Option {
	return optionFunc(func(o *options) { o.query["scan_cnt"] = scanCnt })
}

// WithRowKey row_key获取
func (obj *_AllKvTTLTaskHistoryMgr) WithRowKey(rowKey []byte) Option {
	return optionFunc(func(o *options) { o.query["row_key"] = rowKey })
}

// WithRetCode ret_code获取
func (obj *_AllKvTTLTaskHistoryMgr) WithRetCode(retCode string) Option {
	return optionFunc(func(o *options) { o.query["ret_code"] = retCode })
}

// GetByOption 功能选项模式获取
func (obj *_AllKvTTLTaskHistoryMgr) GetByOption(opts ...Option) (result AllKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllKvTTLTaskHistoryMgr) GetByOptions(opts ...Option) (results []*AllKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllKvTTLTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllKvTTLTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where(options.query)
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
func (obj *_AllKvTTLTaskHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromTaskID(taskID int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromTableID(tableID int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromPartitionID(partitionID int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromTaskStartTime 通过task_start_time获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromTaskStartTime(taskStartTime int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`task_start_time` = ?", taskStartTime).Find(&results).Error

	return
}

// GetBatchFromTaskStartTime 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromTaskStartTime(taskStartTimes []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`task_start_time` IN (?)", taskStartTimes).Find(&results).Error

	return
}

// GetFromTaskUpdateTime 通过task_update_time获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromTaskUpdateTime(taskUpdateTime int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`task_update_time` = ?", taskUpdateTime).Find(&results).Error

	return
}

// GetBatchFromTaskUpdateTime 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromTaskUpdateTime(taskUpdateTimes []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`task_update_time` IN (?)", taskUpdateTimes).Find(&results).Error

	return
}

// GetFromTriggerType 通过trigger_type获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromTriggerType(triggerType int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`trigger_type` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromTriggerType(triggerTypes []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`trigger_type` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromStatus(status int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromStatus(statuss []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTTLDelCnt 通过ttl_del_cnt获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromTTLDelCnt(ttlDelCnt int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`ttl_del_cnt` = ?", ttlDelCnt).Find(&results).Error

	return
}

// GetBatchFromTTLDelCnt 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromTTLDelCnt(ttlDelCnts []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`ttl_del_cnt` IN (?)", ttlDelCnts).Find(&results).Error

	return
}

// GetFromMaxVersionDelCnt 通过max_version_del_cnt获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromMaxVersionDelCnt(maxVersionDelCnt int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`max_version_del_cnt` = ?", maxVersionDelCnt).Find(&results).Error

	return
}

// GetBatchFromMaxVersionDelCnt 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromMaxVersionDelCnt(maxVersionDelCnts []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`max_version_del_cnt` IN (?)", maxVersionDelCnts).Find(&results).Error

	return
}

// GetFromScanCnt 通过scan_cnt获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromScanCnt(scanCnt int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`scan_cnt` = ?", scanCnt).Find(&results).Error

	return
}

// GetBatchFromScanCnt 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromScanCnt(scanCnts []int64) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`scan_cnt` IN (?)", scanCnts).Find(&results).Error

	return
}

// GetFromRowKey 通过row_key获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromRowKey(rowKey []byte) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`row_key` = ?", rowKey).Find(&results).Error

	return
}

// GetBatchFromRowKey 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromRowKey(rowKeys [][]byte) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`row_key` IN (?)", rowKeys).Find(&results).Error

	return
}

// GetFromRetCode 通过ret_code获取内容
func (obj *_AllKvTTLTaskHistoryMgr) GetFromRetCode(retCode string) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`ret_code` = ?", retCode).Find(&results).Error

	return
}

// GetBatchFromRetCode 批量查找
func (obj *_AllKvTTLTaskHistoryMgr) GetBatchFromRetCode(retCodes []string) (results []*AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`ret_code` IN (?)", retCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllKvTTLTaskHistoryMgr) FetchByPrimaryKey(tenantID int64, taskID int64, tableID int64, partitionID int64) (result AllKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllKvTTLTaskHistory{}).Where("`tenant_id` = ? AND `task_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, taskID, tableID, partitionID).First(&result).Error

	return
}
