package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTableHistoryMgr struct {
	*_BaseMgr
}

// AllTableHistoryMgr open func
func AllTableHistoryMgr(db *gorm.DB) *_AllTableHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTableHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTableHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_table_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTableHistoryMgr) GetTableName() string {
	return "__all_table_history"
}

// Reset 重置gorm会话
func (obj *_AllTableHistoryMgr) Reset() *_AllTableHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTableHistoryMgr) Get() (result AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTableHistoryMgr) Gets() (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTableHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTableHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTableHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTableHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllTableHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTableHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTableHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTableName table_name获取
func (obj *_AllTableHistoryMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithDatabaseID database_id获取
func (obj *_AllTableHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithTableType table_type获取
func (obj *_AllTableHistoryMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithLoadType load_type获取
func (obj *_AllTableHistoryMgr) WithLoadType(loadType int64) Option {
	return optionFunc(func(o *options) { o.query["load_type"] = loadType })
}

// WithDefType def_type获取
func (obj *_AllTableHistoryMgr) WithDefType(defType int64) Option {
	return optionFunc(func(o *options) { o.query["def_type"] = defType })
}

// WithRowkeyColumnNum rowkey_column_num获取
func (obj *_AllTableHistoryMgr) WithRowkeyColumnNum(rowkeyColumnNum int64) Option {
	return optionFunc(func(o *options) { o.query["rowkey_column_num"] = rowkeyColumnNum })
}

// WithIndexColumnNum index_column_num获取
func (obj *_AllTableHistoryMgr) WithIndexColumnNum(indexColumnNum int64) Option {
	return optionFunc(func(o *options) { o.query["index_column_num"] = indexColumnNum })
}

// WithMaxUsedColumnID max_used_column_id获取
func (obj *_AllTableHistoryMgr) WithMaxUsedColumnID(maxUsedColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_column_id"] = maxUsedColumnID })
}

// WithReplicaNum replica_num获取
func (obj *_AllTableHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithAutoincColumnID autoinc_column_id获取
func (obj *_AllTableHistoryMgr) WithAutoincColumnID(autoincColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["autoinc_column_id"] = autoincColumnID })
}

// WithAutoIncrement auto_increment获取
func (obj *_AllTableHistoryMgr) WithAutoIncrement(autoIncrement uint64) Option {
	return optionFunc(func(o *options) { o.query["auto_increment"] = autoIncrement })
}

// WithReadOnly read_only获取
func (obj *_AllTableHistoryMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithRowkeySplitPos rowkey_split_pos获取
func (obj *_AllTableHistoryMgr) WithRowkeySplitPos(rowkeySplitPos int64) Option {
	return optionFunc(func(o *options) { o.query["rowkey_split_pos"] = rowkeySplitPos })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllTableHistoryMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithExpireCondition expire_condition获取
func (obj *_AllTableHistoryMgr) WithExpireCondition(expireCondition string) Option {
	return optionFunc(func(o *options) { o.query["expire_condition"] = expireCondition })
}

// WithIsUseBloomfilter is_use_bloomfilter获取
func (obj *_AllTableHistoryMgr) WithIsUseBloomfilter(isUseBloomfilter int64) Option {
	return optionFunc(func(o *options) { o.query["is_use_bloomfilter"] = isUseBloomfilter })
}

// WithComment comment获取
func (obj *_AllTableHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithBlockSize block_size获取
func (obj *_AllTableHistoryMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithCollationType collation_type获取
func (obj *_AllTableHistoryMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithDataTableID data_table_id获取
func (obj *_AllTableHistoryMgr) WithDataTableID(dataTableID int64) Option {
	return optionFunc(func(o *options) { o.query["data_table_id"] = dataTableID })
}

// WithIndexStatus index_status获取
func (obj *_AllTableHistoryMgr) WithIndexStatus(indexStatus int64) Option {
	return optionFunc(func(o *options) { o.query["index_status"] = indexStatus })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllTableHistoryMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithProgressiveMergeNum progressive_merge_num获取
func (obj *_AllTableHistoryMgr) WithProgressiveMergeNum(progressiveMergeNum int64) Option {
	return optionFunc(func(o *options) { o.query["progressive_merge_num"] = progressiveMergeNum })
}

// WithIndexType index_type获取
func (obj *_AllTableHistoryMgr) WithIndexType(indexType int64) Option {
	return optionFunc(func(o *options) { o.query["index_type"] = indexType })
}

// WithPartLevel part_level获取
func (obj *_AllTableHistoryMgr) WithPartLevel(partLevel int64) Option {
	return optionFunc(func(o *options) { o.query["part_level"] = partLevel })
}

// WithPartFuncType part_func_type获取
func (obj *_AllTableHistoryMgr) WithPartFuncType(partFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["part_func_type"] = partFuncType })
}

// WithPartFuncExpr part_func_expr获取
func (obj *_AllTableHistoryMgr) WithPartFuncExpr(partFuncExpr string) Option {
	return optionFunc(func(o *options) { o.query["part_func_expr"] = partFuncExpr })
}

// WithPartNum part_num获取
func (obj *_AllTableHistoryMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithSubPartFuncType sub_part_func_type获取
func (obj *_AllTableHistoryMgr) WithSubPartFuncType(subPartFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_type"] = subPartFuncType })
}

// WithSubPartFuncExpr sub_part_func_expr获取
func (obj *_AllTableHistoryMgr) WithSubPartFuncExpr(subPartFuncExpr string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_expr"] = subPartFuncExpr })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllTableHistoryMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithCreateMemVersion create_mem_version获取
func (obj *_AllTableHistoryMgr) WithCreateMemVersion(createMemVersion int64) Option {
	return optionFunc(func(o *options) { o.query["create_mem_version"] = createMemVersion })
}

// WithViewDefinition view_definition获取
func (obj *_AllTableHistoryMgr) WithViewDefinition(viewDefinition string) Option {
	return optionFunc(func(o *options) { o.query["view_definition"] = viewDefinition })
}

// WithViewCheckOption view_check_option获取
func (obj *_AllTableHistoryMgr) WithViewCheckOption(viewCheckOption int64) Option {
	return optionFunc(func(o *options) { o.query["view_check_option"] = viewCheckOption })
}

// WithViewIsUpdatable view_is_updatable获取
func (obj *_AllTableHistoryMgr) WithViewIsUpdatable(viewIsUpdatable int64) Option {
	return optionFunc(func(o *options) { o.query["view_is_updatable"] = viewIsUpdatable })
}

// WithZoneList zone_list获取
func (obj *_AllTableHistoryMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllTableHistoryMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithIndexUsingType index_using_type获取
func (obj *_AllTableHistoryMgr) WithIndexUsingType(indexUsingType int64) Option {
	return optionFunc(func(o *options) { o.query["index_using_type"] = indexUsingType })
}

// WithParserName parser_name获取
func (obj *_AllTableHistoryMgr) WithParserName(parserName string) Option {
	return optionFunc(func(o *options) { o.query["parser_name"] = parserName })
}

// WithIndexAttributesSet index_attributes_set获取
func (obj *_AllTableHistoryMgr) WithIndexAttributesSet(indexAttributesSet int64) Option {
	return optionFunc(func(o *options) { o.query["index_attributes_set"] = indexAttributesSet })
}

// WithLocality locality获取
func (obj *_AllTableHistoryMgr) WithLocality(locality string) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}

// WithTabletSize tablet_size获取
func (obj *_AllTableHistoryMgr) WithTabletSize(tabletSize int64) Option {
	return optionFunc(func(o *options) { o.query["tablet_size"] = tabletSize })
}

// WithPctfree pctfree获取
func (obj *_AllTableHistoryMgr) WithPctfree(pctfree int64) Option {
	return optionFunc(func(o *options) { o.query["pctfree"] = pctfree })
}

// WithPreviousLocality previous_locality获取
func (obj *_AllTableHistoryMgr) WithPreviousLocality(previousLocality string) Option {
	return optionFunc(func(o *options) { o.query["previous_locality"] = previousLocality })
}

// WithMaxUsedPartID max_used_part_id获取
func (obj *_AllTableHistoryMgr) WithMaxUsedPartID(maxUsedPartID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_part_id"] = maxUsedPartID })
}

// WithPartitionCntWithinPartitionTable partition_cnt_within_partition_table获取
func (obj *_AllTableHistoryMgr) WithPartitionCntWithinPartitionTable(partitionCntWithinPartitionTable int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt_within_partition_table"] = partitionCntWithinPartitionTable })
}

// WithPartitionStatus partition_status获取
func (obj *_AllTableHistoryMgr) WithPartitionStatus(partitionStatus int64) Option {
	return optionFunc(func(o *options) { o.query["partition_status"] = partitionStatus })
}

// WithPartitionSchemaVersion partition_schema_version获取
func (obj *_AllTableHistoryMgr) WithPartitionSchemaVersion(partitionSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["partition_schema_version"] = partitionSchemaVersion })
}

// WithMaxUsedConstraintID max_used_constraint_id获取
func (obj *_AllTableHistoryMgr) WithMaxUsedConstraintID(maxUsedConstraintID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_constraint_id"] = maxUsedConstraintID })
}

// WithSessionID session_id获取
func (obj *_AllTableHistoryMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithPkComment pk_comment获取
func (obj *_AllTableHistoryMgr) WithPkComment(pkComment string) Option {
	return optionFunc(func(o *options) { o.query["pk_comment"] = pkComment })
}

// WithSessActiveTime sess_active_time获取
func (obj *_AllTableHistoryMgr) WithSessActiveTime(sessActiveTime int64) Option {
	return optionFunc(func(o *options) { o.query["sess_active_time"] = sessActiveTime })
}

// WithRowStoreType row_store_type获取
func (obj *_AllTableHistoryMgr) WithRowStoreType(rowStoreType string) Option {
	return optionFunc(func(o *options) { o.query["row_store_type"] = rowStoreType })
}

// WithStoreFormat store_format获取
func (obj *_AllTableHistoryMgr) WithStoreFormat(storeFormat string) Option {
	return optionFunc(func(o *options) { o.query["store_format"] = storeFormat })
}

// WithDuplicateScope duplicate_scope获取
func (obj *_AllTableHistoryMgr) WithDuplicateScope(duplicateScope int64) Option {
	return optionFunc(func(o *options) { o.query["duplicate_scope"] = duplicateScope })
}

// WithBinding binding获取
func (obj *_AllTableHistoryMgr) WithBinding(binding int8) Option {
	return optionFunc(func(o *options) { o.query["binding"] = binding })
}

// WithProgressiveMergeRound progressive_merge_round获取
func (obj *_AllTableHistoryMgr) WithProgressiveMergeRound(progressiveMergeRound int64) Option {
	return optionFunc(func(o *options) { o.query["progressive_merge_round"] = progressiveMergeRound })
}

// WithStorageFormatVersion storage_format_version获取
func (obj *_AllTableHistoryMgr) WithStorageFormatVersion(storageFormatVersion int64) Option {
	return optionFunc(func(o *options) { o.query["storage_format_version"] = storageFormatVersion })
}

// WithTableMode table_mode获取
func (obj *_AllTableHistoryMgr) WithTableMode(tableMode int64) Option {
	return optionFunc(func(o *options) { o.query["table_mode"] = tableMode })
}

// WithEncryption encryption获取
func (obj *_AllTableHistoryMgr) WithEncryption(encryption string) Option {
	return optionFunc(func(o *options) { o.query["encryption"] = encryption })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllTableHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllTableHistoryMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// WithIsSubPartTemplate is_sub_part_template获取
func (obj *_AllTableHistoryMgr) WithIsSubPartTemplate(isSubPartTemplate int8) Option {
	return optionFunc(func(o *options) { o.query["is_sub_part_template"] = isSubPartTemplate })
}

// WithDop dop获取
func (obj *_AllTableHistoryMgr) WithDop(dop int64) Option {
	return optionFunc(func(o *options) { o.query["dop"] = dop })
}

// WithCharacterSetClient character_set_client获取
func (obj *_AllTableHistoryMgr) WithCharacterSetClient(characterSetClient int64) Option {
	return optionFunc(func(o *options) { o.query["character_set_client"] = characterSetClient })
}

// WithCollationConnection collation_connection获取
func (obj *_AllTableHistoryMgr) WithCollationConnection(collationConnection int64) Option {
	return optionFunc(func(o *options) { o.query["collation_connection"] = collationConnection })
}

// WithAutoPartSize auto_part_size获取
func (obj *_AllTableHistoryMgr) WithAutoPartSize(autoPartSize int64) Option {
	return optionFunc(func(o *options) { o.query["auto_part_size"] = autoPartSize })
}

// WithAutoPart auto_part获取
func (obj *_AllTableHistoryMgr) WithAutoPart(autoPart int8) Option {
	return optionFunc(func(o *options) { o.query["auto_part"] = autoPart })
}

// GetByOption 功能选项模式获取
func (obj *_AllTableHistoryMgr) GetByOption(opts ...Option) (result AllTableHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTableHistoryMgr) GetByOptions(opts ...Option) (results []*AllTableHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTableHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTableHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where(options.query)
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
func (obj *_AllTableHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTableHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTableHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTableHistoryMgr) GetFromTableID(tableID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTableHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTableHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllTableHistoryMgr) GetFromTableName(tableName string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTableName(tableNames []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTableHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_AllTableHistoryMgr) GetFromTableType(tableType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTableType(tableTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

// GetFromLoadType 通过load_type获取内容
func (obj *_AllTableHistoryMgr) GetFromLoadType(loadType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`load_type` = ?", loadType).Find(&results).Error

	return
}

// GetBatchFromLoadType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromLoadType(loadTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`load_type` IN (?)", loadTypes).Find(&results).Error

	return
}

// GetFromDefType 通过def_type获取内容
func (obj *_AllTableHistoryMgr) GetFromDefType(defType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`def_type` = ?", defType).Find(&results).Error

	return
}

// GetBatchFromDefType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromDefType(defTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`def_type` IN (?)", defTypes).Find(&results).Error

	return
}

// GetFromRowkeyColumnNum 通过rowkey_column_num获取内容
func (obj *_AllTableHistoryMgr) GetFromRowkeyColumnNum(rowkeyColumnNum int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`rowkey_column_num` = ?", rowkeyColumnNum).Find(&results).Error

	return
}

// GetBatchFromRowkeyColumnNum 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromRowkeyColumnNum(rowkeyColumnNums []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`rowkey_column_num` IN (?)", rowkeyColumnNums).Find(&results).Error

	return
}

// GetFromIndexColumnNum 通过index_column_num获取内容
func (obj *_AllTableHistoryMgr) GetFromIndexColumnNum(indexColumnNum int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_column_num` = ?", indexColumnNum).Find(&results).Error

	return
}

// GetBatchFromIndexColumnNum 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIndexColumnNum(indexColumnNums []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_column_num` IN (?)", indexColumnNums).Find(&results).Error

	return
}

// GetFromMaxUsedColumnID 通过max_used_column_id获取内容
func (obj *_AllTableHistoryMgr) GetFromMaxUsedColumnID(maxUsedColumnID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`max_used_column_id` = ?", maxUsedColumnID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedColumnID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromMaxUsedColumnID(maxUsedColumnIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`max_used_column_id` IN (?)", maxUsedColumnIDs).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllTableHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromAutoincColumnID 通过autoinc_column_id获取内容
func (obj *_AllTableHistoryMgr) GetFromAutoincColumnID(autoincColumnID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`autoinc_column_id` = ?", autoincColumnID).Find(&results).Error

	return
}

// GetBatchFromAutoincColumnID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromAutoincColumnID(autoincColumnIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`autoinc_column_id` IN (?)", autoincColumnIDs).Find(&results).Error

	return
}

// GetFromAutoIncrement 通过auto_increment获取内容
func (obj *_AllTableHistoryMgr) GetFromAutoIncrement(autoIncrement uint64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`auto_increment` = ?", autoIncrement).Find(&results).Error

	return
}

// GetBatchFromAutoIncrement 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromAutoIncrement(autoIncrements []uint64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`auto_increment` IN (?)", autoIncrements).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllTableHistoryMgr) GetFromReadOnly(readOnly int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromRowkeySplitPos 通过rowkey_split_pos获取内容
func (obj *_AllTableHistoryMgr) GetFromRowkeySplitPos(rowkeySplitPos int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`rowkey_split_pos` = ?", rowkeySplitPos).Find(&results).Error

	return
}

// GetBatchFromRowkeySplitPos 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromRowkeySplitPos(rowkeySplitPoss []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`rowkey_split_pos` IN (?)", rowkeySplitPoss).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllTableHistoryMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromExpireCondition 通过expire_condition获取内容
func (obj *_AllTableHistoryMgr) GetFromExpireCondition(expireCondition string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`expire_condition` = ?", expireCondition).Find(&results).Error

	return
}

// GetBatchFromExpireCondition 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromExpireCondition(expireConditions []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`expire_condition` IN (?)", expireConditions).Find(&results).Error

	return
}

// GetFromIsUseBloomfilter 通过is_use_bloomfilter获取内容
func (obj *_AllTableHistoryMgr) GetFromIsUseBloomfilter(isUseBloomfilter int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`is_use_bloomfilter` = ?", isUseBloomfilter).Find(&results).Error

	return
}

// GetBatchFromIsUseBloomfilter 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIsUseBloomfilter(isUseBloomfilters []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`is_use_bloomfilter` IN (?)", isUseBloomfilters).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllTableHistoryMgr) GetFromComment(comment string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromComment(comments []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllTableHistoryMgr) GetFromBlockSize(blockSize int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllTableHistoryMgr) GetFromCollationType(collationType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromDataTableID 通过data_table_id获取内容
func (obj *_AllTableHistoryMgr) GetFromDataTableID(dataTableID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}

// GetBatchFromDataTableID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromDataTableID(dataTableIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`data_table_id` IN (?)", dataTableIDs).Find(&results).Error

	return
}

// GetFromIndexStatus 通过index_status获取内容
func (obj *_AllTableHistoryMgr) GetFromIndexStatus(indexStatus int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_status` = ?", indexStatus).Find(&results).Error

	return
}

// GetBatchFromIndexStatus 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIndexStatus(indexStatuss []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_status` IN (?)", indexStatuss).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllTableHistoryMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromProgressiveMergeNum 通过progressive_merge_num获取内容
func (obj *_AllTableHistoryMgr) GetFromProgressiveMergeNum(progressiveMergeNum int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`progressive_merge_num` = ?", progressiveMergeNum).Find(&results).Error

	return
}

// GetBatchFromProgressiveMergeNum 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromProgressiveMergeNum(progressiveMergeNums []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`progressive_merge_num` IN (?)", progressiveMergeNums).Find(&results).Error

	return
}

// GetFromIndexType 通过index_type获取内容
func (obj *_AllTableHistoryMgr) GetFromIndexType(indexType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_type` = ?", indexType).Find(&results).Error

	return
}

// GetBatchFromIndexType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIndexType(indexTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_type` IN (?)", indexTypes).Find(&results).Error

	return
}

// GetFromPartLevel 通过part_level获取内容
func (obj *_AllTableHistoryMgr) GetFromPartLevel(partLevel int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_level` = ?", partLevel).Find(&results).Error

	return
}

// GetBatchFromPartLevel 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPartLevel(partLevels []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_level` IN (?)", partLevels).Find(&results).Error

	return
}

// GetFromPartFuncType 通过part_func_type获取内容
func (obj *_AllTableHistoryMgr) GetFromPartFuncType(partFuncType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_func_type` = ?", partFuncType).Find(&results).Error

	return
}

// GetBatchFromPartFuncType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPartFuncType(partFuncTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_func_type` IN (?)", partFuncTypes).Find(&results).Error

	return
}

// GetFromPartFuncExpr 通过part_func_expr获取内容
func (obj *_AllTableHistoryMgr) GetFromPartFuncExpr(partFuncExpr string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_func_expr` = ?", partFuncExpr).Find(&results).Error

	return
}

// GetBatchFromPartFuncExpr 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPartFuncExpr(partFuncExprs []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_func_expr` IN (?)", partFuncExprs).Find(&results).Error

	return
}

// GetFromPartNum 通过part_num获取内容
func (obj *_AllTableHistoryMgr) GetFromPartNum(partNum int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_num` = ?", partNum).Find(&results).Error

	return
}

// GetBatchFromPartNum 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPartNum(partNums []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`part_num` IN (?)", partNums).Find(&results).Error

	return
}

// GetFromSubPartFuncType 通过sub_part_func_type获取内容
func (obj *_AllTableHistoryMgr) GetFromSubPartFuncType(subPartFuncType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sub_part_func_type` = ?", subPartFuncType).Find(&results).Error

	return
}

// GetBatchFromSubPartFuncType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromSubPartFuncType(subPartFuncTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sub_part_func_type` IN (?)", subPartFuncTypes).Find(&results).Error

	return
}

// GetFromSubPartFuncExpr 通过sub_part_func_expr获取内容
func (obj *_AllTableHistoryMgr) GetFromSubPartFuncExpr(subPartFuncExpr string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sub_part_func_expr` = ?", subPartFuncExpr).Find(&results).Error

	return
}

// GetBatchFromSubPartFuncExpr 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromSubPartFuncExpr(subPartFuncExprs []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sub_part_func_expr` IN (?)", subPartFuncExprs).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllTableHistoryMgr) GetFromSubPartNum(subPartNum int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromCreateMemVersion 通过create_mem_version获取内容
func (obj *_AllTableHistoryMgr) GetFromCreateMemVersion(createMemVersion int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`create_mem_version` = ?", createMemVersion).Find(&results).Error

	return
}

// GetBatchFromCreateMemVersion 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromCreateMemVersion(createMemVersions []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`create_mem_version` IN (?)", createMemVersions).Find(&results).Error

	return
}

// GetFromViewDefinition 通过view_definition获取内容
func (obj *_AllTableHistoryMgr) GetFromViewDefinition(viewDefinition string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`view_definition` = ?", viewDefinition).Find(&results).Error

	return
}

// GetBatchFromViewDefinition 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromViewDefinition(viewDefinitions []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`view_definition` IN (?)", viewDefinitions).Find(&results).Error

	return
}

// GetFromViewCheckOption 通过view_check_option获取内容
func (obj *_AllTableHistoryMgr) GetFromViewCheckOption(viewCheckOption int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`view_check_option` = ?", viewCheckOption).Find(&results).Error

	return
}

// GetBatchFromViewCheckOption 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromViewCheckOption(viewCheckOptions []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`view_check_option` IN (?)", viewCheckOptions).Find(&results).Error

	return
}

// GetFromViewIsUpdatable 通过view_is_updatable获取内容
func (obj *_AllTableHistoryMgr) GetFromViewIsUpdatable(viewIsUpdatable int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`view_is_updatable` = ?", viewIsUpdatable).Find(&results).Error

	return
}

// GetBatchFromViewIsUpdatable 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromViewIsUpdatable(viewIsUpdatables []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`view_is_updatable` IN (?)", viewIsUpdatables).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllTableHistoryMgr) GetFromZoneList(zoneList string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllTableHistoryMgr) GetFromPrimaryZone(primaryZone string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromIndexUsingType 通过index_using_type获取内容
func (obj *_AllTableHistoryMgr) GetFromIndexUsingType(indexUsingType int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_using_type` = ?", indexUsingType).Find(&results).Error

	return
}

// GetBatchFromIndexUsingType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIndexUsingType(indexUsingTypes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_using_type` IN (?)", indexUsingTypes).Find(&results).Error

	return
}

// GetFromParserName 通过parser_name获取内容
func (obj *_AllTableHistoryMgr) GetFromParserName(parserName string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`parser_name` = ?", parserName).Find(&results).Error

	return
}

// GetBatchFromParserName 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromParserName(parserNames []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`parser_name` IN (?)", parserNames).Find(&results).Error

	return
}

// GetFromIndexAttributesSet 通过index_attributes_set获取内容
func (obj *_AllTableHistoryMgr) GetFromIndexAttributesSet(indexAttributesSet int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_attributes_set` = ?", indexAttributesSet).Find(&results).Error

	return
}

// GetBatchFromIndexAttributesSet 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIndexAttributesSet(indexAttributesSets []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`index_attributes_set` IN (?)", indexAttributesSets).Find(&results).Error

	return
}

// GetFromLocality 通过locality获取内容
func (obj *_AllTableHistoryMgr) GetFromLocality(locality string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`locality` = ?", locality).Find(&results).Error

	return
}

// GetBatchFromLocality 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromLocality(localitys []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`locality` IN (?)", localitys).Find(&results).Error

	return
}

// GetFromTabletSize 通过tablet_size获取内容
func (obj *_AllTableHistoryMgr) GetFromTabletSize(tabletSize int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tablet_size` = ?", tabletSize).Find(&results).Error

	return
}

// GetBatchFromTabletSize 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTabletSize(tabletSizes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tablet_size` IN (?)", tabletSizes).Find(&results).Error

	return
}

// GetFromPctfree 通过pctfree获取内容
func (obj *_AllTableHistoryMgr) GetFromPctfree(pctfree int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`pctfree` = ?", pctfree).Find(&results).Error

	return
}

// GetBatchFromPctfree 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPctfree(pctfrees []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`pctfree` IN (?)", pctfrees).Find(&results).Error

	return
}

// GetFromPreviousLocality 通过previous_locality获取内容
func (obj *_AllTableHistoryMgr) GetFromPreviousLocality(previousLocality string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`previous_locality` = ?", previousLocality).Find(&results).Error

	return
}

// GetBatchFromPreviousLocality 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPreviousLocality(previousLocalitys []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`previous_locality` IN (?)", previousLocalitys).Find(&results).Error

	return
}

