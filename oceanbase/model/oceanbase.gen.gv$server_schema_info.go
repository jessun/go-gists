package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _Gv$serverSchemaInfoMgr struct {
	*_BaseMgr
}

// Gv$serverSchemaInfoMgr open func
func Gv$serverSchemaInfoMgr(db *gorm.DB) *_Gv$serverSchemaInfoMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$serverSchemaInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$serverSchemaInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$server_schema_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$serverSchemaInfoMgr) GetTableName() string {
	return "gv$server_schema_info"
}

// Reset 重置gorm会话
func (obj *_Gv$serverSchemaInfoMgr) Reset() *_Gv$serverSchemaInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$serverSchemaInfoMgr) Get() (result Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$serverSchemaInfoMgr) Gets() (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$serverSchemaInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$serverSchemaInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$serverSchemaInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取 
func (obj *_Gv$serverSchemaInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithRefreshedSchemaVersion REFRESHED_SCHEMA_VERSION获取 
func (obj *_Gv$serverSchemaInfoMgr) WithRefreshedSchemaVersion(refreshedSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["REFRESHED_SCHEMA_VERSION"] = refreshedSchemaVersion })
}

// WithReceivedSchemaVersion RECEIVED_SCHEMA_VERSION获取 
func (obj *_Gv$serverSchemaInfoMgr) WithReceivedSchemaVersion(receivedSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["RECEIVED_SCHEMA_VERSION"] = receivedSchemaVersion })
}

// WithSchemaCount SCHEMA_COUNT获取 
func (obj *_Gv$serverSchemaInfoMgr) WithSchemaCount(schemaCount int64) Option {
	return optionFunc(func(o *options) { o.query["SCHEMA_COUNT"] = schemaCount })
}

// WithSchemaSize SCHEMA_SIZE获取 
func (obj *_Gv$serverSchemaInfoMgr) WithSchemaSize(schemaSize int64) Option {
	return optionFunc(func(o *options) { o.query["SCHEMA_SIZE"] = schemaSize })
}

// WithMinSstableSchemaVersion MIN_SSTABLE_SCHEMA_VERSION获取 
func (obj *_Gv$serverSchemaInfoMgr) WithMinSstableSchemaVersion(minSstableSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["MIN_SSTABLE_SCHEMA_VERSION"] = minSstableSchemaVersion })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$serverSchemaInfoMgr) GetByOption(opts ...Option) (result Gv$serverSchemaInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$serverSchemaInfoMgr) GetByOptions(opts ...Option) (results []*Gv$serverSchemaInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$serverSchemaInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$serverSchemaInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where(options.query)
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
func (obj *_Gv$serverSchemaInfoMgr) GetFromSvrIP(svrIP string) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$serverSchemaInfoMgr) GetFromSvrPort(svrPort int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$serverSchemaInfoMgr) GetFromTenantID(tenantID int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromRefreshedSchemaVersion 通过REFRESHED_SCHEMA_VERSION获取内容  
func (obj *_Gv$serverSchemaInfoMgr) GetFromRefreshedSchemaVersion(refreshedSchemaVersion int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`REFRESHED_SCHEMA_VERSION` = ?", refreshedSchemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromRefreshedSchemaVersion 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromRefreshedSchemaVersion(refreshedSchemaVersions []int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`REFRESHED_SCHEMA_VERSION` IN (?)", refreshedSchemaVersions).Find(&results).Error
	
	return
}
 
// GetFromReceivedSchemaVersion 通过RECEIVED_SCHEMA_VERSION获取内容  
func (obj *_Gv$serverSchemaInfoMgr) GetFromReceivedSchemaVersion(receivedSchemaVersion int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`RECEIVED_SCHEMA_VERSION` = ?", receivedSchemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromReceivedSchemaVersion 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromReceivedSchemaVersion(receivedSchemaVersions []int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`RECEIVED_SCHEMA_VERSION` IN (?)", receivedSchemaVersions).Find(&results).Error
	
	return
}
 
// GetFromSchemaCount 通过SCHEMA_COUNT获取内容  
func (obj *_Gv$serverSchemaInfoMgr) GetFromSchemaCount(schemaCount int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SCHEMA_COUNT` = ?", schemaCount).Find(&results).Error
	
	return
}

// GetBatchFromSchemaCount 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromSchemaCount(schemaCounts []int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SCHEMA_COUNT` IN (?)", schemaCounts).Find(&results).Error
	
	return
}
 
// GetFromSchemaSize 通过SCHEMA_SIZE获取内容  
func (obj *_Gv$serverSchemaInfoMgr) GetFromSchemaSize(schemaSize int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SCHEMA_SIZE` = ?", schemaSize).Find(&results).Error
	
	return
}

// GetBatchFromSchemaSize 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromSchemaSize(schemaSizes []int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`SCHEMA_SIZE` IN (?)", schemaSizes).Find(&results).Error
	
	return
}
 
// GetFromMinSstableSchemaVersion 通过MIN_SSTABLE_SCHEMA_VERSION获取内容  
func (obj *_Gv$serverSchemaInfoMgr) GetFromMinSstableSchemaVersion(minSstableSchemaVersion int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`MIN_SSTABLE_SCHEMA_VERSION` = ?", minSstableSchemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromMinSstableSchemaVersion 批量查找 
func (obj *_Gv$serverSchemaInfoMgr) GetBatchFromMinSstableSchemaVersion(minSstableSchemaVersions []int64) (results []*Gv$serverSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverSchemaInfo{}).Where("`MIN_SSTABLE_SCHEMA_VERSION` IN (?)", minSstableSchemaVersions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

