package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _Gv$sstableMgr struct {
	*_BaseMgr
}

// Gv$sstableMgr open func
func Gv$sstableMgr(db *gorm.DB) *_Gv$sstableMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$sstableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$sstableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$sstable"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$sstableMgr) GetTableName() string {
	return "gv$sstable"
}

// Reset 重置gorm会话
func (obj *_Gv$sstableMgr) Reset() *_Gv$sstableMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$sstableMgr) Get() (result Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$sstableMgr) Gets() (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$sstableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$sstableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$sstableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTableType TABLE_TYPE获取 
func (obj *_Gv$sstableMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_TYPE"] = tableType })
}

// WithTableID TABLE_ID获取 
func (obj *_Gv$sstableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithTableName TABLE_NAME获取 
func (obj *_Gv$sstableMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithTenantID TENANT_ID获取 
func (obj *_Gv$sstableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithPartitionID PARTITION_ID获取 
func (obj *_Gv$sstableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithIndexID INDEX_ID获取 
func (obj *_Gv$sstableMgr) WithIndexID(indexID int64) Option {
	return optionFunc(func(o *options) { o.query["INDEX_ID"] = indexID })
}

// WithBaseVersion BASE_VERSION获取 
func (obj *_Gv$sstableMgr) WithBaseVersion(baseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["BASE_VERSION"] = baseVersion })
}

// WithMultiVersionStart MULTI_VERSION_START获取 
func (obj *_Gv$sstableMgr) WithMultiVersionStart(multiVersionStart int64) Option {
	return optionFunc(func(o *options) { o.query["MULTI_VERSION_START"] = multiVersionStart })
}

// WithSnapshotVersion SNAPSHOT_VERSION获取 
func (obj *_Gv$sstableMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["SNAPSHOT_VERSION"] = snapshotVersion })
}

// WithStartLogTs START_LOG_TS获取 
func (obj *_Gv$sstableMgr) WithStartLogTs(startLogTs uint64) Option {
	return optionFunc(func(o *options) { o.query["START_LOG_TS"] = startLogTs })
}

// WithEndLogTs END_LOG_TS获取 
func (obj *_Gv$sstableMgr) WithEndLogTs(endLogTs uint64) Option {
	return optionFunc(func(o *options) { o.query["END_LOG_TS"] = endLogTs })
}

// WithMaxLogTs MAX_LOG_TS获取 
func (obj *_Gv$sstableMgr) WithMaxLogTs(maxLogTs uint64) Option {
	return optionFunc(func(o *options) { o.query["MAX_LOG_TS"] = maxLogTs })
}

// WithVersion VERSION获取 
func (obj *_Gv$sstableMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["VERSION"] = version })
}

// WithLogicalDataVersion LOGICAL_DATA_VERSION获取 
func (obj *_Gv$sstableMgr) WithLogicalDataVersion(logicalDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["LOGICAL_DATA_VERSION"] = logicalDataVersion })
}

// WithSize SIZE获取 
func (obj *_Gv$sstableMgr) WithSize(size int64) Option {
	return optionFunc(func(o *options) { o.query["SIZE"] = size })
}

// WithIsActive IS_ACTIVE获取 
func (obj *_Gv$sstableMgr) WithIsActive(isActive int64) Option {
	return optionFunc(func(o *options) { o.query["IS_ACTIVE"] = isActive })
}

// WithRef REF获取 
func (obj *_Gv$sstableMgr) WithRef(ref int64) Option {
	return optionFunc(func(o *options) { o.query["REF"] = ref })
}

// WithWriteRef WRITE_REF获取 
func (obj *_Gv$sstableMgr) WithWriteRef(writeRef int64) Option {
	return optionFunc(func(o *options) { o.query["WRITE_REF"] = writeRef })
}

// WithTrxCount TRX_COUNT获取 
func (obj *_Gv$sstableMgr) WithTrxCount(trxCount int64) Option {
	return optionFunc(func(o *options) { o.query["TRX_COUNT"] = trxCount })
}

// WithPendingLogPersistingRowCnt PENDING_LOG_PERSISTING_ROW_CNT获取 
func (obj *_Gv$sstableMgr) WithPendingLogPersistingRowCnt(pendingLogPersistingRowCnt int64) Option {
	return optionFunc(func(o *options) { o.query["PENDING_LOG_PERSISTING_ROW_CNT"] = pendingLogPersistingRowCnt })
}

// WithUpperTransVersion UPPER_TRANS_VERSION获取 
func (obj *_Gv$sstableMgr) WithUpperTransVersion(upperTransVersion int64) Option {
	return optionFunc(func(o *options) { o.query["UPPER_TRANS_VERSION"] = upperTransVersion })
}