// GetFromMaxUsedPartID 通过max_used_part_id获取内容
func (obj *_AllTableHistoryMgr) GetFromMaxUsedPartID(maxUsedPartID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`max_used_part_id` = ?", maxUsedPartID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedPartID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromMaxUsedPartID(maxUsedPartIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`max_used_part_id` IN (?)", maxUsedPartIDs).Find(&results).Error

	return
}

// GetFromPartitionCntWithinPartitionTable 通过partition_cnt_within_partition_table获取内容
func (obj *_AllTableHistoryMgr) GetFromPartitionCntWithinPartitionTable(partitionCntWithinPartitionTable int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`partition_cnt_within_partition_table` = ?", partitionCntWithinPartitionTable).Find(&results).Error

	return
}

// GetBatchFromPartitionCntWithinPartitionTable 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPartitionCntWithinPartitionTable(partitionCntWithinPartitionTables []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`partition_cnt_within_partition_table` IN (?)", partitionCntWithinPartitionTables).Find(&results).Error

	return
}

// GetFromPartitionStatus 通过partition_status获取内容
func (obj *_AllTableHistoryMgr) GetFromPartitionStatus(partitionStatus int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`partition_status` = ?", partitionStatus).Find(&results).Error

	return
}

// GetBatchFromPartitionStatus 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPartitionStatus(partitionStatuss []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`partition_status` IN (?)", partitionStatuss).Find(&results).Error

	return
}

