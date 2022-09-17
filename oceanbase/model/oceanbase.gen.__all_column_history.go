package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllColumnHistoryMgr struct {
	*_BaseMgr
}

// AllColumnHistoryMgr open func
func AllColumnHistoryMgr(db *gorm.DB) *_AllColumnHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllColumnHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllColumnHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_column_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllColumnHistoryMgr) GetTableName() string {
	return "__all_column_history"
}

// Reset 重置gorm会话
func (obj *_AllColumnHistoryMgr) Reset() *_AllColumnHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllColumnHistoryMgr) Get() (result AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllColumnHistoryMgr) Gets() (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllColumnHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllColumnHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllColumnHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllColumnHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllColumnHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithColumnID column_id获取
func (obj *_AllColumnHistoryMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllColumnHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllColumnHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithColumnName column_name获取
func (obj *_AllColumnHistoryMgr) WithColumnName(columnName string) Option {
	return optionFunc(func(o *options) { o.query["column_name"] = columnName })
}

// WithRowkeyPosition rowkey_position获取
func (obj *_AllColumnHistoryMgr) WithRowkeyPosition(rowkeyPosition int64) Option {
	return optionFunc(func(o *options) { o.query["rowkey_position"] = rowkeyPosition })
}

// WithIndexPosition index_position获取
func (obj *_AllColumnHistoryMgr) WithIndexPosition(indexPosition int64) Option {
	return optionFunc(func(o *options) { o.query["index_position"] = indexPosition })
}

// WithOrderInRowkey order_in_rowkey获取
func (obj *_AllColumnHistoryMgr) WithOrderInRowkey(orderInRowkey int64) Option {
	return optionFunc(func(o *options) { o.query["order_in_rowkey"] = orderInRowkey })
}

// WithPartitionKeyPosition partition_key_position获取
func (obj *_AllColumnHistoryMgr) WithPartitionKeyPosition(partitionKeyPosition int64) Option {
	return optionFunc(func(o *options) { o.query["partition_key_position"] = partitionKeyPosition })
}

// WithDataType data_type获取
func (obj *_AllColumnHistoryMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithDataLength data_length获取
func (obj *_AllColumnHistoryMgr) WithDataLength(dataLength int64) Option {
	return optionFunc(func(o *options) { o.query["data_length"] = dataLength })
}

// WithDataPrecision data_precision获取
func (obj *_AllColumnHistoryMgr) WithDataPrecision(dataPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["data_precision"] = dataPrecision })
}

// WithDataScale data_scale获取
func (obj *_AllColumnHistoryMgr) WithDataScale(dataScale int64) Option {
	return optionFunc(func(o *options) { o.query["data_scale"] = dataScale })
}

// WithZeroFill zero_fill获取
func (obj *_AllColumnHistoryMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithNullable nullable获取
func (obj *_AllColumnHistoryMgr) WithNullable(nullable int64) Option {
	return optionFunc(func(o *options) { o.query["nullable"] = nullable })
}

// WithOnUpdateCurrentTimestamp on_update_current_timestamp获取
func (obj *_AllColumnHistoryMgr) WithOnUpdateCurrentTimestamp(onUpdateCurrentTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["on_update_current_timestamp"] = onUpdateCurrentTimestamp })
}

// WithAutoincrement autoincrement获取
func (obj *_AllColumnHistoryMgr) WithAutoincrement(autoincrement int64) Option {
	return optionFunc(func(o *options) { o.query["autoincrement"] = autoincrement })
}

// WithIsHidden is_hidden获取
func (obj *_AllColumnHistoryMgr) WithIsHidden(isHidden int64) Option {
	return optionFunc(func(o *options) { o.query["is_hidden"] = isHidden })
}

// WithCollationType collation_type获取
func (obj *_AllColumnHistoryMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithOrigDefaultValue orig_default_value获取
func (obj *_AllColumnHistoryMgr) WithOrigDefaultValue(origDefaultValue string) Option {
	return optionFunc(func(o *options) { o.query["orig_default_value"] = origDefaultValue })
}

// WithCurDefaultValue cur_default_value获取
func (obj *_AllColumnHistoryMgr) WithCurDefaultValue(curDefaultValue string) Option {
	return optionFunc(func(o *options) { o.query["cur_default_value"] = curDefaultValue })
}

// WithComment comment获取
func (obj *_AllColumnHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithColumnFlags column_flags获取
func (obj *_AllColumnHistoryMgr) WithColumnFlags(columnFlags int64) Option {
	return optionFunc(func(o *options) { o.query["column_flags"] = columnFlags })
}

// WithPrevColumnID prev_column_id获取
func (obj *_AllColumnHistoryMgr) WithPrevColumnID(prevColumnID int64) Option {
	return optionFunc(func(o *options) { o.query["prev_column_id"] = prevColumnID })
}

// WithExtendedTypeInfo extended_type_info获取
func (obj *_AllColumnHistoryMgr) WithExtendedTypeInfo(extendedTypeInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["extended_type_info"] = extendedTypeInfo })
}

