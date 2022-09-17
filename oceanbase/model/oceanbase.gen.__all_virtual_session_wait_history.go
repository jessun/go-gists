package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualSessionWaitHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualSessionWaitHistoryMgr open func
func AllVirtualSessionWaitHistoryMgr(db *gorm.DB) *_AllVirtualSessionWaitHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSessionWaitHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSessionWaitHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_session_wait_history"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSessionWaitHistoryMgr) GetTableName() string {
	return "__all_virtual_session_wait_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSessionWaitHistoryMgr) Reset() *_AllVirtualSessionWaitHistoryMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) Get() (result AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSessionWaitHistoryMgr) Gets() (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSessionWaitHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSessionID session_id获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithSvrIP svr_ip获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSeq# seq#获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithSeq#(seq# int64) Option {
	return optionFunc(func(o *options) { o.query["seq#"] = seq# })
}

// WithTenantID tenant_id获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithEvent# event#获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithEvent#(event# int64) Option {
	return optionFunc(func(o *options) { o.query["event#"] = event# })
}

// WithEvent event获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithP1text p1text获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithP1text(p1text string) Option {
	return optionFunc(func(o *options) { o.query["p1text"] = p1text })
}

// WithP1 p1获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithP1(p1 uint64) Option {
	return optionFunc(func(o *options) { o.query["p1"] = p1 })
}

// WithP2text p2text获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithP2text(p2text string) Option {
	return optionFunc(func(o *options) { o.query["p2text"] = p2text })
}

// WithP2 p2获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithP2(p2 uint64) Option {
	return optionFunc(func(o *options) { o.query["p2"] = p2 })
}

// WithP3text p3text获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithP3text(p3text string) Option {
	return optionFunc(func(o *options) { o.query["p3text"] = p3text })
}

// WithP3 p3获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithP3(p3 uint64) Option {
	return optionFunc(func(o *options) { o.query["p3"] = p3 })
}

// WithLevel level获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithLevel(level int64) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithWaitTimeMicro wait_time_micro获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithWaitTimeMicro(waitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["wait_time_micro"] = waitTimeMicro })
}

// WithTimeSinceLastWaitMicro time_since_last_wait_micro获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) Option {
	return optionFunc(func(o *options) { o.query["time_since_last_wait_micro"] = timeSinceLastWaitMicro })
}

// WithWaitTime wait_time获取 
func (obj *_AllVirtualSessionWaitHistoryMgr) WithWaitTime(waitTime float64) Option {
	return optionFunc(func(o *options) { o.query["wait_time"] = waitTime })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualSessionWaitHistoryMgr) GetByOption(opts ...Option) (result AllVirtualSessionWaitHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSessionWaitHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualSessionWaitHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualSessionWaitHistoryMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSessionWaitHistory,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where(options.query)
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
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSeq# 通过seq#获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromSeq#(seq# int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`seq#` = ?", seq#).Find(&results).Error
	
	return
}

// GetBatchFromSeq# 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromSeq#(seq#s []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`seq#` IN (?)", seq#s).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent# 通过event#获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromEvent#(event# int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`event#` = ?", event#).Find(&results).Error
	
	return
}

// GetBatchFromEvent# 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromEvent#(event#s []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`event#` IN (?)", event#s).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过event获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromEvent(event string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`event` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromEvent(events []string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`event` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromP1text 通过p1text获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromP1text(p1text string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p1text` = ?", p1text).Find(&results).Error
	
	return
}

// GetBatchFromP1text 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromP1text(p1texts []string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p1text` IN (?)", p1texts).Find(&results).Error
	
	return
}
 
// GetFromP1 通过p1获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromP1(p1 uint64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p1` = ?", p1).Find(&results).Error
	
	return
}

// GetBatchFromP1 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromP1(p1s []uint64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p1` IN (?)", p1s).Find(&results).Error
	
	return
}
 
// GetFromP2text 通过p2text获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromP2text(p2text string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p2text` = ?", p2text).Find(&results).Error
	
	return
}

// GetBatchFromP2text 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromP2text(p2texts []string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p2text` IN (?)", p2texts).Find(&results).Error
	
	return
}
 
// GetFromP2 通过p2获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromP2(p2 uint64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p2` = ?", p2).Find(&results).Error
	
	return
}

// GetBatchFromP2 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromP2(p2s []uint64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p2` IN (?)", p2s).Find(&results).Error
	
	return
}
 
// GetFromP3text 通过p3text获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromP3text(p3text string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p3text` = ?", p3text).Find(&results).Error
	
	return
}

// GetBatchFromP3text 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromP3text(p3texts []string) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p3text` IN (?)", p3texts).Find(&results).Error
	
	return
}
 
// GetFromP3 通过p3获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromP3(p3 uint64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p3` = ?", p3).Find(&results).Error
	
	return
}

// GetBatchFromP3 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromP3(p3s []uint64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`p3` IN (?)", p3s).Find(&results).Error
	
	return
}
 
// GetFromLevel 通过level获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromLevel(level int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`level` = ?", level).Find(&results).Error
	
	return
}

// GetBatchFromLevel 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromLevel(levels []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`level` IN (?)", levels).Find(&results).Error
	
	return
}
 
// GetFromWaitTimeMicro 通过wait_time_micro获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromWaitTimeMicro(waitTimeMicro int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`wait_time_micro` = ?", waitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromWaitTimeMicro 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromWaitTimeMicro(waitTimeMicros []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`wait_time_micro` IN (?)", waitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTimeSinceLastWaitMicro 通过time_since_last_wait_micro获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`time_since_last_wait_micro` = ?", timeSinceLastWaitMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeSinceLastWaitMicro 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromTimeSinceLastWaitMicro(timeSinceLastWaitMicros []int64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`time_since_last_wait_micro` IN (?)", timeSinceLastWaitMicros).Find(&results).Error
	
	return
}
 
// GetFromWaitTime 通过wait_time获取内容  
func (obj *_AllVirtualSessionWaitHistoryMgr) GetFromWaitTime(waitTime float64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`wait_time` = ?", waitTime).Find(&results).Error
	
	return
}

// GetBatchFromWaitTime 批量查找 
func (obj *_AllVirtualSessionWaitHistoryMgr) GetBatchFromWaitTime(waitTimes []float64) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`wait_time` IN (?)", waitTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllVirtualSessionWaitHistoryMgr) FetchByPrimaryKey(sessionID int64 ,svrIP string ,svrPort int64 ,seq# int64 ) (result AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`session_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `seq#` = ?", sessionID , svrIP , svrPort , seq#).First(&result).Error
	
	return
}
 

 
 // FetchIndexByAllVirtualSessionWaitHistoryI1  获取多个内容
 func (obj *_AllVirtualSessionWaitHistoryMgr) FetchIndexByAllVirtualSessionWaitHistoryI1(sessionID int64 ) (results []*AllVirtualSessionWaitHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWaitHistory{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}
 

	

