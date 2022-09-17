package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPartitionTableMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionTableMgr open func
func AllVirtualPartitionTableMgr(db *gorm.DB) *_AllVirtualPartitionTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionTableMgr) GetTableName() string {
	return "__all_virtual_partition_table"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionTableMgr) Reset() *_AllVirtualPartitionTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionTableMgr) Get() (result AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionTableMgr) Gets() (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionTableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLPort sql_port获取
func (obj *_AllVirtualPartitionTableMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualPartitionTableMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualPartitionTableMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithZone zone获取
func (obj *_AllVirtualPartitionTableMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRole role获取
func (obj *_AllVirtualPartitionTableMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithMemberList member_list获取
func (obj *_AllVirtualPartitionTableMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithRowCount row_count获取
func (obj *_AllVirtualPartitionTableMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取
func (obj *_AllVirtualPartitionTableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取
func (obj *_AllVirtualPartitionTableMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithDataChecksum data_checksum获取
func (obj *_AllVirtualPartitionTableMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithRowChecksum row_checksum获取
func (obj *_AllVirtualPartitionTableMgr) WithRowChecksum(rowChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["row_checksum"] = rowChecksum })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllVirtualPartitionTableMgr) WithColumnChecksum(columnChecksum string) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithIsOriginalLeader is_original_leader获取
func (obj *_AllVirtualPartitionTableMgr) WithIsOriginalLeader(isOriginalLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_original_leader"] = isOriginalLeader })
}

// WithToLeaderTime to_leader_time获取
func (obj *_AllVirtualPartitionTableMgr) WithToLeaderTime(toLeaderTime int64) Option {
	return optionFunc(func(o *options) { o.query["to_leader_time"] = toLeaderTime })
}

// WithCreateTime create_time获取
func (obj *_AllVirtualPartitionTableMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRebuild rebuild获取
func (obj *_AllVirtualPartitionTableMgr) WithRebuild(rebuild int64) Option {
	return optionFunc(func(o *options) { o.query["rebuild"] = rebuild })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualPartitionTableMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithStatus status获取
func (obj *_AllVirtualPartitionTableMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPartitionChecksum partition_checksum获取
func (obj *_AllVirtualPartitionTableMgr) WithPartitionChecksum(partitionChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["partition_checksum"] = partitionChecksum })
}

// WithRequiredSize required_size获取
func (obj *_AllVirtualPartitionTableMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithIsRestore is_restore获取
func (obj *_AllVirtualPartitionTableMgr) WithIsRestore(isRestore int64) Option {
	return optionFunc(func(o *options) { o.query["is_restore"] = isRestore })
}

// WithQuorum quorum获取
func (obj *_AllVirtualPartitionTableMgr) WithQuorum(quorum int64) Option {
	return optionFunc(func(o *options) { o.query["quorum"] = quorum })
}

// WithFailList fail_list获取
func (obj *_AllVirtualPartitionTableMgr) WithFailList(failList string) Option {
	return optionFunc(func(o *options) { o.query["fail_list"] = failList })
}

// WithRecoveryTimestamp recovery_timestamp获取
func (obj *_AllVirtualPartitionTableMgr) WithRecoveryTimestamp(recoveryTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["recovery_timestamp"] = recoveryTimestamp })
}

// WithMemstorePercent memstore_percent获取
func (obj *_AllVirtualPartitionTableMgr) WithMemstorePercent(memstorePercent int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_percent"] = memstorePercent })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionTableMgr) GetByOption(opts ...Option) (result AllVirtualPartitionTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionTableMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where(options.query)
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
func (obj *_AllVirtualPartitionTableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSQLPort 通过sql_port获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromSQLPort(sqlPort int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error

	return
}

// GetBatchFromSQLPort 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromUnitID(unitID int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromZone(zone string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromZone(zones []string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromRole(role int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromMemberList(memberList string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromMemberList(memberLists []string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDataSize 通过data_size获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromDataSize(dataSize int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`data_size` = ?", dataSize).Find(&results).Error

	return
}

// GetBatchFromDataSize 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromDataVersion(dataVersion int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromRowChecksum 通过row_checksum获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromRowChecksum(rowChecksum int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`row_checksum` = ?", rowChecksum).Find(&results).Error

	return
}

// GetBatchFromRowChecksum 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromRowChecksum(rowChecksums []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`row_checksum` IN (?)", rowChecksums).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromColumnChecksum(columnChecksum string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromColumnChecksum(columnChecksums []string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromIsOriginalLeader 通过is_original_leader获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromIsOriginalLeader(isOriginalLeader int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`is_original_leader` = ?", isOriginalLeader).Find(&results).Error

	return
}

// GetBatchFromIsOriginalLeader 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromIsOriginalLeader(isOriginalLeaders []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`is_original_leader` IN (?)", isOriginalLeaders).Find(&results).Error

	return
}

// GetFromToLeaderTime 通过to_leader_time获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromToLeaderTime(toLeaderTime int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`to_leader_time` = ?", toLeaderTime).Find(&results).Error

	return
}

// GetBatchFromToLeaderTime 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromToLeaderTime(toLeaderTimes []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`to_leader_time` IN (?)", toLeaderTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromCreateTime(createTime int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromCreateTime(createTimes []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRebuild 通过rebuild获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromRebuild(rebuild int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`rebuild` = ?", rebuild).Find(&results).Error

	return
}

// GetBatchFromRebuild 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromRebuild(rebuilds []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`rebuild` IN (?)", rebuilds).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromStatus(status string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPartitionChecksum 通过partition_checksum获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromPartitionChecksum(partitionChecksum int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`partition_checksum` = ?", partitionChecksum).Find(&results).Error

	return
}

// GetBatchFromPartitionChecksum 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromPartitionChecksum(partitionChecksums []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`partition_checksum` IN (?)", partitionChecksums).Find(&results).Error

	return
}

// GetFromRequiredSize 通过required_size获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromRequiredSize(requiredSize int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`required_size` = ?", requiredSize).Find(&results).Error

	return
}

