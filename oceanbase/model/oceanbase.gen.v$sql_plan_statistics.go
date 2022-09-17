package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sqlPlanStatisticsMgr struct {
	*_BaseMgr
}

// V$sqlPlanStatisticsMgr open func
func V$sqlPlanStatisticsMgr(db *gorm.DB) *_V$sqlPlanStatisticsMgr {
	if db == nil {
		panic(fmt.Errorf("V$sqlPlanStatisticsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sqlPlanStatisticsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sql_plan_statistics"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sqlPlanStatisticsMgr) GetTableName() string {
	return "v$sql_plan_statistics"
}

// Reset 重置gorm会话
func (obj *_V$sqlPlanStatisticsMgr) Reset() *_V$sqlPlanStatisticsMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sqlPlanStatisticsMgr) Get() (result V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sqlPlanStatisticsMgr) Gets() (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sqlPlanStatisticsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$sqlPlanStatisticsMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSvrIP SVR_IP获取 
func (obj *_V$sqlPlanStatisticsMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$sqlPlanStatisticsMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithPlanID PLAN_ID获取 
func (obj *_V$sqlPlanStatisticsMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_ID"] = planID })
}

// WithOperationID OPERATION_ID获取 
func (obj *_V$sqlPlanStatisticsMgr) WithOperationID(operationID int64) Option {
	return optionFunc(func(o *options) { o.query["OPERATION_ID"] = operationID })
}

// WithExecutions EXECUTIONS获取 
func (obj *_V$sqlPlanStatisticsMgr) WithExecutions(executions int64) Option {
	return optionFunc(func(o *options) { o.query["EXECUTIONS"] = executions })
}

// WithOutputRows OUTPUT_ROWS获取 
func (obj *_V$sqlPlanStatisticsMgr) WithOutputRows(outputRows int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_ROWS"] = outputRows })
}

// WithInputRows INPUT_ROWS获取 
func (obj *_V$sqlPlanStatisticsMgr) WithInputRows(inputRows int64) Option {
	return optionFunc(func(o *options) { o.query["INPUT_ROWS"] = inputRows })
}

// WithRescanTimes RESCAN_TIMES获取 
func (obj *_V$sqlPlanStatisticsMgr) WithRescanTimes(rescanTimes int64) Option {
	return optionFunc(func(o *options) { o.query["RESCAN_TIMES"] = rescanTimes })
}

// WithBufferGets BUFFER_GETS获取 
func (obj *_V$sqlPlanStatisticsMgr) WithBufferGets(bufferGets int64) Option {
	return optionFunc(func(o *options) { o.query["BUFFER_GETS"] = bufferGets })
}

// WithDiskReads DISK_READS获取 
func (obj *_V$sqlPlanStatisticsMgr) WithDiskReads(diskReads int64) Option {
	return optionFunc(func(o *options) { o.query["DISK_READS"] = diskReads })
}

// WithDiskWrites DISK_WRITES获取 
func (obj *_V$sqlPlanStatisticsMgr) WithDiskWrites(diskWrites int64) Option {
	return optionFunc(func(o *options) { o.query["DISK_WRITES"] = diskWrites })
}

// WithElapsedTime ELAPSED_TIME获取 
func (obj *_V$sqlPlanStatisticsMgr) WithElapsedTime(elapsedTime int64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_TIME"] = elapsedTime })
}

// WithExtendInfo1 EXTEND_INFO1获取 
func (obj *_V$sqlPlanStatisticsMgr) WithExtendInfo1(extendInfo1 string) Option {
	return optionFunc(func(o *options) { o.query["EXTEND_INFO1"] = extendInfo1 })
}

// WithExtendInfo2 EXTEND_INFO2获取 
func (obj *_V$sqlPlanStatisticsMgr) WithExtendInfo2(extendInfo2 string) Option {
	return optionFunc(func(o *options) { o.query["EXTEND_INFO2"] = extendInfo2 })
}


// GetByOption 功能选项模式获取
func (obj *_V$sqlPlanStatisticsMgr) GetByOption(opts ...Option) (result V$sqlPlanStatistics, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sqlPlanStatisticsMgr) GetByOptions(opts ...Option) (results []*V$sqlPlanStatistics, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sqlPlanStatisticsMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sqlPlanStatistics,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where(options.query)
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
func (obj *_V$sqlPlanStatisticsMgr) GetFromConID(conID int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromConID(conIDs []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromSvrIP(svrIP string) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromSvrPort(svrPort int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过PLAN_ID获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromPlanID(planID int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`PLAN_ID` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromPlanID(planIDs []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`PLAN_ID` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromOperationID 通过OPERATION_ID获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromOperationID(operationID int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`OPERATION_ID` = ?", operationID).Find(&results).Error
	
	return
}

// GetBatchFromOperationID 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromOperationID(operationIDs []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`OPERATION_ID` IN (?)", operationIDs).Find(&results).Error
	
	return
}
 
// GetFromExecutions 通过EXECUTIONS获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromExecutions(executions int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`EXECUTIONS` = ?", executions).Find(&results).Error
	
	return
}

// GetBatchFromExecutions 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromExecutions(executionss []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`EXECUTIONS` IN (?)", executionss).Find(&results).Error
	
	return
}
 
// GetFromOutputRows 通过OUTPUT_ROWS获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromOutputRows(outputRows int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`OUTPUT_ROWS` = ?", outputRows).Find(&results).Error
	
	return
}

// GetBatchFromOutputRows 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromOutputRows(outputRowss []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`OUTPUT_ROWS` IN (?)", outputRowss).Find(&results).Error
	
	return
}
 
// GetFromInputRows 通过INPUT_ROWS获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromInputRows(inputRows int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`INPUT_ROWS` = ?", inputRows).Find(&results).Error
	
	return
}

// GetBatchFromInputRows 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromInputRows(inputRowss []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`INPUT_ROWS` IN (?)", inputRowss).Find(&results).Error
	
	return
}
 
// GetFromRescanTimes 通过RESCAN_TIMES获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromRescanTimes(rescanTimes int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`RESCAN_TIMES` = ?", rescanTimes).Find(&results).Error
	
	return
}

// GetBatchFromRescanTimes 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromRescanTimes(rescanTimess []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`RESCAN_TIMES` IN (?)", rescanTimess).Find(&results).Error
	
	return
}
 
// GetFromBufferGets 通过BUFFER_GETS获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromBufferGets(bufferGets int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`BUFFER_GETS` = ?", bufferGets).Find(&results).Error
	
	return
}

