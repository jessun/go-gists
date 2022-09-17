package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$planCacheReferenceInfoMgr struct {
	*_BaseMgr
}

// Gv$planCacheReferenceInfoMgr open func
func Gv$planCacheReferenceInfoMgr(db *gorm.DB) *_Gv$planCacheReferenceInfoMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$planCacheReferenceInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$planCacheReferenceInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$plan_cache_reference_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$planCacheReferenceInfoMgr) GetTableName() string {
	return "gv$plan_cache_reference_info"
}

// Reset 重置gorm会话
func (obj *_Gv$planCacheReferenceInfoMgr) Reset() *_Gv$planCacheReferenceInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$planCacheReferenceInfoMgr) Get() (result Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$planCacheReferenceInfoMgr) Gets() (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$planCacheReferenceInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithPcRefPlanLocal PC_REF_PLAN_LOCAL获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcRefPlanLocal(pcRefPlanLocal int64) Option {
	return optionFunc(func(o *options) { o.query["PC_REF_PLAN_LOCAL"] = pcRefPlanLocal })
}

// WithPcRefPlanRemote PC_REF_PLAN_REMOTE获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcRefPlanRemote(pcRefPlanRemote int64) Option {
	return optionFunc(func(o *options) { o.query["PC_REF_PLAN_REMOTE"] = pcRefPlanRemote })
}

// WithPcRefPlanDist PC_REF_PLAN_DIST获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcRefPlanDist(pcRefPlanDist int64) Option {
	return optionFunc(func(o *options) { o.query["PC_REF_PLAN_DIST"] = pcRefPlanDist })
}

// WithPcRefPlanArr PC_REF_PLAN_ARR获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcRefPlanArr(pcRefPlanArr int64) Option {
	return optionFunc(func(o *options) { o.query["PC_REF_PLAN_ARR"] = pcRefPlanArr })
}

// WithPcRefPl PC_REF_PL获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcRefPl(pcRefPl int64) Option {
	return optionFunc(func(o *options) { o.query["PC_REF_PL"] = pcRefPl })
}

// WithPcRefPlStat PC_REF_PL_STAT获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcRefPlStat(pcRefPlStat int64) Option {
	return optionFunc(func(o *options) { o.query["PC_REF_PL_STAT"] = pcRefPlStat })
}

// WithPlanGen PLAN_GEN获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPlanGen(planGen int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_GEN"] = planGen })
}

// WithCliQuery CLI_QUERY获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithCliQuery(cliQuery int64) Option {
	return optionFunc(func(o *options) { o.query["CLI_QUERY"] = cliQuery })
}

// WithOutlineExec OUTLINE_EXEC获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithOutlineExec(outlineExec int64) Option {
	return optionFunc(func(o *options) { o.query["OUTLINE_EXEC"] = outlineExec })
}

// WithPlanExplain PLAN_EXPLAIN获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPlanExplain(planExplain int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_EXPLAIN"] = planExplain })
}

// WithAsynBaseline ASYN_BASELINE获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithAsynBaseline(asynBaseline int64) Option {
	return optionFunc(func(o *options) { o.query["ASYN_BASELINE"] = asynBaseline })
}

// WithLoadBaseline LOAD_BASELINE获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithLoadBaseline(loadBaseline int64) Option {
	return optionFunc(func(o *options) { o.query["LOAD_BASELINE"] = loadBaseline })
}

// WithPsExec PS_EXEC获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPsExec(psExec int64) Option {
	return optionFunc(func(o *options) { o.query["PS_EXEC"] = psExec })
}

// WithGvSQL GV_SQL获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithGvSQL(gvSQL int64) Option {
	return optionFunc(func(o *options) { o.query["GV_SQL"] = gvSQL })
}

// WithPlAnon PL_ANON获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPlAnon(plAnon int64) Option {
	return optionFunc(func(o *options) { o.query["PL_ANON"] = plAnon })
}

// WithPlRoutine PL_ROUTINE获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPlRoutine(plRoutine int64) Option {
	return optionFunc(func(o *options) { o.query["PL_ROUTINE"] = plRoutine })
}

// WithPackageVar PACKAGE_VAR获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPackageVar(packageVar int64) Option {
	return optionFunc(func(o *options) { o.query["PACKAGE_VAR"] = packageVar })
}

// WithPackageType PACKAGE_TYPE获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPackageType(packageType int64) Option {
	return optionFunc(func(o *options) { o.query["PACKAGE_TYPE"] = packageType })
}

// WithPackageSpec PACKAGE_SPEC获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPackageSpec(packageSpec int64) Option {
	return optionFunc(func(o *options) { o.query["PACKAGE_SPEC"] = packageSpec })
}

