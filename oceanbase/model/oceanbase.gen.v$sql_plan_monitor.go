package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _V$sqlPlanMonitorMgr struct {
	*_BaseMgr
}

// V$sqlPlanMonitorMgr open func
func V$sqlPlanMonitorMgr(db *gorm.DB) *_V$sqlPlanMonitorMgr {
	if db == nil {
		panic(fmt.Errorf("V$sqlPlanMonitorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sqlPlanMonitorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sql_plan_monitor"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sqlPlanMonitorMgr) GetTableName() string {
	return "v$sql_plan_monitor"
}

// Reset 重置gorm会话
func (obj *_V$sqlPlanMonitorMgr) Reset() *_V$sqlPlanMonitorMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sqlPlanMonitorMgr) Get() (result V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sqlPlanMonitorMgr) Gets() (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sqlPlanMonitorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithRequestID REQUEST_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithRequestID(requestID int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST_ID"] = requestID })
}

// WithKey KEY获取 
func (obj *_V$sqlPlanMonitorMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["KEY"] = key })
}

// WithStatus STATUS获取 
func (obj *_V$sqlPlanMonitorMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithSvrIP SVR_IP获取 
func (obj *_V$sqlPlanMonitorMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$sqlPlanMonitorMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTraceID TRACE_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["TRACE_ID"] = traceID })
}

// WithFirstRefreshTime FIRST_REFRESH_TIME获取 
func (obj *_V$sqlPlanMonitorMgr) WithFirstRefreshTime(firstRefreshTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["FIRST_REFRESH_TIME"] = firstRefreshTime })
}

// WithLastRefreshTime LAST_REFRESH_TIME获取 
func (obj *_V$sqlPlanMonitorMgr) WithLastRefreshTime(lastRefreshTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_REFRESH_TIME"] = lastRefreshTime })
}

// WithFirstChangeTime FIRST_CHANGE_TIME获取 
func (obj *_V$sqlPlanMonitorMgr) WithFirstChangeTime(firstChangeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["FIRST_CHANGE_TIME"] = firstChangeTime })
}

// WithLastChangeTime LAST_CHANGE_TIME获取 
func (obj *_V$sqlPlanMonitorMgr) WithLastChangeTime(lastChangeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_CHANGE_TIME"] = lastChangeTime })
}

// WithRefreshCount REFRESH_COUNT获取 
func (obj *_V$sqlPlanMonitorMgr) WithRefreshCount(refreshCount string) Option {
	return optionFunc(func(o *options) { o.query["REFRESH_COUNT"] = refreshCount })
}

// WithSid SID获取 
func (obj *_V$sqlPlanMonitorMgr) WithSid(sid string) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithProcessName PROCESS_NAME获取 
func (obj *_V$sqlPlanMonitorMgr) WithProcessName(processName int64) Option {
	return optionFunc(func(o *options) { o.query["PROCESS_NAME"] = processName })
}

// WithSQLID SQL_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["SQL_ID"] = sqlID })
}

// WithSQLExecStart SQL_EXEC_START获取 
func (obj *_V$sqlPlanMonitorMgr) WithSQLExecStart(sqlExecStart string) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_START"] = sqlExecStart })
}

// WithSQLExecID SQL_EXEC_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithSQLExecID(sqlExecID string) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_ID"] = sqlExecID })
}

// WithSQLPlanHashValue SQL_PLAN_HASH_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithSQLPlanHashValue(sqlPlanHashValue string) Option {
	return optionFunc(func(o *options) { o.query["SQL_PLAN_HASH_VALUE"] = sqlPlanHashValue })
}

// WithSQLChildAddress SQL_CHILD_ADDRESS获取 
func (obj *_V$sqlPlanMonitorMgr) WithSQLChildAddress(sqlChildAddress string) Option {
	return optionFunc(func(o *options) { o.query["SQL_CHILD_ADDRESS"] = sqlChildAddress })
}

// WithPlanParentID PLAN_PARENT_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanParentID(planParentID string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_PARENT_ID"] = planParentID })
}

// WithPlanLineID PLAN_LINE_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanLineID(planLineID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_LINE_ID"] = planLineID })
}

// WithPlanOperation PLAN_OPERATION获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanOperation(planOperation string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_OPERATION"] = planOperation })
}

// WithPlanOptions PLAN_OPTIONS获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanOptions(planOptions string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_OPTIONS"] = planOptions })
}

// WithPlanObjectOwner PLAN_OBJECT_OWNER获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanObjectOwner(planObjectOwner string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_OBJECT_OWNER"] = planObjectOwner })
}

// WithPlanObjectName PLAN_OBJECT_NAME获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanObjectName(planObjectName string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_OBJECT_NAME"] = planObjectName })
}

// WithPlanObjectType PLAN_OBJECT_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanObjectType(planObjectType string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_OBJECT_TYPE"] = planObjectType })
}

