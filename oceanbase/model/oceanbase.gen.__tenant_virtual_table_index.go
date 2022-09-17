package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualTableIndexMgr struct {
	*_BaseMgr
}

// TenantVirtualTableIndexMgr open func
func TenantVirtualTableIndexMgr(db *gorm.DB) *_TenantVirtualTableIndexMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualTableIndexMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualTableIndexMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_table_index"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualTableIndexMgr) GetTableName() string {
	return "__tenant_virtual_table_index"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualTableIndexMgr) Reset() *_TenantVirtualTableIndexMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualTableIndexMgr) Get() (result TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualTableIndexMgr) Gets() (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualTableIndexMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_TenantVirtualTableIndexMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithKeyName key_name获取
func (obj *_TenantVirtualTableIndexMgr) WithKeyName(keyName string) Option {
	return optionFunc(func(o *options) { o.query["key_name"] = keyName })
}

// WithSeqInIndex seq_in_index获取
func (obj *_TenantVirtualTableIndexMgr) WithSeqInIndex(seqInIndex int64) Option {
	return optionFunc(func(o *options) { o.query["seq_in_index"] = seqInIndex })
}

// WithTableSchema table_schema获取
func (obj *_TenantVirtualTableIndexMgr) WithTableSchema(tableSchema string) Option {
	return optionFunc(func(o *options) { o.query["table_schema"] = tableSchema })
}

// WithTable table获取
func (obj *_TenantVirtualTableIndexMgr) WithTable(table string) Option {
	return optionFunc(func(o *options) { o.query["table"] = table })
}

// WithNonUnique non_unique获取
func (obj *_TenantVirtualTableIndexMgr) WithNonUnique(nonUnique int64) Option {
	return optionFunc(func(o *options) { o.query["non_unique"] = nonUnique })
}

// WithIndexSchema index_schema获取
func (obj *_TenantVirtualTableIndexMgr) WithIndexSchema(indexSchema string) Option {
	return optionFunc(func(o *options) { o.query["index_schema"] = indexSchema })
}

// WithColumnName column_name获取
func (obj *_TenantVirtualTableIndexMgr) WithColumnName(columnName string) Option {
	return optionFunc(func(o *options) { o.query["column_name"] = columnName })
}

// WithCollation collation获取
func (obj *_TenantVirtualTableIndexMgr) WithCollation(collation string) Option {
	return optionFunc(func(o *options) { o.query["collation"] = collation })
}

// WithCardinality cardinality获取
func (obj *_TenantVirtualTableIndexMgr) WithCardinality(cardinality int64) Option {
	return optionFunc(func(o *options) { o.query["cardinality"] = cardinality })
}

// WithSubPart sub_part获取
func (obj *_TenantVirtualTableIndexMgr) WithSubPart(subPart string) Option {
	return optionFunc(func(o *options) { o.query["sub_part"] = subPart })
}

// WithPacked packed获取
func (obj *_TenantVirtualTableIndexMgr) WithPacked(packed string) Option {
	return optionFunc(func(o *options) { o.query["packed"] = packed })
}

// WithNull null获取
func (obj *_TenantVirtualTableIndexMgr) WithNull(null string) Option {
	return optionFunc(func(o *options) { o.query["null"] = null })
}

// WithIndexType index_type获取
func (obj *_TenantVirtualTableIndexMgr) WithIndexType(indexType string) Option {
	return optionFunc(func(o *options) { o.query["index_type"] = indexType })
}

// WithComment comment获取
func (obj *_TenantVirtualTableIndexMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithIndexComment index_comment获取
func (obj *_TenantVirtualTableIndexMgr) WithIndexComment(indexComment string) Option {
	return optionFunc(func(o *options) { o.query["index_comment"] = indexComment })
}

// WithIsVisible is_visible获取
func (obj *_TenantVirtualTableIndexMgr) WithIsVisible(isVisible string) Option {
	return optionFunc(func(o *options) { o.query["is_visible"] = isVisible })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualTableIndexMgr) GetByOption(opts ...Option) (result TenantVirtualTableIndex, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualTableIndexMgr) GetByOptions(opts ...Option) (results []*TenantVirtualTableIndex, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualTableIndexMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualTableIndex, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where(options.query)
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

// GetFromTableID 通过table_id获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromTableID(tableID int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromTableID(tableIDs []int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromKeyName 通过key_name获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromKeyName(keyName string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`key_name` = ?", keyName).Find(&results).Error

	return
}

// GetBatchFromKeyName 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromKeyName(keyNames []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`key_name` IN (?)", keyNames).Find(&results).Error

	return
}

// GetFromSeqInIndex 通过seq_in_index获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromSeqInIndex(seqInIndex int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`seq_in_index` = ?", seqInIndex).Find(&results).Error

	return
}

