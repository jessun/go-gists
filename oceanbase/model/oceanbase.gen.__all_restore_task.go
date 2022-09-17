package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllRestoreTaskMgr struct {
	*_BaseMgr
}

// AllRestoreTaskMgr open func
func AllRestoreTaskMgr(db *gorm.DB) *_AllRestoreTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllRestoreTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRestoreTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_restore_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRestoreTaskMgr) GetTableName() string {
	return "__all_restore_task"
}

// Reset 重置gorm会话
func (obj *_AllRestoreTaskMgr) Reset() *_AllRestoreTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRestoreTaskMgr) Get() (result AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRestoreTaskMgr) Gets() (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRestoreTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRestoreTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllRestoreTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllRestoreTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllRestoreTaskMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllRestoreTaskMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithBackupTableID backup_table_id获取
func (obj *_AllRestoreTaskMgr) WithBackupTableID(backupTableID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_table_id"] = backupTableID })
}

// WithIndexMap index_map获取
func (obj *_AllRestoreTaskMgr) WithIndexMap(indexMap string) Option {
	return optionFunc(func(o *options) { o.query["index_map"] = indexMap })
}

// WithStartTime start_time获取
func (obj *_AllRestoreTaskMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithStatus status获取
func (obj *_AllRestoreTaskMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithJobID job_id获取
func (obj *_AllRestoreTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// GetByOption 功能选项模式获取
func (obj *_AllRestoreTaskMgr) GetByOption(opts ...Option) (result AllRestoreTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRestoreTaskMgr) GetByOptions(opts ...Option) (results []*AllRestoreTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRestoreTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRestoreTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where(options.query)
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
func (obj *_AllRestoreTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllRestoreTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllRestoreTaskMgr) GetFromTenantID(tenantID int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllRestoreTaskMgr) GetFromTableID(tableID int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllRestoreTaskMgr) GetFromPartitionID(partitionID int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromBackupTableID 通过backup_table_id获取内容
func (obj *_AllRestoreTaskMgr) GetFromBackupTableID(backupTableID int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`backup_table_id` = ?", backupTableID).Find(&results).Error

	return
}

// GetBatchFromBackupTableID 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromBackupTableID(backupTableIDs []int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`backup_table_id` IN (?)", backupTableIDs).Find(&results).Error

	return
}

// GetFromIndexMap 通过index_map获取内容
func (obj *_AllRestoreTaskMgr) GetFromIndexMap(indexMap string) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`index_map` = ?", indexMap).Find(&results).Error

	return
}

// GetBatchFromIndexMap 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromIndexMap(indexMaps []string) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`index_map` IN (?)", indexMaps).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllRestoreTaskMgr) GetFromStartTime(startTime int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromStartTime(startTimes []int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllRestoreTaskMgr) GetFromStatus(status int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromStatus(statuss []int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllRestoreTaskMgr) GetFromJobID(jobID int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllRestoreTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRestoreTaskMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64) (result AllRestoreTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRestoreTask{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, tableID, partitionID).First(&result).Error

	return
}
