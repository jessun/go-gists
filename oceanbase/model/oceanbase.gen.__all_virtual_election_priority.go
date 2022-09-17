package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualElectionPriorityMgr struct {
	*_BaseMgr
}

// AllVirtualElectionPriorityMgr open func
func AllVirtualElectionPriorityMgr(db *gorm.DB) *_AllVirtualElectionPriorityMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualElectionPriorityMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualElectionPriorityMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_election_priority"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualElectionPriorityMgr) GetTableName() string {
	return "__all_virtual_election_priority"
}

// Reset 重置gorm会话
func (obj *_AllVirtualElectionPriorityMgr) Reset() *_AllVirtualElectionPriorityMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualElectionPriorityMgr) Get() (result AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualElectionPriorityMgr) Gets() (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualElectionPriorityMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualElectionPriorityMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualElectionPriorityMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualElectionPriorityMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualElectionPriorityMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualElectionPriorityMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithRole role获取
func (obj *_AllVirtualElectionPriorityMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithIsCandidate is_candidate获取
func (obj *_AllVirtualElectionPriorityMgr) WithIsCandidate(isCandidate int8) Option {
	return optionFunc(func(o *options) { o.query["is_candidate"] = isCandidate })
}

// WithMembershipVersion membership_version获取
func (obj *_AllVirtualElectionPriorityMgr) WithMembershipVersion(membershipVersion int64) Option {
	return optionFunc(func(o *options) { o.query["membership_version"] = membershipVersion })
}

// WithLogID log_id获取
func (obj *_AllVirtualElectionPriorityMgr) WithLogID(logID uint64) Option {
	return optionFunc(func(o *options) { o.query["log_id"] = logID })
}

// WithLocality locality获取
func (obj *_AllVirtualElectionPriorityMgr) WithLocality(locality uint64) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}

// WithSystemScore system_score获取
func (obj *_AllVirtualElectionPriorityMgr) WithSystemScore(systemScore int64) Option {
	return optionFunc(func(o *options) { o.query["system_score"] = systemScore })
}

// WithIsTenantActive is_tenant_active获取
func (obj *_AllVirtualElectionPriorityMgr) WithIsTenantActive(isTenantActive int8) Option {
	return optionFunc(func(o *options) { o.query["is_tenant_active"] = isTenantActive })
}

// WithOnRevokeBlacklist on_revoke_blacklist获取
func (obj *_AllVirtualElectionPriorityMgr) WithOnRevokeBlacklist(onRevokeBlacklist int8) Option {
	return optionFunc(func(o *options) { o.query["on_revoke_blacklist"] = onRevokeBlacklist })
}

// WithOnLoopBlacklist on_loop_blacklist获取
func (obj *_AllVirtualElectionPriorityMgr) WithOnLoopBlacklist(onLoopBlacklist int8) Option {
	return optionFunc(func(o *options) { o.query["on_loop_blacklist"] = onLoopBlacklist })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualElectionPriorityMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithServerStatus server_status获取
func (obj *_AllVirtualElectionPriorityMgr) WithServerStatus(serverStatus int64) Option {
	return optionFunc(func(o *options) { o.query["server_status"] = serverStatus })
}

// WithIsClogDiskFull is_clog_disk_full获取
func (obj *_AllVirtualElectionPriorityMgr) WithIsClogDiskFull(isClogDiskFull int8) Option {
	return optionFunc(func(o *options) { o.query["is_clog_disk_full"] = isClogDiskFull })
}

// WithIsOffline is_offline获取
func (obj *_AllVirtualElectionPriorityMgr) WithIsOffline(isOffline int8) Option {
	return optionFunc(func(o *options) { o.query["is_offline"] = isOffline })
}

// WithIsNeedRebuild is_need_rebuild获取
func (obj *_AllVirtualElectionPriorityMgr) WithIsNeedRebuild(isNeedRebuild int8) Option {
	return optionFunc(func(o *options) { o.query["is_need_rebuild"] = isNeedRebuild })
}

// WithIsPartitionCandidate is_partition_candidate获取
func (obj *_AllVirtualElectionPriorityMgr) WithIsPartitionCandidate(isPartitionCandidate int8) Option {
	return optionFunc(func(o *options) { o.query["is_partition_candidate"] = isPartitionCandidate })
}

// WithIsDiskError is_disk_error获取
func (obj *_AllVirtualElectionPriorityMgr) WithIsDiskError(isDiskError int8) Option {
	return optionFunc(func(o *options) { o.query["is_disk_error"] = isDiskError })
}

// WithMemstorePercent memstore_percent获取
func (obj *_AllVirtualElectionPriorityMgr) WithMemstorePercent(memstorePercent int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_percent"] = memstorePercent })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualElectionPriorityMgr) GetByOption(opts ...Option) (result AllVirtualElectionPriority, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualElectionPriorityMgr) GetByOptions(opts ...Option) (results []*AllVirtualElectionPriority, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualElectionPriorityMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualElectionPriority, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where(options.query)
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
func (obj *_AllVirtualElectionPriorityMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromTableID(tableID int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromRole(role int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromIsCandidate 通过is_candidate获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromIsCandidate(isCandidate int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_candidate` = ?", isCandidate).Find(&results).Error

	return
}

// GetBatchFromIsCandidate 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromIsCandidate(isCandidates []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_candidate` IN (?)", isCandidates).Find(&results).Error

	return
}

// GetFromMembershipVersion 通过membership_version获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromMembershipVersion(membershipVersion int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`membership_version` = ?", membershipVersion).Find(&results).Error

	return
}

// GetBatchFromMembershipVersion 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromMembershipVersion(membershipVersions []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`membership_version` IN (?)", membershipVersions).Find(&results).Error

	return
}

// GetFromLogID 通过log_id获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromLogID(logID uint64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`log_id` = ?", logID).Find(&results).Error

	return
}

// GetBatchFromLogID 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromLogID(logIDs []uint64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`log_id` IN (?)", logIDs).Find(&results).Error

	return
}

// GetFromLocality 通过locality获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromLocality(locality uint64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`locality` = ?", locality).Find(&results).Error

	return
}

// GetBatchFromLocality 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromLocality(localitys []uint64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`locality` IN (?)", localitys).Find(&results).Error

	return
}