// WithPlanDepth PLAN_DEPTH获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanDepth(planDepth int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_DEPTH"] = planDepth })
}

// WithPlanPosition PLAN_POSITION获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanPosition(planPosition string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_POSITION"] = planPosition })
}

// WithPlanCost PLAN_COST获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanCost(planCost string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_COST"] = planCost })
}

// WithPlanCardinality PLAN_CARDINALITY获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanCardinality(planCardinality string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_CARDINALITY"] = planCardinality })
}

// WithPlanBytes PLAN_BYTES获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanBytes(planBytes string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_BYTES"] = planBytes })
}

// WithPlanTime PLAN_TIME获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanTime(planTime string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_TIME"] = planTime })
}

// WithPlanPartitionStart PLAN_PARTITION_START获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanPartitionStart(planPartitionStart string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_PARTITION_START"] = planPartitionStart })
}

// WithPlanPartitionStop PLAN_PARTITION_STOP获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanPartitionStop(planPartitionStop string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_PARTITION_STOP"] = planPartitionStop })
}

// WithPlanCPUCost PLAN_CPU_COST获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanCPUCost(planCPUCost string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_CPU_COST"] = planCPUCost })
}

// WithPlanIoCost PLAN_IO_COST获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanIoCost(planIoCost string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_IO_COST"] = planIoCost })
}

// WithPlanTempSpace PLAN_TEMP_SPACE获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanTempSpace(planTempSpace string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_TEMP_SPACE"] = planTempSpace })
}

// WithStarts STARTS获取 
func (obj *_V$sqlPlanMonitorMgr) WithStarts(starts int64) Option {
	return optionFunc(func(o *options) { o.query["STARTS"] = starts })
}

// WithOutputRows OUTPUT_ROWS获取 
func (obj *_V$sqlPlanMonitorMgr) WithOutputRows(outputRows int64) Option {
	return optionFunc(func(o *options) { o.query["OUTPUT_ROWS"] = outputRows })
}

// WithIoInterconnectBytes IO_INTERCONNECT_BYTES获取 
func (obj *_V$sqlPlanMonitorMgr) WithIoInterconnectBytes(ioInterconnectBytes string) Option {
	return optionFunc(func(o *options) { o.query["IO_INTERCONNECT_BYTES"] = ioInterconnectBytes })
}

// WithPhysicalReadRequests PHYSICAL_READ_REQUESTS获取 
func (obj *_V$sqlPlanMonitorMgr) WithPhysicalReadRequests(physicalReadRequests string) Option {
	return optionFunc(func(o *options) { o.query["PHYSICAL_READ_REQUESTS"] = physicalReadRequests })
}

// WithPhysicalReadBytes PHYSICAL_READ_BYTES获取 
func (obj *_V$sqlPlanMonitorMgr) WithPhysicalReadBytes(physicalReadBytes string) Option {
	return optionFunc(func(o *options) { o.query["PHYSICAL_READ_BYTES"] = physicalReadBytes })
}

// WithPhysicalWriteRequests PHYSICAL_WRITE_REQUESTS获取 
func (obj *_V$sqlPlanMonitorMgr) WithPhysicalWriteRequests(physicalWriteRequests string) Option {
	return optionFunc(func(o *options) { o.query["PHYSICAL_WRITE_REQUESTS"] = physicalWriteRequests })
}

// WithPhysicalWriteBytes PHYSICAL_WRITE_BYTES获取 
func (obj *_V$sqlPlanMonitorMgr) WithPhysicalWriteBytes(physicalWriteBytes string) Option {
	return optionFunc(func(o *options) { o.query["PHYSICAL_WRITE_BYTES"] = physicalWriteBytes })
}

// WithWorkareaMem WORKAREA_MEM获取 
func (obj *_V$sqlPlanMonitorMgr) WithWorkareaMem(workareaMem string) Option {
	return optionFunc(func(o *options) { o.query["WORKAREA_MEM"] = workareaMem })
}

// WithWorkareaMaxMem WORKAREA_MAX_MEM获取 
func (obj *_V$sqlPlanMonitorMgr) WithWorkareaMaxMem(workareaMaxMem string) Option {
	return optionFunc(func(o *options) { o.query["WORKAREA_MAX_MEM"] = workareaMaxMem })
}

// WithWorkareaTempseg WORKAREA_TEMPSEG获取 
func (obj *_V$sqlPlanMonitorMgr) WithWorkareaTempseg(workareaTempseg string) Option {
	return optionFunc(func(o *options) { o.query["WORKAREA_TEMPSEG"] = workareaTempseg })
}

// WithWorkareaMaxTempseg WORKAREA_MAX_TEMPSEG获取 
func (obj *_V$sqlPlanMonitorMgr) WithWorkareaMaxTempseg(workareaMaxTempseg string) Option {
	return optionFunc(func(o *options) { o.query["WORKAREA_MAX_TEMPSEG"] = workareaMaxTempseg })
}

