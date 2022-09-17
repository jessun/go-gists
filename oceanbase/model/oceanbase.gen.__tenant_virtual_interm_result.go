package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualIntermResultMgr struct {
	*_BaseMgr
}

// TenantVirtualIntermResultMgr open func
func TenantVirtualIntermResultMgr(db *gorm.DB) *_TenantVirtualIntermResultMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualIntermResultMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualIntermResultMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_interm_result"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualIntermResultMgr) GetTableName() string {
	return "__tenant_virtual_interm_result"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualIntermResultMgr) Reset() *_TenantVirtualIntermResultMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualIntermResultMgr) Get() (result TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualIntermResultMgr) Gets() (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualIntermResultMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithJobID job_id获取
func (obj *_TenantVirtualIntermResultMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTaskID task_id获取
func (obj *_TenantVirtualIntermResultMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithSliceID slice_id获取
func (obj *_TenantVirtualIntermResultMgr) WithSliceID(sliceID int64) Option {
	return optionFunc(func(o *options) { o.query["slice_id"] = sliceID })
}

// WithExecutionID execution_id获取
func (obj *_TenantVirtualIntermResultMgr) WithExecutionID(executionID int64) Option {
	return optionFunc(func(o *options) { o.query["execution_id"] = executionID })
}

// WithSvrIP svr_ip获取
func (obj *_TenantVirtualIntermResultMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_TenantVirtualIntermResultMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithExpireTime expire_time获取
func (obj *_TenantVirtualIntermResultMgr) WithExpireTime(expireTime int64) Option {
	return optionFunc(func(o *options) { o.query["expire_time"] = expireTime })
}

// WithRowCount row_count获取
func (obj *_TenantVirtualIntermResultMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithScannerCount scanner_count获取
func (obj *_TenantVirtualIntermResultMgr) WithScannerCount(scannerCount int64) Option {
	return optionFunc(func(o *options) { o.query["scanner_count"] = scannerCount })
}

// WithUsedMemorySize used_memory_size获取
func (obj *_TenantVirtualIntermResultMgr) WithUsedMemorySize(usedMemorySize int64) Option {
	return optionFunc(func(o *options) { o.query["used_memory_size"] = usedMemorySize })
}

// WithUsedDiskSize used_disk_size获取
func (obj *_TenantVirtualIntermResultMgr) WithUsedDiskSize(usedDiskSize int64) Option {
	return optionFunc(func(o *options) { o.query["used_disk_size"] = usedDiskSize })
}

// WithPartitionIP partition_ip获取
func (obj *_TenantVirtualIntermResultMgr) WithPartitionIP(partitionIP string) Option {
	return optionFunc(func(o *options) { o.query["partition_ip"] = partitionIP })
}

// WithPartitionPort partition_port获取
func (obj *_TenantVirtualIntermResultMgr) WithPartitionPort(partitionPort int64) Option {
	return optionFunc(func(o *options) { o.query["partition_port"] = partitionPort })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualIntermResultMgr) GetByOption(opts ...Option) (result TenantVirtualIntermResult, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualIntermResultMgr) GetByOptions(opts ...Option) (results []*TenantVirtualIntermResult, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualIntermResultMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualIntermResult, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where(options.query)
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

// GetFromJobID 通过job_id获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromJobID(jobID int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromJobID(jobIDs []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromTaskID(taskID int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromTaskID(taskIDs []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromSliceID 通过slice_id获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromSliceID(sliceID int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`slice_id` = ?", sliceID).Find(&results).Error

	return
}

// GetBatchFromSliceID 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromSliceID(sliceIDs []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`slice_id` IN (?)", sliceIDs).Find(&results).Error

	return
}

// GetFromExecutionID 通过execution_id获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromExecutionID(executionID int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`execution_id` = ?", executionID).Find(&results).Error

	return
}

// GetBatchFromExecutionID 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromExecutionID(executionIDs []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`execution_id` IN (?)", executionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromSvrIP(svrIP string) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromSvrIP(svrIPs []string) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromSvrPort(svrPort int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromExpireTime 通过expire_time获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromExpireTime(expireTime int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`expire_time` = ?", expireTime).Find(&results).Error

	return
}

// GetBatchFromExpireTime 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromExpireTime(expireTimes []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`expire_time` IN (?)", expireTimes).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromRowCount(rowCount int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromRowCount(rowCounts []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromScannerCount 通过scanner_count获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromScannerCount(scannerCount int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`scanner_count` = ?", scannerCount).Find(&results).Error

	return
}

// GetBatchFromScannerCount 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromScannerCount(scannerCounts []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`scanner_count` IN (?)", scannerCounts).Find(&results).Error

	return
}

// GetFromUsedMemorySize 通过used_memory_size获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromUsedMemorySize(usedMemorySize int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`used_memory_size` = ?", usedMemorySize).Find(&results).Error

	return
}

// GetBatchFromUsedMemorySize 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromUsedMemorySize(usedMemorySizes []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`used_memory_size` IN (?)", usedMemorySizes).Find(&results).Error

	return
}

// GetFromUsedDiskSize 通过used_disk_size获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromUsedDiskSize(usedDiskSize int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`used_disk_size` = ?", usedDiskSize).Find(&results).Error

	return
}

// GetBatchFromUsedDiskSize 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromUsedDiskSize(usedDiskSizes []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`used_disk_size` IN (?)", usedDiskSizes).Find(&results).Error

	return
}

// GetFromPartitionIP 通过partition_ip获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromPartitionIP(partitionIP string) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`partition_ip` = ?", partitionIP).Find(&results).Error

	return
}

// GetBatchFromPartitionIP 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromPartitionIP(partitionIPs []string) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`partition_ip` IN (?)", partitionIPs).Find(&results).Error

	return
}

// GetFromPartitionPort 通过partition_port获取内容
func (obj *_TenantVirtualIntermResultMgr) GetFromPartitionPort(partitionPort int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`partition_port` = ?", partitionPort).Find(&results).Error

	return
}

// GetBatchFromPartitionPort 批量查找
func (obj *_TenantVirtualIntermResultMgr) GetBatchFromPartitionPort(partitionPorts []int64) (results []*TenantVirtualIntermResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualIntermResult{}).Where("`partition_port` IN (?)", partitionPorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
