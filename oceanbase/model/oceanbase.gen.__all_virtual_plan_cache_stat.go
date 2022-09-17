package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPlanCacheStatMgr struct {
	*_BaseMgr
}

// AllVirtualPlanCacheStatMgr open func
func AllVirtualPlanCacheStatMgr(db *gorm.DB) *_AllVirtualPlanCacheStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPlanCacheStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPlanCacheStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_plan_cache_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPlanCacheStatMgr) GetTableName() string {
	return "__all_virtual_plan_cache_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPlanCacheStatMgr) Reset() *_AllVirtualPlanCacheStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPlanCacheStatMgr) Get() (result AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPlanCacheStatMgr) Gets() (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPlanCacheStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPlanCacheStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPlanCacheStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPlanCacheStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLNum sql_num获取
func (obj *_AllVirtualPlanCacheStatMgr) WithSQLNum(sqlNum int64) Option {
	return optionFunc(func(o *options) { o.query["sql_num"] = sqlNum })
}

// WithMemUsed mem_used获取
func (obj *_AllVirtualPlanCacheStatMgr) WithMemUsed(memUsed int64) Option {
	return optionFunc(func(o *options) { o.query["mem_used"] = memUsed })
}

// WithMemHold mem_hold获取
func (obj *_AllVirtualPlanCacheStatMgr) WithMemHold(memHold int64) Option {
	return optionFunc(func(o *options) { o.query["mem_hold"] = memHold })
}

// WithAccessCount access_count获取
func (obj *_AllVirtualPlanCacheStatMgr) WithAccessCount(accessCount int64) Option {
	return optionFunc(func(o *options) { o.query["access_count"] = accessCount })
}

// WithHitCount hit_count获取
func (obj *_AllVirtualPlanCacheStatMgr) WithHitCount(hitCount int64) Option {
	return optionFunc(func(o *options) { o.query["hit_count"] = hitCount })
}

// WithHitRate hit_rate获取
func (obj *_AllVirtualPlanCacheStatMgr) WithHitRate(hitRate int64) Option {
	return optionFunc(func(o *options) { o.query["hit_rate"] = hitRate })
}

// WithPlanNum plan_num获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPlanNum(planNum int64) Option {
	return optionFunc(func(o *options) { o.query["plan_num"] = planNum })
}

// WithMemLimit mem_limit获取
func (obj *_AllVirtualPlanCacheStatMgr) WithMemLimit(memLimit int64) Option {
	return optionFunc(func(o *options) { o.query["mem_limit"] = memLimit })
}

// WithHashBucket hash_bucket获取
func (obj *_AllVirtualPlanCacheStatMgr) WithHashBucket(hashBucket int64) Option {
	return optionFunc(func(o *options) { o.query["hash_bucket"] = hashBucket })
}

// WithStmtkeyNum stmtkey_num获取
func (obj *_AllVirtualPlanCacheStatMgr) WithStmtkeyNum(stmtkeyNum int64) Option {
	return optionFunc(func(o *options) { o.query["stmtkey_num"] = stmtkeyNum })
}

// WithPcRefPlanLocal pc_ref_plan_local获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcRefPlanLocal(pcRefPlanLocal int64) Option {
	return optionFunc(func(o *options) { o.query["pc_ref_plan_local"] = pcRefPlanLocal })
}

// WithPcRefPlanRemote pc_ref_plan_remote获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcRefPlanRemote(pcRefPlanRemote int64) Option {
	return optionFunc(func(o *options) { o.query["pc_ref_plan_remote"] = pcRefPlanRemote })
}

// WithPcRefPlanDist pc_ref_plan_dist获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcRefPlanDist(pcRefPlanDist int64) Option {
	return optionFunc(func(o *options) { o.query["pc_ref_plan_dist"] = pcRefPlanDist })
}

// WithPcRefPlanArr pc_ref_plan_arr获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcRefPlanArr(pcRefPlanArr int64) Option {
	return optionFunc(func(o *options) { o.query["pc_ref_plan_arr"] = pcRefPlanArr })
}

