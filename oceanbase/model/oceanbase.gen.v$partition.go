package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$partitionMgr struct {
	*_BaseMgr
}

// V$partitionMgr open func
func V$partitionMgr(db *gorm.DB) *_V$partitionMgr {
	if db == nil {
		panic(fmt.Errorf("V$partitionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$partitionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$partition"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$partitionMgr) GetTableName() string {
	return "v$partition"
}

// Reset 重置gorm会话
func (obj *_V$partitionMgr) Reset() *_V$partitionMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$partitionMgr) Get() (result V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$partitionMgr) Gets() (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$partitionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$partition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_V$partitionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取 
func (obj *_V$partitionMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTablegroupID tablegroup_id获取 
func (obj *_V$partitionMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithPartitionID partition_id获取 
func (obj *_V$partitionMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取 
func (obj *_V$partitionMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_V$partitionMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLPort sql_port获取 
func (obj *_V$partitionMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithUnitID unit_id获取 
func (obj *_V$partitionMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithPartitionCnt partition_cnt获取 
func (obj *_V$partitionMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithZone zone获取 
func (obj *_V$partitionMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRole role获取 
func (obj *_V$partitionMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithMemberList member_list获取 
func (obj *_V$partitionMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithRowCount row_count获取 
func (obj *_V$partitionMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取 
func (obj *_V$partitionMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取 
func (obj *_V$partitionMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithPartitionChecksum partition_checksum获取 
func (obj *_V$partitionMgr) WithPartitionChecksum(partitionChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["partition_checksum"] = partitionChecksum })
}

// WithDataChecksum data_checksum获取 
func (obj *_V$partitionMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithRowChecksum row_checksum获取 
func (obj *_V$partitionMgr) WithRowChecksum(rowChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["row_checksum"] = rowChecksum })
}

// WithColumnChecksum column_checksum获取 
func (obj *_V$partitionMgr) WithColumnChecksum(columnChecksum string) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithRebuild rebuild获取 
func (obj *_V$partitionMgr) WithRebuild(rebuild int64) Option {
	return optionFunc(func(o *options) { o.query["rebuild"] = rebuild })
}

// WithReplicaType replica_type获取 
func (obj *_V$partitionMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithRequiredSize required_size获取 
func (obj *_V$partitionMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithStatus status获取 
func (obj *_V$partitionMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsRestore is_restore获取 
func (obj *_V$partitionMgr) WithIsRestore(isRestore int64) Option {
	return optionFunc(func(o *options) { o.query["is_restore"] = isRestore })
}

// WithQuorum quorum获取 
func (obj *_V$partitionMgr) WithQuorum(quorum int64) Option {
	return optionFunc(func(o *options) { o.query["quorum"] = quorum })
}


// GetByOption 功能选项模式获取
func (obj *_V$partitionMgr) GetByOption(opts ...Option) (result V$partition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$partitionMgr) GetByOptions(opts ...Option) (results []*V$partition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$partitionMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$partition,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where(options.query)
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
func (obj *_V$partitionMgr) GetFromTenantID(tenantID int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$partitionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过table_id获取内容  
func (obj *_V$partitionMgr) GetFromTableID(tableID int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`table_id` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_V$partitionMgr) GetBatchFromTableID(tableIDs []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromTablegroupID 通过tablegroup_id获取内容  
func (obj *_V$partitionMgr) GetFromTablegroupID(tablegroupID int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error
	
	return
}

// GetBatchFromTablegroupID 批量查找 
func (obj *_V$partitionMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过partition_id获取内容  
func (obj *_V$partitionMgr) GetFromPartitionID(partitionID int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`partition_id` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_V$partitionMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_V$partitionMgr) GetFromSvrIP(svrIP string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$partitionMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_V$partitionMgr) GetFromSvrPort(svrPort int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$partitionMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSQLPort 通过sql_port获取内容  
func (obj *_V$partitionMgr) GetFromSQLPort(sqlPort int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error
	
	return
}

// GetBatchFromSQLPort 批量查找 
func (obj *_V$partitionMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error
	
	return
}
 
// GetFromUnitID 通过unit_id获取内容  
func (obj *_V$partitionMgr) GetFromUnitID(unitID int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`unit_id` = ?", unitID).Find(&results).Error
	
	return
}

// GetBatchFromUnitID 批量查找 
func (obj *_V$partitionMgr) GetBatchFromUnitID(unitIDs []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionCnt 通过partition_cnt获取内容  
func (obj *_V$partitionMgr) GetFromPartitionCnt(partitionCnt int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error
	
	return
}

// GetBatchFromPartitionCnt 批量查找 
func (obj *_V$partitionMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error
	
	return
}
 
// GetFromZone 通过zone获取内容  
func (obj *_V$partitionMgr) GetFromZone(zone string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`zone` = ?", zone).Find(&results).Error
	
	return
}

// GetBatchFromZone 批量查找 
func (obj *_V$partitionMgr) GetBatchFromZone(zones []string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`zone` IN (?)", zones).Find(&results).Error
	
	return
}
 
// GetFromRole 通过role获取内容  
func (obj *_V$partitionMgr) GetFromRole(role int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`role` = ?", role).Find(&results).Error
	
	return
}

// GetBatchFromRole 批量查找 
func (obj *_V$partitionMgr) GetBatchFromRole(roles []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`role` IN (?)", roles).Find(&results).Error
	
	return
}
 
// GetFromMemberList 通过member_list获取内容  
func (obj *_V$partitionMgr) GetFromMemberList(memberList string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`member_list` = ?", memberList).Find(&results).Error
	
	return
}

// GetBatchFromMemberList 批量查找 
func (obj *_V$partitionMgr) GetBatchFromMemberList(memberLists []string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error
	
	return
}
 
// GetFromRowCount 通过row_count获取内容  
func (obj *_V$partitionMgr) GetFromRowCount(rowCount int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`row_count` = ?", rowCount).Find(&results).Error
	
	return
}

// GetBatchFromRowCount 批量查找 
func (obj *_V$partitionMgr) GetBatchFromRowCount(rowCounts []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error
	
	return
}
 
// GetFromDataSize 通过data_size获取内容  
func (obj *_V$partitionMgr) GetFromDataSize(dataSize int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`data_size` = ?", dataSize).Find(&results).Error
	
	return
}

// GetBatchFromDataSize 批量查找 
func (obj *_V$partitionMgr) GetBatchFromDataSize(dataSizes []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error
	
	return
}
 
// GetFromDataVersion 通过data_version获取内容  
func (obj *_V$partitionMgr) GetFromDataVersion(dataVersion int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`data_version` = ?", dataVersion).Find(&results).Error
	
	return
}

// GetBatchFromDataVersion 批量查找 
func (obj *_V$partitionMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error
	
	return
}
 
// GetFromPartitionChecksum 通过partition_checksum获取内容  
func (obj *_V$partitionMgr) GetFromPartitionChecksum(partitionChecksum int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`partition_checksum` = ?", partitionChecksum).Find(&results).Error
	
	return
}

// GetBatchFromPartitionChecksum 批量查找 
func (obj *_V$partitionMgr) GetBatchFromPartitionChecksum(partitionChecksums []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`partition_checksum` IN (?)", partitionChecksums).Find(&results).Error
	
	return
}
 
// GetFromDataChecksum 通过data_checksum获取内容  
func (obj *_V$partitionMgr) GetFromDataChecksum(dataChecksum int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error
	
	return
}

// GetBatchFromDataChecksum 批量查找 
func (obj *_V$partitionMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error
	
	return
}
 
// GetFromRowChecksum 通过row_checksum获取内容  
func (obj *_V$partitionMgr) GetFromRowChecksum(rowChecksum int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`row_checksum` = ?", rowChecksum).Find(&results).Error
	
	return
}

// GetBatchFromRowChecksum 批量查找 
func (obj *_V$partitionMgr) GetBatchFromRowChecksum(rowChecksums []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`row_checksum` IN (?)", rowChecksums).Find(&results).Error
	
	return
}
 
// GetFromColumnChecksum 通过column_checksum获取内容  
func (obj *_V$partitionMgr) GetFromColumnChecksum(columnChecksum string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error
	
	return
}

// GetBatchFromColumnChecksum 批量查找 
func (obj *_V$partitionMgr) GetBatchFromColumnChecksum(columnChecksums []string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error
	
	return
}
 
// GetFromRebuild 通过rebuild获取内容  
func (obj *_V$partitionMgr) GetFromRebuild(rebuild int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`rebuild` = ?", rebuild).Find(&results).Error
	
	return
}

// GetBatchFromRebuild 批量查找 
func (obj *_V$partitionMgr) GetBatchFromRebuild(rebuilds []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`rebuild` IN (?)", rebuilds).Find(&results).Error
	
	return
}
 
// GetFromReplicaType 通过replica_type获取内容  
func (obj *_V$partitionMgr) GetFromReplicaType(replicaType int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`replica_type` = ?", replicaType).Find(&results).Error
	
	return
}

// GetBatchFromReplicaType 批量查找 
func (obj *_V$partitionMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error
	
	return
}
 
// GetFromRequiredSize 通过required_size获取内容  
func (obj *_V$partitionMgr) GetFromRequiredSize(requiredSize int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`required_size` = ?", requiredSize).Find(&results).Error
	
	return
}

// GetBatchFromRequiredSize 批量查找 
func (obj *_V$partitionMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容  
func (obj *_V$partitionMgr) GetFromStatus(status string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`status` = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量查找 
func (obj *_V$partitionMgr) GetBatchFromStatus(statuss []string) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`status` IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromIsRestore 通过is_restore获取内容  
func (obj *_V$partitionMgr) GetFromIsRestore(isRestore int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`is_restore` = ?", isRestore).Find(&results).Error
	
	return
}

// GetBatchFromIsRestore 批量查找 
func (obj *_V$partitionMgr) GetBatchFromIsRestore(isRestores []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`is_restore` IN (?)", isRestores).Find(&results).Error
	
	return
}
 
// GetFromQuorum 通过quorum获取内容  
func (obj *_V$partitionMgr) GetFromQuorum(quorum int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`quorum` = ?", quorum).Find(&results).Error
	
	return
}

// GetBatchFromQuorum 批量查找 
func (obj *_V$partitionMgr) GetBatchFromQuorum(quorums []int64) (results []*V$partition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partition{}).Where("`quorum` IN (?)", quorums).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

