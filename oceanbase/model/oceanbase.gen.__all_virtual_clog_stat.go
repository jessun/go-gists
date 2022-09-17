package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualClogStatMgr struct {
	*_BaseMgr
}

// AllVirtualClogStatMgr open func
func AllVirtualClogStatMgr(db *gorm.DB) *_AllVirtualClogStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualClogStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualClogStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_clog_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualClogStatMgr) GetTableName() string {
	return "__all_virtual_clog_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualClogStatMgr) Reset() *_AllVirtualClogStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualClogStatMgr) Get() (result AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualClogStatMgr) Gets() (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualClogStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualClogStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualClogStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualClogStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualClogStatMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualClogStatMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithRole role获取
func (obj *_AllVirtualClogStatMgr) WithRole(role string) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithStatus status获取
func (obj *_AllVirtualClogStatMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithLeader leader获取
func (obj *_AllVirtualClogStatMgr) WithLeader(leader string) Option {
	return optionFunc(func(o *options) { o.query["leader"] = leader })
}

// WithLastIndexLogID last_index_log_id获取
func (obj *_AllVirtualClogStatMgr) WithLastIndexLogID(lastIndexLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["last_index_log_id"] = lastIndexLogID })
}

// WithLastIndexLogTimestamp last_index_log_timestamp获取
func (obj *_AllVirtualClogStatMgr) WithLastIndexLogTimestamp(lastIndexLogTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_index_log_timestamp"] = lastIndexLogTimestamp })
}

// WithLastLogID last_log_id获取
func (obj *_AllVirtualClogStatMgr) WithLastLogID(lastLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["last_log_id"] = lastLogID })
}

// WithActiveFreezeVersion active_freeze_version获取
func (obj *_AllVirtualClogStatMgr) WithActiveFreezeVersion(activeFreezeVersion string) Option {
	return optionFunc(func(o *options) { o.query["active_freeze_version"] = activeFreezeVersion })
}

// WithCurrMemberList curr_member_list获取
func (obj *_AllVirtualClogStatMgr) WithCurrMemberList(currMemberList string) Option {
	return optionFunc(func(o *options) { o.query["curr_member_list"] = currMemberList })
}

// WithMemberShipLogID member_ship_log_id获取
func (obj *_AllVirtualClogStatMgr) WithMemberShipLogID(memberShipLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["member_ship_log_id"] = memberShipLogID })
}

// WithIsOffline is_offline获取
func (obj *_AllVirtualClogStatMgr) WithIsOffline(isOffline int8) Option {
	return optionFunc(func(o *options) { o.query["is_offline"] = isOffline })
}

// WithIsInSync is_in_sync获取
func (obj *_AllVirtualClogStatMgr) WithIsInSync(isInSync int8) Option {
	return optionFunc(func(o *options) { o.query["is_in_sync"] = isInSync })
}

// WithStartID start_id获取
func (obj *_AllVirtualClogStatMgr) WithStartID(startID uint64) Option {
	return optionFunc(func(o *options) { o.query["start_id"] = startID })
}

// WithParent parent获取
func (obj *_AllVirtualClogStatMgr) WithParent(parent string) Option {
	return optionFunc(func(o *options) { o.query["parent"] = parent })
}

// WithChildrenList children_list获取
func (obj *_AllVirtualClogStatMgr) WithChildrenList(childrenList string) Option {
	return optionFunc(func(o *options) { o.query["children_list"] = childrenList })
}

// WithAccuLogCount accu_log_count获取
func (obj *_AllVirtualClogStatMgr) WithAccuLogCount(accuLogCount uint64) Option {
	return optionFunc(func(o *options) { o.query["accu_log_count"] = accuLogCount })
}

