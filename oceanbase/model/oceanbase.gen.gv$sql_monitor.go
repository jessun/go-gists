package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _Gv$sqlMonitorMgr struct {
	*_BaseMgr
}

// Gv$sqlMonitorMgr open func
func Gv$sqlMonitorMgr(db *gorm.DB) *_Gv$sqlMonitorMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$sqlMonitorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$sqlMonitorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$sql_monitor"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$sqlMonitorMgr) GetTableName() string {
	return "gv$sql_monitor"
}

// Reset 重置gorm会话
func (obj *_Gv$sqlMonitorMgr) Reset() *_Gv$sqlMonitorMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$sqlMonitorMgr) Get() (result Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$sqlMonitorMgr) Gets() (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$sqlMonitorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_Gv$sqlMonitorMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSQLExecID SQL_EXEC_ID获取 
func (obj *_Gv$sqlMonitorMgr) WithSQLExecID(sqlExecID int64) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_ID"] = sqlExecID })
}

// WithJobID JOB_ID获取 
func (obj *_Gv$sqlMonitorMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["JOB_ID"] = jobID })
}

// WithTaskID TASK_ID获取 
func (obj *_Gv$sqlMonitorMgr) WithTaskID(taskID int64) Option {
	return optionFunc(func(o *options) { o.query["TASK_ID"] = taskID })
}

// WithSvrIP SVR_IP获取 
func (obj *_Gv$sqlMonitorMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$sqlMonitorMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithSQLExecStart SQL_EXEC_START获取 
func (obj *_Gv$sqlMonitorMgr) WithSQLExecStart(sqlExecStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_START"] = sqlExecStart })
}

// WithPlanID PLAN_ID获取 
func (obj *_Gv$sqlMonitorMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_ID"] = planID })
}

// WithSchedulerIP SCHEDULER_IP获取 
func (obj *_Gv$sqlMonitorMgr) WithSchedulerIP(schedulerIP string) Option {
	return optionFunc(func(o *options) { o.query["SCHEDULER_IP"] = schedulerIP })
}

// WithSchedulerPort SCHEDULER_PORT获取 
func (obj *_Gv$sqlMonitorMgr) WithSchedulerPort(schedulerPort int64) Option {
	return optionFunc(func(o *options) { o.query["SCHEDULER_PORT"] = schedulerPort })
}

// WithMonitorInfo MONITOR_INFO获取 
func (obj *_Gv$sqlMonitorMgr) WithMonitorInfo(monitorInfo string) Option {
	return optionFunc(func(o *options) { o.query["MONITOR_INFO"] = monitorInfo })
}

// WithExtendInfo EXTEND_INFO获取 
func (obj *_Gv$sqlMonitorMgr) WithExtendInfo(extendInfo string) Option {
	return optionFunc(func(o *options) { o.query["EXTEND_INFO"] = extendInfo })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$sqlMonitorMgr) GetByOption(opts ...Option) (result Gv$sqlMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$sqlMonitorMgr) GetByOptions(opts ...Option) (results []*Gv$sqlMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$sqlMonitorMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$sqlMonitor,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where(options.query)
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
func (obj *_Gv$sqlMonitorMgr) GetFromConID(conID int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromConID(conIDs []int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLExecID 通过SQL_EXEC_ID获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromSQLExecID(sqlExecID int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SQL_EXEC_ID` = ?", sqlExecID).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecID 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromSQLExecID(sqlExecIDs []int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SQL_EXEC_ID` IN (?)", sqlExecIDs).Find(&results).Error
	
	return
}
 
// GetFromJobID 通过JOB_ID获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromJobID(jobID int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`JOB_ID` = ?", jobID).Find(&results).Error
	
	return
}

// GetBatchFromJobID 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromJobID(jobIDs []int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`JOB_ID` IN (?)", jobIDs).Find(&results).Error
	
	return
}
 
// GetFromTaskID 通过TASK_ID获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromTaskID(taskID int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`TASK_ID` = ?", taskID).Find(&results).Error
	
	return
}

// GetBatchFromTaskID 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromTaskID(taskIDs []int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`TASK_ID` IN (?)", taskIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromSvrIP(svrIP string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromSvrPort(svrPort int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSQLExecStart 通过SQL_EXEC_START获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromSQLExecStart(sqlExecStart time.Time) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SQL_EXEC_START` = ?", sqlExecStart).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecStart 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromSQLExecStart(sqlExecStarts []time.Time) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SQL_EXEC_START` IN (?)", sqlExecStarts).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过PLAN_ID获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromPlanID(planID int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`PLAN_ID` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromPlanID(planIDs []int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`PLAN_ID` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromSchedulerIP 通过SCHEDULER_IP获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromSchedulerIP(schedulerIP string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SCHEDULER_IP` = ?", schedulerIP).Find(&results).Error
	
	return
}

// GetBatchFromSchedulerIP 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromSchedulerIP(schedulerIPs []string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SCHEDULER_IP` IN (?)", schedulerIPs).Find(&results).Error
	
	return
}
 
// GetFromSchedulerPort 通过SCHEDULER_PORT获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromSchedulerPort(schedulerPort int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SCHEDULER_PORT` = ?", schedulerPort).Find(&results).Error
	
	return
}

// GetBatchFromSchedulerPort 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromSchedulerPort(schedulerPorts []int64) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`SCHEDULER_PORT` IN (?)", schedulerPorts).Find(&results).Error
	
	return
}
 
// GetFromMonitorInfo 通过MONITOR_INFO获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromMonitorInfo(monitorInfo string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`MONITOR_INFO` = ?", monitorInfo).Find(&results).Error
	
	return
}

// GetBatchFromMonitorInfo 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromMonitorInfo(monitorInfos []string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`MONITOR_INFO` IN (?)", monitorInfos).Find(&results).Error
	
	return
}
 
// GetFromExtendInfo 通过EXTEND_INFO获取内容  
func (obj *_Gv$sqlMonitorMgr) GetFromExtendInfo(extendInfo string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`EXTEND_INFO` = ?", extendInfo).Find(&results).Error
	
	return
}

// GetBatchFromExtendInfo 批量查找 
func (obj *_Gv$sqlMonitorMgr) GetBatchFromExtendInfo(extendInfos []string) (results []*Gv$sqlMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlMonitor{}).Where("`EXTEND_INFO` IN (?)", extendInfos).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

