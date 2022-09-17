package	model	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _V$planCachePlanStatMgr struct {
	*_BaseMgr
}

// V$planCachePlanStatMgr open func
func V$planCachePlanStatMgr(db *gorm.DB) *_V$planCachePlanStatMgr {
	if db == nil {
		panic(fmt.Errorf("V$planCachePlanStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$planCachePlanStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$plan_cache_plan_stat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$planCachePlanStatMgr) GetTableName() string {
	return "v$plan_cache_plan_stat"
}

// Reset 重置gorm会话
func (obj *_V$planCachePlanStatMgr) Reset() *_V$planCachePlanStatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$planCachePlanStatMgr) Get() (result V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$planCachePlanStatMgr) Gets() (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$planCachePlanStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_V$planCachePlanStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取 
func (obj *_V$planCachePlanStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_V$planCachePlanStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPlanID plan_id获取 
func (obj *_V$planCachePlanStatMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithSQLID sql_id获取 
func (obj *_V$planCachePlanStatMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithType type获取 
func (obj *_V$planCachePlanStatMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIsBindSensitive is_bind_sensitive获取 
func (obj *_V$planCachePlanStatMgr) WithIsBindSensitive(isBindSensitive int64) Option {
	return optionFunc(func(o *options) { o.query["is_bind_sensitive"] = isBindSensitive })
}

// WithIsBindAware is_bind_aware获取 
func (obj *_V$planCachePlanStatMgr) WithIsBindAware(isBindAware int64) Option {
	return optionFunc(func(o *options) { o.query["is_bind_aware"] = isBindAware })
}

// WithDbID db_id获取 
func (obj *_V$planCachePlanStatMgr) WithDbID(dbID uint64) Option {
	return optionFunc(func(o *options) { o.query["db_id"] = dbID })
}

// WithStatement statement获取 
func (obj *_V$planCachePlanStatMgr) WithStatement(statement string) Option {
	return optionFunc(func(o *options) { o.query["statement"] = statement })
}

// WithQuerySQL query_sql获取 
func (obj *_V$planCachePlanStatMgr) WithQuerySQL(querySQL string) Option {
	return optionFunc(func(o *options) { o.query["query_sql"] = querySQL })
}

// WithSpecialParams special_params获取 
func (obj *_V$planCachePlanStatMgr) WithSpecialParams(specialParams string) Option {
	return optionFunc(func(o *options) { o.query["special_params"] = specialParams })
}

// WithParamInfos param_infos获取 
func (obj *_V$planCachePlanStatMgr) WithParamInfos(paramInfos string) Option {
	return optionFunc(func(o *options) { o.query["param_infos"] = paramInfos })
}

// WithSysVars sys_vars获取 
func (obj *_V$planCachePlanStatMgr) WithSysVars(sysVars string) Option {
	return optionFunc(func(o *options) { o.query["sys_vars"] = sysVars })
}

// WithPlanHash plan_hash获取 
func (obj *_V$planCachePlanStatMgr) WithPlanHash(planHash uint64) Option {
	return optionFunc(func(o *options) { o.query["plan_hash"] = planHash })
}

// WithFirstLoadTime first_load_time获取 
func (obj *_V$planCachePlanStatMgr) WithFirstLoadTime(firstLoadTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["first_load_time"] = firstLoadTime })
}

// WithSchemaVersion schema_version获取 
func (obj *_V$planCachePlanStatMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithMergedVersion merged_version获取 
func (obj *_V$planCachePlanStatMgr) WithMergedVersion(mergedVersion int64) Option {
	return optionFunc(func(o *options) { o.query["merged_version"] = mergedVersion })
}

// WithLastActiveTime last_active_time获取 
func (obj *_V$planCachePlanStatMgr) WithLastActiveTime(lastActiveTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_active_time"] = lastActiveTime })
}

// WithAvgExeUsec avg_exe_usec获取 
func (obj *_V$planCachePlanStatMgr) WithAvgExeUsec(avgExeUsec int64) Option {
	return optionFunc(func(o *options) { o.query["avg_exe_usec"] = avgExeUsec })
}

// WithSlowestExeTime slowest_exe_time获取 
func (obj *_V$planCachePlanStatMgr) WithSlowestExeTime(slowestExeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["slowest_exe_time"] = slowestExeTime })
}

// WithSlowestExeUsec slowest_exe_usec获取 
func (obj *_V$planCachePlanStatMgr) WithSlowestExeUsec(slowestExeUsec int64) Option {
	return optionFunc(func(o *options) { o.query["slowest_exe_usec"] = slowestExeUsec })
}

// WithSlowCount slow_count获取 
func (obj *_V$planCachePlanStatMgr) WithSlowCount(slowCount int64) Option {
	return optionFunc(func(o *options) { o.query["slow_count"] = slowCount })
}

// WithHitCount hit_count获取 
func (obj *_V$planCachePlanStatMgr) WithHitCount(hitCount int64) Option {
	return optionFunc(func(o *options) { o.query["hit_count"] = hitCount })
}

// WithPlanSize plan_size获取 
func (obj *_V$planCachePlanStatMgr) WithPlanSize(planSize int64) Option {
	return optionFunc(func(o *options) { o.query["plan_size"] = planSize })
}

// WithExecutions executions获取 
func (obj *_V$planCachePlanStatMgr) WithExecutions(executions int64) Option {
	return optionFunc(func(o *options) { o.query["executions"] = executions })
}

// WithDiskReads disk_reads获取 
func (obj *_V$planCachePlanStatMgr) WithDiskReads(diskReads int64) Option {
	return optionFunc(func(o *options) { o.query["disk_reads"] = diskReads })
}

// WithDirectWrites direct_writes获取 
func (obj *_V$planCachePlanStatMgr) WithDirectWrites(directWrites int64) Option {
	return optionFunc(func(o *options) { o.query["direct_writes"] = directWrites })
}

// WithBufferGets buffer_gets获取 
func (obj *_V$planCachePlanStatMgr) WithBufferGets(bufferGets int64) Option {
	return optionFunc(func(o *options) { o.query["buffer_gets"] = bufferGets })
}

// WithApplicationWaitTime application_wait_time获取 
func (obj *_V$planCachePlanStatMgr) WithApplicationWaitTime(applicationWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["application_wait_time"] = applicationWaitTime })
}

// WithConcurrencyWaitTime concurrency_wait_time获取 
func (obj *_V$planCachePlanStatMgr) WithConcurrencyWaitTime(concurrencyWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["concurrency_wait_time"] = concurrencyWaitTime })
}

