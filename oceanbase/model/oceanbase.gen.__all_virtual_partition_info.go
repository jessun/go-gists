package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPartitionInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionInfoMgr open func
func AllVirtualPartitionInfoMgr(db *gorm.DB) *_AllVirtualPartitionInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionInfoMgr) GetTableName() string {
	return "__all_virtual_partition_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionInfoMgr) Reset() *_AllVirtualPartitionInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionInfoMgr) Get() (result AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionInfoMgr) Gets() (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualPartitionInfoMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithMaxDecidedTransVersion max_decided_trans_version获取
func (obj *_AllVirtualPartitionInfoMgr) WithMaxDecidedTransVersion(maxDecidedTransVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_decided_trans_version"] = maxDecidedTransVersion })
}

// WithMaxPassedTransTs max_passed_trans_ts获取
func (obj *_AllVirtualPartitionInfoMgr) WithMaxPassedTransTs(maxPassedTransTs int64) Option {
	return optionFunc(func(o *options) { o.query["max_passed_trans_ts"] = maxPassedTransTs })
}

// WithFreezeTs freeze_ts获取
func (obj *_AllVirtualPartitionInfoMgr) WithFreezeTs(freezeTs int64) Option {
	return optionFunc(func(o *options) { o.query["freeze_ts"] = freezeTs })
}

// WithAllowGc allow_gc获取
func (obj *_AllVirtualPartitionInfoMgr) WithAllowGc(allowGc int8) Option {
	return optionFunc(func(o *options) { o.query["allow_gc"] = allowGc })
}

// WithPartitionState partition_state获取
func (obj *_AllVirtualPartitionInfoMgr) WithPartitionState(partitionState string) Option {
	return optionFunc(func(o *options) { o.query["partition_state"] = partitionState })
}

// WithSstableReadRate sstable_read_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithSstableReadRate(sstableReadRate float64) Option {
	return optionFunc(func(o *options) { o.query["sstable_read_rate"] = sstableReadRate })
}

// WithSstableReadBytesRate sstable_read_bytes_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithSstableReadBytesRate(sstableReadBytesRate float64) Option {
	return optionFunc(func(o *options) { o.query["sstable_read_bytes_rate"] = sstableReadBytesRate })
}

// WithSstableWriteRate sstable_write_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithSstableWriteRate(sstableWriteRate float64) Option {
	return optionFunc(func(o *options) { o.query["sstable_write_rate"] = sstableWriteRate })
}

// WithSstableWriteBytesRate sstable_write_bytes_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithSstableWriteBytesRate(sstableWriteBytesRate float64) Option {
	return optionFunc(func(o *options) { o.query["sstable_write_bytes_rate"] = sstableWriteBytesRate })
}

// WithLogWriteRate log_write_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithLogWriteRate(logWriteRate float64) Option {
	return optionFunc(func(o *options) { o.query["log_write_rate"] = logWriteRate })
}

// WithLogWriteBytesRate log_write_bytes_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithLogWriteBytesRate(logWriteBytesRate float64) Option {
	return optionFunc(func(o *options) { o.query["log_write_bytes_rate"] = logWriteBytesRate })
}

// WithMemtableBytes memtable_bytes获取
func (obj *_AllVirtualPartitionInfoMgr) WithMemtableBytes(memtableBytes int64) Option {
	return optionFunc(func(o *options) { o.query["memtable_bytes"] = memtableBytes })
}

// WithCPUUtimeRate cpu_utime_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithCPUUtimeRate(cpuUtimeRate float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_utime_rate"] = cpuUtimeRate })
}

// WithCPUStimeRate cpu_stime_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithCPUStimeRate(cpuStimeRate float64) Option {
	return optionFunc(func(o *options) { o.query["cpu_stime_rate"] = cpuStimeRate })
}

// WithNetInRate net_in_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithNetInRate(netInRate float64) Option {
	return optionFunc(func(o *options) { o.query["net_in_rate"] = netInRate })
}

// WithNetInBytesRate net_in_bytes_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithNetInBytesRate(netInBytesRate float64) Option {
	return optionFunc(func(o *options) { o.query["net_in_bytes_rate"] = netInBytesRate })
}

// WithNetOutRate net_out_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithNetOutRate(netOutRate float64) Option {
	return optionFunc(func(o *options) { o.query["net_out_rate"] = netOutRate })
}

