package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualPlanStatMgr struct {
	*_BaseMgr
}

// AllVirtualPlanStatMgr open func
func AllVirtualPlanStatMgr(db *gorm.DB) *_AllVirtualPlanStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPlanStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPlanStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_plan_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPlanStatMgr) GetTableName() string {
	return "__all_virtual_plan_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPlanStatMgr) Reset() *_AllVirtualPlanStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPlanStatMgr) Get() (result AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPlanStatMgr) Gets() (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPlanStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPlanStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPlanStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPlanStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPlanID plan_id获取
func (obj *_AllVirtualPlanStatMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithSQLID sql_id获取
func (obj *_AllVirtualPlanStatMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithType type获取
func (obj *_AllVirtualPlanStatMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIsBindSensitive is_bind_sensitive获取
func (obj *_AllVirtualPlanStatMgr) WithIsBindSensitive(isBindSensitive int64) Option {
	return optionFunc(func(o *options) { o.query["is_bind_sensitive"] = isBindSensitive })
}

// WithIsBindAware is_bind_aware获取
func (obj *_AllVirtualPlanStatMgr) WithIsBindAware(isBindAware int64) Option {
	return optionFunc(func(o *options) { o.query["is_bind_aware"] = isBindAware })
}

// WithStatement statement获取
func (obj *_AllVirtualPlanStatMgr) WithStatement(statement string) Option {
	return optionFunc(func(o *options) { o.query["statement"] = statement })
}

// WithQuerySQL query_sql获取
func (obj *_AllVirtualPlanStatMgr) WithQuerySQL(querySQL string) Option {
	return optionFunc(func(o *options) { o.query["query_sql"] = querySQL })
}

// WithSpecialParams special_params获取
func (obj *_AllVirtualPlanStatMgr) WithSpecialParams(specialParams string) Option {
	return optionFunc(func(o *options) { o.query["special_params"] = specialParams })
}

// WithParamInfos param_infos获取
func (obj *_AllVirtualPlanStatMgr) WithParamInfos(paramInfos string) Option {
	return optionFunc(func(o *options) { o.query["param_infos"] = paramInfos })
}

// WithSysVars sys_vars获取
func (obj *_AllVirtualPlanStatMgr) WithSysVars(sysVars string) Option {
	return optionFunc(func(o *options) { o.query["sys_vars"] = sysVars })
}

// WithPlanHash plan_hash获取
func (obj *_AllVirtualPlanStatMgr) WithPlanHash(planHash uint64) Option {
	return optionFunc(func(o *options) { o.query["plan_hash"] = planHash })
}

// WithFirstLoadTime first_load_time获取
func (obj *_AllVirtualPlanStatMgr) WithFirstLoadTime(firstLoadTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["first_load_time"] = firstLoadTime })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualPlanStatMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithMergedVersion merged_version获取
func (obj *_AllVirtualPlanStatMgr) WithMergedVersion(mergedVersion int64) Option {
	return optionFunc(func(o *options) { o.query["merged_version"] = mergedVersion })
}

// WithLastActiveTime last_active_time获取
func (obj *_AllVirtualPlanStatMgr) WithLastActiveTime(lastActiveTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_active_time"] = lastActiveTime })
}

// WithAvgExeUsec avg_exe_usec获取
func (obj *_AllVirtualPlanStatMgr) WithAvgExeUsec(avgExeUsec int64) Option {
	return optionFunc(func(o *options) { o.query["avg_exe_usec"] = avgExeUsec })
}

// WithSlowestExeTime slowest_exe_time获取
func (obj *_AllVirtualPlanStatMgr) WithSlowestExeTime(slowestExeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["slowest_exe_time"] = slowestExeTime })
}

// WithSlowestExeUsec slowest_exe_usec获取
func (obj *_AllVirtualPlanStatMgr) WithSlowestExeUsec(slowestExeUsec int64) Option {
	return optionFunc(func(o *options) { o.query["slowest_exe_usec"] = slowestExeUsec })
}

// WithSlowCount slow_count获取
func (obj *_AllVirtualPlanStatMgr) WithSlowCount(slowCount int64) Option {
	return optionFunc(func(o *options) { o.query["slow_count"] = slowCount })
}

// WithHitCount hit_count获取
func (obj *_AllVirtualPlanStatMgr) WithHitCount(hitCount int64) Option {
	return optionFunc(func(o *options) { o.query["hit_count"] = hitCount })
}

// WithPlanSize plan_size获取
func (obj *_AllVirtualPlanStatMgr) WithPlanSize(planSize int64) Option {
	return optionFunc(func(o *options) { o.query["plan_size"] = planSize })
}

// WithExecutions executions获取
func (obj *_AllVirtualPlanStatMgr) WithExecutions(executions int64) Option {
	return optionFunc(func(o *options) { o.query["executions"] = executions })
}

// WithDiskReads disk_reads获取
func (obj *_AllVirtualPlanStatMgr) WithDiskReads(diskReads int64) Option {
	return optionFunc(func(o *options) { o.query["disk_reads"] = diskReads })
}

// WithDirectWrites direct_writes获取
func (obj *_AllVirtualPlanStatMgr) WithDirectWrites(directWrites int64) Option {
	return optionFunc(func(o *options) { o.query["direct_writes"] = directWrites })
}

// WithBufferGets buffer_gets获取
func (obj *_AllVirtualPlanStatMgr) WithBufferGets(bufferGets int64) Option {
	return optionFunc(func(o *options) { o.query["buffer_gets"] = bufferGets })
}

// WithApplicationWaitTime application_wait_time获取
func (obj *_AllVirtualPlanStatMgr) WithApplicationWaitTime(applicationWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["application_wait_time"] = applicationWaitTime })
}

