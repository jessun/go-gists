package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualProxySchemaMgr struct {
	*_BaseMgr
}

// AllVirtualProxySchemaMgr open func
func AllVirtualProxySchemaMgr(db *gorm.DB) *_AllVirtualProxySchemaMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProxySchemaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProxySchemaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_proxy_schema"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProxySchemaMgr) GetTableName() string {
	return "__all_virtual_proxy_schema"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProxySchemaMgr) Reset() *_AllVirtualProxySchemaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProxySchemaMgr) Get() (result AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProxySchemaMgr) Gets() (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProxySchemaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantName tenant_name获取
func (obj *_AllVirtualProxySchemaMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualProxySchemaMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTableName table_name获取
func (obj *_AllVirtualProxySchemaMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualProxySchemaMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualProxySchemaMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSQLPort sql_port获取
func (obj *_AllVirtualProxySchemaMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualProxySchemaMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithRole role获取
func (obj *_AllVirtualProxySchemaMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithPartNum part_num获取
func (obj *_AllVirtualProxySchemaMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualProxySchemaMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithTableType table_type获取
func (obj *_AllVirtualProxySchemaMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualProxySchemaMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualProxySchemaMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualProxySchemaMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualProxySchemaMgr) WithSpare3(spare3 int64) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithSpare4 spare4获取
func (obj *_AllVirtualProxySchemaMgr) WithSpare4(spare4 string) Option {
	return optionFunc(func(o *options) { o.query["spare4"] = spare4 })
}

// WithSpare5 spare5获取
func (obj *_AllVirtualProxySchemaMgr) WithSpare5(spare5 string) Option {
	return optionFunc(func(o *options) { o.query["spare5"] = spare5 })
}

// WithSpare6 spare6获取
func (obj *_AllVirtualProxySchemaMgr) WithSpare6(spare6 string) Option {
	return optionFunc(func(o *options) { o.query["spare6"] = spare6 })
}

// WithComplexTableType complex_table_type获取
func (obj *_AllVirtualProxySchemaMgr) WithComplexTableType(complexTableType int64) Option {
	return optionFunc(func(o *options) { o.query["complex_table_type"] = complexTableType })
}

// WithLevel1DecodedDbName level1_decoded_db_name获取
func (obj *_AllVirtualProxySchemaMgr) WithLevel1DecodedDbName(level1DecodedDbName string) Option {
	return optionFunc(func(o *options) { o.query["level1_decoded_db_name"] = level1DecodedDbName })
}

// WithLevel1DecodedTableName level1_decoded_table_name获取
func (obj *_AllVirtualProxySchemaMgr) WithLevel1DecodedTableName(level1DecodedTableName string) Option {
	return optionFunc(func(o *options) { o.query["level1_decoded_table_name"] = level1DecodedTableName })
}

// WithLevel2DecodedDbName level2_decoded_db_name获取
func (obj *_AllVirtualProxySchemaMgr) WithLevel2DecodedDbName(level2DecodedDbName string) Option {
	return optionFunc(func(o *options) { o.query["level2_decoded_db_name"] = level2DecodedDbName })
}

// WithLevel2DecodedTableName level2_decoded_table_name获取
func (obj *_AllVirtualProxySchemaMgr) WithLevel2DecodedTableName(level2DecodedTableName string) Option {
	return optionFunc(func(o *options) { o.query["level2_decoded_table_name"] = level2DecodedTableName })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProxySchemaMgr) GetByOption(opts ...Option) (result AllVirtualProxySchema, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProxySchemaMgr) GetByOptions(opts ...Option) (results []*AllVirtualProxySchema, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProxySchemaMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProxySchema, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where(options.query)
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

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromTenantName(tenantName string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromTableName(tableName string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSQLPort 通过sql_port获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSQLPort(sqlPort int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error

	return
}

// GetBatchFromSQLPort 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromTableID(tableID int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromRole(role int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromPartNum 通过part_num获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromPartNum(partNum int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`part_num` = ?", partNum).Find(&results).Error

	return
}

// GetBatchFromPartNum 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromPartNum(partNums []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`part_num` IN (?)", partNums).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromTableType(tableType int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromTableType(tableTypes []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSpare3(spare3 int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSpare3(spare3s []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromSpare4 通过spare4获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSpare4(spare4 string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare4` = ?", spare4).Find(&results).Error

	return
}

// GetBatchFromSpare4 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSpare4(spare4s []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare4` IN (?)", spare4s).Find(&results).Error

	return
}

// GetFromSpare5 通过spare5获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSpare5(spare5 string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare5` = ?", spare5).Find(&results).Error

	return
}

// GetBatchFromSpare5 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSpare5(spare5s []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare5` IN (?)", spare5s).Find(&results).Error

	return
}

// GetFromSpare6 通过spare6获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromSpare6(spare6 string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare6` = ?", spare6).Find(&results).Error

	return
}

// GetBatchFromSpare6 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromSpare6(spare6s []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`spare6` IN (?)", spare6s).Find(&results).Error

	return
}

// GetFromComplexTableType 通过complex_table_type获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromComplexTableType(complexTableType int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`complex_table_type` = ?", complexTableType).Find(&results).Error

	return
}

// GetBatchFromComplexTableType 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromComplexTableType(complexTableTypes []int64) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`complex_table_type` IN (?)", complexTableTypes).Find(&results).Error

	return
}