// WithOtherstatGroupID OTHERSTAT_GROUP_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstatGroupID(otherstatGroupID string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_GROUP_ID"] = otherstatGroupID })
}

// WithOtherstat1ID OTHERSTAT_1_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat1ID(otherstat1ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_1_ID"] = otherstat1ID })
}

// WithOtherstat1Type OTHERSTAT_1_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat1Type(otherstat1Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_1_TYPE"] = otherstat1Type })
}

// WithOtherstat1Value OTHERSTAT_1_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat1Value(otherstat1Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_1_VALUE"] = otherstat1Value })
}

// WithOtherstat2ID OTHERSTAT_2_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat2ID(otherstat2ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_2_ID"] = otherstat2ID })
}

// WithOtherstat2Type OTHERSTAT_2_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat2Type(otherstat2Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_2_TYPE"] = otherstat2Type })
}

// WithOtherstat2Value OTHERSTAT_2_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat2Value(otherstat2Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_2_VALUE"] = otherstat2Value })
}

// WithOtherstat3ID OTHERSTAT_3_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat3ID(otherstat3ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_3_ID"] = otherstat3ID })
}

// WithOtherstat3Type OTHERSTAT_3_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat3Type(otherstat3Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_3_TYPE"] = otherstat3Type })
}

// WithOtherstat3Value OTHERSTAT_3_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat3Value(otherstat3Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_3_VALUE"] = otherstat3Value })
}

// WithOtherstat4ID OTHERSTAT_4_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat4ID(otherstat4ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_4_ID"] = otherstat4ID })
}

// WithOtherstat4Type OTHERSTAT_4_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat4Type(otherstat4Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_4_TYPE"] = otherstat4Type })
}

// WithOtherstat4Value OTHERSTAT_4_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat4Value(otherstat4Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_4_VALUE"] = otherstat4Value })
}

// WithOtherstat5ID OTHERSTAT_5_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat5ID(otherstat5ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_5_ID"] = otherstat5ID })
}

// WithOtherstat5Type OTHERSTAT_5_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat5Type(otherstat5Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_5_TYPE"] = otherstat5Type })
}

// WithOtherstat5Value OTHERSTAT_5_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat5Value(otherstat5Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_5_VALUE"] = otherstat5Value })
}

// WithOtherstat6ID OTHERSTAT_6_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat6ID(otherstat6ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_6_ID"] = otherstat6ID })
}

// WithOtherstat6Type OTHERSTAT_6_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat6Type(otherstat6Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_6_TYPE"] = otherstat6Type })
}

// WithOtherstat6Value OTHERSTAT_6_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat6Value(otherstat6Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_6_VALUE"] = otherstat6Value })
}

// WithOtherstat7ID OTHERSTAT_7_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat7ID(otherstat7ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_7_ID"] = otherstat7ID })
}

// WithOtherstat7Type OTHERSTAT_7_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat7Type(otherstat7Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_7_TYPE"] = otherstat7Type })
}

// WithOtherstat7Value OTHERSTAT_7_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat7Value(otherstat7Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_7_VALUE"] = otherstat7Value })
}

// WithOtherstat8ID OTHERSTAT_8_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat8ID(otherstat8ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_8_ID"] = otherstat8ID })
}

// WithOtherstat8Type OTHERSTAT_8_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat8Type(otherstat8Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_8_TYPE"] = otherstat8Type })
}

// WithOtherstat8Value OTHERSTAT_8_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat8Value(otherstat8Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_8_VALUE"] = otherstat8Value })
}

// WithOtherstat9ID OTHERSTAT_9_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat9ID(otherstat9ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_9_ID"] = otherstat9ID })
}

// WithOtherstat9Type OTHERSTAT_9_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat9Type(otherstat9Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_9_TYPE"] = otherstat9Type })
}

// WithOtherstat9Value OTHERSTAT_9_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat9Value(otherstat9Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_9_VALUE"] = otherstat9Value })
}

// WithOtherstat10ID OTHERSTAT_10_ID获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat10ID(otherstat10ID int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_10_ID"] = otherstat10ID })
}

// WithOtherstat10Type OTHERSTAT_10_TYPE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat10Type(otherstat10Type string) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_10_TYPE"] = otherstat10Type })
}

// WithOtherstat10Value OTHERSTAT_10_VALUE获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherstat10Value(otherstat10Value int64) Option {
	return optionFunc(func(o *options) { o.query["OTHERSTAT_10_VALUE"] = otherstat10Value })
}

// WithOtherXML OTHER_XML获取 
func (obj *_V$sqlPlanMonitorMgr) WithOtherXML(otherXML string) Option {
	return optionFunc(func(o *options) { o.query["OTHER_XML"] = otherXML })
}

// WithPlanOperationInactive PLAN_OPERATION_INACTIVE获取 
func (obj *_V$sqlPlanMonitorMgr) WithPlanOperationInactive(planOperationInactive string) Option {
	return optionFunc(func(o *options) { o.query["PLAN_OPERATION_INACTIVE"] = planOperationInactive })
}


