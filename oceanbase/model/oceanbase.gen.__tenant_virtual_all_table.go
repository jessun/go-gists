package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TenantVirtualAllTableMgr struct {
	*_BaseMgr
}

// TenantVirtualAllTableMgr open func
func TenantVirtualAllTableMgr(db *gorm.DB) *_TenantVirtualAllTableMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualAllTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualAllTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_all_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualAllTableMgr) GetTableName() string {
	return "__tenant_virtual_all_table"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualAllTableMgr) Reset() *_TenantVirtualAllTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualAllTableMgr) Get() (result TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualAllTableMgr) Gets() (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualAllTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDatabaseID database_id获取
func (obj *_TenantVirtualAllTableMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithTableName table_name获取
func (obj *_TenantVirtualAllTableMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithTableType table_type获取
func (obj *_TenantVirtualAllTableMgr) WithTableType(tableType string) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithEngine engine获取
func (obj *_TenantVirtualAllTableMgr) WithEngine(engine string) Option {
	return optionFunc(func(o *options) { o.query["engine"] = engine })
}

// WithVersion version获取
func (obj *_TenantVirtualAllTableMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithRowFormat row_format获取
func (obj *_TenantVirtualAllTableMgr) WithRowFormat(rowFormat string) Option {
	return optionFunc(func(o *options) { o.query["row_format"] = rowFormat })
}

// WithRows rows获取
func (obj *_TenantVirtualAllTableMgr) WithRows(rows int64) Option {
	return optionFunc(func(o *options) { o.query["rows"] = rows })
}

// WithAvgRowLength avg_row_length获取
func (obj *_TenantVirtualAllTableMgr) WithAvgRowLength(avgRowLength int64) Option {
	return optionFunc(func(o *options) { o.query["avg_row_length"] = avgRowLength })
}

// WithDataLength data_length获取
func (obj *_TenantVirtualAllTableMgr) WithDataLength(dataLength int64) Option {
	return optionFunc(func(o *options) { o.query["data_length"] = dataLength })
}

// WithMaxDataLength max_data_length获取
func (obj *_TenantVirtualAllTableMgr) WithMaxDataLength(maxDataLength int64) Option {
	return optionFunc(func(o *options) { o.query["max_data_length"] = maxDataLength })
}

// WithIndexLength index_length获取
func (obj *_TenantVirtualAllTableMgr) WithIndexLength(indexLength int64) Option {
	return optionFunc(func(o *options) { o.query["index_length"] = indexLength })
}

// WithDataFree data_free获取
func (obj *_TenantVirtualAllTableMgr) WithDataFree(dataFree int64) Option {
	return optionFunc(func(o *options) { o.query["data_free"] = dataFree })
}

// WithAutoIncrement auto_increment获取
func (obj *_TenantVirtualAllTableMgr) WithAutoIncrement(autoIncrement uint64) Option {
	return optionFunc(func(o *options) { o.query["auto_increment"] = autoIncrement })
}

// WithCreateTime create_time获取
func (obj *_TenantVirtualAllTableMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_TenantVirtualAllTableMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCheckTime check_time获取
func (obj *_TenantVirtualAllTableMgr) WithCheckTime(checkTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["check_time"] = checkTime })
}

// WithCollation collation获取
func (obj *_TenantVirtualAllTableMgr) WithCollation(collation string) Option {
	return optionFunc(func(o *options) { o.query["collation"] = collation })
}

// WithChecksum checksum获取
func (obj *_TenantVirtualAllTableMgr) WithChecksum(checksum int64) Option {
	return optionFunc(func(o *options) { o.query["checksum"] = checksum })
}

// WithCreateOptions create_options获取
func (obj *_TenantVirtualAllTableMgr) WithCreateOptions(createOptions string) Option {
	return optionFunc(func(o *options) { o.query["create_options"] = createOptions })
}

// WithComment comment获取
func (obj *_TenantVirtualAllTableMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualAllTableMgr) GetByOption(opts ...Option) (result TenantVirtualAllTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualAllTableMgr) GetByOptions(opts ...Option) (results []*TenantVirtualAllTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualAllTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualAllTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where(options.query)
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

// GetFromDatabaseID 通过database_id获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromDatabaseID(databaseID int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromTableName(tableName string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromTableName(tableNames []string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromTableType(tableType string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromTableType(tableTypes []string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

// GetFromEngine 通过engine获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromEngine(engine string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`engine` = ?", engine).Find(&results).Error

	return
}

// GetBatchFromEngine 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromEngine(engines []string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`engine` IN (?)", engines).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromVersion(version int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromVersion(versions []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromRowFormat 通过row_format获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromRowFormat(rowFormat string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`row_format` = ?", rowFormat).Find(&results).Error

	return
}

// GetBatchFromRowFormat 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromRowFormat(rowFormats []string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`row_format` IN (?)", rowFormats).Find(&results).Error

	return
}

// GetFromRows 通过rows获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromRows(rows int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`rows` = ?", rows).Find(&results).Error

	return
}

