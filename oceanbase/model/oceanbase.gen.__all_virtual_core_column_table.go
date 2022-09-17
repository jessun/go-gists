package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualCoreColumnTableMgr struct {
	*_BaseMgr
}

// AllVirtualCoreColumnTableMgr open func
func AllVirtualCoreColumnTableMgr(db *gorm.DB) *_AllVirtualCoreColumnTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualCoreColumnTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualCoreColumnTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_core_column_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualCoreColumnTableMgr) GetTableName() string {
	return "__all_virtual_core_column_table"
}

// Reset 重置gorm会话
func (obj *_AllVirtualCoreColumnTableMgr) Reset() *_AllVirtualCoreColumnTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualCoreColumnTableMgr) Get() (result AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualCoreColumnTableMgr) Gets() (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualCoreColumnTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualCoreColumnTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualCoreColumnTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithColumnID column_id获取
func (obj *_AllVirtualCoreColumnTableMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithColumnName column_name获取
func (obj *_AllVirtualCoreColumnTableMgr) WithColumnName(columnName string) Option {
	return optionFunc(func(o *options) { o.query["column_name"] = columnName })
}

// WithRowkeyPosition rowkey_position获取
func (obj *_AllVirtualCoreColumnTableMgr) WithRowkeyPosition(rowkeyPosition int64) Option {
	return optionFunc(func(o *options) { o.query["rowkey_position"] = rowkeyPosition })
}

// WithIndexPosition index_position获取
func (obj *_AllVirtualCoreColumnTableMgr) WithIndexPosition(indexPosition int64) Option {
	return optionFunc(func(o *options) { o.query["index_position"] = indexPosition })
}

// WithOrderInRowkey order_in_rowkey获取
func (obj *_AllVirtualCoreColumnTableMgr) WithOrderInRowkey(orderInRowkey int64) Option {
	return optionFunc(func(o *options) { o.query["order_in_rowkey"] = orderInRowkey })
}

// WithPartitionKeyPosition partition_key_position获取
func (obj *_AllVirtualCoreColumnTableMgr) WithPartitionKeyPosition(partitionKeyPosition int64) Option {
	return optionFunc(func(o *options) { o.query["partition_key_position"] = partitionKeyPosition })
}

// WithDataType data_type获取
func (obj *_AllVirtualCoreColumnTableMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithDataLength data_length获取
func (obj *_AllVirtualCoreColumnTableMgr) WithDataLength(dataLength int64) Option {
	return optionFunc(func(o *options) { o.query["data_length"] = dataLength })
}

// WithDataPrecision data_precision获取
func (obj *_AllVirtualCoreColumnTableMgr) WithDataPrecision(dataPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["data_precision"] = dataPrecision })
}

// WithDataScale data_scale获取
func (obj *_AllVirtualCoreColumnTableMgr) WithDataScale(dataScale int64) Option {
	return optionFunc(func(o *options) { o.query["data_scale"] = dataScale })
}

// WithZeroFill zero_fill获取
func (obj *_AllVirtualCoreColumnTableMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithNullable nullable获取
func (obj *_AllVirtualCoreColumnTableMgr) WithNullable(nullable int64) Option {
	return optionFunc(func(o *options) { o.query["nullable"] = nullable })
}

// WithOnUpdateCurrentTimestamp on_update_current_timestamp获取
func (obj *_AllVirtualCoreColumnTableMgr) WithOnUpdateCurrentTimestamp(onUpdateCurrentTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["on_update_current_timestamp"] = onUpdateCurrentTimestamp })
}

// WithAutoincrement autoincrement获取
func (obj *_AllVirtualCoreColumnTableMgr) WithAutoincrement(autoincrement int64) Option {
	return optionFunc(func(o *options) { o.query["autoincrement"] = autoincrement })
}

// WithIsHidden is_hidden获取
func (obj *_AllVirtualCoreColumnTableMgr) WithIsHidden(isHidden int64) Option {
	return optionFunc(func(o *options) { o.query["is_hidden"] = isHidden })
}

