package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$outlineMgr struct {
	*_BaseMgr
}

// Gv$outlineMgr open func
func Gv$outlineMgr(db *gorm.DB) *_Gv$outlineMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$outlineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$outlineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$outline"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$outlineMgr) GetTableName() string {
	return "gv$outline"
}

// Reset 重置gorm会话
func (obj *_Gv$outlineMgr) Reset() *_Gv$outlineMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$outlineMgr) Get() (result Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$outlineMgr) Gets() (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$outlineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_Gv$outlineMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取 
func (obj *_Gv$outlineMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithOutlineID outline_id获取 
func (obj *_Gv$outlineMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithDatabaseName database_name获取 
func (obj *_Gv$outlineMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithOutlineName outline_name获取 
func (obj *_Gv$outlineMgr) WithOutlineName(outlineName string) Option {
	return optionFunc(func(o *options) { o.query["outline_name"] = outlineName })
}

// WithVisibleSignature visible_signature获取 
func (obj *_Gv$outlineMgr) WithVisibleSignature(visibleSignature string) Option {
	return optionFunc(func(o *options) { o.query["visible_signature"] = visibleSignature })
}

// WithSQLText sql_text获取 
func (obj *_Gv$outlineMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithOutlineTarget outline_target获取 
func (obj *_Gv$outlineMgr) WithOutlineTarget(outlineTarget string) Option {
	return optionFunc(func(o *options) { o.query["outline_target"] = outlineTarget })
}

// WithOutlineSQL outline_sql获取 
func (obj *_Gv$outlineMgr) WithOutlineSQL(outlineSQL string) Option {
	return optionFunc(func(o *options) { o.query["outline_sql"] = outlineSQL })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$outlineMgr) GetByOption(opts ...Option) (result Gv$outline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$outlineMgr) GetByOptions(opts ...Option) (results []*Gv$outline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$outlineMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$outline,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where(options.query)
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
func (obj *_Gv$outlineMgr) GetFromTenantID(tenantID int64) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromDatabaseID 通过database_id获取内容  
func (obj *_Gv$outlineMgr) GetFromDatabaseID(databaseID int64) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`database_id` = ?", databaseID).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseID 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error
	
	return
}
 
// GetFromOutlineID 通过outline_id获取内容  
func (obj *_Gv$outlineMgr) GetFromOutlineID(outlineID int64) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_id` = ?", outlineID).Find(&results).Error
	
	return
}

// GetBatchFromOutlineID 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error
	
	return
}
 
// GetFromDatabaseName 通过database_name获取内容  
func (obj *_Gv$outlineMgr) GetFromDatabaseName(databaseName string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`database_name` = ?", databaseName).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseName 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error
	
	return
}
 
// GetFromOutlineName 通过outline_name获取内容  
func (obj *_Gv$outlineMgr) GetFromOutlineName(outlineName string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_name` = ?", outlineName).Find(&results).Error
	
	return
}

// GetBatchFromOutlineName 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromOutlineName(outlineNames []string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_name` IN (?)", outlineNames).Find(&results).Error
	
	return
}
 
// GetFromVisibleSignature 通过visible_signature获取内容  
func (obj *_Gv$outlineMgr) GetFromVisibleSignature(visibleSignature string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`visible_signature` = ?", visibleSignature).Find(&results).Error
	
	return
}

// GetBatchFromVisibleSignature 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromVisibleSignature(visibleSignatures []string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`visible_signature` IN (?)", visibleSignatures).Find(&results).Error
	
	return
}
 
// GetFromSQLText 通过sql_text获取内容  
func (obj *_Gv$outlineMgr) GetFromSQLText(sqlText string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`sql_text` = ?", sqlText).Find(&results).Error
	
	return
}

// GetBatchFromSQLText 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromSQLText(sqlTexts []string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error
	
	return
}
 
// GetFromOutlineTarget 通过outline_target获取内容  
func (obj *_Gv$outlineMgr) GetFromOutlineTarget(outlineTarget string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_target` = ?", outlineTarget).Find(&results).Error
	
	return
}

// GetBatchFromOutlineTarget 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromOutlineTarget(outlineTargets []string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_target` IN (?)", outlineTargets).Find(&results).Error
	
	return
}
 
// GetFromOutlineSQL 通过outline_sql获取内容  
func (obj *_Gv$outlineMgr) GetFromOutlineSQL(outlineSQL string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_sql` = ?", outlineSQL).Find(&results).Error
	
	return
}

// GetBatchFromOutlineSQL 批量查找 
func (obj *_Gv$outlineMgr) GetBatchFromOutlineSQL(outlineSQLs []string) (results []*Gv$outline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$outline{}).Where("`outline_sql` IN (?)", outlineSQLs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

