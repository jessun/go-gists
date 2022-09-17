package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualSQLAuditMgr struct {
	*_BaseMgr
}

// AllVirtualSQLAuditMgr open func
func AllVirtualSQLAuditMgr(db *gorm.DB) *_AllVirtualSQLAuditMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_audit"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLAuditMgr) GetTableName() string {
	return "__all_virtual_sql_audit"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLAuditMgr) Reset() *_AllVirtualSQLAuditMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualSQLAuditMgr) Get() (result AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLAuditMgr) Gets() (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取 
func (obj *_AllVirtualSQLAuditMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllVirtualSQLAuditMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRequestID request_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithRequestID(requestID int64) Option {
	return optionFunc(func(o *options) { o.query["request_id"] = requestID })
}

// WithTraceID trace_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// WithClientIP client_ip获取 
func (obj *_AllVirtualSQLAuditMgr) WithClientIP(clientIP string) Option {
	return optionFunc(func(o *options) { o.query["client_ip"] = clientIP })
}

// WithClientPort client_port获取 
func (obj *_AllVirtualSQLAuditMgr) WithClientPort(clientPort int64) Option {
	return optionFunc(func(o *options) { o.query["client_port"] = clientPort })
}

// WithTenantName tenant_name获取 
func (obj *_AllVirtualSQLAuditMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithEffectiveTenantID effective_tenant_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithEffectiveTenantID(effectiveTenantID int64) Option {
	return optionFunc(func(o *options) { o.query["effective_tenant_id"] = effectiveTenantID })
}

// WithUserID user_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserName user_name获取 
func (obj *_AllVirtualSQLAuditMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithDbID db_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithDbID(dbID uint64) Option {
	return optionFunc(func(o *options) { o.query["db_id"] = dbID })
}

// WithDbName db_name获取 
func (obj *_AllVirtualSQLAuditMgr) WithDbName(dbName string) Option {
	return optionFunc(func(o *options) { o.query["db_name"] = dbName })
}

// WithSQLID sql_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithQuerySQL query_sql获取 
func (obj *_AllVirtualSQLAuditMgr) WithQuerySQL(querySQL string) Option {
	return optionFunc(func(o *options) { o.query["query_sql"] = querySQL })
}

// WithPlanID plan_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithAffectedRows affected_rows获取 
func (obj *_AllVirtualSQLAuditMgr) WithAffectedRows(affectedRows int64) Option {
	return optionFunc(func(o *options) { o.query["affected_rows"] = affectedRows })
}

// WithReturnRows return_rows获取 
func (obj *_AllVirtualSQLAuditMgr) WithReturnRows(returnRows int64) Option {
	return optionFunc(func(o *options) { o.query["return_rows"] = returnRows })
}

// WithPartitionCnt partition_cnt获取 
func (obj *_AllVirtualSQLAuditMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithRetCode ret_code获取 
func (obj *_AllVirtualSQLAuditMgr) WithRetCode(retCode int64) Option {
	return optionFunc(func(o *options) { o.query["ret_code"] = retCode })
}

// WithQcID qc_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithQcID(qcID uint64) Option {
	return optionFunc(func(o *options) { o.query["qc_id"] = qcID })
}

// WithDfoID dfo_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithDfoID(dfoID int64) Option {
	return optionFunc(func(o *options) { o.query["dfo_id"] = dfoID })
}

// WithSqcID sqc_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithSqcID(sqcID int64) Option {
	return optionFunc(func(o *options) { o.query["sqc_id"] = sqcID })
}

// WithWorkerID worker_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithWorkerID(workerID int64) Option {
	return optionFunc(func(o *options) { o.query["worker_id"] = workerID })
}

// WithEvent event获取 
func (obj *_AllVirtualSQLAuditMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithP1text p1text获取 
func (obj *_AllVirtualSQLAuditMgr) WithP1text(p1text string) Option {
	return optionFunc(func(o *options) { o.query["p1text"] = p1text })
}

// WithP1 p1获取 
func (obj *_AllVirtualSQLAuditMgr) WithP1(p1 uint64) Option {
	return optionFunc(func(o *options) { o.query["p1"] = p1 })
}

// WithP2text p2text获取 
func (obj *_AllVirtualSQLAuditMgr) WithP2text(p2text string) Option {
	return optionFunc(func(o *options) { o.query["p2text"] = p2text })
}

// WithP2 p2获取 
func (obj *_AllVirtualSQLAuditMgr) WithP2(p2 uint64) Option {
	return optionFunc(func(o *options) { o.query["p2"] = p2 })
}

// WithP3text p3text获取 
func (obj *_AllVirtualSQLAuditMgr) WithP3text(p3text string) Option {
	return optionFunc(func(o *options) { o.query["p3text"] = p3text })
}

// WithP3 p3获取 
func (obj *_AllVirtualSQLAuditMgr) WithP3(p3 uint64) Option {
	return optionFunc(func(o *options) { o.query["p3"] = p3 })
}

// WithLevel level获取 
func (obj *_AllVirtualSQLAuditMgr) WithLevel(level int64) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithWaitClassID wait_class_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class_id"] = waitClassID })
}

// WithWaitClass# wait_class#获取 
func (obj *_AllVirtualSQLAuditMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class#"] = waitClass# })
}

// WithWaitClass wait_class获取 
func (obj *_AllVirtualSQLAuditMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["wait_class"] = waitClass })
}