// GetByOption 功能选项模式获取
func (obj *_V$sqlPlanMonitorMgr) GetByOption(opts ...Option) (result V$sqlPlanMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sqlPlanMonitorMgr) GetByOptions(opts ...Option) (results []*V$sqlPlanMonitor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sqlPlanMonitorMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sqlPlanMonitor,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where(options.query)
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
func (obj *_V$sqlPlanMonitorMgr) GetFromConID(conID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromConID(conIDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromRequestID 通过REQUEST_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromRequestID(requestID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`REQUEST_ID` = ?", requestID).Find(&results).Error
	
	return
}

// GetBatchFromRequestID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromRequestID(requestIDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`REQUEST_ID` IN (?)", requestIDs).Find(&results).Error
	
	return
}
 
// GetFromKey 通过KEY获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromKey(key string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`KEY` = ?", key).Find(&results).Error
	
	return
}

// GetBatchFromKey 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromKey(keys []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`KEY` IN (?)", keys).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过STATUS获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromStatus(status string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`STATUS` = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromStatus(statuss []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSvrIP(svrIP string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSvrPort(svrPort int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTraceID 通过TRACE_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromTraceID(traceID string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`TRACE_ID` = ?", traceID).Find(&results).Error
	
	return
}

// GetBatchFromTraceID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromTraceID(traceIDs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`TRACE_ID` IN (?)", traceIDs).Find(&results).Error
	
	return
}
 
// GetFromFirstRefreshTime 通过FIRST_REFRESH_TIME获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromFirstRefreshTime(firstRefreshTime time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`FIRST_REFRESH_TIME` = ?", firstRefreshTime).Find(&results).Error
	
	return
}

// GetBatchFromFirstRefreshTime 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromFirstRefreshTime(firstRefreshTimes []time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`FIRST_REFRESH_TIME` IN (?)", firstRefreshTimes).Find(&results).Error
	
	return
}
 
// GetFromLastRefreshTime 通过LAST_REFRESH_TIME获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromLastRefreshTime(lastRefreshTime time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`LAST_REFRESH_TIME` = ?", lastRefreshTime).Find(&results).Error
	
	return
}

// GetBatchFromLastRefreshTime 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromLastRefreshTime(lastRefreshTimes []time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`LAST_REFRESH_TIME` IN (?)", lastRefreshTimes).Find(&results).Error
	
	return
}
 
// GetFromFirstChangeTime 通过FIRST_CHANGE_TIME获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromFirstChangeTime(firstChangeTime time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`FIRST_CHANGE_TIME` = ?", firstChangeTime).Find(&results).Error
	
	return
}

// GetBatchFromFirstChangeTime 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromFirstChangeTime(firstChangeTimes []time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`FIRST_CHANGE_TIME` IN (?)", firstChangeTimes).Find(&results).Error
	
	return
}
 
// GetFromLastChangeTime 通过LAST_CHANGE_TIME获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromLastChangeTime(lastChangeTime time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`LAST_CHANGE_TIME` = ?", lastChangeTime).Find(&results).Error
	
	return
}

// GetBatchFromLastChangeTime 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromLastChangeTime(lastChangeTimes []time.Time) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`LAST_CHANGE_TIME` IN (?)", lastChangeTimes).Find(&results).Error
	
	return
}
 
// GetFromRefreshCount 通过REFRESH_COUNT获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromRefreshCount(refreshCount string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`REFRESH_COUNT` = ?", refreshCount).Find(&results).Error
	
	return
}

// GetBatchFromRefreshCount 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromRefreshCount(refreshCounts []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`REFRESH_COUNT` IN (?)", refreshCounts).Find(&results).Error
	
	return
}
 
// GetFromSid 通过SID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSid(sid string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSid(sids []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromProcessName 通过PROCESS_NAME获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromProcessName(processName int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PROCESS_NAME` = ?", processName).Find(&results).Error
	
	return
}

// GetBatchFromProcessName 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromProcessName(processNames []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PROCESS_NAME` IN (?)", processNames).Find(&results).Error
	
	return
}
 
// GetFromSQLID 通过SQL_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSQLID(sqlID string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_ID` = ?", sqlID).Find(&results).Error
	
	return
}

// GetBatchFromSQLID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSQLID(sqlIDs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_ID` IN (?)", sqlIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLExecStart 通过SQL_EXEC_START获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSQLExecStart(sqlExecStart string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_EXEC_START` = ?", sqlExecStart).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecStart 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSQLExecStart(sqlExecStarts []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_EXEC_START` IN (?)", sqlExecStarts).Find(&results).Error
	
	return
}
 
// GetFromSQLExecID 通过SQL_EXEC_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSQLExecID(sqlExecID string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_EXEC_ID` = ?", sqlExecID).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSQLExecID(sqlExecIDs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_EXEC_ID` IN (?)", sqlExecIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLPlanHashValue 通过SQL_PLAN_HASH_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSQLPlanHashValue(sqlPlanHashValue string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_PLAN_HASH_VALUE` = ?", sqlPlanHashValue).Find(&results).Error
	
	return
}