// WithAccuLogDelay accu_log_delay获取
func (obj *_AllVirtualClogStatMgr) WithAccuLogDelay(accuLogDelay uint64) Option {
	return optionFunc(func(o *options) { o.query["accu_log_delay"] = accuLogDelay })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualClogStatMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithAllowGc allow_gc获取
func (obj *_AllVirtualClogStatMgr) WithAllowGc(allowGc int8) Option {
	return optionFunc(func(o *options) { o.query["allow_gc"] = allowGc })
}

// WithQuorum quorum获取
func (obj *_AllVirtualClogStatMgr) WithQuorum(quorum int64) Option {
	return optionFunc(func(o *options) { o.query["quorum"] = quorum })
}

// WithIsNeedRebuild is_need_rebuild获取
func (obj *_AllVirtualClogStatMgr) WithIsNeedRebuild(isNeedRebuild int8) Option {
	return optionFunc(func(o *options) { o.query["is_need_rebuild"] = isNeedRebuild })
}

// WithNextReplayTsDelta next_replay_ts_delta获取
func (obj *_AllVirtualClogStatMgr) WithNextReplayTsDelta(nextReplayTsDelta uint64) Option {
	return optionFunc(func(o *options) { o.query["next_replay_ts_delta"] = nextReplayTsDelta })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualClogStatMgr) GetByOption(opts ...Option) (result AllVirtualClogStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualClogStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualClogStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualClogStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualClogStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where(options.query)
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
func (obj *_AllVirtualClogStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualClogStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualClogStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualClogStatMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualClogStatMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualClogStatMgr) GetFromRole(role string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromRole(roles []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualClogStatMgr) GetFromStatus(status string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromLeader 通过leader获取内容
func (obj *_AllVirtualClogStatMgr) GetFromLeader(leader string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`leader` = ?", leader).Find(&results).Error

	return
}

// GetBatchFromLeader 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromLeader(leaders []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`leader` IN (?)", leaders).Find(&results).Error

	return
}

// GetFromLastIndexLogID 通过last_index_log_id获取内容
func (obj *_AllVirtualClogStatMgr) GetFromLastIndexLogID(lastIndexLogID uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`last_index_log_id` = ?", lastIndexLogID).Find(&results).Error

	return
}

// GetBatchFromLastIndexLogID 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromLastIndexLogID(lastIndexLogIDs []uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`last_index_log_id` IN (?)", lastIndexLogIDs).Find(&results).Error

	return
}

// GetFromLastIndexLogTimestamp 通过last_index_log_timestamp获取内容
func (obj *_AllVirtualClogStatMgr) GetFromLastIndexLogTimestamp(lastIndexLogTimestamp time.Time) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`last_index_log_timestamp` = ?", lastIndexLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromLastIndexLogTimestamp 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromLastIndexLogTimestamp(lastIndexLogTimestamps []time.Time) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`last_index_log_timestamp` IN (?)", lastIndexLogTimestamps).Find(&results).Error

	return
}

// GetFromLastLogID 通过last_log_id获取内容
func (obj *_AllVirtualClogStatMgr) GetFromLastLogID(lastLogID uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`last_log_id` = ?", lastLogID).Find(&results).Error

	return
}

// GetBatchFromLastLogID 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromLastLogID(lastLogIDs []uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`last_log_id` IN (?)", lastLogIDs).Find(&results).Error

	return
}

// GetFromActiveFreezeVersion 通过active_freeze_version获取内容
func (obj *_AllVirtualClogStatMgr) GetFromActiveFreezeVersion(activeFreezeVersion string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`active_freeze_version` = ?", activeFreezeVersion).Find(&results).Error

	return
}

// GetBatchFromActiveFreezeVersion 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromActiveFreezeVersion(activeFreezeVersions []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`active_freeze_version` IN (?)", activeFreezeVersions).Find(&results).Error

	return
}

// GetFromCurrMemberList 通过curr_member_list获取内容
func (obj *_AllVirtualClogStatMgr) GetFromCurrMemberList(currMemberList string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`curr_member_list` = ?", currMemberList).Find(&results).Error

	return
}

// GetBatchFromCurrMemberList 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromCurrMemberList(currMemberLists []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`curr_member_list` IN (?)", currMemberLists).Find(&results).Error

	return
}

// GetFromMemberShipLogID 通过member_ship_log_id获取内容
func (obj *_AllVirtualClogStatMgr) GetFromMemberShipLogID(memberShipLogID uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`member_ship_log_id` = ?", memberShipLogID).Find(&results).Error

	return
}

// GetBatchFromMemberShipLogID 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromMemberShipLogID(memberShipLogIDs []uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`member_ship_log_id` IN (?)", memberShipLogIDs).Find(&results).Error

	return
}

// GetFromIsOffline 通过is_offline获取内容
func (obj *_AllVirtualClogStatMgr) GetFromIsOffline(isOffline int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`is_offline` = ?", isOffline).Find(&results).Error

	return
}

