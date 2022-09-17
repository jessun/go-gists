package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualColumnHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualColumnHistoryMgr open func
func AllVirtualColumnHistoryMgr(db *gorm.DB) *_AllVirtualColumnHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualColumnHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualColumnHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_column_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualColumnHistoryMgr) GetTableName() string {
	return "__all_virtual_column_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualColumnHistoryMgr) Reset() *_AllVirtualColumnHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualColumnHistoryMgr) Get() (result AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualColumnHistoryMgr) Gets() (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualColumnHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualColumnHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualColumnHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithColumnID column_id获取
func (obj *_AllVirtualColumnHistoryMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualColumnHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualColumnHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualColumnHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualColumnHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithColumnName column_name获取
func (obj *_AllVirtualColumnHistoryMgr) WithColumnName(columnName string) Option {
	return optionFunc(func(o *options) { o.query["column_name"] = columnName })
}

// WithRowkeyPosition rowkey_position获取
func (obj *_AllVirtualColumnHistoryMgr) WithRowkeyPosition(rowkeyPosition int64) Option {
	return optionFunc(func(o *options) { o.query["rowkey_position"] = rowkeyPosition })
}

// WithIndexPosition index_position获取
func (obj *_AllVirtualColumnHistoryMgr) WithIndexPosition(indexPosition int64) Option {
	return optionFunc(func(o *options) { o.query["index_position"] = indexPosition })
}

// WithOrderInRowkey order_in_rowkey获取
func (obj *_AllVirtualColumnHistoryMgr) WithOrderInRowkey(orderInRowkey int64) Option {
	return optionFunc(func(o *options) { o.query["order_in_rowkey"] = orderInRowkey })
}

// WithPartitionKeyPosition partition_key_position获取
func (obj *_AllVirtualColumnHistoryMgr) WithPartitionKeyPosition(partitionKeyPosition int64) Option {
	return optionFunc(func(o *options) { o.query["partition_key_position"] = partitionKeyPosition })
}

// WithDataType data_type获取
func (obj *_AllVirtualColumnHistoryMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithDataLength data_length获取
func (obj *_AllVirtualColumnHistoryMgr) WithDataLength(dataLength int64) Option {
	return optionFunc(func(o *options) { o.query["data_length"] = dataLength })
}

// WithDataPrecision data_precision获取
func (obj *_AllVirtualColumnHistoryMgr) WithDataPrecision(dataPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["data_precision"] = dataPrecision })
}

// WithDataScale data_scale获取
func (obj *_AllVirtualColumnHistoryMgr) WithDataScale(dataScale int64) Option {
	return optionFunc(func(o *options) { o.query["data_scale"] = dataScale })
}

// WithZeroFill zero_fill获取
func (obj *_AllVirtualColumnHistoryMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithNullable nullable获取
func (obj *_AllVirtualColumnHistoryMgr) WithNullable(nullable int64) Option {
	return optionFunc(func(o *options) { o.query["nullable"] = nullable })
}

// WithOnUpdateCurrentTimestamp on_update_current_timestamp获取
func (obj *_AllVirtualColumnHistoryMgr) WithOnUpdateCurrentTimestamp(onUpdateCurrentTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["on_update_current_timestamp"] = onUpdateCurrentTimestamp })
}

// WithAutoincrement autoincrement获取
func (obj *_AllVirtualColumnHistoryMgr) WithAutoincrement(autoincrement int64) Option {
	return optionFunc(func(o *options) { o.query["autoincrement"] = autoincrement })
}

// WithIsHidden is_hidden获取
func (obj *_AllVirtualColumnHistoryMgr) WithIsHidden(isHidden int64) Option {
	return optionFunc(func(o *options) { o.query["is_hidden"] = isHidden })
}

