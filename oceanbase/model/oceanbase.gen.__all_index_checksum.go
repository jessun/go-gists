package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllIndexChecksumMgr struct {
	*_BaseMgr
}

// AllIndexChecksumMgr open func
func AllIndexChecksumMgr(db *gorm.DB) *_AllIndexChecksumMgr {
	if db == nil {
		panic(fmt.Errorf("AllIndexChecksumMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllIndexChecksumMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_index_checksum"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllIndexChecksumMgr) GetTableName() string {
	return "__all_index_checksum"
}

// Reset 重置gorm会话
func (obj *_AllIndexChecksumMgr) Reset() *_AllIndexChecksumMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllIndexChecksumMgr) Get() (result AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllIndexChecksumMgr) Gets() (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllIndexChecksumMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllIndexChecksumMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllIndexChecksumMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithExecutionID execution_id获取
func (obj *_AllIndexChecksumMgr) WithExecutionID(executionID int64) Option {
	return optionFunc(func(o *options) { o.query["execution_id"] = executionID })
}

// WithTenantID tenant_id获取
func (obj *_AllIndexChecksumMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllIndexChecksumMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllIndexChecksumMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithColumnID column_id获取
func (obj *_AllIndexChecksumMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithTaskID task_id获取
func (obj *_AllIndexChecksumMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithChecksum checksum获取
func (obj *_AllIndexChecksumMgr) WithChecksum(checksum int64) Option {
	return optionFunc(func(o *options) { o.query["checksum"] = checksum })
}

// WithChecksumMethod checksum_method获取
func (obj *_AllIndexChecksumMgr) WithChecksumMethod(checksumMethod int64) Option {
	return optionFunc(func(o *options) { o.query["checksum_method"] = checksumMethod })
}

// GetByOption 功能选项模式获取
func (obj *_AllIndexChecksumMgr) GetByOption(opts ...Option) (result AllIndexChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllIndexChecksumMgr) GetByOptions(opts ...Option) (results []*AllIndexChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllIndexChecksumMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllIndexChecksum, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where(options.query)
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
func (obj *_AllIndexChecksumMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllIndexChecksumMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromExecutionID 通过execution_id获取内容
func (obj *_AllIndexChecksumMgr) GetFromExecutionID(executionID int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`execution_id` = ?", executionID).Find(&results).Error

	return
}

// GetBatchFromExecutionID 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromExecutionID(executionIDs []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`execution_id` IN (?)", executionIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllIndexChecksumMgr) GetFromTenantID(tenantID int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllIndexChecksumMgr) GetFromTableID(tableID int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllIndexChecksumMgr) GetFromPartitionID(partitionID int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllIndexChecksumMgr) GetFromColumnID(columnID int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_AllIndexChecksumMgr) GetFromTaskID(taskID int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromTaskID(taskIDs []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromChecksum 通过checksum获取内容
func (obj *_AllIndexChecksumMgr) GetFromChecksum(checksum int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`checksum` = ?", checksum).Find(&results).Error

	return
}

// GetBatchFromChecksum 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromChecksum(checksums []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`checksum` IN (?)", checksums).Find(&results).Error

	return
}

// GetFromChecksumMethod 通过checksum_method获取内容
func (obj *_AllIndexChecksumMgr) GetFromChecksumMethod(checksumMethod int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`checksum_method` = ?", checksumMethod).Find(&results).Error

	return
}

// GetBatchFromChecksumMethod 批量查找
func (obj *_AllIndexChecksumMgr) GetBatchFromChecksumMethod(checksumMethods []int64) (results []*AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`checksum_method` IN (?)", checksumMethods).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllIndexChecksumMgr) FetchByPrimaryKey(executionID int64, tenantID int64, tableID int64, partitionID int64, columnID int64, taskID int64) (result AllIndexChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexChecksum{}).Where("`execution_id` = ? AND `tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `column_id` = ? AND `task_id` = ?", executionID, tenantID, tableID, partitionID, columnID, taskID).First(&result).Error

	return
}
