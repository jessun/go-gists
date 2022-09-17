package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualElectionInfoMgr struct {
	*_BaseMgr
}

// AllVirtualElectionInfoMgr open func
func AllVirtualElectionInfoMgr(db *gorm.DB) *_AllVirtualElectionInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualElectionInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualElectionInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_election_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualElectionInfoMgr) GetTableName() string {
	return "__all_virtual_election_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualElectionInfoMgr) Reset() *_AllVirtualElectionInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualElectionInfoMgr) Get() (result AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualElectionInfoMgr) Gets() (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualElectionInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualElectionInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualElectionInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualElectionInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualElectionInfoMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithIsRunning is_running获取
func (obj *_AllVirtualElectionInfoMgr) WithIsRunning(isRunning int64) Option {
	return optionFunc(func(o *options) { o.query["is_running"] = isRunning })
}

// WithIsChangingLeader is_changing_leader获取
func (obj *_AllVirtualElectionInfoMgr) WithIsChangingLeader(isChangingLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_changing_leader"] = isChangingLeader })
}

// WithCurrentLeader current_leader获取
func (obj *_AllVirtualElectionInfoMgr) WithCurrentLeader(currentLeader string) Option {
	return optionFunc(func(o *options) { o.query["current_leader"] = currentLeader })
}

// WithPreviousLeader previous_leader获取
func (obj *_AllVirtualElectionInfoMgr) WithPreviousLeader(previousLeader string) Option {
	return optionFunc(func(o *options) { o.query["previous_leader"] = previousLeader })
}

// WithProposalLeader proposal_leader获取
func (obj *_AllVirtualElectionInfoMgr) WithProposalLeader(proposalLeader string) Option {
	return optionFunc(func(o *options) { o.query["proposal_leader"] = proposalLeader })
}

// WithMemberList member_list获取
func (obj *_AllVirtualElectionInfoMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualElectionInfoMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithLeaseStart lease_start获取
func (obj *_AllVirtualElectionInfoMgr) WithLeaseStart(leaseStart int64) Option {
	return optionFunc(func(o *options) { o.query["lease_start"] = leaseStart })
}

// WithLeaseEnd lease_end获取
func (obj *_AllVirtualElectionInfoMgr) WithLeaseEnd(leaseEnd int64) Option {
	return optionFunc(func(o *options) { o.query["lease_end"] = leaseEnd })
}

// WithTimeOffset time_offset获取
func (obj *_AllVirtualElectionInfoMgr) WithTimeOffset(timeOffset int64) Option {
	return optionFunc(func(o *options) { o.query["time_offset"] = timeOffset })
}

// WithActiveTimestamp active_timestamp获取
func (obj *_AllVirtualElectionInfoMgr) WithActiveTimestamp(activeTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["active_timestamp"] = activeTimestamp })
}

// WithT1Timestamp T1_timestamp获取
func (obj *_AllVirtualElectionInfoMgr) WithT1Timestamp(t1Timestamp int64) Option {
	return optionFunc(func(o *options) { o.query["T1_timestamp"] = t1Timestamp })
}

// WithLeaderEpoch leader_epoch获取
func (obj *_AllVirtualElectionInfoMgr) WithLeaderEpoch(leaderEpoch int64) Option {
	return optionFunc(func(o *options) { o.query["leader_epoch"] = leaderEpoch })
}

// WithState state获取
func (obj *_AllVirtualElectionInfoMgr) WithState(state int64) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithRole role获取
func (obj *_AllVirtualElectionInfoMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithStage stage获取
func (obj *_AllVirtualElectionInfoMgr) WithStage(stage int64) Option {
	return optionFunc(func(o *options) { o.query["stage"] = stage })
}

// WithEgID eg_id获取
func (obj *_AllVirtualElectionInfoMgr) WithEgID(egID uint64) Option {
	return optionFunc(func(o *options) { o.query["eg_id"] = egID })
}

// WithRemainingTimeInBlacklist remaining_time_in_blacklist获取
func (obj *_AllVirtualElectionInfoMgr) WithRemainingTimeInBlacklist(remainingTimeInBlacklist int64) Option {
	return optionFunc(func(o *options) { o.query["remaining_time_in_blacklist"] = remainingTimeInBlacklist })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualElectionInfoMgr) GetByOption(opts ...Option) (result AllVirtualElectionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualElectionInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualElectionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualElectionInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualElectionInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where(options.query)
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
func (obj *_AllVirtualElectionInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromIsRunning 通过is_running获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromIsRunning(isRunning int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`is_running` = ?", isRunning).Find(&results).Error

	return
}

// GetBatchFromIsRunning 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromIsRunning(isRunnings []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`is_running` IN (?)", isRunnings).Find(&results).Error

	return
}

// GetFromIsChangingLeader 通过is_changing_leader获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromIsChangingLeader(isChangingLeader int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`is_changing_leader` = ?", isChangingLeader).Find(&results).Error

	return
}

// GetBatchFromIsChangingLeader 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromIsChangingLeader(isChangingLeaders []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`is_changing_leader` IN (?)", isChangingLeaders).Find(&results).Error

	return
}

