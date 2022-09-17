package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _Gv$serverMemstoreMgr struct {
	*_BaseMgr
}

// Gv$serverMemstoreMgr open func
func Gv$serverMemstoreMgr(db *gorm.DB) *_Gv$serverMemstoreMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$serverMemstoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$serverMemstoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$server_memstore"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$serverMemstoreMgr) GetTableName() string {
	return "gv$server_memstore"
}

// Reset 重置gorm会话
func (obj *_Gv$serverMemstoreMgr) Reset() *_Gv$serverMemstoreMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$serverMemstoreMgr) Get() (result Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$serverMemstoreMgr) Gets() (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$serverMemstoreMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIP IP获取 
func (obj *_Gv$serverMemstoreMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["IP"] = ip })
}

// WithPort PORT获取 
func (obj *_Gv$serverMemstoreMgr) WithPort(port int64) Option {
	return optionFunc(func(o *options) { o.query["PORT"] = port })
}

// WithActive ACTIVE获取 
func (obj *_Gv$serverMemstoreMgr) WithActive(active int64) Option {
	return optionFunc(func(o *options) { o.query["ACTIVE"] = active })
}

// WithTotal TOTAL获取 
func (obj *_Gv$serverMemstoreMgr) WithTotal(total int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL"] = total })
}

// WithFreezeTrigger FREEZE_TRIGGER获取 
func (obj *_Gv$serverMemstoreMgr) WithFreezeTrigger(freezeTrigger int64) Option {
	return optionFunc(func(o *options) { o.query["FREEZE_TRIGGER"] = freezeTrigger })
}

// WithMemLimit MEM_LIMIT获取 
func (obj *_Gv$serverMemstoreMgr) WithMemLimit(memLimit int64) Option {
	return optionFunc(func(o *options) { o.query["MEM_LIMIT"] = memLimit })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$serverMemstoreMgr) GetByOption(opts ...Option) (result Gv$serverMemstore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$serverMemstoreMgr) GetByOptions(opts ...Option) (results []*Gv$serverMemstore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$serverMemstoreMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$serverMemstore,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where(options.query)
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


// GetFromIP 通过IP获取内容  
func (obj *_Gv$serverMemstoreMgr) GetFromIP(ip string) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`IP` = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量查找 
func (obj *_Gv$serverMemstoreMgr) GetBatchFromIP(ips []string) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`IP` IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过PORT获取内容  
func (obj *_Gv$serverMemstoreMgr) GetFromPort(port int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`PORT` = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量查找 
func (obj *_Gv$serverMemstoreMgr) GetBatchFromPort(ports []int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`PORT` IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromActive 通过ACTIVE获取内容  
func (obj *_Gv$serverMemstoreMgr) GetFromActive(active int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`ACTIVE` = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量查找 
func (obj *_Gv$serverMemstoreMgr) GetBatchFromActive(actives []int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`ACTIVE` IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromTotal 通过TOTAL获取内容  
func (obj *_Gv$serverMemstoreMgr) GetFromTotal(total int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`TOTAL` = ?", total).Find(&results).Error
	
	return
}

// GetBatchFromTotal 批量查找 
func (obj *_Gv$serverMemstoreMgr) GetBatchFromTotal(totals []int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`TOTAL` IN (?)", totals).Find(&results).Error
	
	return
}
 
// GetFromFreezeTrigger 通过FREEZE_TRIGGER获取内容  
func (obj *_Gv$serverMemstoreMgr) GetFromFreezeTrigger(freezeTrigger int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`FREEZE_TRIGGER` = ?", freezeTrigger).Find(&results).Error
	
	return
}

// GetBatchFromFreezeTrigger 批量查找 
func (obj *_Gv$serverMemstoreMgr) GetBatchFromFreezeTrigger(freezeTriggers []int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`FREEZE_TRIGGER` IN (?)", freezeTriggers).Find(&results).Error
	
	return
}
 
// GetFromMemLimit 通过MEM_LIMIT获取内容  
func (obj *_Gv$serverMemstoreMgr) GetFromMemLimit(memLimit int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`MEM_LIMIT` = ?", memLimit).Find(&results).Error
	
	return
}

// GetBatchFromMemLimit 批量查找 
func (obj *_Gv$serverMemstoreMgr) GetBatchFromMemLimit(memLimits []int64) (results []*Gv$serverMemstore, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$serverMemstore{}).Where("`MEM_LIMIT` IN (?)", memLimits).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

