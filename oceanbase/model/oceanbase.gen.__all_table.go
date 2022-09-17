package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTableMgr struct {
	*_BaseMgr
}

// AllTableMgr open func
func AllTableMgr(db *gorm.DB) *_AllTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTableMgr) GetTableName() string {
	return "__all_table"
}

// Reset 重置gorm会话
func (obj *_AllTableMgr) Reset() *_AllTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTableMgr) Get() (result AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTableMgr) Gets() (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTableName table_name获取
func (obj *_AllTableMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithDatabaseID database_id获取
func (obj *_AllTableMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithTableType table_type获取
func (obj *_AllTableMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithLoadType load_type获取
func (obj *_AllTableMgr) WithLoadType(loadType int64) Option {
	return optionFunc(func(o *options) { o.query["load_type"] = loadType })
}

// WithDefType def_type获取
func (obj *_AllTableMgr) WithDefType(defType int64) Option {
	return optionFunc(func(o *options) { o.query["def_type"] = defType })
}

// WithRowkeyColumnNum rowkey_column_num获取
func (obj *_AllTableMgr) WithRowkeyColumnNum(rowkeyColumnNum int64) Option {
	return optionFunc(func(o *options) { o.query["rowkey_column_num"] = rowkeyColumnNum })
}

// WithIndexColumnNum index_column_num获取
func (obj *_AllTableMgr) WithIndexColumnNum(indexColumnNum int64) Option {
	return optionFunc(func(o *options) { o.query["index_column_num"] = indexColumnNum })
}

// WithMaxUsedColumnID max_used_column_id获取
func (obj *_AllTableMgr) WithMaxUsedColumnID(maxUsedColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_column_id"] = maxUsedColumnID })
}

// WithReplicaNum replica_num获取
func (obj *_AllTableMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithAutoincColumnID autoinc_column_id获取
func (obj *_AllTableMgr) WithAutoincColumnID(autoincColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["autoinc_column_id"] = autoincColumnID })
}

// WithAutoIncrement auto_increment获取
func (obj *_AllTableMgr) WithAutoIncrement(autoIncrement uint64) Option {
	return optionFunc(func(o *options) { o.query["auto_increment"] = autoIncrement })
}

// WithReadOnly read_only获取
func (obj *_AllTableMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithRowkeySplitPos rowkey_split_pos获取
func (obj *_AllTableMgr) WithRowkeySplitPos(rowkeySplitPos int64) Option {
	return optionFunc(func(o *options) { o.query["rowkey_split_pos"] = rowkeySplitPos })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllTableMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithExpireCondition expire_condition获取
func (obj *_AllTableMgr) WithExpireCondition(expireCondition string) Option {
	return optionFunc(func(o *options) { o.query["expire_condition"] = expireCondition })
}

// WithIsUseBloomfilter is_use_bloomfilter获取
func (obj *_AllTableMgr) WithIsUseBloomfilter(isUseBloomfilter int64) Option {
	return optionFunc(func(o *options) { o.query["is_use_bloomfilter"] = isUseBloomfilter })
}

// WithComment comment获取
func (obj *_AllTableMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithBlockSize block_size获取
func (obj *_AllTableMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithCollationType collation_type获取
func (obj *_AllTableMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithDataTableID data_table_id获取
func (obj *_AllTableMgr) WithDataTableID(dataTableID int64) Option {
	return optionFunc(func(o *options) { o.query["data_table_id"] = dataTableID })
}

// WithIndexStatus index_status获取
func (obj *_AllTableMgr) WithIndexStatus(indexStatus int64) Option {
	return optionFunc(func(o *options) { o.query["index_status"] = indexStatus })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllTableMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithProgressiveMergeNum progressive_merge_num获取
func (obj *_AllTableMgr) WithProgressiveMergeNum(progressiveMergeNum int64) Option {
	return optionFunc(func(o *options) { o.query["progressive_merge_num"] = progressiveMergeNum })
}

// WithIndexType index_type获取
func (obj *_AllTableMgr) WithIndexType(indexType int64) Option {
	return optionFunc(func(o *options) { o.query["index_type"] = indexType })
}

// WithPartLevel part_level获取
func (obj *_AllTableMgr) WithPartLevel(partLevel int64) Option {
	return optionFunc(func(o *options) { o.query["part_level"] = partLevel })
}

// WithPartFuncType part_func_type获取
func (obj *_AllTableMgr) WithPartFuncType(partFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["part_func_type"] = partFuncType })
}

// WithPartFuncExpr part_func_expr获取
func (obj *_AllTableMgr) WithPartFuncExpr(partFuncExpr string) Option {
	return optionFunc(func(o *options) { o.query["part_func_expr"] = partFuncExpr })
}

// WithPartNum part_num获取
func (obj *_AllTableMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithSubPartFuncType sub_part_func_type获取
func (obj *_AllTableMgr) WithSubPartFuncType(subPartFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_type"] = subPartFuncType })
}

// WithSubPartFuncExpr sub_part_func_expr获取
func (obj *_AllTableMgr) WithSubPartFuncExpr(subPartFuncExpr string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_expr"] = subPartFuncExpr })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllTableMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithCreateMemVersion create_mem_version获取
func (obj *_AllTableMgr) WithCreateMemVersion(createMemVersion int64) Option {
	return optionFunc(func(o *options) { o.query["create_mem_version"] = createMemVersion })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTableMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithViewDefinition view_definition获取
func (obj *_AllTableMgr) WithViewDefinition(viewDefinition string) Option {
	return optionFunc(func(o *options) { o.query["view_definition"] = viewDefinition })
}

// WithViewCheckOption view_check_option获取
func (obj *_AllTableMgr) WithViewCheckOption(viewCheckOption int64) Option {
	return optionFunc(func(o *options) { o.query["view_check_option"] = viewCheckOption })
}

// WithViewIsUpdatable view_is_updatable获取
func (obj *_AllTableMgr) WithViewIsUpdatable(viewIsUpdatable int64) Option {
	return optionFunc(func(o *options) { o.query["view_is_updatable"] = viewIsUpdatable })
}

// WithZoneList zone_list获取
func (obj *_AllTableMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllTableMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithIndexUsingType index_using_type获取
func (obj *_AllTableMgr) WithIndexUsingType(indexUsingType int64) Option {
	return optionFunc(func(o *options) { o.query["index_using_type"] = indexUsingType })
}

// WithParserName parser_name获取
func (obj *_AllTableMgr) WithParserName(parserName string) Option {
	return optionFunc(func(o *options) { o.query["parser_name"] = parserName })
}

// WithIndexAttributesSet index_attributes_set获取
func (obj *_AllTableMgr) WithIndexAttributesSet(indexAttributesSet int64) Option {
	return optionFunc(func(o *options) { o.query["index_attributes_set"] = indexAttributesSet })
}

// WithLocality locality获取
func (obj *_AllTableMgr) WithLocality(locality string) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}

// WithTabletSize tablet_size获取
func (obj *_AllTableMgr) WithTabletSize(tabletSize int64) Option {
	return optionFunc(func(o *options) { o.query["tablet_size"] = tabletSize })
}

// WithPctfree pctfree获取
func (obj *_AllTableMgr) WithPctfree(pctfree int64) Option {
	return optionFunc(func(o *options) { o.query["pctfree"] = pctfree })
}

// WithPreviousLocality previous_locality获取
func (obj *_AllTableMgr) WithPreviousLocality(previousLocality string) Option {
	return optionFunc(func(o *options) { o.query["previous_locality"] = previousLocality })
}

// WithMaxUsedPartID max_used_part_id获取
func (obj *_AllTableMgr) WithMaxUsedPartID(maxUsedPartID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_part_id"] = maxUsedPartID })
}

// WithPartitionCntWithinPartitionTable partition_cnt_within_partition_table获取
func (obj *_AllTableMgr) WithPartitionCntWithinPartitionTable(partitionCntWithinPartitionTable int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt_within_partition_table"] = partitionCntWithinPartitionTable })
}

// WithPartitionStatus partition_status获取
func (obj *_AllTableMgr) WithPartitionStatus(partitionStatus int64) Option {
	return optionFunc(func(o *options) { o.query["partition_status"] = partitionStatus })
}

// WithPartitionSchemaVersion partition_schema_version获取
func (obj *_AllTableMgr) WithPartitionSchemaVersion(partitionSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["partition_schema_version"] = partitionSchemaVersion })
}

// WithMaxUsedConstraintID max_used_constraint_id获取
func (obj *_AllTableMgr) WithMaxUsedConstraintID(maxUsedConstraintID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_constraint_id"] = maxUsedConstraintID })
}

// WithSessionID session_id获取
func (obj *_AllTableMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithPkComment pk_comment获取
func (obj *_AllTableMgr) WithPkComment(pkComment string) Option {
	return optionFunc(func(o *options) { o.query["pk_comment"] = pkComment })
}

// WithSessActiveTime sess_active_time获取
func (obj *_AllTableMgr) WithSessActiveTime(sessActiveTime int64) Option {
	return optionFunc(func(o *options) { o.query["sess_active_time"] = sessActiveTime })
}

// WithRowStoreType row_store_type获取
func (obj *_AllTableMgr) WithRowStoreType(rowStoreType string) Option {
	return optionFunc(func(o *options) { o.query["row_store_type"] = rowStoreType })
}

// WithStoreFormat store_format获取
func (obj *_AllTableMgr) WithStoreFormat(storeFormat string) Option {
	return optionFunc(func(o *options) { o.query["store_format"] = storeFormat })
}

// WithDuplicateScope duplicate_scope获取
func (obj *_AllTableMgr) WithDuplicateScope(duplicateScope int64) Option {
	return optionFunc(func(o *options) { o.query["duplicate_scope"] = duplicateScope })
}

// WithBinding binding获取
func (obj *_AllTableMgr) WithBinding(binding int8) Option {
	return optionFunc(func(o *options) { o.query["binding"] = binding })
}

// WithProgressiveMergeRound progressive_merge_round获取
func (obj *_AllTableMgr) WithProgressiveMergeRound(progressiveMergeRound int64) Option {
	return optionFunc(func(o *options) { o.query["progressive_merge_round"] = progressiveMergeRound })
}

// WithStorageFormatVersion storage_format_version获取
func (obj *_AllTableMgr) WithStorageFormatVersion(storageFormatVersion int64) Option {
	return optionFunc(func(o *options) { o.query["storage_format_version"] = storageFormatVersion })
}

// WithTableMode table_mode获取
func (obj *_AllTableMgr) WithTableMode(tableMode int64) Option {
	return optionFunc(func(o *options) { o.query["table_mode"] = tableMode })
}

// WithEncryption encryption获取
func (obj *_AllTableMgr) WithEncryption(encryption string) Option {
	return optionFunc(func(o *options) { o.query["encryption"] = encryption })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllTableMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllTableMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// WithIsSubPartTemplate is_sub_part_template获取
func (obj *_AllTableMgr) WithIsSubPartTemplate(isSubPartTemplate int8) Option {
	return optionFunc(func(o *options) { o.query["is_sub_part_template"] = isSubPartTemplate })
}

// WithDop dop获取
func (obj *_AllTableMgr) WithDop(dop int64) Option {
	return optionFunc(func(o *options) { o.query["dop"] = dop })
}

// WithCharacterSetClient character_set_client获取
func (obj *_AllTableMgr) WithCharacterSetClient(characterSetClient int64) Option {
	return optionFunc(func(o *options) { o.query["character_set_client"] = characterSetClient })
}

// WithCollationConnection collation_connection获取
func (obj *_AllTableMgr) WithCollationConnection(collationConnection int64) Option {
	return optionFunc(func(o *options) { o.query["collation_connection"] = collationConnection })
}

// WithAutoPartSize auto_part_size获取
func (obj *_AllTableMgr) WithAutoPartSize(autoPartSize int64) Option {
	return optionFunc(func(o *options) { o.query["auto_part_size"] = autoPartSize })
}

// WithAutoPart auto_part获取
func (obj *_AllTableMgr) WithAutoPart(autoPart int8) Option {
	return optionFunc(func(o *options) { o.query["auto_part"] = autoPart })
}

// GetByOption 功能选项模式获取
func (obj *_AllTableMgr) GetByOption(opts ...Option) (result AllTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTableMgr) GetByOptions(opts ...Option) (results []*AllTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllTableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTableMgr) GetFromTenantID(tenantID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTableMgr) GetFromTableID(tableID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllTableMgr) GetFromTableName(tableName string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllTableMgr) GetBatchFromTableName(tableNames []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTableMgr) GetFromDatabaseID(databaseID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTableMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_AllTableMgr) GetFromTableType(tableType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_AllTableMgr) GetBatchFromTableType(tableTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

// GetFromLoadType 通过load_type获取内容
func (obj *_AllTableMgr) GetFromLoadType(loadType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`load_type` = ?", loadType).Find(&results).Error

	return
}

// GetBatchFromLoadType 批量查找
func (obj *_AllTableMgr) GetBatchFromLoadType(loadTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`load_type` IN (?)", loadTypes).Find(&results).Error

	return
}

// GetFromDefType 通过def_type获取内容
func (obj *_AllTableMgr) GetFromDefType(defType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`def_type` = ?", defType).Find(&results).Error

	return
}

// GetBatchFromDefType 批量查找
func (obj *_AllTableMgr) GetBatchFromDefType(defTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`def_type` IN (?)", defTypes).Find(&results).Error

	return
}

// GetFromRowkeyColumnNum 通过rowkey_column_num获取内容
func (obj *_AllTableMgr) GetFromRowkeyColumnNum(rowkeyColumnNum int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`rowkey_column_num` = ?", rowkeyColumnNum).Find(&results).Error

	return
}

// GetBatchFromRowkeyColumnNum 批量查找
func (obj *_AllTableMgr) GetBatchFromRowkeyColumnNum(rowkeyColumnNums []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`rowkey_column_num` IN (?)", rowkeyColumnNums).Find(&results).Error

	return
}

// GetFromIndexColumnNum 通过index_column_num获取内容
func (obj *_AllTableMgr) GetFromIndexColumnNum(indexColumnNum int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_column_num` = ?", indexColumnNum).Find(&results).Error

	return
}

// GetBatchFromIndexColumnNum 批量查找
func (obj *_AllTableMgr) GetBatchFromIndexColumnNum(indexColumnNums []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_column_num` IN (?)", indexColumnNums).Find(&results).Error

	return
}

// GetFromMaxUsedColumnID 通过max_used_column_id获取内容
func (obj *_AllTableMgr) GetFromMaxUsedColumnID(maxUsedColumnID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`max_used_column_id` = ?", maxUsedColumnID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedColumnID 批量查找
func (obj *_AllTableMgr) GetBatchFromMaxUsedColumnID(maxUsedColumnIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`max_used_column_id` IN (?)", maxUsedColumnIDs).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllTableMgr) GetFromReplicaNum(replicaNum int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllTableMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromAutoincColumnID 通过autoinc_column_id获取内容
func (obj *_AllTableMgr) GetFromAutoincColumnID(autoincColumnID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`autoinc_column_id` = ?", autoincColumnID).Find(&results).Error

	return
}

// GetBatchFromAutoincColumnID 批量查找
func (obj *_AllTableMgr) GetBatchFromAutoincColumnID(autoincColumnIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`autoinc_column_id` IN (?)", autoincColumnIDs).Find(&results).Error

	return
}

// GetFromAutoIncrement 通过auto_increment获取内容
func (obj *_AllTableMgr) GetFromAutoIncrement(autoIncrement uint64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`auto_increment` = ?", autoIncrement).Find(&results).Error

	return
}

// GetBatchFromAutoIncrement 批量查找
func (obj *_AllTableMgr) GetBatchFromAutoIncrement(autoIncrements []uint64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`auto_increment` IN (?)", autoIncrements).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllTableMgr) GetFromReadOnly(readOnly int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllTableMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromRowkeySplitPos 通过rowkey_split_pos获取内容
func (obj *_AllTableMgr) GetFromRowkeySplitPos(rowkeySplitPos int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`rowkey_split_pos` = ?", rowkeySplitPos).Find(&results).Error

	return
}

// GetBatchFromRowkeySplitPos 批量查找
func (obj *_AllTableMgr) GetBatchFromRowkeySplitPos(rowkeySplitPoss []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`rowkey_split_pos` IN (?)", rowkeySplitPoss).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllTableMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllTableMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromExpireCondition 通过expire_condition获取内容
func (obj *_AllTableMgr) GetFromExpireCondition(expireCondition string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`expire_condition` = ?", expireCondition).Find(&results).Error

	return
}

// GetBatchFromExpireCondition 批量查找
func (obj *_AllTableMgr) GetBatchFromExpireCondition(expireConditions []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`expire_condition` IN (?)", expireConditions).Find(&results).Error

	return
}

// GetFromIsUseBloomfilter 通过is_use_bloomfilter获取内容
func (obj *_AllTableMgr) GetFromIsUseBloomfilter(isUseBloomfilter int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`is_use_bloomfilter` = ?", isUseBloomfilter).Find(&results).Error

	return
}

// GetBatchFromIsUseBloomfilter 批量查找
func (obj *_AllTableMgr) GetBatchFromIsUseBloomfilter(isUseBloomfilters []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`is_use_bloomfilter` IN (?)", isUseBloomfilters).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllTableMgr) GetFromComment(comment string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllTableMgr) GetBatchFromComment(comments []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllTableMgr) GetFromBlockSize(blockSize int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllTableMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllTableMgr) GetFromCollationType(collationType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllTableMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromDataTableID 通过data_table_id获取内容
func (obj *_AllTableMgr) GetFromDataTableID(dataTableID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}

// GetBatchFromDataTableID 批量查找
func (obj *_AllTableMgr) GetBatchFromDataTableID(dataTableIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`data_table_id` IN (?)", dataTableIDs).Find(&results).Error

	return
}

// GetFromIndexStatus 通过index_status获取内容
func (obj *_AllTableMgr) GetFromIndexStatus(indexStatus int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_status` = ?", indexStatus).Find(&results).Error

	return
}

// GetBatchFromIndexStatus 批量查找
func (obj *_AllTableMgr) GetBatchFromIndexStatus(indexStatuss []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_status` IN (?)", indexStatuss).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllTableMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllTableMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromProgressiveMergeNum 通过progressive_merge_num获取内容
func (obj *_AllTableMgr) GetFromProgressiveMergeNum(progressiveMergeNum int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`progressive_merge_num` = ?", progressiveMergeNum).Find(&results).Error

	return
}

// GetBatchFromProgressiveMergeNum 批量查找
func (obj *_AllTableMgr) GetBatchFromProgressiveMergeNum(progressiveMergeNums []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`progressive_merge_num` IN (?)", progressiveMergeNums).Find(&results).Error

	return
}

// GetFromIndexType 通过index_type获取内容
func (obj *_AllTableMgr) GetFromIndexType(indexType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_type` = ?", indexType).Find(&results).Error

	return
}

// GetBatchFromIndexType 批量查找
func (obj *_AllTableMgr) GetBatchFromIndexType(indexTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_type` IN (?)", indexTypes).Find(&results).Error

	return
}

// GetFromPartLevel 通过part_level获取内容
func (obj *_AllTableMgr) GetFromPartLevel(partLevel int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_level` = ?", partLevel).Find(&results).Error

	return
}

// GetBatchFromPartLevel 批量查找
func (obj *_AllTableMgr) GetBatchFromPartLevel(partLevels []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_level` IN (?)", partLevels).Find(&results).Error

	return
}

// GetFromPartFuncType 通过part_func_type获取内容
func (obj *_AllTableMgr) GetFromPartFuncType(partFuncType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_func_type` = ?", partFuncType).Find(&results).Error

	return
}

// GetBatchFromPartFuncType 批量查找
func (obj *_AllTableMgr) GetBatchFromPartFuncType(partFuncTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_func_type` IN (?)", partFuncTypes).Find(&results).Error

	return
}

// GetFromPartFuncExpr 通过part_func_expr获取内容
func (obj *_AllTableMgr) GetFromPartFuncExpr(partFuncExpr string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_func_expr` = ?", partFuncExpr).Find(&results).Error

	return
}

// GetBatchFromPartFuncExpr 批量查找
func (obj *_AllTableMgr) GetBatchFromPartFuncExpr(partFuncExprs []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_func_expr` IN (?)", partFuncExprs).Find(&results).Error

	return
}

// GetFromPartNum 通过part_num获取内容
func (obj *_AllTableMgr) GetFromPartNum(partNum int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_num` = ?", partNum).Find(&results).Error

	return
}

// GetBatchFromPartNum 批量查找
func (obj *_AllTableMgr) GetBatchFromPartNum(partNums []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`part_num` IN (?)", partNums).Find(&results).Error

	return
}

// GetFromSubPartFuncType 通过sub_part_func_type获取内容
func (obj *_AllTableMgr) GetFromSubPartFuncType(subPartFuncType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sub_part_func_type` = ?", subPartFuncType).Find(&results).Error

	return
}

// GetBatchFromSubPartFuncType 批量查找
func (obj *_AllTableMgr) GetBatchFromSubPartFuncType(subPartFuncTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sub_part_func_type` IN (?)", subPartFuncTypes).Find(&results).Error

	return
}

// GetFromSubPartFuncExpr 通过sub_part_func_expr获取内容
func (obj *_AllTableMgr) GetFromSubPartFuncExpr(subPartFuncExpr string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sub_part_func_expr` = ?", subPartFuncExpr).Find(&results).Error

	return
}

// GetBatchFromSubPartFuncExpr 批量查找
func (obj *_AllTableMgr) GetBatchFromSubPartFuncExpr(subPartFuncExprs []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sub_part_func_expr` IN (?)", subPartFuncExprs).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllTableMgr) GetFromSubPartNum(subPartNum int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllTableMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromCreateMemVersion 通过create_mem_version获取内容
func (obj *_AllTableMgr) GetFromCreateMemVersion(createMemVersion int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`create_mem_version` = ?", createMemVersion).Find(&results).Error

	return
}

// GetBatchFromCreateMemVersion 批量查找
func (obj *_AllTableMgr) GetBatchFromCreateMemVersion(createMemVersions []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`create_mem_version` IN (?)", createMemVersions).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTableMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTableMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromViewDefinition 通过view_definition获取内容
func (obj *_AllTableMgr) GetFromViewDefinition(viewDefinition string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`view_definition` = ?", viewDefinition).Find(&results).Error

	return
}

// GetBatchFromViewDefinition 批量查找
func (obj *_AllTableMgr) GetBatchFromViewDefinition(viewDefinitions []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`view_definition` IN (?)", viewDefinitions).Find(&results).Error

	return
}

// GetFromViewCheckOption 通过view_check_option获取内容
func (obj *_AllTableMgr) GetFromViewCheckOption(viewCheckOption int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`view_check_option` = ?", viewCheckOption).Find(&results).Error

	return
}

// GetBatchFromViewCheckOption 批量查找
func (obj *_AllTableMgr) GetBatchFromViewCheckOption(viewCheckOptions []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`view_check_option` IN (?)", viewCheckOptions).Find(&results).Error

	return
}

// GetFromViewIsUpdatable 通过view_is_updatable获取内容
func (obj *_AllTableMgr) GetFromViewIsUpdatable(viewIsUpdatable int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`view_is_updatable` = ?", viewIsUpdatable).Find(&results).Error

	return
}

// GetBatchFromViewIsUpdatable 批量查找
func (obj *_AllTableMgr) GetBatchFromViewIsUpdatable(viewIsUpdatables []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`view_is_updatable` IN (?)", viewIsUpdatables).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllTableMgr) GetFromZoneList(zoneList string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllTableMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllTableMgr) GetFromPrimaryZone(primaryZone string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllTableMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromIndexUsingType 通过index_using_type获取内容
func (obj *_AllTableMgr) GetFromIndexUsingType(indexUsingType int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_using_type` = ?", indexUsingType).Find(&results).Error

	return
}

// GetBatchFromIndexUsingType 批量查找
func (obj *_AllTableMgr) GetBatchFromIndexUsingType(indexUsingTypes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_using_type` IN (?)", indexUsingTypes).Find(&results).Error

	return
}

// GetFromParserName 通过parser_name获取内容
func (obj *_AllTableMgr) GetFromParserName(parserName string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`parser_name` = ?", parserName).Find(&results).Error

	return
}

// GetBatchFromParserName 批量查找
func (obj *_AllTableMgr) GetBatchFromParserName(parserNames []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`parser_name` IN (?)", parserNames).Find(&results).Error

	return
}

// GetFromIndexAttributesSet 通过index_attributes_set获取内容
func (obj *_AllTableMgr) GetFromIndexAttributesSet(indexAttributesSet int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_attributes_set` = ?", indexAttributesSet).Find(&results).Error

	return
}

// GetBatchFromIndexAttributesSet 批量查找
func (obj *_AllTableMgr) GetBatchFromIndexAttributesSet(indexAttributesSets []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`index_attributes_set` IN (?)", indexAttributesSets).Find(&results).Error

	return
}

// GetFromLocality 通过locality获取内容
func (obj *_AllTableMgr) GetFromLocality(locality string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`locality` = ?", locality).Find(&results).Error

	return
}

// GetBatchFromLocality 批量查找
func (obj *_AllTableMgr) GetBatchFromLocality(localitys []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`locality` IN (?)", localitys).Find(&results).Error

	return
}

// GetFromTabletSize 通过tablet_size获取内容
func (obj *_AllTableMgr) GetFromTabletSize(tabletSize int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tablet_size` = ?", tabletSize).Find(&results).Error

	return
}

// GetBatchFromTabletSize 批量查找
func (obj *_AllTableMgr) GetBatchFromTabletSize(tabletSizes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tablet_size` IN (?)", tabletSizes).Find(&results).Error

	return
}

// GetFromPctfree 通过pctfree获取内容
func (obj *_AllTableMgr) GetFromPctfree(pctfree int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`pctfree` = ?", pctfree).Find(&results).Error

	return
}

// GetBatchFromPctfree 批量查找
func (obj *_AllTableMgr) GetBatchFromPctfree(pctfrees []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`pctfree` IN (?)", pctfrees).Find(&results).Error

	return
}

// GetFromPreviousLocality 通过previous_locality获取内容
func (obj *_AllTableMgr) GetFromPreviousLocality(previousLocality string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`previous_locality` = ?", previousLocality).Find(&results).Error

	return
}

// GetBatchFromPreviousLocality 批量查找
func (obj *_AllTableMgr) GetBatchFromPreviousLocality(previousLocalitys []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`previous_locality` IN (?)", previousLocalitys).Find(&results).Error

	return
}

// GetFromMaxUsedPartID 通过max_used_part_id获取内容
func (obj *_AllTableMgr) GetFromMaxUsedPartID(maxUsedPartID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`max_used_part_id` = ?", maxUsedPartID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedPartID 批量查找
func (obj *_AllTableMgr) GetBatchFromMaxUsedPartID(maxUsedPartIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`max_used_part_id` IN (?)", maxUsedPartIDs).Find(&results).Error

	return
}

// GetFromPartitionCntWithinPartitionTable 通过partition_cnt_within_partition_table获取内容
func (obj *_AllTableMgr) GetFromPartitionCntWithinPartitionTable(partitionCntWithinPartitionTable int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`partition_cnt_within_partition_table` = ?", partitionCntWithinPartitionTable).Find(&results).Error

	return
}

// GetBatchFromPartitionCntWithinPartitionTable 批量查找
func (obj *_AllTableMgr) GetBatchFromPartitionCntWithinPartitionTable(partitionCntWithinPartitionTables []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`partition_cnt_within_partition_table` IN (?)", partitionCntWithinPartitionTables).Find(&results).Error

	return
}

// GetFromPartitionStatus 通过partition_status获取内容
func (obj *_AllTableMgr) GetFromPartitionStatus(partitionStatus int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`partition_status` = ?", partitionStatus).Find(&results).Error

	return
}

// GetBatchFromPartitionStatus 批量查找
func (obj *_AllTableMgr) GetBatchFromPartitionStatus(partitionStatuss []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`partition_status` IN (?)", partitionStatuss).Find(&results).Error

	return
}

// GetFromPartitionSchemaVersion 通过partition_schema_version获取内容
func (obj *_AllTableMgr) GetFromPartitionSchemaVersion(partitionSchemaVersion int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`partition_schema_version` = ?", partitionSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromPartitionSchemaVersion 批量查找
func (obj *_AllTableMgr) GetBatchFromPartitionSchemaVersion(partitionSchemaVersions []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`partition_schema_version` IN (?)", partitionSchemaVersions).Find(&results).Error

	return
}

// GetFromMaxUsedConstraintID 通过max_used_constraint_id获取内容
func (obj *_AllTableMgr) GetFromMaxUsedConstraintID(maxUsedConstraintID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`max_used_constraint_id` = ?", maxUsedConstraintID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedConstraintID 批量查找
func (obj *_AllTableMgr) GetBatchFromMaxUsedConstraintID(maxUsedConstraintIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`max_used_constraint_id` IN (?)", maxUsedConstraintIDs).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllTableMgr) GetFromSessionID(sessionID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllTableMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromPkComment 通过pk_comment获取内容
func (obj *_AllTableMgr) GetFromPkComment(pkComment string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`pk_comment` = ?", pkComment).Find(&results).Error

	return
}

// GetBatchFromPkComment 批量查找
func (obj *_AllTableMgr) GetBatchFromPkComment(pkComments []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`pk_comment` IN (?)", pkComments).Find(&results).Error

	return
}

// GetFromSessActiveTime 通过sess_active_time获取内容
func (obj *_AllTableMgr) GetFromSessActiveTime(sessActiveTime int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sess_active_time` = ?", sessActiveTime).Find(&results).Error

	return
}

// GetBatchFromSessActiveTime 批量查找
func (obj *_AllTableMgr) GetBatchFromSessActiveTime(sessActiveTimes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`sess_active_time` IN (?)", sessActiveTimes).Find(&results).Error

	return
}

// GetFromRowStoreType 通过row_store_type获取内容
func (obj *_AllTableMgr) GetFromRowStoreType(rowStoreType string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`row_store_type` = ?", rowStoreType).Find(&results).Error

	return
}

// GetBatchFromRowStoreType 批量查找
func (obj *_AllTableMgr) GetBatchFromRowStoreType(rowStoreTypes []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`row_store_type` IN (?)", rowStoreTypes).Find(&results).Error

	return
}

// GetFromStoreFormat 通过store_format获取内容
func (obj *_AllTableMgr) GetFromStoreFormat(storeFormat string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`store_format` = ?", storeFormat).Find(&results).Error

	return
}

// GetBatchFromStoreFormat 批量查找
func (obj *_AllTableMgr) GetBatchFromStoreFormat(storeFormats []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`store_format` IN (?)", storeFormats).Find(&results).Error

	return
}

// GetFromDuplicateScope 通过duplicate_scope获取内容
func (obj *_AllTableMgr) GetFromDuplicateScope(duplicateScope int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`duplicate_scope` = ?", duplicateScope).Find(&results).Error

	return
}

// GetBatchFromDuplicateScope 批量查找
func (obj *_AllTableMgr) GetBatchFromDuplicateScope(duplicateScopes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`duplicate_scope` IN (?)", duplicateScopes).Find(&results).Error

	return
}

// GetFromBinding 通过binding获取内容
func (obj *_AllTableMgr) GetFromBinding(binding int8) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`binding` = ?", binding).Find(&results).Error

	return
}

// GetBatchFromBinding 批量查找
func (obj *_AllTableMgr) GetBatchFromBinding(bindings []int8) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`binding` IN (?)", bindings).Find(&results).Error

	return
}

// GetFromProgressiveMergeRound 通过progressive_merge_round获取内容
func (obj *_AllTableMgr) GetFromProgressiveMergeRound(progressiveMergeRound int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`progressive_merge_round` = ?", progressiveMergeRound).Find(&results).Error

	return
}

// GetBatchFromProgressiveMergeRound 批量查找
func (obj *_AllTableMgr) GetBatchFromProgressiveMergeRound(progressiveMergeRounds []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`progressive_merge_round` IN (?)", progressiveMergeRounds).Find(&results).Error

	return
}

// GetFromStorageFormatVersion 通过storage_format_version获取内容
func (obj *_AllTableMgr) GetFromStorageFormatVersion(storageFormatVersion int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`storage_format_version` = ?", storageFormatVersion).Find(&results).Error

	return
}

// GetBatchFromStorageFormatVersion 批量查找
func (obj *_AllTableMgr) GetBatchFromStorageFormatVersion(storageFormatVersions []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`storage_format_version` IN (?)", storageFormatVersions).Find(&results).Error

	return
}

// GetFromTableMode 通过table_mode获取内容
func (obj *_AllTableMgr) GetFromTableMode(tableMode int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_mode` = ?", tableMode).Find(&results).Error

	return
}

// GetBatchFromTableMode 批量查找
func (obj *_AllTableMgr) GetBatchFromTableMode(tableModes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`table_mode` IN (?)", tableModes).Find(&results).Error

	return
}

// GetFromEncryption 通过encryption获取内容
func (obj *_AllTableMgr) GetFromEncryption(encryption string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`encryption` = ?", encryption).Find(&results).Error

	return
}

// GetBatchFromEncryption 批量查找
func (obj *_AllTableMgr) GetBatchFromEncryption(encryptions []string) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`encryption` IN (?)", encryptions).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllTableMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllTableMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllTableMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllTableMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

// GetFromIsSubPartTemplate 通过is_sub_part_template获取内容
func (obj *_AllTableMgr) GetFromIsSubPartTemplate(isSubPartTemplate int8) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`is_sub_part_template` = ?", isSubPartTemplate).Find(&results).Error

	return
}

// GetBatchFromIsSubPartTemplate 批量查找
func (obj *_AllTableMgr) GetBatchFromIsSubPartTemplate(isSubPartTemplates []int8) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`is_sub_part_template` IN (?)", isSubPartTemplates).Find(&results).Error

	return
}

// GetFromDop 通过dop获取内容
func (obj *_AllTableMgr) GetFromDop(dop int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`dop` = ?", dop).Find(&results).Error

	return
}

// GetBatchFromDop 批量查找
func (obj *_AllTableMgr) GetBatchFromDop(dops []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`dop` IN (?)", dops).Find(&results).Error

	return
}

// GetFromCharacterSetClient 通过character_set_client获取内容
func (obj *_AllTableMgr) GetFromCharacterSetClient(characterSetClient int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`character_set_client` = ?", characterSetClient).Find(&results).Error

	return
}

// GetBatchFromCharacterSetClient 批量查找
func (obj *_AllTableMgr) GetBatchFromCharacterSetClient(characterSetClients []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`character_set_client` IN (?)", characterSetClients).Find(&results).Error

	return
}

// GetFromCollationConnection 通过collation_connection获取内容
func (obj *_AllTableMgr) GetFromCollationConnection(collationConnection int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`collation_connection` = ?", collationConnection).Find(&results).Error

	return
}

// GetBatchFromCollationConnection 批量查找
func (obj *_AllTableMgr) GetBatchFromCollationConnection(collationConnections []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`collation_connection` IN (?)", collationConnections).Find(&results).Error

	return
}

// GetFromAutoPartSize 通过auto_part_size获取内容
func (obj *_AllTableMgr) GetFromAutoPartSize(autoPartSize int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`auto_part_size` = ?", autoPartSize).Find(&results).Error

	return
}

// GetBatchFromAutoPartSize 批量查找
func (obj *_AllTableMgr) GetBatchFromAutoPartSize(autoPartSizes []int64) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`auto_part_size` IN (?)", autoPartSizes).Find(&results).Error

	return
}

// GetFromAutoPart 通过auto_part获取内容
func (obj *_AllTableMgr) GetFromAutoPart(autoPart int8) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`auto_part` = ?", autoPart).Find(&results).Error

	return
}

// GetBatchFromAutoPart 批量查找
func (obj *_AllTableMgr) GetBatchFromAutoPart(autoParts []int8) (results []*AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`auto_part` IN (?)", autoParts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64) (result AllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTable{}).Where("`tenant_id` = ? AND `table_id` = ?", tenantID, tableID).First(&result).Error

	return
}