// WithCollationType collation_type获取
func (obj *_AllVirtualCoreColumnTableMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithOrigDefaultValue orig_default_value获取
func (obj *_AllVirtualCoreColumnTableMgr) WithOrigDefaultValue(origDefaultValue string) Option {
	return optionFunc(func(o *options) { o.query["orig_default_value"] = origDefaultValue })
}

// WithCurDefaultValue cur_default_value获取
func (obj *_AllVirtualCoreColumnTableMgr) WithCurDefaultValue(curDefaultValue string) Option {
	return optionFunc(func(o *options) { o.query["cur_default_value"] = curDefaultValue })
}

// WithComment comment获取
func (obj *_AllVirtualCoreColumnTableMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualCoreColumnTableMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithColumnFlags column_flags获取
func (obj *_AllVirtualCoreColumnTableMgr) WithColumnFlags(columnFlags int64) Option {
	return optionFunc(func(o *options) { o.query["column_flags"] = columnFlags })
}

// WithPrevColumnID prev_column_id获取
func (obj *_AllVirtualCoreColumnTableMgr) WithPrevColumnID(prevColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_column_id"] = prevColumnID })
}

// WithExtendedTypeInfo extended_type_info获取
func (obj *_AllVirtualCoreColumnTableMgr) WithExtendedTypeInfo(extendedTypeInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["extended_type_info"] = extendedTypeInfo })
}

// WithOrigDefaultValueV2 orig_default_value_v2获取
func (obj *_AllVirtualCoreColumnTableMgr) WithOrigDefaultValueV2(origDefaultValueV2 []byte) Option {
	return optionFunc(func(o *options) { o.query["orig_default_value_v2"] = origDefaultValueV2 })
}

// WithCurDefaultValueV2 cur_default_value_v2获取
func (obj *_AllVirtualCoreColumnTableMgr) WithCurDefaultValueV2(curDefaultValueV2 []byte) Option {
	return optionFunc(func(o *options) { o.query["cur_default_value_v2"] = curDefaultValueV2 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualCoreColumnTableMgr) GetByOption(opts ...Option) (result AllVirtualCoreColumnTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualCoreColumnTableMgr) GetByOptions(opts ...Option) (results []*AllVirtualCoreColumnTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualCoreColumnTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualCoreColumnTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where(options.query)
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
func (obj *_AllVirtualCoreColumnTableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromTableID(tableID int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromColumnID(columnID int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromColumnName 通过column_name获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromColumnName(columnName string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`column_name` = ?", columnName).Find(&results).Error

	return
}

// GetBatchFromColumnName 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromColumnName(columnNames []string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`column_name` IN (?)", columnNames).Find(&results).Error

	return
}

// GetFromRowkeyPosition 通过rowkey_position获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromRowkeyPosition(rowkeyPosition int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`rowkey_position` = ?", rowkeyPosition).Find(&results).Error

	return
}

// GetBatchFromRowkeyPosition 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromRowkeyPosition(rowkeyPositions []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`rowkey_position` IN (?)", rowkeyPositions).Find(&results).Error

	return
}

// GetFromIndexPosition 通过index_position获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromIndexPosition(indexPosition int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`index_position` = ?", indexPosition).Find(&results).Error

	return
}

// GetBatchFromIndexPosition 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromIndexPosition(indexPositions []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`index_position` IN (?)", indexPositions).Find(&results).Error

	return
}

// GetFromOrderInRowkey 通过order_in_rowkey获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromOrderInRowkey(orderInRowkey int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`order_in_rowkey` = ?", orderInRowkey).Find(&results).Error

	return
}

// GetBatchFromOrderInRowkey 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromOrderInRowkey(orderInRowkeys []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`order_in_rowkey` IN (?)", orderInRowkeys).Find(&results).Error

	return
}

// GetFromPartitionKeyPosition 通过partition_key_position获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromPartitionKeyPosition(partitionKeyPosition int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`partition_key_position` = ?", partitionKeyPosition).Find(&results).Error

	return
}

// GetBatchFromPartitionKeyPosition 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromPartitionKeyPosition(partitionKeyPositions []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`partition_key_position` IN (?)", partitionKeyPositions).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromDataType(dataType int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromDataLength 通过data_length获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromDataLength(dataLength int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_length` = ?", dataLength).Find(&results).Error

	return
}

// GetBatchFromDataLength 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromDataLength(dataLengths []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_length` IN (?)", dataLengths).Find(&results).Error

	return
}

// GetFromDataPrecision 通过data_precision获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromDataPrecision(dataPrecision int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_precision` = ?", dataPrecision).Find(&results).Error

	return
}

// GetBatchFromDataPrecision 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromDataPrecision(dataPrecisions []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_precision` IN (?)", dataPrecisions).Find(&results).Error

	return
}

// GetFromDataScale 通过data_scale获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromDataScale(dataScale int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_scale` = ?", dataScale).Find(&results).Error

	return
}

// GetBatchFromDataScale 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromDataScale(dataScales []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`data_scale` IN (?)", dataScales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromZeroFill(zeroFill int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromNullable 通过nullable获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromNullable(nullable int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`nullable` = ?", nullable).Find(&results).Error

	return
}

