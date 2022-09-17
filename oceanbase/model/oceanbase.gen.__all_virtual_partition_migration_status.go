package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualPartitionMigrationStatusMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionMigrationStatusMgr open func
func AllVirtualPartitionMigrationStatusMgr(db *gorm.DB) *_AllVirtualPartitionMigrationStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionMigrationStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionMigrationStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_migration_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetTableName() string {
	return "__all_virtual_partition_migration_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionMigrationStatusMgr) Reset() *_AllVirtualPartitionMigrationStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) Get() (result AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionMigrationStatusMgr) Gets() (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionMigrationStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTaskID task_id获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithTaskID(taskID string) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithMigrateType migrate_type获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithMigrateType(migrateType string) Option {
	return optionFunc(func(o *options) { o.query["migrate_type"] = migrateType })
}

// WithParentIP parent_ip获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithParentIP(parentIP string) Option {
	return optionFunc(func(o *options) { o.query["parent_ip"] = parentIP })
}

// WithParentPort parent_port获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithParentPort(parentPort int64) Option {
	return optionFunc(func(o *options) { o.query["parent_port"] = parentPort })
}

// WithSrcIP src_ip获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithSrcIP(srcIP string) Option {
	return optionFunc(func(o *options) { o.query["src_ip"] = srcIP })
}

// WithSrcPort src_port获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithSrcPort(srcPort int64) Option {
	return optionFunc(func(o *options) { o.query["src_port"] = srcPort })
}

// WithDestIP dest_ip获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithDestIP(destIP string) Option {
	return optionFunc(func(o *options) { o.query["dest_ip"] = destIP })
}

// WithDestPort dest_port获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithDestPort(destPort int64) Option {
	return optionFunc(func(o *options) { o.query["dest_port"] = destPort })
}

// WithResult result获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithStartTime start_time获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithFinishTime finish_time获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithFinishTime(finishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["finish_time"] = finishTime })
}

// WithAction action获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithAction(action string) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

// WithReplicaState replica_state获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithReplicaState(replicaState string) Option {
	return optionFunc(func(o *options) { o.query["replica_state"] = replicaState })
}

// WithRebuildCount rebuild_count获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithRebuildCount(rebuildCount int64) Option {
	return optionFunc(func(o *options) { o.query["rebuild_count"] = rebuildCount })
}

// WithTotalMacroBlock total_macro_block获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithTotalMacroBlock(totalMacroBlock int64) Option {
	return optionFunc(func(o *options) { o.query["total_macro_block"] = totalMacroBlock })
}

// WithReadyMacroBlock ready_macro_block获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithReadyMacroBlock(readyMacroBlock int64) Option {
	return optionFunc(func(o *options) { o.query["ready_macro_block"] = readyMacroBlock })
}

// WithMajorCount major_count获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithMajorCount(majorCount int64) Option {
	return optionFunc(func(o *options) { o.query["major_count"] = majorCount })
}

// WithMiniMinorCount mini_minor_count获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithMiniMinorCount(miniMinorCount int64) Option {
	return optionFunc(func(o *options) { o.query["mini_minor_count"] = miniMinorCount })
}

// WithNormalMinorCount normal_minor_count获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithNormalMinorCount(normalMinorCount int64) Option {
	return optionFunc(func(o *options) { o.query["normal_minor_count"] = normalMinorCount })
}

// WithBufMinorCount buf_minor_count获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithBufMinorCount(bufMinorCount int64) Option {
	return optionFunc(func(o *options) { o.query["buf_minor_count"] = bufMinorCount })
}

// WithReuseCount reuse_count获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithReuseCount(reuseCount int64) Option {
	return optionFunc(func(o *options) { o.query["reuse_count"] = reuseCount })
}

// WithComment comment获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetByOption(opts ...Option) (result AllVirtualPartitionMigrationStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionMigrationStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionMigrationStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionMigrationStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where(options.query)
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

// GetFromTaskID 通过task_id获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromTaskID(taskID string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromTaskID(taskIDs []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromMigrateType 通过migrate_type获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromMigrateType(migrateType string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`migrate_type` = ?", migrateType).Find(&results).Error

	return
}

// GetBatchFromMigrateType 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromMigrateType(migrateTypes []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`migrate_type` IN (?)", migrateTypes).Find(&results).Error

	return
}

// GetFromParentIP 通过parent_ip获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromParentIP(parentIP string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`parent_ip` = ?", parentIP).Find(&results).Error

	return
}

// GetBatchFromParentIP 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromParentIP(parentIPs []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`parent_ip` IN (?)", parentIPs).Find(&results).Error

	return
}

// GetFromParentPort 通过parent_port获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromParentPort(parentPort int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`parent_port` = ?", parentPort).Find(&results).Error

	return
}

// GetBatchFromParentPort 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromParentPort(parentPorts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`parent_port` IN (?)", parentPorts).Find(&results).Error

	return
}

// GetFromSrcIP 通过src_ip获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromSrcIP(srcIP string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`src_ip` = ?", srcIP).Find(&results).Error

	return
}