// GetFromCurrentLeader 通过current_leader获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromCurrentLeader(currentLeader string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`current_leader` = ?", currentLeader).Find(&results).Error

	return
}

// GetBatchFromCurrentLeader 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromCurrentLeader(currentLeaders []string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`current_leader` IN (?)", currentLeaders).Find(&results).Error

	return
}

// GetFromPreviousLeader 通过previous_leader获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromPreviousLeader(previousLeader string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`previous_leader` = ?", previousLeader).Find(&results).Error

	return
}

// GetBatchFromPreviousLeader 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromPreviousLeader(previousLeaders []string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`previous_leader` IN (?)", previousLeaders).Find(&results).Error

	return
}

// GetFromProposalLeader 通过proposal_leader获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromProposalLeader(proposalLeader string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`proposal_leader` = ?", proposalLeader).Find(&results).Error

	return
}

// GetBatchFromProposalLeader 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromProposalLeader(proposalLeaders []string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`proposal_leader` IN (?)", proposalLeaders).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromMemberList(memberList string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromMemberList(memberLists []string) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromLeaseStart 通过lease_start获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromLeaseStart(leaseStart int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`lease_start` = ?", leaseStart).Find(&results).Error

	return
}

// GetBatchFromLeaseStart 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromLeaseStart(leaseStarts []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`lease_start` IN (?)", leaseStarts).Find(&results).Error

	return
}

// GetFromLeaseEnd 通过lease_end获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromLeaseEnd(leaseEnd int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`lease_end` = ?", leaseEnd).Find(&results).Error

	return
}

// GetBatchFromLeaseEnd 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromLeaseEnd(leaseEnds []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`lease_end` IN (?)", leaseEnds).Find(&results).Error

	return
}

// GetFromTimeOffset 通过time_offset获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromTimeOffset(timeOffset int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`time_offset` = ?", timeOffset).Find(&results).Error

	return
}

// GetBatchFromTimeOffset 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromTimeOffset(timeOffsets []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`time_offset` IN (?)", timeOffsets).Find(&results).Error

	return
}

// GetFromActiveTimestamp 通过active_timestamp获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromActiveTimestamp(activeTimestamp int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`active_timestamp` = ?", activeTimestamp).Find(&results).Error

	return
}

// GetBatchFromActiveTimestamp 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromActiveTimestamp(activeTimestamps []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`active_timestamp` IN (?)", activeTimestamps).Find(&results).Error

	return
}

// GetFromT1Timestamp 通过T1_timestamp获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromT1Timestamp(t1Timestamp int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`T1_timestamp` = ?", t1Timestamp).Find(&results).Error

	return
}

// GetBatchFromT1Timestamp 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromT1Timestamp(t1Timestamps []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`T1_timestamp` IN (?)", t1Timestamps).Find(&results).Error

	return
}

// GetFromLeaderEpoch 通过leader_epoch获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromLeaderEpoch(leaderEpoch int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`leader_epoch` = ?", leaderEpoch).Find(&results).Error

	return
}

// GetBatchFromLeaderEpoch 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromLeaderEpoch(leaderEpochs []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`leader_epoch` IN (?)", leaderEpochs).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromState(state int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromState(states []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromRole(role int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromStage 通过stage获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromStage(stage int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`stage` = ?", stage).Find(&results).Error

	return
}

// GetBatchFromStage 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromStage(stages []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`stage` IN (?)", stages).Find(&results).Error

	return
}

// GetFromEgID 通过eg_id获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromEgID(egID uint64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`eg_id` = ?", egID).Find(&results).Error

	return
}

// GetBatchFromEgID 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromEgID(egIDs []uint64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`eg_id` IN (?)", egIDs).Find(&results).Error

	return
}

// GetFromRemainingTimeInBlacklist 通过remaining_time_in_blacklist获取内容
func (obj *_AllVirtualElectionInfoMgr) GetFromRemainingTimeInBlacklist(remainingTimeInBlacklist int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`remaining_time_in_blacklist` = ?", remainingTimeInBlacklist).Find(&results).Error

	return
}

// GetBatchFromRemainingTimeInBlacklist 批量查找
func (obj *_AllVirtualElectionInfoMgr) GetBatchFromRemainingTimeInBlacklist(remainingTimeInBlacklists []int64) (results []*AllVirtualElectionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionInfo{}).Where("`remaining_time_in_blacklist` IN (?)", remainingTimeInBlacklists).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
