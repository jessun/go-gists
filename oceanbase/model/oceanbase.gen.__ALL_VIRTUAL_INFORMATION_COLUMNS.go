package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualInformationColumnsMgr struct {
	*_BaseMgr
}

// AllVirtualInformationColumnsMgr open func
func AllVirtualInformationColumnsMgr(db *gorm.DB) *_AllVirtualInformationColumnsMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualInformationColumnsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualInformationColumnsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__ALL_VIRTUAL_INFORMATION_COLUMNS"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualInformationColumnsMgr) GetTableName() string {
	return "__ALL_VIRTUAL_INFORMATION_COLUMNS"
}

// Reset 重置gorm会话
func (obj *_AllVirtualInformationColumnsMgr) Reset() *_AllVirtualInformationColumnsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualInformationColumnsMgr) Get() (result AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualInformationColumnsMgr) Gets() (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualInformationColumnsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableSchema TABLE_SCHEMA获取
func (obj *_AllVirtualInformationColumnsMgr) WithTableSchema(tableSchema string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_SCHEMA"] = tableSchema })
}

// WithTableName TABLE_NAME获取
func (obj *_AllVirtualInformationColumnsMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithTableCatalog TABLE_CATALOG获取
func (obj *_AllVirtualInformationColumnsMgr) WithTableCatalog(tableCatalog string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_CATALOG"] = tableCatalog })
}

// WithColumnName COLUMN_NAME获取
func (obj *_AllVirtualInformationColumnsMgr) WithColumnName(columnName string) Option {
	return optionFunc(func(o *options) { o.query["COLUMN_NAME"] = columnName })
}

// WithOrdinalPosition ORDINAL_POSITION获取
func (obj *_AllVirtualInformationColumnsMgr) WithOrdinalPosition(ordinalPosition uint64) Option {
	return optionFunc(func(o *options) { o.query["ORDINAL_POSITION"] = ordinalPosition })
}

// WithColumnDefault COLUMN_DEFAULT获取
func (obj *_AllVirtualInformationColumnsMgr) WithColumnDefault(columnDefault string) Option {
	return optionFunc(func(o *options) { o.query["COLUMN_DEFAULT"] = columnDefault })
}

// WithIsNullable IS_NULLABLE获取
func (obj *_AllVirtualInformationColumnsMgr) WithIsNullable(isNullable string) Option {
	return optionFunc(func(o *options) { o.query["IS_NULLABLE"] = isNullable })
}

// WithDataType DATA_TYPE获取
func (obj *_AllVirtualInformationColumnsMgr) WithDataType(dataType string) Option {
	return optionFunc(func(o *options) { o.query["DATA_TYPE"] = dataType })
}

// WithCharacterMaximumLength CHARACTER_MAXIMUM_LENGTH获取
func (obj *_AllVirtualInformationColumnsMgr) WithCharacterMaximumLength(characterMaximumLength uint64) Option {
	return optionFunc(func(o *options) { o.query["CHARACTER_MAXIMUM_LENGTH"] = characterMaximumLength })
}

// WithCharacterOctetLength CHARACTER_OCTET_LENGTH获取
func (obj *_AllVirtualInformationColumnsMgr) WithCharacterOctetLength(characterOctetLength uint64) Option {
	return optionFunc(func(o *options) { o.query["CHARACTER_OCTET_LENGTH"] = characterOctetLength })
}

// WithNumericPrecision NUMERIC_PRECISION获取
func (obj *_AllVirtualInformationColumnsMgr) WithNumericPrecision(numericPrecision uint64) Option {
	return optionFunc(func(o *options) { o.query["NUMERIC_PRECISION"] = numericPrecision })
}

// WithNumericScale NUMERIC_SCALE获取
func (obj *_AllVirtualInformationColumnsMgr) WithNumericScale(numericScale uint64) Option {
	return optionFunc(func(o *options) { o.query["NUMERIC_SCALE"] = numericScale })
}

// WithDatetimePrecision DATETIME_PRECISION获取
func (obj *_AllVirtualInformationColumnsMgr) WithDatetimePrecision(datetimePrecision uint64) Option {
	return optionFunc(func(o *options) { o.query["DATETIME_PRECISION"] = datetimePrecision })
}

// WithCharacterSetName CHARACTER_SET_NAME获取
func (obj *_AllVirtualInformationColumnsMgr) WithCharacterSetName(characterSetName string) Option {
	return optionFunc(func(o *options) { o.query["CHARACTER_SET_NAME"] = characterSetName })
}

// WithCollationName COLLATION_NAME获取
func (obj *_AllVirtualInformationColumnsMgr) WithCollationName(collationName string) Option {
	return optionFunc(func(o *options) { o.query["COLLATION_NAME"] = collationName })
}

// WithColumnType COLUMN_TYPE获取
func (obj *_AllVirtualInformationColumnsMgr) WithColumnType(columnType string) Option {
	return optionFunc(func(o *options) { o.query["COLUMN_TYPE"] = columnType })
}

// WithColumnKey COLUMN_KEY获取
func (obj *_AllVirtualInformationColumnsMgr) WithColumnKey(columnKey string) Option {
	return optionFunc(func(o *options) { o.query["COLUMN_KEY"] = columnKey })
}

// WithExtra EXTRA获取
func (obj *_AllVirtualInformationColumnsMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["EXTRA"] = extra })
}

