package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _V$memoryMgr struct {
	*_BaseMgr
}

// V$memoryMgr open func
func V$memoryMgr(db *gorm.DB) *_V$memoryMgr {
	if db == nil {
		panic(fmt.Errorf("V$memoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$memoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$memory"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$memoryMgr) GetTableName() string {
	return "v$memory"
}

// Reset 重置gorm会话
func (obj *_V$memoryMgr) Reset() *_V$memoryMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$memoryMgr) Get() (result V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$memoryMgr) Gets() (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$memoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$memory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_V$memoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithContext CONTEXT获取 
func (obj *_V$memoryMgr) WithContext(context string) Option {
	return optionFunc(func(o *options) { o.query["CONTEXT"] = context })
}

// WithCount COUNT获取 
func (obj *_V$memoryMgr) WithCount(count float64) Option {
	return optionFunc(func(o *options) { o.query["COUNT"] = count })
}

// WithUsed USED获取 
func (obj *_V$memoryMgr) WithUsed(used float64) Option {
	return optionFunc(func(o *options) { o.query["USED"] = used })
}

// WithAllocCount ALLOC_COUNT获取 
func (obj *_V$memoryMgr) WithAllocCount(allocCount float64) Option {
	return optionFunc(func(o *options) { o.query["ALLOC_COUNT"] = allocCount })
}

// WithFreeCount FREE_COUNT获取 
func (obj *_V$memoryMgr) WithFreeCount(freeCount float64) Option {
	return optionFunc(func(o *options) { o.query["FREE_COUNT"] = freeCount })
}


// GetByOption 功能选项模式获取
func (obj *_V$memoryMgr) GetByOption(opts ...Option) (result V$memory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$memoryMgr) GetByOptions(opts ...Option) (results []*V$memory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$memoryMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$memory,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where(options.query)
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
func (obj *_V$memoryMgr) GetFromTenantID(tenantID int64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$memoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromContext 通过CONTEXT获取内容  
func (obj *_V$memoryMgr) GetFromContext(context string) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`CONTEXT` = ?", context).Find(&results).Error
	
	return
}

// GetBatchFromContext 批量查找 
func (obj *_V$memoryMgr) GetBatchFromContext(contexts []string) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`CONTEXT` IN (?)", contexts).Find(&results).Error
	
	return
}
 
// GetFromCount 通过COUNT获取内容  
func (obj *_V$memoryMgr) GetFromCount(count float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`COUNT` = ?", count).Find(&results).Error
	
	return
}

// GetBatchFromCount 批量查找 
func (obj *_V$memoryMgr) GetBatchFromCount(counts []float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`COUNT` IN (?)", counts).Find(&results).Error
	
	return
}
 
// GetFromUsed 通过USED获取内容  
func (obj *_V$memoryMgr) GetFromUsed(used float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`USED` = ?", used).Find(&results).Error
	
	return
}

// GetBatchFromUsed 批量查找 
func (obj *_V$memoryMgr) GetBatchFromUsed(useds []float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`USED` IN (?)", useds).Find(&results).Error
	
	return
}
 
// GetFromAllocCount 通过ALLOC_COUNT获取内容  
func (obj *_V$memoryMgr) GetFromAllocCount(allocCount float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`ALLOC_COUNT` = ?", allocCount).Find(&results).Error
	
	return
}

// GetBatchFromAllocCount 批量查找 
func (obj *_V$memoryMgr) GetBatchFromAllocCount(allocCounts []float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`ALLOC_COUNT` IN (?)", allocCounts).Find(&results).Error
	
	return
}
 
// GetFromFreeCount 通过FREE_COUNT获取内容  
func (obj *_V$memoryMgr) GetFromFreeCount(freeCount float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`FREE_COUNT` = ?", freeCount).Find(&results).Error
	
	return
}

// GetBatchFromFreeCount 批量查找 
func (obj *_V$memoryMgr) GetBatchFromFreeCount(freeCounts []float64) (results []*V$memory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memory{}).Where("`FREE_COUNT` IN (?)", freeCounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

