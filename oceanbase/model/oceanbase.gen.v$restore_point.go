package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _V$restorePointMgr struct {
	*_BaseMgr
}

// V$restorePointMgr open func
func V$restorePointMgr(db *gorm.DB) *_V$restorePointMgr {
	if db == nil {
		panic(fmt.Errorf("V$restorePointMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$restorePointMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$restore_point"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$restorePointMgr) GetTableName() string {
	return "v$restore_point"
}

// Reset 重置gorm会话
func (obj *_V$restorePointMgr) Reset() *_V$restorePointMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$restorePointMgr) Get() (result V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$restorePointMgr) Gets() (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$restorePointMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_V$restorePointMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithSnapshot SNAPSHOT获取 
func (obj *_V$restorePointMgr) WithSnapshot(snapshot int64) Option {
	return optionFunc(func(o *options) { o.query["SNAPSHOT"] = snapshot })
}

// WithTime TIME获取 
func (obj *_V$restorePointMgr) WithTime(time time.Time) Option {
	return optionFunc(func(o *options) { o.query["TIME"] = time })
}

// WithName NAME获取 
func (obj *_V$restorePointMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_V$restorePointMgr) GetByOption(opts ...Option) (result V$restorePoint, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$restorePointMgr) GetByOptions(opts ...Option) (results []*V$restorePoint, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$restorePointMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$restorePoint,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where(options.query)
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


// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_V$restorePointMgr) GetFromTenantID(tenantID int64) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$restorePointMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSnapshot 通过SNAPSHOT获取内容  
func (obj *_V$restorePointMgr) GetFromSnapshot(snapshot int64) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`SNAPSHOT` = ?", snapshot).Find(&results).Error
	
	return
}

// GetBatchFromSnapshot 批量查找 
func (obj *_V$restorePointMgr) GetBatchFromSnapshot(snapshots []int64) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`SNAPSHOT` IN (?)", snapshots).Find(&results).Error
	
	return
}
 
// GetFromTime 通过TIME获取内容  
func (obj *_V$restorePointMgr) GetFromTime(time time.Time) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`TIME` = ?", time).Find(&results).Error
	
	return
}

// GetBatchFromTime 批量查找 
func (obj *_V$restorePointMgr) GetBatchFromTime(times []time.Time) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`TIME` IN (?)", times).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_V$restorePointMgr) GetFromName(name string) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_V$restorePointMgr) GetBatchFromName(names []string) (results []*V$restorePoint, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$restorePoint{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

