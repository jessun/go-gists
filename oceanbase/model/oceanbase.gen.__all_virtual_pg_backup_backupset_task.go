package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualPgBackupBackupsetTaskMgr struct {
	*_BaseMgr
}

// AllVirtualPgBackupBackupsetTaskMgr open func
func AllVirtualPgBackupBackupsetTaskMgr(db *gorm.DB) *_AllVirtualPgBackupBackupsetTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPgBackupBackupsetTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPgBackupBackupsetTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_pg_backup_backupset_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetTableName() string {
	return "__all_virtual_pg_backup_backupset_task"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) Reset() *_AllVirtualPgBackupBackupsetTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) Get() (result AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) Gets() (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithJobID job_id获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupSetID backup_set_id获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithBackupSetID(backupSetID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_set_id"] = backupSetID })
}

// WithCopyID copy_id获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithStatus status获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTraceID trace_id获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTotalPartitionCount total_partition_count获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithTotalPartitionCount(totalPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_partition_count"] = totalPartitionCount })
}

// WithFinishPartitionCount finish_partition_count获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithFinishPartitionCount(finishPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_partition_count"] = finishPartitionCount })
}

// WithTotalMacroBlockCount total_macro_block_count获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithTotalMacroBlockCount(totalMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_macro_block_count"] = totalMacroBlockCount })
}

// WithFinishMacroBlockCount finish_macro_block_count获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithFinishMacroBlockCount(finishMacroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["finish_macro_block_count"] = finishMacroBlockCount })
}

// WithResult result获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithComment comment获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetByOption(opts ...Option) (result AllVirtualPgBackupBackupsetTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetByOptions(opts ...Option) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPgBackupBackupsetTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where(options.query)
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
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromJobID(jobID int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupSetID 通过backup_set_id获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromBackupSetID(backupSetID int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`backup_set_id` = ?", backupSetID).Find(&results).Error

	return
}

// GetBatchFromBackupSetID 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromBackupSetID(backupSetIDs []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`backup_set_id` IN (?)", backupSetIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromCopyID(copyID int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromTableID(tableID int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromStatus(status string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTraceID 通过trace_id获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromTraceID(traceID string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`trace_id` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTotalPartitionCount 通过total_partition_count获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromTotalPartitionCount(totalPartitionCount int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`total_partition_count` = ?", totalPartitionCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartitionCount 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromTotalPartitionCount(totalPartitionCounts []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`total_partition_count` IN (?)", totalPartitionCounts).Find(&results).Error

	return
}

// GetFromFinishPartitionCount 通过finish_partition_count获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromFinishPartitionCount(finishPartitionCount int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`finish_partition_count` = ?", finishPartitionCount).Find(&results).Error

	return
}

// GetBatchFromFinishPartitionCount 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromFinishPartitionCount(finishPartitionCounts []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`finish_partition_count` IN (?)", finishPartitionCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlockCount 通过total_macro_block_count获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromTotalMacroBlockCount(totalMacroBlockCount int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`total_macro_block_count` = ?", totalMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlockCount 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromTotalMacroBlockCount(totalMacroBlockCounts []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`total_macro_block_count` IN (?)", totalMacroBlockCounts).Find(&results).Error

	return
}

// GetFromFinishMacroBlockCount 通过finish_macro_block_count获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromFinishMacroBlockCount(finishMacroBlockCount int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`finish_macro_block_count` = ?", finishMacroBlockCount).Find(&results).Error

	return
}

// GetBatchFromFinishMacroBlockCount 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromFinishMacroBlockCount(finishMacroBlockCounts []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`finish_macro_block_count` IN (?)", finishMacroBlockCounts).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromResult(result int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromResult(results []int64) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetFromComment(comment string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) GetBatchFromComment(comments []string) (results []*AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPgBackupBackupsetTaskMgr) FetchByPrimaryKey(tenantID int64, jobID int64, incarnation int64, backupSetID int64, copyID int64, tableID int64, partitionID int64) (result AllVirtualPgBackupBackupsetTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupBackupsetTask{}).Where("`tenant_id` = ? AND `job_id` = ? AND `incarnation` = ? AND `backup_set_id` = ? AND `copy_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, jobID, incarnation, backupSetID, copyID, tableID, partitionID).First(&result).Error

	return
}