// GetFromPartitionSchemaVersion 通过partition_schema_version获取内容
func (obj *_AllTableHistoryMgr) GetFromPartitionSchemaVersion(partitionSchemaVersion int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`partition_schema_version` = ?", partitionSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromPartitionSchemaVersion 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPartitionSchemaVersion(partitionSchemaVersions []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`partition_schema_version` IN (?)", partitionSchemaVersions).Find(&results).Error

	return
}

// GetFromMaxUsedConstraintID 通过max_used_constraint_id获取内容
func (obj *_AllTableHistoryMgr) GetFromMaxUsedConstraintID(maxUsedConstraintID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`max_used_constraint_id` = ?", maxUsedConstraintID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedConstraintID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromMaxUsedConstraintID(maxUsedConstraintIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`max_used_constraint_id` IN (?)", maxUsedConstraintIDs).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllTableHistoryMgr) GetFromSessionID(sessionID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromPkComment 通过pk_comment获取内容
func (obj *_AllTableHistoryMgr) GetFromPkComment(pkComment string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`pk_comment` = ?", pkComment).Find(&results).Error

	return
}

// GetBatchFromPkComment 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromPkComment(pkComments []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`pk_comment` IN (?)", pkComments).Find(&results).Error

	return
}

// GetFromSessActiveTime 通过sess_active_time获取内容
func (obj *_AllTableHistoryMgr) GetFromSessActiveTime(sessActiveTime int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sess_active_time` = ?", sessActiveTime).Find(&results).Error

	return
}

// GetBatchFromSessActiveTime 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromSessActiveTime(sessActiveTimes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`sess_active_time` IN (?)", sessActiveTimes).Find(&results).Error

	return
}

// GetFromRowStoreType 通过row_store_type获取内容
func (obj *_AllTableHistoryMgr) GetFromRowStoreType(rowStoreType string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`row_store_type` = ?", rowStoreType).Find(&results).Error

	return
}

// GetBatchFromRowStoreType 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromRowStoreType(rowStoreTypes []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`row_store_type` IN (?)", rowStoreTypes).Find(&results).Error

	return
}

// GetFromStoreFormat 通过store_format获取内容
func (obj *_AllTableHistoryMgr) GetFromStoreFormat(storeFormat string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`store_format` = ?", storeFormat).Find(&results).Error

	return
}

// GetBatchFromStoreFormat 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromStoreFormat(storeFormats []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`store_format` IN (?)", storeFormats).Find(&results).Error

	return
}

// GetFromDuplicateScope 通过duplicate_scope获取内容
func (obj *_AllTableHistoryMgr) GetFromDuplicateScope(duplicateScope int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`duplicate_scope` = ?", duplicateScope).Find(&results).Error

	return
}

// GetBatchFromDuplicateScope 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromDuplicateScope(duplicateScopes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`duplicate_scope` IN (?)", duplicateScopes).Find(&results).Error

	return
}

// GetFromBinding 通过binding获取内容
func (obj *_AllTableHistoryMgr) GetFromBinding(binding int8) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`binding` = ?", binding).Find(&results).Error

	return
}

// GetBatchFromBinding 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromBinding(bindings []int8) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`binding` IN (?)", bindings).Find(&results).Error

	return
}

// GetFromProgressiveMergeRound 通过progressive_merge_round获取内容
func (obj *_AllTableHistoryMgr) GetFromProgressiveMergeRound(progressiveMergeRound int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`progressive_merge_round` = ?", progressiveMergeRound).Find(&results).Error

	return
}

// GetBatchFromProgressiveMergeRound 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromProgressiveMergeRound(progressiveMergeRounds []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`progressive_merge_round` IN (?)", progressiveMergeRounds).Find(&results).Error

	return
}

// GetFromStorageFormatVersion 通过storage_format_version获取内容
func (obj *_AllTableHistoryMgr) GetFromStorageFormatVersion(storageFormatVersion int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`storage_format_version` = ?", storageFormatVersion).Find(&results).Error

	return
}

// GetBatchFromStorageFormatVersion 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromStorageFormatVersion(storageFormatVersions []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`storage_format_version` IN (?)", storageFormatVersions).Find(&results).Error

	return
}

// GetFromTableMode 通过table_mode获取内容
func (obj *_AllTableHistoryMgr) GetFromTableMode(tableMode int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_mode` = ?", tableMode).Find(&results).Error

	return
}

