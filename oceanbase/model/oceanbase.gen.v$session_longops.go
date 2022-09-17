package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sessionLongopsMgr struct {
	*_BaseMgr
}

// V$sessionLongopsMgr open func
func V$sessionLongopsMgr(db *gorm.DB) *_V$sessionLongopsMgr {
	if db == nil {
		panic(fmt.Errorf("V$sessionLongopsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sessionLongopsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$session_longops"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sessionLongopsMgr) GetTableName() string {
	return "v$session_longops"
}

// Reset 重置gorm会话
func (obj *_V$sessionLongopsMgr) Reset() *_V$sessionLongopsMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sessionLongopsMgr) Get() (result V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sessionLongopsMgr) Gets() (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sessionLongopsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSid SID获取 
func (obj *_V$sessionLongopsMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithOpname OPNAME获取 
func (obj *_V$sessionLongopsMgr) WithOpname(opname string) Option {
	return optionFunc(func(o *options) { o.query["OPNAME"] = opname })
}

// WithTarget TARGET获取 
func (obj *_V$sessionLongopsMgr) WithTarget(target string) Option {
	return optionFunc(func(o *options) { o.query["TARGET"] = target })
}

// WithSvrIP SVR_IP获取 
func (obj *_V$sessionLongopsMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$sessionLongopsMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithStartTime START_TIME获取 
func (obj *_V$sessionLongopsMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithElapsedSeconds ELAPSED_SECONDS获取 
func (obj *_V$sessionLongopsMgr) WithElapsedSeconds(elapsedSeconds float64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_SECONDS"] = elapsedSeconds })
}

// WithTimeRemaining TIME_REMAINING获取 
func (obj *_V$sessionLongopsMgr) WithTimeRemaining(timeRemaining int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_REMAINING"] = timeRemaining })
}

// WithLastUpdateTime LAST_UPDATE_TIME获取 
func (obj *_V$sessionLongopsMgr) WithLastUpdateTime(lastUpdateTime int64) Option {
	return optionFunc(func(o *options) { o.query["LAST_UPDATE_TIME"] = lastUpdateTime })
}

// WithMessage MESSAGE获取 
func (obj *_V$sessionLongopsMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["MESSAGE"] = message })
}


// GetByOption 功能选项模式获取
func (obj *_V$sessionLongopsMgr) GetByOption(opts ...Option) (result V$sessionLongops, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sessionLongopsMgr) GetByOptions(opts ...Option) (results []*V$sessionLongops, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sessionLongopsMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sessionLongops,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where(options.query)
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
func (obj *_V$sessionLongopsMgr) GetFromSid(sid int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromSid(sids []int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromOpname 通过OPNAME获取内容  
func (obj *_V$sessionLongopsMgr) GetFromOpname(opname string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`OPNAME` = ?", opname).Find(&results).Error
	
	return
}

// GetBatchFromOpname 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromOpname(opnames []string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`OPNAME` IN (?)", opnames).Find(&results).Error
	
	return
}
 
// GetFromTarget 通过TARGET获取内容  
func (obj *_V$sessionLongopsMgr) GetFromTarget(target string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`TARGET` = ?", target).Find(&results).Error
	
	return
}

// GetBatchFromTarget 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromTarget(targets []string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`TARGET` IN (?)", targets).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_V$sessionLongopsMgr) GetFromSvrIP(svrIP string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$sessionLongopsMgr) GetFromSvrPort(svrPort int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromStartTime 通过START_TIME获取内容  
func (obj *_V$sessionLongopsMgr) GetFromStartTime(startTime int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`START_TIME` = ?", startTime).Find(&results).Error
	
	return
}

// GetBatchFromStartTime 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromStartTime(startTimes []int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error
	
	return
}
 
// GetFromElapsedSeconds 通过ELAPSED_SECONDS获取内容  
func (obj *_V$sessionLongopsMgr) GetFromElapsedSeconds(elapsedSeconds float64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`ELAPSED_SECONDS` = ?", elapsedSeconds).Find(&results).Error
	
	return
}

// GetBatchFromElapsedSeconds 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromElapsedSeconds(elapsedSecondss []float64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`ELAPSED_SECONDS` IN (?)", elapsedSecondss).Find(&results).Error
	
	return
}
 
// GetFromTimeRemaining 通过TIME_REMAINING获取内容  
func (obj *_V$sessionLongopsMgr) GetFromTimeRemaining(timeRemaining int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`TIME_REMAINING` = ?", timeRemaining).Find(&results).Error
	
	return
}

// GetBatchFromTimeRemaining 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromTimeRemaining(timeRemainings []int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`TIME_REMAINING` IN (?)", timeRemainings).Find(&results).Error
	
	return
}
 
// GetFromLastUpdateTime 通过LAST_UPDATE_TIME获取内容  
func (obj *_V$sessionLongopsMgr) GetFromLastUpdateTime(lastUpdateTime int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`LAST_UPDATE_TIME` = ?", lastUpdateTime).Find(&results).Error
	
	return
}

// GetBatchFromLastUpdateTime 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromLastUpdateTime(lastUpdateTimes []int64) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`LAST_UPDATE_TIME` IN (?)", lastUpdateTimes).Find(&results).Error
	
	return
}
 
// GetFromMessage 通过MESSAGE获取内容  
func (obj *_V$sessionLongopsMgr) GetFromMessage(message string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`MESSAGE` = ?", message).Find(&results).Error
	
	return
}

// GetBatchFromMessage 批量查找 
func (obj *_V$sessionLongopsMgr) GetBatchFromMessage(messages []string) (results []*V$sessionLongops, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sessionLongops{}).Where("`MESSAGE` IN (?)", messages).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

