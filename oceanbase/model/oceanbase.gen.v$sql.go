package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _V$sqlMgr struct {
	*_BaseMgr
}

// V$sqlMgr open func
func V$sqlMgr(db *gorm.DB) *_V$sqlMgr {
	if db == nil {
		panic(fmt.Errorf("V$sqlMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sqlMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sql"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sqlMgr) GetTableName() string {
	return "v$sql"
}

// Reset 重置gorm会话
func (obj *_V$sqlMgr) Reset() *_V$sqlMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sqlMgr) Get() (result V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sqlMgr) Gets() (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sqlMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sql{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$sqlMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSvrIP SVR_IP获取 
func (obj *_V$sqlMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$sqlMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithPlanID PLAN_ID获取 
func (obj *_V$sqlMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_ID"] = planID })
}

// WithSQLID SQL_ID获取 
func (obj *_V$sqlMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["SQL_ID"] = sqlID })
}

// WithType TYPE获取 
func (obj *_V$sqlMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithSQLText SQL_TEXT获取 
func (obj *_V$sqlMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["SQL_TEXT"] = sqlText })
}

// WithPlanHashValue PLAN_HASH_VALUE获取 
func (obj *_V$sqlMgr) WithPlanHashValue(planHashValue uint64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_HASH_VALUE"] = planHashValue })
}

// WithFirstLoadTime FIRST_LOAD_TIME获取 
func (obj *_V$sqlMgr) WithFirstLoadTime(firstLoadTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["FIRST_LOAD_TIME"] = firstLoadTime })
}

// WithLastActiveTime LAST_ACTIVE_TIME获取 
func (obj *_V$sqlMgr) WithLastActiveTime(lastActiveTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_ACTIVE_TIME"] = lastActiveTime })
}

// WithAvgExeUsec AVG_EXE_USEC获取 
func (obj *_V$sqlMgr) WithAvgExeUsec(avgExeUsec int64) Option {
	return optionFunc(func(o *options) { o.query["AVG_EXE_USEC"] = avgExeUsec })
}

// WithSlowestExeTime SLOWEST_EXE_TIME获取 
func (obj *_V$sqlMgr) WithSlowestExeTime(slowestExeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["SLOWEST_EXE_TIME"] = slowestExeTime })
}

// WithSlowestExeUsec SLOWEST_EXE_USEC获取 
func (obj *_V$sqlMgr) WithSlowestExeUsec(slowestExeUsec int64) Option {
	return optionFunc(func(o *options) { o.query["SLOWEST_EXE_USEC"] = slowestExeUsec })
}

// WithSlowCount SLOW_COUNT获取 
func (obj *_V$sqlMgr) WithSlowCount(slowCount int64) Option {
	return optionFunc(func(o *options) { o.query["SLOW_COUNT"] = slowCount })
}

// WithHitCount HIT_COUNT获取 
func (obj *_V$sqlMgr) WithHitCount(hitCount int64) Option {
	return optionFunc(func(o *options) { o.query["HIT_COUNT"] = hitCount })
}

// WithPlanSize PLAN_SIZE获取 
func (obj *_V$sqlMgr) WithPlanSize(planSize int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_SIZE"] = planSize })
}

// WithExecutions EXECUTIONS获取 
func (obj *_V$sqlMgr) WithExecutions(executions int64) Option {
	return optionFunc(func(o *options) { o.query["EXECUTIONS"] = executions })
}

// WithDiskReads DISK_READS获取 
func (obj *_V$sqlMgr) WithDiskReads(diskReads int64) Option {
	return optionFunc(func(o *options) { o.query["DISK_READS"] = diskReads })
}

// WithDirectWrites DIRECT_WRITES获取 
func (obj *_V$sqlMgr) WithDirectWrites(directWrites int64) Option {
	return optionFunc(func(o *options) { o.query["DIRECT_WRITES"] = directWrites })
}

// WithBufferGets BUFFER_GETS获取 
func (obj *_V$sqlMgr) WithBufferGets(bufferGets int64) Option {
	return optionFunc(func(o *options) { o.query["BUFFER_GETS"] = bufferGets })
}