// WithUserIoWaitTime user_io_wait_time获取 
func (obj *_V$planCachePlanStatMgr) WithUserIoWaitTime(userIoWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["user_io_wait_time"] = userIoWaitTime })
}

// WithRowsProcessed rows_processed获取 
func (obj *_V$planCachePlanStatMgr) WithRowsProcessed(rowsProcessed int64) Option {
	return optionFunc(func(o *options) { o.query["rows_processed"] = rowsProcessed })
}

// WithElapsedTime elapsed_time获取 
func (obj *_V$planCachePlanStatMgr) WithElapsedTime(elapsedTime uint64) Option {
	return optionFunc(func(o *options) { o.query["elapsed_time"] = elapsedTime })
}

// WithCPUTime cpu_time获取 
func (obj *_V$planCachePlanStatMgr) WithCPUTime(cpuTime uint64) Option {
	return optionFunc(func(o *options) { o.query["cpu_time"] = cpuTime })
}

// WithLargeQuerys large_querys获取 
func (obj *_V$planCachePlanStatMgr) WithLargeQuerys(largeQuerys int64) Option {
	return optionFunc(func(o *options) { o.query["large_querys"] = largeQuerys })
}

// WithDelayedLargeQuerys delayed_large_querys获取 
func (obj *_V$planCachePlanStatMgr) WithDelayedLargeQuerys(delayedLargeQuerys int64) Option {
	return optionFunc(func(o *options) { o.query["delayed_large_querys"] = delayedLargeQuerys })
}

// WithDelayedPxQuerys delayed_px_querys获取 
func (obj *_V$planCachePlanStatMgr) WithDelayedPxQuerys(delayedPxQuerys int64) Option {
	return optionFunc(func(o *options) { o.query["delayed_px_querys"] = delayedPxQuerys })
}

// WithOutlineVersion outline_version获取 
func (obj *_V$planCachePlanStatMgr) WithOutlineVersion(outlineVersion int64) Option {
	return optionFunc(func(o *options) { o.query["outline_version"] = outlineVersion })
}

