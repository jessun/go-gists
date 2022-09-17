package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _V$sysstatMgr struct {
	*_BaseMgr
}

// V$sysstatMgr open func
func V$sysstatMgr(db *gorm.DB) *_V$sysstatMgr {
	if db == nil {
		panic(fmt.Errorf("V$sysstatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sysstatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sysstat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sysstatMgr) GetTableName() string {
	return "v$sysstat"
}

// Reset 重置gorm会话
func (obj *_V$sysstatMgr) Reset() *_V$sysstatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sysstatMgr) Get() (result V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sysstatMgr) Gets() (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sysstatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$sysstatMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithStatistic# STATISTIC#获取 
func (obj *_V$sysstatMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["STATISTIC#"] = statistic# })
}

// WithValue VALUE获取 
func (obj *_V$sysstatMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["VALUE"] = value })
}

// WithStatID STAT_ID获取 
func (obj *_V$sysstatMgr) WithStatID(statID int64) Option {
	return optionFunc(func(o *options) { o.query["STAT_ID"] = statID })
}

// WithName NAME获取 
func (obj *_V$sysstatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithClass CLASS获取 
func (obj *_V$sysstatMgr) WithClass(class int64) Option {
	return optionFunc(func(o *options) { o.query["CLASS"] = class })
}


// GetByOption 功能选项模式获取
func (obj *_V$sysstatMgr) GetByOption(opts ...Option) (result V$sysstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sysstatMgr) GetByOptions(opts ...Option) (results []*V$sysstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sysstatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sysstat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where(options.query)
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


// GetFromConID 通过CON_ID获取内容  
func (obj *_V$sysstatMgr) GetFromConID(conID int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sysstatMgr) GetBatchFromConID(conIDs []int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过STATISTIC#获取内容  
func (obj *_V$sysstatMgr) GetFromStatistic#(statistic# int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`STATISTIC#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_V$sysstatMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`STATISTIC#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromValue 通过VALUE获取内容  
func (obj *_V$sysstatMgr) GetFromValue(value int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`VALUE` = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量查找 
func (obj *_V$sysstatMgr) GetBatchFromValue(values []int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`VALUE` IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromStatID 通过STAT_ID获取内容  
func (obj *_V$sysstatMgr) GetFromStatID(statID int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`STAT_ID` = ?", statID).Find(&results).Error
	
	return
}

// GetBatchFromStatID 批量查找 
func (obj *_V$sysstatMgr) GetBatchFromStatID(statIDs []int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`STAT_ID` IN (?)", statIDs).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_V$sysstatMgr) GetFromName(name string) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_V$sysstatMgr) GetBatchFromName(names []string) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromClass 通过CLASS获取内容  
func (obj *_V$sysstatMgr) GetFromClass(class int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`CLASS` = ?", class).Find(&results).Error
	
	return
}

// GetBatchFromClass 批量查找 
func (obj *_V$sysstatMgr) GetBatchFromClass(classs []int64) (results []*V$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sysstat{}).Where("`CLASS` IN (?)", classs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