// GetBatchFromIsOffline 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromIsOffline(isOfflines []int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`is_offline` IN (?)", isOfflines).Find(&results).Error

	return
}

// GetFromIsInSync 通过is_in_sync获取内容
func (obj *_AllVirtualClogStatMgr) GetFromIsInSync(isInSync int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`is_in_sync` = ?", isInSync).Find(&results).Error

	return
}

// GetBatchFromIsInSync 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromIsInSync(isInSyncs []int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`is_in_sync` IN (?)", isInSyncs).Find(&results).Error

	return
}

// GetFromStartID 通过start_id获取内容
func (obj *_AllVirtualClogStatMgr) GetFromStartID(startID uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`start_id` = ?", startID).Find(&results).Error

	return
}

// GetBatchFromStartID 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromStartID(startIDs []uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`start_id` IN (?)", startIDs).Find(&results).Error

	return
}

// GetFromParent 通过parent获取内容
func (obj *_AllVirtualClogStatMgr) GetFromParent(parent string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`parent` = ?", parent).Find(&results).Error

	return
}

// GetBatchFromParent 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromParent(parents []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`parent` IN (?)", parents).Find(&results).Error

	return
}

// GetFromChildrenList 通过children_list获取内容
func (obj *_AllVirtualClogStatMgr) GetFromChildrenList(childrenList string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`children_list` = ?", childrenList).Find(&results).Error

	return
}

// GetBatchFromChildrenList 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromChildrenList(childrenLists []string) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`children_list` IN (?)", childrenLists).Find(&results).Error

	return
}

// GetFromAccuLogCount 通过accu_log_count获取内容
func (obj *_AllVirtualClogStatMgr) GetFromAccuLogCount(accuLogCount uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`accu_log_count` = ?", accuLogCount).Find(&results).Error

	return
}

// GetBatchFromAccuLogCount 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromAccuLogCount(accuLogCounts []uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`accu_log_count` IN (?)", accuLogCounts).Find(&results).Error

	return
}

// GetFromAccuLogDelay 通过accu_log_delay获取内容
func (obj *_AllVirtualClogStatMgr) GetFromAccuLogDelay(accuLogDelay uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`accu_log_delay` = ?", accuLogDelay).Find(&results).Error

	return
}

// GetBatchFromAccuLogDelay 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromAccuLogDelay(accuLogDelays []uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`accu_log_delay` IN (?)", accuLogDelays).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualClogStatMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromAllowGc 通过allow_gc获取内容
func (obj *_AllVirtualClogStatMgr) GetFromAllowGc(allowGc int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`allow_gc` = ?", allowGc).Find(&results).Error

	return
}

// GetBatchFromAllowGc 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromAllowGc(allowGcs []int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`allow_gc` IN (?)", allowGcs).Find(&results).Error

	return
}

// GetFromQuorum 通过quorum获取内容
func (obj *_AllVirtualClogStatMgr) GetFromQuorum(quorum int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`quorum` = ?", quorum).Find(&results).Error

	return
}

// GetBatchFromQuorum 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromQuorum(quorums []int64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`quorum` IN (?)", quorums).Find(&results).Error

	return
}

// GetFromIsNeedRebuild 通过is_need_rebuild获取内容
func (obj *_AllVirtualClogStatMgr) GetFromIsNeedRebuild(isNeedRebuild int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`is_need_rebuild` = ?", isNeedRebuild).Find(&results).Error

	return
}

// GetBatchFromIsNeedRebuild 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromIsNeedRebuild(isNeedRebuilds []int8) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`is_need_rebuild` IN (?)", isNeedRebuilds).Find(&results).Error

	return
}

// GetFromNextReplayTsDelta 通过next_replay_ts_delta获取内容
func (obj *_AllVirtualClogStatMgr) GetFromNextReplayTsDelta(nextReplayTsDelta uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`next_replay_ts_delta` = ?", nextReplayTsDelta).Find(&results).Error

	return
}

// GetBatchFromNextReplayTsDelta 批量查找
func (obj *_AllVirtualClogStatMgr) GetBatchFromNextReplayTsDelta(nextReplayTsDeltas []uint64) (results []*AllVirtualClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClogStat{}).Where("`next_replay_ts_delta` IN (?)", nextReplayTsDeltas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