// GetBatchFromSQLPlanHashValue 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSQLPlanHashValue(sqlPlanHashValues []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_PLAN_HASH_VALUE` IN (?)", sqlPlanHashValues).Find(&results).Error
	
	return
}
 
// GetFromSQLChildAddress 通过SQL_CHILD_ADDRESS获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromSQLChildAddress(sqlChildAddress string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_CHILD_ADDRESS` = ?", sqlChildAddress).Find(&results).Error
	
	return
}

// GetBatchFromSQLChildAddress 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromSQLChildAddress(sqlChildAddresss []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`SQL_CHILD_ADDRESS` IN (?)", sqlChildAddresss).Find(&results).Error
	
	return
}
 
// GetFromPlanParentID 通过PLAN_PARENT_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanParentID(planParentID string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_PARENT_ID` = ?", planParentID).Find(&results).Error
	
	return
}

// GetBatchFromPlanParentID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanParentID(planParentIDs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_PARENT_ID` IN (?)", planParentIDs).Find(&results).Error
	
	return
}
 
// GetFromPlanLineID 通过PLAN_LINE_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanLineID(planLineID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_LINE_ID` = ?", planLineID).Find(&results).Error
	
	return
}

// GetBatchFromPlanLineID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanLineID(planLineIDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_LINE_ID` IN (?)", planLineIDs).Find(&results).Error
	
	return
}
 
// GetFromPlanOperation 通过PLAN_OPERATION获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanOperation(planOperation string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OPERATION` = ?", planOperation).Find(&results).Error
	
	return
}

// GetBatchFromPlanOperation 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanOperation(planOperations []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OPERATION` IN (?)", planOperations).Find(&results).Error
	
	return
}
 
// GetFromPlanOptions 通过PLAN_OPTIONS获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanOptions(planOptions string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OPTIONS` = ?", planOptions).Find(&results).Error
	
	return
}

// GetBatchFromPlanOptions 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanOptions(planOptionss []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OPTIONS` IN (?)", planOptionss).Find(&results).Error
	
	return
}
 
// GetFromPlanObjectOwner 通过PLAN_OBJECT_OWNER获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanObjectOwner(planObjectOwner string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OBJECT_OWNER` = ?", planObjectOwner).Find(&results).Error
	
	return
}

// GetBatchFromPlanObjectOwner 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanObjectOwner(planObjectOwners []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OBJECT_OWNER` IN (?)", planObjectOwners).Find(&results).Error
	
	return
}
 
// GetFromPlanObjectName 通过PLAN_OBJECT_NAME获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanObjectName(planObjectName string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OBJECT_NAME` = ?", planObjectName).Find(&results).Error
	
	return
}

// GetBatchFromPlanObjectName 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanObjectName(planObjectNames []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OBJECT_NAME` IN (?)", planObjectNames).Find(&results).Error
	
	return
}
 
// GetFromPlanObjectType 通过PLAN_OBJECT_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanObjectType(planObjectType string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OBJECT_TYPE` = ?", planObjectType).Find(&results).Error
	
	return
}

// GetBatchFromPlanObjectType 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanObjectType(planObjectTypes []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OBJECT_TYPE` IN (?)", planObjectTypes).Find(&results).Error
	
	return
}
 
// GetFromPlanDepth 通过PLAN_DEPTH获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanDepth(planDepth int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_DEPTH` = ?", planDepth).Find(&results).Error
	
	return
}

// GetBatchFromPlanDepth 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanDepth(planDepths []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_DEPTH` IN (?)", planDepths).Find(&results).Error
	
	return
}
 
// GetFromPlanPosition 通过PLAN_POSITION获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanPosition(planPosition string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_POSITION` = ?", planPosition).Find(&results).Error
	
	return
}

// GetBatchFromPlanPosition 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanPosition(planPositions []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_POSITION` IN (?)", planPositions).Find(&results).Error
	
	return
}
 
// GetFromPlanCost 通过PLAN_COST获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanCost(planCost string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_COST` = ?", planCost).Find(&results).Error
	
	return
}

// GetBatchFromPlanCost 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanCost(planCosts []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_COST` IN (?)", planCosts).Find(&results).Error
	
	return
}
 
// GetFromPlanCardinality 通过PLAN_CARDINALITY获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanCardinality(planCardinality string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_CARDINALITY` = ?", planCardinality).Find(&results).Error
	
	return
}

// GetBatchFromPlanCardinality 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanCardinality(planCardinalitys []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_CARDINALITY` IN (?)", planCardinalitys).Find(&results).Error
	
	return
}
 
// GetFromPlanBytes 通过PLAN_BYTES获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanBytes(planBytes string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_BYTES` = ?", planBytes).Find(&results).Error
	
	return
}

// GetBatchFromPlanBytes 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanBytes(planBytess []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_BYTES` IN (?)", planBytess).Find(&results).Error
	
	return
}
 