// WithOutlineID outline_id获取 
func (obj *_V$planCachePlanStatMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithOutlineData outline_data获取 
func (obj *_V$planCachePlanStatMgr) WithOutlineData(outlineData string) Option {
	return optionFunc(func(o *options) { o.query["outline_data"] = outlineData })
}

// WithAcsSelInfo acs_sel_info获取 
func (obj *_V$planCachePlanStatMgr) WithAcsSelInfo(acsSelInfo string) Option {
	return optionFunc(func(o *options) { o.query["acs_sel_info"] = acsSelInfo })
}

// WithTableScan table_scan获取 
func (obj *_V$planCachePlanStatMgr) WithTableScan(tableScan int8) Option {
	return optionFunc(func(o *options) { o.query["table_scan"] = tableScan })
}

// WithEvolution evolution获取 
func (obj *_V$planCachePlanStatMgr) WithEvolution(evolution int8) Option {
	return optionFunc(func(o *options) { o.query["evolution"] = evolution })
}

// WithEvoExecutions evo_executions获取 
func (obj *_V$planCachePlanStatMgr) WithEvoExecutions(evoExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["evo_executions"] = evoExecutions })
}

// WithEvoCPUTime evo_cpu_time获取 
func (obj *_V$planCachePlanStatMgr) WithEvoCPUTime(evoCPUTime uint64) Option {
	return optionFunc(func(o *options) { o.query["evo_cpu_time"] = evoCPUTime })
}

// WithTimeoutCount timeout_count获取 
func (obj *_V$planCachePlanStatMgr) WithTimeoutCount(timeoutCount int64) Option {
	return optionFunc(func(o *options) { o.query["timeout_count"] = timeoutCount })
}

// WithPsStmtID ps_stmt_id获取 
func (obj *_V$planCachePlanStatMgr) WithPsStmtID(psStmtID int64) Option {
	return optionFunc(func(o *options) { o.query["ps_stmt_id"] = psStmtID })
}

// WithSessid sessid获取 
func (obj *_V$planCachePlanStatMgr) WithSessid(sessid uint64) Option {
	return optionFunc(func(o *options) { o.query["sessid"] = sessid })
}

// WithTempTables temp_tables获取 
func (obj *_V$planCachePlanStatMgr) WithTempTables(tempTables string) Option {
	return optionFunc(func(o *options) { o.query["temp_tables"] = tempTables })
}

// WithIsUseJit is_use_jit获取 
func (obj *_V$planCachePlanStatMgr) WithIsUseJit(isUseJit int8) Option {
	return optionFunc(func(o *options) { o.query["is_use_jit"] = isUseJit })
}

// WithObjectType object_type获取 
func (obj *_V$planCachePlanStatMgr) WithObjectType(objectType string) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithHintsInfo hints_info获取 
func (obj *_V$planCachePlanStatMgr) WithHintsInfo(hintsInfo string) Option {
	return optionFunc(func(o *options) { o.query["hints_info"] = hintsInfo })
}

// WithHintsAllWorked hints_all_worked获取 
func (obj *_V$planCachePlanStatMgr) WithHintsAllWorked(hintsAllWorked int8) Option {
	return optionFunc(func(o *options) { o.query["hints_all_worked"] = hintsAllWorked })
}

// WithPlSchemaID pl_schema_id获取 
func (obj *_V$planCachePlanStatMgr) WithPlSchemaID(plSchemaID uint64) Option {
	return optionFunc(func(o *options) { o.query["pl_schema_id"] = plSchemaID })
}

// WithIsBatchedMultiStmt is_batched_multi_stmt获取 
func (obj *_V$planCachePlanStatMgr) WithIsBatchedMultiStmt(isBatchedMultiStmt int8) Option {
	return optionFunc(func(o *options) { o.query["is_batched_multi_stmt"] = isBatchedMultiStmt })
}


// GetByOption 功能选项模式获取
func (obj *_V$planCachePlanStatMgr) GetByOption(opts ...Option) (result V$planCachePlanStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$planCachePlanStatMgr) GetByOptions(opts ...Option) (results []*V$planCachePlanStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$planCachePlanStatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$planCachePlanStat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where(options.query)
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


// GetFromTenantID 通过tenant_id获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromTenantID(tenantID int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSvrIP(svrIP string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSvrPort(svrPort int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过plan_id获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromPlanID(planID int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`plan_id` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromPlanID(planIDs []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLID 通过sql_id获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSQLID(sqlID string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`sql_id` = ?", sqlID).Find(&results).Error
	
	return
}

// GetBatchFromSQLID 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSQLID(sqlIDs []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error
	
	return
}
 
