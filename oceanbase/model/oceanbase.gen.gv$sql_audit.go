package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _Gv$sqlAuditMgr struct {
	*_BaseMgr
}

// Gv$sqlAuditMgr open func
func Gv$sqlAuditMgr(db *gorm.DB) *_Gv$sqlAuditMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$sqlAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$sqlAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$sql_audit"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$sqlAuditMgr) GetTableName() string {
	return "gv$sql_audit"
}

// Reset 重置gorm会话
func (obj *_Gv$sqlAuditMgr) Reset() *_Gv$sqlAuditMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$sqlAuditMgr) Get() (result Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$sqlAuditMgr) Gets() (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$sqlAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$sqlAuditMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$sqlAuditMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithRequestID REQUEST_ID获取 
func (obj *_Gv$sqlAuditMgr) WithRequestID(requestID int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST_ID"] = requestID })
}

// WithSQLExecID SQL_EXEC_ID获取 
func (obj *_Gv$sqlAuditMgr) WithSQLExecID(sqlExecID int64) Option {
	return optionFunc(func(o *options) { o.query["SQL_EXEC_ID"] = sqlExecID })
}

// WithTraceID TRACE_ID获取 
func (obj *_Gv$sqlAuditMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["TRACE_ID"] = traceID })
}

// WithSid SID获取 
func (obj *_Gv$sqlAuditMgr) WithSid(sid uint64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithClientIP CLIENT_IP获取 
func (obj *_Gv$sqlAuditMgr) WithClientIP(clientIP string) Option {
	return optionFunc(func(o *options) { o.query["CLIENT_IP"] = clientIP })
}

// WithClientPort CLIENT_PORT获取 
func (obj *_Gv$sqlAuditMgr) WithClientPort(clientPort int64) Option {
	return optionFunc(func(o *options) { o.query["CLIENT_PORT"] = clientPort })
}

// WithTenantID TENANT_ID获取 
func (obj *_Gv$sqlAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTenantName TENANT_NAME获取 
func (obj *_Gv$sqlAuditMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_NAME"] = tenantName })
}

// WithEffectiveTenantID EFFECTIVE_TENANT_ID获取 
func (obj *_Gv$sqlAuditMgr) WithEffectiveTenantID(effectiveTenantID int64) Option {
	return optionFunc(func(o *options) { o.query["EFFECTIVE_TENANT_ID"] = effectiveTenantID })
}

// WithUserID USER_ID获取 
func (obj *_Gv$sqlAuditMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["USER_ID"] = userID })
}

// WithUserName USER_NAME获取 
func (obj *_Gv$sqlAuditMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["USER_NAME"] = userName })
}

// WithUserGroup USER_GROUP获取 
func (obj *_Gv$sqlAuditMgr) WithUserGroup(userGroup int64) Option {
	return optionFunc(func(o *options) { o.query["USER_GROUP"] = userGroup })
}

// WithUserClientIP USER_CLIENT_IP获取 
func (obj *_Gv$sqlAuditMgr) WithUserClientIP(userClientIP string) Option {
	return optionFunc(func(o *options) { o.query["USER_CLIENT_IP"] = userClientIP })
}

// WithDbID DB_ID获取 
func (obj *_Gv$sqlAuditMgr) WithDbID(dbID uint64) Option {
	return optionFunc(func(o *options) { o.query["DB_ID"] = dbID })
}

// WithDbName DB_NAME获取 
func (obj *_Gv$sqlAuditMgr) WithDbName(dbName string) Option {
	return optionFunc(func(o *options) { o.query["DB_NAME"] = dbName })
}

// WithSQLID SQL_ID获取 
func (obj *_Gv$sqlAuditMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["SQL_ID"] = sqlID })
}

// WithQuerySQL QUERY_SQL获取 
func (obj *_Gv$sqlAuditMgr) WithQuerySQL(querySQL string) Option {
	return optionFunc(func(o *options) { o.query["QUERY_SQL"] = querySQL })
}

// WithPlanID PLAN_ID获取 
func (obj *_Gv$sqlAuditMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_ID"] = planID })
}

// WithAffectedRows AFFECTED_ROWS获取 
func (obj *_Gv$sqlAuditMgr) WithAffectedRows(affectedRows int64) Option {
	return optionFunc(func(o *options) { o.query["AFFECTED_ROWS"] = affectedRows })
}

// WithReturnRows RETURN_ROWS获取 
func (obj *_Gv$sqlAuditMgr) WithReturnRows(returnRows int64) Option {
	return optionFunc(func(o *options) { o.query["RETURN_ROWS"] = returnRows })
}

// WithPartitionCnt PARTITION_CNT获取 
func (obj *_Gv$sqlAuditMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_CNT"] = partitionCnt })
}

// WithRetCode RET_CODE获取 
func (obj *_Gv$sqlAuditMgr) WithRetCode(retCode int64) Option {
	return optionFunc(func(o *options) { o.query["RET_CODE"] = retCode })
}

// WithQcID QC_ID获取 
func (obj *_Gv$sqlAuditMgr) WithQcID(qcID uint64) Option {
	return optionFunc(func(o *options) { o.query["QC_ID"] = qcID })
}

// WithDfoID DFO_ID获取 
func (obj *_Gv$sqlAuditMgr) WithDfoID(dfoID int64) Option {
	return optionFunc(func(o *options) { o.query["DFO_ID"] = dfoID })
}

// WithSqcID SQC_ID获取 
func (obj *_Gv$sqlAuditMgr) WithSqcID(sqcID int64) Option {
	return optionFunc(func(o *options) { o.query["SQC_ID"] = sqcID })
}