// WithOrigDefaultValueV2 orig_default_value_v2获取
func (obj *_AllColumnHistoryMgr) WithOrigDefaultValueV2(origDefaultValueV2 []byte) Option {
	return optionFunc(func(o *options) { o.query["orig_default_value_v2"] = origDefaultValueV2 })
}

// WithCurDefaultValueV2 cur_default_value_v2获取
func (obj *_AllColumnHistoryMgr) WithCurDefaultValueV2(curDefaultValueV2 []byte) Option {
	return optionFunc(func(o *options) { o.query["cur_default_value_v2"] = curDefaultValueV2 })
}

// GetByOption 功能选项模式获取
func (obj *_AllColumnHistoryMgr) GetByOption(opts ...Option) (result AllColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllColumnHistoryMgr) GetByOptions(opts ...Option) (results []*AllColumnHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllColumnHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllColumnHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where(options.query)
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
func (obj *_AllColumnHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllColumnHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllColumnHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllColumnHistoryMgr) GetFromTableID(tableID int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllColumnHistoryMgr) GetFromColumnID(columnID int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllColumnHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllColumnHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromColumnName 通过column_name获取内容
func (obj *_AllColumnHistoryMgr) GetFromColumnName(columnName string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`column_name` = ?", columnName).Find(&results).Error

	return
}

// GetBatchFromColumnName 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromColumnName(columnNames []string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`column_name` IN (?)", columnNames).Find(&results).Error

	return
}

// GetFromRowkeyPosition 通过rowkey_position获取内容
func (obj *_AllColumnHistoryMgr) GetFromRowkeyPosition(rowkeyPosition int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`rowkey_position` = ?", rowkeyPosition).Find(&results).Error

	return
}

// GetBatchFromRowkeyPosition 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromRowkeyPosition(rowkeyPositions []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`rowkey_position` IN (?)", rowkeyPositions).Find(&results).Error

	return
}

// GetFromIndexPosition 通过index_position获取内容
func (obj *_AllColumnHistoryMgr) GetFromIndexPosition(indexPosition int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`index_position` = ?", indexPosition).Find(&results).Error

	return
}

// GetBatchFromIndexPosition 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromIndexPosition(indexPositions []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`index_position` IN (?)", indexPositions).Find(&results).Error

	return
}

// GetFromOrderInRowkey 通过order_in_rowkey获取内容
func (obj *_AllColumnHistoryMgr) GetFromOrderInRowkey(orderInRowkey int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`order_in_rowkey` = ?", orderInRowkey).Find(&results).Error

	return
}

// GetBatchFromOrderInRowkey 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromOrderInRowkey(orderInRowkeys []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`order_in_rowkey` IN (?)", orderInRowkeys).Find(&results).Error

	return
}

// GetFromPartitionKeyPosition 通过partition_key_position获取内容
func (obj *_AllColumnHistoryMgr) GetFromPartitionKeyPosition(partitionKeyPosition int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`partition_key_position` = ?", partitionKeyPosition).Find(&results).Error

	return
}

// GetBatchFromPartitionKeyPosition 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromPartitionKeyPosition(partitionKeyPositions []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`partition_key_position` IN (?)", partitionKeyPositions).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllColumnHistoryMgr) GetFromDataType(dataType int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromDataLength 通过data_length获取内容
func (obj *_AllColumnHistoryMgr) GetFromDataLength(dataLength int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_length` = ?", dataLength).Find(&results).Error

	return
}

// GetBatchFromDataLength 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromDataLength(dataLengths []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_length` IN (?)", dataLengths).Find(&results).Error

	return
}

// GetFromDataPrecision 通过data_precision获取内容
func (obj *_AllColumnHistoryMgr) GetFromDataPrecision(dataPrecision int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_precision` = ?", dataPrecision).Find(&results).Error

	return
}

// GetBatchFromDataPrecision 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromDataPrecision(dataPrecisions []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_precision` IN (?)", dataPrecisions).Find(&results).Error

	return
}