// WithApplicationWaitTime APPLICATION_WAIT_TIME获取 
func (obj *_V$sqlMgr) WithApplicationWaitTime(applicationWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["APPLICATION_WAIT_TIME"] = applicationWaitTime })
}

// WithConcurrencyWaitTime CONCURRENCY_WAIT_TIME获取 
func (obj *_V$sqlMgr) WithConcurrencyWaitTime(concurrencyWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["CONCURRENCY_WAIT_TIME"] = concurrencyWaitTime })
}

// WithUserIoWaitTime USER_IO_WAIT_TIME获取 
func (obj *_V$sqlMgr) WithUserIoWaitTime(userIoWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["USER_IO_WAIT_TIME"] = userIoWaitTime })
}

// WithRowsProcessed ROWS_PROCESSED获取 
func (obj *_V$sqlMgr) WithRowsProcessed(rowsProcessed int64) Option {
	return optionFunc(func(o *options) { o.query["ROWS_PROCESSED"] = rowsProcessed })
}

// WithElapsedTime ELAPSED_TIME获取 
func (obj *_V$sqlMgr) WithElapsedTime(elapsedTime uint64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_TIME"] = elapsedTime })
}

// WithCPUTime CPU_TIME获取 
func (obj *_V$sqlMgr) WithCPUTime(cpuTime uint64) Option {
	return optionFunc(func(o *options) { o.query["CPU_TIME"] = cpuTime })
}


// GetByOption 功能选项模式获取
func (obj *_V$sqlMgr) GetByOption(opts ...Option) (result V$sql, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sqlMgr) GetByOptions(opts ...Option) (results []*V$sql, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sqlMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sql,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where(options.query)
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
func (obj *_V$sqlMgr) GetFromConID(conID int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sqlMgr) GetBatchFromConID(conIDs []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_V$sqlMgr) GetFromSvrIP(svrIP string) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$sqlMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$sqlMgr) GetFromSvrPort(svrPort int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$sqlMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过PLAN_ID获取内容  
func (obj *_V$sqlMgr) GetFromPlanID(planID int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`PLAN_ID` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_V$sqlMgr) GetBatchFromPlanID(planIDs []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`PLAN_ID` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLID 通过SQL_ID获取内容  
func (obj *_V$sqlMgr) GetFromSQLID(sqlID string) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SQL_ID` = ?", sqlID).Find(&results).Error
	
	return
}

// GetBatchFromSQLID 批量查找 
func (obj *_V$sqlMgr) GetBatchFromSQLID(sqlIDs []string) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SQL_ID` IN (?)", sqlIDs).Find(&results).Error
	
	return
}
 
// GetFromType 通过TYPE获取内容  
func (obj *_V$sqlMgr) GetFromType(_type int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`TYPE` = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量查找 
func (obj *_V$sqlMgr) GetBatchFromType(_types []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`TYPE` IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromSQLText 通过SQL_TEXT获取内容  
func (obj *_V$sqlMgr) GetFromSQLText(sqlText string) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SQL_TEXT` = ?", sqlText).Find(&results).Error
	
	return
}

// GetBatchFromSQLText 批量查找 
func (obj *_V$sqlMgr) GetBatchFromSQLText(sqlTexts []string) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SQL_TEXT` IN (?)", sqlTexts).Find(&results).Error
	
	return
}
 
// GetFromPlanHashValue 通过PLAN_HASH_VALUE获取内容  
func (obj *_V$sqlMgr) GetFromPlanHashValue(planHashValue uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`PLAN_HASH_VALUE` = ?", planHashValue).Find(&results).Error
	
	return
}

// GetBatchFromPlanHashValue 批量查找 
func (obj *_V$sqlMgr) GetBatchFromPlanHashValue(planHashValues []uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`PLAN_HASH_VALUE` IN (?)", planHashValues).Find(&results).Error
	
	return
}
 
// GetFromFirstLoadTime 通过FIRST_LOAD_TIME获取内容  
func (obj *_V$sqlMgr) GetFromFirstLoadTime(firstLoadTime time.Time) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`FIRST_LOAD_TIME` = ?", firstLoadTime).Find(&results).Error
	
	return
}

// GetBatchFromFirstLoadTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromFirstLoadTime(firstLoadTimes []time.Time) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`FIRST_LOAD_TIME` IN (?)", firstLoadTimes).Find(&results).Error
	
	return
}
 