// WithContainUncommittedRow CONTAIN_UNCOMMITTED_ROW获取 
func (obj *_Gv$sstableMgr) WithContainUncommittedRow(containUncommittedRow int8) Option {
	return optionFunc(func(o *options) { o.query["CONTAIN_UNCOMMITTED_ROW"] = containUncommittedRow })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$sstableMgr) GetByOption(opts ...Option) (result Gv$sstable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$sstableMgr) GetByOptions(opts ...Option) (results []*Gv$sstable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$sstableMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$sstable,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where(options.query)
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


// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$sstableMgr) GetFromSvrIP(svrIP string) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$sstableMgr) GetFromSvrPort(svrPort int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTableType 通过TABLE_TYPE获取内容  
func (obj *_Gv$sstableMgr) GetFromTableType(tableType int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TABLE_TYPE` = ?", tableType).Find(&results).Error
	
	return
}

// GetBatchFromTableType 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromTableType(tableTypes []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TABLE_TYPE` IN (?)", tableTypes).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_Gv$sstableMgr) GetFromTableID(tableID int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromTableID(tableIDs []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromTableName 通过TABLE_NAME获取内容  
func (obj *_Gv$sstableMgr) GetFromTableName(tableName string) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error
	
	return
}

// GetBatchFromTableName 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromTableName(tableNames []string) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$sstableMgr) GetFromTenantID(tenantID int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过PARTITION_ID获取内容  
func (obj *_Gv$sstableMgr) GetFromPartitionID(partitionID int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromIndexID 通过INDEX_ID获取内容  
func (obj *_Gv$sstableMgr) GetFromIndexID(indexID int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`INDEX_ID` = ?", indexID).Find(&results).Error
	
	return
}

// GetBatchFromIndexID 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromIndexID(indexIDs []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`INDEX_ID` IN (?)", indexIDs).Find(&results).Error
	
	return
}
 
// GetFromBaseVersion 通过BASE_VERSION获取内容  
func (obj *_Gv$sstableMgr) GetFromBaseVersion(baseVersion int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`BASE_VERSION` = ?", baseVersion).Find(&results).Error
	
	return
}

// GetBatchFromBaseVersion 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromBaseVersion(baseVersions []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`BASE_VERSION` IN (?)", baseVersions).Find(&results).Error
	
	return
}
 
// GetFromMultiVersionStart 通过MULTI_VERSION_START获取内容  
func (obj *_Gv$sstableMgr) GetFromMultiVersionStart(multiVersionStart int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`MULTI_VERSION_START` = ?", multiVersionStart).Find(&results).Error
	
	return
}

// GetBatchFromMultiVersionStart 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromMultiVersionStart(multiVersionStarts []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`MULTI_VERSION_START` IN (?)", multiVersionStarts).Find(&results).Error
	
	return
}
 
// GetFromSnapshotVersion 通过SNAPSHOT_VERSION获取内容  
func (obj *_Gv$sstableMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SNAPSHOT_VERSION` = ?", snapshotVersion).Find(&results).Error
	
	return
}

// GetBatchFromSnapshotVersion 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SNAPSHOT_VERSION` IN (?)", snapshotVersions).Find(&results).Error
	
	return
}
 
// GetFromStartLogTs 通过START_LOG_TS获取内容  
func (obj *_Gv$sstableMgr) GetFromStartLogTs(startLogTs uint64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`START_LOG_TS` = ?", startLogTs).Find(&results).Error
	
	return
}

// GetBatchFromStartLogTs 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromStartLogTs(startLogTss []uint64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`START_LOG_TS` IN (?)", startLogTss).Find(&results).Error
	
	return
}
 
// GetFromEndLogTs 通过END_LOG_TS获取内容  
func (obj *_Gv$sstableMgr) GetFromEndLogTs(endLogTs uint64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`END_LOG_TS` = ?", endLogTs).Find(&results).Error
	
	return
}

// GetBatchFromEndLogTs 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromEndLogTs(endLogTss []uint64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`END_LOG_TS` IN (?)", endLogTss).Find(&results).Error
	
	return
}
 
// GetFromMaxLogTs 通过MAX_LOG_TS获取内容  
func (obj *_Gv$sstableMgr) GetFromMaxLogTs(maxLogTs uint64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`MAX_LOG_TS` = ?", maxLogTs).Find(&results).Error
	
	return
}

// GetBatchFromMaxLogTs 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromMaxLogTs(maxLogTss []uint64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`MAX_LOG_TS` IN (?)", maxLogTss).Find(&results).Error
	
	return
}
 