// WithState state获取 
func (obj *_AllVirtualSQLAuditMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithWaitTimeMicro wait_time_micro获取 
func (obj *_AllVirtualSQLAuditMgr) WithWaitTimeMicro(waitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["wait_time_micro"] = waitTimeMicro })
}

// WithTotalWaitTimeMicro total_wait_time_micro获取 
func (obj *_AllVirtualSQLAuditMgr) WithTotalWaitTimeMicro(totalWaitTimeMicro int64) Option {
	return optionFunc(func(o *options) { o.query["total_wait_time_micro"] = totalWaitTimeMicro })
}

// WithTotalWaits total_waits获取 
func (obj *_AllVirtualSQLAuditMgr) WithTotalWaits(totalWaits int64) Option {
	return optionFunc(func(o *options) { o.query["total_waits"] = totalWaits })
}

// WithRPCCount rpc_count获取 
func (obj *_AllVirtualSQLAuditMgr) WithRPCCount(rpcCount int64) Option {
	return optionFunc(func(o *options) { o.query["rpc_count"] = rpcCount })
}

// WithPlanType plan_type获取 
func (obj *_AllVirtualSQLAuditMgr) WithPlanType(planType int64) Option {
	return optionFunc(func(o *options) { o.query["plan_type"] = planType })
}

// WithIsInnerSQL is_inner_sql获取 
func (obj *_AllVirtualSQLAuditMgr) WithIsInnerSQL(isInnerSQL int8) Option {
	return optionFunc(func(o *options) { o.query["is_inner_sql"] = isInnerSQL })
}

// WithIsExecutorRPC is_executor_rpc获取 
func (obj *_AllVirtualSQLAuditMgr) WithIsExecutorRPC(isExecutorRPC int8) Option {
	return optionFunc(func(o *options) { o.query["is_executor_rpc"] = isExecutorRPC })
}

// WithIsHitPlan is_hit_plan获取 
func (obj *_AllVirtualSQLAuditMgr) WithIsHitPlan(isHitPlan int8) Option {
	return optionFunc(func(o *options) { o.query["is_hit_plan"] = isHitPlan })
}

// WithRequestTime request_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithRequestTime(requestTime int64) Option {
	return optionFunc(func(o *options) { o.query["request_time"] = requestTime })
}

// WithElapsedTime elapsed_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithElapsedTime(elapsedTime int64) Option {
	return optionFunc(func(o *options) { o.query["elapsed_time"] = elapsedTime })
}

// WithNetTime net_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithNetTime(netTime int64) Option {
	return optionFunc(func(o *options) { o.query["net_time"] = netTime })
}

// WithNetWaitTime net_wait_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithNetWaitTime(netWaitTime int64) Option {
	return optionFunc(func(o *options) { o.query["net_wait_time"] = netWaitTime })
}

// WithQueueTime queue_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithQueueTime(queueTime int64) Option {
	return optionFunc(func(o *options) { o.query["queue_time"] = queueTime })
}

// WithDecodeTime decode_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithDecodeTime(decodeTime int64) Option {
	return optionFunc(func(o *options) { o.query["decode_time"] = decodeTime })
}

// WithGetPlanTime get_plan_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithGetPlanTime(getPlanTime int64) Option {
	return optionFunc(func(o *options) { o.query["get_plan_time"] = getPlanTime })
}

// WithExecuteTime execute_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithExecuteTime(executeTime int64) Option {
	return optionFunc(func(o *options) { o.query["execute_time"] = executeTime })
}

// WithApplicationWaitTime application_wait_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithApplicationWaitTime(applicationWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["application_wait_time"] = applicationWaitTime })
}

// WithConcurrencyWaitTime concurrency_wait_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithConcurrencyWaitTime(concurrencyWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["concurrency_wait_time"] = concurrencyWaitTime })
}

// WithUserIoWaitTime user_io_wait_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithUserIoWaitTime(userIoWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["user_io_wait_time"] = userIoWaitTime })
}

// WithScheduleTime schedule_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithScheduleTime(scheduleTime uint64) Option {
	return optionFunc(func(o *options) { o.query["schedule_time"] = scheduleTime })
}

// WithRowCacheHit row_cache_hit获取 
func (obj *_AllVirtualSQLAuditMgr) WithRowCacheHit(rowCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["row_cache_hit"] = rowCacheHit })
}

// WithBloomFilterCacheHit bloom_filter_cache_hit获取 
func (obj *_AllVirtualSQLAuditMgr) WithBloomFilterCacheHit(bloomFilterCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["bloom_filter_cache_hit"] = bloomFilterCacheHit })
}

// WithBlockCacheHit block_cache_hit获取 
func (obj *_AllVirtualSQLAuditMgr) WithBlockCacheHit(blockCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["block_cache_hit"] = blockCacheHit })
}

// WithBlockIndexCacheHit block_index_cache_hit获取 
func (obj *_AllVirtualSQLAuditMgr) WithBlockIndexCacheHit(blockIndexCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["block_index_cache_hit"] = blockIndexCacheHit })
}

