package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualRebalanceTaskStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceTaskStatMgr open func
func AllVirtualRebalanceTaskStatMgr(db *gorm.DB) *_AllVirtualRebalanceTaskStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceTaskStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceTaskStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_task_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceTaskStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_task_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceTaskStatMgr) Reset() *_AllVirtualRebalanceTaskStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceTaskStatMgr) Get() (result AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceTaskStatMgr) Gets() (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceTaskStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithTenantID(tenantID uint64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithTableID(tableID uint64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionCount partition_count获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithPartitionCount(partitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_count"] = partitionCount })
}

// WithSource source获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithDataSource data_source获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithDataSource(dataSource string) Option {
	return optionFunc(func(o *options) { o.query["data_source"] = dataSource })
}

// WithDestination destination获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithDestination(destination string) Option {
	return optionFunc(func(o *options) { o.query["destination"] = destination })
}

// WithOffline offline获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithOffline(offline string) Option {
	return optionFunc(func(o *options) { o.query["offline"] = offline })
}

// WithIsReplicate is_replicate获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithIsReplicate(isReplicate string) Option {
	return optionFunc(func(o *options) { o.query["is_replicate"] = isReplicate })
}

// WithTaskType task_type获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithTaskType(taskType string) Option {
	return optionFunc(func(o *options) { o.query["task_type"] = taskType })
}

// WithIsScheduled is_scheduled获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithIsScheduled(isScheduled string) Option {
	return optionFunc(func(o *options) { o.query["is_scheduled"] = isScheduled })
}

// WithIsManual is_manual获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithIsManual(isManual string) Option {
	return optionFunc(func(o *options) { o.query["is_manual"] = isManual })
}

// WithWaitingTime waiting_time获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithWaitingTime(waitingTime int64) Option {
	return optionFunc(func(o *options) { o.query["waiting_time"] = waitingTime })
}

// WithExecutingTime executing_time获取
func (obj *_AllVirtualRebalanceTaskStatMgr) WithExecutingTime(executingTime int64) Option {
	return optionFunc(func(o *options) { o.query["executing_time"] = executingTime })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceTaskStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceTaskStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceTaskStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceTaskStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceTaskStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceTaskStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where(options.query)
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
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromTenantID(tenantID uint64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromTenantID(tenantIDs []uint64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromTableID(tableID uint64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromTableID(tableIDs []uint64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromPartitionCount 通过partition_count获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromPartitionCount(partitionCount int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`partition_count` = ?", partitionCount).Find(&results).Error

	return
}

// GetBatchFromPartitionCount 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromPartitionCount(partitionCounts []int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`partition_count` IN (?)", partitionCounts).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromSource(source string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromSource(sources []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromDataSource 通过data_source获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromDataSource(dataSource string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`data_source` = ?", dataSource).Find(&results).Error

	return
}

// GetBatchFromDataSource 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromDataSource(dataSources []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`data_source` IN (?)", dataSources).Find(&results).Error

	return
}

// GetFromDestination 通过destination获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromDestination(destination string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`destination` = ?", destination).Find(&results).Error

	return
}

// GetBatchFromDestination 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromDestination(destinations []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`destination` IN (?)", destinations).Find(&results).Error

	return
}

// GetFromOffline 通过offline获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromOffline(offline string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`offline` = ?", offline).Find(&results).Error

	return
}

// GetBatchFromOffline 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromOffline(offlines []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`offline` IN (?)", offlines).Find(&results).Error

	return
}

// GetFromIsReplicate 通过is_replicate获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromIsReplicate(isReplicate string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`is_replicate` = ?", isReplicate).Find(&results).Error

	return
}

// GetBatchFromIsReplicate 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromIsReplicate(isReplicates []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`is_replicate` IN (?)", isReplicates).Find(&results).Error

	return
}

// GetFromTaskType 通过task_type获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromTaskType(taskType string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`task_type` = ?", taskType).Find(&results).Error

	return
}

// GetBatchFromTaskType 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromTaskType(taskTypes []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`task_type` IN (?)", taskTypes).Find(&results).Error

	return
}

// GetFromIsScheduled 通过is_scheduled获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromIsScheduled(isScheduled string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`is_scheduled` = ?", isScheduled).Find(&results).Error

	return
}

// GetBatchFromIsScheduled 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromIsScheduled(isScheduleds []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`is_scheduled` IN (?)", isScheduleds).Find(&results).Error

	return
}

// GetFromIsManual 通过is_manual获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromIsManual(isManual string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`is_manual` = ?", isManual).Find(&results).Error

	return
}

// GetBatchFromIsManual 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromIsManual(isManuals []string) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`is_manual` IN (?)", isManuals).Find(&results).Error

	return
}

// GetFromWaitingTime 通过waiting_time获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromWaitingTime(waitingTime int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`waiting_time` = ?", waitingTime).Find(&results).Error

	return
}

// GetBatchFromWaitingTime 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromWaitingTime(waitingTimes []int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`waiting_time` IN (?)", waitingTimes).Find(&results).Error

	return
}

// GetFromExecutingTime 通过executing_time获取内容
func (obj *_AllVirtualRebalanceTaskStatMgr) GetFromExecutingTime(executingTime int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`executing_time` = ?", executingTime).Find(&results).Error

	return
}

// GetBatchFromExecutingTime 批量查找
func (obj *_AllVirtualRebalanceTaskStatMgr) GetBatchFromExecutingTime(executingTimes []int64) (results []*AllVirtualRebalanceTaskStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceTaskStat{}).Where("`executing_time` IN (?)", executingTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