// GetBatchFromTableMode 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTableMode(tableModes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`table_mode` IN (?)", tableModes).Find(&results).Error

	return
}

// GetFromEncryption 通过encryption获取内容
func (obj *_AllTableHistoryMgr) GetFromEncryption(encryption string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`encryption` = ?", encryption).Find(&results).Error

	return
}

// GetBatchFromEncryption 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromEncryption(encryptions []string) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`encryption` IN (?)", encryptions).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllTableHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllTableHistoryMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

// GetFromIsSubPartTemplate 通过is_sub_part_template获取内容
func (obj *_AllTableHistoryMgr) GetFromIsSubPartTemplate(isSubPartTemplate int8) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`is_sub_part_template` = ?", isSubPartTemplate).Find(&results).Error

	return
}

// GetBatchFromIsSubPartTemplate 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromIsSubPartTemplate(isSubPartTemplates []int8) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`is_sub_part_template` IN (?)", isSubPartTemplates).Find(&results).Error

	return
}

// GetFromDop 通过dop获取内容
func (obj *_AllTableHistoryMgr) GetFromDop(dop int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`dop` = ?", dop).Find(&results).Error

	return
}

// GetBatchFromDop 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromDop(dops []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`dop` IN (?)", dops).Find(&results).Error

	return
}