// WithPackageBody PACKAGE_BODY获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPackageBody(packageBody int64) Option {
	return optionFunc(func(o *options) { o.query["PACKAGE_BODY"] = packageBody })
}

// WithPackageResv PACKAGE_RESV获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPackageResv(packageResv int64) Option {
	return optionFunc(func(o *options) { o.query["PACKAGE_RESV"] = packageResv })
}

// WithGetPkg GET_PKG获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithGetPkg(getPkg int64) Option {
	return optionFunc(func(o *options) { o.query["GET_PKG"] = getPkg })
}

// WithIndexBuilder INDEX_BUILDER获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithIndexBuilder(indexBuilder int64) Option {
	return optionFunc(func(o *options) { o.query["INDEX_BUILDER"] = indexBuilder })
}

// WithPcvSet PCV_SET获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcvSet(pcvSet int64) Option {
	return optionFunc(func(o *options) { o.query["PCV_SET"] = pcvSet })
}

// WithPcvRd PCV_RD获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcvRd(pcvRd int64) Option {
	return optionFunc(func(o *options) { o.query["PCV_RD"] = pcvRd })
}

// WithPcvWr PCV_WR获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcvWr(pcvWr int64) Option {
	return optionFunc(func(o *options) { o.query["PCV_WR"] = pcvWr })
}

// WithPcvGetPlanKey PCV_GET_PLAN_KEY获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcvGetPlanKey(pcvGetPlanKey int64) Option {
	return optionFunc(func(o *options) { o.query["PCV_GET_PLAN_KEY"] = pcvGetPlanKey })
}

// WithPcvGetPlKey PCV_GET_PL_KEY获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcvGetPlKey(pcvGetPlKey int64) Option {
	return optionFunc(func(o *options) { o.query["PCV_GET_PL_KEY"] = pcvGetPlKey })
}

// WithPcvExpireByUsed PCV_EXPIRE_BY_USED获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcvExpireByUsed(pcvExpireByUsed int64) Option {
	return optionFunc(func(o *options) { o.query["PCV_EXPIRE_BY_USED"] = pcvExpireByUsed })
}

// WithPcvExpireByMem PCV_EXPIRE_BY_MEM获取 
func (obj *_Gv$planCacheReferenceInfoMgr) WithPcvExpireByMem(pcvExpireByMem int64) Option {
	return optionFunc(func(o *options) { o.query["PCV_EXPIRE_BY_MEM"] = pcvExpireByMem })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$planCacheReferenceInfoMgr) GetByOption(opts ...Option) (result Gv$planCacheReferenceInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$planCacheReferenceInfoMgr) GetByOptions(opts ...Option) (results []*Gv$planCacheReferenceInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$planCacheReferenceInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$planCacheReferenceInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where(options.query)
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
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromSvrIP(svrIP string) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromSvrPort(svrPort int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromTenantID(tenantID int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromPcRefPlanLocal 通过PC_REF_PLAN_LOCAL获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcRefPlanLocal(pcRefPlanLocal int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_LOCAL` = ?", pcRefPlanLocal).Find(&results).Error
	
	return
}

// GetBatchFromPcRefPlanLocal 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcRefPlanLocal(pcRefPlanLocals []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_LOCAL` IN (?)", pcRefPlanLocals).Find(&results).Error
	
	return
}
 
// GetFromPcRefPlanRemote 通过PC_REF_PLAN_REMOTE获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcRefPlanRemote(pcRefPlanRemote int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_REMOTE` = ?", pcRefPlanRemote).Find(&results).Error
	
	return
}

// GetBatchFromPcRefPlanRemote 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcRefPlanRemote(pcRefPlanRemotes []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_REMOTE` IN (?)", pcRefPlanRemotes).Find(&results).Error
	
	return
}
 
// GetFromPcRefPlanDist 通过PC_REF_PLAN_DIST获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcRefPlanDist(pcRefPlanDist int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_DIST` = ?", pcRefPlanDist).Find(&results).Error
	
	return
}

// GetBatchFromPcRefPlanDist 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcRefPlanDist(pcRefPlanDists []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_DIST` IN (?)", pcRefPlanDists).Find(&results).Error
	
	return
}
 
// GetFromPcRefPlanArr 通过PC_REF_PLAN_ARR获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcRefPlanArr(pcRefPlanArr int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_ARR` = ?", pcRefPlanArr).Find(&results).Error
	
	return
}

// GetBatchFromPcRefPlanArr 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcRefPlanArr(pcRefPlanArrs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PLAN_ARR` IN (?)", pcRefPlanArrs).Find(&results).Error
	
	return
}
 
// GetFromPcRefPl 通过PC_REF_PL获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcRefPl(pcRefPl int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PL` = ?", pcRefPl).Find(&results).Error
	
	return
}

// GetBatchFromPcRefPl 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcRefPl(pcRefPls []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PL` IN (?)", pcRefPls).Find(&results).Error
	
	return
}
 