// WithDiskReads disk_reads获取 
func (obj *_AllVirtualSQLAuditMgr) WithDiskReads(diskReads int64) Option {
	return optionFunc(func(o *options) { o.query["disk_reads"] = diskReads })
}

// WithExecutionID execution_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithExecutionID(executionID int64) Option {
	return optionFunc(func(o *options) { o.query["execution_id"] = executionID })
}

// WithSessionID session_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithSessionID(sessionID uint64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithRetryCnt retry_cnt获取 
func (obj *_AllVirtualSQLAuditMgr) WithRetryCnt(retryCnt int64) Option {
	return optionFunc(func(o *options) { o.query["retry_cnt"] = retryCnt })
}

// WithTableScan table_scan获取 
func (obj *_AllVirtualSQLAuditMgr) WithTableScan(tableScan int8) Option {
	return optionFunc(func(o *options) { o.query["table_scan"] = tableScan })
}

// WithConsistencyLevel consistency_level获取 
func (obj *_AllVirtualSQLAuditMgr) WithConsistencyLevel(consistencyLevel int64) Option {
	return optionFunc(func(o *options) { o.query["consistency_level"] = consistencyLevel })
}

// WithMemstoreReadRowCount memstore_read_row_count获取 
func (obj *_AllVirtualSQLAuditMgr) WithMemstoreReadRowCount(memstoreReadRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_read_row_count"] = memstoreReadRowCount })
}

// WithSsstoreReadRowCount ssstore_read_row_count获取 
func (obj *_AllVirtualSQLAuditMgr) WithSsstoreReadRowCount(ssstoreReadRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["ssstore_read_row_count"] = ssstoreReadRowCount })
}

// WithRequestMemoryUsed request_memory_used获取 
func (obj *_AllVirtualSQLAuditMgr) WithRequestMemoryUsed(requestMemoryUsed int64) Option {
	return optionFunc(func(o *options) { o.query["request_memory_used"] = requestMemoryUsed })
}

// WithExpectedWorkerCount expected_worker_count获取 
func (obj *_AllVirtualSQLAuditMgr) WithExpectedWorkerCount(expectedWorkerCount int64) Option {
	return optionFunc(func(o *options) { o.query["expected_worker_count"] = expectedWorkerCount })
}

// WithUsedWorkerCount used_worker_count获取 
func (obj *_AllVirtualSQLAuditMgr) WithUsedWorkerCount(usedWorkerCount int64) Option {
	return optionFunc(func(o *options) { o.query["used_worker_count"] = usedWorkerCount })
}

// WithSchedInfo sched_info获取 
func (obj *_AllVirtualSQLAuditMgr) WithSchedInfo(schedInfo string) Option {
	return optionFunc(func(o *options) { o.query["sched_info"] = schedInfo })
}

// WithFuseRowCacheHit fuse_row_cache_hit获取 
func (obj *_AllVirtualSQLAuditMgr) WithFuseRowCacheHit(fuseRowCacheHit int64) Option {
	return optionFunc(func(o *options) { o.query["fuse_row_cache_hit"] = fuseRowCacheHit })
}

// WithUserClientIP user_client_ip获取 
func (obj *_AllVirtualSQLAuditMgr) WithUserClientIP(userClientIP string) Option {
	return optionFunc(func(o *options) { o.query["user_client_ip"] = userClientIP })
}

// WithPsStmtID ps_stmt_id获取 
func (obj *_AllVirtualSQLAuditMgr) WithPsStmtID(psStmtID int64) Option {
	return optionFunc(func(o *options) { o.query["ps_stmt_id"] = psStmtID })
}

// WithTransactionHash transaction_hash获取 
func (obj *_AllVirtualSQLAuditMgr) WithTransactionHash(transactionHash uint64) Option {
	return optionFunc(func(o *options) { o.query["transaction_hash"] = transactionHash })
}

// WithRequestType request_type获取 
func (obj *_AllVirtualSQLAuditMgr) WithRequestType(requestType int64) Option {
	return optionFunc(func(o *options) { o.query["request_type"] = requestType })
}

// WithIsBatchedMultiStmt is_batched_multi_stmt获取 
func (obj *_AllVirtualSQLAuditMgr) WithIsBatchedMultiStmt(isBatchedMultiStmt int8) Option {
	return optionFunc(func(o *options) { o.query["is_batched_multi_stmt"] = isBatchedMultiStmt })
}

// WithObTraceInfo ob_trace_info获取 
func (obj *_AllVirtualSQLAuditMgr) WithObTraceInfo(obTraceInfo string) Option {
	return optionFunc(func(o *options) { o.query["ob_trace_info"] = obTraceInfo })
}

// WithPlanHash plan_hash获取 
func (obj *_AllVirtualSQLAuditMgr) WithPlanHash(planHash uint64) Option {
	return optionFunc(func(o *options) { o.query["plan_hash"] = planHash })
}

// WithUserGroup user_group获取 
func (obj *_AllVirtualSQLAuditMgr) WithUserGroup(userGroup int64) Option {
	return optionFunc(func(o *options) { o.query["user_group"] = userGroup })
}

