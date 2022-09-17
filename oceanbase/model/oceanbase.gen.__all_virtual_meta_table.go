package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualMetaTableMgr struct {
	*_BaseMgr
}

// AllVirtualMetaTableMgr open func
func AllVirtualMetaTableMgr(db *gorm.DB) *_AllVirtualMetaTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualMetaTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualMetaTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_meta_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualMetaTableMgr) GetTableName() string {
	return "__all_virtual_meta_table"
}

// Reset 重置gorm会话
func (obj *_AllVirtualMetaTableMgr) Reset() *_AllVirtualMetaTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualMetaTableMgr) Get() (result AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualMetaTableMgr) Gets() (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualMetaTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualMetaTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualMetaTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualMetaTableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualMetaTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualMetaTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualMetaTableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualMetaTableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSQLPort sql_port获取
func (obj *_AllVirtualMetaTableMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualMetaTableMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualMetaTableMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithZone zone获取
func (obj *_AllVirtualMetaTableMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRole role获取
func (obj *_AllVirtualMetaTableMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithMemberList member_list获取
func (obj *_AllVirtualMetaTableMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithRowCount row_count获取
func (obj *_AllVirtualMetaTableMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取
func (obj *_AllVirtualMetaTableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取
func (obj *_AllVirtualMetaTableMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithDataChecksum data_checksum获取
func (obj *_AllVirtualMetaTableMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithRowChecksum row_checksum获取
func (obj *_AllVirtualMetaTableMgr) WithRowChecksum(rowChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["row_checksum"] = rowChecksum })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllVirtualMetaTableMgr) WithColumnChecksum(columnChecksum string) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithIsOriginalLeader is_original_leader获取
func (obj *_AllVirtualMetaTableMgr) WithIsOriginalLeader(isOriginalLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_original_leader"] = isOriginalLeader })
}

// WithIsPreviousLeader is_previous_leader获取
func (obj *_AllVirtualMetaTableMgr) WithIsPreviousLeader(isPreviousLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_previous_leader"] = isPreviousLeader })
}

// WithCreateTime create_time获取
func (obj *_AllVirtualMetaTableMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRebuild rebuild获取
func (obj *_AllVirtualMetaTableMgr) WithRebuild(rebuild int64) Option {
	return optionFunc(func(o *options) { o.query["rebuild"] = rebuild })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualMetaTableMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithRequiredSize required_size获取
func (obj *_AllVirtualMetaTableMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithStatus status获取
func (obj *_AllVirtualMetaTableMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsRestore is_restore获取
func (obj *_AllVirtualMetaTableMgr) WithIsRestore(isRestore int64) Option {
	return optionFunc(func(o *options) { o.query["is_restore"] = isRestore })
}

// WithPartitionChecksum partition_checksum获取
func (obj *_AllVirtualMetaTableMgr) WithPartitionChecksum(partitionChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["partition_checksum"] = partitionChecksum })
}

// WithQuorum quorum获取
func (obj *_AllVirtualMetaTableMgr) WithQuorum(quorum int64) Option {
	return optionFunc(func(o *options) { o.query["quorum"] = quorum })
}

// WithFailList fail_list获取
func (obj *_AllVirtualMetaTableMgr) WithFailList(failList string) Option {
	return optionFunc(func(o *options) { o.query["fail_list"] = failList })
}

// WithRecoveryTimestamp recovery_timestamp获取
func (obj *_AllVirtualMetaTableMgr) WithRecoveryTimestamp(recoveryTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["recovery_timestamp"] = recoveryTimestamp })
}

// WithMemstorePercent memstore_percent获取
func (obj *_AllVirtualMetaTableMgr) WithMemstorePercent(memstorePercent int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_percent"] = memstorePercent })
}

// WithDataFileID data_file_id获取
func (obj *_AllVirtualMetaTableMgr) WithDataFileID(dataFileID int64) Option {
	return optionFunc(func(o *options) { o.query["data_file_id"] = dataFileID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualMetaTableMgr) GetByOption(opts ...Option) (result AllVirtualMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualMetaTableMgr) GetByOptions(opts ...Option) (results []*AllVirtualMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualMetaTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualMetaTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where(options.query)
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
func (obj *_AllVirtualMetaTableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromTableID(tableID int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSQLPort 通过sql_port获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromSQLPort(sqlPort int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error

	return
}

// GetBatchFromSQLPort 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromUnitID(unitID int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromZone(zone string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromZone(zones []string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromRole(role int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromMemberList(memberList string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromMemberList(memberLists []string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDataSize 通过data_size获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromDataSize(dataSize int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_size` = ?", dataSize).Find(&results).Error

	return
}

// GetBatchFromDataSize 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromDataVersion(dataVersion int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromRowChecksum 通过row_checksum获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromRowChecksum(rowChecksum int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`row_checksum` = ?", rowChecksum).Find(&results).Error

	return
}

// GetBatchFromRowChecksum 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromRowChecksum(rowChecksums []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`row_checksum` IN (?)", rowChecksums).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromColumnChecksum(columnChecksum string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromColumnChecksum(columnChecksums []string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromIsOriginalLeader 通过is_original_leader获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromIsOriginalLeader(isOriginalLeader int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`is_original_leader` = ?", isOriginalLeader).Find(&results).Error

	return
}

// GetBatchFromIsOriginalLeader 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromIsOriginalLeader(isOriginalLeaders []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`is_original_leader` IN (?)", isOriginalLeaders).Find(&results).Error

	return
}

// GetFromIsPreviousLeader 通过is_previous_leader获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromIsPreviousLeader(isPreviousLeader int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`is_previous_leader` = ?", isPreviousLeader).Find(&results).Error

	return
}

// GetBatchFromIsPreviousLeader 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromIsPreviousLeader(isPreviousLeaders []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`is_previous_leader` IN (?)", isPreviousLeaders).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromCreateTime(createTime int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromCreateTime(createTimes []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRebuild 通过rebuild获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromRebuild(rebuild int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`rebuild` = ?", rebuild).Find(&results).Error

	return
}

// GetBatchFromRebuild 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromRebuild(rebuilds []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`rebuild` IN (?)", rebuilds).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromRequiredSize 通过required_size获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromRequiredSize(requiredSize int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`required_size` = ?", requiredSize).Find(&results).Error

	return
}

// GetBatchFromRequiredSize 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromStatus(status string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsRestore 通过is_restore获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromIsRestore(isRestore int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`is_restore` = ?", isRestore).Find(&results).Error

	return
}

// GetBatchFromIsRestore 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromIsRestore(isRestores []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`is_restore` IN (?)", isRestores).Find(&results).Error

	return
}

// GetFromPartitionChecksum 通过partition_checksum获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromPartitionChecksum(partitionChecksum int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`partition_checksum` = ?", partitionChecksum).Find(&results).Error

	return
}

// GetBatchFromPartitionChecksum 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromPartitionChecksum(partitionChecksums []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`partition_checksum` IN (?)", partitionChecksums).Find(&results).Error

	return
}

// GetFromQuorum 通过quorum获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromQuorum(quorum int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`quorum` = ?", quorum).Find(&results).Error

	return
}