// WithWorkerID WORKER_ID获取 
func (obj *_Gv$sqlAuditMgr) WithWorkerID(workerID int64) Option {
	return optionFunc(func(o *options) { o.query["WORKER_ID"] = workerID })
}

// WithEvent EVENT获取 
func (obj *_Gv$sqlAuditMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["EVENT"] = event })
}

// WithP1text P1TEXT获取 
func (obj *_Gv$sqlAuditMgr) WithP1text(p1text string) Option {
	return optionFunc(func(o *options) { o.query["P1TEXT"] = p1text })
}

// WithP1 P1获取 
func (obj *_Gv$sqlAuditMgr) WithP1(p1 uint64) Option {
	return optionFunc(func(o *options) { o.query["P1"] = p1 })
}

// WithP2text P2TEXT获取 
func (obj *_Gv$sqlAuditMgr) WithP2text(p2text string) Option {
	return optionFunc(func(o *options) { o.query["P2TEXT"] = p2text })
}

// WithP2 P2获取 
func (obj *_Gv$sqlAuditMgr) WithP2(p2 uint64) Option {
	return optionFunc(func(o *options) { o.query["P2"] = p2 })
}

// WithP3text P3TEXT获取 
func (obj *_Gv$sqlAuditMgr) WithP3text(p3text string) Option {
	return optionFunc(func(o *options) { o.query["P3TEXT"] = p3text })
}

// WithP3 P3获取 
func (obj *_Gv$sqlAuditMgr) WithP3(p3 uint64) Option {
	return optionFunc(func(o *options) { o.query["P3"] = p3 })
}

// WithLevel LEVEL获取 
func (obj *_Gv$sqlAuditMgr) WithLevel(level int64) Option {
	return optionFunc(func(o *options) { o.query["LEVEL"] = level })
}

// WithWaitClassID WAIT_CLASS_ID获取 
func (obj *_Gv$sqlAuditMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS_ID"] = waitClassID })
}

// WithWaitClass# WAIT_CLASS#获取 
func (obj *_Gv$sqlAuditMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS#"] = waitClass# })
}

// WithWaitClass WAIT_CLASS获取 
func (obj *_Gv$sqlAuditMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS"] = waitClass })
}

// WithState STATE获取 
func (obj *_Gv$sqlAuditMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["STATE"] = state })
}

// WithWaitTimeMicro WAIT_TIME_MICRO获取 
func (obj *_Gv$sqlAuditMgr) WithWaitTimeMicro(waitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TIME_MICRO"] = waitTimeMicro })
}

// WithTotalWaitTimeMicro TOTAL_WAIT_TIME_MICRO获取 
func (obj *_Gv$sqlAuditMgr) WithTotalWaitTimeMicro(totalWaitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_WAIT_TIME_MICRO"] = totalWaitTimeMicro })
}

// WithTotalWaits TOTAL_WAITS获取 
func (obj *_Gv$sqlAuditMgr) WithTotalWaits(totalWaits int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_WAITS"] = totalWaits })
}

// WithRPCCount RPC_COUNT获取 
func (obj *_Gv$sqlAuditMgr) WithRPCCount(rpcCount int64) Option {
	return optionFunc(func(o *options) { o.query["RPC_COUNT"] = rpcCount })
}

// WithPlanType PLAN_TYPE获取 
func (obj *_Gv$sqlAuditMgr) WithPlanType(planType int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_TYPE"] = planType })
}

// WithIsInnerSQL IS_INNER_SQL获取 
func (obj *_Gv$sqlAuditMgr) WithIsInnerSQL(isInnerSQL int8) Option {
	return optionFunc(func(o *options) { o.query["IS_INNER_SQL"] = isInnerSQL })
}

// WithIsExecutorRPC IS_EXECUTOR_RPC获取 
func (obj *_Gv$sqlAuditMgr) WithIsExecutorRPC(isExecutorRPC int8) Option {
	return optionFunc(func(o *options) { o.query["IS_EXECUTOR_RPC"] = isExecutorRPC })
}

// WithIsHitPlan IS_HIT_PLAN获取 
func (obj *_Gv$sqlAuditMgr) WithIsHitPlan(isHitPlan int8) Option {
	return optionFunc(func(o *options) { o.query["IS_HIT_PLAN"] = isHitPlan })
}

// WithRequestTime REQUEST_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithRequestTime(requestTime int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST_TIME"] = requestTime })
}

// WithElapsedTime ELAPSED_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithElapsedTime(elapsedTime int64) Option {
	return optionFunc(func(o *options) { o.query["ELAPSED_TIME"] = elapsedTime })
}

// WithNetTime NET_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithNetTime(netTime int64) Option {
	return optionFunc(func(o *options) { o.query["NET_TIME"] = netTime })
}

// WithNetWaitTime NET_WAIT_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithNetWaitTime(netWaitTime int64) Option {
	return optionFunc(func(o *options) { o.query["NET_WAIT_TIME"] = netWaitTime })
}

// WithQueueTime QUEUE_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithQueueTime(queueTime int64) Option {
	return optionFunc(func(o *options) { o.query["QUEUE_TIME"] = queueTime })
}

// WithDecodeTime DECODE_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithDecodeTime(decodeTime int64) Option {
	return optionFunc(func(o *options) { o.query["DECODE_TIME"] = decodeTime })
}

// WithGetPlanTime GET_PLAN_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithGetPlanTime(getPlanTime int64) Option {
	return optionFunc(func(o *options) { o.query["GET_PLAN_TIME"] = getPlanTime })
}

// WithExecuteTime EXECUTE_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithExecuteTime(executeTime int64) Option {
	return optionFunc(func(o *options) { o.query["EXECUTE_TIME"] = executeTime })
}

