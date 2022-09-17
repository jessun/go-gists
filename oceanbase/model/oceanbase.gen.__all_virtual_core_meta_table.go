package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualCoreMetaTableMgr struct {
	*_BaseMgr
}

// AllVirtualCoreMetaTableMgr open func
func AllVirtualCoreMetaTableMgr(db *gorm.DB) *_AllVirtualCoreMetaTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualCoreMetaTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualCoreMetaTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_core_meta_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualCoreMetaTableMgr) GetTableName() string {
	return "__all_virtual_core_meta_table"
}

// Reset 重置gorm会话
func (obj *_AllVirtualCoreMetaTableMgr) Reset() *_AllVirtualCoreMetaTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualCoreMetaTableMgr) Get() (result AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualCoreMetaTableMgr) Gets() (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualCoreMetaTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualCoreMetaTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualCoreMetaTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualCoreMetaTableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualCoreMetaTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualCoreMetaTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLPort sql_port获取
func (obj *_AllVirtualCoreMetaTableMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualCoreMetaTableMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualCoreMetaTableMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithZone zone获取
func (obj *_AllVirtualCoreMetaTableMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRole role获取
func (obj *_AllVirtualCoreMetaTableMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithMemberList member_list获取
func (obj *_AllVirtualCoreMetaTableMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithRowCount row_count获取
func (obj *_AllVirtualCoreMetaTableMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取
func (obj *_AllVirtualCoreMetaTableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取
func (obj *_AllVirtualCoreMetaTableMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithDataChecksum data_checksum获取
func (obj *_AllVirtualCoreMetaTableMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithRowChecksum row_checksum获取
func (obj *_AllVirtualCoreMetaTableMgr) WithRowChecksum(rowChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["row_checksum"] = rowChecksum })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllVirtualCoreMetaTableMgr) WithColumnChecksum(columnChecksum string) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithIsOriginalLeader is_original_leader获取
func (obj *_AllVirtualCoreMetaTableMgr) WithIsOriginalLeader(isOriginalLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_original_leader"] = isOriginalLeader })
}

// WithIsPreviousLeader is_previous_leader获取
func (obj *_AllVirtualCoreMetaTableMgr) WithIsPreviousLeader(isPreviousLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_previous_leader"] = isPreviousLeader })
}

// WithCreateTime create_time获取
func (obj *_AllVirtualCoreMetaTableMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRebuild rebuild获取
func (obj *_AllVirtualCoreMetaTableMgr) WithRebuild(rebuild int64) Option {
	return optionFunc(func(o *options) { o.query["rebuild"] = rebuild })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualCoreMetaTableMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithRequiredSize required_size获取
func (obj *_AllVirtualCoreMetaTableMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithStatus status获取
func (obj *_AllVirtualCoreMetaTableMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsRestore is_restore获取
func (obj *_AllVirtualCoreMetaTableMgr) WithIsRestore(isRestore int64) Option {
	return optionFunc(func(o *options) { o.query["is_restore"] = isRestore })
}

// WithPartitionChecksum partition_checksum获取
func (obj *_AllVirtualCoreMetaTableMgr) WithPartitionChecksum(partitionChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["partition_checksum"] = partitionChecksum })
}

// WithQuorum quorum获取
func (obj *_AllVirtualCoreMetaTableMgr) WithQuorum(quorum int64) Option {
	return optionFunc(func(o *options) { o.query["quorum"] = quorum })
}

// WithFailList fail_list获取
func (obj *_AllVirtualCoreMetaTableMgr) WithFailList(failList string) Option {
	return optionFunc(func(o *options) { o.query["fail_list"] = failList })
}

// WithRecoveryTimestamp recovery_timestamp获取
func (obj *_AllVirtualCoreMetaTableMgr) WithRecoveryTimestamp(recoveryTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["recovery_timestamp"] = recoveryTimestamp })
}

// WithMemstorePercent memstore_percent获取
func (obj *_AllVirtualCoreMetaTableMgr) WithMemstorePercent(memstorePercent int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_percent"] = memstorePercent })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualCoreMetaTableMgr) GetByOption(opts ...Option) (result AllVirtualCoreMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualCoreMetaTableMgr) GetByOptions(opts ...Option) (results []*AllVirtualCoreMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualCoreMetaTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualCoreMetaTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where(options.query)
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
func (obj *_AllVirtualCoreMetaTableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromTableID(tableID int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSQLPort 通过sql_port获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromSQLPort(sqlPort int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error

	return
}

// GetBatchFromSQLPort 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromUnitID(unitID int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromZone(zone string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromZone(zones []string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromRole(role int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromMemberList(memberList string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromMemberList(memberLists []string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDataSize 通过data_size获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromDataSize(dataSize int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`data_size` = ?", dataSize).Find(&results).Error

	return
}

// GetBatchFromDataSize 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromDataVersion(dataVersion int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromRowChecksum 通过row_checksum获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromRowChecksum(rowChecksum int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`row_checksum` = ?", rowChecksum).Find(&results).Error

	return
}

// GetBatchFromRowChecksum 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromRowChecksum(rowChecksums []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`row_checksum` IN (?)", rowChecksums).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromColumnChecksum(columnChecksum string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromColumnChecksum(columnChecksums []string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromIsOriginalLeader 通过is_original_leader获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromIsOriginalLeader(isOriginalLeader int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`is_original_leader` = ?", isOriginalLeader).Find(&results).Error

	return
}

// GetBatchFromIsOriginalLeader 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromIsOriginalLeader(isOriginalLeaders []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`is_original_leader` IN (?)", isOriginalLeaders).Find(&results).Error

	return
}

// GetFromIsPreviousLeader 通过is_previous_leader获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromIsPreviousLeader(isPreviousLeader int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`is_previous_leader` = ?", isPreviousLeader).Find(&results).Error

	return
}

// GetBatchFromIsPreviousLeader 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromIsPreviousLeader(isPreviousLeaders []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`is_previous_leader` IN (?)", isPreviousLeaders).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromCreateTime(createTime int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromCreateTime(createTimes []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRebuild 通过rebuild获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromRebuild(rebuild int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`rebuild` = ?", rebuild).Find(&results).Error

	return
}

// GetBatchFromRebuild 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromRebuild(rebuilds []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`rebuild` IN (?)", rebuilds).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromRequiredSize 通过required_size获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromRequiredSize(requiredSize int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`required_size` = ?", requiredSize).Find(&results).Error

	return
}

// GetBatchFromRequiredSize 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromStatus(status string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsRestore 通过is_restore获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromIsRestore(isRestore int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`is_restore` = ?", isRestore).Find(&results).Error

	return
}

// GetBatchFromIsRestore 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromIsRestore(isRestores []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`is_restore` IN (?)", isRestores).Find(&results).Error

	return
}

// GetFromPartitionChecksum 通过partition_checksum获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromPartitionChecksum(partitionChecksum int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`partition_checksum` = ?", partitionChecksum).Find(&results).Error

	return
}

