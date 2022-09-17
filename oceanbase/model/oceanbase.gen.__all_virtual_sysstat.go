package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualSysstatMgr struct {
	*_BaseMgr
}

// AllVirtualSysstatMgr open func
func AllVirtualSysstatMgr(db *gorm.DB) *_AllVirtualSysstatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysstatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysstatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sysstat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysstatMgr) GetTableName() string {
	return "__all_virtual_sysstat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysstatMgr) Reset() *_AllVirtualSysstatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualSysstatMgr) Get() (result AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysstatMgr) Gets() (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysstatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_AllVirtualSysstatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取 
func (obj *_AllVirtualSysstatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllVirtualSysstatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithStatistic# statistic#获取 
func (obj *_AllVirtualSysstatMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["statistic#"] = statistic# })
}

// WithValue value获取 
func (obj *_AllVirtualSysstatMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithStatID stat_id获取 
func (obj *_AllVirtualSysstatMgr) WithStatID(statID int64) Option {
	return optionFunc(func(o *options) { o.query["stat_id"] = statID })
}

// WithName name获取 
func (obj *_AllVirtualSysstatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithClass class获取 
func (obj *_AllVirtualSysstatMgr) WithClass(class int64) Option {
	return optionFunc(func(o *options) { o.query["class"] = class })
}

// WithCanVisible can_visible获取 
func (obj *_AllVirtualSysstatMgr) WithCanVisible(canVisible int8) Option {
	return optionFunc(func(o *options) { o.query["can_visible"] = canVisible })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysstatMgr) GetByOption(opts ...Option) (result AllVirtualSysstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysstatMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualSysstatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysstat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where(options.query)
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
func (obj *_AllVirtualSysstatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过statistic#获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromStatistic#(statistic# int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`statistic#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`statistic#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromValue(value int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`value` = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromValue(values []int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`value` IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromStatID 通过stat_id获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromStatID(statID int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`stat_id` = ?", statID).Find(&results).Error
	
	return
}

// GetBatchFromStatID 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromStatID(statIDs []int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`stat_id` IN (?)", statIDs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromName(name string) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromName(names []string) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromClass 通过class获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromClass(class int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`class` = ?", class).Find(&results).Error
	
	return
}

// GetBatchFromClass 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromClass(classs []int64) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`class` IN (?)", classs).Find(&results).Error
	
	return
}
 
// GetFromCanVisible 通过can_visible获取内容  
func (obj *_AllVirtualSysstatMgr) GetFromCanVisible(canVisible int8) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`can_visible` = ?", canVisible).Find(&results).Error
	
	return
}

// GetBatchFromCanVisible 批量查找 
func (obj *_AllVirtualSysstatMgr) GetBatchFromCanVisible(canVisibles []int8) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`can_visible` IN (?)", canVisibles).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllVirtualSysstatMgr) FetchByPrimaryKey(tenantID int64 ,svrIP string ,svrPort int64 ,statistic# int64 ) (result AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `statistic#` = ?", tenantID , svrIP , svrPort , statistic#).First(&result).Error
	
	return
}
 

 
 // FetchIndexByAllVirtualSysstatI1  获取多个内容
 func (obj *_AllVirtualSysstatMgr) FetchIndexByAllVirtualSysstatI1(tenantID int64 ) (results []*AllVirtualSysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysstat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}
 

	

