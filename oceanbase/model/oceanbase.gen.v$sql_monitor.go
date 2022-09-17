package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _V$sqlMonitorMgr struct {
	*_BaseMgr
}

// V$sqlMonitorMgr open func
func V$sqlMonitorMgr(db *gorm.DB) *_V$sqlMonitorMgr {
	if db == nil {
		panic(fmt.Errorf("V$sqlMonitorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sqlMonitorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sql_monitor"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sqlMonitorMgr) GetTableName() string {
	return "v$sql_monitor"
}

// Reset 重置gorm会话
func (obj *_V$sqlMonitorMgr) Reset() *_V$sqlMonitorMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sqlMonitorMgr) Get() (result V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sqlMonitorMgr) Gets() (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sqlMonitorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$sqlMonitorMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSQLExecID SQL_EXEC_ID获取 
func (obj *_V$sqlMonitorMgr) WithSQLExecID(sqlExecID int64) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_ID"] = sqlExecID })
}

// WithJobID JOB_ID获取 
func (obj *_V$sqlMonitorMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTaskID TASK_ID获取 
func (obj *_V$sqlMonitorMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["TASK_ID"] = taskID })
}

// WithSvrIP SVR_IP获取 
func (obj *_V$sqlMonitorMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$sqlMonitorMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithSQLExecStart SQL_EXEC_START获取 
func (obj *_V$sqlMonitorMgr) WithSQLExecStart(sqlExecStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_START"] = sqlExecStart })
}

// WithPlanID PLAN_ID获取 
func (obj *_V$sqlMonitorMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_ID"] = planID })
}

// WithSchedulerIP SCHEDULER_IP获取 
func (obj *_V$sqlMonitorMgr) WithSchedulerIP(schedulerIP string) Option {
	return optionFunc(func(o *options) { o.query["SCHEDULER_IP"] = schedulerIP })
}

// WithSchedulerPort SCHEDULER_PORT获取 
func (obj *_V$sqlMonitorMgr) WithSchedulerPort(schedulerPort int64) Option {
	return optionFunc(func(o *options) { o.query["SCHEDULER_PORT"] = schedulerPort })
}

// WithMonitorInfo MONITOR_INFO获取 
func (obj *_V$sqlMonitorMgr) WithMonitorInfo(monitorInfo string) Option {
	return optionFunc(func(o *options) { o.query["MONITOR_INFO"] = monitorInfo })
}

// WithExtendInfo EXTEND_INFO获取 
func (obj *_V$sqlMonitorMgr) WithExtendInfo(extendInfo string) Option {
	return optionFunc(func(o *options) { o.query["EXTEND_INFO"] = extendInfo })
}


// GetByOption 功能选项模式获取
func (obj *_V$sqlMonitorMgr) GetByOption(opts ...Option) (result V$sqlMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sqlMonitorMgr) GetByOptions(opts ...Option) (results []*V$sqlMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sqlMonitorMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sqlMonitor,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where(options.query)
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
func (obj *_V$sqlMonitorMgr) GetFromConID(conID int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromConID(conIDs []int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLExecID 通过SQL_EXEC_ID获取内容  
func (obj *_V$sqlMonitorMgr) GetFromSQLExecID(sqlExecID int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SQL_EXEC_ID` = ?", sqlExecID).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecID 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromSQLExecID(sqlExecIDs []int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SQL_EXEC_ID` IN (?)", sqlExecIDs).Find(&results).Error
	
	return
}
 
// GetFromJobID 通过JOB_ID获取内容  
func (obj *_V$sqlMonitorMgr) GetFromJobID(jobID int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error
	
	return
}

// GetBatchFromJobID 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromJobID(jobIDs []int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error
	
	return
}
 
// GetFromTaskID 通过TASK_ID获取内容  
func (obj *_V$sqlMonitorMgr) GetFromTaskID(taskID int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`TASK_ID` = ?", taskID).Find(&results).Error
	
	return
}

// GetBatchFromTaskID 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromTaskID(taskIDs []int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`TASK_ID` IN (?)", taskIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_V$sqlMonitorMgr) GetFromSvrIP(svrIP string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$sqlMonitorMgr) GetFromSvrPort(svrPort int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSQLExecStart 通过SQL_EXEC_START获取内容  
func (obj *_V$sqlMonitorMgr) GetFromSQLExecStart(sqlExecStart time.Time) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SQL_EXEC_START` = ?", sqlExecStart).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecStart 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromSQLExecStart(sqlExecStarts []time.Time) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SQL_EXEC_START` IN (?)", sqlExecStarts).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过PLAN_ID获取内容  
func (obj *_V$sqlMonitorMgr) GetFromPlanID(planID int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`PLAN_ID` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromPlanID(planIDs []int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`PLAN_ID` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromSchedulerIP 通过SCHEDULER_IP获取内容  
func (obj *_V$sqlMonitorMgr) GetFromSchedulerIP(schedulerIP string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SCHEDULER_IP` = ?", schedulerIP).Find(&results).Error
	
	return
}

// GetBatchFromSchedulerIP 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromSchedulerIP(schedulerIPs []string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SCHEDULER_IP` IN (?)", schedulerIPs).Find(&results).Error
	
	return
}
 
// GetFromSchedulerPort 通过SCHEDULER_PORT获取内容  
func (obj *_V$sqlMonitorMgr) GetFromSchedulerPort(schedulerPort int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SCHEDULER_PORT` = ?", schedulerPort).Find(&results).Error
	
	return
}

// GetBatchFromSchedulerPort 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromSchedulerPort(schedulerPorts []int64) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`SCHEDULER_PORT` IN (?)", schedulerPorts).Find(&results).Error
	
	return
}
 
// GetFromMonitorInfo 通过MONITOR_INFO获取内容  
func (obj *_V$sqlMonitorMgr) GetFromMonitorInfo(monitorInfo string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`MONITOR_INFO` = ?", monitorInfo).Find(&results).Error
	
	return
}

// GetBatchFromMonitorInfo 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromMonitorInfo(monitorInfos []string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`MONITOR_INFO` IN (?)", monitorInfos).Find(&results).Error
	
	return
}
 
// GetFromExtendInfo 通过EXTEND_INFO获取内容  
func (obj *_V$sqlMonitorMgr) GetFromExtendInfo(extendInfo string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`EXTEND_INFO` = ?", extendInfo).Find(&results).Error
	
	return
}

// GetBatchFromExtendInfo 批量查找 
func (obj *_V$sqlMonitorMgr) GetBatchFromExtendInfo(extendInfos []string) (results []*V$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitor{}).Where("`EXTEND_INFO` IN (?)", extendInfos).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