// WithConcurrencyWaitTime concurrency_wait_time获取
func (obj *_AllVirtualPlanStatMgr) WithConcurrencyWaitTime(concurrencyWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["concurrency_wait_time"] = concurrencyWaitTime })
}

// WithUserIoWaitTime user_io_wait_time获取
func (obj *_AllVirtualPlanStatMgr) WithUserIoWaitTime(userIoWaitTime uint64) Option {
	return optionFunc(func(o *options) { o.query["user_io_wait_time"] = userIoWaitTime })
}

// WithRowsProcessed rows_processed获取
func (obj *_AllVirtualPlanStatMgr) WithRowsProcessed(rowsProcessed int64) Option {
	return optionFunc(func(o *options) { o.query["rows_processed"] = rowsProcessed })
}

// WithElapsedTime elapsed_time获取
func (obj *_AllVirtualPlanStatMgr) WithElapsedTime(elapsedTime uint64) Option {
	return optionFunc(func(o *options) { o.query["elapsed_time"] = elapsedTime })
}

// WithCPUTime cpu_time获取
func (obj *_AllVirtualPlanStatMgr) WithCPUTime(cpuTime uint64) Option {
	return optionFunc(func(o *options) { o.query["cpu_time"] = cpuTime })
}

// WithLargeQuerys large_querys获取
func (obj *_AllVirtualPlanStatMgr) WithLargeQuerys(largeQuerys int64) Option {
	return optionFunc(func(o *options) { o.query["large_querys"] = largeQuerys })
}

// WithDelayedLargeQuerys delayed_large_querys获取
func (obj *_AllVirtualPlanStatMgr) WithDelayedLargeQuerys(delayedLargeQuerys int64) Option {
	return optionFunc(func(o *options) { o.query["delayed_large_querys"] = delayedLargeQuerys })
}

// WithOutlineVersion outline_version获取
func (obj *_AllVirtualPlanStatMgr) WithOutlineVersion(outlineVersion int64) Option {
	return optionFunc(func(o *options) { o.query["outline_version"] = outlineVersion })
}

// WithOutlineID outline_id获取
func (obj *_AllVirtualPlanStatMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithOutlineData outline_data获取
func (obj *_AllVirtualPlanStatMgr) WithOutlineData(outlineData string) Option {
	return optionFunc(func(o *options) { o.query["outline_data"] = outlineData })
}

// WithAcsSelInfo acs_sel_info获取
func (obj *_AllVirtualPlanStatMgr) WithAcsSelInfo(acsSelInfo string) Option {
	return optionFunc(func(o *options) { o.query["acs_sel_info"] = acsSelInfo })
}

// WithTableScan table_scan获取
func (obj *_AllVirtualPlanStatMgr) WithTableScan(tableScan int8) Option {
	return optionFunc(func(o *options) { o.query["table_scan"] = tableScan })
}

// WithDbID db_id获取
func (obj *_AllVirtualPlanStatMgr) WithDbID(dbID uint64) Option {
	return optionFunc(func(o *options) { o.query["db_id"] = dbID })
}

// WithEvolution evolution获取
func (obj *_AllVirtualPlanStatMgr) WithEvolution(evolution int8) Option {
	return optionFunc(func(o *options) { o.query["evolution"] = evolution })
}

// WithEvoExecutions evo_executions获取
func (obj *_AllVirtualPlanStatMgr) WithEvoExecutions(evoExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["evo_executions"] = evoExecutions })
}

