package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$memstoreMgr struct {
	*_BaseMgr
}

// V$memstoreMgr open func
func V$memstoreMgr(db *gorm.DB) *_V$memstoreMgr {
	if db == nil {
		panic(fmt.Errorf("V$memstoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$memstoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$memstore"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$memstoreMgr) GetTableName() string {
	return "v$memstore"
}

// Reset 重置gorm会话
func (obj *_V$memstoreMgr) Reset() *_V$memstoreMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$memstoreMgr) Get() (result V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$memstoreMgr) Gets() (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$memstoreMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_V$memstoreMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithActive ACTIVE获取 
func (obj *_V$memstoreMgr) WithActive(active int64) Option {
	return optionFunc(func(o *options) { o.query["ACTIVE"] = active })
}

// WithTotal TOTAL获取 
func (obj *_V$memstoreMgr) WithTotal(total int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL"] = total })
}

// WithFreezeTrigger FREEZE_TRIGGER获取 
func (obj *_V$memstoreMgr) WithFreezeTrigger(freezeTrigger int64) Option {
	return optionFunc(func(o *options) { o.query["FREEZE_TRIGGER"] = freezeTrigger })
}

// WithMemLimit MEM_LIMIT获取 
func (obj *_V$memstoreMgr) WithMemLimit(memLimit int64) Option {
	return optionFunc(func(o *options) { o.query["MEM_LIMIT"] = memLimit })
}


// GetByOption 功能选项模式获取
func (obj *_V$memstoreMgr) GetByOption(opts ...Option) (result V$memstore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$memstoreMgr) GetByOptions(opts ...Option) (results []*V$memstore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$memstoreMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$memstore,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where(options.query)
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
func (obj *_V$memstoreMgr) GetFromTenantID(tenantID int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$memstoreMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromActive 通过ACTIVE获取内容  
func (obj *_V$memstoreMgr) GetFromActive(active int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`ACTIVE` = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量查找 
func (obj *_V$memstoreMgr) GetBatchFromActive(actives []int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`ACTIVE` IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromTotal 通过TOTAL获取内容  
func (obj *_V$memstoreMgr) GetFromTotal(total int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`TOTAL` = ?", total).Find(&results).Error
	
	return
}

// GetBatchFromTotal 批量查找 
func (obj *_V$memstoreMgr) GetBatchFromTotal(totals []int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`TOTAL` IN (?)", totals).Find(&results).Error
	
	return
}
 
// GetFromFreezeTrigger 通过FREEZE_TRIGGER获取内容  
func (obj *_V$memstoreMgr) GetFromFreezeTrigger(freezeTrigger int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`FREEZE_TRIGGER` = ?", freezeTrigger).Find(&results).Error
	
	return
}

// GetBatchFromFreezeTrigger 批量查找 
func (obj *_V$memstoreMgr) GetBatchFromFreezeTrigger(freezeTriggers []int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`FREEZE_TRIGGER` IN (?)", freezeTriggers).Find(&results).Error
	
	return
}
 
// GetFromMemLimit 通过MEM_LIMIT获取内容  
func (obj *_V$memstoreMgr) GetFromMemLimit(memLimit int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`MEM_LIMIT` = ?", memLimit).Find(&results).Error
	
	return
}

// GetBatchFromMemLimit 批量查找 
func (obj *_V$memstoreMgr) GetBatchFromMemLimit(memLimits []int64) (results []*V$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$memstore{}).Where("`MEM_LIMIT` IN (?)", memLimits).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