// GetFromType 通过type获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromType(_type int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`type` = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromType(_types []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`type` IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromIsBindSensitive 通过is_bind_sensitive获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromIsBindSensitive(isBindSensitive int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_bind_sensitive` = ?", isBindSensitive).Find(&results).Error
	
	return
}

// GetBatchFromIsBindSensitive 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromIsBindSensitive(isBindSensitives []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_bind_sensitive` IN (?)", isBindSensitives).Find(&results).Error
	
	return
}
 
// GetFromIsBindAware 通过is_bind_aware获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromIsBindAware(isBindAware int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_bind_aware` = ?", isBindAware).Find(&results).Error
	
	return
}

// GetBatchFromIsBindAware 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromIsBindAware(isBindAwares []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_bind_aware` IN (?)", isBindAwares).Find(&results).Error
	
	return
}
 
// GetFromDbID 通过db_id获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromDbID(dbID uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`db_id` = ?", dbID).Find(&results).Error
	
	return
}

// GetBatchFromDbID 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromDbID(dbIDs []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`db_id` IN (?)", dbIDs).Find(&results).Error
	
	return
}
 
// GetFromStatement 通过statement获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromStatement(statement string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`statement` = ?", statement).Find(&results).Error
	
	return
}

// GetBatchFromStatement 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromStatement(statements []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`statement` IN (?)", statements).Find(&results).Error
	
	return
}
 
// GetFromQuerySQL 通过query_sql获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromQuerySQL(querySQL string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`query_sql` = ?", querySQL).Find(&results).Error
	
	return
}

// GetBatchFromQuerySQL 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromQuerySQL(querySQLs []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`query_sql` IN (?)", querySQLs).Find(&results).Error
	
	return
}
 
// GetFromSpecialParams 通过special_params获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSpecialParams(specialParams string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`special_params` = ?", specialParams).Find(&results).Error
	
	return
}

// GetBatchFromSpecialParams 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSpecialParams(specialParamss []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`special_params` IN (?)", specialParamss).Find(&results).Error
	
	return
}
 
// GetFromParamInfos 通过param_infos获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromParamInfos(paramInfos string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`param_infos` = ?", paramInfos).Find(&results).Error
	
	return
}

// GetBatchFromParamInfos 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromParamInfos(paramInfoss []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`param_infos` IN (?)", paramInfoss).Find(&results).Error
	
	return
}
 
// GetFromSysVars 通过sys_vars获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSysVars(sysVars string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`sys_vars` = ?", sysVars).Find(&results).Error
	
	return
}

// GetBatchFromSysVars 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSysVars(sysVarss []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`sys_vars` IN (?)", sysVarss).Find(&results).Error
	
	return
}
 
// GetFromPlanHash 通过plan_hash获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromPlanHash(planHash uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`plan_hash` = ?", planHash).Find(&results).Error
	
	return
}

// GetBatchFromPlanHash 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromPlanHash(planHashs []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`plan_hash` IN (?)", planHashs).Find(&results).Error
	
	return
}
 
// GetFromFirstLoadTime 通过first_load_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromFirstLoadTime(firstLoadTime time.Time) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`first_load_time` = ?", firstLoadTime).Find(&results).Error
	
	return
}

// GetBatchFromFirstLoadTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromFirstLoadTime(firstLoadTimes []time.Time) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`first_load_time` IN (?)", firstLoadTimes).Find(&results).Error
	
	return
}
 
// GetFromSchemaVersion 通过schema_version获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSchemaVersion(schemaVersion int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromSchemaVersion 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error
	
	return
}
 
// GetFromMergedVersion 通过merged_version获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromMergedVersion(mergedVersion int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`merged_version` = ?", mergedVersion).Find(&results).Error
	
	return
}

// GetBatchFromMergedVersion 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromMergedVersion(mergedVersions []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`merged_version` IN (?)", mergedVersions).Find(&results).Error
	
	return
}
 
// GetFromLastActiveTime 通过last_active_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromLastActiveTime(lastActiveTime time.Time) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`last_active_time` = ?", lastActiveTime).Find(&results).Error
	
	return
}

// GetBatchFromLastActiveTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromLastActiveTime(lastActiveTimes []time.Time) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`last_active_time` IN (?)", lastActiveTimes).Find(&results).Error
	
	return
}
 
