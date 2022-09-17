package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sessionWaitMgr struct {
	*_BaseMgr
}

// V$sessionWaitMgr open func
func V$sessionWaitMgr(db *gorm.DB) *_V$sessionWaitMgr {
	if db == nil {
		panic(fmt.Errorf("V$sessionWaitMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sessionWaitMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$session_wait"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sessionWaitMgr) GetTableName() string {
	return "v$session_wait"
}

// Reset 重置gorm会话
func (obj *_V$sessionWaitMgr) Reset() *_V$sessionWaitMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sessionWaitMgr) Get() (result V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sessionWaitMgr) Gets() (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sessionWaitMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSid SID获取 
func (obj *_V$sessionWaitMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithConID CON_ID获取 
func (obj *_V$sessionWaitMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithEvent EVENT获取 
func (obj *_V$sessionWaitMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["EVENT"] = event })
}

// WithP1text P1TEXT获取 
func (obj *_V$sessionWaitMgr) WithP1text(p1text string) Option {
	return optionFunc(func(o *options) { o.query["P1TEXT"] = p1text })
}

// WithP1 P1获取 
func (obj *_V$sessionWaitMgr) WithP1(p1 uint64) Option {
	return optionFunc(func(o *options) { o.query["P1"] = p1 })
}

// WithP2text P2TEXT获取 
func (obj *_V$sessionWaitMgr) WithP2text(p2text string) Option {
	return optionFunc(func(o *options) { o.query["P2TEXT"] = p2text })
}

// WithP2 P2获取 
func (obj *_V$sessionWaitMgr) WithP2(p2 uint64) Option {
	return optionFunc(func(o *options) { o.query["P2"] = p2 })
}

// WithP3text P3TEXT获取 
func (obj *_V$sessionWaitMgr) WithP3text(p3text string) Option {
	return optionFunc(func(o *options) { o.query["P3TEXT"] = p3text })
}

// WithP3 P3获取 
func (obj *_V$sessionWaitMgr) WithP3(p3 uint64) Option {
	return optionFunc(func(o *options) { o.query["P3"] = p3 })
}

// WithWaitClassID WAIT_CLASS_ID获取 
func (obj *_V$sessionWaitMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS_ID"] = waitClassID })
}

// WithWaitClass# WAIT_CLASS#获取 
func (obj *_V$sessionWaitMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS#"] = waitClass# })
}

// WithWaitClass WAIT_CLASS获取 
func (obj *_V$sessionWaitMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS"] = waitClass })
}

// WithState STATE获取 
func (obj *_V$sessionWaitMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["STATE"] = state })
}

// WithWaitTimeMicro WAIT_TIME_MICRO获取 
func (obj *_V$sessionWaitMgr) WithWaitTimeMicro(waitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TIME_MICRO"] = waitTimeMicro })
}

// WithTimeRemainingMicro TIME_REMAINING_MICRO获取 
func (obj *_V$sessionWaitMgr) WithTimeRemainingMicro(timeRemainingMicro int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_REMAINING_MICRO"] = timeRemainingMicro })
}

// WithTimeSinceLastWaitMicro TIME_SINCE_LAST_WAIT_MICRO获取 
func (obj *_V$sessionWaitMgr) WithTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_SINCE_LAST_WAIT_MICRO"] = timeSinceLastWaitMicro })
}