// WithPcRefPlanStat pc_ref_plan_stat获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcRefPlanStat(pcRefPlanStat int64) Option {
	return optionFunc(func(o *options) { o.query["pc_ref_plan_stat"] = pcRefPlanStat })
}

// WithPcRefPl pc_ref_pl获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcRefPl(pcRefPl int64) Option {
	return optionFunc(func(o *options) { o.query["pc_ref_pl"] = pcRefPl })
}

// WithPcRefPlStat pc_ref_pl_stat获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcRefPlStat(pcRefPlStat int64) Option {
	return optionFunc(func(o *options) { o.query["pc_ref_pl_stat"] = pcRefPlStat })
}

// WithPlanGen plan_gen获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPlanGen(planGen int64) Option {
	return optionFunc(func(o *options) { o.query["plan_gen"] = planGen })
}

// WithCliQuery cli_query获取
func (obj *_AllVirtualPlanCacheStatMgr) WithCliQuery(cliQuery int64) Option {
	return optionFunc(func(o *options) { o.query["cli_query"] = cliQuery })
}

// WithOutlineExec outline_exec获取
func (obj *_AllVirtualPlanCacheStatMgr) WithOutlineExec(outlineExec int64) Option {
	return optionFunc(func(o *options) { o.query["outline_exec"] = outlineExec })
}

// WithPlanExplain plan_explain获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPlanExplain(planExplain int64) Option {
	return optionFunc(func(o *options) { o.query["plan_explain"] = planExplain })
}

// WithAsynBaseline asyn_baseline获取
func (obj *_AllVirtualPlanCacheStatMgr) WithAsynBaseline(asynBaseline int64) Option {
	return optionFunc(func(o *options) { o.query["asyn_baseline"] = asynBaseline })
}

// WithLoadBaseline load_baseline获取
func (obj *_AllVirtualPlanCacheStatMgr) WithLoadBaseline(loadBaseline int64) Option {
	return optionFunc(func(o *options) { o.query["load_baseline"] = loadBaseline })
}

// WithPsExec ps_exec获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPsExec(psExec int64) Option {
	return optionFunc(func(o *options) { o.query["ps_exec"] = psExec })
}

// WithGvSQL gv_sql获取
func (obj *_AllVirtualPlanCacheStatMgr) WithGvSQL(gvSQL int64) Option {
	return optionFunc(func(o *options) { o.query["gv_sql"] = gvSQL })
}

// WithPlAnon pl_anon获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPlAnon(plAnon int64) Option {
	return optionFunc(func(o *options) { o.query["pl_anon"] = plAnon })
}

// WithPlRoutine pl_routine获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPlRoutine(plRoutine int64) Option {
	return optionFunc(func(o *options) { o.query["pl_routine"] = plRoutine })
}

// WithPackageVar package_var获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPackageVar(packageVar int64) Option {
	return optionFunc(func(o *options) { o.query["package_var"] = packageVar })
}

// WithPackageType package_type获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPackageType(packageType int64) Option {
	return optionFunc(func(o *options) { o.query["package_type"] = packageType })
}

// WithPackageSpec package_spec获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPackageSpec(packageSpec int64) Option {
	return optionFunc(func(o *options) { o.query["package_spec"] = packageSpec })
}

// WithPackageBody package_body获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPackageBody(packageBody int64) Option {
	return optionFunc(func(o *options) { o.query["package_body"] = packageBody })
}

// WithPackageResv package_resv获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPackageResv(packageResv int64) Option {
	return optionFunc(func(o *options) { o.query["package_resv"] = packageResv })
}

// WithGetPkg get_pkg获取
func (obj *_AllVirtualPlanCacheStatMgr) WithGetPkg(getPkg int64) Option {
	return optionFunc(func(o *options) { o.query["get_pkg"] = getPkg })
}

// WithIndexBuilder index_builder获取
func (obj *_AllVirtualPlanCacheStatMgr) WithIndexBuilder(indexBuilder int64) Option {
	return optionFunc(func(o *options) { o.query["index_builder"] = indexBuilder })
}