// GetFromPlanTime 通过PLAN_TIME获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanTime(planTime string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_TIME` = ?", planTime).Find(&results).Error
	
	return
}

// GetBatchFromPlanTime 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanTime(planTimes []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_TIME` IN (?)", planTimes).Find(&results).Error
	
	return
}
 
// GetFromPlanPartitionStart 通过PLAN_PARTITION_START获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanPartitionStart(planPartitionStart string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_PARTITION_START` = ?", planPartitionStart).Find(&results).Error
	
	return
}

// GetBatchFromPlanPartitionStart 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanPartitionStart(planPartitionStarts []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_PARTITION_START` IN (?)", planPartitionStarts).Find(&results).Error
	
	return
}
 
// GetFromPlanPartitionStop 通过PLAN_PARTITION_STOP获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanPartitionStop(planPartitionStop string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_PARTITION_STOP` = ?", planPartitionStop).Find(&results).Error
	
	return
}

// GetBatchFromPlanPartitionStop 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanPartitionStop(planPartitionStops []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_PARTITION_STOP` IN (?)", planPartitionStops).Find(&results).Error
	
	return
}
 
// GetFromPlanCPUCost 通过PLAN_CPU_COST获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanCPUCost(planCPUCost string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_CPU_COST` = ?", planCPUCost).Find(&results).Error
	
	return
}

// GetBatchFromPlanCPUCost 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanCPUCost(planCPUCosts []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_CPU_COST` IN (?)", planCPUCosts).Find(&results).Error
	
	return
}
 
// GetFromPlanIoCost 通过PLAN_IO_COST获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanIoCost(planIoCost string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_IO_COST` = ?", planIoCost).Find(&results).Error
	
	return
}

// GetBatchFromPlanIoCost 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanIoCost(planIoCosts []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_IO_COST` IN (?)", planIoCosts).Find(&results).Error
	
	return
}
 
// GetFromPlanTempSpace 通过PLAN_TEMP_SPACE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanTempSpace(planTempSpace string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_TEMP_SPACE` = ?", planTempSpace).Find(&results).Error
	
	return
}

// GetBatchFromPlanTempSpace 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanTempSpace(planTempSpaces []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_TEMP_SPACE` IN (?)", planTempSpaces).Find(&results).Error
	
	return
}
 
// GetFromStarts 通过STARTS获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromStarts(starts int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`STARTS` = ?", starts).Find(&results).Error
	
	return
}

// GetBatchFromStarts 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromStarts(startss []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`STARTS` IN (?)", startss).Find(&results).Error
	
	return
}
 
// GetFromOutputRows 通过OUTPUT_ROWS获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOutputRows(outputRows int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OUTPUT_ROWS` = ?", outputRows).Find(&results).Error
	
	return
}

// GetBatchFromOutputRows 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOutputRows(outputRowss []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OUTPUT_ROWS` IN (?)", outputRowss).Find(&results).Error
	
	return
}
 
// GetFromIoInterconnectBytes 通过IO_INTERCONNECT_BYTES获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromIoInterconnectBytes(ioInterconnectBytes string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`IO_INTERCONNECT_BYTES` = ?", ioInterconnectBytes).Find(&results).Error
	
	return
}

// GetBatchFromIoInterconnectBytes 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromIoInterconnectBytes(ioInterconnectBytess []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`IO_INTERCONNECT_BYTES` IN (?)", ioInterconnectBytess).Find(&results).Error
	
	return
}
 
// GetFromPhysicalReadRequests 通过PHYSICAL_READ_REQUESTS获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPhysicalReadRequests(physicalReadRequests string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_READ_REQUESTS` = ?", physicalReadRequests).Find(&results).Error
	
	return
}

// GetBatchFromPhysicalReadRequests 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPhysicalReadRequests(physicalReadRequestss []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_READ_REQUESTS` IN (?)", physicalReadRequestss).Find(&results).Error
	
	return
}
 
// GetFromPhysicalReadBytes 通过PHYSICAL_READ_BYTES获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPhysicalReadBytes(physicalReadBytes string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_READ_BYTES` = ?", physicalReadBytes).Find(&results).Error
	
	return
}

// GetBatchFromPhysicalReadBytes 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPhysicalReadBytes(physicalReadBytess []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_READ_BYTES` IN (?)", physicalReadBytess).Find(&results).Error
	
	return
}
 
// GetFromPhysicalWriteRequests 通过PHYSICAL_WRITE_REQUESTS获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPhysicalWriteRequests(physicalWriteRequests string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_WRITE_REQUESTS` = ?", physicalWriteRequests).Find(&results).Error
	
	return
}

// GetBatchFromPhysicalWriteRequests 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPhysicalWriteRequests(physicalWriteRequestss []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_WRITE_REQUESTS` IN (?)", physicalWriteRequestss).Find(&results).Error
	
	return
}
 