// GetFromAvgExeUsec 通过avg_exe_usec获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromAvgExeUsec(avgExeUsec int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`avg_exe_usec` = ?", avgExeUsec).Find(&results).Error
	
	return
}

// GetBatchFromAvgExeUsec 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromAvgExeUsec(avgExeUsecs []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`avg_exe_usec` IN (?)", avgExeUsecs).Find(&results).Error
	
	return
}
 
// GetFromSlowestExeTime 通过slowest_exe_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSlowestExeTime(slowestExeTime time.Time) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`slowest_exe_time` = ?", slowestExeTime).Find(&results).Error
	
	return
}

// GetBatchFromSlowestExeTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSlowestExeTime(slowestExeTimes []time.Time) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`slowest_exe_time` IN (?)", slowestExeTimes).Find(&results).Error
	
	return
}
 
// GetFromSlowestExeUsec 通过slowest_exe_usec获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSlowestExeUsec(slowestExeUsec int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`slowest_exe_usec` = ?", slowestExeUsec).Find(&results).Error
	
	return
}

// GetBatchFromSlowestExeUsec 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSlowestExeUsec(slowestExeUsecs []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`slowest_exe_usec` IN (?)", slowestExeUsecs).Find(&results).Error
	
	return
}
 
// GetFromSlowCount 通过slow_count获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSlowCount(slowCount int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`slow_count` = ?", slowCount).Find(&results).Error
	
	return
}

// GetBatchFromSlowCount 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSlowCount(slowCounts []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`slow_count` IN (?)", slowCounts).Find(&results).Error
	
	return
}
 
// GetFromHitCount 通过hit_count获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromHitCount(hitCount int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`hit_count` = ?", hitCount).Find(&results).Error
	
	return
}

// GetBatchFromHitCount 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromHitCount(hitCounts []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`hit_count` IN (?)", hitCounts).Find(&results).Error
	
	return
}
 
// GetFromPlanSize 通过plan_size获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromPlanSize(planSize int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`plan_size` = ?", planSize).Find(&results).Error
	
	return
}

// GetBatchFromPlanSize 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromPlanSize(planSizes []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`plan_size` IN (?)", planSizes).Find(&results).Error
	
	return
}
 
// GetFromExecutions 通过executions获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromExecutions(executions int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`executions` = ?", executions).Find(&results).Error
	
	return
}

// GetBatchFromExecutions 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromExecutions(executionss []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`executions` IN (?)", executionss).Find(&results).Error
	
	return
}
 
// GetFromDiskReads 通过disk_reads获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromDiskReads(diskReads int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`disk_reads` = ?", diskReads).Find(&results).Error
	
	return
}

// GetBatchFromDiskReads 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromDiskReads(diskReadss []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`disk_reads` IN (?)", diskReadss).Find(&results).Error
	
	return
}
 
// GetFromDirectWrites 通过direct_writes获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromDirectWrites(directWrites int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`direct_writes` = ?", directWrites).Find(&results).Error
	
	return
}

// GetBatchFromDirectWrites 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromDirectWrites(directWritess []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`direct_writes` IN (?)", directWritess).Find(&results).Error
	
	return
}
 
// GetFromBufferGets 通过buffer_gets获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromBufferGets(bufferGets int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`buffer_gets` = ?", bufferGets).Find(&results).Error
	
	return
}

// GetBatchFromBufferGets 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromBufferGets(bufferGetss []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`buffer_gets` IN (?)", bufferGetss).Find(&results).Error
	
	return
}
 
// GetFromApplicationWaitTime 通过application_wait_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromApplicationWaitTime(applicationWaitTime uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`application_wait_time` = ?", applicationWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromApplicationWaitTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromApplicationWaitTime(applicationWaitTimes []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`application_wait_time` IN (?)", applicationWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromConcurrencyWaitTime 通过concurrency_wait_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromConcurrencyWaitTime(concurrencyWaitTime uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`concurrency_wait_time` = ?", concurrencyWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromConcurrencyWaitTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromConcurrencyWaitTime(concurrencyWaitTimes []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`concurrency_wait_time` IN (?)", concurrencyWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromUserIoWaitTime 通过user_io_wait_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromUserIoWaitTime(userIoWaitTime uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`user_io_wait_time` = ?", userIoWaitTime).Find(&results).Error
	
	return
}

// GetBatchFromUserIoWaitTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromUserIoWaitTime(userIoWaitTimes []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`user_io_wait_time` IN (?)", userIoWaitTimes).Find(&results).Error
	
	return
}
 
