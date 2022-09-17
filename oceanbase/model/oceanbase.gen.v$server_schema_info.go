package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$serverSchemaInfoMgr struct {
	*_BaseMgr
}

// V$serverSchemaInfoMgr open func
func V$serverSchemaInfoMgr(db *gorm.DB) *_V$serverSchemaInfoMgr {
	if db == nil {
		panic(fmt.Errorf("V$serverSchemaInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$serverSchemaInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$server_schema_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$serverSchemaInfoMgr) GetTableName() string {
	return "v$server_schema_info"
}

// Reset 重置gorm会话
func (obj *_V$serverSchemaInfoMgr) Reset() *_V$serverSchemaInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$serverSchemaInfoMgr) Get() (result V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$serverSchemaInfoMgr) Gets() (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$serverSchemaInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_V$serverSchemaInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithRefreshedSchemaVersion REFRESHED_SCHEMA_VERSION获取 
func (obj *_V$serverSchemaInfoMgr) WithRefreshedSchemaVersion(refreshedSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["REFRESHED_SCHEMA_VERSION"] = refreshedSchemaVersion })
}

// WithReceivedSchemaVersion RECEIVED_SCHEMA_VERSION获取 
func (obj *_V$serverSchemaInfoMgr) WithReceivedSchemaVersion(receivedSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["RECEIVED_SCHEMA_VERSION"] = receivedSchemaVersion })
}

// WithSchemaCount SCHEMA_COUNT获取 
func (obj *_V$serverSchemaInfoMgr) WithSchemaCount(schemaCount int64) Option {
	return optionFunc(func(o *options) { o.query["SCHEMA_COUNT"] = schemaCount })
}

// WithSchemaSize SCHEMA_SIZE获取 
func (obj *_V$serverSchemaInfoMgr) WithSchemaSize(schemaSize int64) Option {
	return optionFunc(func(o *options) { o.query["SCHEMA_SIZE"] = schemaSize })
}

// WithMinSstableSchemaVersion MIN_SSTABLE_SCHEMA_VERSION获取 
func (obj *_V$serverSchemaInfoMgr) WithMinSstableSchemaVersion(minSstableSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["MIN_SSTABLE_SCHEMA_VERSION"] = minSstableSchemaVersion })
}


// GetByOption 功能选项模式获取
func (obj *_V$serverSchemaInfoMgr) GetByOption(opts ...Option) (result V$serverSchemaInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$serverSchemaInfoMgr) GetByOptions(opts ...Option) (results []*V$serverSchemaInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$serverSchemaInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$serverSchemaInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where(options.query)
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
func (obj *_V$serverSchemaInfoMgr) GetFromTenantID(tenantID int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$serverSchemaInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromRefreshedSchemaVersion 通过REFRESHED_SCHEMA_VERSION获取内容  
func (obj *_V$serverSchemaInfoMgr) GetFromRefreshedSchemaVersion(refreshedSchemaVersion int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`REFRESHED_SCHEMA_VERSION` = ?", refreshedSchemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromRefreshedSchemaVersion 批量查找 
func (obj *_V$serverSchemaInfoMgr) GetBatchFromRefreshedSchemaVersion(refreshedSchemaVersions []int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`REFRESHED_SCHEMA_VERSION` IN (?)", refreshedSchemaVersions).Find(&results).Error
	
	return
}
 
// GetFromReceivedSchemaVersion 通过RECEIVED_SCHEMA_VERSION获取内容  
func (obj *_V$serverSchemaInfoMgr) GetFromReceivedSchemaVersion(receivedSchemaVersion int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`RECEIVED_SCHEMA_VERSION` = ?", receivedSchemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromReceivedSchemaVersion 批量查找 
func (obj *_V$serverSchemaInfoMgr) GetBatchFromReceivedSchemaVersion(receivedSchemaVersions []int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`RECEIVED_SCHEMA_VERSION` IN (?)", receivedSchemaVersions).Find(&results).Error
	
	return
}
 
// GetFromSchemaCount 通过SCHEMA_COUNT获取内容  
func (obj *_V$serverSchemaInfoMgr) GetFromSchemaCount(schemaCount int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`SCHEMA_COUNT` = ?", schemaCount).Find(&results).Error
	
	return
}

// GetBatchFromSchemaCount 批量查找 
func (obj *_V$serverSchemaInfoMgr) GetBatchFromSchemaCount(schemaCounts []int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`SCHEMA_COUNT` IN (?)", schemaCounts).Find(&results).Error
	
	return
}
 
// GetFromSchemaSize 通过SCHEMA_SIZE获取内容  
func (obj *_V$serverSchemaInfoMgr) GetFromSchemaSize(schemaSize int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`SCHEMA_SIZE` = ?", schemaSize).Find(&results).Error
	
	return
}

// GetBatchFromSchemaSize 批量查找 
func (obj *_V$serverSchemaInfoMgr) GetBatchFromSchemaSize(schemaSizes []int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`SCHEMA_SIZE` IN (?)", schemaSizes).Find(&results).Error
	
	return
}
 
// GetFromMinSstableSchemaVersion 通过MIN_SSTABLE_SCHEMA_VERSION获取内容  
func (obj *_V$serverSchemaInfoMgr) GetFromMinSstableSchemaVersion(minSstableSchemaVersion int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`MIN_SSTABLE_SCHEMA_VERSION` = ?", minSstableSchemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromMinSstableSchemaVersion 批量查找 
func (obj *_V$serverSchemaInfoMgr) GetBatchFromMinSstableSchemaVersion(minSstableSchemaVersions []int64) (results []*V$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$serverSchemaInfo{}).Where("`MIN_SSTABLE_SCHEMA_VERSION` IN (?)", minSstableSchemaVersions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