// GetFromPhysicalWriteBytes 通过PHYSICAL_WRITE_BYTES获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPhysicalWriteBytes(physicalWriteBytes string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_WRITE_BYTES` = ?", physicalWriteBytes).Find(&results).Error
	
	return
}

// GetBatchFromPhysicalWriteBytes 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPhysicalWriteBytes(physicalWriteBytess []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PHYSICAL_WRITE_BYTES` IN (?)", physicalWriteBytess).Find(&results).Error
	
	return
}
 
// GetFromWorkareaMem 通过WORKAREA_MEM获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromWorkareaMem(workareaMem string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_MEM` = ?", workareaMem).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaMem 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromWorkareaMem(workareaMems []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_MEM` IN (?)", workareaMems).Find(&results).Error
	
	return
}
 
// GetFromWorkareaMaxMem 通过WORKAREA_MAX_MEM获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromWorkareaMaxMem(workareaMaxMem string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_MAX_MEM` = ?", workareaMaxMem).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaMaxMem 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromWorkareaMaxMem(workareaMaxMems []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_MAX_MEM` IN (?)", workareaMaxMems).Find(&results).Error
	
	return
}
 
// GetFromWorkareaTempseg 通过WORKAREA_TEMPSEG获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromWorkareaTempseg(workareaTempseg string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_TEMPSEG` = ?", workareaTempseg).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaTempseg 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromWorkareaTempseg(workareaTempsegs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_TEMPSEG` IN (?)", workareaTempsegs).Find(&results).Error
	
	return
}
 
// GetFromWorkareaMaxTempseg 通过WORKAREA_MAX_TEMPSEG获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromWorkareaMaxTempseg(workareaMaxTempseg string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_MAX_TEMPSEG` = ?", workareaMaxTempseg).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaMaxTempseg 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromWorkareaMaxTempseg(workareaMaxTempsegs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`WORKAREA_MAX_TEMPSEG` IN (?)", workareaMaxTempsegs).Find(&results).Error
	
	return
}
 
// GetFromOtherstatGroupID 通过OTHERSTAT_GROUP_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstatGroupID(otherstatGroupID string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_GROUP_ID` = ?", otherstatGroupID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstatGroupID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstatGroupID(otherstatGroupIDs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_GROUP_ID` IN (?)", otherstatGroupIDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat1ID 通过OTHERSTAT_1_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat1ID(otherstat1ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_1_ID` = ?", otherstat1ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat1ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat1ID(otherstat1IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_1_ID` IN (?)", otherstat1IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat1Type 通过OTHERSTAT_1_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat1Type(otherstat1Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_1_TYPE` = ?", otherstat1Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat1Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat1Type(otherstat1Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_1_TYPE` IN (?)", otherstat1Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat1Value 通过OTHERSTAT_1_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat1Value(otherstat1Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_1_VALUE` = ?", otherstat1Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat1Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat1Value(otherstat1Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_1_VALUE` IN (?)", otherstat1Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat2ID 通过OTHERSTAT_2_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat2ID(otherstat2ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_2_ID` = ?", otherstat2ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat2ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat2ID(otherstat2IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_2_ID` IN (?)", otherstat2IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat2Type 通过OTHERSTAT_2_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat2Type(otherstat2Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_2_TYPE` = ?", otherstat2Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat2Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat2Type(otherstat2Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_2_TYPE` IN (?)", otherstat2Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat2Value 通过OTHERSTAT_2_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat2Value(otherstat2Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_2_VALUE` = ?", otherstat2Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat2Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat2Value(otherstat2Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_2_VALUE` IN (?)", otherstat2Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat3ID 通过OTHERSTAT_3_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat3ID(otherstat3ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_3_ID` = ?", otherstat3ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat3ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat3ID(otherstat3IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_3_ID` IN (?)", otherstat3IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat3Type 通过OTHERSTAT_3_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat3Type(otherstat3Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_3_TYPE` = ?", otherstat3Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat3Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat3Type(otherstat3Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_3_TYPE` IN (?)", otherstat3Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat3Value 通过OTHERSTAT_3_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat3Value(otherstat3Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_3_VALUE` = ?", otherstat3Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat3Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat3Value(otherstat3Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_3_VALUE` IN (?)", otherstat3Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat4ID 通过OTHERSTAT_4_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat4ID(otherstat4ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_4_ID` = ?", otherstat4ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat4ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat4ID(otherstat4IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_4_ID` IN (?)", otherstat4IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat4Type 通过OTHERSTAT_4_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat4Type(otherstat4Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_4_TYPE` = ?", otherstat4Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat4Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat4Type(otherstat4Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_4_TYPE` IN (?)", otherstat4Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat4Value 通过OTHERSTAT_4_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat4Value(otherstat4Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_4_VALUE` = ?", otherstat4Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat4Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat4Value(otherstat4Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_4_VALUE` IN (?)", otherstat4Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat5ID 通过OTHERSTAT_5_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat5ID(otherstat5ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_5_ID` = ?", otherstat5ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat5ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat5ID(otherstat5IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_5_ID` IN (?)", otherstat5IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat5Type 通过OTHERSTAT_5_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat5Type(otherstat5Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_5_TYPE` = ?", otherstat5Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat5Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat5Type(otherstat5Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_5_TYPE` IN (?)", otherstat5Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat5Value 通过OTHERSTAT_5_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat5Value(otherstat5Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_5_VALUE` = ?", otherstat5Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat5Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat5Value(otherstat5Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_5_VALUE` IN (?)", otherstat5Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat6ID 通过OTHERSTAT_6_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat6ID(otherstat6ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_6_ID` = ?", otherstat6ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat6ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat6ID(otherstat6IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_6_ID` IN (?)", otherstat6IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat6Type 通过OTHERSTAT_6_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat6Type(otherstat6Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_6_TYPE` = ?", otherstat6Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat6Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat6Type(otherstat6Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_6_TYPE` IN (?)", otherstat6Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat6Value 通过OTHERSTAT_6_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat6Value(otherstat6Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_6_VALUE` = ?", otherstat6Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat6Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat6Value(otherstat6Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_6_VALUE` IN (?)", otherstat6Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat7ID 通过OTHERSTAT_7_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat7ID(otherstat7ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_7_ID` = ?", otherstat7ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat7ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat7ID(otherstat7IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_7_ID` IN (?)", otherstat7IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat7Type 通过OTHERSTAT_7_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat7Type(otherstat7Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_7_TYPE` = ?", otherstat7Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat7Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat7Type(otherstat7Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_7_TYPE` IN (?)", otherstat7Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat7Value 通过OTHERSTAT_7_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat7Value(otherstat7Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_7_VALUE` = ?", otherstat7Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat7Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat7Value(otherstat7Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_7_VALUE` IN (?)", otherstat7Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat8ID 通过OTHERSTAT_8_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat8ID(otherstat8ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_8_ID` = ?", otherstat8ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat8ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat8ID(otherstat8IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_8_ID` IN (?)", otherstat8IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat8Type 通过OTHERSTAT_8_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat8Type(otherstat8Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_8_TYPE` = ?", otherstat8Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat8Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat8Type(otherstat8Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_8_TYPE` IN (?)", otherstat8Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat8Value 通过OTHERSTAT_8_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat8Value(otherstat8Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_8_VALUE` = ?", otherstat8Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat8Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat8Value(otherstat8Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_8_VALUE` IN (?)", otherstat8Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat9ID 通过OTHERSTAT_9_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat9ID(otherstat9ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_9_ID` = ?", otherstat9ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat9ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat9ID(otherstat9IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_9_ID` IN (?)", otherstat9IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat9Type 通过OTHERSTAT_9_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat9Type(otherstat9Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_9_TYPE` = ?", otherstat9Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat9Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat9Type(otherstat9Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_9_TYPE` IN (?)", otherstat9Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat9Value 通过OTHERSTAT_9_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat9Value(otherstat9Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_9_VALUE` = ?", otherstat9Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat9Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat9Value(otherstat9Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_9_VALUE` IN (?)", otherstat9Values).Find(&results).Error
	
	return
}
 