// WithNetOutBytesRate net_out_bytes_rate获取
func (obj *_AllVirtualPartitionInfoMgr) WithNetOutBytesRate(netOutBytesRate float64) Option {
	return optionFunc(func(o *options) { o.query["net_out_bytes_rate"] = netOutBytesRate })
}

// WithMinLogServiceTs min_log_service_ts获取
func (obj *_AllVirtualPartitionInfoMgr) WithMinLogServiceTs(minLogServiceTs int64) Option {
	return optionFunc(func(o *options) { o.query["min_log_service_ts"] = minLogServiceTs })
}

// WithMinTransServiceTs min_trans_service_ts获取
func (obj *_AllVirtualPartitionInfoMgr) WithMinTransServiceTs(minTransServiceTs int64) Option {
	return optionFunc(func(o *options) { o.query["min_trans_service_ts"] = minTransServiceTs })
}

// WithMinReplayEngineTs min_replay_engine_ts获取
func (obj *_AllVirtualPartitionInfoMgr) WithMinReplayEngineTs(minReplayEngineTs int64) Option {
	return optionFunc(func(o *options) { o.query["min_replay_engine_ts"] = minReplayEngineTs })
}

// WithIsNeedRebuild is_need_rebuild获取
func (obj *_AllVirtualPartitionInfoMgr) WithIsNeedRebuild(isNeedRebuild int8) Option {
	return optionFunc(func(o *options) { o.query["is_need_rebuild"] = isNeedRebuild })
}

// WithPgPartitionCount pg_partition_count获取
func (obj *_AllVirtualPartitionInfoMgr) WithPgPartitionCount(pgPartitionCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_partition_count"] = pgPartitionCount })
}

// WithIsPg is_pg获取
func (obj *_AllVirtualPartitionInfoMgr) WithIsPg(isPg int8) Option {
	return optionFunc(func(o *options) { o.query["is_pg"] = isPg })
}

// WithWeakReadTimestamp weak_read_timestamp获取
func (obj *_AllVirtualPartitionInfoMgr) WithWeakReadTimestamp(weakReadTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["weak_read_timestamp"] = weakReadTimestamp })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualPartitionInfoMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithLastReplayLogID last_replay_log_id获取
func (obj *_AllVirtualPartitionInfoMgr) WithLastReplayLogID(lastReplayLogID int64) Option {
	return optionFunc(func(o *options) { o.query["last_replay_log_id"] = lastReplayLogID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualPartitionInfoMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithLastReplayLogTs last_replay_log_ts获取
func (obj *_AllVirtualPartitionInfoMgr) WithLastReplayLogTs(lastReplayLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["last_replay_log_ts"] = lastReplayLogTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionInfoMgr) GetByOption(opts ...Option) (result AllVirtualPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where(options.query)
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
func (obj *_AllVirtualPartitionInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromMaxDecidedTransVersion 通过max_decided_trans_version获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromMaxDecidedTransVersion(maxDecidedTransVersion int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`max_decided_trans_version` = ?", maxDecidedTransVersion).Find(&results).Error

	return
}

// GetBatchFromMaxDecidedTransVersion 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromMaxDecidedTransVersion(maxDecidedTransVersions []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`max_decided_trans_version` IN (?)", maxDecidedTransVersions).Find(&results).Error

	return
}

// GetFromMaxPassedTransTs 通过max_passed_trans_ts获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromMaxPassedTransTs(maxPassedTransTs int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`max_passed_trans_ts` = ?", maxPassedTransTs).Find(&results).Error

	return
}

// GetBatchFromMaxPassedTransTs 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromMaxPassedTransTs(maxPassedTransTss []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`max_passed_trans_ts` IN (?)", maxPassedTransTss).Find(&results).Error

	return
}

// GetFromFreezeTs 通过freeze_ts获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromFreezeTs(freezeTs int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`freeze_ts` = ?", freezeTs).Find(&results).Error

	return
}

// GetBatchFromFreezeTs 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromFreezeTs(freezeTss []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`freeze_ts` IN (?)", freezeTss).Find(&results).Error

	return
}

// GetFromAllowGc 通过allow_gc获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromAllowGc(allowGc int8) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`allow_gc` = ?", allowGc).Find(&results).Error

	return
}

// GetBatchFromAllowGc 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromAllowGc(allowGcs []int8) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`allow_gc` IN (?)", allowGcs).Find(&results).Error

	return
}

// GetFromPartitionState 通过partition_state获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromPartitionState(partitionState string) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`partition_state` = ?", partitionState).Find(&results).Error

	return
}