// WithLockForReadTime lock_for_read_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithLockForReadTime(lockForReadTime int64) Option {
	return optionFunc(func(o *options) { o.query["lock_for_read_time"] = lockForReadTime })
}

// WithWaitTrxMigrateTime wait_trx_migrate_time获取 
func (obj *_AllVirtualSQLAuditMgr) WithWaitTrxMigrateTime(waitTrxMigrateTime int64) Option {
	return optionFunc(func(o *options) { o.query["wait_trx_migrate_time"] = waitTrxMigrateTime })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLAuditMgr) GetByOption(opts ...Option) (result AllVirtualSQLAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLAuditMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualSQLAuditMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLAudit,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where(options.query)
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


// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromRequestID 通过request_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRequestID(requestID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_id` = ?", requestID).Find(&results).Error
	
	return
}

// GetBatchFromRequestID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRequestID(requestIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_id` IN (?)", requestIDs).Find(&results).Error
	
	return
}
 
// GetFromTraceID 通过trace_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromTraceID(traceID string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`trace_id` = ?", traceID).Find(&results).Error
	
	return
}

// GetBatchFromTraceID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error
	
	return
}
 
// GetFromClientIP 通过client_ip获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromClientIP(clientIP string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`client_ip` = ?", clientIP).Find(&results).Error
	
	return
}

// GetBatchFromClientIP 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromClientIP(clientIPs []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`client_ip` IN (?)", clientIPs).Find(&results).Error
	
	return
}
 
// GetFromClientPort 通过client_port获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromClientPort(clientPort int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`client_port` = ?", clientPort).Find(&results).Error
	
	return
}

// GetBatchFromClientPort 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromClientPort(clientPorts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`client_port` IN (?)", clientPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantName 通过tenant_name获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromTenantName(tenantName string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error
	
	return
}

// GetBatchFromTenantName 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error
	
	return
}
 
// GetFromEffectiveTenantID 通过effective_tenant_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromEffectiveTenantID(effectiveTenantID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`effective_tenant_id` = ?", effectiveTenantID).Find(&results).Error
	
	return
}

// GetBatchFromEffectiveTenantID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromEffectiveTenantID(effectiveTenantIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`effective_tenant_id` IN (?)", effectiveTenantIDs).Find(&results).Error
	
	return
}
 
// GetFromUserID 通过user_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromUserID(userID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_id` = ?", userID).Find(&results).Error
	
	return
}

// GetBatchFromUserID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	
	return
}
 
// GetFromUserName 通过user_name获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromUserName(userName string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_name` = ?", userName).Find(&results).Error
	
	return
}

// GetBatchFromUserName 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_name` IN (?)", userNames).Find(&results).Error
	
	return
}
 
// GetFromDbID 通过db_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromDbID(dbID uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`db_id` = ?", dbID).Find(&results).Error
	
	return
}

// GetBatchFromDbID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromDbID(dbIDs []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`db_id` IN (?)", dbIDs).Find(&results).Error
	
	return
}
 
// GetFromDbName 通过db_name获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromDbName(dbName string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`db_name` = ?", dbName).Find(&results).Error
	
	return
}

// GetBatchFromDbName 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromDbName(dbNames []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`db_name` IN (?)", dbNames).Find(&results).Error
	
	return
}
 
// GetFromSQLID 通过sql_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromSQLID(sqlID string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`sql_id` = ?", sqlID).Find(&results).Error
	
	return
}

// GetBatchFromSQLID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromSQLID(sqlIDs []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error
	
	return
}
 
// GetFromQuerySQL 通过query_sql获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromQuerySQL(querySQL string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`query_sql` = ?", querySQL).Find(&results).Error
	
	return
}

// GetBatchFromQuerySQL 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromQuerySQL(querySQLs []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`query_sql` IN (?)", querySQLs).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过plan_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromPlanID(planID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`plan_id` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromPlanID(planIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromAffectedRows 通过affected_rows获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromAffectedRows(affectedRows int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`affected_rows` = ?", affectedRows).Find(&results).Error
	
	return
}

// GetBatchFromAffectedRows 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromAffectedRows(affectedRowss []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`affected_rows` IN (?)", affectedRowss).Find(&results).Error
	
	return
}
 
// GetFromReturnRows 通过return_rows获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromReturnRows(returnRows int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`return_rows` = ?", returnRows).Find(&results).Error
	
	return
}

// GetBatchFromReturnRows 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromReturnRows(returnRowss []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`return_rows` IN (?)", returnRowss).Find(&results).Error
	
	return
}
 
// GetFromPartitionCnt 通过partition_cnt获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error
	
	return
}

// GetBatchFromPartitionCnt 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error
	
	return
}
 
// GetFromRetCode 通过ret_code获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRetCode(retCode int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ret_code` = ?", retCode).Find(&results).Error
	
	return
}

// GetBatchFromRetCode 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRetCode(retCodes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ret_code` IN (?)", retCodes).Find(&results).Error
	
	return
}
 
// GetFromQcID 通过qc_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromQcID(qcID uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`qc_id` = ?", qcID).Find(&results).Error
	
	return
}

// GetBatchFromQcID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromQcID(qcIDs []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`qc_id` IN (?)", qcIDs).Find(&results).Error
	
	return
}
 