// WithApplicationWaitTime APPLICATION_WAIT_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithApplicationWaitTime(applicationWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["APPLICATION_WAIT_TIME"] = applicationWaitTime })
}

// WithConcurrencyWaitTime CONCURRENCY_WAIT_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithConcurrencyWaitTime(concurrencyWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["CONCURRENCY_WAIT_TIME"] = concurrencyWaitTime })
}

// WithUserIoWaitTime USER_IO_WAIT_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithUserIoWaitTime(userIoWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["USER_IO_WAIT_TIME"] = userIoWaitTime })
}

// WithScheduleTime SCHEDULE_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithScheduleTime(scheduleTime uint64) Option {
	return optionFunc(func(o *options) { o.query["SCHEDULE_TIME"] = scheduleTime })
}

// WithRowCacheHit ROW_CACHE_HIT获取 
func (obj *_Gv$sqlAuditMgr) WithRowCacheHit(rowCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["ROW_CACHE_HIT"] = rowCacheHit })
}

// WithBloomFilterCacheHit BLOOM_FILTER_CACHE_HIT获取 
func (obj *_Gv$sqlAuditMgr) WithBloomFilterCacheHit(bloomFilterCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["BLOOM_FILTER_CACHE_HIT"] = bloomFilterCacheHit })
}

// WithBlockCacheHit BLOCK_CACHE_HIT获取 
func (obj *_Gv$sqlAuditMgr) WithBlockCacheHit(blockCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["BLOCK_CACHE_HIT"] = blockCacheHit })
}

// WithBlockIndexCacheHit BLOCK_INDEX_CACHE_HIT获取 
func (obj *_Gv$sqlAuditMgr) WithBlockIndexCacheHit(blockIndexCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["BLOCK_INDEX_CACHE_HIT"] = blockIndexCacheHit })
}

// WithDiskReads DISK_READS获取 
func (obj *_Gv$sqlAuditMgr) WithDiskReads(diskReads int64) Option {
	return optionFunc(func(o *options) { o.query["DISK_READS"] = diskReads })
}

// WithRetryCnt RETRY_CNT获取 
func (obj *_Gv$sqlAuditMgr) WithRetryCnt(retryCnt int64) Option {
	return optionFunc(func(o *options) { o.query["RETRY_CNT"] = retryCnt })
}

// WithTableScan TABLE_SCAN获取 
func (obj *_Gv$sqlAuditMgr) WithTableScan(tableScan int8) Option {
	return optionFunc(func(o *options) { o.query["TABLE_SCAN"] = tableScan })
}

// WithConsistencyLevel CONSISTENCY_LEVEL获取 
func (obj *_Gv$sqlAuditMgr) WithConsistencyLevel(consistencyLevel int64) Option {
	return optionFunc(func(o *options) { o.query["CONSISTENCY_LEVEL"] = consistencyLevel })
}

// WithMemstoreReadRowCount MEMSTORE_READ_ROW_COUNT获取 
func (obj *_Gv$sqlAuditMgr) WithMemstoreReadRowCount(memstoreReadRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["MEMSTORE_READ_ROW_COUNT"] = memstoreReadRowCount })
}

// WithSsstoreReadRowCount SSSTORE_READ_ROW_COUNT获取 
func (obj *_Gv$sqlAuditMgr) WithSsstoreReadRowCount(ssstoreReadRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["SSSTORE_READ_ROW_COUNT"] = ssstoreReadRowCount })
}

// WithRequestMemoryUsed REQUEST_MEMORY_USED获取 
func (obj *_Gv$sqlAuditMgr) WithRequestMemoryUsed(requestMemoryUsed int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST_MEMORY_USED"] = requestMemoryUsed })
}

// WithExpectedWorkerCount EXPECTED_WORKER_COUNT获取 
func (obj *_Gv$sqlAuditMgr) WithExpectedWorkerCount(expectedWorkerCount int64) Option {
	return optionFunc(func(o *options) { o.query["EXPECTED_WORKER_COUNT"] = expectedWorkerCount })
}

// WithUsedWorkerCount USED_WORKER_COUNT获取 
func (obj *_Gv$sqlAuditMgr) WithUsedWorkerCount(usedWorkerCount int64) Option {
	return optionFunc(func(o *options) { o.query["USED_WORKER_COUNT"] = usedWorkerCount })
}

// WithSchedInfo SCHED_INFO获取 
func (obj *_Gv$sqlAuditMgr) WithSchedInfo(schedInfo string) Option {
	return optionFunc(func(o *options) { o.query["SCHED_INFO"] = schedInfo })
}

// WithFuseRowCacheHit FUSE_ROW_CACHE_HIT获取 
func (obj *_Gv$sqlAuditMgr) WithFuseRowCacheHit(fuseRowCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["FUSE_ROW_CACHE_HIT"] = fuseRowCacheHit })
}

// WithPsStmtID PS_STMT_ID获取 
func (obj *_Gv$sqlAuditMgr) WithPsStmtID(psStmtID int64) Option {
	return optionFunc(func(o *options) { o.query["PS_STMT_ID"] = psStmtID })
}

// WithTransactionHash TRANSACTION_HASH获取 
func (obj *_Gv$sqlAuditMgr) WithTransactionHash(transactionHash uint64) Option {
	return optionFunc(func(o *options) { o.query["TRANSACTION_HASH"] = transactionHash })
}

// WithRequestType REQUEST_TYPE获取 
func (obj *_Gv$sqlAuditMgr) WithRequestType(requestType int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST_TYPE"] = requestType })
}

// WithIsBatchedMultiStmt IS_BATCHED_MULTI_STMT获取 
func (obj *_Gv$sqlAuditMgr) WithIsBatchedMultiStmt(isBatchedMultiStmt int8) Option {
	return optionFunc(func(o *options) { o.query["IS_BATCHED_MULTI_STMT"] = isBatchedMultiStmt })
}