// WithPcvSet pcv_set获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcvSet(pcvSet int64) Option {
	return optionFunc(func(o *options) { o.query["pcv_set"] = pcvSet })
}

// WithPcvRd pcv_rd获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcvRd(pcvRd int64) Option {
	return optionFunc(func(o *options) { o.query["pcv_rd"] = pcvRd })
}

// WithPcvWr pcv_wr获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcvWr(pcvWr int64) Option {
	return optionFunc(func(o *options) { o.query["pcv_wr"] = pcvWr })
}

// WithPcvGetPlanKey pcv_get_plan_key获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcvGetPlanKey(pcvGetPlanKey int64) Option {
	return optionFunc(func(o *options) { o.query["pcv_get_plan_key"] = pcvGetPlanKey })
}

// WithPcvGetPlKey pcv_get_pl_key获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcvGetPlKey(pcvGetPlKey int64) Option {
	return optionFunc(func(o *options) { o.query["pcv_get_pl_key"] = pcvGetPlKey })
}

// WithPcvExpireByUsed pcv_expire_by_used获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcvExpireByUsed(pcvExpireByUsed int64) Option {
	return optionFunc(func(o *options) { o.query["pcv_expire_by_used"] = pcvExpireByUsed })
}

// WithPcvExpireByMem pcv_expire_by_mem获取
func (obj *_AllVirtualPlanCacheStatMgr) WithPcvExpireByMem(pcvExpireByMem int64) Option {
	return optionFunc(func(o *options) { o.query["pcv_expire_by_mem"] = pcvExpireByMem })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPlanCacheStatMgr) GetByOption(opts ...Option) (result AllVirtualPlanCacheStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPlanCacheStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualPlanCacheStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPlanCacheStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPlanCacheStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where(options.query)
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
func (obj *_AllVirtualPlanCacheStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSQLNum 通过sql_num获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromSQLNum(sqlNum int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`sql_num` = ?", sqlNum).Find(&results).Error

	return
}

// GetBatchFromSQLNum 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromSQLNum(sqlNums []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`sql_num` IN (?)", sqlNums).Find(&results).Error

	return
}

// GetFromMemUsed 通过mem_used获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromMemUsed(memUsed int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`mem_used` = ?", memUsed).Find(&results).Error

	return
}

// GetBatchFromMemUsed 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromMemUsed(memUseds []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`mem_used` IN (?)", memUseds).Find(&results).Error

	return
}

// GetFromMemHold 通过mem_hold获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromMemHold(memHold int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`mem_hold` = ?", memHold).Find(&results).Error

	return
}

// GetBatchFromMemHold 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromMemHold(memHolds []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`mem_hold` IN (?)", memHolds).Find(&results).Error

	return
}

// GetFromAccessCount 通过access_count获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromAccessCount(accessCount int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`access_count` = ?", accessCount).Find(&results).Error

	return
}

// GetBatchFromAccessCount 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromAccessCount(accessCounts []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`access_count` IN (?)", accessCounts).Find(&results).Error

	return
}

// GetFromHitCount 通过hit_count获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromHitCount(hitCount int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`hit_count` = ?", hitCount).Find(&results).Error

	return
}

// GetBatchFromHitCount 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromHitCount(hitCounts []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`hit_count` IN (?)", hitCounts).Find(&results).Error

	return
}

// GetFromHitRate 通过hit_rate获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromHitRate(hitRate int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`hit_rate` = ?", hitRate).Find(&results).Error

	return
}

// GetBatchFromHitRate 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromHitRate(hitRates []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`hit_rate` IN (?)", hitRates).Find(&results).Error

	return
}

// GetFromPlanNum 通过plan_num获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPlanNum(planNum int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`plan_num` = ?", planNum).Find(&results).Error

	return
}

// GetBatchFromPlanNum 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPlanNum(planNums []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`plan_num` IN (?)", planNums).Find(&results).Error

	return
}

// GetFromMemLimit 通过mem_limit获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromMemLimit(memLimit int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`mem_limit` = ?", memLimit).Find(&results).Error

	return
}