// WithPrivileges PRIVILEGES获取
func (obj *_AllVirtualInformationColumnsMgr) WithPrivileges(privileges string) Option {
	return optionFunc(func(o *options) { o.query["PRIVILEGES"] = privileges })
}

// WithColumnComment COLUMN_COMMENT获取
func (obj *_AllVirtualInformationColumnsMgr) WithColumnComment(columnComment string) Option {
	return optionFunc(func(o *options) { o.query["COLUMN_COMMENT"] = columnComment })
}

// WithGenerationExpression GENERATION_EXPRESSION获取
func (obj *_AllVirtualInformationColumnsMgr) WithGenerationExpression(generationExpression string) Option {
	return optionFunc(func(o *options) { o.query["GENERATION_EXPRESSION"] = generationExpression })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualInformationColumnsMgr) GetByOption(opts ...Option) (result AllVirtualInformationColumns, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualInformationColumnsMgr) GetByOptions(opts ...Option) (results []*AllVirtualInformationColumns, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualInformationColumnsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualInformationColumns, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where(options.query)
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

// GetFromTableSchema 通过TABLE_SCHEMA获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromTableSchema(tableSchema string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`TABLE_SCHEMA` = ?", tableSchema).Find(&results).Error

	return
}

// GetBatchFromTableSchema 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromTableSchema(tableSchemas []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`TABLE_SCHEMA` IN (?)", tableSchemas).Find(&results).Error

	return
}

// GetFromTableName 通过TABLE_NAME获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromTableName(tableName string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromTableCatalog 通过TABLE_CATALOG获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromTableCatalog(tableCatalog string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`TABLE_CATALOG` = ?", tableCatalog).Find(&results).Error

	return
}

// GetBatchFromTableCatalog 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromTableCatalog(tableCatalogs []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`TABLE_CATALOG` IN (?)", tableCatalogs).Find(&results).Error

	return
}

// GetFromColumnName 通过COLUMN_NAME获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromColumnName(columnName string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_NAME` = ?", columnName).Find(&results).Error

	return
}

// GetBatchFromColumnName 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromColumnName(columnNames []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_NAME` IN (?)", columnNames).Find(&results).Error

	return
}

// GetFromOrdinalPosition 通过ORDINAL_POSITION获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromOrdinalPosition(ordinalPosition uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`ORDINAL_POSITION` = ?", ordinalPosition).Find(&results).Error

	return
}

// GetBatchFromOrdinalPosition 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromOrdinalPosition(ordinalPositions []uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`ORDINAL_POSITION` IN (?)", ordinalPositions).Find(&results).Error

	return
}

// GetFromColumnDefault 通过COLUMN_DEFAULT获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromColumnDefault(columnDefault string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_DEFAULT` = ?", columnDefault).Find(&results).Error

	return
}

// GetBatchFromColumnDefault 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromColumnDefault(columnDefaults []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_DEFAULT` IN (?)", columnDefaults).Find(&results).Error

	return
}

// GetFromIsNullable 通过IS_NULLABLE获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromIsNullable(isNullable string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`IS_NULLABLE` = ?", isNullable).Find(&results).Error

	return
}

// GetBatchFromIsNullable 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromIsNullable(isNullables []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`IS_NULLABLE` IN (?)", isNullables).Find(&results).Error

	return
}

// GetFromDataType 通过DATA_TYPE获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromDataType(dataType string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`DATA_TYPE` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromDataType(dataTypes []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`DATA_TYPE` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromCharacterMaximumLength 通过CHARACTER_MAXIMUM_LENGTH获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromCharacterMaximumLength(characterMaximumLength uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`CHARACTER_MAXIMUM_LENGTH` = ?", characterMaximumLength).Find(&results).Error

	return
}

// GetBatchFromCharacterMaximumLength 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromCharacterMaximumLength(characterMaximumLengths []uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`CHARACTER_MAXIMUM_LENGTH` IN (?)", characterMaximumLengths).Find(&results).Error

	return
}

