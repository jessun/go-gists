package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AllVirtualRootserviceStatMgr struct {
	*_BaseMgr
}

// AllVirtualRootserviceStatMgr open func
func AllVirtualRootserviceStatMgr(db *gorm.DB) *_AllVirtualRootserviceStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRootserviceStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRootserviceStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rootservice_stat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRootserviceStatMgr) GetTableName() string {
	return "__all_virtual_rootservice_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRootserviceStatMgr) Reset() *_AllVirtualRootserviceStatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllVirtualRootserviceStatMgr) Get() (result AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRootserviceStatMgr) Gets() (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRootserviceStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithStatistic# statistic#获取 
func (obj *_AllVirtualRootserviceStatMgr) WithStatistic#(statistic# int64) Option {
	return optionFunc(func(o *options) { o.query["statistic#"] = statistic# })
}

// WithValue value获取 
func (obj *_AllVirtualRootserviceStatMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithStatID stat_id获取 
func (obj *_AllVirtualRootserviceStatMgr) WithStatID(statID int64) Option {
	return optionFunc(func(o *options) { o.query["stat_id"] = statID })
}

// WithName name获取 
func (obj *_AllVirtualRootserviceStatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithClass class获取 
func (obj *_AllVirtualRootserviceStatMgr) WithClass(class int64) Option {
	return optionFunc(func(o *options) { o.query["class"] = class })
}

// WithCanVisible can_visible获取 
func (obj *_AllVirtualRootserviceStatMgr) WithCanVisible(canVisible int8) Option {
	return optionFunc(func(o *options) { o.query["can_visible"] = canVisible })
}


// GetByOption 功能选项模式获取
func (obj *_AllVirtualRootserviceStatMgr) GetByOption(opts ...Option) (result AllVirtualRootserviceStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRootserviceStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRootserviceStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllVirtualRootserviceStatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRootserviceStat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where(options.query)
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


// GetFromStatistic# 通过statistic#获取内容  
func (obj *_AllVirtualRootserviceStatMgr) GetFromStatistic#(statistic# int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`statistic#` = ?", statistic#).Find(&results).Error
	
	return
}

// GetBatchFromStatistic# 批量查找 
func (obj *_AllVirtualRootserviceStatMgr) GetBatchFromStatistic#(statistic#s []int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`statistic#` IN (?)", statistic#s).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_AllVirtualRootserviceStatMgr) GetFromValue(value int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`value` = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量查找 
func (obj *_AllVirtualRootserviceStatMgr) GetBatchFromValue(values []int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`value` IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromStatID 通过stat_id获取内容  
func (obj *_AllVirtualRootserviceStatMgr) GetFromStatID(statID int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`stat_id` = ?", statID).Find(&results).Error
	
	return
}

// GetBatchFromStatID 批量查找 
func (obj *_AllVirtualRootserviceStatMgr) GetBatchFromStatID(statIDs []int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`stat_id` IN (?)", statIDs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_AllVirtualRootserviceStatMgr) GetFromName(name string) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_AllVirtualRootserviceStatMgr) GetBatchFromName(names []string) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromClass 通过class获取内容  
func (obj *_AllVirtualRootserviceStatMgr) GetFromClass(class int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`class` = ?", class).Find(&results).Error
	
	return
}

// GetBatchFromClass 批量查找 
func (obj *_AllVirtualRootserviceStatMgr) GetBatchFromClass(classs []int64) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`class` IN (?)", classs).Find(&results).Error
	
	return
}
 
// GetFromCanVisible 通过can_visible获取内容  
func (obj *_AllVirtualRootserviceStatMgr) GetFromCanVisible(canVisible int8) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`can_visible` = ?", canVisible).Find(&results).Error
	
	return
}

// GetBatchFromCanVisible 批量查找 
func (obj *_AllVirtualRootserviceStatMgr) GetBatchFromCanVisible(canVisibles []int8) (results []*AllVirtualRootserviceStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRootserviceStat{}).Where("`can_visible` IN (?)", canVisibles).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