// GetBatchFromMemLimit 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromMemLimit(memLimits []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`mem_limit` IN (?)", memLimits).Find(&results).Error

	return
}

// GetFromHashBucket 通过hash_bucket获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromHashBucket(hashBucket int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`hash_bucket` = ?", hashBucket).Find(&results).Error

	return
}

// GetBatchFromHashBucket 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromHashBucket(hashBuckets []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`hash_bucket` IN (?)", hashBuckets).Find(&results).Error

	return
}

// GetFromStmtkeyNum 通过stmtkey_num获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromStmtkeyNum(stmtkeyNum int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`stmtkey_num` = ?", stmtkeyNum).Find(&results).Error

	return
}

// GetBatchFromStmtkeyNum 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromStmtkeyNum(stmtkeyNums []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`stmtkey_num` IN (?)", stmtkeyNums).Find(&results).Error

	return
}

// GetFromPcRefPlanLocal 通过pc_ref_plan_local获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcRefPlanLocal(pcRefPlanLocal int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_local` = ?", pcRefPlanLocal).Find(&results).Error

	return
}

// GetBatchFromPcRefPlanLocal 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcRefPlanLocal(pcRefPlanLocals []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_local` IN (?)", pcRefPlanLocals).Find(&results).Error

	return
}

// GetFromPcRefPlanRemote 通过pc_ref_plan_remote获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcRefPlanRemote(pcRefPlanRemote int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_remote` = ?", pcRefPlanRemote).Find(&results).Error

	return
}

// GetBatchFromPcRefPlanRemote 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcRefPlanRemote(pcRefPlanRemotes []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_remote` IN (?)", pcRefPlanRemotes).Find(&results).Error

	return
}

// GetFromPcRefPlanDist 通过pc_ref_plan_dist获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcRefPlanDist(pcRefPlanDist int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_dist` = ?", pcRefPlanDist).Find(&results).Error

	return
}

// GetBatchFromPcRefPlanDist 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcRefPlanDist(pcRefPlanDists []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_dist` IN (?)", pcRefPlanDists).Find(&results).Error

	return
}

// GetFromPcRefPlanArr 通过pc_ref_plan_arr获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcRefPlanArr(pcRefPlanArr int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_arr` = ?", pcRefPlanArr).Find(&results).Error

	return
}

// GetBatchFromPcRefPlanArr 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcRefPlanArr(pcRefPlanArrs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_arr` IN (?)", pcRefPlanArrs).Find(&results).Error

	return
}

// GetFromPcRefPlanStat 通过pc_ref_plan_stat获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcRefPlanStat(pcRefPlanStat int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_stat` = ?", pcRefPlanStat).Find(&results).Error

	return
}

// GetBatchFromPcRefPlanStat 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcRefPlanStat(pcRefPlanStats []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_plan_stat` IN (?)", pcRefPlanStats).Find(&results).Error

	return
}

// GetFromPcRefPl 通过pc_ref_pl获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcRefPl(pcRefPl int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_pl` = ?", pcRefPl).Find(&results).Error

	return
}

// GetBatchFromPcRefPl 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcRefPl(pcRefPls []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_pl` IN (?)", pcRefPls).Find(&results).Error

	return
}

// GetFromPcRefPlStat 通过pc_ref_pl_stat获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcRefPlStat(pcRefPlStat int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_pl_stat` = ?", pcRefPlStat).Find(&results).Error

	return
}

// GetBatchFromPcRefPlStat 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcRefPlStat(pcRefPlStats []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pc_ref_pl_stat` IN (?)", pcRefPlStats).Find(&results).Error

	return
}

// GetFromPlanGen 通过plan_gen获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPlanGen(planGen int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`plan_gen` = ?", planGen).Find(&results).Error

	return
}

// GetBatchFromPlanGen 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPlanGen(planGens []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`plan_gen` IN (?)", planGens).Find(&results).Error

	return
}

// GetFromCliQuery 通过cli_query获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromCliQuery(cliQuery int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`cli_query` = ?", cliQuery).Find(&results).Error

	return
}