// GetFromDfoID 通过dfo_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromDfoID(dfoID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`dfo_id` = ?", dfoID).Find(&results).Error
	
	return
}

// GetBatchFromDfoID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromDfoID(dfoIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`dfo_id` IN (?)", dfoIDs).Find(&results).Error
	
	return
}
 
// GetFromSqcID 通过sqc_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromSqcID(sqcID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`sqc_id` = ?", sqcID).Find(&results).Error
	
	return
}

// GetBatchFromSqcID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromSqcID(sqcIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`sqc_id` IN (?)", sqcIDs).Find(&results).Error
	
	return
}
 
// GetFromWorkerID 通过worker_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromWorkerID(workerID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`worker_id` = ?", workerID).Find(&results).Error
	
	return
}

// GetBatchFromWorkerID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromWorkerID(workerIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`worker_id` IN (?)", workerIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent 通过event获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromEvent(event string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`event` = ?", event).Find(&results).Error
	
	return
}

// GetBatchFromEvent 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromEvent(events []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`event` IN (?)", events).Find(&results).Error
	
	return
}
 
// GetFromP1text 通过p1text获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromP1text(p1text string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p1text` = ?", p1text).Find(&results).Error
	
	return
}

// GetBatchFromP1text 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromP1text(p1texts []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p1text` IN (?)", p1texts).Find(&results).Error
	
	return
}
 
// GetFromP1 通过p1获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromP1(p1 uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p1` = ?", p1).Find(&results).Error
	
	return
}

// GetBatchFromP1 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromP1(p1s []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p1` IN (?)", p1s).Find(&results).Error
	
	return
}
 
// GetFromP2text 通过p2text获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromP2text(p2text string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p2text` = ?", p2text).Find(&results).Error
	
	return
}

// GetBatchFromP2text 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromP2text(p2texts []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p2text` IN (?)", p2texts).Find(&results).Error
	
	return
}
 
// GetFromP2 通过p2获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromP2(p2 uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p2` = ?", p2).Find(&results).Error
	
	return
}

// GetBatchFromP2 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromP2(p2s []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p2` IN (?)", p2s).Find(&results).Error
	
	return
}
 
// GetFromP3text 通过p3text获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromP3text(p3text string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p3text` = ?", p3text).Find(&results).Error
	
	return
}

// GetBatchFromP3text 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromP3text(p3texts []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p3text` IN (?)", p3texts).Find(&results).Error
	
	return
}
 
// GetFromP3 通过p3获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromP3(p3 uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p3` = ?", p3).Find(&results).Error
	
	return
}

// GetBatchFromP3 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromP3(p3s []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`p3` IN (?)", p3s).Find(&results).Error
	
	return
}
 
// GetFromLevel 通过level获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromLevel(level int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`level` = ?", level).Find(&results).Error
	
	return
}

// GetBatchFromLevel 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromLevel(levels []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`level` IN (?)", levels).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过wait_class_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromWaitClassID(waitClassID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_class_id` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_class_id` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过wait_class#获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromWaitClass#(waitClass# int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_class#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_class#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过wait_class获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromWaitClass(waitClass string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_class` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromWaitClass(waitClasss []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_class` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
// GetFromState 通过state获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromState(state string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`state` = ?", state).Find(&results).Error
	
	return
}

// GetBatchFromState 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromState(states []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`state` IN (?)", states).Find(&results).Error
	
	return
}
 
// GetFromWaitTimeMicro 通过wait_time_micro获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromWaitTimeMicro(waitTimeMicro int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_time_micro` = ?", waitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromWaitTimeMicro 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromWaitTimeMicro(waitTimeMicros []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_time_micro` IN (?)", waitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTotalWaitTimeMicro 通过total_wait_time_micro获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromTotalWaitTimeMicro(totalWaitTimeMicro int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`total_wait_time_micro` = ?", totalWaitTimeMicro).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaitTimeMicro 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromTotalWaitTimeMicro(totalWaitTimeMicros []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`total_wait_time_micro` IN (?)", totalWaitTimeMicros).Find(&results).Error
	
	return
}
 
// GetFromTotalWaits 通过total_waits获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromTotalWaits(totalWaits int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`total_waits` = ?", totalWaits).Find(&results).Error
	
	return
}

// GetBatchFromTotalWaits 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromTotalWaits(totalWaitss []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`total_waits` IN (?)", totalWaitss).Find(&results).Error
	
	return
}
 
// GetFromRPCCount 通过rpc_count获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRPCCount(rpcCount int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`rpc_count` = ?", rpcCount).Find(&results).Error
	
	return
}

// GetBatchFromRPCCount 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRPCCount(rpcCounts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`rpc_count` IN (?)", rpcCounts).Find(&results).Error
	
	return
}
 
// GetFromPlanType 通过plan_type获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromPlanType(planType int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`plan_type` = ?", planType).Find(&results).Error
	
	return
}

// GetBatchFromPlanType 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromPlanType(planTypes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`plan_type` IN (?)", planTypes).Find(&results).Error
	
	return
}
 
// GetFromIsInnerSQL 通过is_inner_sql获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromIsInnerSQL(isInnerSQL int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_inner_sql` = ?", isInnerSQL).Find(&results).Error
	
	return
}

