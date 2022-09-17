package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$planCacheStatMgr struct {
	*_BaseMgr
}

// V$planCacheStatMgr open func
func V$planCacheStatMgr(db *gorm.DB) *_V$planCacheStatMgr {
	if db == nil {
		panic(fmt.Errorf("V$planCacheStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$planCacheStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$plan_cache_stat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$planCacheStatMgr) GetTableName() string {
	return "v$plan_cache_stat"
}

// Reset 重置gorm会话
func (obj *_V$planCacheStatMgr) Reset() *_V$planCacheStatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$planCacheStatMgr) Get() (result V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$planCacheStatMgr) Gets() (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$planCacheStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_V$planCacheStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取 
func (obj *_V$planCacheStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_V$planCacheStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLNum sql_num获取 
func (obj *_V$planCacheStatMgr) WithSQLNum(sqlNum int64) Option {
	return optionFunc(func(o *options) { o.query["sql_num"] = sqlNum })
}

// WithMemUsed mem_used获取 
func (obj *_V$planCacheStatMgr) WithMemUsed(memUsed int64) Option {
	return optionFunc(func(o *options) { o.query["mem_used"] = memUsed })
}

// WithMemHold mem_hold获取 
func (obj *_V$planCacheStatMgr) WithMemHold(memHold int64) Option {
	return optionFunc(func(o *options) { o.query["mem_hold"] = memHold })
}

// WithAccessCount access_count获取 
func (obj *_V$planCacheStatMgr) WithAccessCount(accessCount int64) Option {
	return optionFunc(func(o *options) { o.query["access_count"] = accessCount })
}

// WithHitCount hit_count获取 
func (obj *_V$planCacheStatMgr) WithHitCount(hitCount int64) Option {
	return optionFunc(func(o *options) { o.query["hit_count"] = hitCount })
}

// WithHitRate hit_rate获取 
func (obj *_V$planCacheStatMgr) WithHitRate(hitRate int64) Option {
	return optionFunc(func(o *options) { o.query["hit_rate"] = hitRate })
}

// WithPlanNum plan_num获取 
func (obj *_V$planCacheStatMgr) WithPlanNum(planNum int64) Option {
	return optionFunc(func(o *options) { o.query["plan_num"] = planNum })
}

// WithMemLimit mem_limit获取 
func (obj *_V$planCacheStatMgr) WithMemLimit(memLimit int64) Option {
	return optionFunc(func(o *options) { o.query["mem_limit"] = memLimit })
}

// WithHashBucket hash_bucket获取 
func (obj *_V$planCacheStatMgr) WithHashBucket(hashBucket int64) Option {
	return optionFunc(func(o *options) { o.query["hash_bucket"] = hashBucket })
}

// WithStmtkeyNum stmtkey_num获取 
func (obj *_V$planCacheStatMgr) WithStmtkeyNum(stmtkeyNum int64) Option {
	return optionFunc(func(o *options) { o.query["stmtkey_num"] = stmtkeyNum })
}


// GetByOption 功能选项模式获取
func (obj *_V$planCacheStatMgr) GetByOption(opts ...Option) (result V$planCacheStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$planCacheStatMgr) GetByOptions(opts ...Option) (results []*V$planCacheStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$planCacheStatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$planCacheStat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where(options.query)
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
func (obj *_V$planCacheStatMgr) GetFromTenantID(tenantID int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_V$planCacheStatMgr) GetFromSvrIP(svrIP string) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_V$planCacheStatMgr) GetFromSvrPort(svrPort int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSQLNum 通过sql_num获取内容  
func (obj *_V$planCacheStatMgr) GetFromSQLNum(sqlNum int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`sql_num` = ?", sqlNum).Find(&results).Error
	
	return
}

// GetBatchFromSQLNum 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromSQLNum(sqlNums []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`sql_num` IN (?)", sqlNums).Find(&results).Error
	
	return
}
 
// GetFromMemUsed 通过mem_used获取内容  
func (obj *_V$planCacheStatMgr) GetFromMemUsed(memUsed int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`mem_used` = ?", memUsed).Find(&results).Error
	
	return
}

// GetBatchFromMemUsed 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromMemUsed(memUseds []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`mem_used` IN (?)", memUseds).Find(&results).Error
	
	return
}
 
// GetFromMemHold 通过mem_hold获取内容  
func (obj *_V$planCacheStatMgr) GetFromMemHold(memHold int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`mem_hold` = ?", memHold).Find(&results).Error
	
	return
}

// GetBatchFromMemHold 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromMemHold(memHolds []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`mem_hold` IN (?)", memHolds).Find(&results).Error
	
	return
}
 
// GetFromAccessCount 通过access_count获取内容  
func (obj *_V$planCacheStatMgr) GetFromAccessCount(accessCount int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`access_count` = ?", accessCount).Find(&results).Error
	
	return
}

// GetBatchFromAccessCount 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromAccessCount(accessCounts []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`access_count` IN (?)", accessCounts).Find(&results).Error
	
	return
}
 
// GetFromHitCount 通过hit_count获取内容  
func (obj *_V$planCacheStatMgr) GetFromHitCount(hitCount int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`hit_count` = ?", hitCount).Find(&results).Error
	
	return
}

// GetBatchFromHitCount 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromHitCount(hitCounts []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`hit_count` IN (?)", hitCounts).Find(&results).Error
	
	return
}
 
// GetFromHitRate 通过hit_rate获取内容  
func (obj *_V$planCacheStatMgr) GetFromHitRate(hitRate int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`hit_rate` = ?", hitRate).Find(&results).Error
	
	return
}

// GetBatchFromHitRate 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromHitRate(hitRates []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`hit_rate` IN (?)", hitRates).Find(&results).Error
	
	return
}
 
// GetFromPlanNum 通过plan_num获取内容  
func (obj *_V$planCacheStatMgr) GetFromPlanNum(planNum int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`plan_num` = ?", planNum).Find(&results).Error
	
	return
}

// GetBatchFromPlanNum 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromPlanNum(planNums []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`plan_num` IN (?)", planNums).Find(&results).Error
	
	return
}
 
// GetFromMemLimit 通过mem_limit获取内容  
func (obj *_V$planCacheStatMgr) GetFromMemLimit(memLimit int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`mem_limit` = ?", memLimit).Find(&results).Error
	
	return
}

// GetBatchFromMemLimit 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromMemLimit(memLimits []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`mem_limit` IN (?)", memLimits).Find(&results).Error
	
	return
}
 
// GetFromHashBucket 通过hash_bucket获取内容  
func (obj *_V$planCacheStatMgr) GetFromHashBucket(hashBucket int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`hash_bucket` = ?", hashBucket).Find(&results).Error
	
	return
}

// GetBatchFromHashBucket 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromHashBucket(hashBuckets []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`hash_bucket` IN (?)", hashBuckets).Find(&results).Error
	
	return
}
 
// GetFromStmtkeyNum 通过stmtkey_num获取内容  
func (obj *_V$planCacheStatMgr) GetFromStmtkeyNum(stmtkeyNum int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`stmtkey_num` = ?", stmtkeyNum).Find(&results).Error
	
	return
}

// GetBatchFromStmtkeyNum 批量查找 
func (obj *_V$planCacheStatMgr) GetBatchFromStmtkeyNum(stmtkeyNums []int64) (results []*V$planCacheStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCacheStat{}).Where("`stmtkey_num` IN (?)", stmtkeyNums).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

