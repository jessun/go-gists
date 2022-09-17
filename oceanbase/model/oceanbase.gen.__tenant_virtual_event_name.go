package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _TenantVirtualEventNameMgr struct {
	*_BaseMgr
}

// TenantVirtualEventNameMgr open func
func TenantVirtualEventNameMgr(db *gorm.DB) *_TenantVirtualEventNameMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualEventNameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualEventNameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_event_name"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualEventNameMgr) GetTableName() string {
	return "__tenant_virtual_event_name"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualEventNameMgr) Reset() *_TenantVirtualEventNameMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_TenantVirtualEventNameMgr) Get() (result TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualEventNameMgr) Gets() (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualEventNameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_TenantVirtualEventNameMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithEventID event_id获取 
func (obj *_TenantVirtualEventNameMgr) WithEventID(eventID int64) Option {
	return optionFunc(func(o *options) { o.query["event_id"] = eventID })
}

// WithEvent# event#获取 
func (obj *_TenantVirtualEventNameMgr) WithEvent#(event# int64) Option {
	return optionFunc(func(o *options) { o.query["event#"] = event# })
}

// WithName name获取 
func (obj *_TenantVirtualEventNameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDisplayName display_name获取 
func (obj *_TenantVirtualEventNameMgr) WithDisplayName(displayName string) Option {
	return optionFunc(func(o *options) { o.query["display_name"] = displayName })
}

// WithParameter1 parameter1获取 
func (obj *_TenantVirtualEventNameMgr) WithParameter1(parameter1 string) Option {
	return optionFunc(func(o *options) { o.query["parameter1"] = parameter1 })
}

// WithParameter2 parameter2获取 
func (obj *_TenantVirtualEventNameMgr) WithParameter2(parameter2 string) Option {
	return optionFunc(func(o *options) { o.query["parameter2"] = parameter2 })
}

// WithParameter3 parameter3获取 
func (obj *_TenantVirtualEventNameMgr) WithParameter3(parameter3 string) Option {
	return optionFunc(func(o *options) { o.query["parameter3"] = parameter3 })
}

// WithWaitClassID wait_class_id获取 
func (obj *_TenantVirtualEventNameMgr) WithWaitClassID(waitClassID int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class_id"] = waitClassID })
}

// WithWaitClass# wait_class#获取 
func (obj *_TenantVirtualEventNameMgr) WithWaitClass#(waitClass# int64) Option {
	return optionFunc(func(o *options) { o.query["wait_class#"] = waitClass# })
}

// WithWaitClass wait_class获取 
func (obj *_TenantVirtualEventNameMgr) WithWaitClass(waitClass string) Option {
	return optionFunc(func(o *options) { o.query["wait_class"] = waitClass })
}


// GetByOption 功能选项模式获取
func (obj *_TenantVirtualEventNameMgr) GetByOption(opts ...Option) (result TenantVirtualEventName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualEventNameMgr) GetByOptions(opts ...Option) (results []*TenantVirtualEventName, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_TenantVirtualEventNameMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualEventName,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where(options.query)
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


// GetFromTenantID 通过tenant_id获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromTenantID(tenantID int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromEventID 通过event_id获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromEventID(eventID int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`event_id` = ?", eventID).Find(&results).Error
	
	return
}

// GetBatchFromEventID 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromEventID(eventIDs []int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`event_id` IN (?)", eventIDs).Find(&results).Error
	
	return
}
 
// GetFromEvent# 通过event#获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromEvent#(event# int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`event#` = ?", event#).Find(&results).Error
	
	return
}

// GetBatchFromEvent# 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromEvent#(event#s []int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`event#` IN (?)", event#s).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromName(name string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromName(names []string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDisplayName 通过display_name获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromDisplayName(displayName string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`display_name` = ?", displayName).Find(&results).Error
	
	return
}

// GetBatchFromDisplayName 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromDisplayName(displayNames []string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`display_name` IN (?)", displayNames).Find(&results).Error
	
	return
}
 
// GetFromParameter1 通过parameter1获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromParameter1(parameter1 string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`parameter1` = ?", parameter1).Find(&results).Error
	
	return
}

// GetBatchFromParameter1 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromParameter1(parameter1s []string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`parameter1` IN (?)", parameter1s).Find(&results).Error
	
	return
}
 
// GetFromParameter2 通过parameter2获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromParameter2(parameter2 string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`parameter2` = ?", parameter2).Find(&results).Error
	
	return
}

// GetBatchFromParameter2 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromParameter2(parameter2s []string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`parameter2` IN (?)", parameter2s).Find(&results).Error
	
	return
}
 
// GetFromParameter3 通过parameter3获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromParameter3(parameter3 string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`parameter3` = ?", parameter3).Find(&results).Error
	
	return
}

// GetBatchFromParameter3 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromParameter3(parameter3s []string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`parameter3` IN (?)", parameter3s).Find(&results).Error
	
	return
}
 
// GetFromWaitClassID 通过wait_class_id获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromWaitClassID(waitClassID int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`wait_class_id` = ?", waitClassID).Find(&results).Error
	
	return
}

// GetBatchFromWaitClassID 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromWaitClassID(waitClassIDs []int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`wait_class_id` IN (?)", waitClassIDs).Find(&results).Error
	
	return
}
 
// GetFromWaitClass# 通过wait_class#获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromWaitClass#(waitClass# int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`wait_class#` = ?", waitClass#).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass# 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromWaitClass#(waitClass#s []int64) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`wait_class#` IN (?)", waitClass#s).Find(&results).Error
	
	return
}
 
// GetFromWaitClass 通过wait_class获取内容  
func (obj *_TenantVirtualEventNameMgr) GetFromWaitClass(waitClass string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`wait_class` = ?", waitClass).Find(&results).Error
	
	return
}

// GetBatchFromWaitClass 批量查找 
func (obj *_TenantVirtualEventNameMgr) GetBatchFromWaitClass(waitClasss []string) (results []*TenantVirtualEventName, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualEventName{}).Where("`wait_class` IN (?)", waitClasss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