// GetBatchFromIsInnerSQL 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromIsInnerSQL(isInnerSQLs []int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_inner_sql` IN (?)", isInnerSQLs).Find(&results).Error
	
	return
}
 
// GetFromIsExecutorRPC 通过is_executor_rpc获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromIsExecutorRPC(isExecutorRPC int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_executor_rpc` = ?", isExecutorRPC).Find(&results).Error
	
	return
}

// GetBatchFromIsExecutorRPC 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromIsExecutorRPC(isExecutorRPCs []int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_executor_rpc` IN (?)", isExecutorRPCs).Find(&results).Error
	
	return
}
 
// GetFromIsHitPlan 通过is_hit_plan获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromIsHitPlan(isHitPlan int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_hit_plan` = ?", isHitPlan).Find(&results).Error
	
	return
}

// GetBatchFromIsHitPlan 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromIsHitPlan(isHitPlans []int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_hit_plan` IN (?)", isHitPlans).Find(&results).Error
	
	return
}
 
// GetFromRequestTime 通过request_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRequestTime(requestTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_time` = ?", requestTime).Find(&results).Error
	
	return
}

// GetBatchFromRequestTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRequestTime(requestTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_time` IN (?)", requestTimes).Find(&results).Error
	
	return
}
 
// GetFromElapsedTime 通过elapsed_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromElapsedTime(elapsedTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`elapsed_time` = ?", elapsedTime).Find(&results).Error
	
	return
}

// GetBatchFromElapsedTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromElapsedTime(elapsedTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`elapsed_time` IN (?)", elapsedTimes).Find(&results).Error
	
	return
}
 
// GetFromNetTime 通过net_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromNetTime(netTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`net_time` = ?", netTime).Find(&results).Error
	
	return
}

// GetBatchFromNetTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromNetTime(netTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`net_time` IN (?)", netTimes).Find(&results).Error
	
	return
}
 
// GetFromNetWaitTime 通过net_wait_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromNetWaitTime(netWaitTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`net_wait_time` = ?", netWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromNetWaitTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromNetWaitTime(netWaitTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`net_wait_time` IN (?)", netWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromQueueTime 通过queue_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromQueueTime(queueTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`queue_time` = ?", queueTime).Find(&results).Error
	
	return
}

// GetBatchFromQueueTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromQueueTime(queueTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`queue_time` IN (?)", queueTimes).Find(&results).Error
	
	return
}
 
// GetFromDecodeTime 通过decode_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromDecodeTime(decodeTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`decode_time` = ?", decodeTime).Find(&results).Error
	
	return
}

// GetBatchFromDecodeTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromDecodeTime(decodeTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`decode_time` IN (?)", decodeTimes).Find(&results).Error
	
	return
}
 
// GetFromGetPlanTime 通过get_plan_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromGetPlanTime(getPlanTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`get_plan_time` = ?", getPlanTime).Find(&results).Error
	
	return
}

// GetBatchFromGetPlanTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromGetPlanTime(getPlanTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`get_plan_time` IN (?)", getPlanTimes).Find(&results).Error
	
	return
}
 
// GetFromExecuteTime 通过execute_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromExecuteTime(executeTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`execute_time` = ?", executeTime).Find(&results).Error
	
	return
}

// GetBatchFromExecuteTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromExecuteTime(executeTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`execute_time` IN (?)", executeTimes).Find(&results).Error
	
	return
}
 
// GetFromApplicationWaitTime 通过application_wait_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromApplicationWaitTime(applicationWaitTime uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`application_wait_time` = ?", applicationWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromApplicationWaitTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromApplicationWaitTime(applicationWaitTimes []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`application_wait_time` IN (?)", applicationWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromConcurrencyWaitTime 通过concurrency_wait_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromConcurrencyWaitTime(concurrencyWaitTime uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`concurrency_wait_time` = ?", concurrencyWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromConcurrencyWaitTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromConcurrencyWaitTime(concurrencyWaitTimes []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`concurrency_wait_time` IN (?)", concurrencyWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromUserIoWaitTime 通过user_io_wait_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromUserIoWaitTime(userIoWaitTime uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_io_wait_time` = ?", userIoWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromUserIoWaitTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromUserIoWaitTime(userIoWaitTimes []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_io_wait_time` IN (?)", userIoWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromScheduleTime 通过schedule_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromScheduleTime(scheduleTime uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`schedule_time` = ?", scheduleTime).Find(&results).Error
	
	return
}

// GetBatchFromScheduleTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromScheduleTime(scheduleTimes []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`schedule_time` IN (?)", scheduleTimes).Find(&results).Error
	
	return
}
 
// GetFromRowCacheHit 通过row_cache_hit获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRowCacheHit(rowCacheHit int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`row_cache_hit` = ?", rowCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromRowCacheHit 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRowCacheHit(rowCacheHits []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`row_cache_hit` IN (?)", rowCacheHits).Find(&results).Error
	
	return
}
 