// GetBatchFromPartitionState 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromPartitionState(partitionStates []string) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`partition_state` IN (?)", partitionStates).Find(&results).Error

	return
}

// GetFromSstableReadRate 通过sstable_read_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromSstableReadRate(sstableReadRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_read_rate` = ?", sstableReadRate).Find(&results).Error

	return
}

// GetBatchFromSstableReadRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromSstableReadRate(sstableReadRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_read_rate` IN (?)", sstableReadRates).Find(&results).Error

	return
}

// GetFromSstableReadBytesRate 通过sstable_read_bytes_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromSstableReadBytesRate(sstableReadBytesRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_read_bytes_rate` = ?", sstableReadBytesRate).Find(&results).Error

	return
}

// GetBatchFromSstableReadBytesRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromSstableReadBytesRate(sstableReadBytesRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_read_bytes_rate` IN (?)", sstableReadBytesRates).Find(&results).Error

	return
}

// GetFromSstableWriteRate 通过sstable_write_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromSstableWriteRate(sstableWriteRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_write_rate` = ?", sstableWriteRate).Find(&results).Error

	return
}

// GetBatchFromSstableWriteRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromSstableWriteRate(sstableWriteRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_write_rate` IN (?)", sstableWriteRates).Find(&results).Error

	return
}

// GetFromSstableWriteBytesRate 通过sstable_write_bytes_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromSstableWriteBytesRate(sstableWriteBytesRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_write_bytes_rate` = ?", sstableWriteBytesRate).Find(&results).Error

	return
}

// GetBatchFromSstableWriteBytesRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromSstableWriteBytesRate(sstableWriteBytesRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`sstable_write_bytes_rate` IN (?)", sstableWriteBytesRates).Find(&results).Error

	return
}

// GetFromLogWriteRate 通过log_write_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromLogWriteRate(logWriteRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`log_write_rate` = ?", logWriteRate).Find(&results).Error

	return
}

// GetBatchFromLogWriteRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromLogWriteRate(logWriteRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`log_write_rate` IN (?)", logWriteRates).Find(&results).Error

	return
}

// GetFromLogWriteBytesRate 通过log_write_bytes_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromLogWriteBytesRate(logWriteBytesRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`log_write_bytes_rate` = ?", logWriteBytesRate).Find(&results).Error

	return
}

// GetBatchFromLogWriteBytesRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromLogWriteBytesRate(logWriteBytesRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`log_write_bytes_rate` IN (?)", logWriteBytesRates).Find(&results).Error

	return
}

// GetFromMemtableBytes 通过memtable_bytes获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromMemtableBytes(memtableBytes int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`memtable_bytes` = ?", memtableBytes).Find(&results).Error

	return
}

// GetBatchFromMemtableBytes 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromMemtableBytes(memtableBytess []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`memtable_bytes` IN (?)", memtableBytess).Find(&results).Error

	return
}

// GetFromCPUUtimeRate 通过cpu_utime_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromCPUUtimeRate(cpuUtimeRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`cpu_utime_rate` = ?", cpuUtimeRate).Find(&results).Error

	return
}

// GetBatchFromCPUUtimeRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromCPUUtimeRate(cpuUtimeRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`cpu_utime_rate` IN (?)", cpuUtimeRates).Find(&results).Error

	return
}

// GetFromCPUStimeRate 通过cpu_stime_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromCPUStimeRate(cpuStimeRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`cpu_stime_rate` = ?", cpuStimeRate).Find(&results).Error

	return
}

// GetBatchFromCPUStimeRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromCPUStimeRate(cpuStimeRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`cpu_stime_rate` IN (?)", cpuStimeRates).Find(&results).Error

	return
}

// GetFromNetInRate 通过net_in_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromNetInRate(netInRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_in_rate` = ?", netInRate).Find(&results).Error

	return
}

// GetBatchFromNetInRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromNetInRate(netInRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_in_rate` IN (?)", netInRates).Find(&results).Error

	return
}

// GetFromNetInBytesRate 通过net_in_bytes_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromNetInBytesRate(netInBytesRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_in_bytes_rate` = ?", netInBytesRate).Find(&results).Error

	return
}

// GetBatchFromNetInBytesRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromNetInBytesRate(netInBytesRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_in_bytes_rate` IN (?)", netInBytesRates).Find(&results).Error

	return
}