// WithEvoCPUTime evo_cpu_time获取
func (obj *_AllVirtualPlanStatMgr) WithEvoCPUTime(evoCPUTime uint64) Option {
	return optionFunc(func(o *options) { o.query["evo_cpu_time"] = evoCPUTime })
}

// WithTimeoutCount timeout_count获取
func (obj *_AllVirtualPlanStatMgr) WithTimeoutCount(timeoutCount int64) Option {
	return optionFunc(func(o *options) { o.query["timeout_count"] = timeoutCount })
}

// WithPsStmtID ps_stmt_id获取
func (obj *_AllVirtualPlanStatMgr) WithPsStmtID(psStmtID int64) Option {
	return optionFunc(func(o *options) { o.query["ps_stmt_id"] = psStmtID })
}

// WithDelayedPxQuerys delayed_px_querys获取
func (obj *_AllVirtualPlanStatMgr) WithDelayedPxQuerys(delayedPxQuerys int64) Option {
	return optionFunc(func(o *options) { o.query["delayed_px_querys"] = delayedPxQuerys })
}

// WithSessid sessid获取
func (obj *_AllVirtualPlanStatMgr) WithSessid(sessid uint64) Option {
	return optionFunc(func(o *options) { o.query["sessid"] = sessid })
}

// WithTempTables temp_tables获取
func (obj *_AllVirtualPlanStatMgr) WithTempTables(tempTables string) Option {
	return optionFunc(func(o *options) { o.query["temp_tables"] = tempTables })
}

// WithIsUseJit is_use_jit获取
func (obj *_AllVirtualPlanStatMgr) WithIsUseJit(isUseJit int8) Option {
	return optionFunc(func(o *options) { o.query["is_use_jit"] = isUseJit })
}

// WithObjectType object_type获取
func (obj *_AllVirtualPlanStatMgr) WithObjectType(objectType string) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithEnableBfCache enable_bf_cache获取
func (obj *_AllVirtualPlanStatMgr) WithEnableBfCache(enableBfCache int8) Option {
	return optionFunc(func(o *options) { o.query["enable_bf_cache"] = enableBfCache })
}

// WithBfFilterCnt bf_filter_cnt获取
func (obj *_AllVirtualPlanStatMgr) WithBfFilterCnt(bfFilterCnt int64) Option {
	return optionFunc(func(o *options) { o.query["bf_filter_cnt"] = bfFilterCnt })
}

// WithBfAccessCnt bf_access_cnt获取
func (obj *_AllVirtualPlanStatMgr) WithBfAccessCnt(bfAccessCnt int64) Option {
	return optionFunc(func(o *options) { o.query["bf_access_cnt"] = bfAccessCnt })
}

// WithEnableRowCache enable_row_cache获取
func (obj *_AllVirtualPlanStatMgr) WithEnableRowCache(enableRowCache int8) Option {
	return optionFunc(func(o *options) { o.query["enable_row_cache"] = enableRowCache })
}

// WithRowCacheHitCnt row_cache_hit_cnt获取
func (obj *_AllVirtualPlanStatMgr) WithRowCacheHitCnt(rowCacheHitCnt int64) Option {
	return optionFunc(func(o *options) { o.query["row_cache_hit_cnt"] = rowCacheHitCnt })
}

// WithRowCacheMissCnt row_cache_miss_cnt获取
func (obj *_AllVirtualPlanStatMgr) WithRowCacheMissCnt(rowCacheMissCnt int64) Option {
	return optionFunc(func(o *options) { o.query["row_cache_miss_cnt"] = rowCacheMissCnt })
}

// WithEnableFuseRowCache enable_fuse_row_cache获取
func (obj *_AllVirtualPlanStatMgr) WithEnableFuseRowCache(enableFuseRowCache int8) Option {
	return optionFunc(func(o *options) { o.query["enable_fuse_row_cache"] = enableFuseRowCache })
}

// WithFuseRowCacheHitCnt fuse_row_cache_hit_cnt获取
func (obj *_AllVirtualPlanStatMgr) WithFuseRowCacheHitCnt(fuseRowCacheHitCnt int64) Option {
	return optionFunc(func(o *options) { o.query["fuse_row_cache_hit_cnt"] = fuseRowCacheHitCnt })
}

