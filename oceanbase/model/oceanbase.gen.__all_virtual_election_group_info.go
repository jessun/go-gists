package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualElectionGroupInfoMgr struct {
	*_BaseMgr
}

// AllVirtualElectionGroupInfoMgr open func
func AllVirtualElectionGroupInfoMgr(db *gorm.DB) *_AllVirtualElectionGroupInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualElectionGroupInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualElectionGroupInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_election_group_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualElectionGroupInfoMgr) GetTableName() string {
	return "__all_virtual_election_group_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualElectionGroupInfoMgr) Reset() *_AllVirtualElectionGroupInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualElectionGroupInfoMgr) Get() (result AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualElectionGroupInfoMgr) Gets() (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualElectionGroupInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithEgIDHash eg_id_hash获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithEgIDHash(egIDHash uint64) Option {
	return optionFunc(func(o *options) { o.query["eg_id_hash"] = egIDHash })
}

// WithIsRunning is_running获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithIsRunning(isRunning int8) Option {
	return optionFunc(func(o *options) { o.query["is_running"] = isRunning })
}

// WithEgCreateTime eg_create_time获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithEgCreateTime(egCreateTime int64) Option {
	return optionFunc(func(o *options) { o.query["eg_create_time"] = egCreateTime })
}

// WithEgVersion eg_version获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithEgVersion(egVersion int64) Option {
	return optionFunc(func(o *options) { o.query["eg_version"] = egVersion })
}

// WithEgPartCnt eg_part_cnt获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithEgPartCnt(egPartCnt int64) Option {
	return optionFunc(func(o *options) { o.query["eg_part_cnt"] = egPartCnt })
}

// WithIsAllPartMergedIn is_all_part_merged_in获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithIsAllPartMergedIn(isAllPartMergedIn int8) Option {
	return optionFunc(func(o *options) { o.query["is_all_part_merged_in"] = isAllPartMergedIn })
}

// WithIsPriorityAllowReappoint is_priority_allow_reappoint获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithIsPriorityAllowReappoint(isPriorityAllowReappoint int8) Option {
	return optionFunc(func(o *options) { o.query["is_priority_allow_reappoint"] = isPriorityAllowReappoint })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIsCandidate is_candidate获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithIsCandidate(isCandidate int8) Option {
	return optionFunc(func(o *options) { o.query["is_candidate"] = isCandidate })
}

// WithSystemScore system_score获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithSystemScore(systemScore int64) Option {
	return optionFunc(func(o *options) { o.query["system_score"] = systemScore })
}

// WithCurrentLeader current_leader获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithCurrentLeader(currentLeader string) Option {
	return optionFunc(func(o *options) { o.query["current_leader"] = currentLeader })
}

// WithMemberList member_list获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithTakeoverT1Timestamp takeover_t1_timestamp_获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithTakeoverT1Timestamp(takeoverT1Timestamp int64) Option {
	return optionFunc(func(o *options) { o.query["takeover_t1_timestamp_"] = takeoverT1Timestamp })
}

// WithT1Timestamp T1_timestamp获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithT1Timestamp(t1Timestamp int64) Option {
	return optionFunc(func(o *options) { o.query["T1_timestamp"] = t1Timestamp })
}

// WithLeaseStart lease_start获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithLeaseStart(leaseStart int64) Option {
	return optionFunc(func(o *options) { o.query["lease_start"] = leaseStart })
}

// WithLeaseEnd lease_end获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithLeaseEnd(leaseEnd int64) Option {
	return optionFunc(func(o *options) { o.query["lease_end"] = leaseEnd })
}

// WithRole role获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithState state获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithState(state int64) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithPreDestroyState pre_destroy_state获取
func (obj *_AllVirtualElectionGroupInfoMgr) WithPreDestroyState(preDestroyState int8) Option {
	return optionFunc(func(o *options) { o.query["pre_destroy_state"] = preDestroyState })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualElectionGroupInfoMgr) GetByOption(opts ...Option) (result AllVirtualElectionGroupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualElectionGroupInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualElectionGroupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualElectionGroupInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualElectionGroupInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where(options.query)
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
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromEgIDHash 通过eg_id_hash获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromEgIDHash(egIDHash uint64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_id_hash` = ?", egIDHash).Find(&results).Error

	return
}

// GetBatchFromEgIDHash 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromEgIDHash(egIDHashs []uint64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_id_hash` IN (?)", egIDHashs).Find(&results).Error

	return
}

// GetFromIsRunning 通过is_running获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromIsRunning(isRunning int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_running` = ?", isRunning).Find(&results).Error

	return
}

// GetBatchFromIsRunning 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromIsRunning(isRunnings []int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_running` IN (?)", isRunnings).Find(&results).Error

	return
}

// GetFromEgCreateTime 通过eg_create_time获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromEgCreateTime(egCreateTime int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_create_time` = ?", egCreateTime).Find(&results).Error

	return
}