// WithObTraceInfo OB_TRACE_INFO获取 
func (obj *_Gv$sqlAuditMgr) WithObTraceInfo(obTraceInfo string) Option {
	return optionFunc(func(o *options) { o.query["OB_TRACE_INFO"] = obTraceInfo })
}

// WithPlanHash PLAN_HASH获取 
func (obj *_Gv$sqlAuditMgr) WithPlanHash(planHash uint64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_HASH"] = planHash })
}

// WithLockForReadTime LOCK_FOR_READ_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithLockForReadTime(lockForReadTime int64) Option {
	return optionFunc(func(o *options) { o.query["LOCK_FOR_READ_TIME"] = lockForReadTime })
}

// WithWaitTrxMigrateTime WAIT_TRX_MIGRATE_TIME获取 
func (obj *_Gv$sqlAuditMgr) WithWaitTrxMigrateTime(waitTrxMigrateTime int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TRX_MIGRATE_TIME"] = waitTrxMigrateTime })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$sqlAuditMgr) GetByOption(opts ...Option) (result Gv$sqlAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$sqlAuditMgr) GetByOptions(opts ...Option) (results []*Gv$sqlAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$sqlAuditMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$sqlAudit,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where(options.query)
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


// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSvrIP(svrIP string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSvrPort(svrPort int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromRequestID 通过REQUEST_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRequestID(requestID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_ID` = ?", requestID).Find(&results).Error
	
	return
}

// GetBatchFromRequestID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRequestID(requestIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_ID` IN (?)", requestIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLExecID 通过SQL_EXEC_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSQLExecID(sqlExecID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SQL_EXEC_ID` = ?", sqlExecID).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSQLExecID(sqlExecIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SQL_EXEC_ID` IN (?)", sqlExecIDs).Find(&results).Error
	
	return
}
 
// GetFromTraceID 通过TRACE_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromTraceID(traceID string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TRACE_ID` = ?", traceID).Find(&results).Error
	
	return
}

// GetBatchFromTraceID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromTraceID(traceIDs []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TRACE_ID` IN (?)", traceIDs).Find(&results).Error
	
	return
}
 
// GetFromSid 通过SID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSid(sid uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSid(sids []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromClientIP 通过CLIENT_IP获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromClientIP(clientIP string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CLIENT_IP` = ?", clientIP).Find(&results).Error
	
	return
}

// GetBatchFromClientIP 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromClientIP(clientIPs []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CLIENT_IP` IN (?)", clientIPs).Find(&results).Error
	
	return
}
 
// GetFromClientPort 通过CLIENT_PORT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromClientPort(clientPort int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CLIENT_PORT` = ?", clientPort).Find(&results).Error
	
	return
}

// GetBatchFromClientPort 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromClientPort(clientPorts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CLIENT_PORT` IN (?)", clientPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromTenantID(tenantID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantName 通过TENANT_NAME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromTenantName(tenantName string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TENANT_NAME` = ?", tenantName).Find(&results).Error
	
	return
}

// GetBatchFromTenantName 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromTenantName(tenantNames []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TENANT_NAME` IN (?)", tenantNames).Find(&results).Error
	
	return
}
 
// GetFromEffectiveTenantID 通过EFFECTIVE_TENANT_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromEffectiveTenantID(effectiveTenantID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EFFECTIVE_TENANT_ID` = ?", effectiveTenantID).Find(&results).Error
	
	return
}

// GetBatchFromEffectiveTenantID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromEffectiveTenantID(effectiveTenantIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EFFECTIVE_TENANT_ID` IN (?)", effectiveTenantIDs).Find(&results).Error
	
	return
}
 
// GetFromUserID 通过USER_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromUserID(userID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_ID` = ?", userID).Find(&results).Error
	
	return
}

// GetBatchFromUserID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromUserID(userIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_ID` IN (?)", userIDs).Find(&results).Error
	
	return
}
 
// GetFromUserName 通过USER_NAME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromUserName(userName string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_NAME` = ?", userName).Find(&results).Error
	
	return
}

// GetBatchFromUserName 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromUserName(userNames []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_NAME` IN (?)", userNames).Find(&results).Error
	
	return
}
 
// GetFromUserGroup 通过USER_GROUP获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromUserGroup(userGroup int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_GROUP` = ?", userGroup).Find(&results).Error
	
	return
}

// GetBatchFromUserGroup 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromUserGroup(userGroups []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_GROUP` IN (?)", userGroups).Find(&results).Error
	
	return
}
 
// GetFromUserClientIP 通过USER_CLIENT_IP获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromUserClientIP(userClientIP string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_CLIENT_IP` = ?", userClientIP).Find(&results).Error
	
	return
}

// GetBatchFromUserClientIP 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromUserClientIP(userClientIPs []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_CLIENT_IP` IN (?)", userClientIPs).Find(&results).Error
	
	return
}
 
// GetFromDbID 通过DB_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromDbID(dbID uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DB_ID` = ?", dbID).Find(&results).Error
	
	return
}

// GetBatchFromDbID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromDbID(dbIDs []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DB_ID` IN (?)", dbIDs).Find(&results).Error
	
	return
}
 
// GetFromDbName 通过DB_NAME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromDbName(dbName string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DB_NAME` = ?", dbName).Find(&results).Error
	
	return
}

// GetBatchFromDbName 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromDbName(dbNames []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DB_NAME` IN (?)", dbNames).Find(&results).Error
	
	return
}
 
// GetFromSQLID 通过SQL_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSQLID(sqlID string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SQL_ID` = ?", sqlID).Find(&results).Error
	
	return
}