// GetBatchFromQuorum 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromQuorum(quorums []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`quorum` IN (?)", quorums).Find(&results).Error

	return
}

// GetFromFailList 通过fail_list获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromFailList(failList string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`fail_list` = ?", failList).Find(&results).Error

	return
}

// GetBatchFromFailList 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromFailList(failLists []string) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`fail_list` IN (?)", failLists).Find(&results).Error

	return
}

// GetFromRecoveryTimestamp 通过recovery_timestamp获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromRecoveryTimestamp(recoveryTimestamp int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`recovery_timestamp` = ?", recoveryTimestamp).Find(&results).Error

	return
}

// GetBatchFromRecoveryTimestamp 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromRecoveryTimestamp(recoveryTimestamps []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`recovery_timestamp` IN (?)", recoveryTimestamps).Find(&results).Error

	return
}

// GetFromMemstorePercent 通过memstore_percent获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromMemstorePercent(memstorePercent int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`memstore_percent` = ?", memstorePercent).Find(&results).Error

	return
}

// GetBatchFromMemstorePercent 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromMemstorePercent(memstorePercents []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`memstore_percent` IN (?)", memstorePercents).Find(&results).Error

	return
}

// GetFromDataFileID 通过data_file_id获取内容
func (obj *_AllVirtualMetaTableMgr) GetFromDataFileID(dataFileID int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_file_id` = ?", dataFileID).Find(&results).Error

	return
}

// GetBatchFromDataFileID 批量查找
func (obj *_AllVirtualMetaTableMgr) GetBatchFromDataFileID(dataFileIDs []int64) (results []*AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`data_file_id` IN (?)", dataFileIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualMetaTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, svrIP string, svrPort int64) (result AllVirtualMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMetaTable{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