// WithFuseRowCacheMissCnt fuse_row_cache_miss_cnt获取
func (obj *_AllVirtualPlanStatMgr) WithFuseRowCacheMissCnt(fuseRowCacheMissCnt int64) Option {
	return optionFunc(func(o *options) { o.query["fuse_row_cache_miss_cnt"] = fuseRowCacheMissCnt })
}

// WithHintsInfo hints_info获取
func (obj *_AllVirtualPlanStatMgr) WithHintsInfo(hintsInfo string) Option {
	return optionFunc(func(o *options) { o.query["hints_info"] = hintsInfo })
}

// WithHintsAllWorked hints_all_worked获取
func (obj *_AllVirtualPlanStatMgr) WithHintsAllWorked(hintsAllWorked int8) Option {
	return optionFunc(func(o *options) { o.query["hints_all_worked"] = hintsAllWorked })
}

// WithPlSchemaID pl_schema_id获取
func (obj *_AllVirtualPlanStatMgr) WithPlSchemaID(plSchemaID uint64) Option {
	return optionFunc(func(o *options) { o.query["pl_schema_id"] = plSchemaID })
}

// WithIsBatchedMultiStmt is_batched_multi_stmt获取
func (obj *_AllVirtualPlanStatMgr) WithIsBatchedMultiStmt(isBatchedMultiStmt int8) Option {
	return optionFunc(func(o *options) { o.query["is_batched_multi_stmt"] = isBatchedMultiStmt })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPlanStatMgr) GetByOption(opts ...Option) (result AllVirtualPlanStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPlanStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualPlanStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPlanStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPlanStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where(options.query)
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
func (obj *_AllVirtualPlanStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPlanID 通过plan_id获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromPlanID(planID int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`plan_id` = ?", planID).Find(&results).Error

	return
}

// GetBatchFromPlanID 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromPlanID(planIDs []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSQLID(sqlID string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSQLID(sqlIDs []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromType(_type int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromType(_types []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromIsBindSensitive 通过is_bind_sensitive获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromIsBindSensitive(isBindSensitive int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_bind_sensitive` = ?", isBindSensitive).Find(&results).Error

	return
}

// GetBatchFromIsBindSensitive 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromIsBindSensitive(isBindSensitives []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_bind_sensitive` IN (?)", isBindSensitives).Find(&results).Error

	return
}

// GetFromIsBindAware 通过is_bind_aware获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromIsBindAware(isBindAware int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_bind_aware` = ?", isBindAware).Find(&results).Error

	return
}

// GetBatchFromIsBindAware 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromIsBindAware(isBindAwares []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_bind_aware` IN (?)", isBindAwares).Find(&results).Error

	return
}

// GetFromStatement 通过statement获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromStatement(statement string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`statement` = ?", statement).Find(&results).Error

	return
}

// GetBatchFromStatement 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromStatement(statements []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`statement` IN (?)", statements).Find(&results).Error

	return
}

// GetFromQuerySQL 通过query_sql获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromQuerySQL(querySQL string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`query_sql` = ?", querySQL).Find(&results).Error

	return
}

// GetBatchFromQuerySQL 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromQuerySQL(querySQLs []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`query_sql` IN (?)", querySQLs).Find(&results).Error

	return
}

// GetFromSpecialParams 通过special_params获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSpecialParams(specialParams string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`special_params` = ?", specialParams).Find(&results).Error

	return
}

// GetBatchFromSpecialParams 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSpecialParams(specialParamss []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`special_params` IN (?)", specialParamss).Find(&results).Error

	return
}

// GetFromParamInfos 通过param_infos获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromParamInfos(paramInfos string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`param_infos` = ?", paramInfos).Find(&results).Error

	return
}

// GetBatchFromParamInfos 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromParamInfos(paramInfoss []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`param_infos` IN (?)", paramInfoss).Find(&results).Error

	return
}

// GetFromSysVars 通过sys_vars获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSysVars(sysVars string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`sys_vars` = ?", sysVars).Find(&results).Error

	return
}

// GetBatchFromSysVars 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSysVars(sysVarss []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`sys_vars` IN (?)", sysVarss).Find(&results).Error

	return
}

// GetFromPlanHash 通过plan_hash获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromPlanHash(planHash uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`plan_hash` = ?", planHash).Find(&results).Error

	return
}

// GetBatchFromPlanHash 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromPlanHash(planHashs []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`plan_hash` IN (?)", planHashs).Find(&results).Error

	return
}

// GetFromFirstLoadTime 通过first_load_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromFirstLoadTime(firstLoadTime time.Time) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`first_load_time` = ?", firstLoadTime).Find(&results).Error

	return
}

// GetBatchFromFirstLoadTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromFirstLoadTime(firstLoadTimes []time.Time) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`first_load_time` IN (?)", firstLoadTimes).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromMergedVersion 通过merged_version获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromMergedVersion(mergedVersion int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`merged_version` = ?", mergedVersion).Find(&results).Error

	return
}

// GetBatchFromMergedVersion 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromMergedVersion(mergedVersions []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`merged_version` IN (?)", mergedVersions).Find(&results).Error

	return
}

// GetFromLastActiveTime 通过last_active_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromLastActiveTime(lastActiveTime time.Time) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`last_active_time` = ?", lastActiveTime).Find(&results).Error

	return
}

// GetBatchFromLastActiveTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromLastActiveTime(lastActiveTimes []time.Time) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`last_active_time` IN (?)", lastActiveTimes).Find(&results).Error

	return
}

// GetFromAvgExeUsec 通过avg_exe_usec获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromAvgExeUsec(avgExeUsec int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`avg_exe_usec` = ?", avgExeUsec).Find(&results).Error

	return
}

// GetBatchFromAvgExeUsec 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromAvgExeUsec(avgExeUsecs []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`avg_exe_usec` IN (?)", avgExeUsecs).Find(&results).Error

	return
}

// GetFromSlowestExeTime 通过slowest_exe_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSlowestExeTime(slowestExeTime time.Time) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`slowest_exe_time` = ?", slowestExeTime).Find(&results).Error

	return
}

// GetBatchFromSlowestExeTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSlowestExeTime(slowestExeTimes []time.Time) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`slowest_exe_time` IN (?)", slowestExeTimes).Find(&results).Error

	return
}

// GetFromSlowestExeUsec 通过slowest_exe_usec获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSlowestExeUsec(slowestExeUsec int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`slowest_exe_usec` = ?", slowestExeUsec).Find(&results).Error

	return
}

// GetBatchFromSlowestExeUsec 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSlowestExeUsec(slowestExeUsecs []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`slowest_exe_usec` IN (?)", slowestExeUsecs).Find(&results).Error

	return
}

// GetFromSlowCount 通过slow_count获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSlowCount(slowCount int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`slow_count` = ?", slowCount).Find(&results).Error

	return
}

// GetBatchFromSlowCount 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSlowCount(slowCounts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`slow_count` IN (?)", slowCounts).Find(&results).Error

	return
}

// GetFromHitCount 通过hit_count获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromHitCount(hitCount int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`hit_count` = ?", hitCount).Find(&results).Error

	return
}

// GetBatchFromHitCount 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromHitCount(hitCounts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`hit_count` IN (?)", hitCounts).Find(&results).Error

	return
}

// GetFromPlanSize 通过plan_size获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromPlanSize(planSize int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`plan_size` = ?", planSize).Find(&results).Error

	return
}

// GetBatchFromPlanSize 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromPlanSize(planSizes []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`plan_size` IN (?)", planSizes).Find(&results).Error

	return
}

// GetFromExecutions 通过executions获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromExecutions(executions int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`executions` = ?", executions).Find(&results).Error

	return
}

// GetBatchFromExecutions 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromExecutions(executionss []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`executions` IN (?)", executionss).Find(&results).Error

	return
}

// GetFromDiskReads 通过disk_reads获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromDiskReads(diskReads int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`disk_reads` = ?", diskReads).Find(&results).Error

	return
}

// GetBatchFromDiskReads 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromDiskReads(diskReadss []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`disk_reads` IN (?)", diskReadss).Find(&results).Error

	return
}

// GetFromDirectWrites 通过direct_writes获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromDirectWrites(directWrites int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`direct_writes` = ?", directWrites).Find(&results).Error

	return
}

// GetBatchFromDirectWrites 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromDirectWrites(directWritess []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`direct_writes` IN (?)", directWritess).Find(&results).Error

	return
}

// GetFromBufferGets 通过buffer_gets获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromBufferGets(bufferGets int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`buffer_gets` = ?", bufferGets).Find(&results).Error

	return
}

// GetBatchFromBufferGets 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromBufferGets(bufferGetss []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`buffer_gets` IN (?)", bufferGetss).Find(&results).Error

	return
}

// GetFromApplicationWaitTime 通过application_wait_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromApplicationWaitTime(applicationWaitTime uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`application_wait_time` = ?", applicationWaitTime).Find(&results).Error

	return
}

// GetBatchFromApplicationWaitTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromApplicationWaitTime(applicationWaitTimes []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`application_wait_time` IN (?)", applicationWaitTimes).Find(&results).Error

	return
}

// GetFromConcurrencyWaitTime 通过concurrency_wait_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromConcurrencyWaitTime(concurrencyWaitTime uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`concurrency_wait_time` = ?", concurrencyWaitTime).Find(&results).Error

	return
}

// GetBatchFromConcurrencyWaitTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromConcurrencyWaitTime(concurrencyWaitTimes []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`concurrency_wait_time` IN (?)", concurrencyWaitTimes).Find(&results).Error

	return
}

// GetFromUserIoWaitTime 通过user_io_wait_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromUserIoWaitTime(userIoWaitTime uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`user_io_wait_time` = ?", userIoWaitTime).Find(&results).Error

	return
}

// GetBatchFromUserIoWaitTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromUserIoWaitTime(userIoWaitTimes []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`user_io_wait_time` IN (?)", userIoWaitTimes).Find(&results).Error

	return
}

// GetFromRowsProcessed 通过rows_processed获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromRowsProcessed(rowsProcessed int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`rows_processed` = ?", rowsProcessed).Find(&results).Error

	return
}

// GetBatchFromRowsProcessed 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromRowsProcessed(rowsProcesseds []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`rows_processed` IN (?)", rowsProcesseds).Find(&results).Error

	return
}

// GetFromElapsedTime 通过elapsed_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromElapsedTime(elapsedTime uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`elapsed_time` = ?", elapsedTime).Find(&results).Error

	return
}

// GetBatchFromElapsedTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromElapsedTime(elapsedTimes []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`elapsed_time` IN (?)", elapsedTimes).Find(&results).Error

	return
}

// GetFromCPUTime 通过cpu_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromCPUTime(cpuTime uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`cpu_time` = ?", cpuTime).Find(&results).Error

	return
}

// GetBatchFromCPUTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromCPUTime(cpuTimes []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`cpu_time` IN (?)", cpuTimes).Find(&results).Error

	return
}

// GetFromLargeQuerys 通过large_querys获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromLargeQuerys(largeQuerys int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`large_querys` = ?", largeQuerys).Find(&results).Error

	return
}

// GetBatchFromLargeQuerys 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromLargeQuerys(largeQueryss []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`large_querys` IN (?)", largeQueryss).Find(&results).Error

	return
}

// GetFromDelayedLargeQuerys 通过delayed_large_querys获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromDelayedLargeQuerys(delayedLargeQuerys int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`delayed_large_querys` = ?", delayedLargeQuerys).Find(&results).Error

	return
}

// GetBatchFromDelayedLargeQuerys 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromDelayedLargeQuerys(delayedLargeQueryss []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`delayed_large_querys` IN (?)", delayedLargeQueryss).Find(&results).Error

	return
}

// GetFromOutlineVersion 通过outline_version获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromOutlineVersion(outlineVersion int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`outline_version` = ?", outlineVersion).Find(&results).Error

	return
}

// GetBatchFromOutlineVersion 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromOutlineVersion(outlineVersions []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`outline_version` IN (?)", outlineVersions).Find(&results).Error

	return
}

// GetFromOutlineID 通过outline_id获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromOutlineID(outlineID int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`outline_id` = ?", outlineID).Find(&results).Error

	return
}

// GetBatchFromOutlineID 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error

	return
}

// GetFromOutlineData 通过outline_data获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromOutlineData(outlineData string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`outline_data` = ?", outlineData).Find(&results).Error

	return
}

// GetBatchFromOutlineData 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromOutlineData(outlineDatas []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`outline_data` IN (?)", outlineDatas).Find(&results).Error

	return
}

// GetFromAcsSelInfo 通过acs_sel_info获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromAcsSelInfo(acsSelInfo string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`acs_sel_info` = ?", acsSelInfo).Find(&results).Error

	return
}

// GetBatchFromAcsSelInfo 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromAcsSelInfo(acsSelInfos []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`acs_sel_info` IN (?)", acsSelInfos).Find(&results).Error

	return
}

// GetFromTableScan 通过table_scan获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromTableScan(tableScan int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`table_scan` = ?", tableScan).Find(&results).Error

	return
}

// GetBatchFromTableScan 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromTableScan(tableScans []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`table_scan` IN (?)", tableScans).Find(&results).Error

	return
}

// GetFromDbID 通过db_id获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromDbID(dbID uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`db_id` = ?", dbID).Find(&results).Error

	return
}

