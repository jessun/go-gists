package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$statnameMgr struct {
	*_BaseMgr
}

// V$statnameMgr open func
func V$statnameMgr(db *gorm.DB) *_V$statnameMgr {
	if db == nil {
		panic(fmt.Errorf("V$statnameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$statnameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$statname"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$statnameMgr) GetTableName() string {
	return "v$statname"
}

// Reset 重置gorm会话
func (obj *_V$statnameMgr) Reset() *_V$statnameMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$statnameMgr) Get() (result V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$statnameMgr) Gets() (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$statnameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$statname{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$statnameMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithStatID STAT_ID获取 
func (obj *_V$statnameMgr) WithStatID(statID int64) Option {
	return optionFunc(func(o *options) { o.query["STAT_ID"] = statID })
}

// WithStatistic# STATISTIC#获取 
func (obj *_V$statnameMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["STATISTIC#"] = statistic# })
}

// WithName NAME获取 
func (obj *_V$statnameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithDisplayName DISPLAY_NAME获取 
func (obj *_V$statnameMgr) WithDisplayName(displayName string) Option {
	return optionFunc(func(o *options) { o.query["DISPLAY_NAME"] = displayName })
}

// WithClass CLASS获取 
func (obj *_V$statnameMgr) WithClass(class int64) Option {
	return optionFunc(func(o *options) { o.query["CLASS"] = class })
}


// GetByOption 功能选项模式获取
func (obj *_V$statnameMgr) GetByOption(opts ...Option) (result V$statname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$statnameMgr) GetByOptions(opts ...Option) (results []*V$statname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$statnameMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$statname,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where(options.query)
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
func (obj *_V$statnameMgr) GetFromConID(conID int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$statnameMgr) GetBatchFromConID(conIDs []int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromStatID 通过STAT_ID获取内容  
func (obj *_V$statnameMgr) GetFromStatID(statID int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`STAT_ID` = ?", statID).Find(&results).Error
	
	return
}

// GetBatchFromStatID 批量查找 
func (obj *_V$statnameMgr) GetBatchFromStatID(statIDs []int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`STAT_ID` IN (?)", statIDs).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过STATISTIC#获取内容  
func (obj *_V$statnameMgr) GetFromStatistic#(statistic# int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`STATISTIC#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_V$statnameMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`STATISTIC#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_V$statnameMgr) GetFromName(name string) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_V$statnameMgr) GetBatchFromName(names []string) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDisplayName 通过DISPLAY_NAME获取内容  
func (obj *_V$statnameMgr) GetFromDisplayName(displayName string) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`DISPLAY_NAME` = ?", displayName).Find(&results).Error
	
	return
}

// GetBatchFromDisplayName 批量查找 
func (obj *_V$statnameMgr) GetBatchFromDisplayName(displayNames []string) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`DISPLAY_NAME` IN (?)", displayNames).Find(&results).Error
	
	return
}
 
// GetFromClass 通过CLASS获取内容  
func (obj *_V$statnameMgr) GetFromClass(class int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`CLASS` = ?", class).Find(&results).Error
	
	return
}

// GetBatchFromClass 批量查找 
func (obj *_V$statnameMgr) GetBatchFromClass(classs []int64) (results []*V$statname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$statname{}).Where("`CLASS` IN (?)", classs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

