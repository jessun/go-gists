package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$sysstatMgr struct {
	*_BaseMgr
}

// Gv$sysstatMgr open func
func Gv$sysstatMgr(db *gorm.DB) *_Gv$sysstatMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$sysstatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$sysstatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$sysstat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$sysstatMgr) GetTableName() string {
	return "gv$sysstat"
}

// Reset 重置gorm会话
func (obj *_Gv$sysstatMgr) Reset() *_Gv$sysstatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$sysstatMgr) Get() (result Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$sysstatMgr) Gets() (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$sysstatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_Gv$sysstatMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSvrIP SVR_IP获取 
func (obj *_Gv$sysstatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$sysstatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithStatistic# STATISTIC#获取 
func (obj *_Gv$sysstatMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["STATISTIC#"] = statistic# })
}

// WithValue VALUE获取 
func (obj *_Gv$sysstatMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["VALUE"] = value })
}

// WithStatID STAT_ID获取 
func (obj *_Gv$sysstatMgr) WithStatID(statID int64) Option {
	return optionFunc(func(o *options) { o.query["STAT_ID"] = statID })
}

// WithName NAME获取 
func (obj *_Gv$sysstatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithClass CLASS获取 
func (obj *_Gv$sysstatMgr) WithClass(class int64) Option {
	return optionFunc(func(o *options) { o.query["CLASS"] = class })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$sysstatMgr) GetByOption(opts ...Option) (result Gv$sysstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$sysstatMgr) GetByOptions(opts ...Option) (results []*Gv$sysstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$sysstatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$sysstat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where(options.query)
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
func (obj *_Gv$sysstatMgr) GetFromConID(conID int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromConID(conIDs []int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$sysstatMgr) GetFromSvrIP(svrIP string) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$sysstatMgr) GetFromSvrPort(svrPort int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过STATISTIC#获取内容  
func (obj *_Gv$sysstatMgr) GetFromStatistic#(statistic# int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`STATISTIC#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`STATISTIC#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromValue 通过VALUE获取内容  
func (obj *_Gv$sysstatMgr) GetFromValue(value int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`VALUE` = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromValue(values []int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`VALUE` IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromStatID 通过STAT_ID获取内容  
func (obj *_Gv$sysstatMgr) GetFromStatID(statID int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`STAT_ID` = ?", statID).Find(&results).Error
	
	return
}

// GetBatchFromStatID 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromStatID(statIDs []int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`STAT_ID` IN (?)", statIDs).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_Gv$sysstatMgr) GetFromName(name string) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromName(names []string) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromClass 通过CLASS获取内容  
func (obj *_Gv$sysstatMgr) GetFromClass(class int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`CLASS` = ?", class).Find(&results).Error
	
	return
}

// GetBatchFromClass 批量查找 
func (obj *_Gv$sysstatMgr) GetBatchFromClass(classs []int64) (results []*Gv$sysstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sysstat{}).Where("`CLASS` IN (?)", classs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