// GetFromCharacterOctetLength 通过CHARACTER_OCTET_LENGTH获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromCharacterOctetLength(characterOctetLength uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`CHARACTER_OCTET_LENGTH` = ?", characterOctetLength).Find(&results).Error

	return
}

// GetBatchFromCharacterOctetLength 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromCharacterOctetLength(characterOctetLengths []uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`CHARACTER_OCTET_LENGTH` IN (?)", characterOctetLengths).Find(&results).Error

	return
}

// GetFromNumericPrecision 通过NUMERIC_PRECISION获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromNumericPrecision(numericPrecision uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`NUMERIC_PRECISION` = ?", numericPrecision).Find(&results).Error

	return
}

// GetBatchFromNumericPrecision 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromNumericPrecision(numericPrecisions []uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`NUMERIC_PRECISION` IN (?)", numericPrecisions).Find(&results).Error

	return
}

// GetFromNumericScale 通过NUMERIC_SCALE获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromNumericScale(numericScale uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`NUMERIC_SCALE` = ?", numericScale).Find(&results).Error

	return
}

// GetBatchFromNumericScale 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromNumericScale(numericScales []uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`NUMERIC_SCALE` IN (?)", numericScales).Find(&results).Error

	return
}

// GetFromDatetimePrecision 通过DATETIME_PRECISION获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromDatetimePrecision(datetimePrecision uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`DATETIME_PRECISION` = ?", datetimePrecision).Find(&results).Error

	return
}

// GetBatchFromDatetimePrecision 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromDatetimePrecision(datetimePrecisions []uint64) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`DATETIME_PRECISION` IN (?)", datetimePrecisions).Find(&results).Error

	return
}

// GetFromCharacterSetName 通过CHARACTER_SET_NAME获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromCharacterSetName(characterSetName string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`CHARACTER_SET_NAME` = ?", characterSetName).Find(&results).Error

	return
}

// GetBatchFromCharacterSetName 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromCharacterSetName(characterSetNames []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`CHARACTER_SET_NAME` IN (?)", characterSetNames).Find(&results).Error

	return
}

// GetFromCollationName 通过COLLATION_NAME获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromCollationName(collationName string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLLATION_NAME` = ?", collationName).Find(&results).Error

	return
}

// GetBatchFromCollationName 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromCollationName(collationNames []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLLATION_NAME` IN (?)", collationNames).Find(&results).Error

	return
}

// GetFromColumnType 通过COLUMN_TYPE获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromColumnType(columnType string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_TYPE` = ?", columnType).Find(&results).Error

	return
}

// GetBatchFromColumnType 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromColumnType(columnTypes []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_TYPE` IN (?)", columnTypes).Find(&results).Error

	return
}

// GetFromColumnKey 通过COLUMN_KEY获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromColumnKey(columnKey string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_KEY` = ?", columnKey).Find(&results).Error

	return
}

// GetBatchFromColumnKey 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromColumnKey(columnKeys []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_KEY` IN (?)", columnKeys).Find(&results).Error

	return
}

// GetFromExtra 通过EXTRA获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromExtra(extra string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`EXTRA` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromExtra(extras []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`EXTRA` IN (?)", extras).Find(&results).Error

	return
}

// GetFromPrivileges 通过PRIVILEGES获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromPrivileges(privileges string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`PRIVILEGES` = ?", privileges).Find(&results).Error

	return
}

// GetBatchFromPrivileges 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromPrivileges(privilegess []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`PRIVILEGES` IN (?)", privilegess).Find(&results).Error

	return
}

// GetFromColumnComment 通过COLUMN_COMMENT获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromColumnComment(columnComment string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_COMMENT` = ?", columnComment).Find(&results).Error

	return
}

// GetBatchFromColumnComment 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromColumnComment(columnComments []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`COLUMN_COMMENT` IN (?)", columnComments).Find(&results).Error

	return
}

// GetFromGenerationExpression 通过GENERATION_EXPRESSION获取内容
func (obj *_AllVirtualInformationColumnsMgr) GetFromGenerationExpression(generationExpression string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`GENERATION_EXPRESSION` = ?", generationExpression).Find(&results).Error

	return
}

// GetBatchFromGenerationExpression 批量查找
func (obj *_AllVirtualInformationColumnsMgr) GetBatchFromGenerationExpression(generationExpressions []string) (results []*AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`GENERATION_EXPRESSION` IN (?)", generationExpressions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualInformationColumnsMgr) FetchByPrimaryKey(tableSchema string, tableName string) (result AllVirtualInformationColumns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualInformationColumns{}).Where("`TABLE_SCHEMA` = ? AND `TABLE_NAME` = ?", tableSchema, tableName).First(&result).Error

	return
}
