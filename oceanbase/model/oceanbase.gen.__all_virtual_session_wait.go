package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualSessionWaitMgr struct {
	*_BaseMgr
}

// AllVirtualSessionWaitMgr open func
func AllVirtualSessionWaitMgr(db *gorm.DB) *_AllVirtualSessionWaitMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSessionWaitMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSessionWaitMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_session_wait"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSessionWaitMgr) GetTableName() string {
	return "__all_virtual_session_wait"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSessionWaitMgr) Reset() *_AllVirtualSessionWaitMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualSessionWaitMgr) Get() (result AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSessionWaitMgr) Gets() (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSessionWaitMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSessionID session_id获取 
func (obj *_AllVirtualSessionWaitMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithSvrIP svr_ip获取 
func (obj *_AllVirtualSessionWaitMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllVirtualSessionWaitMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取 
func (obj *_AllVirtualSessionWaitMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithEvent event获取 
func (obj *_AllVirtualSessionWaitMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithP1text p1text获取 
func (obj *_AllVirtualSessionWaitMgr) WithP1text(p1text string) Option {
	return optionFunc(func(o *options) { o.query["p1text"] = p1text })
}

// WithP1 p1获取 
func (obj *_AllVirtualSessionWaitMgr) WithP1(p1 uint64) Option {
	return optionFunc(func(o *options) { o.query["p1"] = p1 })
}

// WithP2text p2text获取 
func (obj *_AllVirtualSessionWaitMgr) WithP2text(p2text string) Option {
	return optionFunc(func(o *options) { o.query["p2text"] = p2text })
}

// WithP2 p2获取 
func (obj *_AllVirtualSessionWaitMgr) WithP2(p2 uint64) Option {
	return optionFunc(func(o *options) { o.query["p2"] = p2 })
}

// WithP3text p3text获取 
func (obj *_AllVirtualSessionWaitMgr) WithP3text(p3text string) Option {
	return optionFunc(func(o *options) { o.query["p3text"] = p3text })
}

// WithP3 p3获取 
func (obj *_AllVirtualSessionWaitMgr) WithP3(p3 uint64) Option {
	return optionFunc(func(o *options) { o.query["p3"] = p3 })
}

// WithLevel level获取 
func (obj *_AllVirtualSessionWaitMgr) WithLevel(level int64) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithWaitClassID wait_class_id获取 
func (obj *_AllVirtualSessionWaitMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class_id"] = waitClassID })
}

// WithWaitClass# wait_class#获取 
func (obj *_AllVirtualSessionWaitMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class#"] = waitClass# })
}

// WithWaitClass wait_class获取 
func (obj *_AllVirtualSessionWaitMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["wait_class"] = waitClass })
}

// WithState state获取 
func (obj *_AllVirtualSessionWaitMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithWaitTimeMicro wait_time_micro获取 
func (obj *_AllVirtualSessionWaitMgr) WithWaitTimeMicro(waitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["wait_time_micro"] = waitTimeMicro })
}

// WithTimeRemainingMicro time_remaining_micro获取 
func (obj *_AllVirtualSessionWaitMgr) WithTimeRemainingMicro(timeRemainingMicro int64) Option {
	return optionFunc(func(o *options) { o.query["time_remaining_micro"] = timeRemainingMicro })
}

// WithTimeSinceLastWaitMicro time_since_last_wait_micro获取 
func (obj *_AllVirtualSessionWaitMgr) WithTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) Option {
	return optionFunc(func(o *options) { o.query["time_since_last_wait_micro"] = timeSinceLastWaitMicro })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualSessionWaitMgr) GetByOption(opts ...Option) (result AllVirtualSessionWait, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSessionWaitMgr) GetByOptions(opts ...Option) (results []*AllVirtualSessionWait, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualSessionWaitMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSessionWait,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where(options.query)
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
func (obj *_AllVirtualSessionWaitMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过event获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromEvent(event string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`event` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromEvent(events []string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`event` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromP1text 通过p1text获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromP1text(p1text string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p1text` = ?", p1text).Find(&results).Error
	
	return
}

// GetBatchFromP1text 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromP1text(p1texts []string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p1text` IN (?)", p1texts).Find(&results).Error
	
	return
}
 
// GetFromP1 通过p1获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromP1(p1 uint64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p1` = ?", p1).Find(&results).Error
	
	return
}

// GetBatchFromP1 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromP1(p1s []uint64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p1` IN (?)", p1s).Find(&results).Error
	
	return
}
 
// GetFromP2text 通过p2text获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromP2text(p2text string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p2text` = ?", p2text).Find(&results).Error
	
	return
}

// GetBatchFromP2text 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromP2text(p2texts []string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p2text` IN (?)", p2texts).Find(&results).Error
	
	return
}
 
// GetFromP2 通过p2获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromP2(p2 uint64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p2` = ?", p2).Find(&results).Error
	
	return
}

