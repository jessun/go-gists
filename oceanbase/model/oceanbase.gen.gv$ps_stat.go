package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _Gv$psStatMgr struct {
	*_BaseMgr
}

// Gv$psStatMgr open func
func Gv$psStatMgr(db *gorm.DB) *_Gv$psStatMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$psStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$psStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$ps_stat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$psStatMgr) GetTableName() string {
	return "gv$ps_stat"
}

// Reset 重置gorm会话
func (obj *_Gv$psStatMgr) Reset() *_Gv$psStatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$psStatMgr) Get() (result Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$psStatMgr) Gets() (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$psStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_Gv$psStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取 
func (obj *_Gv$psStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_Gv$psStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithStmtCount stmt_count获取 
func (obj *_Gv$psStatMgr) WithStmtCount(stmtCount int64) Option {
	return optionFunc(func(o *options) { o.query["stmt_count"] = stmtCount })
}

// WithHitCount hit_count获取 
func (obj *_Gv$psStatMgr) WithHitCount(hitCount int64) Option {
	return optionFunc(func(o *options) { o.query["hit_count"] = hitCount })
}

// WithAccessCount access_count获取 
func (obj *_Gv$psStatMgr) WithAccessCount(accessCount int64) Option {
	return optionFunc(func(o *options) { o.query["access_count"] = accessCount })
}

// WithMemHold mem_hold获取 
func (obj *_Gv$psStatMgr) WithMemHold(memHold int64) Option {
	return optionFunc(func(o *options) { o.query["mem_hold"] = memHold })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$psStatMgr) GetByOption(opts ...Option) (result Gv$psStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$psStatMgr) GetByOptions(opts ...Option) (results []*Gv$psStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$psStatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$psStat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where(options.query)
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
func (obj *_Gv$psStatMgr) GetFromTenantID(tenantID int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$psStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_Gv$psStatMgr) GetFromSvrIP(svrIP string) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$psStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_Gv$psStatMgr) GetFromSvrPort(svrPort int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$psStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromStmtCount 通过stmt_count获取内容  
func (obj *_Gv$psStatMgr) GetFromStmtCount(stmtCount int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`stmt_count` = ?", stmtCount).Find(&results).Error
	
	return
}

// GetBatchFromStmtCount 批量查找 
func (obj *_Gv$psStatMgr) GetBatchFromStmtCount(stmtCounts []int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`stmt_count` IN (?)", stmtCounts).Find(&results).Error
	
	return
}
 
// GetFromHitCount 通过hit_count获取内容  
func (obj *_Gv$psStatMgr) GetFromHitCount(hitCount int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`hit_count` = ?", hitCount).Find(&results).Error
	
	return
}

// GetBatchFromHitCount 批量查找 
func (obj *_Gv$psStatMgr) GetBatchFromHitCount(hitCounts []int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`hit_count` IN (?)", hitCounts).Find(&results).Error
	
	return
}
 
// GetFromAccessCount 通过access_count获取内容  
func (obj *_Gv$psStatMgr) GetFromAccessCount(accessCount int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`access_count` = ?", accessCount).Find(&results).Error
	
	return
}

// GetBatchFromAccessCount 批量查找 
func (obj *_Gv$psStatMgr) GetBatchFromAccessCount(accessCounts []int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`access_count` IN (?)", accessCounts).Find(&results).Error
	
	return
}
 
// GetFromMemHold 通过mem_hold获取内容  
func (obj *_Gv$psStatMgr) GetFromMemHold(memHold int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`mem_hold` = ?", memHold).Find(&results).Error
	
	return
}

// GetBatchFromMemHold 批量查找 
func (obj *_Gv$psStatMgr) GetBatchFromMemHold(memHolds []int64) (results []*Gv$psStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$psStat{}).Where("`mem_hold` IN (?)", memHolds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

