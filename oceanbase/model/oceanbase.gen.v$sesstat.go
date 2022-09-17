package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sesstatMgr struct {
	*_BaseMgr
}

// V$sesstatMgr open func
func V$sesstatMgr(db *gorm.DB) *_V$sesstatMgr {
	if db == nil {
		panic(fmt.Errorf("V$sesstatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sesstatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sesstat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sesstatMgr) GetTableName() string {
	return "v$sesstat"
}

// Reset 重置gorm会话
func (obj *_V$sesstatMgr) Reset() *_V$sesstatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sesstatMgr) Get() (result V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sesstatMgr) Gets() (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sesstatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSid SID获取 
func (obj *_V$sesstatMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithConID CON_ID获取 
func (obj *_V$sesstatMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithStatistic# STATISTIC#获取 
func (obj *_V$sesstatMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["STATISTIC#"] = statistic# })
}

// WithValue VALUE获取 
func (obj *_V$sesstatMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["VALUE"] = value })
}


// GetByOption 功能选项模式获取
func (obj *_V$sesstatMgr) GetByOption(opts ...Option) (result V$sesstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sesstatMgr) GetByOptions(opts ...Option) (results []*V$sesstat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sesstatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sesstat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where(options.query)
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
func (obj *_V$sesstatMgr) GetFromSid(sid int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_V$sesstatMgr) GetBatchFromSid(sids []int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromConID 通过CON_ID获取内容  
func (obj *_V$sesstatMgr) GetFromConID(conID int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sesstatMgr) GetBatchFromConID(conIDs []int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过STATISTIC#获取内容  
func (obj *_V$sesstatMgr) GetFromStatistic#(statistic# int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`STATISTIC#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_V$sesstatMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`STATISTIC#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromValue 通过VALUE获取内容  
func (obj *_V$sesstatMgr) GetFromValue(value int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`VALUE` = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量查找 
func (obj *_V$sesstatMgr) GetBatchFromValue(values []int64) (results []*V$sesstat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sesstat{}).Where("`VALUE` IN (?)", values).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