// GetFromCharacterSetClient 通过character_set_client获取内容
func (obj *_AllTableHistoryMgr) GetFromCharacterSetClient(characterSetClient int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`character_set_client` = ?", characterSetClient).Find(&results).Error

	return
}

// GetBatchFromCharacterSetClient 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromCharacterSetClient(characterSetClients []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`character_set_client` IN (?)", characterSetClients).Find(&results).Error

	return
}

// GetFromCollationConnection 通过collation_connection获取内容
func (obj *_AllTableHistoryMgr) GetFromCollationConnection(collationConnection int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`collation_connection` = ?", collationConnection).Find(&results).Error

	return
}

// GetBatchFromCollationConnection 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromCollationConnection(collationConnections []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`collation_connection` IN (?)", collationConnections).Find(&results).Error

	return
}

// GetFromAutoPartSize 通过auto_part_size获取内容
func (obj *_AllTableHistoryMgr) GetFromAutoPartSize(autoPartSize int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`auto_part_size` = ?", autoPartSize).Find(&results).Error

	return
}

// GetBatchFromAutoPartSize 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromAutoPartSize(autoPartSizes []int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`auto_part_size` IN (?)", autoPartSizes).Find(&results).Error

	return
}

// GetFromAutoPart 通过auto_part获取内容
func (obj *_AllTableHistoryMgr) GetFromAutoPart(autoPart int8) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`auto_part` = ?", autoPart).Find(&results).Error

	return
}

// GetBatchFromAutoPart 批量查找
func (obj *_AllTableHistoryMgr) GetBatchFromAutoPart(autoParts []int8) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`auto_part` IN (?)", autoParts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTableHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, schemaVersion int64) (result AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `schema_version` = ?", tenantID, tableID, schemaVersion).First(&result).Error

	return
}

// FetchIndexByIDxDataTableID  获取多个内容
func (obj *_AllTableHistoryMgr) FetchIndexByIDxDataTableID(dataTableID int64) (results []*AllTableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTableHistory{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}
