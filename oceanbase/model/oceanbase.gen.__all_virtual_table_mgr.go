package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTableMgrMgr struct {
	*_BaseMgr
}

// AllVirtualTableMgrMgr open func
func AllVirtualTableMgrMgr(db *gorm.DB) *_AllVirtualTableMgrMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTableMgrMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTableMgrMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_table_mgr"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTableMgrMgr) GetTableName() string {
	return "__all_virtual_table_mgr"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTableMgrMgr) Reset() *_AllVirtualTableMgrMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTableMgrMgr) Get() (result AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTableMgrMgr) Gets() (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTableMgrMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTableMgrMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTableMgrMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualTableMgrMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableType table_type获取
func (obj *_AllVirtualTableMgrMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithTableID table_id获取
func (obj *_AllVirtualTableMgrMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualTableMgrMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithIndexID index_id获取
func (obj *_AllVirtualTableMgrMgr) WithIndexID(indexID int64) Option {
	return optionFunc(func(o *options) { o.query["index_id"] = indexID })
}

// WithBaseVersion base_version获取
func (obj *_AllVirtualTableMgrMgr) WithBaseVersion(baseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["base_version"] = baseVersion })
}

// WithMultiVersionStart multi_version_start获取
func (obj *_AllVirtualTableMgrMgr) WithMultiVersionStart(multiVersionStart int64) Option {
	return optionFunc(func(o *options) { o.query["multi_version_start"] = multiVersionStart })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualTableMgrMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithMaxMergedVersion max_merged_version获取
func (obj *_AllVirtualTableMgrMgr) WithMaxMergedVersion(maxMergedVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_merged_version"] = maxMergedVersion })
}

// WithUpperTransVersion upper_trans_version获取
func (obj *_AllVirtualTableMgrMgr) WithUpperTransVersion(upperTransVersion int64) Option {
	return optionFunc(func(o *options) { o.query["upper_trans_version"] = upperTransVersion })
}

// WithStartLogTs start_log_ts获取
func (obj *_AllVirtualTableMgrMgr) WithStartLogTs(startLogTs uint64) Option {
	return optionFunc(func(o *options) { o.query["start_log_ts"] = startLogTs })
}

// WithEndLogTs end_log_ts获取
func (obj *_AllVirtualTableMgrMgr) WithEndLogTs(endLogTs uint64) Option {
	return optionFunc(func(o *options) { o.query["end_log_ts"] = endLogTs })
}

// WithMaxLogTs max_log_ts获取
func (obj *_AllVirtualTableMgrMgr) WithMaxLogTs(maxLogTs uint64) Option {
	return optionFunc(func(o *options) { o.query["max_log_ts"] = maxLogTs })
}

// WithVersion version获取
func (obj *_AllVirtualTableMgrMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithLogicalDataVersion logical_data_version获取
func (obj *_AllVirtualTableMgrMgr) WithLogicalDataVersion(logicalDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["logical_data_version"] = logicalDataVersion })
}

// WithSize size获取
func (obj *_AllVirtualTableMgrMgr) WithSize(size int64) Option {
	return optionFunc(func(o *options) { o.query["size"] = size })
}

// WithCompactRow compact_row获取
func (obj *_AllVirtualTableMgrMgr) WithCompactRow(compactRow int64) Option {
	return optionFunc(func(o *options) { o.query["compact_row"] = compactRow })
}

// WithIsActive is_active获取
func (obj *_AllVirtualTableMgrMgr) WithIsActive(isActive int64) Option {
	return optionFunc(func(o *options) { o.query["is_active"] = isActive })
}

// WithTimestamp timestamp获取
func (obj *_AllVirtualTableMgrMgr) WithTimestamp(timestamp int64) Option {
	return optionFunc(func(o *options) { o.query["timestamp"] = timestamp })
}

// WithRef ref获取
func (obj *_AllVirtualTableMgrMgr) WithRef(ref int64) Option {
	return optionFunc(func(o *options) { o.query["ref"] = ref })
}

// WithWriteRef write_ref获取
func (obj *_AllVirtualTableMgrMgr) WithWriteRef(writeRef int64) Option {
	return optionFunc(func(o *options) { o.query["write_ref"] = writeRef })
}

// WithTrxCount trx_count获取
func (obj *_AllVirtualTableMgrMgr) WithTrxCount(trxCount int64) Option {
	return optionFunc(func(o *options) { o.query["trx_count"] = trxCount })
}

// WithPendingLogPersistingRowCnt pending_log_persisting_row_cnt获取
func (obj *_AllVirtualTableMgrMgr) WithPendingLogPersistingRowCnt(pendingLogPersistingRowCnt int64) Option {
	return optionFunc(func(o *options) { o.query["pending_log_persisting_row_cnt"] = pendingLogPersistingRowCnt })
}

// WithContainUncommittedRow contain_uncommitted_row获取
func (obj *_AllVirtualTableMgrMgr) WithContainUncommittedRow(containUncommittedRow int8) Option {
	return optionFunc(func(o *options) { o.query["contain_uncommitted_row"] = containUncommittedRow })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTableMgrMgr) GetByOption(opts ...Option) (result AllVirtualTableMgr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTableMgrMgr) GetByOptions(opts ...Option) (results []*AllVirtualTableMgr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTableMgrMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTableMgr, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromTableType(tableType int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromTableType(tableTypes []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromTableID(tableID int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromIndexID 通过index_id获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromIndexID(indexID int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`index_id` = ?", indexID).Find(&results).Error

	return
}

// GetBatchFromIndexID 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromIndexID(indexIDs []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`index_id` IN (?)", indexIDs).Find(&results).Error

	return
}

// GetFromBaseVersion 通过base_version获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromBaseVersion(baseVersion int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`base_version` = ?", baseVersion).Find(&results).Error

	return
}

// GetBatchFromBaseVersion 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromBaseVersion(baseVersions []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`base_version` IN (?)", baseVersions).Find(&results).Error

	return
}

// GetFromMultiVersionStart 通过multi_version_start获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromMultiVersionStart(multiVersionStart int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`multi_version_start` = ?", multiVersionStart).Find(&results).Error

	return
}

// GetBatchFromMultiVersionStart 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromMultiVersionStart(multiVersionStarts []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`multi_version_start` IN (?)", multiVersionStarts).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromMaxMergedVersion 通过max_merged_version获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromMaxMergedVersion(maxMergedVersion int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`max_merged_version` = ?", maxMergedVersion).Find(&results).Error

	return
}