// GetFromPcRefPlStat 通过PC_REF_PL_STAT获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcRefPlStat(pcRefPlStat int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PL_STAT` = ?", pcRefPlStat).Find(&results).Error
	
	return
}

// GetBatchFromPcRefPlStat 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcRefPlStat(pcRefPlStats []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PC_REF_PL_STAT` IN (?)", pcRefPlStats).Find(&results).Error
	
	return
}
 
// GetFromPlanGen 通过PLAN_GEN获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPlanGen(planGen int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PLAN_GEN` = ?", planGen).Find(&results).Error
	
	return
}

// GetBatchFromPlanGen 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPlanGen(planGens []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PLAN_GEN` IN (?)", planGens).Find(&results).Error
	
	return
}
 
// GetFromCliQuery 通过CLI_QUERY获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromCliQuery(cliQuery int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`CLI_QUERY` = ?", cliQuery).Find(&results).Error
	
	return
}

// GetBatchFromCliQuery 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromCliQuery(cliQuerys []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`CLI_QUERY` IN (?)", cliQuerys).Find(&results).Error
	
	return
}
 
// GetFromOutlineExec 通过OUTLINE_EXEC获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromOutlineExec(outlineExec int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`OUTLINE_EXEC` = ?", outlineExec).Find(&results).Error
	
	return
}

// GetBatchFromOutlineExec 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromOutlineExec(outlineExecs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`OUTLINE_EXEC` IN (?)", outlineExecs).Find(&results).Error
	
	return
}
 
// GetFromPlanExplain 通过PLAN_EXPLAIN获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPlanExplain(planExplain int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PLAN_EXPLAIN` = ?", planExplain).Find(&results).Error
	
	return
}

// GetBatchFromPlanExplain 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPlanExplain(planExplains []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PLAN_EXPLAIN` IN (?)", planExplains).Find(&results).Error
	
	return
}
 
// GetFromAsynBaseline 通过ASYN_BASELINE获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromAsynBaseline(asynBaseline int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`ASYN_BASELINE` = ?", asynBaseline).Find(&results).Error
	
	return
}

// GetBatchFromAsynBaseline 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromAsynBaseline(asynBaselines []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`ASYN_BASELINE` IN (?)", asynBaselines).Find(&results).Error
	
	return
}
 
// GetFromLoadBaseline 通过LOAD_BASELINE获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromLoadBaseline(loadBaseline int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`LOAD_BASELINE` = ?", loadBaseline).Find(&results).Error
	
	return
}

// GetBatchFromLoadBaseline 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromLoadBaseline(loadBaselines []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`LOAD_BASELINE` IN (?)", loadBaselines).Find(&results).Error
	
	return
}
 
// GetFromPsExec 通过PS_EXEC获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPsExec(psExec int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PS_EXEC` = ?", psExec).Find(&results).Error
	
	return
}

// GetBatchFromPsExec 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPsExec(psExecs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PS_EXEC` IN (?)", psExecs).Find(&results).Error
	
	return
}
 
// GetFromGvSQL 通过GV_SQL获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromGvSQL(gvSQL int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`GV_SQL` = ?", gvSQL).Find(&results).Error
	
	return
}

// GetBatchFromGvSQL 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromGvSQL(gvSQLs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`GV_SQL` IN (?)", gvSQLs).Find(&results).Error
	
	return
}
 
// GetFromPlAnon 通过PL_ANON获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPlAnon(plAnon int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PL_ANON` = ?", plAnon).Find(&results).Error
	
	return
}

// GetBatchFromPlAnon 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPlAnon(plAnons []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PL_ANON` IN (?)", plAnons).Find(&results).Error
	
	return
}
 
// GetFromPlRoutine 通过PL_ROUTINE获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPlRoutine(plRoutine int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PL_ROUTINE` = ?", plRoutine).Find(&results).Error
	
	return
}

// GetBatchFromPlRoutine 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPlRoutine(plRoutines []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PL_ROUTINE` IN (?)", plRoutines).Find(&results).Error
	
	return
}
 
// GetFromPackageVar 通过PACKAGE_VAR获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPackageVar(packageVar int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_VAR` = ?", packageVar).Find(&results).Error
	
	return
}

// GetBatchFromPackageVar 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPackageVar(packageVars []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_VAR` IN (?)", packageVars).Find(&results).Error
	
	return
}
 
// GetFromPackageType 通过PACKAGE_TYPE获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPackageType(packageType int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_TYPE` = ?", packageType).Find(&results).Error
	
	return
}

// GetBatchFromPackageType 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPackageType(packageTypes []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_TYPE` IN (?)", packageTypes).Find(&results).Error
	
	return
}
 