// GetBatchFromSrcIP 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromSrcIP(srcIPs []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`src_ip` IN (?)", srcIPs).Find(&results).Error

	return
}

// GetFromSrcPort 通过src_port获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromSrcPort(srcPort int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`src_port` = ?", srcPort).Find(&results).Error

	return
}

// GetBatchFromSrcPort 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromSrcPort(srcPorts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`src_port` IN (?)", srcPorts).Find(&results).Error

	return
}

// GetFromDestIP 通过dest_ip获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromDestIP(destIP string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`dest_ip` = ?", destIP).Find(&results).Error

	return
}

// GetBatchFromDestIP 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromDestIP(destIPs []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`dest_ip` IN (?)", destIPs).Find(&results).Error

	return
}

// GetFromDestPort 通过dest_port获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromDestPort(destPort int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`dest_port` = ?", destPort).Find(&results).Error

	return
}

// GetBatchFromDestPort 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromDestPort(destPorts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`dest_port` IN (?)", destPorts).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromResult(result int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromResult(results []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromStartTime(startTime time.Time) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromFinishTime 通过finish_time获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromFinishTime(finishTime time.Time) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`finish_time` = ?", finishTime).Find(&results).Error

	return
}

// GetBatchFromFinishTime 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromFinishTime(finishTimes []time.Time) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`finish_time` IN (?)", finishTimes).Find(&results).Error

	return
}

// GetFromAction 通过action获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromAction(action string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`action` = ?", action).Find(&results).Error

	return
}

// GetBatchFromAction 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromAction(actions []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`action` IN (?)", actions).Find(&results).Error

	return
}

// GetFromReplicaState 通过replica_state获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromReplicaState(replicaState string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`replica_state` = ?", replicaState).Find(&results).Error

	return
}

// GetBatchFromReplicaState 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromReplicaState(replicaStates []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`replica_state` IN (?)", replicaStates).Find(&results).Error

	return
}

// GetFromRebuildCount 通过rebuild_count获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromRebuildCount(rebuildCount int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`rebuild_count` = ?", rebuildCount).Find(&results).Error

	return
}

// GetBatchFromRebuildCount 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromRebuildCount(rebuildCounts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`rebuild_count` IN (?)", rebuildCounts).Find(&results).Error

	return
}

// GetFromTotalMacroBlock 通过total_macro_block获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromTotalMacroBlock(totalMacroBlock int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`total_macro_block` = ?", totalMacroBlock).Find(&results).Error

	return
}

// GetBatchFromTotalMacroBlock 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromTotalMacroBlock(totalMacroBlocks []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`total_macro_block` IN (?)", totalMacroBlocks).Find(&results).Error

	return
}

// GetFromReadyMacroBlock 通过ready_macro_block获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromReadyMacroBlock(readyMacroBlock int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`ready_macro_block` = ?", readyMacroBlock).Find(&results).Error

	return
}

// GetBatchFromReadyMacroBlock 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromReadyMacroBlock(readyMacroBlocks []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`ready_macro_block` IN (?)", readyMacroBlocks).Find(&results).Error

	return
}

// GetFromMajorCount 通过major_count获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromMajorCount(majorCount int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`major_count` = ?", majorCount).Find(&results).Error

	return
}

// GetBatchFromMajorCount 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromMajorCount(majorCounts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`major_count` IN (?)", majorCounts).Find(&results).Error

	return
}

// GetFromMiniMinorCount 通过mini_minor_count获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromMiniMinorCount(miniMinorCount int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`mini_minor_count` = ?", miniMinorCount).Find(&results).Error

	return
}

// GetBatchFromMiniMinorCount 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromMiniMinorCount(miniMinorCounts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`mini_minor_count` IN (?)", miniMinorCounts).Find(&results).Error

	return
}

// GetFromNormalMinorCount 通过normal_minor_count获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromNormalMinorCount(normalMinorCount int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`normal_minor_count` = ?", normalMinorCount).Find(&results).Error

	return
}

// GetBatchFromNormalMinorCount 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromNormalMinorCount(normalMinorCounts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`normal_minor_count` IN (?)", normalMinorCounts).Find(&results).Error

	return
}

// GetFromBufMinorCount 通过buf_minor_count获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromBufMinorCount(bufMinorCount int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`buf_minor_count` = ?", bufMinorCount).Find(&results).Error

	return
}

// GetBatchFromBufMinorCount 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromBufMinorCount(bufMinorCounts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`buf_minor_count` IN (?)", bufMinorCounts).Find(&results).Error

	return
}

// GetFromReuseCount 通过reuse_count获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromReuseCount(reuseCount int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`reuse_count` = ?", reuseCount).Find(&results).Error

	return
}

// GetBatchFromReuseCount 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromReuseCount(reuseCounts []int64) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`reuse_count` IN (?)", reuseCounts).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetFromComment(comment string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualPartitionMigrationStatusMgr) GetBatchFromComment(comments []string) (results []*AllVirtualPartitionMigrationStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionMigrationStatus{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
