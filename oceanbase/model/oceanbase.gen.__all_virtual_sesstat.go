package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualSesstatMgr struct {
	*_BaseMgr
}

// AllVirtualSesstatMgr open func
func AllVirtualSesstatMgr(db *gorm.DB) *_AllVirtualSesstatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSesstatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSesstatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sesstat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSesstatMgr) GetTableName() string {
	return "__all_virtual_sesstat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSesstatMgr) Reset() *_AllVirtualSesstatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualSesstatMgr) Get() (result AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSesstatMgr) Gets() (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSesstatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSessionID session_id获取 
func (obj *_AllVirtualSesstatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithSvrIP svr_ip获取 
func (obj *_AllVirtualSesstatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllVirtualSesstatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithStatistic# statistic#获取 
func (obj *_AllVirtualSesstatMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["statistic#"] = statistic# })
}

// WithTenantID tenant_id获取 
func (obj *_AllVirtualSesstatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithValue value获取 
func (obj *_AllVirtualSesstatMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithCanVisible can_visible获取 
func (obj *_AllVirtualSesstatMgr) WithCanVisible(canVisible int8) Option {
	return optionFunc(func(o *options) { o.query["can_visible"] = canVisible })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualSesstatMgr) GetByOption(opts ...Option) (result AllVirtualSesstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSesstatMgr) GetByOptions(opts ...Option) (results []*AllVirtualSesstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualSesstatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSesstat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where(options.query)
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


// GetFromSessionID 通过session_id获取内容  
func (obj *_AllVirtualSesstatMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_AllVirtualSesstatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllVirtualSesstatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllVirtualSesstatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllVirtualSesstatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllVirtualSesstatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过statistic#获取内容  
func (obj *_AllVirtualSesstatMgr) GetFromStatistic#(statistic# int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`statistic#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_AllVirtualSesstatMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`statistic#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_AllVirtualSesstatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllVirtualSesstatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_AllVirtualSesstatMgr) GetFromValue(value int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`value` = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量查找 
func (obj *_AllVirtualSesstatMgr) GetBatchFromValue(values []int64) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`value` IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromCanVisible 通过can_visible获取内容  
func (obj *_AllVirtualSesstatMgr) GetFromCanVisible(canVisible int8) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`can_visible` = ?", canVisible).Find(&results).Error
	
	return
}

// GetBatchFromCanVisible 批量查找 
func (obj *_AllVirtualSesstatMgr) GetBatchFromCanVisible(canVisibles []int8) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`can_visible` IN (?)", canVisibles).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllVirtualSesstatMgr) FetchByPrimaryKey(sessionID int64 ,svrIP string ,svrPort int64 ,statistic# int64 ) (result AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`session_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `statistic#` = ?", sessionID , svrIP , svrPort , statistic#).First(&result).Error
	
	return
}
 

 
 // FetchIndexByAllVirtualSesstatI1  获取多个内容
 func (obj *_AllVirtualSesstatMgr) FetchIndexByAllVirtualSesstatI1(sessionID int64 ) (results []*AllVirtualSesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSesstat{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}
 

	