// GetBatchFromCliQuery 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromCliQuery(cliQuerys []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`cli_query` IN (?)", cliQuerys).Find(&results).Error

	return
}

// GetFromOutlineExec 通过outline_exec获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromOutlineExec(outlineExec int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`outline_exec` = ?", outlineExec).Find(&results).Error

	return
}

// GetBatchFromOutlineExec 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromOutlineExec(outlineExecs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`outline_exec` IN (?)", outlineExecs).Find(&results).Error

	return
}

// GetFromPlanExplain 通过plan_explain获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPlanExplain(planExplain int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`plan_explain` = ?", planExplain).Find(&results).Error

	return
}

// GetBatchFromPlanExplain 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPlanExplain(planExplains []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`plan_explain` IN (?)", planExplains).Find(&results).Error

	return
}

// GetFromAsynBaseline 通过asyn_baseline获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromAsynBaseline(asynBaseline int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`asyn_baseline` = ?", asynBaseline).Find(&results).Error

	return
}

// GetBatchFromAsynBaseline 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromAsynBaseline(asynBaselines []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`asyn_baseline` IN (?)", asynBaselines).Find(&results).Error

	return
}

// GetFromLoadBaseline 通过load_baseline获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromLoadBaseline(loadBaseline int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`load_baseline` = ?", loadBaseline).Find(&results).Error

	return
}

// GetBatchFromLoadBaseline 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromLoadBaseline(loadBaselines []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`load_baseline` IN (?)", loadBaselines).Find(&results).Error

	return
}

// GetFromPsExec 通过ps_exec获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPsExec(psExec int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`ps_exec` = ?", psExec).Find(&results).Error

	return
}

// GetBatchFromPsExec 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPsExec(psExecs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`ps_exec` IN (?)", psExecs).Find(&results).Error

	return
}

// GetFromGvSQL 通过gv_sql获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromGvSQL(gvSQL int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`gv_sql` = ?", gvSQL).Find(&results).Error

	return
}

// GetBatchFromGvSQL 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromGvSQL(gvSQLs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`gv_sql` IN (?)", gvSQLs).Find(&results).Error

	return
}

// GetFromPlAnon 通过pl_anon获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPlAnon(plAnon int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pl_anon` = ?", plAnon).Find(&results).Error

	return
}

// GetBatchFromPlAnon 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPlAnon(plAnons []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pl_anon` IN (?)", plAnons).Find(&results).Error

	return
}

// GetFromPlRoutine 通过pl_routine获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPlRoutine(plRoutine int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pl_routine` = ?", plRoutine).Find(&results).Error

	return
}

// GetBatchFromPlRoutine 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPlRoutine(plRoutines []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pl_routine` IN (?)", plRoutines).Find(&results).Error

	return
}

// GetFromPackageVar 通过package_var获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPackageVar(packageVar int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_var` = ?", packageVar).Find(&results).Error

	return
}

// GetBatchFromPackageVar 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPackageVar(packageVars []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_var` IN (?)", packageVars).Find(&results).Error

	return
}

// GetFromPackageType 通过package_type获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPackageType(packageType int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_type` = ?", packageType).Find(&results).Error

	return
}

// GetBatchFromPackageType 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPackageType(packageTypes []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_type` IN (?)", packageTypes).Find(&results).Error

	return
}

// GetFromPackageSpec 通过package_spec获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPackageSpec(packageSpec int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_spec` = ?", packageSpec).Find(&results).Error

	return
}

// GetBatchFromPackageSpec 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPackageSpec(packageSpecs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_spec` IN (?)", packageSpecs).Find(&results).Error

	return
}

// GetFromPackageBody 通过package_body获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPackageBody(packageBody int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_body` = ?", packageBody).Find(&results).Error

	return
}

// GetBatchFromPackageBody 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPackageBody(packageBodys []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_body` IN (?)", packageBodys).Find(&results).Error

	return
}

// GetFromPackageResv 通过package_resv获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPackageResv(packageResv int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_resv` = ?", packageResv).Find(&results).Error

	return
}