// GetBatchFromPartitionChecksum 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromPartitionChecksum(partitionChecksums []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`partition_checksum` IN (?)", partitionChecksums).Find(&results).Error

	return
}

// GetFromQuorum 通过quorum获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromQuorum(quorum int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`quorum` = ?", quorum).Find(&results).Error

	return
}

// GetBatchFromQuorum 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromQuorum(quorums []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`quorum` IN (?)", quorums).Find(&results).Error

	return
}

// GetFromFailList 通过fail_list获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromFailList(failList string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`fail_list` = ?", failList).Find(&results).Error

	return
}

// GetBatchFromFailList 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromFailList(failLists []string) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`fail_list` IN (?)", failLists).Find(&results).Error

	return
}

// GetFromRecoveryTimestamp 通过recovery_timestamp获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromRecoveryTimestamp(recoveryTimestamp int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`recovery_timestamp` = ?", recoveryTimestamp).Find(&results).Error

	return
}

// GetBatchFromRecoveryTimestamp 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromRecoveryTimestamp(recoveryTimestamps []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`recovery_timestamp` IN (?)", recoveryTimestamps).Find(&results).Error

	return
}

// GetFromMemstorePercent 通过memstore_percent获取内容
func (obj *_AllVirtualCoreMetaTableMgr) GetFromMemstorePercent(memstorePercent int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`memstore_percent` = ?", memstorePercent).Find(&results).Error

	return
}

// GetBatchFromMemstorePercent 批量查找
func (obj *_AllVirtualCoreMetaTableMgr) GetBatchFromMemstorePercent(memstorePercents []int64) (results []*AllVirtualCoreMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreMetaTable{}).Where("`memstore_percent` IN (?)", memstorePercents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
