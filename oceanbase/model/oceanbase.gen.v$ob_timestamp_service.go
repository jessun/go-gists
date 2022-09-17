package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$obTimestampServiceMgr struct {
	*_BaseMgr
}

// V$obTimestampServiceMgr open func
func V$obTimestampServiceMgr(db *gorm.DB) *_V$obTimestampServiceMgr {
	if db == nil {
		panic(fmt.Errorf("V$obTimestampServiceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$obTimestampServiceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$ob_timestamp_service"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$obTimestampServiceMgr) GetTableName() string {
	return "v$ob_timestamp_service"
}

// Reset 重置gorm会话
func (obj *_V$obTimestampServiceMgr) Reset() *_V$obTimestampServiceMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$obTimestampServiceMgr) Get() (result V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$obTimestampServiceMgr) Gets() (results []*V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$obTimestampServiceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_V$obTimestampServiceMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTsType ts_type获取 
func (obj *_V$obTimestampServiceMgr) WithTsType(tsType string) Option {
	return optionFunc(func(o *options) { o.query["ts_type"] = tsType })
}

// WithTsValue ts_value获取 
func (obj *_V$obTimestampServiceMgr) WithTsValue(tsValue int64) Option {
	return optionFunc(func(o *options) { o.query["ts_value"] = tsValue })
}


// GetByOption 功能选项模式获取
func (obj *_V$obTimestampServiceMgr) GetByOption(opts ...Option) (result V$obTimestampService, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$obTimestampServiceMgr) GetByOptions(opts ...Option) (results []*V$obTimestampService, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$obTimestampServiceMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$obTimestampService,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where(options.query)
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
func (obj *_V$obTimestampServiceMgr) GetFromTenantID(tenantID int64) (results []*V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$obTimestampServiceMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTsType 通过ts_type获取内容  
func (obj *_V$obTimestampServiceMgr) GetFromTsType(tsType string) (results []*V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where("`ts_type` = ?", tsType).Find(&results).Error
	
	return
}

// GetBatchFromTsType 批量查找 
func (obj *_V$obTimestampServiceMgr) GetBatchFromTsType(tsTypes []string) (results []*V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where("`ts_type` IN (?)", tsTypes).Find(&results).Error
	
	return
}
 
// GetFromTsValue 通过ts_value获取内容  
func (obj *_V$obTimestampServiceMgr) GetFromTsValue(tsValue int64) (results []*V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where("`ts_value` = ?", tsValue).Find(&results).Error
	
	return
}

// GetBatchFromTsValue 批量查找 
func (obj *_V$obTimestampServiceMgr) GetBatchFromTsValue(tsValues []int64) (results []*V$obTimestampService, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTimestampService{}).Where("`ts_value` IN (?)", tsValues).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

