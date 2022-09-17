package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _DbaObKvTTLTasksMgr struct {
	*_BaseMgr
}

// DbaObKvTTLTasksMgr open func
func DbaObKvTTLTasksMgr(db *gorm.DB) *_DbaObKvTTLTasksMgr {
	if db == nil {
		panic(fmt.Errorf("DbaObKvTTLTasksMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DbaObKvTTLTasksMgr{_BaseMgr: &_BaseMgr{DB: db.Table("DBA_OB_KV_TTL_TASKS"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DbaObKvTTLTasksMgr) GetTableName() string {
	return "DBA_OB_KV_TTL_TASKS"
}

// Reset 重置gorm会话
func (obj *_DbaObKvTTLTasksMgr) Reset() *_DbaObKvTTLTasksMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DbaObKvTTLTasksMgr) Get() (result DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DbaObKvTTLTasksMgr) Gets() (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DbaObKvTTLTasksMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableName TABLE_NAME获取
func (obj *_DbaObKvTTLTasksMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithTableID TABLE_ID获取
func (obj *_DbaObKvTTLTasksMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取
func (obj *_DbaObKvTTLTasksMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithTaskID TASK_ID获取
func (obj *_DbaObKvTTLTasksMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["TASK_ID"] = taskID })
}

// WithStartTime START_TIME获取
func (obj *_DbaObKvTTLTasksMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取
func (obj *_DbaObKvTTLTasksMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithTriggerType TRIGGER_TYPE获取
func (obj *_DbaObKvTTLTasksMgr) WithTriggerType(triggerType string) Option {
	return optionFunc(func(o *options) { o.query["TRIGGER_TYPE"] = triggerType })
}

// WithStatus STATUS获取
func (obj *_DbaObKvTTLTasksMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithTTLDelCnt TTL_DEL_CNT获取
func (obj *_DbaObKvTTLTasksMgr) WithTTLDelCnt(ttlDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["TTL_DEL_CNT"] = ttlDelCnt })
}

// WithMaxVersionDelCnt MAX_VERSION_DEL_CNT获取
func (obj *_DbaObKvTTLTasksMgr) WithMaxVersionDelCnt(maxVersionDelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["MAX_VERSION_DEL_CNT"] = maxVersionDelCnt })
}

// WithScanCnt SCAN_CNT获取
func (obj *_DbaObKvTTLTasksMgr) WithScanCnt(scanCnt int64) Option {
	return optionFunc(func(o *options) { o.query["SCAN_CNT"] = scanCnt })
}

// WithRetCode RET_CODE获取
func (obj *_DbaObKvTTLTasksMgr) WithRetCode(retCode string) Option {
	return optionFunc(func(o *options) { o.query["RET_CODE"] = retCode })
}

// GetByOption 功能选项模式获取
func (obj *_DbaObKvTTLTasksMgr) GetByOption(opts ...Option) (result DbaObKvTTLTasks, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DbaObKvTTLTasksMgr) GetByOptions(opts ...Option) (results []*DbaObKvTTLTasks, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_DbaObKvTTLTasksMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]DbaObKvTTLTasks, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where(options.query)
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
func (obj *_DbaObKvTTLTasksMgr) GetFromTableName(tableName string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromTableName(tableNames []string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromTableID 通过TABLE_ID获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromTableID(tableID int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromTableID(tableIDs []int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过PARTITION_ID获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromPartitionID(partitionID int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过TASK_ID获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromTaskID(taskID int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TASK_ID` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromTaskID(taskIDs []int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TASK_ID` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过START_TIME获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromStartTime(startTime time.Time) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`START_TIME` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过END_TIME获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromEndTime(endTime time.Time) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`END_TIME` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTriggerType 通过TRIGGER_TYPE获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromTriggerType(triggerType string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TRIGGER_TYPE` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromTriggerType(triggerTypes []string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TRIGGER_TYPE` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromStatus(status string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromStatus(statuss []string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTTLDelCnt 通过TTL_DEL_CNT获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromTTLDelCnt(ttlDelCnt int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TTL_DEL_CNT` = ?", ttlDelCnt).Find(&results).Error

	return
}

// GetBatchFromTTLDelCnt 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromTTLDelCnt(ttlDelCnts []int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`TTL_DEL_CNT` IN (?)", ttlDelCnts).Find(&results).Error

	return
}

// GetFromMaxVersionDelCnt 通过MAX_VERSION_DEL_CNT获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromMaxVersionDelCnt(maxVersionDelCnt int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`MAX_VERSION_DEL_CNT` = ?", maxVersionDelCnt).Find(&results).Error

	return
}

// GetBatchFromMaxVersionDelCnt 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromMaxVersionDelCnt(maxVersionDelCnts []int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`MAX_VERSION_DEL_CNT` IN (?)", maxVersionDelCnts).Find(&results).Error

	return
}

// GetFromScanCnt 通过SCAN_CNT获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromScanCnt(scanCnt int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`SCAN_CNT` = ?", scanCnt).Find(&results).Error

	return
}

// GetBatchFromScanCnt 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromScanCnt(scanCnts []int64) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`SCAN_CNT` IN (?)", scanCnts).Find(&results).Error

	return
}

// GetFromRetCode 通过RET_CODE获取内容
func (obj *_DbaObKvTTLTasksMgr) GetFromRetCode(retCode string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`RET_CODE` = ?", retCode).Find(&results).Error

	return
}

// GetBatchFromRetCode 批量查找
func (obj *_DbaObKvTTLTasksMgr) GetBatchFromRetCode(retCodes []string) (results []*DbaObKvTTLTasks, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DbaObKvTTLTasks{}).Where("`RET_CODE` IN (?)", retCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
