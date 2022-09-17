package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _Gv$memstoreInfoMgr struct {
	*_BaseMgr
}

// Gv$memstoreInfoMgr open func
func Gv$memstoreInfoMgr(db *gorm.DB) *_Gv$memstoreInfoMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$memstoreInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$memstoreInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$memstore_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$memstoreInfoMgr) GetTableName() string {
	return "gv$memstore_info"
}

// Reset 重置gorm会话
func (obj *_Gv$memstoreInfoMgr) Reset() *_Gv$memstoreInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$memstoreInfoMgr) Get() (result Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$memstoreInfoMgr) Gets() (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$memstoreInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_Gv$memstoreInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIP IP获取 
func (obj *_Gv$memstoreInfoMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["IP"] = ip })
}

// WithPort PORT获取 
func (obj *_Gv$memstoreInfoMgr) WithPort(port int64) Option {
	return optionFunc(func(o *options) { o.query["PORT"] = port })
}

// WithTableID TABLE_ID获取 
func (obj *_Gv$memstoreInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取 
func (obj *_Gv$memstoreInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithVersion VERSION获取 
func (obj *_Gv$memstoreInfoMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["VERSION"] = version })
}

// WithBaseVersion BASE_VERSION获取 
func (obj *_Gv$memstoreInfoMgr) WithBaseVersion(baseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["BASE_VERSION"] = baseVersion })
}

// WithMultiVersionStart MULTI_VERSION_START获取 
func (obj *_Gv$memstoreInfoMgr) WithMultiVersionStart(multiVersionStart int64) Option {
	return optionFunc(func(o *options) { o.query["MULTI_VERSION_START"] = multiVersionStart })
}

// WithSnapshotVersion SNAPSHOT_VERSION获取 
func (obj *_Gv$memstoreInfoMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["SNAPSHOT_VERSION"] = snapshotVersion })
}

// WithIsActive IS_ACTIVE获取 
func (obj *_Gv$memstoreInfoMgr) WithIsActive(isActive int64) Option {
	return optionFunc(func(o *options) { o.query["IS_ACTIVE"] = isActive })
}

// WithUsed USED获取 
func (obj *_Gv$memstoreInfoMgr) WithUsed(used int64) Option {
	return optionFunc(func(o *options) { o.query["USED"] = used })
}

// WithHashItems HASH_ITEMS获取 
func (obj *_Gv$memstoreInfoMgr) WithHashItems(hashItems int64) Option {
	return optionFunc(func(o *options) { o.query["HASH_ITEMS"] = hashItems })
}

// WithBtreeItems BTREE_ITEMS获取 
func (obj *_Gv$memstoreInfoMgr) WithBtreeItems(btreeItems int64) Option {
	return optionFunc(func(o *options) { o.query["BTREE_ITEMS"] = btreeItems })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$memstoreInfoMgr) GetByOption(opts ...Option) (result Gv$memstoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$memstoreInfoMgr) GetByOptions(opts ...Option) (results []*Gv$memstoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$memstoreInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$memstoreInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where(options.query)
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


// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromTenantID(tenantID int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromIP 通过IP获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromIP(ip string) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`IP` = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromIP(ips []string) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`IP` IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过PORT获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromPort(port int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`PORT` = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromPort(ports []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`PORT` IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromTableID(tableID int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过PARTITION_ID获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromPartitionID(partitionID int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromVersion 通过VERSION获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromVersion(version string) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`VERSION` = ?", version).Find(&results).Error
	
	return
}

// GetBatchFromVersion 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromVersion(versions []string) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`VERSION` IN (?)", versions).Find(&results).Error
	
	return
}
 
// GetFromBaseVersion 通过BASE_VERSION获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromBaseVersion(baseVersion int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`BASE_VERSION` = ?", baseVersion).Find(&results).Error
	
	return
}

// GetBatchFromBaseVersion 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromBaseVersion(baseVersions []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`BASE_VERSION` IN (?)", baseVersions).Find(&results).Error
	
	return
}
 
// GetFromMultiVersionStart 通过MULTI_VERSION_START获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromMultiVersionStart(multiVersionStart int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`MULTI_VERSION_START` = ?", multiVersionStart).Find(&results).Error
	
	return
}

// GetBatchFromMultiVersionStart 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromMultiVersionStart(multiVersionStarts []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`MULTI_VERSION_START` IN (?)", multiVersionStarts).Find(&results).Error
	
	return
}
 
// GetFromSnapshotVersion 通过SNAPSHOT_VERSION获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`SNAPSHOT_VERSION` = ?", snapshotVersion).Find(&results).Error
	
	return
}

// GetBatchFromSnapshotVersion 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`SNAPSHOT_VERSION` IN (?)", snapshotVersions).Find(&results).Error
	
	return
}
 
// GetFromIsActive 通过IS_ACTIVE获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromIsActive(isActive int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`IS_ACTIVE` = ?", isActive).Find(&results).Error
	
	return
}

// GetBatchFromIsActive 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromIsActive(isActives []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`IS_ACTIVE` IN (?)", isActives).Find(&results).Error
	
	return
}
 
// GetFromUsed 通过USED获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromUsed(used int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`USED` = ?", used).Find(&results).Error
	
	return
}

// GetBatchFromUsed 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromUsed(useds []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`USED` IN (?)", useds).Find(&results).Error
	
	return
}
 
// GetFromHashItems 通过HASH_ITEMS获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromHashItems(hashItems int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`HASH_ITEMS` = ?", hashItems).Find(&results).Error
	
	return
}

// GetBatchFromHashItems 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromHashItems(hashItemss []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`HASH_ITEMS` IN (?)", hashItemss).Find(&results).Error
	
	return
}
 
// GetFromBtreeItems 通过BTREE_ITEMS获取内容  
func (obj *_Gv$memstoreInfoMgr) GetFromBtreeItems(btreeItems int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`BTREE_ITEMS` = ?", btreeItems).Find(&results).Error
	
	return
}

// GetBatchFromBtreeItems 批量查找 
func (obj *_Gv$memstoreInfoMgr) GetBatchFromBtreeItems(btreeItemss []int64) (results []*Gv$memstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstoreInfo{}).Where("`BTREE_ITEMS` IN (?)", btreeItemss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

