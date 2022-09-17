package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$databaseMgr struct {
	*_BaseMgr
}

// Gv$databaseMgr open func
func Gv$databaseMgr(db *gorm.DB) *_Gv$databaseMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$databaseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$databaseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$database"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$databaseMgr) GetTableName() string {
	return "gv$database"
}

// Reset 重置gorm会话
func (obj *_Gv$databaseMgr) Reset() *_Gv$databaseMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$databaseMgr) Get() (result Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$databaseMgr) Gets() (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$databaseMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_Gv$databaseMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取 
func (obj *_Gv$databaseMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithDatabaseID database_id获取 
func (obj *_Gv$databaseMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithDatabaseName database_name获取 
func (obj *_Gv$databaseMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithZoneList zone_list获取 
func (obj *_Gv$databaseMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取 
func (obj *_Gv$databaseMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithCollationType collation_type获取 
func (obj *_Gv$databaseMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithComment comment获取 
func (obj *_Gv$databaseMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithReadOnly read_only获取 
func (obj *_Gv$databaseMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithDefaultTablegroupID default_tablegroup_id获取 
func (obj *_Gv$databaseMgr) WithDefaultTablegroupID(defaultTablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["default_tablegroup_id"] = defaultTablegroupID })
}

// WithInRecyclebin in_recyclebin获取 
func (obj *_Gv$databaseMgr) WithInRecyclebin(inRecyclebin int64) Option {
	return optionFunc(func(o *options) { o.query["in_recyclebin"] = inRecyclebin })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$databaseMgr) GetByOption(opts ...Option) (result Gv$database, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$databaseMgr) GetByOptions(opts ...Option) (results []*Gv$database, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$databaseMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$database,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where(options.query)
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
func (obj *_Gv$databaseMgr) GetFromTenantID(tenantID int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantName 通过tenant_name获取内容  
func (obj *_Gv$databaseMgr) GetFromTenantName(tenantName string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error
	
	return
}

// GetBatchFromTenantName 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromTenantName(tenantNames []string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error
	
	return
}
 
// GetFromDatabaseID 通过database_id获取内容  
func (obj *_Gv$databaseMgr) GetFromDatabaseID(databaseID int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`database_id` = ?", databaseID).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseID 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error
	
	return
}
 
// GetFromDatabaseName 通过database_name获取内容  
func (obj *_Gv$databaseMgr) GetFromDatabaseName(databaseName string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`database_name` = ?", databaseName).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseName 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error
	
	return
}
 
// GetFromZoneList 通过zone_list获取内容  
func (obj *_Gv$databaseMgr) GetFromZoneList(zoneList string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`zone_list` = ?", zoneList).Find(&results).Error
	
	return
}

// GetBatchFromZoneList 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromZoneList(zoneLists []string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error
	
	return
}
 
// GetFromPrimaryZone 通过primary_zone获取内容  
func (obj *_Gv$databaseMgr) GetFromPrimaryZone(primaryZone string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error
	
	return
}

// GetBatchFromPrimaryZone 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error
	
	return
}
 
// GetFromCollationType 通过collation_type获取内容  
func (obj *_Gv$databaseMgr) GetFromCollationType(collationType int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`collation_type` = ?", collationType).Find(&results).Error
	
	return
}

// GetBatchFromCollationType 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromCollationType(collationTypes []int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error
	
	return
}
 
// GetFromComment 通过comment获取内容  
func (obj *_Gv$databaseMgr) GetFromComment(comment string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`comment` = ?", comment).Find(&results).Error
	
	return
}

// GetBatchFromComment 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromComment(comments []string) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`comment` IN (?)", comments).Find(&results).Error
	
	return
}
 
// GetFromReadOnly 通过read_only获取内容  
func (obj *_Gv$databaseMgr) GetFromReadOnly(readOnly int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`read_only` = ?", readOnly).Find(&results).Error
	
	return
}

// GetBatchFromReadOnly 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error
	
	return
}
 
// GetFromDefaultTablegroupID 通过default_tablegroup_id获取内容  
func (obj *_Gv$databaseMgr) GetFromDefaultTablegroupID(defaultTablegroupID int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`default_tablegroup_id` = ?", defaultTablegroupID).Find(&results).Error
	
	return
}

// GetBatchFromDefaultTablegroupID 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromDefaultTablegroupID(defaultTablegroupIDs []int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`default_tablegroup_id` IN (?)", defaultTablegroupIDs).Find(&results).Error
	
	return
}
 
// GetFromInRecyclebin 通过in_recyclebin获取内容  
func (obj *_Gv$databaseMgr) GetFromInRecyclebin(inRecyclebin int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`in_recyclebin` = ?", inRecyclebin).Find(&results).Error
	
	return
}

// GetBatchFromInRecyclebin 批量查找 
func (obj *_Gv$databaseMgr) GetBatchFromInRecyclebin(inRecyclebins []int64) (results []*Gv$database, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$database{}).Where("`in_recyclebin` IN (?)", inRecyclebins).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