// GetBatchFromP2 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromP2(p2s []uint64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p2` IN (?)", p2s).Find(&results).Error
	
	return
}
 
// GetFromP3text 通过p3text获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromP3text(p3text string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p3text` = ?", p3text).Find(&results).Error
	
	return
}

// GetBatchFromP3text 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromP3text(p3texts []string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p3text` IN (?)", p3texts).Find(&results).Error
	
	return
}
 
// GetFromP3 通过p3获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromP3(p3 uint64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p3` = ?", p3).Find(&results).Error
	
	return
}

// GetBatchFromP3 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromP3(p3s []uint64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`p3` IN (?)", p3s).Find(&results).Error
	
	return
}
 
// GetFromLevel 通过level获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromLevel(level int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`level` = ?", level).Find(&results).Error
	
	return
}

// GetBatchFromLevel 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromLevel(levels []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`level` IN (?)", levels).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过wait_class_id获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromWaitClassID(waitClassID int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_class_id` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_class_id` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过wait_class#获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromWaitClass#(waitClass# int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_class#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_class#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过wait_class获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromWaitClass(waitClass string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_class` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromWaitClass(waitClasss []string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_class` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromState 通过state获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromState(state string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`state` = ?", state).Find(&results).Error
	
	return
}

// GetBatchFromState 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromState(states []string) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`state` IN (?)", states).Find(&results).Error
	
	return
}
 
// GetFromWaitTimeMicro 通过wait_time_micro获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromWaitTimeMicro(waitTimeMicro int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_time_micro` = ?", waitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromWaitTimeMicro 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromWaitTimeMicro(waitTimeMicros []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`wait_time_micro` IN (?)", waitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTimeRemainingMicro 通过time_remaining_micro获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromTimeRemainingMicro(timeRemainingMicro int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`time_remaining_micro` = ?", timeRemainingMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeRemainingMicro 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromTimeRemainingMicro(timeRemainingMicros []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`time_remaining_micro` IN (?)", timeRemainingMicros).Find(&results).Error
	
	return
}
 
// GetFromTimeSinceLastWaitMicro 通过time_since_last_wait_micro获取内容  
func (obj *_AllVirtualSessionWaitMgr) GetFromTimeSinceLastWaitMicro(timeSinceLastWaitMicro int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`time_since_last_wait_micro` = ?", timeSinceLastWaitMicro).Find(&results).Error
	
	return
}

// GetBatchFromTimeSinceLastWaitMicro 批量查找 
func (obj *_AllVirtualSessionWaitMgr) GetBatchFromTimeSinceLastWaitMicro(timeSinceLastWaitMicros []int64) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`time_since_last_wait_micro` IN (?)", timeSinceLastWaitMicros).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllVirtualSessionWaitMgr) FetchByPrimaryKey(sessionID int64 ,svrIP string ,svrPort int64 ) (result AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`session_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", sessionID , svrIP , svrPort).First(&result).Error
	
	return
}
 

 
 // FetchIndexByAllVirtualSessionWaitI1  获取多个内容
 func (obj *_AllVirtualSessionWaitMgr) FetchIndexByAllVirtualSessionWaitI1(sessionID int64 ) (results []*AllVirtualSessionWait, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSessionWait{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}
 

	

