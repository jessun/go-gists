package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _DbaObKvTTLTaskHistoryMgr struct {
	*_BaseMgr
}

// DbaObKvTTLTaskHistoryMgr open func
func DbaObKvTTLTaskHistoryMgr(db *gorm.DB) *_DbaObKvTTLTaskHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("DbaObKvTTLTaskHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DbaObKvTTLTaskHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("DBA_OB_KV_TTL_TASK_HISTORY"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DbaObKvTTLTaskHistoryMgr) GetTableName() string {
	return "DBA_OB_KV_TTL_TASK_HISTORY"
}

// Reset 重置gorm会话
func (obj *_DbaObKvTTLTaskHistoryMgr) Reset() *_DbaObKvTTLTaskHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DbaObKvTTLTaskHistoryMgr) Get() (result DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DbaObKvTTLTaskHistoryMgr) Gets() (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DbaObKvTTLTaskHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableName TABLE_NAME获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithTableID TABLE_ID获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithTaskID TASK_ID获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["TASK_ID"] = taskID })
}

// WithStartTime START_TIME获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithTriggerType TRIGGER_TYPE获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithTriggerType(triggerType string) Option {
	return optionFunc(func(o *options) { o.query["TRIGGER_TYPE"] = triggerType })
}

// WithStatus STATUS获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithTTLDelCnt TTL_DEL_CNT获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithTTLDelCnt(ttlDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["TTL_DEL_CNT"] = ttlDelCnt })
}

// WithMaxVersionDelCnt MAX_VERSION_DEL_CNT获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithMaxVersionDelCnt(maxVersionDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["MAX_VERSION_DEL_CNT"] = maxVersionDelCnt })
}

// WithScanCnt SCAN_CNT获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithScanCnt(scanCnt int64) Option {
	return optionFunc(func(o *options) { o.query["SCAN_CNT"] = scanCnt })
}

// WithRetCode RET_CODE获取
func (obj *_DbaObKvTTLTaskHistoryMgr) WithRetCode(retCode string) Option {
	return optionFunc(func(o *options) { o.query["RET_CODE"] = retCode })
}

// GetByOption 功能选项模式获取
func (obj *_DbaObKvTTLTaskHistoryMgr) GetByOption(opts ...Option) (result DbaObKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DbaObKvTTLTaskHistoryMgr) GetByOptions(opts ...Option) (results []*DbaObKvTTLTaskHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_DbaObKvTTLTaskHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]DbaObKvTTLTaskHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where(options.query)
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

// GetFromTableName 通过TABLE_NAME获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromTableName(tableName string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromTableName(tableNames []string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromTableID 通过TABLE_ID获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromTableID(tableID int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过PARTITION_ID获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromPartitionID(partitionID int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过TASK_ID获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromTaskID(taskID int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TASK_ID` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromTaskID(taskIDs []int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TASK_ID` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromStartTime(startTime time.Time) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromEndTime(endTime time.Time) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTriggerType 通过TRIGGER_TYPE获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromTriggerType(triggerType string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TRIGGER_TYPE` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromTriggerType(triggerTypes []string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TRIGGER_TYPE` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromStatus(status string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromStatus(statuss []string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTTLDelCnt 通过TTL_DEL_CNT获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromTTLDelCnt(ttlDelCnt int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TTL_DEL_CNT` = ?", ttlDelCnt).Find(&results).Error

	return
}

// GetBatchFromTTLDelCnt 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromTTLDelCnt(ttlDelCnts []int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`TTL_DEL_CNT` IN (?)", ttlDelCnts).Find(&results).Error

	return
}

// GetFromMaxVersionDelCnt 通过MAX_VERSION_DEL_CNT获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromMaxVersionDelCnt(maxVersionDelCnt int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`MAX_VERSION_DEL_CNT` = ?", maxVersionDelCnt).Find(&results).Error

	return
}

// GetBatchFromMaxVersionDelCnt 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromMaxVersionDelCnt(maxVersionDelCnts []int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`MAX_VERSION_DEL_CNT` IN (?)", maxVersionDelCnts).Find(&results).Error

	return
}

// GetFromScanCnt 通过SCAN_CNT获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromScanCnt(scanCnt int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`SCAN_CNT` = ?", scanCnt).Find(&results).Error

	return
}

// GetBatchFromScanCnt 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromScanCnt(scanCnts []int64) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`SCAN_CNT` IN (?)", scanCnts).Find(&results).Error

	return
}

// GetFromRetCode 通过RET_CODE获取内容
func (obj *_DbaObKvTTLTaskHistoryMgr) GetFromRetCode(retCode string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`RET_CODE` = ?", retCode).Find(&results).Error

	return
}

// GetBatchFromRetCode 批量查找
func (obj *_DbaObKvTTLTaskHistoryMgr) GetBatchFromRetCode(retCodes []string) (results []*DbaObKvTTLTaskHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTaskHistory{}).Where("`RET_CODE` IN (?)", retCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
