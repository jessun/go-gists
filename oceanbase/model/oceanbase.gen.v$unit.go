package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$unitMgr struct {
	*_BaseMgr
}

// V$unitMgr open func
func V$unitMgr(db *gorm.DB) *_V$unitMgr {
	if db == nil {
		panic(fmt.Errorf("V$unitMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$unitMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$unit"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$unitMgr) GetTableName() string {
	return "v$unit"
}

// Reset 重置gorm会话
func (obj *_V$unitMgr) Reset() *_V$unitMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$unitMgr) Get() (result V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$unitMgr) Gets() (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$unitMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$unit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUnitID unit_id获取 
func (obj *_V$unitMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithUnitConfigID unit_config_id获取 
func (obj *_V$unitMgr) WithUnitConfigID(unitConfigID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_config_id"] = unitConfigID })
}

// WithUnitConfigName unit_config_name获取 
func (obj *_V$unitMgr) WithUnitConfigName(unitConfigName string) Option {
	return optionFunc(func(o *options) { o.query["unit_config_name"] = unitConfigName })
}

// WithResourcePoolID resource_pool_id获取 
func (obj *_V$unitMgr) WithResourcePoolID(resourcePoolID int64) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_id"] = resourcePoolID })
}

// WithResourcePoolName resource_pool_name获取 
func (obj *_V$unitMgr) WithResourcePoolName(resourcePoolName string) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_name"] = resourcePoolName })
}

// WithZone zone获取 
func (obj *_V$unitMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithTenantID tenant_id获取 
func (obj *_V$unitMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取 
func (obj *_V$unitMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithSvrIP svr_ip获取 
func (obj *_V$unitMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_V$unitMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithMigrateFromSvrIP migrate_from_svr_ip获取 
func (obj *_V$unitMgr) WithMigrateFromSvrIP(migrateFromSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["migrate_from_svr_ip"] = migrateFromSvrIP })
}

// WithMigrateFromSvrPort migrate_from_svr_port获取 
func (obj *_V$unitMgr) WithMigrateFromSvrPort(migrateFromSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["migrate_from_svr_port"] = migrateFromSvrPort })
}

// WithMaxCPU max_cpu获取 
func (obj *_V$unitMgr) WithMaxCPU(maxCPU float64) Option {
	return optionFunc(func(o *options) { o.query["max_cpu"] = maxCPU })
}

// WithMinCPU min_cpu获取 
func (obj *_V$unitMgr) WithMinCPU(minCPU float64) Option {
	return optionFunc(func(o *options) { o.query["min_cpu"] = minCPU })
}

// WithMaxMemory max_memory获取 
func (obj *_V$unitMgr) WithMaxMemory(maxMemory int64) Option {
	return optionFunc(func(o *options) { o.query["max_memory"] = maxMemory })
}

// WithMinMemory min_memory获取 
func (obj *_V$unitMgr) WithMinMemory(minMemory int64) Option {
	return optionFunc(func(o *options) { o.query["min_memory"] = minMemory })
}

// WithMaxIops max_iops获取 
func (obj *_V$unitMgr) WithMaxIops(maxIops int64) Option {
	return optionFunc(func(o *options) { o.query["max_iops"] = maxIops })
}

// WithMinIops min_iops获取 
func (obj *_V$unitMgr) WithMinIops(minIops int64) Option {
	return optionFunc(func(o *options) { o.query["min_iops"] = minIops })
}

// WithMaxDiskSize max_disk_size获取 
func (obj *_V$unitMgr) WithMaxDiskSize(maxDiskSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_disk_size"] = maxDiskSize })
}

// WithMaxSessionNum max_session_num获取 
func (obj *_V$unitMgr) WithMaxSessionNum(maxSessionNum int64) Option {
	return optionFunc(func(o *options) { o.query["max_session_num"] = maxSessionNum })
}


// GetByOption 功能选项模式获取
func (obj *_V$unitMgr) GetByOption(opts ...Option) (result V$unit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$unitMgr) GetByOptions(opts ...Option) (results []*V$unit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$unitMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$unit,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where(options.query)
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


// GetFromUnitID 通过unit_id获取内容  
func (obj *_V$unitMgr) GetFromUnitID(unitID int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`unit_id` = ?", unitID).Find(&results).Error
	
	return
}

// GetBatchFromUnitID 批量查找 
func (obj *_V$unitMgr) GetBatchFromUnitID(unitIDs []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error
	
	return
}
 
// GetFromUnitConfigID 通过unit_config_id获取内容  
func (obj *_V$unitMgr) GetFromUnitConfigID(unitConfigID int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`unit_config_id` = ?", unitConfigID).Find(&results).Error
	
	return
}

// GetBatchFromUnitConfigID 批量查找 
func (obj *_V$unitMgr) GetBatchFromUnitConfigID(unitConfigIDs []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`unit_config_id` IN (?)", unitConfigIDs).Find(&results).Error
	
	return
}
 
// GetFromUnitConfigName 通过unit_config_name获取内容  
func (obj *_V$unitMgr) GetFromUnitConfigName(unitConfigName string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`unit_config_name` = ?", unitConfigName).Find(&results).Error
	
	return
}

// GetBatchFromUnitConfigName 批量查找 
func (obj *_V$unitMgr) GetBatchFromUnitConfigName(unitConfigNames []string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`unit_config_name` IN (?)", unitConfigNames).Find(&results).Error
	
	return
}
 
// GetFromResourcePoolID 通过resource_pool_id获取内容  
func (obj *_V$unitMgr) GetFromResourcePoolID(resourcePoolID int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`resource_pool_id` = ?", resourcePoolID).Find(&results).Error
	
	return
}

// GetBatchFromResourcePoolID 批量查找 
func (obj *_V$unitMgr) GetBatchFromResourcePoolID(resourcePoolIDs []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`resource_pool_id` IN (?)", resourcePoolIDs).Find(&results).Error
	
	return
}
 
// GetFromResourcePoolName 通过resource_pool_name获取内容  
func (obj *_V$unitMgr) GetFromResourcePoolName(resourcePoolName string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`resource_pool_name` = ?", resourcePoolName).Find(&results).Error
	
	return
}

// GetBatchFromResourcePoolName 批量查找 
func (obj *_V$unitMgr) GetBatchFromResourcePoolName(resourcePoolNames []string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`resource_pool_name` IN (?)", resourcePoolNames).Find(&results).Error
	
	return
}
 
// GetFromZone 通过zone获取内容  
func (obj *_V$unitMgr) GetFromZone(zone string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`zone` = ?", zone).Find(&results).Error
	
	return
}

// GetBatchFromZone 批量查找 
func (obj *_V$unitMgr) GetBatchFromZone(zones []string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`zone` IN (?)", zones).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_V$unitMgr) GetFromTenantID(tenantID int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$unitMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantName 通过tenant_name获取内容  
func (obj *_V$unitMgr) GetFromTenantName(tenantName string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error
	
	return
}

// GetBatchFromTenantName 批量查找 
func (obj *_V$unitMgr) GetBatchFromTenantName(tenantNames []string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_V$unitMgr) GetFromSvrIP(svrIP string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$unitMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_V$unitMgr) GetFromSvrPort(svrPort int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$unitMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromMigrateFromSvrIP 通过migrate_from_svr_ip获取内容  
func (obj *_V$unitMgr) GetFromMigrateFromSvrIP(migrateFromSvrIP string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`migrate_from_svr_ip` = ?", migrateFromSvrIP).Find(&results).Error
	
	return
}

// GetBatchFromMigrateFromSvrIP 批量查找 
func (obj *_V$unitMgr) GetBatchFromMigrateFromSvrIP(migrateFromSvrIPs []string) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`migrate_from_svr_ip` IN (?)", migrateFromSvrIPs).Find(&results).Error
	
	return
}
 
// GetFromMigrateFromSvrPort 通过migrate_from_svr_port获取内容  
func (obj *_V$unitMgr) GetFromMigrateFromSvrPort(migrateFromSvrPort int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`migrate_from_svr_port` = ?", migrateFromSvrPort).Find(&results).Error
	
	return
}

// GetBatchFromMigrateFromSvrPort 批量查找 
func (obj *_V$unitMgr) GetBatchFromMigrateFromSvrPort(migrateFromSvrPorts []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`migrate_from_svr_port` IN (?)", migrateFromSvrPorts).Find(&results).Error
	
	return
}
 
// GetFromMaxCPU 通过max_cpu获取内容  
func (obj *_V$unitMgr) GetFromMaxCPU(maxCPU float64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_cpu` = ?", maxCPU).Find(&results).Error
	
	return
}

// GetBatchFromMaxCPU 批量查找 
func (obj *_V$unitMgr) GetBatchFromMaxCPU(maxCPUs []float64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_cpu` IN (?)", maxCPUs).Find(&results).Error
	
	return
}
 
// GetFromMinCPU 通过min_cpu获取内容  
func (obj *_V$unitMgr) GetFromMinCPU(minCPU float64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`min_cpu` = ?", minCPU).Find(&results).Error
	
	return
}

// GetBatchFromMinCPU 批量查找 
func (obj *_V$unitMgr) GetBatchFromMinCPU(minCPUs []float64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`min_cpu` IN (?)", minCPUs).Find(&results).Error
	
	return
}
 
// GetFromMaxMemory 通过max_memory获取内容  
func (obj *_V$unitMgr) GetFromMaxMemory(maxMemory int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_memory` = ?", maxMemory).Find(&results).Error
	
	return
}

// GetBatchFromMaxMemory 批量查找 
func (obj *_V$unitMgr) GetBatchFromMaxMemory(maxMemorys []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_memory` IN (?)", maxMemorys).Find(&results).Error
	
	return
}
 
// GetFromMinMemory 通过min_memory获取内容  
func (obj *_V$unitMgr) GetFromMinMemory(minMemory int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`min_memory` = ?", minMemory).Find(&results).Error
	
	return
}

// GetBatchFromMinMemory 批量查找 
func (obj *_V$unitMgr) GetBatchFromMinMemory(minMemorys []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`min_memory` IN (?)", minMemorys).Find(&results).Error
	
	return
}
 
// GetFromMaxIops 通过max_iops获取内容  
func (obj *_V$unitMgr) GetFromMaxIops(maxIops int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_iops` = ?", maxIops).Find(&results).Error
	
	return
}

// GetBatchFromMaxIops 批量查找 
func (obj *_V$unitMgr) GetBatchFromMaxIops(maxIopss []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_iops` IN (?)", maxIopss).Find(&results).Error
	
	return
}
 
// GetFromMinIops 通过min_iops获取内容  
func (obj *_V$unitMgr) GetFromMinIops(minIops int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`min_iops` = ?", minIops).Find(&results).Error
	
	return
}

// GetBatchFromMinIops 批量查找 
func (obj *_V$unitMgr) GetBatchFromMinIops(minIopss []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`min_iops` IN (?)", minIopss).Find(&results).Error
	
	return
}
 
// GetFromMaxDiskSize 通过max_disk_size获取内容  
func (obj *_V$unitMgr) GetFromMaxDiskSize(maxDiskSize int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_disk_size` = ?", maxDiskSize).Find(&results).Error
	
	return
}

// GetBatchFromMaxDiskSize 批量查找 
func (obj *_V$unitMgr) GetBatchFromMaxDiskSize(maxDiskSizes []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_disk_size` IN (?)", maxDiskSizes).Find(&results).Error
	
	return
}
 
// GetFromMaxSessionNum 通过max_session_num获取内容  
func (obj *_V$unitMgr) GetFromMaxSessionNum(maxSessionNum int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_session_num` = ?", maxSessionNum).Find(&results).Error
	
	return
}

// GetBatchFromMaxSessionNum 批量查找 
func (obj *_V$unitMgr) GetBatchFromMaxSessionNum(maxSessionNums []int64) (results []*V$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$unit{}).Where("`max_session_num` IN (?)", maxSessionNums).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