// GetFromSystemScore 通过system_score获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromSystemScore(systemScore int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`system_score` = ?", systemScore).Find(&results).Error

	return
}

// GetBatchFromSystemScore 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromSystemScore(systemScores []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`system_score` IN (?)", systemScores).Find(&results).Error

	return
}

// GetFromIsTenantActive 通过is_tenant_active获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromIsTenantActive(isTenantActive int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_tenant_active` = ?", isTenantActive).Find(&results).Error

	return
}

// GetBatchFromIsTenantActive 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromIsTenantActive(isTenantActives []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_tenant_active` IN (?)", isTenantActives).Find(&results).Error

	return
}

// GetFromOnRevokeBlacklist 通过on_revoke_blacklist获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromOnRevokeBlacklist(onRevokeBlacklist int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`on_revoke_blacklist` = ?", onRevokeBlacklist).Find(&results).Error

	return
}

// GetBatchFromOnRevokeBlacklist 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromOnRevokeBlacklist(onRevokeBlacklists []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`on_revoke_blacklist` IN (?)", onRevokeBlacklists).Find(&results).Error

	return
}

// GetFromOnLoopBlacklist 通过on_loop_blacklist获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromOnLoopBlacklist(onLoopBlacklist int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`on_loop_blacklist` = ?", onLoopBlacklist).Find(&results).Error

	return
}

// GetBatchFromOnLoopBlacklist 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromOnLoopBlacklist(onLoopBlacklists []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`on_loop_blacklist` IN (?)", onLoopBlacklists).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromServerStatus 通过server_status获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromServerStatus(serverStatus int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`server_status` = ?", serverStatus).Find(&results).Error

	return
}

// GetBatchFromServerStatus 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromServerStatus(serverStatuss []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`server_status` IN (?)", serverStatuss).Find(&results).Error

	return
}

// GetFromIsClogDiskFull 通过is_clog_disk_full获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromIsClogDiskFull(isClogDiskFull int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_clog_disk_full` = ?", isClogDiskFull).Find(&results).Error

	return
}

// GetBatchFromIsClogDiskFull 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromIsClogDiskFull(isClogDiskFulls []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_clog_disk_full` IN (?)", isClogDiskFulls).Find(&results).Error

	return
}

// GetFromIsOffline 通过is_offline获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromIsOffline(isOffline int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_offline` = ?", isOffline).Find(&results).Error

	return
}

// GetBatchFromIsOffline 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromIsOffline(isOfflines []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_offline` IN (?)", isOfflines).Find(&results).Error

	return
}

// GetFromIsNeedRebuild 通过is_need_rebuild获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromIsNeedRebuild(isNeedRebuild int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_need_rebuild` = ?", isNeedRebuild).Find(&results).Error

	return
}

// GetBatchFromIsNeedRebuild 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromIsNeedRebuild(isNeedRebuilds []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_need_rebuild` IN (?)", isNeedRebuilds).Find(&results).Error

	return
}

// GetFromIsPartitionCandidate 通过is_partition_candidate获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromIsPartitionCandidate(isPartitionCandidate int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_partition_candidate` = ?", isPartitionCandidate).Find(&results).Error

	return
}

// GetBatchFromIsPartitionCandidate 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromIsPartitionCandidate(isPartitionCandidates []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_partition_candidate` IN (?)", isPartitionCandidates).Find(&results).Error

	return
}

// GetFromIsDiskError 通过is_disk_error获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromIsDiskError(isDiskError int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_disk_error` = ?", isDiskError).Find(&results).Error

	return
}

// GetBatchFromIsDiskError 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromIsDiskError(isDiskErrors []int8) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`is_disk_error` IN (?)", isDiskErrors).Find(&results).Error

	return
}

// GetFromMemstorePercent 通过memstore_percent获取内容
func (obj *_AllVirtualElectionPriorityMgr) GetFromMemstorePercent(memstorePercent int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`memstore_percent` = ?", memstorePercent).Find(&results).Error

	return
}

// GetBatchFromMemstorePercent 批量查找
func (obj *_AllVirtualElectionPriorityMgr) GetBatchFromMemstorePercent(memstorePercents []int64) (results []*AllVirtualElectionPriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionPriority{}).Where("`memstore_percent` IN (?)", memstorePercents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
