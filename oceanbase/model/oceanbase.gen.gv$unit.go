package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$unitMgr struct {
	*_BaseMgr
}

// Gv$unitMgr open func
func Gv$unitMgr(db *gorm.DB) *_Gv$unitMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$unitMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$unitMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$unit"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$unitMgr) GetTableName() string {
	return "gv$unit"
}

// Reset 重置gorm会话
func (obj *_Gv$unitMgr) Reset() *_Gv$unitMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$unitMgr) Get() (result Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$unitMgr) Gets() (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$unitMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUnitID unit_id获取 
func (obj *_Gv$unitMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithUnitConfigID unit_config_id获取 
func (obj *_Gv$unitMgr) WithUnitConfigID(unitConfigID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_config_id"] = unitConfigID })
}

// WithUnitConfigName unit_config_name获取 
func (obj *_Gv$unitMgr) WithUnitConfigName(unitConfigName string) Option {
	return optionFunc(func(o *options) { o.query["unit_config_name"] = unitConfigName })
}

// WithResourcePoolID resource_pool_id获取 
func (obj *_Gv$unitMgr) WithResourcePoolID(resourcePoolID int64) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_id"] = resourcePoolID })
}

// WithResourcePoolName resource_pool_name获取 
func (obj *_Gv$unitMgr) WithResourcePoolName(resourcePoolName string) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_name"] = resourcePoolName })
}

// WithZone zone获取 
func (obj *_Gv$unitMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithTenantID tenant_id获取 
func (obj *_Gv$unitMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取 
func (obj *_Gv$unitMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithSvrIP svr_ip获取 
func (obj *_Gv$unitMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_Gv$unitMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithMigrateFromSvrIP migrate_from_svr_ip获取 
func (obj *_Gv$unitMgr) WithMigrateFromSvrIP(migrateFromSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["migrate_from_svr_ip"] = migrateFromSvrIP })
}

// WithMigrateFromSvrPort migrate_from_svr_port获取 
func (obj *_Gv$unitMgr) WithMigrateFromSvrPort(migrateFromSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["migrate_from_svr_port"] = migrateFromSvrPort })
}

// WithMaxCPU max_cpu获取 
func (obj *_Gv$unitMgr) WithMaxCPU(maxCPU float64) Option {
	return optionFunc(func(o *options) { o.query["max_cpu"] = maxCPU })
}

// WithMinCPU min_cpu获取 
func (obj *_Gv$unitMgr) WithMinCPU(minCPU float64) Option {
	return optionFunc(func(o *options) { o.query["min_cpu"] = minCPU })
}

// WithMaxMemory max_memory获取 
func (obj *_Gv$unitMgr) WithMaxMemory(maxMemory int64) Option {
	return optionFunc(func(o *options) { o.query["max_memory"] = maxMemory })
}

// WithMinMemory min_memory获取 
func (obj *_Gv$unitMgr) WithMinMemory(minMemory int64) Option {
	return optionFunc(func(o *options) { o.query["min_memory"] = minMemory })
}

// WithMaxIops max_iops获取 
func (obj *_Gv$unitMgr) WithMaxIops(maxIops int64) Option {
	return optionFunc(func(o *options) { o.query["max_iops"] = maxIops })
}

// WithMinIops min_iops获取 
func (obj *_Gv$unitMgr) WithMinIops(minIops int64) Option {
	return optionFunc(func(o *options) { o.query["min_iops"] = minIops })
}

// WithMaxDiskSize max_disk_size获取 
func (obj *_Gv$unitMgr) WithMaxDiskSize(maxDiskSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_disk_size"] = maxDiskSize })
}

// WithMaxSessionNum max_session_num获取 
func (obj *_Gv$unitMgr) WithMaxSessionNum(maxSessionNum int64) Option {
	return optionFunc(func(o *options) { o.query["max_session_num"] = maxSessionNum })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$unitMgr) GetByOption(opts ...Option) (result Gv$unit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$unitMgr) GetByOptions(opts ...Option) (results []*Gv$unit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$unitMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$unit,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where(options.query)
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
func (obj *_Gv$unitMgr) GetFromUnitID(unitID int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`unit_id` = ?", unitID).Find(&results).Error
	
	return
}

// GetBatchFromUnitID 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromUnitID(unitIDs []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error
	
	return
}
 
// GetFromUnitConfigID 通过unit_config_id获取内容  
func (obj *_Gv$unitMgr) GetFromUnitConfigID(unitConfigID int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`unit_config_id` = ?", unitConfigID).Find(&results).Error
	
	return
}

// GetBatchFromUnitConfigID 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromUnitConfigID(unitConfigIDs []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`unit_config_id` IN (?)", unitConfigIDs).Find(&results).Error
	
	return
}
 
// GetFromUnitConfigName 通过unit_config_name获取内容  
func (obj *_Gv$unitMgr) GetFromUnitConfigName(unitConfigName string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`unit_config_name` = ?", unitConfigName).Find(&results).Error
	
	return
}

// GetBatchFromUnitConfigName 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromUnitConfigName(unitConfigNames []string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`unit_config_name` IN (?)", unitConfigNames).Find(&results).Error
	
	return
}
 
// GetFromResourcePoolID 通过resource_pool_id获取内容  
func (obj *_Gv$unitMgr) GetFromResourcePoolID(resourcePoolID int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`resource_pool_id` = ?", resourcePoolID).Find(&results).Error
	
	return
}

// GetBatchFromResourcePoolID 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromResourcePoolID(resourcePoolIDs []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`resource_pool_id` IN (?)", resourcePoolIDs).Find(&results).Error
	
	return
}
 
