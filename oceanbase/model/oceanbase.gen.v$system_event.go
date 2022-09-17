package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$systemEventMgr struct {
	*_BaseMgr
}

// V$systemEventMgr open func
func V$systemEventMgr(db *gorm.DB) *_V$systemEventMgr {
	if db == nil {
		panic(fmt.Errorf("V$systemEventMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$systemEventMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$system_event"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$systemEventMgr) GetTableName() string {
	return "v$system_event"
}

// Reset 重置gorm会话
func (obj *_V$systemEventMgr) Reset() *_V$systemEventMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$systemEventMgr) Get() (result V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$systemEventMgr) Gets() (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$systemEventMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$systemEventMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithEventID EVENT_ID获取 
func (obj *_V$systemEventMgr) WithEventID(eventID int64) Option {
	return optionFunc(func(o *options) { o.query["EVENT_ID"] = eventID })
}

// WithEvent EVENT获取 
func (obj *_V$systemEventMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["EVENT"] = event })
}

// WithWaitClassID WAIT_CLASS_ID获取 
func (obj *_V$systemEventMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS_ID"] = waitClassID })
}

// WithWaitClass# WAIT_CLASS#获取 
func (obj *_V$systemEventMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS#"] = waitClass# })
}

// WithWaitClass WAIT_CLASS获取 
func (obj *_V$systemEventMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS"] = waitClass })
}

// WithTotalWaits TOTAL_WAITS获取 
func (obj *_V$systemEventMgr) WithTotalWaits(totalWaits int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_WAITS"] = totalWaits })
}

// WithTotalTimeouts TOTAL_TIMEOUTS获取 
func (obj *_V$systemEventMgr) WithTotalTimeouts(totalTimeouts int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_TIMEOUTS"] = totalTimeouts })
}

// WithTimeWaited TIME_WAITED获取 
func (obj *_V$systemEventMgr) WithTimeWaited(timeWaited float64) Option {
	return optionFunc(func(o *options) { o.query["TIME_WAITED"] = timeWaited })
}

// WithAverageWait AVERAGE_WAIT获取 
func (obj *_V$systemEventMgr) WithAverageWait(averageWait float64) Option {
	return optionFunc(func(o *options) { o.query["AVERAGE_WAIT"] = averageWait })
}

// WithTimeWaitedMicro TIME_WAITED_MICRO获取 
func (obj *_V$systemEventMgr) WithTimeWaitedMicro(timeWaitedMicro int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_WAITED_MICRO"] = timeWaitedMicro })
}


// GetByOption 功能选项模式获取
func (obj *_V$systemEventMgr) GetByOption(opts ...Option) (result V$systemEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$systemEventMgr) GetByOptions(opts ...Option) (results []*V$systemEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$systemEventMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$systemEvent,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where(options.query)
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


// GetFromConID 通过CON_ID获取内容  
func (obj *_V$systemEventMgr) GetFromConID(conID int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromConID(conIDs []int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromEventID 通过EVENT_ID获取内容  
func (obj *_V$systemEventMgr) GetFromEventID(eventID int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`EVENT_ID` = ?", eventID).Find(&results).Error
	
	return
}

// GetBatchFromEventID 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromEventID(eventIDs []int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`EVENT_ID` IN (?)", eventIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过EVENT获取内容  
func (obj *_V$systemEventMgr) GetFromEvent(event string) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`EVENT` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromEvent(events []string) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`EVENT` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过WAIT_CLASS_ID获取内容  
func (obj *_V$systemEventMgr) GetFromWaitClassID(waitClassID int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`WAIT_CLASS_ID` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`WAIT_CLASS_ID` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过WAIT_CLASS#获取内容  
func (obj *_V$systemEventMgr) GetFromWaitClass#(waitClass# int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`WAIT_CLASS#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`WAIT_CLASS#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过WAIT_CLASS获取内容  
func (obj *_V$systemEventMgr) GetFromWaitClass(waitClass string) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`WAIT_CLASS` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromWaitClass(waitClasss []string) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`WAIT_CLASS` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromTotalWaits 通过TOTAL_WAITS获取内容  
func (obj *_V$systemEventMgr) GetFromTotalWaits(totalWaits int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TOTAL_WAITS` = ?", totalWaits).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaits 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromTotalWaits(totalWaitss []int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TOTAL_WAITS` IN (?)", totalWaitss).Find(&results).Error
	
	return
}
 
// GetFromTotalTimeouts 通过TOTAL_TIMEOUTS获取内容  
func (obj *_V$systemEventMgr) GetFromTotalTimeouts(totalTimeouts int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TOTAL_TIMEOUTS` = ?", totalTimeouts).Find(&results).Error
	
	return
}

// GetBatchFromTotalTimeouts 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromTotalTimeouts(totalTimeoutss []int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TOTAL_TIMEOUTS` IN (?)", totalTimeoutss).Find(&results).Error
	
	return
}
 
// GetFromTimeWaited 通过TIME_WAITED获取内容  
func (obj *_V$systemEventMgr) GetFromTimeWaited(timeWaited float64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TIME_WAITED` = ?", timeWaited).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaited 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromTimeWaited(timeWaiteds []float64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TIME_WAITED` IN (?)", timeWaiteds).Find(&results).Error
	
	return
}
 
// GetFromAverageWait 通过AVERAGE_WAIT获取内容  
func (obj *_V$systemEventMgr) GetFromAverageWait(averageWait float64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`AVERAGE_WAIT` = ?", averageWait).Find(&results).Error
	
	return
}

// GetBatchFromAverageWait 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromAverageWait(averageWaits []float64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`AVERAGE_WAIT` IN (?)", averageWaits).Find(&results).Error
	
	return
}
 
// GetFromTimeWaitedMicro 通过TIME_WAITED_MICRO获取内容  
func (obj *_V$systemEventMgr) GetFromTimeWaitedMicro(timeWaitedMicro int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TIME_WAITED_MICRO` = ?", timeWaitedMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaitedMicro 批量查找 
func (obj *_V$systemEventMgr) GetBatchFromTimeWaitedMicro(timeWaitedMicros []int64) (results []*V$systemEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$systemEvent{}).Where("`TIME_WAITED_MICRO` IN (?)", timeWaitedMicros).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