// GetBatchFromMaxMergedVersion 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromMaxMergedVersion(maxMergedVersions []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`max_merged_version` IN (?)", maxMergedVersions).Find(&results).Error

	return
}

// GetFromUpperTransVersion 通过upper_trans_version获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromUpperTransVersion(upperTransVersion int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`upper_trans_version` = ?", upperTransVersion).Find(&results).Error

	return
}

// GetBatchFromUpperTransVersion 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromUpperTransVersion(upperTransVersions []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`upper_trans_version` IN (?)", upperTransVersions).Find(&results).Error

	return
}

// GetFromStartLogTs 通过start_log_ts获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromStartLogTs(startLogTs uint64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`start_log_ts` = ?", startLogTs).Find(&results).Error

	return
}

// GetBatchFromStartLogTs 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromStartLogTs(startLogTss []uint64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`start_log_ts` IN (?)", startLogTss).Find(&results).Error

	return
}

// GetFromEndLogTs 通过end_log_ts获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromEndLogTs(endLogTs uint64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`end_log_ts` = ?", endLogTs).Find(&results).Error

	return
}

// GetBatchFromEndLogTs 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromEndLogTs(endLogTss []uint64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`end_log_ts` IN (?)", endLogTss).Find(&results).Error

	return
}

// GetFromMaxLogTs 通过max_log_ts获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromMaxLogTs(maxLogTs uint64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`max_log_ts` = ?", maxLogTs).Find(&results).Error

	return
}

