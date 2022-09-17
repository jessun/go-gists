package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CdbObKvTTLTaskHistoryMgr struct {
	*_BaseMgr
}

// CdbObKvTTLTaskHistoryMgr open func
func CdbObKvTTLTaskHistoryMgr(db *gorm.DB) *_CdbObKvTTLTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObKvTTLTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObKvTTLTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_KV_TTL_TASK_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObKvTTLTaskHistoryMgr) GetTableName() string {
	return "CDB_OB_KV_TTL_TASK_HISTORY"
}

// Reset 重置gorm会话
func (obj *_CdbObKvTTLTaskHistoryMgr) Reset() *_CdbObKvTTLTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObKvTTLTaskHistoryMgr) Get() (result CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObKvTTLTaskHistoryMgr) Gets() (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObKvTTLTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTableName TABLE_NAME获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithTableID TABLE_ID获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithTaskID TASK_ID获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["TASK_ID"] = taskID })
}

// WithStartTime START_TIME获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithTriggerType TRIGGER_TYPE获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithTriggerType(triggerType string) Option {
	return optionFunc(func(o *options) { o.query["TRIGGER_TYPE"] = triggerType })
}

// WithStatus STATUS获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithTTLDelCnt TTL_DEL_CNT获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithTTLDelCnt(ttlDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["TTL_DEL_CNT"] = ttlDelCnt })
}

// WithMaxVersionDelCnt MAX_VERSION_DEL_CNT获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithMaxVersionDelCnt(maxVersionDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["MAX_VERSION_DEL_CNT"] = maxVersionDelCnt })
}

// WithScanCnt SCAN_CNT获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithScanCnt(scanCnt int64) Option {
	return optionFunc(func(o *options) { o.query["SCAN_CNT"] = scanCnt })
}

// WithRetCode RET_CODE获取
func (obj *_CdbObKvTTLTaskHistoryMgr) WithRetCode(retCode string) Option {
	return optionFunc(func(o *options) { o.query["RET_CODE"] = retCode })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObKvTTLTaskHistoryMgr) GetByOption(opts ...Option) (result CdbObKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObKvTTLTaskHistoryMgr) GetByOptions(opts ...Option) (results []*CdbObKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObKvTTLTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObKvTTLTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where(options.query)
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

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromTenantID(tenantID int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableName 通过TABLE_NAME获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromTableName(tableName string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromTableName(tableNames []string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromTableID 通过TABLE_ID获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromTableID(tableID int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过PARTITION_ID获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromPartitionID(partitionID int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过TASK_ID获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromTaskID(taskID int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TASK_ID` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromTaskID(taskIDs []int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TASK_ID` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromEndTime(endTime time.Time) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTriggerType 通过TRIGGER_TYPE获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromTriggerType(triggerType string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TRIGGER_TYPE` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromTriggerType(triggerTypes []string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TRIGGER_TYPE` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromStatus(status string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTTLDelCnt 通过TTL_DEL_CNT获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromTTLDelCnt(ttlDelCnt int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TTL_DEL_CNT` = ?", ttlDelCnt).Find(&results).Error

	return
}

// GetBatchFromTTLDelCnt 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromTTLDelCnt(ttlDelCnts []int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`TTL_DEL_CNT` IN (?)", ttlDelCnts).Find(&results).Error

	return
}

// GetFromMaxVersionDelCnt 通过MAX_VERSION_DEL_CNT获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromMaxVersionDelCnt(maxVersionDelCnt int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`MAX_VERSION_DEL_CNT` = ?", maxVersionDelCnt).Find(&results).Error

	return
}

// GetBatchFromMaxVersionDelCnt 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromMaxVersionDelCnt(maxVersionDelCnts []int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`MAX_VERSION_DEL_CNT` IN (?)", maxVersionDelCnts).Find(&results).Error

	return
}

// GetFromScanCnt 通过SCAN_CNT获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromScanCnt(scanCnt int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`SCAN_CNT` = ?", scanCnt).Find(&results).Error

	return
}

// GetBatchFromScanCnt 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromScanCnt(scanCnts []int64) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`SCAN_CNT` IN (?)", scanCnts).Find(&results).Error

	return
}

// GetFromRetCode 通过RET_CODE获取内容
func (obj *_CdbObKvTTLTaskHistoryMgr) GetFromRetCode(retCode string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`RET_CODE` = ?", retCode).Find(&results).Error

	return
}

// GetBatchFromRetCode 批量查找
func (obj *_CdbObKvTTLTaskHistoryMgr) GetBatchFromRetCode(retCodes []string) (results []*CdbObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObKvTTLTaskHistory{}).Where("`RET_CODE` IN (?)", retCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