// GetBatchFromBufferGets 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromBufferGets(bufferGetss []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`BUFFER_GETS` IN (?)", bufferGetss).Find(&results).Error
	
	return
}
 
// GetFromDiskReads 通过DISK_READS获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromDiskReads(diskReads int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`DISK_READS` = ?", diskReads).Find(&results).Error
	
	return
}

// GetBatchFromDiskReads 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromDiskReads(diskReadss []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`DISK_READS` IN (?)", diskReadss).Find(&results).Error
	
	return
}
 
// GetFromDiskWrites 通过DISK_WRITES获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromDiskWrites(diskWrites int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`DISK_WRITES` = ?", diskWrites).Find(&results).Error
	
	return
}

// GetBatchFromDiskWrites 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromDiskWrites(diskWritess []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`DISK_WRITES` IN (?)", diskWritess).Find(&results).Error
	
	return
}
 
// GetFromElapsedTime 通过ELAPSED_TIME获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromElapsedTime(elapsedTime int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`ELAPSED_TIME` = ?", elapsedTime).Find(&results).Error
	
	return
}

// GetBatchFromElapsedTime 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromElapsedTime(elapsedTimes []int64) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`ELAPSED_TIME` IN (?)", elapsedTimes).Find(&results).Error
	
	return
}
 
// GetFromExtendInfo1 通过EXTEND_INFO1获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromExtendInfo1(extendInfo1 string) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`EXTEND_INFO1` = ?", extendInfo1).Find(&results).Error
	
	return
}

// GetBatchFromExtendInfo1 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromExtendInfo1(extendInfo1s []string) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`EXTEND_INFO1` IN (?)", extendInfo1s).Find(&results).Error
	
	return
}
 
// GetFromExtendInfo2 通过EXTEND_INFO2获取内容  
func (obj *_V$sqlPlanStatisticsMgr) GetFromExtendInfo2(extendInfo2 string) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`EXTEND_INFO2` = ?", extendInfo2).Find(&results).Error
	
	return
}

// GetBatchFromExtendInfo2 批量查找 
func (obj *_V$sqlPlanStatisticsMgr) GetBatchFromExtendInfo2(extendInfo2s []string) (results []*V$sqlPlanStatistics, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanStatistics{}).Where("`EXTEND_INFO2` IN (?)", extendInfo2s).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