// GetBatchFromMaxLogTs 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromMaxLogTs(maxLogTss []uint64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`max_log_ts` IN (?)", maxLogTss).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromVersion(version int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromVersion(versions []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromLogicalDataVersion 通过logical_data_version获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromLogicalDataVersion(logicalDataVersion int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`logical_data_version` = ?", logicalDataVersion).Find(&results).Error

	return
}

// GetBatchFromLogicalDataVersion 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromLogicalDataVersion(logicalDataVersions []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`logical_data_version` IN (?)", logicalDataVersions).Find(&results).Error

	return
}

// GetFromSize 通过size获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromSize(size int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`size` = ?", size).Find(&results).Error

	return
}

// GetBatchFromSize 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromSize(sizes []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`size` IN (?)", sizes).Find(&results).Error

	return
}

// GetFromCompactRow 通过compact_row获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromCompactRow(compactRow int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`compact_row` = ?", compactRow).Find(&results).Error

	return
}

// GetBatchFromCompactRow 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromCompactRow(compactRows []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`compact_row` IN (?)", compactRows).Find(&results).Error

	return
}

// GetFromIsActive 通过is_active获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromIsActive(isActive int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`is_active` = ?", isActive).Find(&results).Error

	return
}

// GetBatchFromIsActive 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromIsActive(isActives []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`is_active` IN (?)", isActives).Find(&results).Error

	return
}

// GetFromTimestamp 通过timestamp获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromTimestamp(timestamp int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`timestamp` = ?", timestamp).Find(&results).Error

	return
}

// GetBatchFromTimestamp 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromTimestamp(timestamps []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`timestamp` IN (?)", timestamps).Find(&results).Error

	return
}

// GetFromRef 通过ref获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromRef(ref int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`ref` = ?", ref).Find(&results).Error

	return
}

// GetBatchFromRef 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromRef(refs []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`ref` IN (?)", refs).Find(&results).Error

	return
}

// GetFromWriteRef 通过write_ref获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromWriteRef(writeRef int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`write_ref` = ?", writeRef).Find(&results).Error

	return
}

// GetBatchFromWriteRef 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromWriteRef(writeRefs []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`write_ref` IN (?)", writeRefs).Find(&results).Error

	return
}

// GetFromTrxCount 通过trx_count获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromTrxCount(trxCount int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`trx_count` = ?", trxCount).Find(&results).Error

	return
}

// GetBatchFromTrxCount 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromTrxCount(trxCounts []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`trx_count` IN (?)", trxCounts).Find(&results).Error

	return
}

// GetFromPendingLogPersistingRowCnt 通过pending_log_persisting_row_cnt获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromPendingLogPersistingRowCnt(pendingLogPersistingRowCnt int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`pending_log_persisting_row_cnt` = ?", pendingLogPersistingRowCnt).Find(&results).Error

	return
}

// GetBatchFromPendingLogPersistingRowCnt 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromPendingLogPersistingRowCnt(pendingLogPersistingRowCnts []int64) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`pending_log_persisting_row_cnt` IN (?)", pendingLogPersistingRowCnts).Find(&results).Error

	return
}

// GetFromContainUncommittedRow 通过contain_uncommitted_row获取内容
func (obj *_AllVirtualTableMgrMgr) GetFromContainUncommittedRow(containUncommittedRow int8) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`contain_uncommitted_row` = ?", containUncommittedRow).Find(&results).Error

	return
}

// GetBatchFromContainUncommittedRow 批量查找
func (obj *_AllVirtualTableMgrMgr) GetBatchFromContainUncommittedRow(containUncommittedRows []int8) (results []*AllVirtualTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTableMgr{}).Where("`contain_uncommitted_row` IN (?)", containUncommittedRows).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
