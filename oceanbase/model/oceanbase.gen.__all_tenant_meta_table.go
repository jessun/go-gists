package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantMetaTableMgr struct {
	*_BaseMgr
}

// AllTenantMetaTableMgr open func
func AllTenantMetaTableMgr(db *gorm.DB) *_AllTenantMetaTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantMetaTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantMetaTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_meta_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantMetaTableMgr) GetTableName() string {
	return "__all_tenant_meta_table"
}

// Reset 重置gorm会话
func (obj *_AllTenantMetaTableMgr) Reset() *_AllTenantMetaTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantMetaTableMgr) Get() (result AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantMetaTableMgr) Gets() (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantMetaTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantMetaTableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantMetaTableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantMetaTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllTenantMetaTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllTenantMetaTableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllTenantMetaTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllTenantMetaTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLPort sql_port获取
func (obj *_AllTenantMetaTableMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithUnitID unit_id获取
func (obj *_AllTenantMetaTableMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllTenantMetaTableMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithZone zone获取
func (obj *_AllTenantMetaTableMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRole role获取
func (obj *_AllTenantMetaTableMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithMemberList member_list获取
func (obj *_AllTenantMetaTableMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithRowCount row_count获取
func (obj *_AllTenantMetaTableMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取
func (obj *_AllTenantMetaTableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取
func (obj *_AllTenantMetaTableMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithDataChecksum data_checksum获取
func (obj *_AllTenantMetaTableMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithRowChecksum row_checksum获取
func (obj *_AllTenantMetaTableMgr) WithRowChecksum(rowChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["row_checksum"] = rowChecksum })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllTenantMetaTableMgr) WithColumnChecksum(columnChecksum string) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithIsOriginalLeader is_original_leader获取
func (obj *_AllTenantMetaTableMgr) WithIsOriginalLeader(isOriginalLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_original_leader"] = isOriginalLeader })
}

// WithIsPreviousLeader is_previous_leader获取
func (obj *_AllTenantMetaTableMgr) WithIsPreviousLeader(isPreviousLeader int64) Option {
	return optionFunc(func(o *options) { o.query["is_previous_leader"] = isPreviousLeader })
}

// WithCreateTime create_time获取
func (obj *_AllTenantMetaTableMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRebuild rebuild获取
func (obj *_AllTenantMetaTableMgr) WithRebuild(rebuild int64) Option {
	return optionFunc(func(o *options) { o.query["rebuild"] = rebuild })
}

// WithReplicaType replica_type获取
func (obj *_AllTenantMetaTableMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithRequiredSize required_size获取
func (obj *_AllTenantMetaTableMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithStatus status获取
func (obj *_AllTenantMetaTableMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsRestore is_restore获取
func (obj *_AllTenantMetaTableMgr) WithIsRestore(isRestore int64) Option {
	return optionFunc(func(o *options) { o.query["is_restore"] = isRestore })
}

// WithPartitionChecksum partition_checksum获取
func (obj *_AllTenantMetaTableMgr) WithPartitionChecksum(partitionChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["partition_checksum"] = partitionChecksum })
}

// WithQuorum quorum获取
func (obj *_AllTenantMetaTableMgr) WithQuorum(quorum int64) Option {
	return optionFunc(func(o *options) { o.query["quorum"] = quorum })
}

// WithFailList fail_list获取
func (obj *_AllTenantMetaTableMgr) WithFailList(failList string) Option {
	return optionFunc(func(o *options) { o.query["fail_list"] = failList })
}

// WithRecoveryTimestamp recovery_timestamp获取
func (obj *_AllTenantMetaTableMgr) WithRecoveryTimestamp(recoveryTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["recovery_timestamp"] = recoveryTimestamp })
}

// WithMemstorePercent memstore_percent获取
func (obj *_AllTenantMetaTableMgr) WithMemstorePercent(memstorePercent int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_percent"] = memstorePercent })
}

// WithDataFileID data_file_id获取
func (obj *_AllTenantMetaTableMgr) WithDataFileID(dataFileID int64) Option {
	return optionFunc(func(o *options) { o.query["data_file_id"] = dataFileID })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantMetaTableMgr) GetByOption(opts ...Option) (result AllTenantMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantMetaTableMgr) GetByOptions(opts ...Option) (results []*AllTenantMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantMetaTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantMetaTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllTenantMetaTableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantMetaTableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantMetaTableMgr) GetFromTenantID(tenantID int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTenantMetaTableMgr) GetFromTableID(tableID int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllTenantMetaTableMgr) GetFromPartitionID(partitionID int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllTenantMetaTableMgr) GetFromSvrIP(svrIP string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllTenantMetaTableMgr) GetFromSvrPort(svrPort int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSQLPort 通过sql_port获取内容
func (obj *_AllTenantMetaTableMgr) GetFromSQLPort(sqlPort int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error

	return
}

// GetBatchFromSQLPort 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllTenantMetaTableMgr) GetFromUnitID(unitID int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllTenantMetaTableMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllTenantMetaTableMgr) GetFromZone(zone string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromZone(zones []string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllTenantMetaTableMgr) GetFromRole(role int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromRole(roles []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllTenantMetaTableMgr) GetFromMemberList(memberList string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromMemberList(memberLists []string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllTenantMetaTableMgr) GetFromRowCount(rowCount int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDataSize 通过data_size获取内容
func (obj *_AllTenantMetaTableMgr) GetFromDataSize(dataSize int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_size` = ?", dataSize).Find(&results).Error

	return
}

// GetBatchFromDataSize 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllTenantMetaTableMgr) GetFromDataVersion(dataVersion int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllTenantMetaTableMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromRowChecksum 通过row_checksum获取内容
func (obj *_AllTenantMetaTableMgr) GetFromRowChecksum(rowChecksum int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`row_checksum` = ?", rowChecksum).Find(&results).Error

	return
}

// GetBatchFromRowChecksum 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromRowChecksum(rowChecksums []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`row_checksum` IN (?)", rowChecksums).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllTenantMetaTableMgr) GetFromColumnChecksum(columnChecksum string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromColumnChecksum(columnChecksums []string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromIsOriginalLeader 通过is_original_leader获取内容
func (obj *_AllTenantMetaTableMgr) GetFromIsOriginalLeader(isOriginalLeader int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`is_original_leader` = ?", isOriginalLeader).Find(&results).Error

	return
}

// GetBatchFromIsOriginalLeader 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromIsOriginalLeader(isOriginalLeaders []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`is_original_leader` IN (?)", isOriginalLeaders).Find(&results).Error

	return
}

// GetFromIsPreviousLeader 通过is_previous_leader获取内容
func (obj *_AllTenantMetaTableMgr) GetFromIsPreviousLeader(isPreviousLeader int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`is_previous_leader` = ?", isPreviousLeader).Find(&results).Error

	return
}

// GetBatchFromIsPreviousLeader 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromIsPreviousLeader(isPreviousLeaders []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`is_previous_leader` IN (?)", isPreviousLeaders).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_AllTenantMetaTableMgr) GetFromCreateTime(createTime int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromCreateTime(createTimes []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRebuild 通过rebuild获取内容
func (obj *_AllTenantMetaTableMgr) GetFromRebuild(rebuild int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`rebuild` = ?", rebuild).Find(&results).Error

	return
}

// GetBatchFromRebuild 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromRebuild(rebuilds []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`rebuild` IN (?)", rebuilds).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllTenantMetaTableMgr) GetFromReplicaType(replicaType int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromRequiredSize 通过required_size获取内容
func (obj *_AllTenantMetaTableMgr) GetFromRequiredSize(requiredSize int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`required_size` = ?", requiredSize).Find(&results).Error

	return
}

// GetBatchFromRequiredSize 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantMetaTableMgr) GetFromStatus(status string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsRestore 通过is_restore获取内容
func (obj *_AllTenantMetaTableMgr) GetFromIsRestore(isRestore int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`is_restore` = ?", isRestore).Find(&results).Error

	return
}

// GetBatchFromIsRestore 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromIsRestore(isRestores []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`is_restore` IN (?)", isRestores).Find(&results).Error

	return
}

// GetFromPartitionChecksum 通过partition_checksum获取内容
func (obj *_AllTenantMetaTableMgr) GetFromPartitionChecksum(partitionChecksum int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`partition_checksum` = ?", partitionChecksum).Find(&results).Error

	return
}

// GetBatchFromPartitionChecksum 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromPartitionChecksum(partitionChecksums []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`partition_checksum` IN (?)", partitionChecksums).Find(&results).Error

	return
}

// GetFromQuorum 通过quorum获取内容
func (obj *_AllTenantMetaTableMgr) GetFromQuorum(quorum int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`quorum` = ?", quorum).Find(&results).Error

	return
}

// GetBatchFromQuorum 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromQuorum(quorums []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`quorum` IN (?)", quorums).Find(&results).Error

	return
}

// GetFromFailList 通过fail_list获取内容
func (obj *_AllTenantMetaTableMgr) GetFromFailList(failList string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`fail_list` = ?", failList).Find(&results).Error

	return
}

// GetBatchFromFailList 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromFailList(failLists []string) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`fail_list` IN (?)", failLists).Find(&results).Error

	return
}

// GetFromRecoveryTimestamp 通过recovery_timestamp获取内容
func (obj *_AllTenantMetaTableMgr) GetFromRecoveryTimestamp(recoveryTimestamp int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`recovery_timestamp` = ?", recoveryTimestamp).Find(&results).Error

	return
}

// GetBatchFromRecoveryTimestamp 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromRecoveryTimestamp(recoveryTimestamps []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`recovery_timestamp` IN (?)", recoveryTimestamps).Find(&results).Error

	return
}

// GetFromMemstorePercent 通过memstore_percent获取内容
func (obj *_AllTenantMetaTableMgr) GetFromMemstorePercent(memstorePercent int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`memstore_percent` = ?", memstorePercent).Find(&results).Error

	return
}

// GetBatchFromMemstorePercent 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromMemstorePercent(memstorePercents []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`memstore_percent` IN (?)", memstorePercents).Find(&results).Error

	return
}

// GetFromDataFileID 通过data_file_id获取内容
func (obj *_AllTenantMetaTableMgr) GetFromDataFileID(dataFileID int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_file_id` = ?", dataFileID).Find(&results).Error

	return
}

// GetBatchFromDataFileID 批量查找
func (obj *_AllTenantMetaTableMgr) GetBatchFromDataFileID(dataFileIDs []int64) (results []*AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`data_file_id` IN (?)", dataFileIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantMetaTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, svrIP string, svrPort int64) (result AllTenantMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantMetaTable{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