// GetBatchFromRows 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromRows(rowss []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`rows` IN (?)", rowss).Find(&results).Error

	return
}

// GetFromAvgRowLength 通过avg_row_length获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromAvgRowLength(avgRowLength int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`avg_row_length` = ?", avgRowLength).Find(&results).Error

	return
}

// GetBatchFromAvgRowLength 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromAvgRowLength(avgRowLengths []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`avg_row_length` IN (?)", avgRowLengths).Find(&results).Error

	return
}

// GetFromDataLength 通过data_length获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromDataLength(dataLength int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`data_length` = ?", dataLength).Find(&results).Error

	return
}

// GetBatchFromDataLength 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromDataLength(dataLengths []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`data_length` IN (?)", dataLengths).Find(&results).Error

	return
}

// GetFromMaxDataLength 通过max_data_length获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromMaxDataLength(maxDataLength int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`max_data_length` = ?", maxDataLength).Find(&results).Error

	return
}

// GetBatchFromMaxDataLength 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromMaxDataLength(maxDataLengths []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`max_data_length` IN (?)", maxDataLengths).Find(&results).Error

	return
}

// GetFromIndexLength 通过index_length获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromIndexLength(indexLength int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`index_length` = ?", indexLength).Find(&results).Error

	return
}

// GetBatchFromIndexLength 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromIndexLength(indexLengths []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`index_length` IN (?)", indexLengths).Find(&results).Error

	return
}

// GetFromDataFree 通过data_free获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromDataFree(dataFree int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`data_free` = ?", dataFree).Find(&results).Error

	return
}

// GetBatchFromDataFree 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromDataFree(dataFrees []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`data_free` IN (?)", dataFrees).Find(&results).Error

	return
}

// GetFromAutoIncrement 通过auto_increment获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromAutoIncrement(autoIncrement uint64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`auto_increment` = ?", autoIncrement).Find(&results).Error

	return
}

// GetBatchFromAutoIncrement 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromAutoIncrement(autoIncrements []uint64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`auto_increment` IN (?)", autoIncrements).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromCreateTime(createTime time.Time) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromUpdateTime(updateTime time.Time) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCheckTime 通过check_time获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromCheckTime(checkTime time.Time) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`check_time` = ?", checkTime).Find(&results).Error

	return
}

// GetBatchFromCheckTime 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromCheckTime(checkTimes []time.Time) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`check_time` IN (?)", checkTimes).Find(&results).Error

	return
}

// GetFromCollation 通过collation获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromCollation(collation string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`collation` = ?", collation).Find(&results).Error

	return
}

// GetBatchFromCollation 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromCollation(collations []string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`collation` IN (?)", collations).Find(&results).Error

	return
}

// GetFromChecksum 通过checksum获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromChecksum(checksum int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`checksum` = ?", checksum).Find(&results).Error

	return
}

// GetBatchFromChecksum 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromChecksum(checksums []int64) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`checksum` IN (?)", checksums).Find(&results).Error

	return
}

// GetFromCreateOptions 通过create_options获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromCreateOptions(createOptions string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`create_options` = ?", createOptions).Find(&results).Error

	return
}

// GetBatchFromCreateOptions 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromCreateOptions(createOptionss []string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`create_options` IN (?)", createOptionss).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_TenantVirtualAllTableMgr) GetFromComment(comment string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_TenantVirtualAllTableMgr) GetBatchFromComment(comments []string) (results []*TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualAllTableMgr) FetchByPrimaryKey(databaseID int64, tableName string) (result TenantVirtualAllTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualAllTable{}).Where("`database_id` = ? AND `table_name` = ?", databaseID, tableName).First(&result).Error

	return
}