// GetByOption 功能选项模式获取
func (obj *_V$sessionWaitMgr) GetByOption(opts ...Option) (result V$sessionWait, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sessionWaitMgr) GetByOptions(opts ...Option) (results []*V$sessionWait, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sessionWaitMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sessionWait,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where(options.query)
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
func (obj *_V$sessionWaitMgr) GetFromSid(sid int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromSid(sids []int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromConID 通过CON_ID获取内容  
func (obj *_V$sessionWaitMgr) GetFromConID(conID int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromConID(conIDs []int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过EVENT获取内容  
func (obj *_V$sessionWaitMgr) GetFromEvent(event string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`EVENT` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromEvent(events []string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`EVENT` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromP1text 通过P1TEXT获取内容  
func (obj *_V$sessionWaitMgr) GetFromP1text(p1text string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P1TEXT` = ?", p1text).Find(&results).Error
	
	return
}

// GetBatchFromP1text 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromP1text(p1texts []string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P1TEXT` IN (?)", p1texts).Find(&results).Error
	
	return
}
 
// GetFromP1 通过P1获取内容  
func (obj *_V$sessionWaitMgr) GetFromP1(p1 uint64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P1` = ?", p1).Find(&results).Error
	
	return
}

// GetBatchFromP1 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromP1(p1s []uint64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P1` IN (?)", p1s).Find(&results).Error
	
	return
}
 
// GetFromP2text 通过P2TEXT获取内容  
func (obj *_V$sessionWaitMgr) GetFromP2text(p2text string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P2TEXT` = ?", p2text).Find(&results).Error
	
	return
}

// GetBatchFromP2text 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromP2text(p2texts []string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P2TEXT` IN (?)", p2texts).Find(&results).Error
	
	return
}
 
// GetFromP2 通过P2获取内容  
func (obj *_V$sessionWaitMgr) GetFromP2(p2 uint64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P2` = ?", p2).Find(&results).Error
	
	return
}

// GetBatchFromP2 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromP2(p2s []uint64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P2` IN (?)", p2s).Find(&results).Error
	
	return
}
 
// GetFromP3text 通过P3TEXT获取内容  
func (obj *_V$sessionWaitMgr) GetFromP3text(p3text string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P3TEXT` = ?", p3text).Find(&results).Error
	
	return
}

// GetBatchFromP3text 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromP3text(p3texts []string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P3TEXT` IN (?)", p3texts).Find(&results).Error
	
	return
}
 
// GetFromP3 通过P3获取内容  
func (obj *_V$sessionWaitMgr) GetFromP3(p3 uint64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P3` = ?", p3).Find(&results).Error
	
	return
}

// GetBatchFromP3 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromP3(p3s []uint64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`P3` IN (?)", p3s).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过WAIT_CLASS_ID获取内容  
func (obj *_V$sessionWaitMgr) GetFromWaitClassID(waitClassID int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_CLASS_ID` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_CLASS_ID` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过WAIT_CLASS#获取内容  
func (obj *_V$sessionWaitMgr) GetFromWaitClass#(waitClass# int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_CLASS#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_CLASS#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过WAIT_CLASS获取内容  
func (obj *_V$sessionWaitMgr) GetFromWaitClass(waitClass string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_CLASS` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromWaitClass(waitClasss []string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_CLASS` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromState 通过STATE获取内容  
func (obj *_V$sessionWaitMgr) GetFromState(state string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`STATE` = ?", state).Find(&results).Error
	
	return
}

// GetBatchFromState 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromState(states []string) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`STATE` IN (?)", states).Find(&results).Error
	
	return
}
 
// GetFromWaitTimeMicro 通过WAIT_TIME_MICRO获取内容  
func (obj *_V$sessionWaitMgr) GetFromWaitTimeMicro(waitTimeMicro int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_TIME_MICRO` = ?", waitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromWaitTimeMicro 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromWaitTimeMicro(waitTimeMicros []int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`WAIT_TIME_MICRO` IN (?)", waitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTimeRemainingMicro 通过TIME_REMAINING_MICRO获取内容  
func (obj *_V$sessionWaitMgr) GetFromTimeRemainingMicro(timeRemainingMicro int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`TIME_REMAINING_MICRO` = ?", timeRemainingMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeRemainingMicro 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromTimeRemainingMicro(timeRemainingMicros []int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`TIME_REMAINING_MICRO` IN (?)", timeRemainingMicros).Find(&results).Error
	
	return
}
 
// GetFromTimeSinceLastWaitMicro 通过TIME_SINCE_LAST_WAIT_MICRO获取内容  
func (obj *_V$sessionWaitMgr) GetFromTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`TIME_SINCE_LAST_WAIT_MICRO` = ?", timeSinceLastWaitMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeSinceLastWaitMicro 批量查找 
func (obj *_V$sessionWaitMgr) GetBatchFromTimeSinceLastWaitMicro(timeSinceLastWaitMicros []int64) (results []*V$sessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWait{}).Where("`TIME_SINCE_LAST_WAIT_MICRO` IN (?)", timeSinceLastWaitMicros).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

