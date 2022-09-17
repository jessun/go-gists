package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualSystemEventMgr struct {
	*_BaseMgr
}

// AllVirtualSystemEventMgr open func
func AllVirtualSystemEventMgr(db *gorm.DB) *_AllVirtualSystemEventMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSystemEventMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSystemEventMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_system_event"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSystemEventMgr) GetTableName() string {
	return "__all_virtual_system_event"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSystemEventMgr) Reset() *_AllVirtualSystemEventMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualSystemEventMgr) Get() (result AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSystemEventMgr) Gets() (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSystemEventMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_AllVirtualSystemEventMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取 
func (obj *_AllVirtualSystemEventMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllVirtualSystemEventMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithEventID event_id获取 
func (obj *_AllVirtualSystemEventMgr) WithEventID(eventID int64) Option {
	return optionFunc(func(o *options) { o.query["event_id"] = eventID })
}

// WithEvent event获取 
func (obj *_AllVirtualSystemEventMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithWaitClassID wait_class_id获取 
func (obj *_AllVirtualSystemEventMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class_id"] = waitClassID })
}

// WithWaitClass# wait_class#获取 
func (obj *_AllVirtualSystemEventMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class#"] = waitClass# })
}

// WithWaitClass wait_class获取 
func (obj *_AllVirtualSystemEventMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["wait_class"] = waitClass })
}

// WithTotalWaits total_waits获取 
func (obj *_AllVirtualSystemEventMgr) WithTotalWaits(totalWaits int64) Option {
	return optionFunc(func(o *options) { o.query["total_waits"] = totalWaits })
}

// WithTotalTimeouts total_timeouts获取 
func (obj *_AllVirtualSystemEventMgr) WithTotalTimeouts(totalTimeouts int64) Option {
	return optionFunc(func(o *options) { o.query["total_timeouts"] = totalTimeouts })
}

// WithTimeWaited time_waited获取 
func (obj *_AllVirtualSystemEventMgr) WithTimeWaited(timeWaited float64) Option {
	return optionFunc(func(o *options) { o.query["time_waited"] = timeWaited })
}

// WithMaxWait max_wait获取 
func (obj *_AllVirtualSystemEventMgr) WithMaxWait(maxWait float64) Option {
	return optionFunc(func(o *options) { o.query["max_wait"] = maxWait })
}

// WithAverageWait average_wait获取 
func (obj *_AllVirtualSystemEventMgr) WithAverageWait(averageWait float64) Option {
	return optionFunc(func(o *options) { o.query["average_wait"] = averageWait })
}

// WithTimeWaitedMicro time_waited_micro获取 
func (obj *_AllVirtualSystemEventMgr) WithTimeWaitedMicro(timeWaitedMicro int64) Option {
	return optionFunc(func(o *options) { o.query["time_waited_micro"] = timeWaitedMicro })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualSystemEventMgr) GetByOption(opts ...Option) (result AllVirtualSystemEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSystemEventMgr) GetByOptions(opts ...Option) (results []*AllVirtualSystemEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualSystemEventMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSystemEvent,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where(options.query)
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
func (obj *_AllVirtualSystemEventMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromEventID 通过event_id获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromEventID(eventID int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`event_id` = ?", eventID).Find(&results).Error
	
	return
}

// GetBatchFromEventID 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromEventID(eventIDs []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`event_id` IN (?)", eventIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过event获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromEvent(event string) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`event` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromEvent(events []string) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`event` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过wait_class_id获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromWaitClassID(waitClassID int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`wait_class_id` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`wait_class_id` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过wait_class#获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromWaitClass#(waitClass# int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`wait_class#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`wait_class#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过wait_class获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromWaitClass(waitClass string) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`wait_class` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromWaitClass(waitClasss []string) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`wait_class` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromTotalWaits 通过total_waits获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromTotalWaits(totalWaits int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`total_waits` = ?", totalWaits).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaits 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromTotalWaits(totalWaitss []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`total_waits` IN (?)", totalWaitss).Find(&results).Error
	
	return
}
 
// GetFromTotalTimeouts 通过total_timeouts获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromTotalTimeouts(totalTimeouts int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`total_timeouts` = ?", totalTimeouts).Find(&results).Error
	
	return
}

// GetBatchFromTotalTimeouts 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromTotalTimeouts(totalTimeoutss []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`total_timeouts` IN (?)", totalTimeoutss).Find(&results).Error
	
	return
}
 
// GetFromTimeWaited 通过time_waited获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromTimeWaited(timeWaited float64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`time_waited` = ?", timeWaited).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaited 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromTimeWaited(timeWaiteds []float64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`time_waited` IN (?)", timeWaiteds).Find(&results).Error
	
	return
}
 
// GetFromMaxWait 通过max_wait获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromMaxWait(maxWait float64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`max_wait` = ?", maxWait).Find(&results).Error
	
	return
}

// GetBatchFromMaxWait 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromMaxWait(maxWaits []float64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`max_wait` IN (?)", maxWaits).Find(&results).Error
	
	return
}
 
// GetFromAverageWait 通过average_wait获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromAverageWait(averageWait float64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`average_wait` = ?", averageWait).Find(&results).Error
	
	return
}

// GetBatchFromAverageWait 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromAverageWait(averageWaits []float64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`average_wait` IN (?)", averageWaits).Find(&results).Error
	
	return
}
 
// GetFromTimeWaitedMicro 通过time_waited_micro获取内容  
func (obj *_AllVirtualSystemEventMgr) GetFromTimeWaitedMicro(timeWaitedMicro int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`time_waited_micro` = ?", timeWaitedMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaitedMicro 批量查找 
func (obj *_AllVirtualSystemEventMgr) GetBatchFromTimeWaitedMicro(timeWaitedMicros []int64) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`time_waited_micro` IN (?)", timeWaitedMicros).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllVirtualSystemEventMgr) FetchByPrimaryKey(tenantID int64 ,svrIP string ,svrPort int64 ,eventID int64 ) (result AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `event_id` = ?", tenantID , svrIP , svrPort , eventID).First(&result).Error
	
	return
}
 

 
 // FetchIndexByAllVirtualSystemEventI1  获取多个内容
 func (obj *_AllVirtualSystemEventMgr) FetchIndexByAllVirtualSystemEventI1(tenantID int64 ) (results []*AllVirtualSystemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSystemEvent{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}
 

	