// GetFromRowsProcessed 通过rows_processed获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromRowsProcessed(rowsProcessed int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`rows_processed` = ?", rowsProcessed).Find(&results).Error
	
	return
}

// GetBatchFromRowsProcessed 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromRowsProcessed(rowsProcesseds []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`rows_processed` IN (?)", rowsProcesseds).Find(&results).Error
	
	return
}
 
// GetFromElapsedTime 通过elapsed_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromElapsedTime(elapsedTime uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`elapsed_time` = ?", elapsedTime).Find(&results).Error
	
	return
}

// GetBatchFromElapsedTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromElapsedTime(elapsedTimes []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`elapsed_time` IN (?)", elapsedTimes).Find(&results).Error
	
	return
}
 
// GetFromCPUTime 通过cpu_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromCPUTime(cpuTime uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`cpu_time` = ?", cpuTime).Find(&results).Error
	
	return
}

// GetBatchFromCPUTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromCPUTime(cpuTimes []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`cpu_time` IN (?)", cpuTimes).Find(&results).Error
	
	return
}
 
// GetFromLargeQuerys 通过large_querys获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromLargeQuerys(largeQuerys int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`large_querys` = ?", largeQuerys).Find(&results).Error
	
	return
}

// GetBatchFromLargeQuerys 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromLargeQuerys(largeQueryss []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`large_querys` IN (?)", largeQueryss).Find(&results).Error
	
	return
}
 
// GetFromDelayedLargeQuerys 通过delayed_large_querys获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromDelayedLargeQuerys(delayedLargeQuerys int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`delayed_large_querys` = ?", delayedLargeQuerys).Find(&results).Error
	
	return
}

// GetBatchFromDelayedLargeQuerys 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromDelayedLargeQuerys(delayedLargeQueryss []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`delayed_large_querys` IN (?)", delayedLargeQueryss).Find(&results).Error
	
	return
}
 
// GetFromDelayedPxQuerys 通过delayed_px_querys获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromDelayedPxQuerys(delayedPxQuerys int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`delayed_px_querys` = ?", delayedPxQuerys).Find(&results).Error
	
	return
}

// GetBatchFromDelayedPxQuerys 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromDelayedPxQuerys(delayedPxQueryss []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`delayed_px_querys` IN (?)", delayedPxQueryss).Find(&results).Error
	
	return
}
 
// GetFromOutlineVersion 通过outline_version获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromOutlineVersion(outlineVersion int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`outline_version` = ?", outlineVersion).Find(&results).Error
	
	return
}

// GetBatchFromOutlineVersion 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromOutlineVersion(outlineVersions []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`outline_version` IN (?)", outlineVersions).Find(&results).Error
	
	return
}
 
// GetFromOutlineID 通过outline_id获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromOutlineID(outlineID int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`outline_id` = ?", outlineID).Find(&results).Error
	
	return
}

// GetBatchFromOutlineID 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error
	
	return
}
 
// GetFromOutlineData 通过outline_data获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromOutlineData(outlineData string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`outline_data` = ?", outlineData).Find(&results).Error
	
	return
}

// GetBatchFromOutlineData 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromOutlineData(outlineDatas []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`outline_data` IN (?)", outlineDatas).Find(&results).Error
	
	return
}
 
// GetFromAcsSelInfo 通过acs_sel_info获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromAcsSelInfo(acsSelInfo string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`acs_sel_info` = ?", acsSelInfo).Find(&results).Error
	
	return
}

// GetBatchFromAcsSelInfo 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromAcsSelInfo(acsSelInfos []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`acs_sel_info` IN (?)", acsSelInfos).Find(&results).Error
	
	return
}
 
// GetFromTableScan 通过table_scan获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromTableScan(tableScan int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`table_scan` = ?", tableScan).Find(&results).Error
	
	return
}

// GetBatchFromTableScan 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromTableScan(tableScans []int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`table_scan` IN (?)", tableScans).Find(&results).Error
	
	return
}
 
// GetFromEvolution 通过evolution获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromEvolution(evolution int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`evolution` = ?", evolution).Find(&results).Error
	
	return
}

// GetBatchFromEvolution 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromEvolution(evolutions []int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`evolution` IN (?)", evolutions).Find(&results).Error
	
	return
}
 
// GetFromEvoExecutions 通过evo_executions获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromEvoExecutions(evoExecutions int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`evo_executions` = ?", evoExecutions).Find(&results).Error
	
	return
}