// GetBatchFromRequiredSize 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error

	return
}

// GetFromIsRestore 通过is_restore获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromIsRestore(isRestore int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`is_restore` = ?", isRestore).Find(&results).Error

	return
}

// GetBatchFromIsRestore 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromIsRestore(isRestores []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`is_restore` IN (?)", isRestores).Find(&results).Error

	return
}

// GetFromQuorum 通过quorum获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromQuorum(quorum int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`quorum` = ?", quorum).Find(&results).Error

	return
}

// GetBatchFromQuorum 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromQuorum(quorums []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`quorum` IN (?)", quorums).Find(&results).Error

	return
}

// GetFromFailList 通过fail_list获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromFailList(failList string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`fail_list` = ?", failList).Find(&results).Error

	return
}

// GetBatchFromFailList 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromFailList(failLists []string) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`fail_list` IN (?)", failLists).Find(&results).Error

	return
}

// GetFromRecoveryTimestamp 通过recovery_timestamp获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromRecoveryTimestamp(recoveryTimestamp int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`recovery_timestamp` = ?", recoveryTimestamp).Find(&results).Error

	return
}

// GetBatchFromRecoveryTimestamp 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromRecoveryTimestamp(recoveryTimestamps []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`recovery_timestamp` IN (?)", recoveryTimestamps).Find(&results).Error

	return
}

// GetFromMemstorePercent 通过memstore_percent获取内容
func (obj *_AllVirtualPartitionTableMgr) GetFromMemstorePercent(memstorePercent int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`memstore_percent` = ?", memstorePercent).Find(&results).Error

	return
}

// GetBatchFromMemstorePercent 批量查找
func (obj *_AllVirtualPartitionTableMgr) GetBatchFromMemstorePercent(memstorePercents []int64) (results []*AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`memstore_percent` IN (?)", memstorePercents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPartitionTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, svrIP string, svrPort int64) (result AllVirtualPartitionTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionTable{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
