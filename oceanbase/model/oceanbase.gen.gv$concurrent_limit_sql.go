package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$concurrentLimitSQLMgr struct {
	*_BaseMgr
}

// Gv$concurrentLimitSQLMgr open func
func Gv$concurrentLimitSQLMgr(db *gorm.DB) *_Gv$concurrentLimitSQLMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$concurrentLimitSQLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$concurrentLimitSQLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$concurrent_limit_sql"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$concurrentLimitSQLMgr) GetTableName() string {
	return "gv$concurrent_limit_sql"
}

// Reset 重置gorm会话
func (obj *_Gv$concurrentLimitSQLMgr) Reset() *_Gv$concurrentLimitSQLMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$concurrentLimitSQLMgr) Get() (result Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$concurrentLimitSQLMgr) Gets() (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$concurrentLimitSQLMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithOutlineID outline_id获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithDatabaseName database_name获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithOutlineName outline_name获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithOutlineName(outlineName string) Option {
	return optionFunc(func(o *options) { o.query["outline_name"] = outlineName })
}

// WithOutlineContent outline_content获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithOutlineContent(outlineContent string) Option {
	return optionFunc(func(o *options) { o.query["outline_content"] = outlineContent })
}

// WithVisibleSignature visible_signature获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithVisibleSignature(visibleSignature string) Option {
	return optionFunc(func(o *options) { o.query["visible_signature"] = visibleSignature })
}

// WithSQLText sql_text获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithConcurrentNum concurrent_num获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithConcurrentNum(concurrentNum int64) Option {
	return optionFunc(func(o *options) { o.query["concurrent_num"] = concurrentNum })
}

// WithLimitTarget limit_target获取 
func (obj *_Gv$concurrentLimitSQLMgr) WithLimitTarget(limitTarget string) Option {
	return optionFunc(func(o *options) { o.query["limit_target"] = limitTarget })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$concurrentLimitSQLMgr) GetByOption(opts ...Option) (result Gv$concurrentLimitSQL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$concurrentLimitSQLMgr) GetByOptions(opts ...Option) (results []*Gv$concurrentLimitSQL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$concurrentLimitSQLMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$concurrentLimitSQL,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where(options.query)
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
func (obj *_Gv$concurrentLimitSQLMgr) GetFromTenantID(tenantID int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromDatabaseID 通过database_id获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromDatabaseID(databaseID int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`database_id` = ?", databaseID).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseID 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error
	
	return
}
 
// GetFromOutlineID 通过outline_id获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromOutlineID(outlineID int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`outline_id` = ?", outlineID).Find(&results).Error
	
	return
}

// GetBatchFromOutlineID 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error
	
	return
}
 
// GetFromDatabaseName 通过database_name获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromDatabaseName(databaseName string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`database_name` = ?", databaseName).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseName 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error
	
	return
}
 
// GetFromOutlineName 通过outline_name获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromOutlineName(outlineName string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`outline_name` = ?", outlineName).Find(&results).Error
	
	return
}

// GetBatchFromOutlineName 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromOutlineName(outlineNames []string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`outline_name` IN (?)", outlineNames).Find(&results).Error
	
	return
}
 
// GetFromOutlineContent 通过outline_content获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromOutlineContent(outlineContent string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`outline_content` = ?", outlineContent).Find(&results).Error
	
	return
}

// GetBatchFromOutlineContent 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromOutlineContent(outlineContents []string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`outline_content` IN (?)", outlineContents).Find(&results).Error
	
	return
}
 
// GetFromVisibleSignature 通过visible_signature获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromVisibleSignature(visibleSignature string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`visible_signature` = ?", visibleSignature).Find(&results).Error
	
	return
}

// GetBatchFromVisibleSignature 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromVisibleSignature(visibleSignatures []string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`visible_signature` IN (?)", visibleSignatures).Find(&results).Error
	
	return
}
 
// GetFromSQLText 通过sql_text获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromSQLText(sqlText string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`sql_text` = ?", sqlText).Find(&results).Error
	
	return
}

// GetBatchFromSQLText 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromSQLText(sqlTexts []string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error
	
	return
}
 
// GetFromConcurrentNum 通过concurrent_num获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromConcurrentNum(concurrentNum int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`concurrent_num` = ?", concurrentNum).Find(&results).Error
	
	return
}

// GetBatchFromConcurrentNum 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromConcurrentNum(concurrentNums []int64) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`concurrent_num` IN (?)", concurrentNums).Find(&results).Error
	
	return
}
 
// GetFromLimitTarget 通过limit_target获取内容  
func (obj *_Gv$concurrentLimitSQLMgr) GetFromLimitTarget(limitTarget string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`limit_target` = ?", limitTarget).Find(&results).Error
	
	return
}

// GetBatchFromLimitTarget 批量查找 
func (obj *_Gv$concurrentLimitSQLMgr) GetBatchFromLimitTarget(limitTargets []string) (results []*Gv$concurrentLimitSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$concurrentLimitSQL{}).Where("`limit_target` IN (?)", limitTargets).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