// GetFromDataScale 通过data_scale获取内容
func (obj *_AllColumnHistoryMgr) GetFromDataScale(dataScale int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_scale` = ?", dataScale).Find(&results).Error

	return
}

// GetBatchFromDataScale 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromDataScale(dataScales []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`data_scale` IN (?)", dataScales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllColumnHistoryMgr) GetFromZeroFill(zeroFill int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromNullable 通过nullable获取内容
func (obj *_AllColumnHistoryMgr) GetFromNullable(nullable int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`nullable` = ?", nullable).Find(&results).Error

	return
}

// GetBatchFromNullable 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromNullable(nullables []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`nullable` IN (?)", nullables).Find(&results).Error

	return
}

// GetFromOnUpdateCurrentTimestamp 通过on_update_current_timestamp获取内容
func (obj *_AllColumnHistoryMgr) GetFromOnUpdateCurrentTimestamp(onUpdateCurrentTimestamp int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`on_update_current_timestamp` = ?", onUpdateCurrentTimestamp).Find(&results).Error

	return
}

// GetBatchFromOnUpdateCurrentTimestamp 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromOnUpdateCurrentTimestamp(onUpdateCurrentTimestamps []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`on_update_current_timestamp` IN (?)", onUpdateCurrentTimestamps).Find(&results).Error

	return
}

// GetFromAutoincrement 通过autoincrement获取内容
func (obj *_AllColumnHistoryMgr) GetFromAutoincrement(autoincrement int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`autoincrement` = ?", autoincrement).Find(&results).Error

	return
}

// GetBatchFromAutoincrement 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromAutoincrement(autoincrements []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`autoincrement` IN (?)", autoincrements).Find(&results).Error

	return
}

// GetFromIsHidden 通过is_hidden获取内容
func (obj *_AllColumnHistoryMgr) GetFromIsHidden(isHidden int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`is_hidden` = ?", isHidden).Find(&results).Error

	return
}

// GetBatchFromIsHidden 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromIsHidden(isHiddens []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`is_hidden` IN (?)", isHiddens).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllColumnHistoryMgr) GetFromCollationType(collationType int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromOrigDefaultValue 通过orig_default_value获取内容
func (obj *_AllColumnHistoryMgr) GetFromOrigDefaultValue(origDefaultValue string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`orig_default_value` = ?", origDefaultValue).Find(&results).Error

	return
}

// GetBatchFromOrigDefaultValue 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromOrigDefaultValue(origDefaultValues []string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`orig_default_value` IN (?)", origDefaultValues).Find(&results).Error

	return
}

// GetFromCurDefaultValue 通过cur_default_value获取内容
func (obj *_AllColumnHistoryMgr) GetFromCurDefaultValue(curDefaultValue string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`cur_default_value` = ?", curDefaultValue).Find(&results).Error

	return
}

// GetBatchFromCurDefaultValue 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromCurDefaultValue(curDefaultValues []string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`cur_default_value` IN (?)", curDefaultValues).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllColumnHistoryMgr) GetFromComment(comment string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromComment(comments []string) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromColumnFlags 通过column_flags获取内容
func (obj *_AllColumnHistoryMgr) GetFromColumnFlags(columnFlags int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`column_flags` = ?", columnFlags).Find(&results).Error

	return
}

// GetBatchFromColumnFlags 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromColumnFlags(columnFlagss []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`column_flags` IN (?)", columnFlagss).Find(&results).Error

	return
}

// GetFromPrevColumnID 通过prev_column_id获取内容
func (obj *_AllColumnHistoryMgr) GetFromPrevColumnID(prevColumnID int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`prev_column_id` = ?", prevColumnID).Find(&results).Error

	return
}

// GetBatchFromPrevColumnID 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromPrevColumnID(prevColumnIDs []int64) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`prev_column_id` IN (?)", prevColumnIDs).Find(&results).Error

	return
}

// GetFromExtendedTypeInfo 通过extended_type_info获取内容
func (obj *_AllColumnHistoryMgr) GetFromExtendedTypeInfo(extendedTypeInfo []byte) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`extended_type_info` = ?", extendedTypeInfo).Find(&results).Error

	return
}

// GetBatchFromExtendedTypeInfo 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromExtendedTypeInfo(extendedTypeInfos [][]byte) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`extended_type_info` IN (?)", extendedTypeInfos).Find(&results).Error

	return
}

// GetFromOrigDefaultValueV2 通过orig_default_value_v2获取内容
func (obj *_AllColumnHistoryMgr) GetFromOrigDefaultValueV2(origDefaultValueV2 []byte) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`orig_default_value_v2` = ?", origDefaultValueV2).Find(&results).Error

	return
}

// GetBatchFromOrigDefaultValueV2 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromOrigDefaultValueV2(origDefaultValueV2s [][]byte) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`orig_default_value_v2` IN (?)", origDefaultValueV2s).Find(&results).Error

	return
}

// GetFromCurDefaultValueV2 通过cur_default_value_v2获取内容
func (obj *_AllColumnHistoryMgr) GetFromCurDefaultValueV2(curDefaultValueV2 []byte) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`cur_default_value_v2` = ?", curDefaultValueV2).Find(&results).Error

	return
}

// GetBatchFromCurDefaultValueV2 批量查找
func (obj *_AllColumnHistoryMgr) GetBatchFromCurDefaultValueV2(curDefaultValueV2s [][]byte) (results []*AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`cur_default_value_v2` IN (?)", curDefaultValueV2s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllColumnHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, columnID int64, schemaVersion int64) (result AllColumnHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `column_id` = ? AND `schema_version` = ?", tenantID, tableID, columnID, schemaVersion).First(&result).Error

	return
}