// GetFromVersion 通过VERSION获取内容  
func (obj *_Gv$sstableMgr) GetFromVersion(version int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`VERSION` = ?", version).Find(&results).Error
	
	return
}

// GetBatchFromVersion 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromVersion(versions []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`VERSION` IN (?)", versions).Find(&results).Error
	
	return
}
 
// GetFromLogicalDataVersion 通过LOGICAL_DATA_VERSION获取内容  
func (obj *_Gv$sstableMgr) GetFromLogicalDataVersion(logicalDataVersion int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`LOGICAL_DATA_VERSION` = ?", logicalDataVersion).Find(&results).Error
	
	return
}

// GetBatchFromLogicalDataVersion 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromLogicalDataVersion(logicalDataVersions []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`LOGICAL_DATA_VERSION` IN (?)", logicalDataVersions).Find(&results).Error
	
	return
}
 
// GetFromSize 通过SIZE获取内容  
func (obj *_Gv$sstableMgr) GetFromSize(size int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SIZE` = ?", size).Find(&results).Error
	
	return
}

// GetBatchFromSize 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromSize(sizes []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`SIZE` IN (?)", sizes).Find(&results).Error
	
	return
}
 
// GetFromIsActive 通过IS_ACTIVE获取内容  
func (obj *_Gv$sstableMgr) GetFromIsActive(isActive int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`IS_ACTIVE` = ?", isActive).Find(&results).Error
	
	return
}

// GetBatchFromIsActive 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromIsActive(isActives []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`IS_ACTIVE` IN (?)", isActives).Find(&results).Error
	
	return
}
 
// GetFromRef 通过REF获取内容  
func (obj *_Gv$sstableMgr) GetFromRef(ref int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`REF` = ?", ref).Find(&results).Error
	
	return
}

// GetBatchFromRef 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromRef(refs []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`REF` IN (?)", refs).Find(&results).Error
	
	return
}
 
// GetFromWriteRef 通过WRITE_REF获取内容  
func (obj *_Gv$sstableMgr) GetFromWriteRef(writeRef int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`WRITE_REF` = ?", writeRef).Find(&results).Error
	
	return
}

// GetBatchFromWriteRef 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromWriteRef(writeRefs []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`WRITE_REF` IN (?)", writeRefs).Find(&results).Error
	
	return
}
 
// GetFromTrxCount 通过TRX_COUNT获取内容  
func (obj *_Gv$sstableMgr) GetFromTrxCount(trxCount int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TRX_COUNT` = ?", trxCount).Find(&results).Error
	
	return
}

// GetBatchFromTrxCount 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromTrxCount(trxCounts []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`TRX_COUNT` IN (?)", trxCounts).Find(&results).Error
	
	return
}
 
// GetFromPendingLogPersistingRowCnt 通过PENDING_LOG_PERSISTING_ROW_CNT获取内容  
func (obj *_Gv$sstableMgr) GetFromPendingLogPersistingRowCnt(pendingLogPersistingRowCnt int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`PENDING_LOG_PERSISTING_ROW_CNT` = ?", pendingLogPersistingRowCnt).Find(&results).Error
	
	return
}

// GetBatchFromPendingLogPersistingRowCnt 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromPendingLogPersistingRowCnt(pendingLogPersistingRowCnts []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`PENDING_LOG_PERSISTING_ROW_CNT` IN (?)", pendingLogPersistingRowCnts).Find(&results).Error
	
	return
}
 
// GetFromUpperTransVersion 通过UPPER_TRANS_VERSION获取内容  
func (obj *_Gv$sstableMgr) GetFromUpperTransVersion(upperTransVersion int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`UPPER_TRANS_VERSION` = ?", upperTransVersion).Find(&results).Error
	
	return
}

// GetBatchFromUpperTransVersion 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromUpperTransVersion(upperTransVersions []int64) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`UPPER_TRANS_VERSION` IN (?)", upperTransVersions).Find(&results).Error
	
	return
}
 
// GetFromContainUncommittedRow 通过CONTAIN_UNCOMMITTED_ROW获取内容  
func (obj *_Gv$sstableMgr) GetFromContainUncommittedRow(containUncommittedRow int8) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`CONTAIN_UNCOMMITTED_ROW` = ?", containUncommittedRow).Find(&results).Error
	
	return
}

// GetBatchFromContainUncommittedRow 批量查找 
func (obj *_Gv$sstableMgr) GetBatchFromContainUncommittedRow(containUncommittedRows []int8) (results []*Gv$sstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sstable{}).Where("`CONTAIN_UNCOMMITTED_ROW` IN (?)", containUncommittedRows).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

