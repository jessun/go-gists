package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPgPartitionInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPgPartitionInfoMgr open func
func AllVirtualPgPartitionInfoMgr(db *gorm.DB) *_AllVirtualPgPartitionInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPgPartitionInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPgPartitionInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_pg_partition_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPgPartitionInfoMgr) GetTableName() string {
	return "__all_virtual_pg_partition_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPgPartitionInfoMgr) Reset() *_AllVirtualPgPartitionInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPgPartitionInfoMgr) Get() (result AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPgPartitionInfoMgr) Gets() (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPgPartitionInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithTgID tg_id获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithTgID(tgID int64) Option {
	return optionFunc(func(o *options) { o.query["tg_id"] = tgID })
}

// WithPgIDx pg_idx获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithPgIDx(pgIDx int64) Option {
	return optionFunc(func(o *options) { o.query["pg_idx"] = pgIDx })
}

// WithMaxDecidedTransVersion max_decided_trans_version获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithMaxDecidedTransVersion(maxDecidedTransVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_decided_trans_version"] = maxDecidedTransVersion })
}

// WithMaxPassedTransTs max_passed_trans_ts获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithMaxPassedTransTs(maxPassedTransTs int64) Option {
	return optionFunc(func(o *options) { o.query["max_passed_trans_ts"] = maxPassedTransTs })
}

// WithFreezeTs freeze_ts获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithFreezeTs(freezeTs int64) Option {
	return optionFunc(func(o *options) { o.query["freeze_ts"] = freezeTs })
}

// WithAllowGc allow_gc获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithAllowGc(allowGc int8) Option {
	return optionFunc(func(o *options) { o.query["allow_gc"] = allowGc })
}

// WithPartitionState partition_state获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithPartitionState(partitionState string) Option {
	return optionFunc(func(o *options) { o.query["partition_state"] = partitionState })
}

// WithMinLogServiceTs min_log_service_ts获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithMinLogServiceTs(minLogServiceTs int64) Option {
	return optionFunc(func(o *options) { o.query["min_log_service_ts"] = minLogServiceTs })
}

// WithMinTransServiceTs min_trans_service_ts获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithMinTransServiceTs(minTransServiceTs int64) Option {
	return optionFunc(func(o *options) { o.query["min_trans_service_ts"] = minTransServiceTs })
}

// WithMinReplayEngineTs min_replay_engine_ts获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithMinReplayEngineTs(minReplayEngineTs int64) Option {
	return optionFunc(func(o *options) { o.query["min_replay_engine_ts"] = minReplayEngineTs })
}

// WithIsPg is_pg获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithIsPg(isPg int8) Option {
	return optionFunc(func(o *options) { o.query["is_pg"] = isPg })
}

// WithWeakReadTimestamp weak_read_timestamp获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithWeakReadTimestamp(weakReadTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["weak_read_timestamp"] = weakReadTimestamp })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualPgPartitionInfoMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPgPartitionInfoMgr) GetByOption(opts ...Option) (result AllVirtualPgPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPgPartitionInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPgPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPgPartitionInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPgPartitionInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where(options.query)
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
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromTgID 通过tg_id获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromTgID(tgID int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`tg_id` = ?", tgID).Find(&results).Error

	return
}

// GetBatchFromTgID 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromTgID(tgIDs []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`tg_id` IN (?)", tgIDs).Find(&results).Error

	return
}

// GetFromPgIDx 通过pg_idx获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromPgIDx(pgIDx int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`pg_idx` = ?", pgIDx).Find(&results).Error

	return
}

// GetBatchFromPgIDx 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromPgIDx(pgIDxs []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`pg_idx` IN (?)", pgIDxs).Find(&results).Error

	return
}

// GetFromMaxDecidedTransVersion 通过max_decided_trans_version获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromMaxDecidedTransVersion(maxDecidedTransVersion int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`max_decided_trans_version` = ?", maxDecidedTransVersion).Find(&results).Error

	return
}

// GetBatchFromMaxDecidedTransVersion 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromMaxDecidedTransVersion(maxDecidedTransVersions []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`max_decided_trans_version` IN (?)", maxDecidedTransVersions).Find(&results).Error

	return
}

// GetFromMaxPassedTransTs 通过max_passed_trans_ts获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromMaxPassedTransTs(maxPassedTransTs int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`max_passed_trans_ts` = ?", maxPassedTransTs).Find(&results).Error

	return
}

// GetBatchFromMaxPassedTransTs 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromMaxPassedTransTs(maxPassedTransTss []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`max_passed_trans_ts` IN (?)", maxPassedTransTss).Find(&results).Error

	return
}

// GetFromFreezeTs 通过freeze_ts获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromFreezeTs(freezeTs int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`freeze_ts` = ?", freezeTs).Find(&results).Error

	return
}

// GetBatchFromFreezeTs 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromFreezeTs(freezeTss []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`freeze_ts` IN (?)", freezeTss).Find(&results).Error

	return
}

// GetFromAllowGc 通过allow_gc获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromAllowGc(allowGc int8) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`allow_gc` = ?", allowGc).Find(&results).Error

	return
}

// GetBatchFromAllowGc 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromAllowGc(allowGcs []int8) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`allow_gc` IN (?)", allowGcs).Find(&results).Error

	return
}

// GetFromPartitionState 通过partition_state获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromPartitionState(partitionState string) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`partition_state` = ?", partitionState).Find(&results).Error

	return
}

// GetBatchFromPartitionState 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromPartitionState(partitionStates []string) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`partition_state` IN (?)", partitionStates).Find(&results).Error

	return
}

// GetFromMinLogServiceTs 通过min_log_service_ts获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromMinLogServiceTs(minLogServiceTs int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`min_log_service_ts` = ?", minLogServiceTs).Find(&results).Error

	return
}

// GetBatchFromMinLogServiceTs 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromMinLogServiceTs(minLogServiceTss []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`min_log_service_ts` IN (?)", minLogServiceTss).Find(&results).Error

	return
}

// GetFromMinTransServiceTs 通过min_trans_service_ts获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromMinTransServiceTs(minTransServiceTs int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`min_trans_service_ts` = ?", minTransServiceTs).Find(&results).Error

	return
}

// GetBatchFromMinTransServiceTs 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromMinTransServiceTs(minTransServiceTss []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`min_trans_service_ts` IN (?)", minTransServiceTss).Find(&results).Error

	return
}

// GetFromMinReplayEngineTs 通过min_replay_engine_ts获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromMinReplayEngineTs(minReplayEngineTs int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`min_replay_engine_ts` = ?", minReplayEngineTs).Find(&results).Error

	return
}

// GetBatchFromMinReplayEngineTs 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromMinReplayEngineTs(minReplayEngineTss []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`min_replay_engine_ts` IN (?)", minReplayEngineTss).Find(&results).Error

	return
}

// GetFromIsPg 通过is_pg获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromIsPg(isPg int8) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`is_pg` = ?", isPg).Find(&results).Error

	return
}

// GetBatchFromIsPg 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromIsPg(isPgs []int8) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`is_pg` IN (?)", isPgs).Find(&results).Error

	return
}

// GetFromWeakReadTimestamp 通过weak_read_timestamp获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromWeakReadTimestamp(weakReadTimestamp int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`weak_read_timestamp` = ?", weakReadTimestamp).Find(&results).Error

	return
}

// GetBatchFromWeakReadTimestamp 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromWeakReadTimestamp(weakReadTimestamps []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`weak_read_timestamp` IN (?)", weakReadTimestamps).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualPgPartitionInfoMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualPgPartitionInfoMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualPgPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgPartitionInfo{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