// GetFromLastActiveTime 通过LAST_ACTIVE_TIME获取内容  
func (obj *_V$sqlMgr) GetFromLastActiveTime(lastActiveTime time.Time) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`LAST_ACTIVE_TIME` = ?", lastActiveTime).Find(&results).Error
	
	return
}

// GetBatchFromLastActiveTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromLastActiveTime(lastActiveTimes []time.Time) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`LAST_ACTIVE_TIME` IN (?)", lastActiveTimes).Find(&results).Error
	
	return
}
 
// GetFromAvgExeUsec 通过AVG_EXE_USEC获取内容  
func (obj *_V$sqlMgr) GetFromAvgExeUsec(avgExeUsec int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`AVG_EXE_USEC` = ?", avgExeUsec).Find(&results).Error
	
	return
}

// GetBatchFromAvgExeUsec 批量查找 
func (obj *_V$sqlMgr) GetBatchFromAvgExeUsec(avgExeUsecs []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`AVG_EXE_USEC` IN (?)", avgExeUsecs).Find(&results).Error
	
	return
}
 
// GetFromSlowestExeTime 通过SLOWEST_EXE_TIME获取内容  
func (obj *_V$sqlMgr) GetFromSlowestExeTime(slowestExeTime time.Time) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SLOWEST_EXE_TIME` = ?", slowestExeTime).Find(&results).Error
	
	return
}

// GetBatchFromSlowestExeTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromSlowestExeTime(slowestExeTimes []time.Time) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SLOWEST_EXE_TIME` IN (?)", slowestExeTimes).Find(&results).Error
	
	return
}
 
// GetFromSlowestExeUsec 通过SLOWEST_EXE_USEC获取内容  
func (obj *_V$sqlMgr) GetFromSlowestExeUsec(slowestExeUsec int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SLOWEST_EXE_USEC` = ?", slowestExeUsec).Find(&results).Error
	
	return
}

// GetBatchFromSlowestExeUsec 批量查找 
func (obj *_V$sqlMgr) GetBatchFromSlowestExeUsec(slowestExeUsecs []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SLOWEST_EXE_USEC` IN (?)", slowestExeUsecs).Find(&results).Error
	
	return
}
 
