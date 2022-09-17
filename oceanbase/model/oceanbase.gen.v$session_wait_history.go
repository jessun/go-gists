package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sessionWaitHistoryMgr struct {
	*_BaseMgr
}

// V$sessionWaitHistoryMgr open func
func V$sessionWaitHistoryMgr(db *gorm.DB) *_V$sessionWaitHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("V$sessionWaitHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sessionWaitHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$session_wait_history"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sessionWaitHistoryMgr) GetTableName() string {
	return "v$session_wait_history"
}

// Reset 重置gorm会话
func (obj *_V$sessionWaitHistoryMgr) Reset() *_V$sessionWaitHistoryMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sessionWaitHistoryMgr) Get() (result V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sessionWaitHistoryMgr) Gets() (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sessionWaitHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSid SID获取 
func (obj *_V$sessionWaitHistoryMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithConID CON_ID获取 
func (obj *_V$sessionWaitHistoryMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSeq# SEQ#获取 
func (obj *_V$sessionWaitHistoryMgr) WithSeq#(seq# int64) Option {
	return optionFunc(func(o *options) { o.query["SEQ#"] = seq# })
}

// WithEvent# EVENT#获取 
func (obj *_V$sessionWaitHistoryMgr) WithEvent#(event# int64) Option {
	return optionFunc(func(o *options) { o.query["EVENT#"] = event# })
}

// WithEvent EVENT获取 
func (obj *_V$sessionWaitHistoryMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["EVENT"] = event })
}

// WithP1text P1TEXT获取 
func (obj *_V$sessionWaitHistoryMgr) WithP1text(p1text string) Option {
	return optionFunc(func(o *options) { o.query["P1TEXT"] = p1text })
}

// WithP1 P1获取 
func (obj *_V$sessionWaitHistoryMgr) WithP1(p1 uint64) Option {
	return optionFunc(func(o *options) { o.query["P1"] = p1 })
}

// WithP2text P2TEXT获取 
func (obj *_V$sessionWaitHistoryMgr) WithP2text(p2text string) Option {
	return optionFunc(func(o *options) { o.query["P2TEXT"] = p2text })
}

// WithP2 P2获取 
func (obj *_V$sessionWaitHistoryMgr) WithP2(p2 uint64) Option {
	return optionFunc(func(o *options) { o.query["P2"] = p2 })
}

// WithP3text P3TEXT获取 
func (obj *_V$sessionWaitHistoryMgr) WithP3text(p3text string) Option {
	return optionFunc(func(o *options) { o.query["P3TEXT"] = p3text })
}

// WithP3 P3获取 
func (obj *_V$sessionWaitHistoryMgr) WithP3(p3 uint64) Option {
	return optionFunc(func(o *options) { o.query["P3"] = p3 })
}

// WithWaitTimeMicro WAIT_TIME_MICRO获取 
func (obj *_V$sessionWaitHistoryMgr) WithWaitTimeMicro(waitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TIME_MICRO"] = waitTimeMicro })
}

// WithTimeSinceLastWaitMicro TIME_SINCE_LAST_WAIT_MICRO获取 
func (obj *_V$sessionWaitHistoryMgr) WithTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_SINCE_LAST_WAIT_MICRO"] = timeSinceLastWaitMicro })
}

// WithWaitTime WAIT_TIME获取 
func (obj *_V$sessionWaitHistoryMgr) WithWaitTime(waitTime float64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TIME"] = waitTime })
}


// GetByOption 功能选项模式获取
func (obj *_V$sessionWaitHistoryMgr) GetByOption(opts ...Option) (result V$sessionWaitHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sessionWaitHistoryMgr) GetByOptions(opts ...Option) (results []*V$sessionWaitHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sessionWaitHistoryMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sessionWaitHistory,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where(options.query)
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
func (obj *_V$sessionWaitHistoryMgr) GetFromSid(sid int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromSid(sids []int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromConID 通过CON_ID获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromConID(conID int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromConID(conIDs []int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSeq# 通过SEQ#获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromSeq#(seq# int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`SEQ#` = ?", seq#).Find(&results).Error
	
	return
}

// GetBatchFromSeq# 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromSeq#(seq#s []int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`SEQ#` IN (?)", seq#s).Find(&results).Error
	
	return
}
 
// GetFromEvent# 通过EVENT#获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromEvent#(event# int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`EVENT#` = ?", event#).Find(&results).Error
	
	return
}

// GetBatchFromEvent# 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromEvent#(event#s []int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`EVENT#` IN (?)", event#s).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过EVENT获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromEvent(event string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`EVENT` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromEvent(events []string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`EVENT` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromP1text 通过P1TEXT获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromP1text(p1text string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P1TEXT` = ?", p1text).Find(&results).Error
	
	return
}

// GetBatchFromP1text 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromP1text(p1texts []string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P1TEXT` IN (?)", p1texts).Find(&results).Error
	
	return
}
 
// GetFromP1 通过P1获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromP1(p1 uint64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P1` = ?", p1).Find(&results).Error
	
	return
}

// GetBatchFromP1 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromP1(p1s []uint64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P1` IN (?)", p1s).Find(&results).Error
	
	return
}
 
// GetFromP2text 通过P2TEXT获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromP2text(p2text string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P2TEXT` = ?", p2text).Find(&results).Error
	
	return
}

// GetBatchFromP2text 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromP2text(p2texts []string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P2TEXT` IN (?)", p2texts).Find(&results).Error
	
	return
}
 
// GetFromP2 通过P2获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromP2(p2 uint64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P2` = ?", p2).Find(&results).Error
	
	return
}

// GetBatchFromP2 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromP2(p2s []uint64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P2` IN (?)", p2s).Find(&results).Error
	
	return
}
 
// GetFromP3text 通过P3TEXT获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromP3text(p3text string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P3TEXT` = ?", p3text).Find(&results).Error
	
	return
}

// GetBatchFromP3text 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromP3text(p3texts []string) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P3TEXT` IN (?)", p3texts).Find(&results).Error
	
	return
}
 
// GetFromP3 通过P3获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromP3(p3 uint64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P3` = ?", p3).Find(&results).Error
	
	return
}

// GetBatchFromP3 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromP3(p3s []uint64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`P3` IN (?)", p3s).Find(&results).Error
	
	return
}
 
// GetFromWaitTimeMicro 通过WAIT_TIME_MICRO获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromWaitTimeMicro(waitTimeMicro int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`WAIT_TIME_MICRO` = ?", waitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromWaitTimeMicro 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromWaitTimeMicro(waitTimeMicros []int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`WAIT_TIME_MICRO` IN (?)", waitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTimeSinceLastWaitMicro 通过TIME_SINCE_LAST_WAIT_MICRO获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`TIME_SINCE_LAST_WAIT_MICRO` = ?", timeSinceLastWaitMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeSinceLastWaitMicro 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromTimeSinceLastWaitMicro(timeSinceLastWaitMicros []int64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`TIME_SINCE_LAST_WAIT_MICRO` IN (?)", timeSinceLastWaitMicros).Find(&results).Error
	
	return
}
 
// GetFromWaitTime 通过WAIT_TIME获取内容  
func (obj *_V$sessionWaitHistoryMgr) GetFromWaitTime(waitTime float64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`WAIT_TIME` = ?", waitTime).Find(&results).Error
	
	return
}

// GetBatchFromWaitTime 批量查找 
func (obj *_V$sessionWaitHistoryMgr) GetBatchFromWaitTime(waitTimes []float64) (results []*V$sessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionWaitHistory{}).Where("`WAIT_TIME` IN (?)", waitTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