// GetBatchFromSQLID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSQLID(sqlIDs []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SQL_ID` IN (?)", sqlIDs).Find(&results).Error
	
	return
}
 
// GetFromQuerySQL 通过QUERY_SQL获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromQuerySQL(querySQL string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`QUERY_SQL` = ?", querySQL).Find(&results).Error
	
	return
}

// GetBatchFromQuerySQL 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromQuerySQL(querySQLs []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`QUERY_SQL` IN (?)", querySQLs).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过PLAN_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromPlanID(planID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PLAN_ID` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromPlanID(planIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PLAN_ID` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromAffectedRows 通过AFFECTED_ROWS获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromAffectedRows(affectedRows int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`AFFECTED_ROWS` = ?", affectedRows).Find(&results).Error
	
	return
}

// GetBatchFromAffectedRows 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromAffectedRows(affectedRowss []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`AFFECTED_ROWS` IN (?)", affectedRowss).Find(&results).Error
	
	return
}
 
// GetFromReturnRows 通过RETURN_ROWS获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromReturnRows(returnRows int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RETURN_ROWS` = ?", returnRows).Find(&results).Error
	
	return
}

// GetBatchFromReturnRows 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromReturnRows(returnRowss []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RETURN_ROWS` IN (?)", returnRowss).Find(&results).Error
	
	return
}
 
// GetFromPartitionCnt 通过PARTITION_CNT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromPartitionCnt(partitionCnt int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PARTITION_CNT` = ?", partitionCnt).Find(&results).Error
	
	return
}

// GetBatchFromPartitionCnt 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PARTITION_CNT` IN (?)", partitionCnts).Find(&results).Error
	
	return
}
 
// GetFromRetCode 通过RET_CODE获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRetCode(retCode int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RET_CODE` = ?", retCode).Find(&results).Error
	
	return
}

// GetBatchFromRetCode 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRetCode(retCodes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RET_CODE` IN (?)", retCodes).Find(&results).Error
	
	return
}
 
// GetFromQcID 通过QC_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromQcID(qcID uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`QC_ID` = ?", qcID).Find(&results).Error
	
	return
}

// GetBatchFromQcID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromQcID(qcIDs []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`QC_ID` IN (?)", qcIDs).Find(&results).Error
	
	return
}
 
// GetFromDfoID 通过DFO_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromDfoID(dfoID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DFO_ID` = ?", dfoID).Find(&results).Error
	
	return
}

// GetBatchFromDfoID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromDfoID(dfoIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DFO_ID` IN (?)", dfoIDs).Find(&results).Error
	
	return
}
 
// GetFromSqcID 通过SQC_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSqcID(sqcID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SQC_ID` = ?", sqcID).Find(&results).Error
	
	return
}

// GetBatchFromSqcID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSqcID(sqcIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SQC_ID` IN (?)", sqcIDs).Find(&results).Error
	
	return
}
 
// GetFromWorkerID 通过WORKER_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromWorkerID(workerID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WORKER_ID` = ?", workerID).Find(&results).Error
	
	return
}

// GetBatchFromWorkerID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromWorkerID(workerIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WORKER_ID` IN (?)", workerIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过EVENT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromEvent(event string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EVENT` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromEvent(events []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EVENT` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromP1text 通过P1TEXT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromP1text(p1text string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P1TEXT` = ?", p1text).Find(&results).Error
	
	return
}

// GetBatchFromP1text 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromP1text(p1texts []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P1TEXT` IN (?)", p1texts).Find(&results).Error
	
	return
}
 
// GetFromP1 通过P1获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromP1(p1 uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P1` = ?", p1).Find(&results).Error
	
	return
}

// GetBatchFromP1 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromP1(p1s []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P1` IN (?)", p1s).Find(&results).Error
	
	return
}
 
// GetFromP2text 通过P2TEXT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromP2text(p2text string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P2TEXT` = ?", p2text).Find(&results).Error
	
	return
}

// GetBatchFromP2text 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromP2text(p2texts []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P2TEXT` IN (?)", p2texts).Find(&results).Error
	
	return
}
 
// GetFromP2 通过P2获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromP2(p2 uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P2` = ?", p2).Find(&results).Error
	
	return
}

// GetBatchFromP2 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromP2(p2s []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P2` IN (?)", p2s).Find(&results).Error
	
	return
}
 
// GetFromP3text 通过P3TEXT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromP3text(p3text string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P3TEXT` = ?", p3text).Find(&results).Error
	
	return
}

// GetBatchFromP3text 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromP3text(p3texts []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P3TEXT` IN (?)", p3texts).Find(&results).Error
	
	return
}
 
// GetFromP3 通过P3获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromP3(p3 uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P3` = ?", p3).Find(&results).Error
	
	return
}

// GetBatchFromP3 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromP3(p3s []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`P3` IN (?)", p3s).Find(&results).Error
	
	return
}
 
// GetFromLevel 通过LEVEL获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromLevel(level int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`LEVEL` = ?", level).Find(&results).Error
	
	return
}

// GetBatchFromLevel 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromLevel(levels []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`LEVEL` IN (?)", levels).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过WAIT_CLASS_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromWaitClassID(waitClassID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_CLASS_ID` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_CLASS_ID` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过WAIT_CLASS#获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromWaitClass#(waitClass# int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_CLASS#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_CLASS#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过WAIT_CLASS获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromWaitClass(waitClass string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_CLASS` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromWaitClass(waitClasss []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_CLASS` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromState 通过STATE获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromState(state string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`STATE` = ?", state).Find(&results).Error
	
	return
}

// GetBatchFromState 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromState(states []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`STATE` IN (?)", states).Find(&results).Error
	
	return
}
 