// GetFromResourcePoolName 通过resource_pool_name获取内容  
func (obj *_Gv$unitMgr) GetFromResourcePoolName(resourcePoolName string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`resource_pool_name` = ?", resourcePoolName).Find(&results).Error
	
	return
}

// GetBatchFromResourcePoolName 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromResourcePoolName(resourcePoolNames []string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`resource_pool_name` IN (?)", resourcePoolNames).Find(&results).Error
	
	return
}
 
// GetFromZone 通过zone获取内容  
func (obj *_Gv$unitMgr) GetFromZone(zone string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`zone` = ?", zone).Find(&results).Error
	
	return
}

// GetBatchFromZone 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromZone(zones []string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`zone` IN (?)", zones).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_Gv$unitMgr) GetFromTenantID(tenantID int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantName 通过tenant_name获取内容  
func (obj *_Gv$unitMgr) GetFromTenantName(tenantName string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error
	
	return
}

// GetBatchFromTenantName 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromTenantName(tenantNames []string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_Gv$unitMgr) GetFromSvrIP(svrIP string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_Gv$unitMgr) GetFromSvrPort(svrPort int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromMigrateFromSvrIP 通过migrate_from_svr_ip获取内容  
func (obj *_Gv$unitMgr) GetFromMigrateFromSvrIP(migrateFromSvrIP string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`migrate_from_svr_ip` = ?", migrateFromSvrIP).Find(&results).Error
	
	return
}

// GetBatchFromMigrateFromSvrIP 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMigrateFromSvrIP(migrateFromSvrIPs []string) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`migrate_from_svr_ip` IN (?)", migrateFromSvrIPs).Find(&results).Error
	
	return
}
 
// GetFromMigrateFromSvrPort 通过migrate_from_svr_port获取内容  
func (obj *_Gv$unitMgr) GetFromMigrateFromSvrPort(migrateFromSvrPort int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`migrate_from_svr_port` = ?", migrateFromSvrPort).Find(&results).Error
	
	return
}

// GetBatchFromMigrateFromSvrPort 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMigrateFromSvrPort(migrateFromSvrPorts []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`migrate_from_svr_port` IN (?)", migrateFromSvrPorts).Find(&results).Error
	
	return
}
 
// GetFromMaxCPU 通过max_cpu获取内容  
func (obj *_Gv$unitMgr) GetFromMaxCPU(maxCPU float64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_cpu` = ?", maxCPU).Find(&results).Error
	
	return
}

// GetBatchFromMaxCPU 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMaxCPU(maxCPUs []float64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_cpu` IN (?)", maxCPUs).Find(&results).Error
	
	return
}
 
// GetFromMinCPU 通过min_cpu获取内容  
func (obj *_Gv$unitMgr) GetFromMinCPU(minCPU float64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`min_cpu` = ?", minCPU).Find(&results).Error
	
	return
}

// GetBatchFromMinCPU 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMinCPU(minCPUs []float64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`min_cpu` IN (?)", minCPUs).Find(&results).Error
	
	return
}
 
// GetFromMaxMemory 通过max_memory获取内容  
func (obj *_Gv$unitMgr) GetFromMaxMemory(maxMemory int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_memory` = ?", maxMemory).Find(&results).Error
	
	return
}

// GetBatchFromMaxMemory 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMaxMemory(maxMemorys []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_memory` IN (?)", maxMemorys).Find(&results).Error
	
	return
}
 
// GetFromMinMemory 通过min_memory获取内容  
func (obj *_Gv$unitMgr) GetFromMinMemory(minMemory int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`min_memory` = ?", minMemory).Find(&results).Error
	
	return
}

// GetBatchFromMinMemory 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMinMemory(minMemorys []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`min_memory` IN (?)", minMemorys).Find(&results).Error
	
	return
}
 
// GetFromMaxIops 通过max_iops获取内容  
func (obj *_Gv$unitMgr) GetFromMaxIops(maxIops int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_iops` = ?", maxIops).Find(&results).Error
	
	return
}

// GetBatchFromMaxIops 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMaxIops(maxIopss []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_iops` IN (?)", maxIopss).Find(&results).Error
	
	return
}
 
// GetFromMinIops 通过min_iops获取内容  
func (obj *_Gv$unitMgr) GetFromMinIops(minIops int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`min_iops` = ?", minIops).Find(&results).Error
	
	return
}

// GetBatchFromMinIops 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMinIops(minIopss []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`min_iops` IN (?)", minIopss).Find(&results).Error
	
	return
}
 
// GetFromMaxDiskSize 通过max_disk_size获取内容  
func (obj *_Gv$unitMgr) GetFromMaxDiskSize(maxDiskSize int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_disk_size` = ?", maxDiskSize).Find(&results).Error
	
	return
}

// GetBatchFromMaxDiskSize 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMaxDiskSize(maxDiskSizes []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_disk_size` IN (?)", maxDiskSizes).Find(&results).Error
	
	return
}
 
// GetFromMaxSessionNum 通过max_session_num获取内容  
func (obj *_Gv$unitMgr) GetFromMaxSessionNum(maxSessionNum int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_session_num` = ?", maxSessionNum).Find(&results).Error
	
	return
}

// GetBatchFromMaxSessionNum 批量查找 
func (obj *_Gv$unitMgr) GetBatchFromMaxSessionNum(maxSessionNums []int64) (results []*Gv$unit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unit{}).Where("`max_session_num` IN (?)", maxSessionNums).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

