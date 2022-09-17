package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$memoryMgr struct {
	*_BaseMgr
}

// Gv$memoryMgr open func
func Gv$memoryMgr(db *gorm.DB) *_Gv$memoryMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$memoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$memoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$memory"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$memoryMgr) GetTableName() string {
	return "gv$memory"
}

// Reset 重置gorm会话
func (obj *_Gv$memoryMgr) Reset() *_Gv$memoryMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$memoryMgr) Get() (result Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$memoryMgr) Gets() (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$memoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_Gv$memoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIP IP获取 
func (obj *_Gv$memoryMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["IP"] = ip })
}

// WithPort PORT获取 
func (obj *_Gv$memoryMgr) WithPort(port int64) Option {
	return optionFunc(func(o *options) { o.query["PORT"] = port })
}

// WithContext CONTEXT获取 
func (obj *_Gv$memoryMgr) WithContext(context string) Option {
	return optionFunc(func(o *options) { o.query["CONTEXT"] = context })
}

// WithCount COUNT获取 
func (obj *_Gv$memoryMgr) WithCount(count float64) Option {
	return optionFunc(func(o *options) { o.query["COUNT"] = count })
}

// WithUsed USED获取 
func (obj *_Gv$memoryMgr) WithUsed(used float64) Option {
	return optionFunc(func(o *options) { o.query["USED"] = used })
}

// WithAllocCount ALLOC_COUNT获取 
func (obj *_Gv$memoryMgr) WithAllocCount(allocCount float64) Option {
	return optionFunc(func(o *options) { o.query["ALLOC_COUNT"] = allocCount })
}

// WithFreeCount FREE_COUNT获取 
func (obj *_Gv$memoryMgr) WithFreeCount(freeCount float64) Option {
	return optionFunc(func(o *options) { o.query["FREE_COUNT"] = freeCount })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$memoryMgr) GetByOption(opts ...Option) (result Gv$memory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$memoryMgr) GetByOptions(opts ...Option) (results []*Gv$memory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$memoryMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$memory,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where(options.query)
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
func (obj *_Gv$memoryMgr) GetFromTenantID(tenantID int64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromIP 通过IP获取内容  
func (obj *_Gv$memoryMgr) GetFromIP(ip string) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`IP` = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromIP(ips []string) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`IP` IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过PORT获取内容  
func (obj *_Gv$memoryMgr) GetFromPort(port int64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`PORT` = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromPort(ports []int64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`PORT` IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromContext 通过CONTEXT获取内容  
func (obj *_Gv$memoryMgr) GetFromContext(context string) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`CONTEXT` = ?", context).Find(&results).Error
	
	return
}

// GetBatchFromContext 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromContext(contexts []string) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`CONTEXT` IN (?)", contexts).Find(&results).Error
	
	return
}
 
// GetFromCount 通过COUNT获取内容  
func (obj *_Gv$memoryMgr) GetFromCount(count float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`COUNT` = ?", count).Find(&results).Error
	
	return
}

// GetBatchFromCount 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromCount(counts []float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`COUNT` IN (?)", counts).Find(&results).Error
	
	return
}
 
// GetFromUsed 通过USED获取内容  
func (obj *_Gv$memoryMgr) GetFromUsed(used float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`USED` = ?", used).Find(&results).Error
	
	return
}

// GetBatchFromUsed 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromUsed(useds []float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`USED` IN (?)", useds).Find(&results).Error
	
	return
}
 
// GetFromAllocCount 通过ALLOC_COUNT获取内容  
func (obj *_Gv$memoryMgr) GetFromAllocCount(allocCount float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`ALLOC_COUNT` = ?", allocCount).Find(&results).Error
	
	return
}

// GetBatchFromAllocCount 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromAllocCount(allocCounts []float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`ALLOC_COUNT` IN (?)", allocCounts).Find(&results).Error
	
	return
}
 
// GetFromFreeCount 通过FREE_COUNT获取内容  
func (obj *_Gv$memoryMgr) GetFromFreeCount(freeCount float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`FREE_COUNT` = ?", freeCount).Find(&results).Error
	
	return
}

// GetBatchFromFreeCount 批量查找 
func (obj *_Gv$memoryMgr) GetBatchFromFreeCount(freeCounts []float64) (results []*Gv$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memory{}).Where("`FREE_COUNT` IN (?)", freeCounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

