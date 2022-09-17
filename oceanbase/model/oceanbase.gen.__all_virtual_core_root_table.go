package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualCoreRootTableMgr struct {
	*_BaseMgr
}

// AllVirtualCoreRootTableMgr open func
func AllVirtualCoreRootTableMgr(db *gorm.DB) *_AllVirtualCoreRootTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualCoreRootTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualCoreRootTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_core_root_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualCoreRootTableMgr) GetTableName() string {
	return "__all_virtual_core_root_table"
}

// Reset 重置gorm会话
func (obj *_AllVirtualCoreRootTableMgr) Reset() *_AllVirtualCoreRootTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualCoreRootTableMgr) Get() (result AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualCoreRootTableMgr) Gets() (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualCoreRootTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualCoreRootTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualCoreRootTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualCoreRootTableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualCoreRootTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualCoreRootTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLPort sql_port获取
func (obj *_AllVirtualCoreRootTableMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualCoreRootTableMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualCoreRootTableMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithZone zone获取
func (obj *_AllVirtualCoreRootTableMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRole role获取
func (obj *_AllVirtualCoreRootTableMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithMemberList member_list获取
func (obj *_AllVirtualCoreRootTableMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithRowCount row_count获取
func (obj *_AllVirtualCoreRootTableMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取
func (obj *_AllVirtualCoreRootTableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取
func (obj *_AllVirtualCoreRootTableMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithDataChecksum data_checksum获取
func (obj *_AllVirtualCoreRootTableMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithRowChecksum row_checksum获取
func (obj *_AllVirtualCoreRootTableMgr) WithRowChecksum(rowChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["row_checksum"] = rowChecksum })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllVirtualCoreRootTableMgr) WithColumnChecksum(columnChecksum string) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithIsOriginalLeader is_original_leader获取
func (obj *_AllVirtualCoreRootTableMgr) WithIsOriginalLeader(isOriginalLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_original_leader"] = isOriginalLeader })
}

// WithIsPreviousLeader is_previous_leader获取
func (obj *_AllVirtualCoreRootTableMgr) WithIsPreviousLeader(isPreviousLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_previous_leader"] = isPreviousLeader })
}

// WithCreateTime create_time获取
func (obj *_AllVirtualCoreRootTableMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRebuild rebuild获取
func (obj *_AllVirtualCoreRootTableMgr) WithRebuild(rebuild int64) Option {
	return optionFunc(func(o *options) { o.query["rebuild"] = rebuild })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualCoreRootTableMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithRequiredSize required_size获取
func (obj *_AllVirtualCoreRootTableMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithStatus status获取
func (obj *_AllVirtualCoreRootTableMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsRestore is_restore获取
func (obj *_AllVirtualCoreRootTableMgr) WithIsRestore(isRestore int64) Option {
	return optionFunc(func(o *options) { o.query["is_restore"] = isRestore })
}

// WithPartitionChecksum partition_checksum获取
func (obj *_AllVirtualCoreRootTableMgr) WithPartitionChecksum(partitionChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["partition_checksum"] = partitionChecksum })
}

// WithQuorum quorum获取
func (obj *_AllVirtualCoreRootTableMgr) WithQuorum(quorum int64) Option {
	return optionFunc(func(o *options) { o.query["quorum"] = quorum })
}

// WithFailList fail_list获取
func (obj *_AllVirtualCoreRootTableMgr) WithFailList(failList string) Option {
	return optionFunc(func(o *options) { o.query["fail_list"] = failList })
}

// WithRecoveryTimestamp recovery_timestamp获取
func (obj *_AllVirtualCoreRootTableMgr) WithRecoveryTimestamp(recoveryTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["recovery_timestamp"] = recoveryTimestamp })
}

// WithMemstorePercent memstore_percent获取
func (obj *_AllVirtualCoreRootTableMgr) WithMemstorePercent(memstorePercent int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_percent"] = memstorePercent })
}

// WithDataFileID data_file_id获取
func (obj *_AllVirtualCoreRootTableMgr) WithDataFileID(dataFileID int64) Option {
	return optionFunc(func(o *options) { o.query["data_file_id"] = dataFileID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualCoreRootTableMgr) GetByOption(opts ...Option) (result AllVirtualCoreRootTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualCoreRootTableMgr) GetByOptions(opts ...Option) (results []*AllVirtualCoreRootTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualCoreRootTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualCoreRootTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where(options.query)
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
func (obj *_AllVirtualCoreRootTableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromTableID(tableID int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSQLPort 通过sql_port获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromSQLPort(sqlPort int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error

	return
}

// GetBatchFromSQLPort 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromUnitID(unitID int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromZone(zone string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromZone(zones []string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromRole(role int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromMemberList(memberList string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromMemberList(memberLists []string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDataSize 通过data_size获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromDataSize(dataSize int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_size` = ?", dataSize).Find(&results).Error

	return
}

// GetBatchFromDataSize 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromDataVersion(dataVersion int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromRowChecksum 通过row_checksum获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromRowChecksum(rowChecksum int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`row_checksum` = ?", rowChecksum).Find(&results).Error

	return
}

// GetBatchFromRowChecksum 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromRowChecksum(rowChecksums []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`row_checksum` IN (?)", rowChecksums).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromColumnChecksum(columnChecksum string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromColumnChecksum(columnChecksums []string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromIsOriginalLeader 通过is_original_leader获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromIsOriginalLeader(isOriginalLeader int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`is_original_leader` = ?", isOriginalLeader).Find(&results).Error

	return
}

// GetBatchFromIsOriginalLeader 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromIsOriginalLeader(isOriginalLeaders []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`is_original_leader` IN (?)", isOriginalLeaders).Find(&results).Error

	return
}

// GetFromIsPreviousLeader 通过is_previous_leader获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromIsPreviousLeader(isPreviousLeader int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`is_previous_leader` = ?", isPreviousLeader).Find(&results).Error

	return
}

// GetBatchFromIsPreviousLeader 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromIsPreviousLeader(isPreviousLeaders []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`is_previous_leader` IN (?)", isPreviousLeaders).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromCreateTime(createTime int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromCreateTime(createTimes []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRebuild 通过rebuild获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromRebuild(rebuild int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`rebuild` = ?", rebuild).Find(&results).Error

	return
}

// GetBatchFromRebuild 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromRebuild(rebuilds []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`rebuild` IN (?)", rebuilds).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromRequiredSize 通过required_size获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromRequiredSize(requiredSize int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`required_size` = ?", requiredSize).Find(&results).Error

	return
}

// GetBatchFromRequiredSize 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromStatus(status string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsRestore 通过is_restore获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromIsRestore(isRestore int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`is_restore` = ?", isRestore).Find(&results).Error

	return
}

// GetBatchFromIsRestore 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromIsRestore(isRestores []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`is_restore` IN (?)", isRestores).Find(&results).Error

	return
}

// GetFromPartitionChecksum 通过partition_checksum获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromPartitionChecksum(partitionChecksum int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`partition_checksum` = ?", partitionChecksum).Find(&results).Error

	return
}

// GetBatchFromPartitionChecksum 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromPartitionChecksum(partitionChecksums []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`partition_checksum` IN (?)", partitionChecksums).Find(&results).Error

	return
}

// GetFromQuorum 通过quorum获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromQuorum(quorum int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`quorum` = ?", quorum).Find(&results).Error

	return
}