// GetFromBloomFilterCacheHit 通过bloom_filter_cache_hit获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromBloomFilterCacheHit(bloomFilterCacheHit int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`bloom_filter_cache_hit` = ?", bloomFilterCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromBloomFilterCacheHit 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromBloomFilterCacheHit(bloomFilterCacheHits []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`bloom_filter_cache_hit` IN (?)", bloomFilterCacheHits).Find(&results).Error
	
	return
}
 
// GetFromBlockCacheHit 通过block_cache_hit获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromBlockCacheHit(blockCacheHit int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`block_cache_hit` = ?", blockCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromBlockCacheHit 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromBlockCacheHit(blockCacheHits []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`block_cache_hit` IN (?)", blockCacheHits).Find(&results).Error
	
	return
}
 
// GetFromBlockIndexCacheHit 通过block_index_cache_hit获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromBlockIndexCacheHit(blockIndexCacheHit int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`block_index_cache_hit` = ?", blockIndexCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromBlockIndexCacheHit 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromBlockIndexCacheHit(blockIndexCacheHits []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`block_index_cache_hit` IN (?)", blockIndexCacheHits).Find(&results).Error
	
	return
}
 
// GetFromDiskReads 通过disk_reads获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromDiskReads(diskReads int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`disk_reads` = ?", diskReads).Find(&results).Error
	
	return
}

// GetBatchFromDiskReads 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromDiskReads(diskReadss []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`disk_reads` IN (?)", diskReadss).Find(&results).Error
	
	return
}
 
// GetFromExecutionID 通过execution_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromExecutionID(executionID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`execution_id` = ?", executionID).Find(&results).Error
	
	return
}

// GetBatchFromExecutionID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromExecutionID(executionIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`execution_id` IN (?)", executionIDs).Find(&results).Error
	
	return
}
 
// GetFromSessionID 通过session_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromSessionID(sessionID uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromSessionID(sessionIDs []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromRetryCnt 通过retry_cnt获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRetryCnt(retryCnt int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`retry_cnt` = ?", retryCnt).Find(&results).Error
	
	return
}

// GetBatchFromRetryCnt 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRetryCnt(retryCnts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`retry_cnt` IN (?)", retryCnts).Find(&results).Error
	
	return
}
 
// GetFromTableScan 通过table_scan获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromTableScan(tableScan int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`table_scan` = ?", tableScan).Find(&results).Error
	
	return
}

// GetBatchFromTableScan 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromTableScan(tableScans []int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`table_scan` IN (?)", tableScans).Find(&results).Error
	
	return
}
 
// GetFromConsistencyLevel 通过consistency_level获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromConsistencyLevel(consistencyLevel int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`consistency_level` = ?", consistencyLevel).Find(&results).Error
	
	return
}

// GetBatchFromConsistencyLevel 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromConsistencyLevel(consistencyLevels []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`consistency_level` IN (?)", consistencyLevels).Find(&results).Error
	
	return
}
 
// GetFromMemstoreReadRowCount 通过memstore_read_row_count获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromMemstoreReadRowCount(memstoreReadRowCount int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`memstore_read_row_count` = ?", memstoreReadRowCount).Find(&results).Error
	
	return
}

// GetBatchFromMemstoreReadRowCount 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromMemstoreReadRowCount(memstoreReadRowCounts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`memstore_read_row_count` IN (?)", memstoreReadRowCounts).Find(&results).Error
	
	return
}
 
// GetFromSsstoreReadRowCount 通过ssstore_read_row_count获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromSsstoreReadRowCount(ssstoreReadRowCount int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ssstore_read_row_count` = ?", ssstoreReadRowCount).Find(&results).Error
	
	return
}

// GetBatchFromSsstoreReadRowCount 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromSsstoreReadRowCount(ssstoreReadRowCounts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ssstore_read_row_count` IN (?)", ssstoreReadRowCounts).Find(&results).Error
	
	return
}
 
// GetFromRequestMemoryUsed 通过request_memory_used获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRequestMemoryUsed(requestMemoryUsed int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_memory_used` = ?", requestMemoryUsed).Find(&results).Error
	
	return
}

// GetBatchFromRequestMemoryUsed 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRequestMemoryUsed(requestMemoryUseds []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_memory_used` IN (?)", requestMemoryUseds).Find(&results).Error
	
	return
}
 
// GetFromExpectedWorkerCount 通过expected_worker_count获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromExpectedWorkerCount(expectedWorkerCount int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`expected_worker_count` = ?", expectedWorkerCount).Find(&results).Error
	
	return
}

// GetBatchFromExpectedWorkerCount 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromExpectedWorkerCount(expectedWorkerCounts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`expected_worker_count` IN (?)", expectedWorkerCounts).Find(&results).Error
	
	return
}
 
// GetFromUsedWorkerCount 通过used_worker_count获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromUsedWorkerCount(usedWorkerCount int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`used_worker_count` = ?", usedWorkerCount).Find(&results).Error
	
	return
}

// GetBatchFromUsedWorkerCount 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromUsedWorkerCount(usedWorkerCounts []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`used_worker_count` IN (?)", usedWorkerCounts).Find(&results).Error
	
	return
}
 
// GetFromSchedInfo 通过sched_info获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromSchedInfo(schedInfo string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`sched_info` = ?", schedInfo).Find(&results).Error
	
	return
}

