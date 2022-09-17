package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$planCachePlanExplainMgr struct {
	*_BaseMgr
}

// V$planCachePlanExplainMgr open func
func V$planCachePlanExplainMgr(db *gorm.DB) *_V$planCachePlanExplainMgr {
	if db == nil {
		panic(fmt.Errorf("V$planCachePlanExplainMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$planCachePlanExplainMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$plan_cache_plan_explain"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$planCachePlanExplainMgr) GetTableName() string {
	return "v$plan_cache_plan_explain"
}

// Reset 重置gorm会话
func (obj *_V$planCachePlanExplainMgr) Reset() *_V$planCachePlanExplainMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$planCachePlanExplainMgr) Get() (result V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$planCachePlanExplainMgr) Gets() (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$planCachePlanExplainMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_V$planCachePlanExplainMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIP IP获取 
func (obj *_V$planCachePlanExplainMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["IP"] = ip })
}

// WithPort PORT获取 
func (obj *_V$planCachePlanExplainMgr) WithPort(port int64) Option {
	return optionFunc(func(o *options) { o.query["PORT"] = port })
}

// WithPlanID PLAN_ID获取 
func (obj *_V$planCachePlanExplainMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_ID"] = planID })
}

// WithPlanDepth PLAN_DEPTH获取 
func (obj *_V$planCachePlanExplainMgr) WithPlanDepth(planDepth int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_DEPTH"] = planDepth })
}

// WithPlanLineID PLAN_LINE_ID获取 
func (obj *_V$planCachePlanExplainMgr) WithPlanLineID(planLineID int64) Option {
	return optionFunc(func(o *options) { o.query["PLAN_LINE_ID"] = planLineID })
}

// WithOperator OPERATOR获取 
func (obj *_V$planCachePlanExplainMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["OPERATOR"] = operator })
}

// WithName NAME获取 
func (obj *_V$planCachePlanExplainMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithRows ROWS获取 
func (obj *_V$planCachePlanExplainMgr) WithRows(rows int64) Option {
	return optionFunc(func(o *options) { o.query["ROWS"] = rows })
}

// WithCost COST获取 
func (obj *_V$planCachePlanExplainMgr) WithCost(cost int64) Option {
	return optionFunc(func(o *options) { o.query["COST"] = cost })
}

// WithProperty PROPERTY获取 
func (obj *_V$planCachePlanExplainMgr) WithProperty(property string) Option {
	return optionFunc(func(o *options) { o.query["PROPERTY"] = property })
}


// GetByOption 功能选项模式获取
func (obj *_V$planCachePlanExplainMgr) GetByOption(opts ...Option) (result V$planCachePlanExplain, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$planCachePlanExplainMgr) GetByOptions(opts ...Option) (results []*V$planCachePlanExplain, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$planCachePlanExplainMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$planCachePlanExplain,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where(options.query)
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


// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromTenantID(tenantID int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromIP 通过IP获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromIP(ip string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`IP` = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromIP(ips []string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`IP` IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过PORT获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromPort(port int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PORT` = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromPort(ports []int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PORT` IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromPlanID 通过PLAN_ID获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromPlanID(planID int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PLAN_ID` = ?", planID).Find(&results).Error
	
	return
}

// GetBatchFromPlanID 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromPlanID(planIDs []int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PLAN_ID` IN (?)", planIDs).Find(&results).Error
	
	return
}
 
// GetFromPlanDepth 通过PLAN_DEPTH获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromPlanDepth(planDepth int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PLAN_DEPTH` = ?", planDepth).Find(&results).Error
	
	return
}

// GetBatchFromPlanDepth 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromPlanDepth(planDepths []int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PLAN_DEPTH` IN (?)", planDepths).Find(&results).Error
	
	return
}
 
// GetFromPlanLineID 通过PLAN_LINE_ID获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromPlanLineID(planLineID int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PLAN_LINE_ID` = ?", planLineID).Find(&results).Error
	
	return
}

// GetBatchFromPlanLineID 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromPlanLineID(planLineIDs []int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PLAN_LINE_ID` IN (?)", planLineIDs).Find(&results).Error
	
	return
}
 
// GetFromOperator 通过OPERATOR获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromOperator(operator string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`OPERATOR` = ?", operator).Find(&results).Error
	
	return
}

// GetBatchFromOperator 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromOperator(operators []string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`OPERATOR` IN (?)", operators).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromName(name string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromName(names []string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromRows 通过ROWS获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromRows(rows int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`ROWS` = ?", rows).Find(&results).Error
	
	return
}

// GetBatchFromRows 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromRows(rowss []int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`ROWS` IN (?)", rowss).Find(&results).Error
	
	return
}
 
// GetFromCost 通过COST获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromCost(cost int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`COST` = ?", cost).Find(&results).Error
	
	return
}

// GetBatchFromCost 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromCost(costs []int64) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`COST` IN (?)", costs).Find(&results).Error
	
	return
}
 
// GetFromProperty 通过PROPERTY获取内容  
func (obj *_V$planCachePlanExplainMgr) GetFromProperty(property string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PROPERTY` = ?", property).Find(&results).Error
	
	return
}

// GetBatchFromProperty 批量查找 
func (obj *_V$planCachePlanExplainMgr) GetBatchFromProperty(propertys []string) (results []*V$planCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$planCachePlanExplain{}).Where("`PROPERTY` IN (?)", propertys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