// WithCollationType collation_type获取
func (obj *_AllVirtualColumnHistoryMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithOrigDefaultValue orig_default_value获取
func (obj *_AllVirtualColumnHistoryMgr) WithOrigDefaultValue(origDefaultValue string) Option {
	return optionFunc(func(o *options) { o.query["orig_default_value"] = origDefaultValue })
}

// WithCurDefaultValue cur_default_value获取
func (obj *_AllVirtualColumnHistoryMgr) WithCurDefaultValue(curDefaultValue string) Option {
	return optionFunc(func(o *options) { o.query["cur_default_value"] = curDefaultValue })
}

// WithComment comment获取
func (obj *_AllVirtualColumnHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithColumnFlags column_flags获取
func (obj *_AllVirtualColumnHistoryMgr) WithColumnFlags(columnFlags int64) Option {
	return optionFunc(func(o *options) { o.query["column_flags"] = columnFlags })
}

// WithPrevColumnID prev_column_id获取
func (obj *_AllVirtualColumnHistoryMgr) WithPrevColumnID(prevColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_column_id"] = prevColumnID })
}

// WithExtendedTypeInfo extended_type_info获取
func (obj *_AllVirtualColumnHistoryMgr) WithExtendedTypeInfo(extendedTypeInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["extended_type_info"] = extendedTypeInfo })
}

// WithOrigDefaultValueV2 orig_default_value_v2获取
func (obj *_AllVirtualColumnHistoryMgr) WithOrigDefaultValueV2(origDefaultValueV2 []byte) Option {
	return optionFunc(func(o *options) { o.query["orig_default_value_v2"] = origDefaultValueV2 })
}

// WithCurDefaultValueV2 cur_default_value_v2获取
func (obj *_AllVirtualColumnHistoryMgr) WithCurDefaultValueV2(curDefaultValueV2 []byte) Option {
	return optionFunc(func(o *options) { o.query["cur_default_value_v2"] = curDefaultValueV2 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualColumnHistoryMgr) GetByOption(opts ...Option) (result AllVirtualColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualColumnHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualColumnHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualColumnHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where(options.query)
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
func (obj *_AllVirtualColumnHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromTableID(tableID int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromColumnID(columnID int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromColumnName 通过column_name获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromColumnName(columnName string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`column_name` = ?", columnName).Find(&results).Error

	return
}

// GetBatchFromColumnName 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromColumnName(columnNames []string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`column_name` IN (?)", columnNames).Find(&results).Error

	return
}

// GetFromRowkeyPosition 通过rowkey_position获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromRowkeyPosition(rowkeyPosition int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`rowkey_position` = ?", rowkeyPosition).Find(&results).Error

	return
}

// GetBatchFromRowkeyPosition 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromRowkeyPosition(rowkeyPositions []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`rowkey_position` IN (?)", rowkeyPositions).Find(&results).Error

	return
}

// GetFromIndexPosition 通过index_position获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromIndexPosition(indexPosition int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`index_position` = ?", indexPosition).Find(&results).Error

	return
}

// GetBatchFromIndexPosition 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromIndexPosition(indexPositions []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`index_position` IN (?)", indexPositions).Find(&results).Error

	return
}

// GetFromOrderInRowkey 通过order_in_rowkey获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromOrderInRowkey(orderInRowkey int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`order_in_rowkey` = ?", orderInRowkey).Find(&results).Error

	return
}

// GetBatchFromOrderInRowkey 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromOrderInRowkey(orderInRowkeys []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`order_in_rowkey` IN (?)", orderInRowkeys).Find(&results).Error

	return
}

// GetFromPartitionKeyPosition 通过partition_key_position获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromPartitionKeyPosition(partitionKeyPosition int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`partition_key_position` = ?", partitionKeyPosition).Find(&results).Error

	return
}

// GetBatchFromPartitionKeyPosition 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromPartitionKeyPosition(partitionKeyPositions []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`partition_key_position` IN (?)", partitionKeyPositions).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromDataType(dataType int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromDataLength 通过data_length获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromDataLength(dataLength int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_length` = ?", dataLength).Find(&results).Error

	return
}

// GetBatchFromDataLength 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromDataLength(dataLengths []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_length` IN (?)", dataLengths).Find(&results).Error

	return
}

// GetFromDataPrecision 通过data_precision获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromDataPrecision(dataPrecision int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_precision` = ?", dataPrecision).Find(&results).Error

	return
}

// GetBatchFromDataPrecision 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromDataPrecision(dataPrecisions []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_precision` IN (?)", dataPrecisions).Find(&results).Error

	return
}

// GetFromDataScale 通过data_scale获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromDataScale(dataScale int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_scale` = ?", dataScale).Find(&results).Error

	return
}

// GetBatchFromDataScale 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromDataScale(dataScales []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`data_scale` IN (?)", dataScales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromZeroFill(zeroFill int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromNullable 通过nullable获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromNullable(nullable int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`nullable` = ?", nullable).Find(&results).Error

	return
}

// GetBatchFromNullable 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromNullable(nullables []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`nullable` IN (?)", nullables).Find(&results).Error

	return
}