// GetBatchFromSchedInfo 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromSchedInfo(schedInfos []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`sched_info` IN (?)", schedInfos).Find(&results).Error
	
	return
}
 
// GetFromFuseRowCacheHit 通过fuse_row_cache_hit获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromFuseRowCacheHit(fuseRowCacheHit int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`fuse_row_cache_hit` = ?", fuseRowCacheHit).Find(&results).Error
	
	return
}

// GetBatchFromFuseRowCacheHit 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromFuseRowCacheHit(fuseRowCacheHits []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`fuse_row_cache_hit` IN (?)", fuseRowCacheHits).Find(&results).Error
	
	return
}
 
// GetFromUserClientIP 通过user_client_ip获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromUserClientIP(userClientIP string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_client_ip` = ?", userClientIP).Find(&results).Error
	
	return
}

// GetBatchFromUserClientIP 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromUserClientIP(userClientIPs []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_client_ip` IN (?)", userClientIPs).Find(&results).Error
	
	return
}
 
// GetFromPsStmtID 通过ps_stmt_id获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromPsStmtID(psStmtID int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ps_stmt_id` = ?", psStmtID).Find(&results).Error
	
	return
}

// GetBatchFromPsStmtID 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromPsStmtID(psStmtIDs []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ps_stmt_id` IN (?)", psStmtIDs).Find(&results).Error
	
	return
}
 
// GetFromTransactionHash 通过transaction_hash获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromTransactionHash(transactionHash uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`transaction_hash` = ?", transactionHash).Find(&results).Error
	
	return
}

// GetBatchFromTransactionHash 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromTransactionHash(transactionHashs []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`transaction_hash` IN (?)", transactionHashs).Find(&results).Error
	
	return
}
 
// GetFromRequestType 通过request_type获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromRequestType(requestType int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_type` = ?", requestType).Find(&results).Error
	
	return
}

// GetBatchFromRequestType 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromRequestType(requestTypes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`request_type` IN (?)", requestTypes).Find(&results).Error
	
	return
}
 
// GetFromIsBatchedMultiStmt 通过is_batched_multi_stmt获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromIsBatchedMultiStmt(isBatchedMultiStmt int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_batched_multi_stmt` = ?", isBatchedMultiStmt).Find(&results).Error
	
	return
}

// GetBatchFromIsBatchedMultiStmt 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromIsBatchedMultiStmt(isBatchedMultiStmts []int8) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`is_batched_multi_stmt` IN (?)", isBatchedMultiStmts).Find(&results).Error
	
	return
}
 
// GetFromObTraceInfo 通过ob_trace_info获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromObTraceInfo(obTraceInfo string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ob_trace_info` = ?", obTraceInfo).Find(&results).Error
	
	return
}

// GetBatchFromObTraceInfo 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromObTraceInfo(obTraceInfos []string) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`ob_trace_info` IN (?)", obTraceInfos).Find(&results).Error
	
	return
}
 
// GetFromPlanHash 通过plan_hash获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromPlanHash(planHash uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`plan_hash` = ?", planHash).Find(&results).Error
	
	return
}

// GetBatchFromPlanHash 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromPlanHash(planHashs []uint64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`plan_hash` IN (?)", planHashs).Find(&results).Error
	
	return
}
 
// GetFromUserGroup 通过user_group获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromUserGroup(userGroup int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_group` = ?", userGroup).Find(&results).Error
	
	return
}

// GetBatchFromUserGroup 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromUserGroup(userGroups []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`user_group` IN (?)", userGroups).Find(&results).Error
	
	return
}
 
// GetFromLockForReadTime 通过lock_for_read_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromLockForReadTime(lockForReadTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`lock_for_read_time` = ?", lockForReadTime).Find(&results).Error
	
	return
}

// GetBatchFromLockForReadTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromLockForReadTime(lockForReadTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`lock_for_read_time` IN (?)", lockForReadTimes).Find(&results).Error
	
	return
}
 
// GetFromWaitTrxMigrateTime 通过wait_trx_migrate_time获取内容  
func (obj *_AllVirtualSQLAuditMgr) GetFromWaitTrxMigrateTime(waitTrxMigrateTime int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_trx_migrate_time` = ?", waitTrxMigrateTime).Find(&results).Error
	
	return
}

// GetBatchFromWaitTrxMigrateTime 批量查找 
func (obj *_AllVirtualSQLAuditMgr) GetBatchFromWaitTrxMigrateTime(waitTrxMigrateTimes []int64) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`wait_trx_migrate_time` IN (?)", waitTrxMigrateTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllVirtualSQLAuditMgr) FetchByPrimaryKey(svrIP string ,svrPort int64 ,tenantID int64 ,requestID int64 ) (result AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `request_id` = ?", svrIP , svrPort , tenantID , requestID).First(&result).Error
	
	return
}
 

 
 // FetchIndexByAllVirtualSQLAuditI1  获取多个内容
 func (obj *_AllVirtualSQLAuditMgr) FetchIndexByAllVirtualSQLAuditI1(tenantID int64 ,requestID int64 ) (results []*AllVirtualSQLAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLAudit{}).Where("`tenant_id` = ? AND `request_id` = ?", tenantID , requestID).Find(&results).Error
	
	return
}
 

	