// GetFromWaitTimeMicro 通过WAIT_TIME_MICRO获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromWaitTimeMicro(waitTimeMicro int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_TIME_MICRO` = ?", waitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromWaitTimeMicro 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromWaitTimeMicro(waitTimeMicros []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_TIME_MICRO` IN (?)", waitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTotalWaitTimeMicro 通过TOTAL_WAIT_TIME_MICRO获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromTotalWaitTimeMicro(totalWaitTimeMicro int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TOTAL_WAIT_TIME_MICRO` = ?", totalWaitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaitTimeMicro 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromTotalWaitTimeMicro(totalWaitTimeMicros []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TOTAL_WAIT_TIME_MICRO` IN (?)", totalWaitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTotalWaits 通过TOTAL_WAITS获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromTotalWaits(totalWaits int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TOTAL_WAITS` = ?", totalWaits).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaits 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromTotalWaits(totalWaitss []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TOTAL_WAITS` IN (?)", totalWaitss).Find(&results).Error
	
	return
}
 
// GetFromRPCCount 通过RPC_COUNT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRPCCount(rpcCount int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RPC_COUNT` = ?", rpcCount).Find(&results).Error
	
	return
}

// GetBatchFromRPCCount 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRPCCount(rpcCounts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RPC_COUNT` IN (?)", rpcCounts).Find(&results).Error
	
	return
}
 
// GetFromPlanType 通过PLAN_TYPE获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromPlanType(planType int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PLAN_TYPE` = ?", planType).Find(&results).Error
	
	return
}

// GetBatchFromPlanType 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromPlanType(planTypes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PLAN_TYPE` IN (?)", planTypes).Find(&results).Error
	
	return
}
 
// GetFromIsInnerSQL 通过IS_INNER_SQL获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromIsInnerSQL(isInnerSQL int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_INNER_SQL` = ?", isInnerSQL).Find(&results).Error
	
	return
}

// GetBatchFromIsInnerSQL 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromIsInnerSQL(isInnerSQLs []int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_INNER_SQL` IN (?)", isInnerSQLs).Find(&results).Error
	
	return
}
 
// GetFromIsExecutorRPC 通过IS_EXECUTOR_RPC获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromIsExecutorRPC(isExecutorRPC int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_EXECUTOR_RPC` = ?", isExecutorRPC).Find(&results).Error
	
	return
}

// GetBatchFromIsExecutorRPC 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromIsExecutorRPC(isExecutorRPCs []int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_EXECUTOR_RPC` IN (?)", isExecutorRPCs).Find(&results).Error
	
	return
}
 
// GetFromIsHitPlan 通过IS_HIT_PLAN获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromIsHitPlan(isHitPlan int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_HIT_PLAN` = ?", isHitPlan).Find(&results).Error
	
	return
}

// GetBatchFromIsHitPlan 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromIsHitPlan(isHitPlans []int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_HIT_PLAN` IN (?)", isHitPlans).Find(&results).Error
	
	return
}
 
// GetFromRequestTime 通过REQUEST_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRequestTime(requestTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_TIME` = ?", requestTime).Find(&results).Error
	
	return
}

// GetBatchFromRequestTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRequestTime(requestTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_TIME` IN (?)", requestTimes).Find(&results).Error
	
	return
}
 
// GetFromElapsedTime 通过ELAPSED_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromElapsedTime(elapsedTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`ELAPSED_TIME` = ?", elapsedTime).Find(&results).Error
	
	return
}

// GetBatchFromElapsedTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromElapsedTime(elapsedTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`ELAPSED_TIME` IN (?)", elapsedTimes).Find(&results).Error
	
	return
}
 
// GetFromNetTime 通过NET_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromNetTime(netTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`NET_TIME` = ?", netTime).Find(&results).Error
	
	return
}

// GetBatchFromNetTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromNetTime(netTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`NET_TIME` IN (?)", netTimes).Find(&results).Error
	
	return
}
 
// GetFromNetWaitTime 通过NET_WAIT_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromNetWaitTime(netWaitTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`NET_WAIT_TIME` = ?", netWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromNetWaitTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromNetWaitTime(netWaitTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`NET_WAIT_TIME` IN (?)", netWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromQueueTime 通过QUEUE_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromQueueTime(queueTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`QUEUE_TIME` = ?", queueTime).Find(&results).Error
	
	return
}

// GetBatchFromQueueTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromQueueTime(queueTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`QUEUE_TIME` IN (?)", queueTimes).Find(&results).Error
	
	return
}
 
// GetFromDecodeTime 通过DECODE_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromDecodeTime(decodeTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DECODE_TIME` = ?", decodeTime).Find(&results).Error
	
	return
}

// GetBatchFromDecodeTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromDecodeTime(decodeTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DECODE_TIME` IN (?)", decodeTimes).Find(&results).Error
	
	return
}
 
// GetFromGetPlanTime 通过GET_PLAN_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromGetPlanTime(getPlanTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`GET_PLAN_TIME` = ?", getPlanTime).Find(&results).Error
	
	return
}

// GetBatchFromGetPlanTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromGetPlanTime(getPlanTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`GET_PLAN_TIME` IN (?)", getPlanTimes).Find(&results).Error
	
	return
}
 
// GetFromExecuteTime 通过EXECUTE_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromExecuteTime(executeTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EXECUTE_TIME` = ?", executeTime).Find(&results).Error
	
	return
}

// GetBatchFromExecuteTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromExecuteTime(executeTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EXECUTE_TIME` IN (?)", executeTimes).Find(&results).Error
	
	return
}
 