// GetBatchFromEgCreateTime 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromEgCreateTime(egCreateTimes []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_create_time` IN (?)", egCreateTimes).Find(&results).Error

	return
}

// GetFromEgVersion 通过eg_version获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromEgVersion(egVersion int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_version` = ?", egVersion).Find(&results).Error

	return
}

// GetBatchFromEgVersion 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromEgVersion(egVersions []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_version` IN (?)", egVersions).Find(&results).Error

	return
}

// GetFromEgPartCnt 通过eg_part_cnt获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromEgPartCnt(egPartCnt int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_part_cnt` = ?", egPartCnt).Find(&results).Error

	return
}

// GetBatchFromEgPartCnt 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromEgPartCnt(egPartCnts []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`eg_part_cnt` IN (?)", egPartCnts).Find(&results).Error

	return
}

// GetFromIsAllPartMergedIn 通过is_all_part_merged_in获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromIsAllPartMergedIn(isAllPartMergedIn int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_all_part_merged_in` = ?", isAllPartMergedIn).Find(&results).Error

	return
}

// GetBatchFromIsAllPartMergedIn 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromIsAllPartMergedIn(isAllPartMergedIns []int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_all_part_merged_in` IN (?)", isAllPartMergedIns).Find(&results).Error

	return
}

// GetFromIsPriorityAllowReappoint 通过is_priority_allow_reappoint获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromIsPriorityAllowReappoint(isPriorityAllowReappoint int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_priority_allow_reappoint` = ?", isPriorityAllowReappoint).Find(&results).Error

	return
}

// GetBatchFromIsPriorityAllowReappoint 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromIsPriorityAllowReappoint(isPriorityAllowReappoints []int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_priority_allow_reappoint` IN (?)", isPriorityAllowReappoints).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIsCandidate 通过is_candidate获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromIsCandidate(isCandidate int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_candidate` = ?", isCandidate).Find(&results).Error

	return
}

// GetBatchFromIsCandidate 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromIsCandidate(isCandidates []int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`is_candidate` IN (?)", isCandidates).Find(&results).Error

	return
}

// GetFromSystemScore 通过system_score获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromSystemScore(systemScore int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`system_score` = ?", systemScore).Find(&results).Error

	return
}

// GetBatchFromSystemScore 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromSystemScore(systemScores []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`system_score` IN (?)", systemScores).Find(&results).Error

	return
}

// GetFromCurrentLeader 通过current_leader获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromCurrentLeader(currentLeader string) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`current_leader` = ?", currentLeader).Find(&results).Error

	return
}

// GetBatchFromCurrentLeader 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromCurrentLeader(currentLeaders []string) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`current_leader` IN (?)", currentLeaders).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromMemberList(memberList string) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromMemberList(memberLists []string) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromTakeoverT1Timestamp 通过takeover_t1_timestamp_获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromTakeoverT1Timestamp(takeoverT1Timestamp int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`takeover_t1_timestamp_` = ?", takeoverT1Timestamp).Find(&results).Error

	return
}

// GetBatchFromTakeoverT1Timestamp 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromTakeoverT1Timestamp(takeoverT1Timestamps []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`takeover_t1_timestamp_` IN (?)", takeoverT1Timestamps).Find(&results).Error

	return
}

// GetFromT1Timestamp 通过T1_timestamp获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromT1Timestamp(t1Timestamp int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`T1_timestamp` = ?", t1Timestamp).Find(&results).Error

	return
}

// GetBatchFromT1Timestamp 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromT1Timestamp(t1Timestamps []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`T1_timestamp` IN (?)", t1Timestamps).Find(&results).Error

	return
}

// GetFromLeaseStart 通过lease_start获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromLeaseStart(leaseStart int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`lease_start` = ?", leaseStart).Find(&results).Error

	return
}

// GetBatchFromLeaseStart 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromLeaseStart(leaseStarts []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`lease_start` IN (?)", leaseStarts).Find(&results).Error

	return
}

// GetFromLeaseEnd 通过lease_end获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromLeaseEnd(leaseEnd int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`lease_end` = ?", leaseEnd).Find(&results).Error

	return
}

// GetBatchFromLeaseEnd 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromLeaseEnd(leaseEnds []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`lease_end` IN (?)", leaseEnds).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromRole(role int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromState(state int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromState(states []int64) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromPreDestroyState 通过pre_destroy_state获取内容
func (obj *_AllVirtualElectionGroupInfoMgr) GetFromPreDestroyState(preDestroyState int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`pre_destroy_state` = ?", preDestroyState).Find(&results).Error

	return
}

// GetBatchFromPreDestroyState 批量查找
func (obj *_AllVirtualElectionGroupInfoMgr) GetBatchFromPreDestroyState(preDestroyStates []int8) (results []*AllVirtualElectionGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionGroupInfo{}).Where("`pre_destroy_state` IN (?)", preDestroyStates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