// GetBatchFromQuorum 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromQuorum(quorums []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`quorum` IN (?)", quorums).Find(&results).Error

	return
}

// GetFromFailList 通过fail_list获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromFailList(failList string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`fail_list` = ?", failList).Find(&results).Error

	return
}

// GetBatchFromFailList 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromFailList(failLists []string) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`fail_list` IN (?)", failLists).Find(&results).Error

	return
}

// GetFromRecoveryTimestamp 通过recovery_timestamp获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromRecoveryTimestamp(recoveryTimestamp int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`recovery_timestamp` = ?", recoveryTimestamp).Find(&results).Error

	return
}

// GetBatchFromRecoveryTimestamp 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromRecoveryTimestamp(recoveryTimestamps []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`recovery_timestamp` IN (?)", recoveryTimestamps).Find(&results).Error

	return
}

// GetFromMemstorePercent 通过memstore_percent获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromMemstorePercent(memstorePercent int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`memstore_percent` = ?", memstorePercent).Find(&results).Error

	return
}

// GetBatchFromMemstorePercent 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromMemstorePercent(memstorePercents []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`memstore_percent` IN (?)", memstorePercents).Find(&results).Error

	return
}

// GetFromDataFileID 通过data_file_id获取内容
func (obj *_AllVirtualCoreRootTableMgr) GetFromDataFileID(dataFileID int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_file_id` = ?", dataFileID).Find(&results).Error

	return
}

// GetBatchFromDataFileID 批量查找
func (obj *_AllVirtualCoreRootTableMgr) GetBatchFromDataFileID(dataFileIDs []int64) (results []*AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`data_file_id` IN (?)", dataFileIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualCoreRootTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, svrIP string, svrPort int64) (result AllVirtualCoreRootTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreRootTable{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
