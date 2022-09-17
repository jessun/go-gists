package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$sesstatMgr struct {
	*_BaseMgr
}

// Gv$sesstatMgr open func
func Gv$sesstatMgr(db *gorm.DB) *_Gv$sesstatMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$sesstatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$sesstatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$sesstat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$sesstatMgr) GetTableName() string {
	return "gv$sesstat"
}

// Reset 重置gorm会话
func (obj *_Gv$sesstatMgr) Reset() *_Gv$sesstatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$sesstatMgr) Get() (result Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$sesstatMgr) Gets() (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$sesstatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSid SID获取 
func (obj *_Gv$sesstatMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithConID CON_ID获取 
func (obj *_Gv$sesstatMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSvrIP SVR_IP获取 
func (obj *_Gv$sesstatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$sesstatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithStatistic# STATISTIC#获取 
func (obj *_Gv$sesstatMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["STATISTIC#"] = statistic# })
}

// WithValue VALUE获取 
func (obj *_Gv$sesstatMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["VALUE"] = value })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$sesstatMgr) GetByOption(opts ...Option) (result Gv$sesstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$sesstatMgr) GetByOptions(opts ...Option) (results []*Gv$sesstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$sesstatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$sesstat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where(options.query)
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


// GetFromSid 通过SID获取内容  
func (obj *_Gv$sesstatMgr) GetFromSid(sid int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_Gv$sesstatMgr) GetBatchFromSid(sids []int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromConID 通过CON_ID获取内容  
func (obj *_Gv$sesstatMgr) GetFromConID(conID int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_Gv$sesstatMgr) GetBatchFromConID(conIDs []int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$sesstatMgr) GetFromSvrIP(svrIP string) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$sesstatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$sesstatMgr) GetFromSvrPort(svrPort int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$sesstatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过STATISTIC#获取内容  
func (obj *_Gv$sesstatMgr) GetFromStatistic#(statistic# int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`STATISTIC#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_Gv$sesstatMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`STATISTIC#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromValue 通过VALUE获取内容  
func (obj *_Gv$sesstatMgr) GetFromValue(value int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`VALUE` = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量查找 
func (obj *_Gv$sesstatMgr) GetBatchFromValue(values []int64) (results []*Gv$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sesstat{}).Where("`VALUE` IN (?)", values).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