// GetFromPackageSpec 通过PACKAGE_SPEC获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPackageSpec(packageSpec int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_SPEC` = ?", packageSpec).Find(&results).Error
	
	return
}

// GetBatchFromPackageSpec 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPackageSpec(packageSpecs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_SPEC` IN (?)", packageSpecs).Find(&results).Error
	
	return
}
 
// GetFromPackageBody 通过PACKAGE_BODY获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPackageBody(packageBody int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_BODY` = ?", packageBody).Find(&results).Error
	
	return
}

// GetBatchFromPackageBody 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPackageBody(packageBodys []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_BODY` IN (?)", packageBodys).Find(&results).Error
	
	return
}
 
// GetFromPackageResv 通过PACKAGE_RESV获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPackageResv(packageResv int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_RESV` = ?", packageResv).Find(&results).Error
	
	return
}

// GetBatchFromPackageResv 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPackageResv(packageResvs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PACKAGE_RESV` IN (?)", packageResvs).Find(&results).Error
	
	return
}
 
// GetFromGetPkg 通过GET_PKG获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromGetPkg(getPkg int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`GET_PKG` = ?", getPkg).Find(&results).Error
	
	return
}

// GetBatchFromGetPkg 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromGetPkg(getPkgs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`GET_PKG` IN (?)", getPkgs).Find(&results).Error
	
	return
}
 
// GetFromIndexBuilder 通过INDEX_BUILDER获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromIndexBuilder(indexBuilder int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`INDEX_BUILDER` = ?", indexBuilder).Find(&results).Error
	
	return
}

// GetBatchFromIndexBuilder 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromIndexBuilder(indexBuilders []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`INDEX_BUILDER` IN (?)", indexBuilders).Find(&results).Error
	
	return
}
 
// GetFromPcvSet 通过PCV_SET获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcvSet(pcvSet int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_SET` = ?", pcvSet).Find(&results).Error
	
	return
}

// GetBatchFromPcvSet 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcvSet(pcvSets []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_SET` IN (?)", pcvSets).Find(&results).Error
	
	return
}
 
// GetFromPcvRd 通过PCV_RD获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcvRd(pcvRd int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_RD` = ?", pcvRd).Find(&results).Error
	
	return
}

// GetBatchFromPcvRd 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcvRd(pcvRds []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_RD` IN (?)", pcvRds).Find(&results).Error
	
	return
}
 
// GetFromPcvWr 通过PCV_WR获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcvWr(pcvWr int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_WR` = ?", pcvWr).Find(&results).Error
	
	return
}

// GetBatchFromPcvWr 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcvWr(pcvWrs []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_WR` IN (?)", pcvWrs).Find(&results).Error
	
	return
}
 
// GetFromPcvGetPlanKey 通过PCV_GET_PLAN_KEY获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcvGetPlanKey(pcvGetPlanKey int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_GET_PLAN_KEY` = ?", pcvGetPlanKey).Find(&results).Error
	
	return
}

// GetBatchFromPcvGetPlanKey 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcvGetPlanKey(pcvGetPlanKeys []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_GET_PLAN_KEY` IN (?)", pcvGetPlanKeys).Find(&results).Error
	
	return
}
 
// GetFromPcvGetPlKey 通过PCV_GET_PL_KEY获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcvGetPlKey(pcvGetPlKey int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_GET_PL_KEY` = ?", pcvGetPlKey).Find(&results).Error
	
	return
}

// GetBatchFromPcvGetPlKey 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcvGetPlKey(pcvGetPlKeys []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_GET_PL_KEY` IN (?)", pcvGetPlKeys).Find(&results).Error
	
	return
}
 
// GetFromPcvExpireByUsed 通过PCV_EXPIRE_BY_USED获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcvExpireByUsed(pcvExpireByUsed int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_EXPIRE_BY_USED` = ?", pcvExpireByUsed).Find(&results).Error
	
	return
}

// GetBatchFromPcvExpireByUsed 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcvExpireByUsed(pcvExpireByUseds []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_EXPIRE_BY_USED` IN (?)", pcvExpireByUseds).Find(&results).Error
	
	return
}
 
// GetFromPcvExpireByMem 通过PCV_EXPIRE_BY_MEM获取内容  
func (obj *_Gv$planCacheReferenceInfoMgr) GetFromPcvExpireByMem(pcvExpireByMem int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_EXPIRE_BY_MEM` = ?", pcvExpireByMem).Find(&results).Error
	
	return
}

// GetBatchFromPcvExpireByMem 批量查找 
func (obj *_Gv$planCacheReferenceInfoMgr) GetBatchFromPcvExpireByMem(pcvExpireByMems []int64) (results []*Gv$planCacheReferenceInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$planCacheReferenceInfo{}).Where("`PCV_EXPIRE_BY_MEM` IN (?)", pcvExpireByMems).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

