package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sqlMonitorStatnameMgr struct {
	*_BaseMgr
}

// V$sqlMonitorStatnameMgr open func
func V$sqlMonitorStatnameMgr(db *gorm.DB) *_V$sqlMonitorStatnameMgr {
	if db == nil {
		panic(fmt.Errorf("V$sqlMonitorStatnameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sqlMonitorStatnameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sql_monitor_statname"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sqlMonitorStatnameMgr) GetTableName() string {
	return "v$sql_monitor_statname"
}

// Reset 重置gorm会话
func (obj *_V$sqlMonitorStatnameMgr) Reset() *_V$sqlMonitorStatnameMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sqlMonitorStatnameMgr) Get() (result V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sqlMonitorStatnameMgr) Gets() (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sqlMonitorStatnameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$sqlMonitorStatnameMgr) WithConID(conID string) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithID ID获取 
func (obj *_V$sqlMonitorStatnameMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithGroupID GROUP_ID获取 
func (obj *_V$sqlMonitorStatnameMgr) WithGroupID(groupID int64) Option {
	return optionFunc(func(o *options) { o.query["GROUP_ID"] = groupID })
}

// WithName NAME获取 
func (obj *_V$sqlMonitorStatnameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithDescription DESCRIPTION获取 
func (obj *_V$sqlMonitorStatnameMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["DESCRIPTION"] = description })
}

// WithType TYPE获取 
func (obj *_V$sqlMonitorStatnameMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithFlags FLAGS获取 
func (obj *_V$sqlMonitorStatnameMgr) WithFlags(flags int64) Option {
	return optionFunc(func(o *options) { o.query["FLAGS"] = flags })
}


// GetByOption 功能选项模式获取
func (obj *_V$sqlMonitorStatnameMgr) GetByOption(opts ...Option) (result V$sqlMonitorStatname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sqlMonitorStatnameMgr) GetByOptions(opts ...Option) (results []*V$sqlMonitorStatname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sqlMonitorStatnameMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sqlMonitorStatname,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where(options.query)
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
func (obj *_V$sqlMonitorStatnameMgr) GetFromConID(conID string) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sqlMonitorStatnameMgr) GetBatchFromConID(conIDs []string) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromID 通过ID获取内容  
func (obj *_V$sqlMonitorStatnameMgr) GetFromID(id int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`ID` = ?", id).Find(&results).Error
	
	return
}

// GetBatchFromID 批量查找 
func (obj *_V$sqlMonitorStatnameMgr) GetBatchFromID(ids []int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`ID` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromGroupID 通过GROUP_ID获取内容  
func (obj *_V$sqlMonitorStatnameMgr) GetFromGroupID(groupID int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`GROUP_ID` = ?", groupID).Find(&results).Error
	
	return
}

// GetBatchFromGroupID 批量查找 
func (obj *_V$sqlMonitorStatnameMgr) GetBatchFromGroupID(groupIDs []int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`GROUP_ID` IN (?)", groupIDs).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_V$sqlMonitorStatnameMgr) GetFromName(name string) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_V$sqlMonitorStatnameMgr) GetBatchFromName(names []string) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过DESCRIPTION获取内容  
func (obj *_V$sqlMonitorStatnameMgr) GetFromDescription(description string) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`DESCRIPTION` = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量查找 
func (obj *_V$sqlMonitorStatnameMgr) GetBatchFromDescription(descriptions []string) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`DESCRIPTION` IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromType 通过TYPE获取内容  
func (obj *_V$sqlMonitorStatnameMgr) GetFromType(_type int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`TYPE` = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量查找 
func (obj *_V$sqlMonitorStatnameMgr) GetBatchFromType(_types []int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`TYPE` IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromFlags 通过FLAGS获取内容  
func (obj *_V$sqlMonitorStatnameMgr) GetFromFlags(flags int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`FLAGS` = ?", flags).Find(&results).Error
	
	return
}

// GetBatchFromFlags 批量查找 
func (obj *_V$sqlMonitorStatnameMgr) GetBatchFromFlags(flagss []int64) (results []*V$sqlMonitorStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlMonitorStatname{}).Where("`FLAGS` IN (?)", flagss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