// GetFromApplicationWaitTime 通过APPLICATION_WAIT_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromApplicationWaitTime(applicationWaitTime uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`APPLICATION_WAIT_TIME` = ?", applicationWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromApplicationWaitTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromApplicationWaitTime(applicationWaitTimes []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`APPLICATION_WAIT_TIME` IN (?)", applicationWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromConcurrencyWaitTime 通过CONCURRENCY_WAIT_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromConcurrencyWaitTime(concurrencyWaitTime uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CONCURRENCY_WAIT_TIME` = ?", concurrencyWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromConcurrencyWaitTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromConcurrencyWaitTime(concurrencyWaitTimes []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CONCURRENCY_WAIT_TIME` IN (?)", concurrencyWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromUserIoWaitTime 通过USER_IO_WAIT_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromUserIoWaitTime(userIoWaitTime uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_IO_WAIT_TIME` = ?", userIoWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromUserIoWaitTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromUserIoWaitTime(userIoWaitTimes []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USER_IO_WAIT_TIME` IN (?)", userIoWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromScheduleTime 通过SCHEDULE_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromScheduleTime(scheduleTime uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SCHEDULE_TIME` = ?", scheduleTime).Find(&results).Error
	
	return
}

// GetBatchFromScheduleTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromScheduleTime(scheduleTimes []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SCHEDULE_TIME` IN (?)", scheduleTimes).Find(&results).Error
	
	return
}
 
// GetFromRowCacheHit 通过ROW_CACHE_HIT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRowCacheHit(rowCacheHit int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`ROW_CACHE_HIT` = ?", rowCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromRowCacheHit 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRowCacheHit(rowCacheHits []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`ROW_CACHE_HIT` IN (?)", rowCacheHits).Find(&results).Error
	
	return
}
 
// GetFromBloomFilterCacheHit 通过BLOOM_FILTER_CACHE_HIT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromBloomFilterCacheHit(bloomFilterCacheHit int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`BLOOM_FILTER_CACHE_HIT` = ?", bloomFilterCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromBloomFilterCacheHit 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromBloomFilterCacheHit(bloomFilterCacheHits []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`BLOOM_FILTER_CACHE_HIT` IN (?)", bloomFilterCacheHits).Find(&results).Error
	
	return
}
 
// GetFromBlockCacheHit 通过BLOCK_CACHE_HIT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromBlockCacheHit(blockCacheHit int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`BLOCK_CACHE_HIT` = ?", blockCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromBlockCacheHit 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromBlockCacheHit(blockCacheHits []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`BLOCK_CACHE_HIT` IN (?)", blockCacheHits).Find(&results).Error
	
	return
}
 
// GetFromBlockIndexCacheHit 通过BLOCK_INDEX_CACHE_HIT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromBlockIndexCacheHit(blockIndexCacheHit int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`BLOCK_INDEX_CACHE_HIT` = ?", blockIndexCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromBlockIndexCacheHit 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromBlockIndexCacheHit(blockIndexCacheHits []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`BLOCK_INDEX_CACHE_HIT` IN (?)", blockIndexCacheHits).Find(&results).Error
	
	return
}
 
// GetFromDiskReads 通过DISK_READS获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromDiskReads(diskReads int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DISK_READS` = ?", diskReads).Find(&results).Error
	
	return
}

// GetBatchFromDiskReads 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromDiskReads(diskReadss []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`DISK_READS` IN (?)", diskReadss).Find(&results).Error
	
	return
}
 
// GetFromRetryCnt 通过RETRY_CNT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRetryCnt(retryCnt int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RETRY_CNT` = ?", retryCnt).Find(&results).Error
	
	return
}

// GetBatchFromRetryCnt 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRetryCnt(retryCnts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`RETRY_CNT` IN (?)", retryCnts).Find(&results).Error
	
	return
}
 
// GetFromTableScan 通过TABLE_SCAN获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromTableScan(tableScan int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TABLE_SCAN` = ?", tableScan).Find(&results).Error
	
	return
}

// GetBatchFromTableScan 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromTableScan(tableScans []int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TABLE_SCAN` IN (?)", tableScans).Find(&results).Error
	
	return
}
 
// GetFromConsistencyLevel 通过CONSISTENCY_LEVEL获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromConsistencyLevel(consistencyLevel int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CONSISTENCY_LEVEL` = ?", consistencyLevel).Find(&results).Error
	
	return
}

// GetBatchFromConsistencyLevel 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromConsistencyLevel(consistencyLevels []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`CONSISTENCY_LEVEL` IN (?)", consistencyLevels).Find(&results).Error
	
	return
}
 
// GetFromMemstoreReadRowCount 通过MEMSTORE_READ_ROW_COUNT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromMemstoreReadRowCount(memstoreReadRowCount int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`MEMSTORE_READ_ROW_COUNT` = ?", memstoreReadRowCount).Find(&results).Error
	
	return
}

// GetBatchFromMemstoreReadRowCount 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromMemstoreReadRowCount(memstoreReadRowCounts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`MEMSTORE_READ_ROW_COUNT` IN (?)", memstoreReadRowCounts).Find(&results).Error
	
	return
}
 
// GetFromSsstoreReadRowCount 通过SSSTORE_READ_ROW_COUNT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSsstoreReadRowCount(ssstoreReadRowCount int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SSSTORE_READ_ROW_COUNT` = ?", ssstoreReadRowCount).Find(&results).Error
	
	return
}

// GetBatchFromSsstoreReadRowCount 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSsstoreReadRowCount(ssstoreReadRowCounts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SSSTORE_READ_ROW_COUNT` IN (?)", ssstoreReadRowCounts).Find(&results).Error
	
	return
}
 