// GetBatchFromSeqInIndex 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromSeqInIndex(seqInIndexs []int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`seq_in_index` IN (?)", seqInIndexs).Find(&results).Error

	return
}

// GetFromTableSchema 通过table_schema获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromTableSchema(tableSchema string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`table_schema` = ?", tableSchema).Find(&results).Error

	return
}

// GetBatchFromTableSchema 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromTableSchema(tableSchemas []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`table_schema` IN (?)", tableSchemas).Find(&results).Error

	return
}

// GetFromTable 通过table获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromTable(table string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`table` = ?", table).Find(&results).Error

	return
}

// GetBatchFromTable 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromTable(tables []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`table` IN (?)", tables).Find(&results).Error

	return
}

// GetFromNonUnique 通过non_unique获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromNonUnique(nonUnique int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`non_unique` = ?", nonUnique).Find(&results).Error

	return
}

// GetBatchFromNonUnique 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromNonUnique(nonUniques []int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`non_unique` IN (?)", nonUniques).Find(&results).Error

	return
}

// GetFromIndexSchema 通过index_schema获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromIndexSchema(indexSchema string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`index_schema` = ?", indexSchema).Find(&results).Error

	return
}

// GetBatchFromIndexSchema 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromIndexSchema(indexSchemas []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`index_schema` IN (?)", indexSchemas).Find(&results).Error

	return
}

// GetFromColumnName 通过column_name获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromColumnName(columnName string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`column_name` = ?", columnName).Find(&results).Error

	return
}

// GetBatchFromColumnName 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromColumnName(columnNames []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`column_name` IN (?)", columnNames).Find(&results).Error

	return
}

// GetFromCollation 通过collation获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromCollation(collation string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`collation` = ?", collation).Find(&results).Error

	return
}

// GetBatchFromCollation 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromCollation(collations []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`collation` IN (?)", collations).Find(&results).Error

	return
}

// GetFromCardinality 通过cardinality获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromCardinality(cardinality int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`cardinality` = ?", cardinality).Find(&results).Error

	return
}

// GetBatchFromCardinality 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromCardinality(cardinalitys []int64) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`cardinality` IN (?)", cardinalitys).Find(&results).Error

	return
}

// GetFromSubPart 通过sub_part获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromSubPart(subPart string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`sub_part` = ?", subPart).Find(&results).Error

	return
}

// GetBatchFromSubPart 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromSubPart(subParts []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`sub_part` IN (?)", subParts).Find(&results).Error

	return
}

// GetFromPacked 通过packed获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromPacked(packed string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`packed` = ?", packed).Find(&results).Error

	return
}

// GetBatchFromPacked 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromPacked(packeds []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`packed` IN (?)", packeds).Find(&results).Error

	return
}

// GetFromNull 通过null获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromNull(null string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`null` = ?", null).Find(&results).Error

	return
}

// GetBatchFromNull 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromNull(nulls []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`null` IN (?)", nulls).Find(&results).Error

	return
}

// GetFromIndexType 通过index_type获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromIndexType(indexType string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`index_type` = ?", indexType).Find(&results).Error

	return
}

// GetBatchFromIndexType 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromIndexType(indexTypes []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`index_type` IN (?)", indexTypes).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromComment(comment string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromComment(comments []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromIndexComment 通过index_comment获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromIndexComment(indexComment string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`index_comment` = ?", indexComment).Find(&results).Error

	return
}

// GetBatchFromIndexComment 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromIndexComment(indexComments []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`index_comment` IN (?)", indexComments).Find(&results).Error

	return
}

// GetFromIsVisible 通过is_visible获取内容
func (obj *_TenantVirtualTableIndexMgr) GetFromIsVisible(isVisible string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`is_visible` = ?", isVisible).Find(&results).Error

	return
}

// GetBatchFromIsVisible 批量查找
func (obj *_TenantVirtualTableIndexMgr) GetBatchFromIsVisible(isVisibles []string) (results []*TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`is_visible` IN (?)", isVisibles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualTableIndexMgr) FetchByPrimaryKey(tableID int64, keyName string, seqInIndex int64) (result TenantVirtualTableIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTableIndex{}).Where("`table_id` = ? AND `key_name` = ? AND `seq_in_index` = ?", tableID, keyName, seqInIndex).First(&result).Error

	return
}
