package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualSessionEventMgr struct {
	*_BaseMgr
}

// AllVirtualSessionEventMgr open func
func AllVirtualSessionEventMgr(db *gorm.DB) *_AllVirtualSessionEventMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSessionEventMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSessionEventMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_session_event"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSessionEventMgr) GetTableName() string {
	return "__all_virtual_session_event"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSessionEventMgr) Reset() *_AllVirtualSessionEventMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualSessionEventMgr) Get() (result AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSessionEventMgr) Gets() (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSessionEventMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSessionID session_id获取 
func (obj *_AllVirtualSessionEventMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithSvrIP svr_ip获取 
func (obj *_AllVirtualSessionEventMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllVirtualSessionEventMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithEventID event_id获取 
func (obj *_AllVirtualSessionEventMgr) WithEventID(eventID int64) Option {
	return optionFunc(func(o *options) { o.query["event_id"] = eventID })
}

// WithTenantID tenant_id获取 
func (obj *_AllVirtualSessionEventMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithEvent event获取 
func (obj *_AllVirtualSessionEventMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithWaitClassID wait_class_id获取 
func (obj *_AllVirtualSessionEventMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class_id"] = waitClassID })
}

// WithWaitClass# wait_class#获取 
func (obj *_AllVirtualSessionEventMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class#"] = waitClass# })
}

// WithWaitClass wait_class获取 
func (obj *_AllVirtualSessionEventMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["wait_class"] = waitClass })
}

// WithTotalWaits total_waits获取 
func (obj *_AllVirtualSessionEventMgr) WithTotalWaits(totalWaits int64) Option {
	return optionFunc(func(o *options) { o.query["total_waits"] = totalWaits })
}

// WithTotalTimeouts total_timeouts获取 
func (obj *_AllVirtualSessionEventMgr) WithTotalTimeouts(totalTimeouts int64) Option {
	return optionFunc(func(o *options) { o.query["total_timeouts"] = totalTimeouts })
}

// WithTimeWaited time_waited获取 
func (obj *_AllVirtualSessionEventMgr) WithTimeWaited(timeWaited float64) Option {
	return optionFunc(func(o *options) { o.query["time_waited"] = timeWaited })
}

// WithMaxWait max_wait获取 
func (obj *_AllVirtualSessionEventMgr) WithMaxWait(maxWait float64) Option {
	return optionFunc(func(o *options) { o.query["max_wait"] = maxWait })
}

// WithAverageWait average_wait获取 
func (obj *_AllVirtualSessionEventMgr) WithAverageWait(averageWait float64) Option {
	return optionFunc(func(o *options) { o.query["average_wait"] = averageWait })
}

// WithTimeWaitedMicro time_waited_micro获取 
func (obj *_AllVirtualSessionEventMgr) WithTimeWaitedMicro(timeWaitedMicro int64) Option {
	return optionFunc(func(o *options) { o.query["time_waited_micro"] = timeWaitedMicro })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualSessionEventMgr) GetByOption(opts ...Option) (result AllVirtualSessionEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSessionEventMgr) GetByOptions(opts ...Option) (results []*AllVirtualSessionEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualSessionEventMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSessionEvent,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where(options.query)
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
func (obj *_AllVirtualSessionEventMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromEventID 通过event_id获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromEventID(eventID int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`event_id` = ?", eventID).Find(&results).Error
	
	return
}

// GetBatchFromEventID 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromEventID(eventIDs []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`event_id` IN (?)", eventIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过event获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromEvent(event string) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`event` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromEvent(events []string) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`event` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过wait_class_id获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromWaitClassID(waitClassID int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`wait_class_id` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`wait_class_id` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过wait_class#获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromWaitClass#(waitClass# int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`wait_class#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`wait_class#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过wait_class获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromWaitClass(waitClass string) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`wait_class` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromWaitClass(waitClasss []string) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`wait_class` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromTotalWaits 通过total_waits获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromTotalWaits(totalWaits int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`total_waits` = ?", totalWaits).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaits 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromTotalWaits(totalWaitss []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`total_waits` IN (?)", totalWaitss).Find(&results).Error
	
	return
}
 
// GetFromTotalTimeouts 通过total_timeouts获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromTotalTimeouts(totalTimeouts int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`total_timeouts` = ?", totalTimeouts).Find(&results).Error
	
	return
}

// GetBatchFromTotalTimeouts 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromTotalTimeouts(totalTimeoutss []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`total_timeouts` IN (?)", totalTimeoutss).Find(&results).Error
	
	return
}
 
// GetFromTimeWaited 通过time_waited获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromTimeWaited(timeWaited float64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`time_waited` = ?", timeWaited).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaited 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromTimeWaited(timeWaiteds []float64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`time_waited` IN (?)", timeWaiteds).Find(&results).Error
	
	return
}
 
// GetFromMaxWait 通过max_wait获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromMaxWait(maxWait float64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`max_wait` = ?", maxWait).Find(&results).Error
	
	return
}

// GetBatchFromMaxWait 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromMaxWait(maxWaits []float64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`max_wait` IN (?)", maxWaits).Find(&results).Error
	
	return
}
 
// GetFromAverageWait 通过average_wait获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromAverageWait(averageWait float64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`average_wait` = ?", averageWait).Find(&results).Error
	
	return
}

// GetBatchFromAverageWait 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromAverageWait(averageWaits []float64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`average_wait` IN (?)", averageWaits).Find(&results).Error
	
	return
}
 
// GetFromTimeWaitedMicro 通过time_waited_micro获取内容  
func (obj *_AllVirtualSessionEventMgr) GetFromTimeWaitedMicro(timeWaitedMicro int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`time_waited_micro` = ?", timeWaitedMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaitedMicro 批量查找 
func (obj *_AllVirtualSessionEventMgr) GetBatchFromTimeWaitedMicro(timeWaitedMicros []int64) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`time_waited_micro` IN (?)", timeWaitedMicros).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllVirtualSessionEventMgr) FetchByPrimaryKey(sessionID int64 ,svrIP string ,svrPort int64 ,eventID int64 ) (result AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`session_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `event_id` = ?", sessionID , svrIP , svrPort , eventID).First(&result).Error
	
	return
}
 

 
 // FetchIndexByAllVirtualSessionEventI1  获取多个内容
 func (obj *_AllVirtualSessionEventMgr) FetchIndexByAllVirtualSessionEventI1(sessionID int64 ) (results []*AllVirtualSessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionEvent{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}
 

	