// GetBatchFromNullable 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromNullable(nullables []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`nullable` IN (?)", nullables).Find(&results).Error

	return
}

// GetFromOnUpdateCurrentTimestamp 通过on_update_current_timestamp获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromOnUpdateCurrentTimestamp(onUpdateCurrentTimestamp int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`on_update_current_timestamp` = ?", onUpdateCurrentTimestamp).Find(&results).Error

	return
}

// GetBatchFromOnUpdateCurrentTimestamp 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromOnUpdateCurrentTimestamp(onUpdateCurrentTimestamps []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`on_update_current_timestamp` IN (?)", onUpdateCurrentTimestamps).Find(&results).Error

	return
}

// GetFromAutoincrement 通过autoincrement获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromAutoincrement(autoincrement int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`autoincrement` = ?", autoincrement).Find(&results).Error

	return
}

// GetBatchFromAutoincrement 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromAutoincrement(autoincrements []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`autoincrement` IN (?)", autoincrements).Find(&results).Error

	return
}

// GetFromIsHidden 通过is_hidden获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromIsHidden(isHidden int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`is_hidden` = ?", isHidden).Find(&results).Error

	return
}

// GetBatchFromIsHidden 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromIsHidden(isHiddens []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`is_hidden` IN (?)", isHiddens).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromCollationType(collationType int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromOrigDefaultValue 通过orig_default_value获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromOrigDefaultValue(origDefaultValue string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`orig_default_value` = ?", origDefaultValue).Find(&results).Error

	return
}

// GetBatchFromOrigDefaultValue 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromOrigDefaultValue(origDefaultValues []string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`orig_default_value` IN (?)", origDefaultValues).Find(&results).Error

	return
}

// GetFromCurDefaultValue 通过cur_default_value获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromCurDefaultValue(curDefaultValue string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`cur_default_value` = ?", curDefaultValue).Find(&results).Error

	return
}

// GetBatchFromCurDefaultValue 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromCurDefaultValue(curDefaultValues []string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`cur_default_value` IN (?)", curDefaultValues).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromComment(comment string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromComment(comments []string) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromColumnFlags 通过column_flags获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromColumnFlags(columnFlags int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`column_flags` = ?", columnFlags).Find(&results).Error

	return
}

// GetBatchFromColumnFlags 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromColumnFlags(columnFlagss []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`column_flags` IN (?)", columnFlagss).Find(&results).Error

	return
}

// GetFromPrevColumnID 通过prev_column_id获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromPrevColumnID(prevColumnID int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`prev_column_id` = ?", prevColumnID).Find(&results).Error

	return
}

// GetBatchFromPrevColumnID 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromPrevColumnID(prevColumnIDs []int64) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`prev_column_id` IN (?)", prevColumnIDs).Find(&results).Error

	return
}

// GetFromExtendedTypeInfo 通过extended_type_info获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromExtendedTypeInfo(extendedTypeInfo []byte) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`extended_type_info` = ?", extendedTypeInfo).Find(&results).Error

	return
}

// GetBatchFromExtendedTypeInfo 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromExtendedTypeInfo(extendedTypeInfos [][]byte) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`extended_type_info` IN (?)", extendedTypeInfos).Find(&results).Error

	return
}

// GetFromOrigDefaultValueV2 通过orig_default_value_v2获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromOrigDefaultValueV2(origDefaultValueV2 []byte) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`orig_default_value_v2` = ?", origDefaultValueV2).Find(&results).Error

	return
}

// GetBatchFromOrigDefaultValueV2 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromOrigDefaultValueV2(origDefaultValueV2s [][]byte) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`orig_default_value_v2` IN (?)", origDefaultValueV2s).Find(&results).Error

	return
}

// GetFromCurDefaultValueV2 通过cur_default_value_v2获取内容
func (obj *_AllVirtualCoreColumnTableMgr) GetFromCurDefaultValueV2(curDefaultValueV2 []byte) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`cur_default_value_v2` = ?", curDefaultValueV2).Find(&results).Error

	return
}

// GetBatchFromCurDefaultValueV2 批量查找
func (obj *_AllVirtualCoreColumnTableMgr) GetBatchFromCurDefaultValueV2(curDefaultValueV2s [][]byte) (results []*AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`cur_default_value_v2` IN (?)", curDefaultValueV2s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualCoreColumnTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64, columnID int64) (result AllVirtualCoreColumnTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCoreColumnTable{}).Where("`tenant_id` = ? AND `table_id` = ? AND `column_id` = ?", tenantID, tableID, columnID).First(&result).Error

	return
}
