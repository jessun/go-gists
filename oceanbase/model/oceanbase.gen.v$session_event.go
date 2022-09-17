package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sessionEventMgr struct {
	*_BaseMgr
}

// V$sessionEventMgr open func
func V$sessionEventMgr(db *gorm.DB) *_V$sessionEventMgr {
	if db == nil {
		panic(fmt.Errorf("V$sessionEventMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sessionEventMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$session_event"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sessionEventMgr) GetTableName() string {
	return "v$session_event"
}

// Reset 重置gorm会话
func (obj *_V$sessionEventMgr) Reset() *_V$sessionEventMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sessionEventMgr) Get() (result V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sessionEventMgr) Gets() (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sessionEventMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSid SID获取 
func (obj *_V$sessionEventMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithConID CON_ID获取 
func (obj *_V$sessionEventMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithEventID EVENT_ID获取 
func (obj *_V$sessionEventMgr) WithEventID(eventID int64) Option {
	return optionFunc(func(o *options) { o.query["EVENT_ID"] = eventID })
}

// WithEvent EVENT获取 
func (obj *_V$sessionEventMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["EVENT"] = event })
}

// WithWaitClassID WAIT_CLASS_ID获取 
func (obj *_V$sessionEventMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS_ID"] = waitClassID })
}

// WithWaitClass# WAIT_CLASS#获取 
func (obj *_V$sessionEventMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS#"] = waitClass# })
}

// WithWaitClass WAIT_CLASS获取 
func (obj *_V$sessionEventMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS"] = waitClass })
}

// WithTotalWaits TOTAL_WAITS获取 
func (obj *_V$sessionEventMgr) WithTotalWaits(totalWaits int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_WAITS"] = totalWaits })
}

// WithTotalTimeouts TOTAL_TIMEOUTS获取 
func (obj *_V$sessionEventMgr) WithTotalTimeouts(totalTimeouts int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_TIMEOUTS"] = totalTimeouts })
}

// WithTimeWaited TIME_WAITED获取 
func (obj *_V$sessionEventMgr) WithTimeWaited(timeWaited float64) Option {
	return optionFunc(func(o *options) { o.query["TIME_WAITED"] = timeWaited })
}

// WithMaxWait MAX_WAIT获取 
func (obj *_V$sessionEventMgr) WithMaxWait(maxWait float64) Option {
	return optionFunc(func(o *options) { o.query["MAX_WAIT"] = maxWait })
}

// WithAverageWait AVERAGE_WAIT获取 
func (obj *_V$sessionEventMgr) WithAverageWait(averageWait float64) Option {
	return optionFunc(func(o *options) { o.query["AVERAGE_WAIT"] = averageWait })
}

// WithTimeWaitedMicro TIME_WAITED_MICRO获取 
func (obj *_V$sessionEventMgr) WithTimeWaitedMicro(timeWaitedMicro int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_WAITED_MICRO"] = timeWaitedMicro })
}


// GetByOption 功能选项模式获取
func (obj *_V$sessionEventMgr) GetByOption(opts ...Option) (result V$sessionEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sessionEventMgr) GetByOptions(opts ...Option) (results []*V$sessionEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sessionEventMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sessionEvent,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where(options.query)
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


// GetFromSid 通过SID获取内容  
func (obj *_V$sessionEventMgr) GetFromSid(sid int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromSid(sids []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromConID 通过CON_ID获取内容  
func (obj *_V$sessionEventMgr) GetFromConID(conID int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromConID(conIDs []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromEventID 通过EVENT_ID获取内容  
func (obj *_V$sessionEventMgr) GetFromEventID(eventID int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`EVENT_ID` = ?", eventID).Find(&results).Error
	
	return
}

// GetBatchFromEventID 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromEventID(eventIDs []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`EVENT_ID` IN (?)", eventIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过EVENT获取内容  
func (obj *_V$sessionEventMgr) GetFromEvent(event string) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`EVENT` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromEvent(events []string) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`EVENT` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过WAIT_CLASS_ID获取内容  
func (obj *_V$sessionEventMgr) GetFromWaitClassID(waitClassID int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`WAIT_CLASS_ID` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`WAIT_CLASS_ID` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过WAIT_CLASS#获取内容  
func (obj *_V$sessionEventMgr) GetFromWaitClass#(waitClass# int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`WAIT_CLASS#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`WAIT_CLASS#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过WAIT_CLASS获取内容  
func (obj *_V$sessionEventMgr) GetFromWaitClass(waitClass string) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`WAIT_CLASS` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromWaitClass(waitClasss []string) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`WAIT_CLASS` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromTotalWaits 通过TOTAL_WAITS获取内容  
func (obj *_V$sessionEventMgr) GetFromTotalWaits(totalWaits int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TOTAL_WAITS` = ?", totalWaits).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaits 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromTotalWaits(totalWaitss []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TOTAL_WAITS` IN (?)", totalWaitss).Find(&results).Error
	
	return
}
 
// GetFromTotalTimeouts 通过TOTAL_TIMEOUTS获取内容  
func (obj *_V$sessionEventMgr) GetFromTotalTimeouts(totalTimeouts int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TOTAL_TIMEOUTS` = ?", totalTimeouts).Find(&results).Error
	
	return
}

// GetBatchFromTotalTimeouts 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromTotalTimeouts(totalTimeoutss []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TOTAL_TIMEOUTS` IN (?)", totalTimeoutss).Find(&results).Error
	
	return
}
 
// GetFromTimeWaited 通过TIME_WAITED获取内容  
func (obj *_V$sessionEventMgr) GetFromTimeWaited(timeWaited float64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TIME_WAITED` = ?", timeWaited).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaited 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromTimeWaited(timeWaiteds []float64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TIME_WAITED` IN (?)", timeWaiteds).Find(&results).Error
	
	return
}
 
// GetFromMaxWait 通过MAX_WAIT获取内容  
func (obj *_V$sessionEventMgr) GetFromMaxWait(maxWait float64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`MAX_WAIT` = ?", maxWait).Find(&results).Error
	
	return
}

// GetBatchFromMaxWait 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromMaxWait(maxWaits []float64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`MAX_WAIT` IN (?)", maxWaits).Find(&results).Error
	
	return
}
 
// GetFromAverageWait 通过AVERAGE_WAIT获取内容  
func (obj *_V$sessionEventMgr) GetFromAverageWait(averageWait float64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`AVERAGE_WAIT` = ?", averageWait).Find(&results).Error
	
	return
}

// GetBatchFromAverageWait 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromAverageWait(averageWaits []float64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`AVERAGE_WAIT` IN (?)", averageWaits).Find(&results).Error
	
	return
}
 
// GetFromTimeWaitedMicro 通过TIME_WAITED_MICRO获取内容  
func (obj *_V$sessionEventMgr) GetFromTimeWaitedMicro(timeWaitedMicro int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TIME_WAITED_MICRO` = ?", timeWaitedMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeWaitedMicro 批量查找 
func (obj *_V$sessionEventMgr) GetBatchFromTimeWaitedMicro(timeWaitedMicros []int64) (results []*V$sessionEvent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionEvent{}).Where("`TIME_WAITED_MICRO` IN (?)", timeWaitedMicros).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

