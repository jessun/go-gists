package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _V$eventNameMgr struct {
	*_BaseMgr
}

// V$eventNameMgr open func
func V$eventNameMgr(db *gorm.DB) *_V$eventNameMgr {
	if db == nil {
		panic(fmt.Errorf("V$eventNameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$eventNameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$event_name"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$eventNameMgr) GetTableName() string {
	return "v$event_name"
}

// Reset 重置gorm会话
func (obj *_V$eventNameMgr) Reset() *_V$eventNameMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$eventNameMgr) Get() (result V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$eventNameMgr) Gets() (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$eventNameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$eventNameMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithEventID EVENT_ID获取 
func (obj *_V$eventNameMgr) WithEventID(eventID int64) Option {
	return optionFunc(func(o *options) { o.query["EVENT_ID"] = eventID })
}

// WithEvent# EVENT#获取 
func (obj *_V$eventNameMgr) WithEvent#(event# int64) Option {
	return optionFunc(func(o *options) { o.query["EVENT#"] = event# })
}

// WithName NAME获取 
func (obj *_V$eventNameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithDisplayName DISPLAY_NAME获取 
func (obj *_V$eventNameMgr) WithDisplayName(displayName string) Option {
	return optionFunc(func(o *options) { o.query["DISPLAY_NAME"] = displayName })
}

// WithParameter1 PARAMETER1获取 
func (obj *_V$eventNameMgr) WithParameter1(parameter1 string) Option {
	return optionFunc(func(o *options) { o.query["PARAMETER1"] = parameter1 })
}

// WithParameter2 PARAMETER2获取 
func (obj *_V$eventNameMgr) WithParameter2(parameter2 string) Option {
	return optionFunc(func(o *options) { o.query["PARAMETER2"] = parameter2 })
}

// WithParameter3 PARAMETER3获取 
func (obj *_V$eventNameMgr) WithParameter3(parameter3 string) Option {
	return optionFunc(func(o *options) { o.query["PARAMETER3"] = parameter3 })
}

// WithWaitClassID WAIT_CLASS_ID获取 
func (obj *_V$eventNameMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS_ID"] = waitClassID })
}

// WithWaitClass# WAIT_CLASS#获取 
func (obj *_V$eventNameMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS#"] = waitClass# })
}

// WithWaitClass WAIT_CLASS获取 
func (obj *_V$eventNameMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["WAIT_CLASS"] = waitClass })
}


// GetByOption 功能选项模式获取
func (obj *_V$eventNameMgr) GetByOption(opts ...Option) (result V$eventName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$eventNameMgr) GetByOptions(opts ...Option) (results []*V$eventName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$eventNameMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$eventName,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where(options.query)
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
func (obj *_V$eventNameMgr) GetFromConID(conID int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromConID(conIDs []int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromEventID 通过EVENT_ID获取内容  
func (obj *_V$eventNameMgr) GetFromEventID(eventID int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`EVENT_ID` = ?", eventID).Find(&results).Error
	
	return
}

// GetBatchFromEventID 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromEventID(eventIDs []int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`EVENT_ID` IN (?)", eventIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent# 通过EVENT#获取内容  
func (obj *_V$eventNameMgr) GetFromEvent#(event# int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`EVENT#` = ?", event#).Find(&results).Error
	
	return
}

// GetBatchFromEvent# 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromEvent#(event#s []int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`EVENT#` IN (?)", event#s).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_V$eventNameMgr) GetFromName(name string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromName(names []string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDisplayName 通过DISPLAY_NAME获取内容  
func (obj *_V$eventNameMgr) GetFromDisplayName(displayName string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`DISPLAY_NAME` = ?", displayName).Find(&results).Error
	
	return
}

// GetBatchFromDisplayName 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromDisplayName(displayNames []string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`DISPLAY_NAME` IN (?)", displayNames).Find(&results).Error
	
	return
}
 
// GetFromParameter1 通过PARAMETER1获取内容  
func (obj *_V$eventNameMgr) GetFromParameter1(parameter1 string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`PARAMETER1` = ?", parameter1).Find(&results).Error
	
	return
}

// GetBatchFromParameter1 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromParameter1(parameter1s []string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`PARAMETER1` IN (?)", parameter1s).Find(&results).Error
	
	return
}
 
// GetFromParameter2 通过PARAMETER2获取内容  
func (obj *_V$eventNameMgr) GetFromParameter2(parameter2 string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`PARAMETER2` = ?", parameter2).Find(&results).Error
	
	return
}

// GetBatchFromParameter2 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromParameter2(parameter2s []string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`PARAMETER2` IN (?)", parameter2s).Find(&results).Error
	
	return
}
 
// GetFromParameter3 通过PARAMETER3获取内容  
func (obj *_V$eventNameMgr) GetFromParameter3(parameter3 string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`PARAMETER3` = ?", parameter3).Find(&results).Error
	
	return
}

// GetBatchFromParameter3 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromParameter3(parameter3s []string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`PARAMETER3` IN (?)", parameter3s).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过WAIT_CLASS_ID获取内容  
func (obj *_V$eventNameMgr) GetFromWaitClassID(waitClassID int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`WAIT_CLASS_ID` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`WAIT_CLASS_ID` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过WAIT_CLASS#获取内容  
func (obj *_V$eventNameMgr) GetFromWaitClass#(waitClass# int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`WAIT_CLASS#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`WAIT_CLASS#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过WAIT_CLASS获取内容  
func (obj *_V$eventNameMgr) GetFromWaitClass(waitClass string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`WAIT_CLASS` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_V$eventNameMgr) GetBatchFromWaitClass(waitClasss []string) (results []*V$eventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$eventName{}).Where("`WAIT_CLASS` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