// GetFromSlowCount 通过SLOW_COUNT获取内容  
func (obj *_V$sqlMgr) GetFromSlowCount(slowCount int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SLOW_COUNT` = ?", slowCount).Find(&results).Error
	
	return
}

// GetBatchFromSlowCount 批量查找 
func (obj *_V$sqlMgr) GetBatchFromSlowCount(slowCounts []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`SLOW_COUNT` IN (?)", slowCounts).Find(&results).Error
	
	return
}
 
// GetFromHitCount 通过HIT_COUNT获取内容  
func (obj *_V$sqlMgr) GetFromHitCount(hitCount int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`HIT_COUNT` = ?", hitCount).Find(&results).Error
	
	return
}

// GetBatchFromHitCount 批量查找 
func (obj *_V$sqlMgr) GetBatchFromHitCount(hitCounts []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`HIT_COUNT` IN (?)", hitCounts).Find(&results).Error
	
	return
}
 
// GetFromPlanSize 通过PLAN_SIZE获取内容  
func (obj *_V$sqlMgr) GetFromPlanSize(planSize int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`PLAN_SIZE` = ?", planSize).Find(&results).Error
	
	return
}

// GetBatchFromPlanSize 批量查找 
func (obj *_V$sqlMgr) GetBatchFromPlanSize(planSizes []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`PLAN_SIZE` IN (?)", planSizes).Find(&results).Error
	
	return
}
 
// GetFromExecutions 通过EXECUTIONS获取内容  
func (obj *_V$sqlMgr) GetFromExecutions(executions int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`EXECUTIONS` = ?", executions).Find(&results).Error
	
	return
}

// GetBatchFromExecutions 批量查找 
func (obj *_V$sqlMgr) GetBatchFromExecutions(executionss []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`EXECUTIONS` IN (?)", executionss).Find(&results).Error
	
	return
}
 
// GetFromDiskReads 通过DISK_READS获取内容  
func (obj *_V$sqlMgr) GetFromDiskReads(diskReads int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`DISK_READS` = ?", diskReads).Find(&results).Error
	
	return
}

// GetBatchFromDiskReads 批量查找 
func (obj *_V$sqlMgr) GetBatchFromDiskReads(diskReadss []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`DISK_READS` IN (?)", diskReadss).Find(&results).Error
	
	return
}
 
// GetFromDirectWrites 通过DIRECT_WRITES获取内容  
func (obj *_V$sqlMgr) GetFromDirectWrites(directWrites int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`DIRECT_WRITES` = ?", directWrites).Find(&results).Error
	
	return
}

// GetBatchFromDirectWrites 批量查找 
func (obj *_V$sqlMgr) GetBatchFromDirectWrites(directWritess []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`DIRECT_WRITES` IN (?)", directWritess).Find(&results).Error
	
	return
}
 
// GetFromBufferGets 通过BUFFER_GETS获取内容  
func (obj *_V$sqlMgr) GetFromBufferGets(bufferGets int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`BUFFER_GETS` = ?", bufferGets).Find(&results).Error
	
	return
}

// GetBatchFromBufferGets 批量查找 
func (obj *_V$sqlMgr) GetBatchFromBufferGets(bufferGetss []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`BUFFER_GETS` IN (?)", bufferGetss).Find(&results).Error
	
	return
}
 
// GetFromApplicationWaitTime 通过APPLICATION_WAIT_TIME获取内容  
func (obj *_V$sqlMgr) GetFromApplicationWaitTime(applicationWaitTime uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`APPLICATION_WAIT_TIME` = ?", applicationWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromApplicationWaitTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromApplicationWaitTime(applicationWaitTimes []uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`APPLICATION_WAIT_TIME` IN (?)", applicationWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromConcurrencyWaitTime 通过CONCURRENCY_WAIT_TIME获取内容  
func (obj *_V$sqlMgr) GetFromConcurrencyWaitTime(concurrencyWaitTime uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`CONCURRENCY_WAIT_TIME` = ?", concurrencyWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromConcurrencyWaitTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromConcurrencyWaitTime(concurrencyWaitTimes []uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`CONCURRENCY_WAIT_TIME` IN (?)", concurrencyWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromUserIoWaitTime 通过USER_IO_WAIT_TIME获取内容  
func (obj *_V$sqlMgr) GetFromUserIoWaitTime(userIoWaitTime uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`USER_IO_WAIT_TIME` = ?", userIoWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromUserIoWaitTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromUserIoWaitTime(userIoWaitTimes []uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`USER_IO_WAIT_TIME` IN (?)", userIoWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromRowsProcessed 通过ROWS_PROCESSED获取内容  
func (obj *_V$sqlMgr) GetFromRowsProcessed(rowsProcessed int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`ROWS_PROCESSED` = ?", rowsProcessed).Find(&results).Error
	
	return
}

// GetBatchFromRowsProcessed 批量查找 
func (obj *_V$sqlMgr) GetBatchFromRowsProcessed(rowsProcesseds []int64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`ROWS_PROCESSED` IN (?)", rowsProcesseds).Find(&results).Error
	
	return
}
 
// GetFromElapsedTime 通过ELAPSED_TIME获取内容  
func (obj *_V$sqlMgr) GetFromElapsedTime(elapsedTime uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`ELAPSED_TIME` = ?", elapsedTime).Find(&results).Error
	
	return
}

// GetBatchFromElapsedTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromElapsedTime(elapsedTimes []uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`ELAPSED_TIME` IN (?)", elapsedTimes).Find(&results).Error
	
	return
}
 
// GetFromCPUTime 通过CPU_TIME获取内容  
func (obj *_V$sqlMgr) GetFromCPUTime(cpuTime uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`CPU_TIME` = ?", cpuTime).Find(&results).Error
	
	return
}

// GetBatchFromCPUTime 批量查找 
func (obj *_V$sqlMgr) GetBatchFromCPUTime(cpuTimes []uint64) (results []*V$sql, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sql{}).Where("`CPU_TIME` IN (?)", cpuTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

