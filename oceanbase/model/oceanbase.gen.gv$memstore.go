package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$memstoreMgr struct {
	*_BaseMgr
}

// Gv$memstoreMgr open func
func Gv$memstoreMgr(db *gorm.DB) *_Gv$memstoreMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$memstoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$memstoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$memstore"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$memstoreMgr) GetTableName() string {
	return "gv$memstore"
}

// Reset 重置gorm会话
func (obj *_Gv$memstoreMgr) Reset() *_Gv$memstoreMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$memstoreMgr) Get() (result Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$memstoreMgr) Gets() (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$memstoreMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_Gv$memstoreMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIP IP获取 
func (obj *_Gv$memstoreMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["IP"] = ip })
}

// WithPort PORT获取 
func (obj *_Gv$memstoreMgr) WithPort(port int64) Option {
	return optionFunc(func(o *options) { o.query["PORT"] = port })
}

// WithActive ACTIVE获取 
func (obj *_Gv$memstoreMgr) WithActive(active int64) Option {
	return optionFunc(func(o *options) { o.query["ACTIVE"] = active })
}

// WithTotal TOTAL获取 
func (obj *_Gv$memstoreMgr) WithTotal(total int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL"] = total })
}

// WithFreezeTrigger FREEZE_TRIGGER获取 
func (obj *_Gv$memstoreMgr) WithFreezeTrigger(freezeTrigger int64) Option {
	return optionFunc(func(o *options) { o.query["FREEZE_TRIGGER"] = freezeTrigger })
}

// WithMemLimit MEM_LIMIT获取 
func (obj *_Gv$memstoreMgr) WithMemLimit(memLimit int64) Option {
	return optionFunc(func(o *options) { o.query["MEM_LIMIT"] = memLimit })
}

// WithFreezeCnt FREEZE_CNT获取 
func (obj *_Gv$memstoreMgr) WithFreezeCnt(freezeCnt int64) Option {
	return optionFunc(func(o *options) { o.query["FREEZE_CNT"] = freezeCnt })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$memstoreMgr) GetByOption(opts ...Option) (result Gv$memstore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$memstoreMgr) GetByOptions(opts ...Option) (results []*Gv$memstore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$memstoreMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$memstore,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where(options.query)
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
func (obj *_Gv$memstoreMgr) GetFromTenantID(tenantID int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromIP 通过IP获取内容  
func (obj *_Gv$memstoreMgr) GetFromIP(ip string) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`IP` = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromIP(ips []string) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`IP` IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过PORT获取内容  
func (obj *_Gv$memstoreMgr) GetFromPort(port int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`PORT` = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromPort(ports []int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`PORT` IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromActive 通过ACTIVE获取内容  
func (obj *_Gv$memstoreMgr) GetFromActive(active int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`ACTIVE` = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromActive(actives []int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`ACTIVE` IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromTotal 通过TOTAL获取内容  
func (obj *_Gv$memstoreMgr) GetFromTotal(total int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`TOTAL` = ?", total).Find(&results).Error
	
	return
}

// GetBatchFromTotal 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromTotal(totals []int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`TOTAL` IN (?)", totals).Find(&results).Error
	
	return
}
 
// GetFromFreezeTrigger 通过FREEZE_TRIGGER获取内容  
func (obj *_Gv$memstoreMgr) GetFromFreezeTrigger(freezeTrigger int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`FREEZE_TRIGGER` = ?", freezeTrigger).Find(&results).Error
	
	return
}

// GetBatchFromFreezeTrigger 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromFreezeTrigger(freezeTriggers []int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`FREEZE_TRIGGER` IN (?)", freezeTriggers).Find(&results).Error
	
	return
}
 
// GetFromMemLimit 通过MEM_LIMIT获取内容  
func (obj *_Gv$memstoreMgr) GetFromMemLimit(memLimit int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`MEM_LIMIT` = ?", memLimit).Find(&results).Error
	
	return
}

// GetBatchFromMemLimit 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromMemLimit(memLimits []int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`MEM_LIMIT` IN (?)", memLimits).Find(&results).Error
	
	return
}
 
// GetFromFreezeCnt 通过FREEZE_CNT获取内容  
func (obj *_Gv$memstoreMgr) GetFromFreezeCnt(freezeCnt int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`FREEZE_CNT` = ?", freezeCnt).Find(&results).Error
	
	return
}

// GetBatchFromFreezeCnt 批量查找 
func (obj *_Gv$memstoreMgr) GetBatchFromFreezeCnt(freezeCnts []int64) (results []*Gv$memstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$memstore{}).Where("`FREEZE_CNT` IN (?)", freezeCnts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