// GetBatchFromDbID 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromDbID(dbIDs []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`db_id` IN (?)", dbIDs).Find(&results).Error

	return
}

// GetFromEvolution 通过evolution获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromEvolution(evolution int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`evolution` = ?", evolution).Find(&results).Error

	return
}

// GetBatchFromEvolution 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromEvolution(evolutions []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`evolution` IN (?)", evolutions).Find(&results).Error

	return
}

// GetFromEvoExecutions 通过evo_executions获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromEvoExecutions(evoExecutions int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`evo_executions` = ?", evoExecutions).Find(&results).Error

	return
}

// GetBatchFromEvoExecutions 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromEvoExecutions(evoExecutionss []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`evo_executions` IN (?)", evoExecutionss).Find(&results).Error

	return
}

// GetFromEvoCPUTime 通过evo_cpu_time获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromEvoCPUTime(evoCPUTime uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`evo_cpu_time` = ?", evoCPUTime).Find(&results).Error

	return
}

// GetBatchFromEvoCPUTime 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromEvoCPUTime(evoCPUTimes []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`evo_cpu_time` IN (?)", evoCPUTimes).Find(&results).Error

	return
}

// GetFromTimeoutCount 通过timeout_count获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromTimeoutCount(timeoutCount int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`timeout_count` = ?", timeoutCount).Find(&results).Error

	return
}

// GetBatchFromTimeoutCount 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromTimeoutCount(timeoutCounts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`timeout_count` IN (?)", timeoutCounts).Find(&results).Error

	return
}

// GetFromPsStmtID 通过ps_stmt_id获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromPsStmtID(psStmtID int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`ps_stmt_id` = ?", psStmtID).Find(&results).Error

	return
}

// GetBatchFromPsStmtID 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromPsStmtID(psStmtIDs []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`ps_stmt_id` IN (?)", psStmtIDs).Find(&results).Error

	return
}

// GetFromDelayedPxQuerys 通过delayed_px_querys获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromDelayedPxQuerys(delayedPxQuerys int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`delayed_px_querys` = ?", delayedPxQuerys).Find(&results).Error

	return
}

// GetBatchFromDelayedPxQuerys 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromDelayedPxQuerys(delayedPxQueryss []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`delayed_px_querys` IN (?)", delayedPxQueryss).Find(&results).Error

	return
}

// GetFromSessid 通过sessid获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromSessid(sessid uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`sessid` = ?", sessid).Find(&results).Error

	return
}

// GetBatchFromSessid 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromSessid(sessids []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`sessid` IN (?)", sessids).Find(&results).Error

	return
}

// GetFromTempTables 通过temp_tables获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromTempTables(tempTables string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`temp_tables` = ?", tempTables).Find(&results).Error

	return
}

// GetBatchFromTempTables 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromTempTables(tempTabless []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`temp_tables` IN (?)", tempTabless).Find(&results).Error

	return
}

// GetFromIsUseJit 通过is_use_jit获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromIsUseJit(isUseJit int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_use_jit` = ?", isUseJit).Find(&results).Error

	return
}

// GetBatchFromIsUseJit 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromIsUseJit(isUseJits []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_use_jit` IN (?)", isUseJits).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromObjectType(objectType string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromObjectType(objectTypes []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromEnableBfCache 通过enable_bf_cache获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromEnableBfCache(enableBfCache int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`enable_bf_cache` = ?", enableBfCache).Find(&results).Error

	return
}

// GetBatchFromEnableBfCache 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromEnableBfCache(enableBfCaches []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`enable_bf_cache` IN (?)", enableBfCaches).Find(&results).Error

	return
}

// GetFromBfFilterCnt 通过bf_filter_cnt获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromBfFilterCnt(bfFilterCnt int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`bf_filter_cnt` = ?", bfFilterCnt).Find(&results).Error

	return
}

// GetBatchFromBfFilterCnt 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromBfFilterCnt(bfFilterCnts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`bf_filter_cnt` IN (?)", bfFilterCnts).Find(&results).Error

	return
}

// GetFromBfAccessCnt 通过bf_access_cnt获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromBfAccessCnt(bfAccessCnt int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`bf_access_cnt` = ?", bfAccessCnt).Find(&results).Error

	return
}

// GetBatchFromBfAccessCnt 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromBfAccessCnt(bfAccessCnts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`bf_access_cnt` IN (?)", bfAccessCnts).Find(&results).Error

	return
}

// GetFromEnableRowCache 通过enable_row_cache获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromEnableRowCache(enableRowCache int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`enable_row_cache` = ?", enableRowCache).Find(&results).Error

	return
}