// GetBatchFromPackageResv 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPackageResv(packageResvs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`package_resv` IN (?)", packageResvs).Find(&results).Error

	return
}

// GetFromGetPkg 通过get_pkg获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromGetPkg(getPkg int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`get_pkg` = ?", getPkg).Find(&results).Error

	return
}

// GetBatchFromGetPkg 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromGetPkg(getPkgs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`get_pkg` IN (?)", getPkgs).Find(&results).Error

	return
}

// GetFromIndexBuilder 通过index_builder获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromIndexBuilder(indexBuilder int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`index_builder` = ?", indexBuilder).Find(&results).Error

	return
}

// GetBatchFromIndexBuilder 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromIndexBuilder(indexBuilders []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`index_builder` IN (?)", indexBuilders).Find(&results).Error

	return
}

// GetFromPcvSet 通过pcv_set获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcvSet(pcvSet int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_set` = ?", pcvSet).Find(&results).Error

	return
}

// GetBatchFromPcvSet 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcvSet(pcvSets []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_set` IN (?)", pcvSets).Find(&results).Error

	return
}

// GetFromPcvRd 通过pcv_rd获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcvRd(pcvRd int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_rd` = ?", pcvRd).Find(&results).Error

	return
}

// GetBatchFromPcvRd 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcvRd(pcvRds []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_rd` IN (?)", pcvRds).Find(&results).Error

	return
}

// GetFromPcvWr 通过pcv_wr获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcvWr(pcvWr int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_wr` = ?", pcvWr).Find(&results).Error

	return
}

// GetBatchFromPcvWr 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcvWr(pcvWrs []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_wr` IN (?)", pcvWrs).Find(&results).Error

	return
}

// GetFromPcvGetPlanKey 通过pcv_get_plan_key获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcvGetPlanKey(pcvGetPlanKey int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_get_plan_key` = ?", pcvGetPlanKey).Find(&results).Error

	return
}

// GetBatchFromPcvGetPlanKey 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcvGetPlanKey(pcvGetPlanKeys []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_get_plan_key` IN (?)", pcvGetPlanKeys).Find(&results).Error

	return
}

// GetFromPcvGetPlKey 通过pcv_get_pl_key获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcvGetPlKey(pcvGetPlKey int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_get_pl_key` = ?", pcvGetPlKey).Find(&results).Error

	return
}

// GetBatchFromPcvGetPlKey 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcvGetPlKey(pcvGetPlKeys []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_get_pl_key` IN (?)", pcvGetPlKeys).Find(&results).Error

	return
}

// GetFromPcvExpireByUsed 通过pcv_expire_by_used获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcvExpireByUsed(pcvExpireByUsed int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_expire_by_used` = ?", pcvExpireByUsed).Find(&results).Error

	return
}

// GetBatchFromPcvExpireByUsed 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcvExpireByUsed(pcvExpireByUseds []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_expire_by_used` IN (?)", pcvExpireByUseds).Find(&results).Error

	return
}

// GetFromPcvExpireByMem 通过pcv_expire_by_mem获取内容
func (obj *_AllVirtualPlanCacheStatMgr) GetFromPcvExpireByMem(pcvExpireByMem int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_expire_by_mem` = ?", pcvExpireByMem).Find(&results).Error

	return
}

// GetBatchFromPcvExpireByMem 批量查找
func (obj *_AllVirtualPlanCacheStatMgr) GetBatchFromPcvExpireByMem(pcvExpireByMems []int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`pcv_expire_by_mem` IN (?)", pcvExpireByMems).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPlanCacheStatMgr) FetchByPrimaryKey(tenantID int64, svrIP string, svrPort int64) (result AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, svrIP, svrPort).First(&result).Error

	return
}

// FetchIndexByAllVirtualPlanCacheStatI1  获取多个内容
func (obj *_AllVirtualPlanCacheStatMgr) FetchIndexByAllVirtualPlanCacheStatI1(tenantID int64) (results []*AllVirtualPlanCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCacheStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}
