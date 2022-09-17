package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _TenantVirtualStatnameMgr struct {
	*_BaseMgr
}

// TenantVirtualStatnameMgr open func
func TenantVirtualStatnameMgr(db *gorm.DB) *_TenantVirtualStatnameMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualStatnameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualStatnameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_statname"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualStatnameMgr) GetTableName() string {
	return "__tenant_virtual_statname"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualStatnameMgr) Reset() *_TenantVirtualStatnameMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_TenantVirtualStatnameMgr) Get() (result TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualStatnameMgr) Gets() (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualStatnameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_TenantVirtualStatnameMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithStatID stat_id获取 
func (obj *_TenantVirtualStatnameMgr) WithStatID(statID int64) Option {
	return optionFunc(func(o *options) { o.query["stat_id"] = statID })
}

// WithStatistic# statistic#获取 
func (obj *_TenantVirtualStatnameMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["statistic#"] = statistic# })
}

// WithName name获取 
func (obj *_TenantVirtualStatnameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDisplayName display_name获取 
func (obj *_TenantVirtualStatnameMgr) WithDisplayName(displayName string) Option {
	return optionFunc(func(o *options) { o.query["display_name"] = displayName })
}

// WithClass class获取 
func (obj *_TenantVirtualStatnameMgr) WithClass(class int64) Option {
	return optionFunc(func(o *options) { o.query["class"] = class })
}


// GetByOption 功能选项模式获取
func (obj *_TenantVirtualStatnameMgr) GetByOption(opts ...Option) (result TenantVirtualStatname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualStatnameMgr) GetByOptions(opts ...Option) (results []*TenantVirtualStatname, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_TenantVirtualStatnameMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualStatname,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where(options.query)
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
func (obj *_TenantVirtualStatnameMgr) GetFromTenantID(tenantID int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_TenantVirtualStatnameMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromStatID 通过stat_id获取内容  
func (obj *_TenantVirtualStatnameMgr) GetFromStatID(statID int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`stat_id` = ?", statID).Find(&results).Error
	
	return
}

// GetBatchFromStatID 批量查找 
func (obj *_TenantVirtualStatnameMgr) GetBatchFromStatID(statIDs []int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`stat_id` IN (?)", statIDs).Find(&results).Error
	
	return
}
 
// GetFromStatistic# 通过statistic#获取内容  
func (obj *_TenantVirtualStatnameMgr) GetFromStatistic#(statistic# int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`statistic#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_TenantVirtualStatnameMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`statistic#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_TenantVirtualStatnameMgr) GetFromName(name string) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_TenantVirtualStatnameMgr) GetBatchFromName(names []string) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDisplayName 通过display_name获取内容  
func (obj *_TenantVirtualStatnameMgr) GetFromDisplayName(displayName string) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`display_name` = ?", displayName).Find(&results).Error
	
	return
}

// GetBatchFromDisplayName 批量查找 
func (obj *_TenantVirtualStatnameMgr) GetBatchFromDisplayName(displayNames []string) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`display_name` IN (?)", displayNames).Find(&results).Error
	
	return
}
 
// GetFromClass 通过class获取内容  
func (obj *_TenantVirtualStatnameMgr) GetFromClass(class int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`class` = ?", class).Find(&results).Error
	
	return
}

// GetBatchFromClass 批量查找 
func (obj *_TenantVirtualStatnameMgr) GetBatchFromClass(classs []int64) (results []*TenantVirtualStatname, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualStatname{}).Where("`class` IN (?)", classs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

