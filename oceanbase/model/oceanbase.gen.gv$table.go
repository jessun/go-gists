package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$tableMgr struct {
	*_BaseMgr
}

// Gv$tableMgr open func
func Gv$tableMgr(db *gorm.DB) *_Gv$tableMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$tableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$tableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$table"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$tableMgr) GetTableName() string {
	return "gv$table"
}

// Reset 重置gorm会话
func (obj *_Gv$tableMgr) Reset() *_Gv$tableMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$tableMgr) Get() (result Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$tableMgr) Gets() (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$tableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_Gv$tableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取 
func (obj *_Gv$tableMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithTableID table_id获取 
func (obj *_Gv$tableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTableName table_name获取 
func (obj *_Gv$tableMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithDatabaseID database_id获取 
func (obj *_Gv$tableMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithDatabaseName database_name获取 
func (obj *_Gv$tableMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTablegroupID tablegroup_id获取 
func (obj *_Gv$tableMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithTablegroupName tablegroup_name获取 
func (obj *_Gv$tableMgr) WithTablegroupName(tablegroupName string) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_name"] = tablegroupName })
}

// WithTableType table_type获取 
func (obj *_Gv$tableMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithZoneList zone_list获取 
func (obj *_Gv$tableMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取 
func (obj *_Gv$tableMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithCollationType collation_type获取 
func (obj *_Gv$tableMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithLocality locality获取 
func (obj *_Gv$tableMgr) WithLocality(locality string) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}

// WithSchemaVersion schema_version获取 
func (obj *_Gv$tableMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithReadOnly read_only获取 
func (obj *_Gv$tableMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithComment comment获取 
func (obj *_Gv$tableMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithIndexStatus index_status获取 
func (obj *_Gv$tableMgr) WithIndexStatus(indexStatus int64) Option {
	return optionFunc(func(o *options) { o.query["index_status"] = indexStatus })
}

// WithIndexType index_type获取 
func (obj *_Gv$tableMgr) WithIndexType(indexType int64) Option {
	return optionFunc(func(o *options) { o.query["index_type"] = indexType })
}

// WithPartLevel part_level获取 
func (obj *_Gv$tableMgr) WithPartLevel(partLevel int64) Option {
	return optionFunc(func(o *options) { o.query["part_level"] = partLevel })
}

// WithPartFuncType part_func_type获取 
func (obj *_Gv$tableMgr) WithPartFuncType(partFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["part_func_type"] = partFuncType })
}

// WithPartFuncExpr part_func_expr获取 
func (obj *_Gv$tableMgr) WithPartFuncExpr(partFuncExpr string) Option {
	return optionFunc(func(o *options) { o.query["part_func_expr"] = partFuncExpr })
}

// WithPartNum part_num获取 
func (obj *_Gv$tableMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithSubPartFuncType sub_part_func_type获取 
func (obj *_Gv$tableMgr) WithSubPartFuncType(subPartFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_type"] = subPartFuncType })
}

// WithSubPartFuncExpr sub_part_func_expr获取 
func (obj *_Gv$tableMgr) WithSubPartFuncExpr(subPartFuncExpr string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_expr"] = subPartFuncExpr })
}

// WithSubPartNum sub_part_num获取 
func (obj *_Gv$tableMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithDop dop获取 
func (obj *_Gv$tableMgr) WithDop(dop int64) Option {
	return optionFunc(func(o *options) { o.query["dop"] = dop })
}

// WithAutoPart auto_part获取 
func (obj *_Gv$tableMgr) WithAutoPart(autoPart int8) Option {
	return optionFunc(func(o *options) { o.query["auto_part"] = autoPart })
}

// WithAutoPartSize auto_part_size获取 
func (obj *_Gv$tableMgr) WithAutoPartSize(autoPartSize int64) Option {
	return optionFunc(func(o *options) { o.query["auto_part_size"] = autoPartSize })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$tableMgr) GetByOption(opts ...Option) (result Gv$table, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$tableMgr) GetByOptions(opts ...Option) (results []*Gv$table, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$tableMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$table,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where(options.query)
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
func (obj *_Gv$tableMgr) GetFromTenantID(tenantID int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantName 通过tenant_name获取内容  
func (obj *_Gv$tableMgr) GetFromTenantName(tenantName string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error
	
	return
}

// GetBatchFromTenantName 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromTenantName(tenantNames []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过table_id获取内容  
func (obj *_Gv$tableMgr) GetFromTableID(tableID int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`table_id` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromTableID(tableIDs []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromTableName 通过table_name获取内容  
func (obj *_Gv$tableMgr) GetFromTableName(tableName string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`table_name` = ?", tableName).Find(&results).Error
	
	return
}

// GetBatchFromTableName 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromTableName(tableNames []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error
	
	return
}
 
// GetFromDatabaseID 通过database_id获取内容  
func (obj *_Gv$tableMgr) GetFromDatabaseID(databaseID int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`database_id` = ?", databaseID).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseID 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error
	
	return
}
 
// GetFromDatabaseName 通过database_name获取内容  
func (obj *_Gv$tableMgr) GetFromDatabaseName(databaseName string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`database_name` = ?", databaseName).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseName 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error
	
	return
}
 
// GetFromTablegroupID 通过tablegroup_id获取内容  
func (obj *_Gv$tableMgr) GetFromTablegroupID(tablegroupID int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error
	
	return
}

// GetBatchFromTablegroupID 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error
	
	return
}
 
// GetFromTablegroupName 通过tablegroup_name获取内容  
func (obj *_Gv$tableMgr) GetFromTablegroupName(tablegroupName string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tablegroup_name` = ?", tablegroupName).Find(&results).Error
	
	return
}

// GetBatchFromTablegroupName 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromTablegroupName(tablegroupNames []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`tablegroup_name` IN (?)", tablegroupNames).Find(&results).Error
	
	return
}
 
// GetFromTableType 通过table_type获取内容  
func (obj *_Gv$tableMgr) GetFromTableType(tableType int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`table_type` = ?", tableType).Find(&results).Error
	
	return
}

// GetBatchFromTableType 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromTableType(tableTypes []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error
	
	return
}
 
// GetFromZoneList 通过zone_list获取内容  
func (obj *_Gv$tableMgr) GetFromZoneList(zoneList string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`zone_list` = ?", zoneList).Find(&results).Error
	
	return
}

// GetBatchFromZoneList 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromZoneList(zoneLists []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error
	
	return
}
 
// GetFromPrimaryZone 通过primary_zone获取内容  
func (obj *_Gv$tableMgr) GetFromPrimaryZone(primaryZone string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error
	
	return
}

// GetBatchFromPrimaryZone 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error
	
	return
}
 
// GetFromCollationType 通过collation_type获取内容  
func (obj *_Gv$tableMgr) GetFromCollationType(collationType int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`collation_type` = ?", collationType).Find(&results).Error
	
	return
}

// GetBatchFromCollationType 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromCollationType(collationTypes []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error
	
	return
}
 
// GetFromLocality 通过locality获取内容  
func (obj *_Gv$tableMgr) GetFromLocality(locality string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`locality` = ?", locality).Find(&results).Error
	
	return
}

// GetBatchFromLocality 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromLocality(localitys []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`locality` IN (?)", localitys).Find(&results).Error
	
	return
}
 
// GetFromSchemaVersion 通过schema_version获取内容  
func (obj *_Gv$tableMgr) GetFromSchemaVersion(schemaVersion int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromSchemaVersion 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error
	
	return
}
 
// GetFromReadOnly 通过read_only获取内容  
func (obj *_Gv$tableMgr) GetFromReadOnly(readOnly int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`read_only` = ?", readOnly).Find(&results).Error
	
	return
}

// GetBatchFromReadOnly 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error
	
	return
}
 
// GetFromComment 通过comment获取内容  
func (obj *_Gv$tableMgr) GetFromComment(comment string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`comment` = ?", comment).Find(&results).Error
	
	return
}

// GetBatchFromComment 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromComment(comments []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`comment` IN (?)", comments).Find(&results).Error
	
	return
}
 
// GetFromIndexStatus 通过index_status获取内容  
func (obj *_Gv$tableMgr) GetFromIndexStatus(indexStatus int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`index_status` = ?", indexStatus).Find(&results).Error
	
	return
}

// GetBatchFromIndexStatus 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromIndexStatus(indexStatuss []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`index_status` IN (?)", indexStatuss).Find(&results).Error
	
	return
}
 
// GetFromIndexType 通过index_type获取内容  
func (obj *_Gv$tableMgr) GetFromIndexType(indexType int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`index_type` = ?", indexType).Find(&results).Error
	
	return
}

// GetBatchFromIndexType 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromIndexType(indexTypes []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`index_type` IN (?)", indexTypes).Find(&results).Error
	
	return
}
 
// GetFromPartLevel 通过part_level获取内容  
func (obj *_Gv$tableMgr) GetFromPartLevel(partLevel int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_level` = ?", partLevel).Find(&results).Error
	
	return
}

// GetBatchFromPartLevel 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromPartLevel(partLevels []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_level` IN (?)", partLevels).Find(&results).Error
	
	return
}
 
// GetFromPartFuncType 通过part_func_type获取内容  
func (obj *_Gv$tableMgr) GetFromPartFuncType(partFuncType int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_func_type` = ?", partFuncType).Find(&results).Error
	
	return
}

// GetBatchFromPartFuncType 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromPartFuncType(partFuncTypes []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_func_type` IN (?)", partFuncTypes).Find(&results).Error
	
	return
}
 
// GetFromPartFuncExpr 通过part_func_expr获取内容  
func (obj *_Gv$tableMgr) GetFromPartFuncExpr(partFuncExpr string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_func_expr` = ?", partFuncExpr).Find(&results).Error
	
	return
}

// GetBatchFromPartFuncExpr 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromPartFuncExpr(partFuncExprs []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_func_expr` IN (?)", partFuncExprs).Find(&results).Error
	
	return
}
 
// GetFromPartNum 通过part_num获取内容  
func (obj *_Gv$tableMgr) GetFromPartNum(partNum int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_num` = ?", partNum).Find(&results).Error
	
	return
}

// GetBatchFromPartNum 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromPartNum(partNums []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`part_num` IN (?)", partNums).Find(&results).Error
	
	return
}
 
// GetFromSubPartFuncType 通过sub_part_func_type获取内容  
func (obj *_Gv$tableMgr) GetFromSubPartFuncType(subPartFuncType int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`sub_part_func_type` = ?", subPartFuncType).Find(&results).Error
	
	return
}

// GetBatchFromSubPartFuncType 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromSubPartFuncType(subPartFuncTypes []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`sub_part_func_type` IN (?)", subPartFuncTypes).Find(&results).Error
	
	return
}
 
// GetFromSubPartFuncExpr 通过sub_part_func_expr获取内容  
func (obj *_Gv$tableMgr) GetFromSubPartFuncExpr(subPartFuncExpr string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`sub_part_func_expr` = ?", subPartFuncExpr).Find(&results).Error
	
	return
}

// GetBatchFromSubPartFuncExpr 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromSubPartFuncExpr(subPartFuncExprs []string) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`sub_part_func_expr` IN (?)", subPartFuncExprs).Find(&results).Error
	
	return
}
 
// GetFromSubPartNum 通过sub_part_num获取内容  
func (obj *_Gv$tableMgr) GetFromSubPartNum(subPartNum int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error
	
	return
}

// GetBatchFromSubPartNum 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error
	
	return
}
 
// GetFromDop 通过dop获取内容  
func (obj *_Gv$tableMgr) GetFromDop(dop int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`dop` = ?", dop).Find(&results).Error
	
	return
}

// GetBatchFromDop 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromDop(dops []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`dop` IN (?)", dops).Find(&results).Error
	
	return
}
 
// GetFromAutoPart 通过auto_part获取内容  
func (obj *_Gv$tableMgr) GetFromAutoPart(autoPart int8) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`auto_part` = ?", autoPart).Find(&results).Error
	
	return
}

// GetBatchFromAutoPart 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromAutoPart(autoParts []int8) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`auto_part` IN (?)", autoParts).Find(&results).Error
	
	return
}
 
// GetFromAutoPartSize 通过auto_part_size获取内容  
func (obj *_Gv$tableMgr) GetFromAutoPartSize(autoPartSize int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`auto_part_size` = ?", autoPartSize).Find(&results).Error
	
	return
}

// GetBatchFromAutoPartSize 批量查找 
func (obj *_Gv$tableMgr) GetBatchFromAutoPartSize(autoPartSizes []int64) (results []*Gv$table, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$table{}).Where("`auto_part_size` IN (?)", autoPartSizes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