// GetBatchFromEnableRowCache 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromEnableRowCache(enableRowCaches []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`enable_row_cache` IN (?)", enableRowCaches).Find(&results).Error

	return
}

// GetFromRowCacheHitCnt 通过row_cache_hit_cnt获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromRowCacheHitCnt(rowCacheHitCnt int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`row_cache_hit_cnt` = ?", rowCacheHitCnt).Find(&results).Error

	return
}

// GetBatchFromRowCacheHitCnt 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromRowCacheHitCnt(rowCacheHitCnts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`row_cache_hit_cnt` IN (?)", rowCacheHitCnts).Find(&results).Error

	return
}

// GetFromRowCacheMissCnt 通过row_cache_miss_cnt获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromRowCacheMissCnt(rowCacheMissCnt int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`row_cache_miss_cnt` = ?", rowCacheMissCnt).Find(&results).Error

	return
}

// GetBatchFromRowCacheMissCnt 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromRowCacheMissCnt(rowCacheMissCnts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`row_cache_miss_cnt` IN (?)", rowCacheMissCnts).Find(&results).Error

	return
}

// GetFromEnableFuseRowCache 通过enable_fuse_row_cache获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromEnableFuseRowCache(enableFuseRowCache int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`enable_fuse_row_cache` = ?", enableFuseRowCache).Find(&results).Error

	return
}

// GetBatchFromEnableFuseRowCache 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromEnableFuseRowCache(enableFuseRowCaches []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`enable_fuse_row_cache` IN (?)", enableFuseRowCaches).Find(&results).Error

	return
}

// GetFromFuseRowCacheHitCnt 通过fuse_row_cache_hit_cnt获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromFuseRowCacheHitCnt(fuseRowCacheHitCnt int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`fuse_row_cache_hit_cnt` = ?", fuseRowCacheHitCnt).Find(&results).Error

	return
}

// GetBatchFromFuseRowCacheHitCnt 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromFuseRowCacheHitCnt(fuseRowCacheHitCnts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`fuse_row_cache_hit_cnt` IN (?)", fuseRowCacheHitCnts).Find(&results).Error

	return
}

// GetFromFuseRowCacheMissCnt 通过fuse_row_cache_miss_cnt获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromFuseRowCacheMissCnt(fuseRowCacheMissCnt int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`fuse_row_cache_miss_cnt` = ?", fuseRowCacheMissCnt).Find(&results).Error

	return
}

// GetBatchFromFuseRowCacheMissCnt 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromFuseRowCacheMissCnt(fuseRowCacheMissCnts []int64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`fuse_row_cache_miss_cnt` IN (?)", fuseRowCacheMissCnts).Find(&results).Error

	return
}

// GetFromHintsInfo 通过hints_info获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromHintsInfo(hintsInfo string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`hints_info` = ?", hintsInfo).Find(&results).Error

	return
}

// GetBatchFromHintsInfo 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromHintsInfo(hintsInfos []string) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`hints_info` IN (?)", hintsInfos).Find(&results).Error

	return
}

// GetFromHintsAllWorked 通过hints_all_worked获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromHintsAllWorked(hintsAllWorked int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`hints_all_worked` = ?", hintsAllWorked).Find(&results).Error

	return
}

// GetBatchFromHintsAllWorked 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromHintsAllWorked(hintsAllWorkeds []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`hints_all_worked` IN (?)", hintsAllWorkeds).Find(&results).Error

	return
}

// GetFromPlSchemaID 通过pl_schema_id获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromPlSchemaID(plSchemaID uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`pl_schema_id` = ?", plSchemaID).Find(&results).Error

	return
}

// GetBatchFromPlSchemaID 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromPlSchemaID(plSchemaIDs []uint64) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`pl_schema_id` IN (?)", plSchemaIDs).Find(&results).Error

	return
}

// GetFromIsBatchedMultiStmt 通过is_batched_multi_stmt获取内容
func (obj *_AllVirtualPlanStatMgr) GetFromIsBatchedMultiStmt(isBatchedMultiStmt int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_batched_multi_stmt` = ?", isBatchedMultiStmt).Find(&results).Error

	return
}

// GetBatchFromIsBatchedMultiStmt 批量查找
func (obj *_AllVirtualPlanStatMgr) GetBatchFromIsBatchedMultiStmt(isBatchedMultiStmts []int8) (results []*AllVirtualPlanStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanStat{}).Where("`is_batched_multi_stmt` IN (?)", isBatchedMultiStmts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