// GetFromOtherstat10ID 通过OTHERSTAT_10_ID获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat10ID(otherstat10ID int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_10_ID` = ?", otherstat10ID).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat10ID 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat10ID(otherstat10IDs []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_10_ID` IN (?)", otherstat10IDs).Find(&results).Error
	
	return
}
 
// GetFromOtherstat10Type 通过OTHERSTAT_10_TYPE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat10Type(otherstat10Type string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_10_TYPE` = ?", otherstat10Type).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat10Type 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat10Type(otherstat10Types []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_10_TYPE` IN (?)", otherstat10Types).Find(&results).Error
	
	return
}
 
// GetFromOtherstat10Value 通过OTHERSTAT_10_VALUE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherstat10Value(otherstat10Value int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_10_VALUE` = ?", otherstat10Value).Find(&results).Error
	
	return
}

// GetBatchFromOtherstat10Value 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherstat10Value(otherstat10Values []int64) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHERSTAT_10_VALUE` IN (?)", otherstat10Values).Find(&results).Error
	
	return
}
 
// GetFromOtherXML 通过OTHER_XML获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromOtherXML(otherXML string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHER_XML` = ?", otherXML).Find(&results).Error
	
	return
}

// GetBatchFromOtherXML 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromOtherXML(otherXMLs []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`OTHER_XML` IN (?)", otherXMLs).Find(&results).Error
	
	return
}
 
// GetFromPlanOperationInactive 通过PLAN_OPERATION_INACTIVE获取内容  
func (obj *_V$sqlPlanMonitorMgr) GetFromPlanOperationInactive(planOperationInactive string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OPERATION_INACTIVE` = ?", planOperationInactive).Find(&results).Error
	
	return
}

// GetBatchFromPlanOperationInactive 批量查找 
func (obj *_V$sqlPlanMonitorMgr) GetBatchFromPlanOperationInactive(planOperationInactives []string) (results []*V$sqlPlanMonitor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlPlanMonitor{}).Where("`PLAN_OPERATION_INACTIVE` IN (?)", planOperationInactives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