// GetBatchFromEvoExecutions 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromEvoExecutions(evoExecutionss []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`evo_executions` IN (?)", evoExecutionss).Find(&results).Error
	
	return
}
 
// GetFromEvoCPUTime 通过evo_cpu_time获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromEvoCPUTime(evoCPUTime uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`evo_cpu_time` = ?", evoCPUTime).Find(&results).Error
	
	return
}

// GetBatchFromEvoCPUTime 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromEvoCPUTime(evoCPUTimes []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`evo_cpu_time` IN (?)", evoCPUTimes).Find(&results).Error
	
	return
}
 
// GetFromTimeoutCount 通过timeout_count获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromTimeoutCount(timeoutCount int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`timeout_count` = ?", timeoutCount).Find(&results).Error
	
	return
}

// GetBatchFromTimeoutCount 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromTimeoutCount(timeoutCounts []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`timeout_count` IN (?)", timeoutCounts).Find(&results).Error
	
	return
}
 
// GetFromPsStmtID 通过ps_stmt_id获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromPsStmtID(psStmtID int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`ps_stmt_id` = ?", psStmtID).Find(&results).Error
	
	return
}

// GetBatchFromPsStmtID 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromPsStmtID(psStmtIDs []int64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`ps_stmt_id` IN (?)", psStmtIDs).Find(&results).Error
	
	return
}
 
// GetFromSessid 通过sessid获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromSessid(sessid uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`sessid` = ?", sessid).Find(&results).Error
	
	return
}

// GetBatchFromSessid 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromSessid(sessids []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`sessid` IN (?)", sessids).Find(&results).Error
	
	return
}
 
// GetFromTempTables 通过temp_tables获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromTempTables(tempTables string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`temp_tables` = ?", tempTables).Find(&results).Error
	
	return
}

// GetBatchFromTempTables 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromTempTables(tempTabless []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`temp_tables` IN (?)", tempTabless).Find(&results).Error
	
	return
}
 
// GetFromIsUseJit 通过is_use_jit获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromIsUseJit(isUseJit int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_use_jit` = ?", isUseJit).Find(&results).Error
	
	return
}

// GetBatchFromIsUseJit 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromIsUseJit(isUseJits []int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_use_jit` IN (?)", isUseJits).Find(&results).Error
	
	return
}
 
// GetFromObjectType 通过object_type获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromObjectType(objectType string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`object_type` = ?", objectType).Find(&results).Error
	
	return
}

// GetBatchFromObjectType 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromObjectType(objectTypes []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error
	
	return
}
 
// GetFromHintsInfo 通过hints_info获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromHintsInfo(hintsInfo string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`hints_info` = ?", hintsInfo).Find(&results).Error
	
	return
}

// GetBatchFromHintsInfo 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromHintsInfo(hintsInfos []string) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`hints_info` IN (?)", hintsInfos).Find(&results).Error
	
	return
}
 
// GetFromHintsAllWorked 通过hints_all_worked获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromHintsAllWorked(hintsAllWorked int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`hints_all_worked` = ?", hintsAllWorked).Find(&results).Error
	
	return
}

// GetBatchFromHintsAllWorked 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromHintsAllWorked(hintsAllWorkeds []int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`hints_all_worked` IN (?)", hintsAllWorkeds).Find(&results).Error
	
	return
}
 
// GetFromPlSchemaID 通过pl_schema_id获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromPlSchemaID(plSchemaID uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`pl_schema_id` = ?", plSchemaID).Find(&results).Error
	
	return
}

// GetBatchFromPlSchemaID 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromPlSchemaID(plSchemaIDs []uint64) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`pl_schema_id` IN (?)", plSchemaIDs).Find(&results).Error
	
	return
}
 
// GetFromIsBatchedMultiStmt 通过is_batched_multi_stmt获取内容  
func (obj *_V$planCachePlanStatMgr) GetFromIsBatchedMultiStmt(isBatchedMultiStmt int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_batched_multi_stmt` = ?", isBatchedMultiStmt).Find(&results).Error
	
	return
}

// GetBatchFromIsBatchedMultiStmt 批量查找 
func (obj *_V$planCachePlanStatMgr) GetBatchFromIsBatchedMultiStmt(isBatchedMultiStmts []int8) (results []*V$planCachePlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanStat{}).Where("`is_batched_multi_stmt` IN (?)", isBatchedMultiStmts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