// GetFromOnUpdateCurrentTimestamp 通过on_update_current_timestamp获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromOnUpdateCurrentTimestamp(onUpdateCurrentTimestamp int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`on_update_current_timestamp` = ?", onUpdateCurrentTimestamp).Find(&results).Error

	return
}

// GetBatchFromOnUpdateCurrentTimestamp 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromOnUpdateCurrentTimestamp(onUpdateCurrentTimestamps []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`on_update_current_timestamp` IN (?)", onUpdateCurrentTimestamps).Find(&results).Error

	return
}

// GetFromAutoincrement 通过autoincrement获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromAutoincrement(autoincrement int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`autoincrement` = ?", autoincrement).Find(&results).Error

	return
}

// GetBatchFromAutoincrement 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromAutoincrement(autoincrements []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`autoincrement` IN (?)", autoincrements).Find(&results).Error

	return
}

// GetFromIsHidden 通过is_hidden获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromIsHidden(isHidden int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`is_hidden` = ?", isHidden).Find(&results).Error

	return
}

// GetBatchFromIsHidden 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromIsHidden(isHiddens []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`is_hidden` IN (?)", isHiddens).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromCollationType(collationType int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromOrigDefaultValue 通过orig_default_value获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromOrigDefaultValue(origDefaultValue string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`orig_default_value` = ?", origDefaultValue).Find(&results).Error

	return
}

// GetBatchFromOrigDefaultValue 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromOrigDefaultValue(origDefaultValues []string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`orig_default_value` IN (?)", origDefaultValues).Find(&results).Error

	return
}

// GetFromCurDefaultValue 通过cur_default_value获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromCurDefaultValue(curDefaultValue string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`cur_default_value` = ?", curDefaultValue).Find(&results).Error

	return
}

// GetBatchFromCurDefaultValue 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromCurDefaultValue(curDefaultValues []string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`cur_default_value` IN (?)", curDefaultValues).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromComment(comment string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromComment(comments []string) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromColumnFlags 通过column_flags获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromColumnFlags(columnFlags int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`column_flags` = ?", columnFlags).Find(&results).Error

	return
}

// GetBatchFromColumnFlags 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromColumnFlags(columnFlagss []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`column_flags` IN (?)", columnFlagss).Find(&results).Error

	return
}

// GetFromPrevColumnID 通过prev_column_id获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromPrevColumnID(prevColumnID int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`prev_column_id` = ?", prevColumnID).Find(&results).Error

	return
}

// GetBatchFromPrevColumnID 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromPrevColumnID(prevColumnIDs []int64) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`prev_column_id` IN (?)", prevColumnIDs).Find(&results).Error

	return
}

// GetFromExtendedTypeInfo 通过extended_type_info获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromExtendedTypeInfo(extendedTypeInfo []byte) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`extended_type_info` = ?", extendedTypeInfo).Find(&results).Error

	return
}

// GetBatchFromExtendedTypeInfo 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromExtendedTypeInfo(extendedTypeInfos [][]byte) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`extended_type_info` IN (?)", extendedTypeInfos).Find(&results).Error

	return
}

// GetFromOrigDefaultValueV2 通过orig_default_value_v2获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromOrigDefaultValueV2(origDefaultValueV2 []byte) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`orig_default_value_v2` = ?", origDefaultValueV2).Find(&results).Error

	return
}

// GetBatchFromOrigDefaultValueV2 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromOrigDefaultValueV2(origDefaultValueV2s [][]byte) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`orig_default_value_v2` IN (?)", origDefaultValueV2s).Find(&results).Error

	return
}

// GetFromCurDefaultValueV2 通过cur_default_value_v2获取内容
func (obj *_AllVirtualColumnHistoryMgr) GetFromCurDefaultValueV2(curDefaultValueV2 []byte) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`cur_default_value_v2` = ?", curDefaultValueV2).Find(&results).Error

	return
}

// GetBatchFromCurDefaultValueV2 批量查找
func (obj *_AllVirtualColumnHistoryMgr) GetBatchFromCurDefaultValueV2(curDefaultValueV2s [][]byte) (results []*AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`cur_default_value_v2` IN (?)", curDefaultValueV2s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualColumnHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, columnID int64, schemaVersion int64) (result AllVirtualColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `column_id` = ? AND `schema_version` = ?", tenantID, tableID, columnID, schemaVersion).First(&result).Error

	return
}
