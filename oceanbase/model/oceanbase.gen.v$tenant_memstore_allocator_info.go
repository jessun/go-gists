package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$tenantMemstoreAllocatorInfoMgr struct {
	*_BaseMgr
}

// V$tenantMemstoreAllocatorInfoMgr open func
func V$tenantMemstoreAllocatorInfoMgr(db *gorm.DB) *_V$tenantMemstoreAllocatorInfoMgr {
	if db == nil {
		panic(fmt.Errorf("V$tenantMemstoreAllocatorInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$tenantMemstoreAllocatorInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$tenant_memstore_allocator_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetTableName() string {
	return "v$tenant_memstore_allocator_info"
}

// Reset 重置gorm会话
func (obj *_V$tenantMemstoreAllocatorInfoMgr) Reset() *_V$tenantMemstoreAllocatorInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) Get() (result V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$tenantMemstoreAllocatorInfoMgr) Gets() (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$tenantMemstoreAllocatorInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTableID TABLE_ID获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithMtBaseVersion MT_BASE_VERSION获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithMtBaseVersion(mtBaseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["MT_BASE_VERSION"] = mtBaseVersion })
}

// WithRetireClock RETIRE_CLOCK获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithRetireClock(retireClock int64) Option {
	return optionFunc(func(o *options) { o.query["RETIRE_CLOCK"] = retireClock })
}

// WithMtIsFrozen MT_IS_FROZEN获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithMtIsFrozen(mtIsFrozen int64) Option {
	return optionFunc(func(o *options) { o.query["MT_IS_FROZEN"] = mtIsFrozen })
}

// WithMtProtectionClock MT_PROTECTION_CLOCK获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithMtProtectionClock(mtProtectionClock int64) Option {
	return optionFunc(func(o *options) { o.query["MT_PROTECTION_CLOCK"] = mtProtectionClock })
}

// WithMtSnapshotVersion MT_SNAPSHOT_VERSION获取 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) WithMtSnapshotVersion(mtSnapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["MT_SNAPSHOT_VERSION"] = mtSnapshotVersion })
}


// GetByOption 功能选项模式获取
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetByOption(opts ...Option) (result V$tenantMemstoreAllocatorInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetByOptions(opts ...Option) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$tenantMemstoreAllocatorInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$tenantMemstoreAllocatorInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where(options.query)
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
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromSvrIP(svrIP string) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromSvrPort(svrPort int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromTenantID(tenantID int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromTableID(tableID int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过PARTITION_ID获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromPartitionID(partitionID int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromMtBaseVersion 通过MT_BASE_VERSION获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromMtBaseVersion(mtBaseVersion int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_BASE_VERSION` = ?", mtBaseVersion).Find(&results).Error
	
	return
}

// GetBatchFromMtBaseVersion 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromMtBaseVersion(mtBaseVersions []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_BASE_VERSION` IN (?)", mtBaseVersions).Find(&results).Error
	
	return
}
 
// GetFromRetireClock 通过RETIRE_CLOCK获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromRetireClock(retireClock int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`RETIRE_CLOCK` = ?", retireClock).Find(&results).Error
	
	return
}

// GetBatchFromRetireClock 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromRetireClock(retireClocks []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`RETIRE_CLOCK` IN (?)", retireClocks).Find(&results).Error
	
	return
}
 
// GetFromMtIsFrozen 通过MT_IS_FROZEN获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromMtIsFrozen(mtIsFrozen int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_IS_FROZEN` = ?", mtIsFrozen).Find(&results).Error
	
	return
}

// GetBatchFromMtIsFrozen 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromMtIsFrozen(mtIsFrozens []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_IS_FROZEN` IN (?)", mtIsFrozens).Find(&results).Error
	
	return
}
 
// GetFromMtProtectionClock 通过MT_PROTECTION_CLOCK获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromMtProtectionClock(mtProtectionClock int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_PROTECTION_CLOCK` = ?", mtProtectionClock).Find(&results).Error
	
	return
}

// GetBatchFromMtProtectionClock 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromMtProtectionClock(mtProtectionClocks []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_PROTECTION_CLOCK` IN (?)", mtProtectionClocks).Find(&results).Error
	
	return
}
 
// GetFromMtSnapshotVersion 通过MT_SNAPSHOT_VERSION获取内容  
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetFromMtSnapshotVersion(mtSnapshotVersion int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_SNAPSHOT_VERSION` = ?", mtSnapshotVersion).Find(&results).Error
	
	return
}

// GetBatchFromMtSnapshotVersion 批量查找 
func (obj *_V$tenantMemstoreAllocatorInfoMgr) GetBatchFromMtSnapshotVersion(mtSnapshotVersions []int64) (results []*V$tenantMemstoreAllocatorInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$tenantMemstoreAllocatorInfo{}).Where("`MT_SNAPSHOT_VERSION` IN (?)", mtSnapshotVersions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