// GetFromRequestMemoryUsed 通过REQUEST_MEMORY_USED获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRequestMemoryUsed(requestMemoryUsed int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_MEMORY_USED` = ?", requestMemoryUsed).Find(&results).Error
	
	return
}

// GetBatchFromRequestMemoryUsed 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRequestMemoryUsed(requestMemoryUseds []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_MEMORY_USED` IN (?)", requestMemoryUseds).Find(&results).Error
	
	return
}
 
// GetFromExpectedWorkerCount 通过EXPECTED_WORKER_COUNT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromExpectedWorkerCount(expectedWorkerCount int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EXPECTED_WORKER_COUNT` = ?", expectedWorkerCount).Find(&results).Error
	
	return
}

// GetBatchFromExpectedWorkerCount 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromExpectedWorkerCount(expectedWorkerCounts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`EXPECTED_WORKER_COUNT` IN (?)", expectedWorkerCounts).Find(&results).Error
	
	return
}
 
// GetFromUsedWorkerCount 通过USED_WORKER_COUNT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromUsedWorkerCount(usedWorkerCount int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USED_WORKER_COUNT` = ?", usedWorkerCount).Find(&results).Error
	
	return
}

// GetBatchFromUsedWorkerCount 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromUsedWorkerCount(usedWorkerCounts []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`USED_WORKER_COUNT` IN (?)", usedWorkerCounts).Find(&results).Error
	
	return
}
 
// GetFromSchedInfo 通过SCHED_INFO获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromSchedInfo(schedInfo string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SCHED_INFO` = ?", schedInfo).Find(&results).Error
	
	return
}

// GetBatchFromSchedInfo 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromSchedInfo(schedInfos []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`SCHED_INFO` IN (?)", schedInfos).Find(&results).Error
	
	return
}
 
// GetFromFuseRowCacheHit 通过FUSE_ROW_CACHE_HIT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromFuseRowCacheHit(fuseRowCacheHit int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`FUSE_ROW_CACHE_HIT` = ?", fuseRowCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromFuseRowCacheHit 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromFuseRowCacheHit(fuseRowCacheHits []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`FUSE_ROW_CACHE_HIT` IN (?)", fuseRowCacheHits).Find(&results).Error
	
	return
}
 
// GetFromPsStmtID 通过PS_STMT_ID获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromPsStmtID(psStmtID int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PS_STMT_ID` = ?", psStmtID).Find(&results).Error
	
	return
}

// GetBatchFromPsStmtID 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromPsStmtID(psStmtIDs []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PS_STMT_ID` IN (?)", psStmtIDs).Find(&results).Error
	
	return
}
 
// GetFromTransactionHash 通过TRANSACTION_HASH获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromTransactionHash(transactionHash uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TRANSACTION_HASH` = ?", transactionHash).Find(&results).Error
	
	return
}

// GetBatchFromTransactionHash 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromTransactionHash(transactionHashs []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`TRANSACTION_HASH` IN (?)", transactionHashs).Find(&results).Error
	
	return
}
 
// GetFromRequestType 通过REQUEST_TYPE获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromRequestType(requestType int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_TYPE` = ?", requestType).Find(&results).Error
	
	return
}

// GetBatchFromRequestType 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromRequestType(requestTypes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`REQUEST_TYPE` IN (?)", requestTypes).Find(&results).Error
	
	return
}
 
// GetFromIsBatchedMultiStmt 通过IS_BATCHED_MULTI_STMT获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromIsBatchedMultiStmt(isBatchedMultiStmt int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_BATCHED_MULTI_STMT` = ?", isBatchedMultiStmt).Find(&results).Error
	
	return
}

// GetBatchFromIsBatchedMultiStmt 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromIsBatchedMultiStmt(isBatchedMultiStmts []int8) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`IS_BATCHED_MULTI_STMT` IN (?)", isBatchedMultiStmts).Find(&results).Error
	
	return
}
 
// GetFromObTraceInfo 通过OB_TRACE_INFO获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromObTraceInfo(obTraceInfo string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`OB_TRACE_INFO` = ?", obTraceInfo).Find(&results).Error
	
	return
}

// GetBatchFromObTraceInfo 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromObTraceInfo(obTraceInfos []string) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`OB_TRACE_INFO` IN (?)", obTraceInfos).Find(&results).Error
	
	return
}
 
// GetFromPlanHash 通过PLAN_HASH获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromPlanHash(planHash uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PLAN_HASH` = ?", planHash).Find(&results).Error
	
	return
}

// GetBatchFromPlanHash 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromPlanHash(planHashs []uint64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`PLAN_HASH` IN (?)", planHashs).Find(&results).Error
	
	return
}
 
// GetFromLockForReadTime 通过LOCK_FOR_READ_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromLockForReadTime(lockForReadTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`LOCK_FOR_READ_TIME` = ?", lockForReadTime).Find(&results).Error
	
	return
}

// GetBatchFromLockForReadTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromLockForReadTime(lockForReadTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`LOCK_FOR_READ_TIME` IN (?)", lockForReadTimes).Find(&results).Error
	
	return
}
 
// GetFromWaitTrxMigrateTime 通过WAIT_TRX_MIGRATE_TIME获取内容  
func (obj *_Gv$sqlAuditMgr) GetFromWaitTrxMigrateTime(waitTrxMigrateTime int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_TRX_MIGRATE_TIME` = ?", waitTrxMigrateTime).Find(&results).Error
	
	return
}

// GetBatchFromWaitTrxMigrateTime 批量查找 
func (obj *_Gv$sqlAuditMgr) GetBatchFromWaitTrxMigrateTime(waitTrxMigrateTimes []int64) (results []*Gv$sqlAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlAudit{}).Where("`WAIT_TRX_MIGRATE_TIME` IN (?)", waitTrxMigrateTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