// GetFromLevel1DecodedDbName 通过level1_decoded_db_name获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromLevel1DecodedDbName(level1DecodedDbName string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level1_decoded_db_name` = ?", level1DecodedDbName).Find(&results).Error

	return
}

// GetBatchFromLevel1DecodedDbName 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromLevel1DecodedDbName(level1DecodedDbNames []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level1_decoded_db_name` IN (?)", level1DecodedDbNames).Find(&results).Error

	return
}

// GetFromLevel1DecodedTableName 通过level1_decoded_table_name获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromLevel1DecodedTableName(level1DecodedTableName string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level1_decoded_table_name` = ?", level1DecodedTableName).Find(&results).Error

	return
}

// GetBatchFromLevel1DecodedTableName 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromLevel1DecodedTableName(level1DecodedTableNames []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level1_decoded_table_name` IN (?)", level1DecodedTableNames).Find(&results).Error

	return
}

// GetFromLevel2DecodedDbName 通过level2_decoded_db_name获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromLevel2DecodedDbName(level2DecodedDbName string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level2_decoded_db_name` = ?", level2DecodedDbName).Find(&results).Error

	return
}

// GetBatchFromLevel2DecodedDbName 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromLevel2DecodedDbName(level2DecodedDbNames []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level2_decoded_db_name` IN (?)", level2DecodedDbNames).Find(&results).Error

	return
}

// GetFromLevel2DecodedTableName 通过level2_decoded_table_name获取内容
func (obj *_AllVirtualProxySchemaMgr) GetFromLevel2DecodedTableName(level2DecodedTableName string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level2_decoded_table_name` = ?", level2DecodedTableName).Find(&results).Error

	return
}

// GetBatchFromLevel2DecodedTableName 批量查找
func (obj *_AllVirtualProxySchemaMgr) GetBatchFromLevel2DecodedTableName(level2DecodedTableNames []string) (results []*AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`level2_decoded_table_name` IN (?)", level2DecodedTableNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualProxySchemaMgr) FetchByPrimaryKey(tenantName string, databaseName string, tableName string, partitionID int64, svrIP string, sqlPort int64) (result AllVirtualProxySchema, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySchema{}).Where("`tenant_name` = ? AND `database_name` = ? AND `table_name` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `sql_port` = ?", tenantName, databaseName, tableName, partitionID, svrIP, sqlPort).First(&result).Error

	return
}