// GetFromNetOutRate 通过net_out_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromNetOutRate(netOutRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_out_rate` = ?", netOutRate).Find(&results).Error

	return
}

// GetBatchFromNetOutRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromNetOutRate(netOutRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_out_rate` IN (?)", netOutRates).Find(&results).Error

	return
}

// GetFromNetOutBytesRate 通过net_out_bytes_rate获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromNetOutBytesRate(netOutBytesRate float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_out_bytes_rate` = ?", netOutBytesRate).Find(&results).Error

	return
}

// GetBatchFromNetOutBytesRate 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromNetOutBytesRate(netOutBytesRates []float64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`net_out_bytes_rate` IN (?)", netOutBytesRates).Find(&results).Error

	return
}

// GetFromMinLogServiceTs 通过min_log_service_ts获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromMinLogServiceTs(minLogServiceTs int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`min_log_service_ts` = ?", minLogServiceTs).Find(&results).Error

	return
}

// GetBatchFromMinLogServiceTs 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromMinLogServiceTs(minLogServiceTss []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`min_log_service_ts` IN (?)", minLogServiceTss).Find(&results).Error

	return
}

// GetFromMinTransServiceTs 通过min_trans_service_ts获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromMinTransServiceTs(minTransServiceTs int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`min_trans_service_ts` = ?", minTransServiceTs).Find(&results).Error

	return
}

// GetBatchFromMinTransServiceTs 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromMinTransServiceTs(minTransServiceTss []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`min_trans_service_ts` IN (?)", minTransServiceTss).Find(&results).Error

	return
}

// GetFromMinReplayEngineTs 通过min_replay_engine_ts获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromMinReplayEngineTs(minReplayEngineTs int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`min_replay_engine_ts` = ?", minReplayEngineTs).Find(&results).Error

	return
}

// GetBatchFromMinReplayEngineTs 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromMinReplayEngineTs(minReplayEngineTss []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`min_replay_engine_ts` IN (?)", minReplayEngineTss).Find(&results).Error

	return
}

// GetFromIsNeedRebuild 通过is_need_rebuild获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromIsNeedRebuild(isNeedRebuild int8) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`is_need_rebuild` = ?", isNeedRebuild).Find(&results).Error

	return
}

// GetBatchFromIsNeedRebuild 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromIsNeedRebuild(isNeedRebuilds []int8) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`is_need_rebuild` IN (?)", isNeedRebuilds).Find(&results).Error

	return
}

// GetFromPgPartitionCount 通过pg_partition_count获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromPgPartitionCount(pgPartitionCount int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`pg_partition_count` = ?", pgPartitionCount).Find(&results).Error

	return
}

// GetBatchFromPgPartitionCount 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromPgPartitionCount(pgPartitionCounts []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`pg_partition_count` IN (?)", pgPartitionCounts).Find(&results).Error

	return
}

// GetFromIsPg 通过is_pg获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromIsPg(isPg int8) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`is_pg` = ?", isPg).Find(&results).Error

	return
}

// GetBatchFromIsPg 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromIsPg(isPgs []int8) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`is_pg` IN (?)", isPgs).Find(&results).Error

	return
}

// GetFromWeakReadTimestamp 通过weak_read_timestamp获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromWeakReadTimestamp(weakReadTimestamp int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`weak_read_timestamp` = ?", weakReadTimestamp).Find(&results).Error

	return
}

// GetBatchFromWeakReadTimestamp 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromWeakReadTimestamp(weakReadTimestamps []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`weak_read_timestamp` IN (?)", weakReadTimestamps).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromLastReplayLogID 通过last_replay_log_id获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromLastReplayLogID(lastReplayLogID int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`last_replay_log_id` = ?", lastReplayLogID).Find(&results).Error

	return
}

// GetBatchFromLastReplayLogID 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromLastReplayLogID(lastReplayLogIDs []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`last_replay_log_id` IN (?)", lastReplayLogIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromLastReplayLogTs 通过last_replay_log_ts获取内容
func (obj *_AllVirtualPartitionInfoMgr) GetFromLastReplayLogTs(lastReplayLogTs int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`last_replay_log_ts` = ?", lastReplayLogTs).Find(&results).Error

	return
}

// GetBatchFromLastReplayLogTs 批量查找
func (obj *_AllVirtualPartitionInfoMgr) GetBatchFromLastReplayLogTs(lastReplayLogTss []int64) (results []*AllVirtualPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionInfo{}).Where("`last_replay_log_ts` IN (?)", lastReplayLogTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
