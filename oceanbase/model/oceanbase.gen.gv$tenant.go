package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$tenantMgr struct {
	*_BaseMgr
}

// Gv$tenantMgr open func
func Gv$tenantMgr(db *gorm.DB) *_Gv$tenantMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$tenantMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$tenantMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$tenant"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$tenantMgr) GetTableName() string {
	return "gv$tenant"
}

// Reset 重置gorm会话
func (obj *_Gv$tenantMgr) Reset() *_Gv$tenantMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$tenantMgr) Get() (result Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$tenantMgr) Gets() (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$tenantMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_Gv$tenantMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取 
func (obj *_Gv$tenantMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithZoneList zone_list获取 
func (obj *_Gv$tenantMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取 
func (obj *_Gv$tenantMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithCollationType collation_type获取 
func (obj *_Gv$tenantMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithInfo info获取 
func (obj *_Gv$tenantMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithReadOnly read_only获取 
func (obj *_Gv$tenantMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithLocality locality获取 
func (obj *_Gv$tenantMgr) WithLocality(locality string) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$tenantMgr) GetByOption(opts ...Option) (result Gv$tenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$tenantMgr) GetByOptions(opts ...Option) (results []*Gv$tenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$tenantMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$tenant,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where(options.query)
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
func (obj *_Gv$tenantMgr) GetFromTenantID(tenantID int64) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantName 通过tenant_name获取内容  
func (obj *_Gv$tenantMgr) GetFromTenantName(tenantName string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error
	
	return
}

// GetBatchFromTenantName 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromTenantName(tenantNames []string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error
	
	return
}
 
// GetFromZoneList 通过zone_list获取内容  
func (obj *_Gv$tenantMgr) GetFromZoneList(zoneList string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`zone_list` = ?", zoneList).Find(&results).Error
	
	return
}

// GetBatchFromZoneList 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromZoneList(zoneLists []string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error
	
	return
}
 
// GetFromPrimaryZone 通过primary_zone获取内容  
func (obj *_Gv$tenantMgr) GetFromPrimaryZone(primaryZone string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error
	
	return
}

// GetBatchFromPrimaryZone 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error
	
	return
}
 
// GetFromCollationType 通过collation_type获取内容  
func (obj *_Gv$tenantMgr) GetFromCollationType(collationType int64) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`collation_type` = ?", collationType).Find(&results).Error
	
	return
}

// GetBatchFromCollationType 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromCollationType(collationTypes []int64) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error
	
	return
}
 
// GetFromInfo 通过info获取内容  
func (obj *_Gv$tenantMgr) GetFromInfo(info string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`info` = ?", info).Find(&results).Error
	
	return
}

// GetBatchFromInfo 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromInfo(infos []string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`info` IN (?)", infos).Find(&results).Error
	
	return
}
 
// GetFromReadOnly 通过read_only获取内容  
func (obj *_Gv$tenantMgr) GetFromReadOnly(readOnly int64) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`read_only` = ?", readOnly).Find(&results).Error
	
	return
}

// GetBatchFromReadOnly 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error
	
	return
}
 
// GetFromLocality 通过locality获取内容  
func (obj *_Gv$tenantMgr) GetFromLocality(locality string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`locality` = ?", locality).Find(&results).Error
	
	return
}

// GetBatchFromLocality 批量查找 
func (obj *_Gv$tenantMgr) GetBatchFromLocality(localitys []string) (results []*Gv$tenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenant{}).Where("`locality` IN (?)", localitys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

