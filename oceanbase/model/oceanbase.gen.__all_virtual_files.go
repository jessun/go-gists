package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualFilesMgr struct {
	*_BaseMgr
}

// AllVirtualFilesMgr open func
func AllVirtualFilesMgr(db *gorm.DB) *_AllVirtualFilesMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualFilesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualFilesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_files"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualFilesMgr) GetTableName() string {
	return "__all_virtual_files"
}

// Reset 重置gorm会话
func (obj *_AllVirtualFilesMgr) Reset() *_AllVirtualFilesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualFilesMgr) Get() (result AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualFilesMgr) Gets() (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualFilesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithFileID FILE_ID获取
func (obj *_AllVirtualFilesMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["FILE_ID"] = fileID })
}

// WithFileName FILE_NAME获取
func (obj *_AllVirtualFilesMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["FILE_NAME"] = fileName })
}

// WithFileType FILE_TYPE获取
func (obj *_AllVirtualFilesMgr) WithFileType(fileType string) Option {
	return optionFunc(func(o *options) { o.query["FILE_TYPE"] = fileType })
}

// WithTablespaceName TABLESPACE_NAME获取
func (obj *_AllVirtualFilesMgr) WithTablespaceName(tablespaceName string) Option {
	return optionFunc(func(o *options) { o.query["TABLESPACE_NAME"] = tablespaceName })
}

// WithTableCatalog TABLE_CATALOG获取
func (obj *_AllVirtualFilesMgr) WithTableCatalog(tableCatalog string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_CATALOG"] = tableCatalog })
}

// WithTableSchema TABLE_SCHEMA获取
func (obj *_AllVirtualFilesMgr) WithTableSchema(tableSchema string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_SCHEMA"] = tableSchema })
}

// WithTableName TABLE_NAME获取
func (obj *_AllVirtualFilesMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithLogfileGroupName LOGFILE_GROUP_NAME获取
func (obj *_AllVirtualFilesMgr) WithLogfileGroupName(logfileGroupName string) Option {
	return optionFunc(func(o *options) { o.query["LOGFILE_GROUP_NAME"] = logfileGroupName })
}

// WithLogfileGroupNumber LOGFILE_GROUP_NUMBER获取
func (obj *_AllVirtualFilesMgr) WithLogfileGroupNumber(logfileGroupNumber int64) Option {
	return optionFunc(func(o *options) { o.query["LOGFILE_GROUP_NUMBER"] = logfileGroupNumber })
}

// WithEngine ENGINE获取
func (obj *_AllVirtualFilesMgr) WithEngine(engine string) Option {
	return optionFunc(func(o *options) { o.query["ENGINE"] = engine })
}

// WithFulltextKeys FULLTEXT_KEYS获取
func (obj *_AllVirtualFilesMgr) WithFulltextKeys(fulltextKeys string) Option {
	return optionFunc(func(o *options) { o.query["FULLTEXT_KEYS"] = fulltextKeys })
}

// WithDeletedRows DELETED_ROWS获取
func (obj *_AllVirtualFilesMgr) WithDeletedRows(deletedRows int64) Option {
	return optionFunc(func(o *options) { o.query["DELETED_ROWS"] = deletedRows })
}

// WithUpdateCount UPDATE_COUNT获取
func (obj *_AllVirtualFilesMgr) WithUpdateCount(updateCount int64) Option {
	return optionFunc(func(o *options) { o.query["UPDATE_COUNT"] = updateCount })
}

// WithFreeExtents FREE_EXTENTS获取
func (obj *_AllVirtualFilesMgr) WithFreeExtents(freeExtents int64) Option {
	return optionFunc(func(o *options) { o.query["FREE_EXTENTS"] = freeExtents })
}

// WithTotalExtents TOTAL_EXTENTS获取
func (obj *_AllVirtualFilesMgr) WithTotalExtents(totalExtents int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_EXTENTS"] = totalExtents })
}

// WithExtentSize EXTENT_SIZE获取
func (obj *_AllVirtualFilesMgr) WithExtentSize(extentSize int64) Option {
	return optionFunc(func(o *options) { o.query["EXTENT_SIZE"] = extentSize })
}

// WithInitialSize INITIAL_SIZE获取
func (obj *_AllVirtualFilesMgr) WithInitialSize(initialSize uint64) Option {
	return optionFunc(func(o *options) { o.query["INITIAL_SIZE"] = initialSize })
}

// WithMaximumSize MAXIMUM_SIZE获取
func (obj *_AllVirtualFilesMgr) WithMaximumSize(maximumSize uint64) Option {
	return optionFunc(func(o *options) { o.query["MAXIMUM_SIZE"] = maximumSize })
}

// WithAutoextendSize AUTOEXTEND_SIZE获取
func (obj *_AllVirtualFilesMgr) WithAutoextendSize(autoextendSize uint64) Option {
	return optionFunc(func(o *options) { o.query["AUTOEXTEND_SIZE"] = autoextendSize })
}

// WithCreationTime CREATION_TIME获取
func (obj *_AllVirtualFilesMgr) WithCreationTime(creationTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CREATION_TIME"] = creationTime })
}

// WithLastUpdateTime LAST_UPDATE_TIME获取
func (obj *_AllVirtualFilesMgr) WithLastUpdateTime(lastUpdateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_UPDATE_TIME"] = lastUpdateTime })
}

// WithLastAccessTime LAST_ACCESS_TIME获取
func (obj *_AllVirtualFilesMgr) WithLastAccessTime(lastAccessTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_ACCESS_TIME"] = lastAccessTime })
}

// WithRecoverTime RECOVER_TIME获取
func (obj *_AllVirtualFilesMgr) WithRecoverTime(recoverTime int64) Option {
	return optionFunc(func(o *options) { o.query["RECOVER_TIME"] = recoverTime })
}

// WithTransactionCounter TRANSACTION_COUNTER获取
func (obj *_AllVirtualFilesMgr) WithTransactionCounter(transactionCounter int64) Option {
	return optionFunc(func(o *options) { o.query["TRANSACTION_COUNTER"] = transactionCounter })
}

// WithVersion VERSION获取
func (obj *_AllVirtualFilesMgr) WithVersion(version uint64) Option {
	return optionFunc(func(o *options) { o.query["VERSION"] = version })
}

// WithRowFormat ROW_FORMAT获取
func (obj *_AllVirtualFilesMgr) WithRowFormat(rowFormat string) Option {
	return optionFunc(func(o *options) { o.query["ROW_FORMAT"] = rowFormat })
}

// WithTableRows TABLE_ROWS获取
func (obj *_AllVirtualFilesMgr) WithTableRows(tableRows uint64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ROWS"] = tableRows })
}

// WithAvgRowLength AVG_ROW_LENGTH获取
func (obj *_AllVirtualFilesMgr) WithAvgRowLength(avgRowLength uint64) Option {
	return optionFunc(func(o *options) { o.query["AVG_ROW_LENGTH"] = avgRowLength })
}

// WithDataLength DATA_LENGTH获取
func (obj *_AllVirtualFilesMgr) WithDataLength(dataLength uint64) Option {
	return optionFunc(func(o *options) { o.query["DATA_LENGTH"] = dataLength })
}

// WithMaxDataLength MAX_DATA_LENGTH获取
func (obj *_AllVirtualFilesMgr) WithMaxDataLength(maxDataLength uint64) Option {
	return optionFunc(func(o *options) { o.query["MAX_DATA_LENGTH"] = maxDataLength })
}

// WithIndexLength INDEX_LENGTH获取
func (obj *_AllVirtualFilesMgr) WithIndexLength(indexLength uint64) Option {
	return optionFunc(func(o *options) { o.query["INDEX_LENGTH"] = indexLength })
}

// WithDataFree DATA_FREE获取
func (obj *_AllVirtualFilesMgr) WithDataFree(dataFree uint64) Option {
	return optionFunc(func(o *options) { o.query["DATA_FREE"] = dataFree })
}

// WithCreateTime CREATE_TIME获取
func (obj *_AllVirtualFilesMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CREATE_TIME"] = createTime })
}

// WithUpdateTime UPDATE_TIME获取
func (obj *_AllVirtualFilesMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["UPDATE_TIME"] = updateTime })
}

// WithCheckTime CHECK_TIME获取
func (obj *_AllVirtualFilesMgr) WithCheckTime(checkTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CHECK_TIME"] = checkTime })
}

// WithChecksum CHECKSUM获取
func (obj *_AllVirtualFilesMgr) WithChecksum(checksum uint64) Option {
	return optionFunc(func(o *options) { o.query["CHECKSUM"] = checksum })
}

// WithStatus STATUS获取
func (obj *_AllVirtualFilesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithExtra EXTRA获取
func (obj *_AllVirtualFilesMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["EXTRA"] = extra })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualFilesMgr) GetByOption(opts ...Option) (result AllVirtualFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualFilesMgr) GetByOptions(opts ...Option) (results []*AllVirtualFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualFilesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualFiles, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where(options.query)
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

// GetFromFileID 通过FILE_ID获取内容
func (obj *_AllVirtualFilesMgr) GetFromFileID(fileID int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FILE_ID` = ?", fileID).Find(&results).Error

	return
}

// GetBatchFromFileID 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromFileID(fileIDs []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FILE_ID` IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromFileName 通过FILE_NAME获取内容
func (obj *_AllVirtualFilesMgr) GetFromFileName(fileName string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FILE_NAME` = ?", fileName).Find(&results).Error

	return
}

// GetBatchFromFileName 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromFileName(fileNames []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FILE_NAME` IN (?)", fileNames).Find(&results).Error

	return
}

// GetFromFileType 通过FILE_TYPE获取内容
func (obj *_AllVirtualFilesMgr) GetFromFileType(fileType string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FILE_TYPE` = ?", fileType).Find(&results).Error

	return
}

// GetBatchFromFileType 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromFileType(fileTypes []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FILE_TYPE` IN (?)", fileTypes).Find(&results).Error

	return
}

// GetFromTablespaceName 通过TABLESPACE_NAME获取内容
func (obj *_AllVirtualFilesMgr) GetFromTablespaceName(tablespaceName string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLESPACE_NAME` = ?", tablespaceName).Find(&results).Error

	return
}

// GetBatchFromTablespaceName 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromTablespaceName(tablespaceNames []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLESPACE_NAME` IN (?)", tablespaceNames).Find(&results).Error

	return
}

// GetFromTableCatalog 通过TABLE_CATALOG获取内容
func (obj *_AllVirtualFilesMgr) GetFromTableCatalog(tableCatalog string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_CATALOG` = ?", tableCatalog).Find(&results).Error

	return
}

// GetBatchFromTableCatalog 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromTableCatalog(tableCatalogs []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_CATALOG` IN (?)", tableCatalogs).Find(&results).Error

	return
}

// GetFromTableSchema 通过TABLE_SCHEMA获取内容
func (obj *_AllVirtualFilesMgr) GetFromTableSchema(tableSchema string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_SCHEMA` = ?", tableSchema).Find(&results).Error

	return
}

// GetBatchFromTableSchema 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromTableSchema(tableSchemas []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_SCHEMA` IN (?)", tableSchemas).Find(&results).Error

	return
}

// GetFromTableName 通过TABLE_NAME获取内容
func (obj *_AllVirtualFilesMgr) GetFromTableName(tableName string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromLogfileGroupName 通过LOGFILE_GROUP_NAME获取内容
func (obj *_AllVirtualFilesMgr) GetFromLogfileGroupName(logfileGroupName string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LOGFILE_GROUP_NAME` = ?", logfileGroupName).Find(&results).Error

	return
}

// GetBatchFromLogfileGroupName 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromLogfileGroupName(logfileGroupNames []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LOGFILE_GROUP_NAME` IN (?)", logfileGroupNames).Find(&results).Error

	return
}

// GetFromLogfileGroupNumber 通过LOGFILE_GROUP_NUMBER获取内容
func (obj *_AllVirtualFilesMgr) GetFromLogfileGroupNumber(logfileGroupNumber int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LOGFILE_GROUP_NUMBER` = ?", logfileGroupNumber).Find(&results).Error

	return
}

// GetBatchFromLogfileGroupNumber 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromLogfileGroupNumber(logfileGroupNumbers []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LOGFILE_GROUP_NUMBER` IN (?)", logfileGroupNumbers).Find(&results).Error

	return
}

// GetFromEngine 通过ENGINE获取内容
func (obj *_AllVirtualFilesMgr) GetFromEngine(engine string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`ENGINE` = ?", engine).Find(&results).Error

	return
}

// GetBatchFromEngine 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromEngine(engines []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`ENGINE` IN (?)", engines).Find(&results).Error

	return
}

// GetFromFulltextKeys 通过FULLTEXT_KEYS获取内容
func (obj *_AllVirtualFilesMgr) GetFromFulltextKeys(fulltextKeys string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FULLTEXT_KEYS` = ?", fulltextKeys).Find(&results).Error

	return
}

// GetBatchFromFulltextKeys 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromFulltextKeys(fulltextKeyss []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FULLTEXT_KEYS` IN (?)", fulltextKeyss).Find(&results).Error

	return
}

// GetFromDeletedRows 通过DELETED_ROWS获取内容
func (obj *_AllVirtualFilesMgr) GetFromDeletedRows(deletedRows int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`DELETED_ROWS` = ?", deletedRows).Find(&results).Error

	return
}

// GetBatchFromDeletedRows 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromDeletedRows(deletedRowss []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`DELETED_ROWS` IN (?)", deletedRowss).Find(&results).Error

	return
}

// GetFromUpdateCount 通过UPDATE_COUNT获取内容
func (obj *_AllVirtualFilesMgr) GetFromUpdateCount(updateCount int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`UPDATE_COUNT` = ?", updateCount).Find(&results).Error

	return
}

// GetBatchFromUpdateCount 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromUpdateCount(updateCounts []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`UPDATE_COUNT` IN (?)", updateCounts).Find(&results).Error

	return
}

// GetFromFreeExtents 通过FREE_EXTENTS获取内容
func (obj *_AllVirtualFilesMgr) GetFromFreeExtents(freeExtents int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FREE_EXTENTS` = ?", freeExtents).Find(&results).Error

	return
}

// GetBatchFromFreeExtents 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromFreeExtents(freeExtentss []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`FREE_EXTENTS` IN (?)", freeExtentss).Find(&results).Error

	return
}

// GetFromTotalExtents 通过TOTAL_EXTENTS获取内容
func (obj *_AllVirtualFilesMgr) GetFromTotalExtents(totalExtents int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TOTAL_EXTENTS` = ?", totalExtents).Find(&results).Error

	return
}

// GetBatchFromTotalExtents 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromTotalExtents(totalExtentss []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TOTAL_EXTENTS` IN (?)", totalExtentss).Find(&results).Error

	return
}

// GetFromExtentSize 通过EXTENT_SIZE获取内容
func (obj *_AllVirtualFilesMgr) GetFromExtentSize(extentSize int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`EXTENT_SIZE` = ?", extentSize).Find(&results).Error

	return
}

// GetBatchFromExtentSize 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromExtentSize(extentSizes []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`EXTENT_SIZE` IN (?)", extentSizes).Find(&results).Error

	return
}

// GetFromInitialSize 通过INITIAL_SIZE获取内容
func (obj *_AllVirtualFilesMgr) GetFromInitialSize(initialSize uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`INITIAL_SIZE` = ?", initialSize).Find(&results).Error

	return
}

// GetBatchFromInitialSize 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromInitialSize(initialSizes []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`INITIAL_SIZE` IN (?)", initialSizes).Find(&results).Error

	return
}

// GetFromMaximumSize 通过MAXIMUM_SIZE获取内容
func (obj *_AllVirtualFilesMgr) GetFromMaximumSize(maximumSize uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`MAXIMUM_SIZE` = ?", maximumSize).Find(&results).Error

	return
}

// GetBatchFromMaximumSize 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromMaximumSize(maximumSizes []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`MAXIMUM_SIZE` IN (?)", maximumSizes).Find(&results).Error

	return
}

// GetFromAutoextendSize 通过AUTOEXTEND_SIZE获取内容
func (obj *_AllVirtualFilesMgr) GetFromAutoextendSize(autoextendSize uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`AUTOEXTEND_SIZE` = ?", autoextendSize).Find(&results).Error

	return
}

// GetBatchFromAutoextendSize 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromAutoextendSize(autoextendSizes []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`AUTOEXTEND_SIZE` IN (?)", autoextendSizes).Find(&results).Error

	return
}

// GetFromCreationTime 通过CREATION_TIME获取内容
func (obj *_AllVirtualFilesMgr) GetFromCreationTime(creationTime time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CREATION_TIME` = ?", creationTime).Find(&results).Error

	return
}

// GetBatchFromCreationTime 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromCreationTime(creationTimes []time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CREATION_TIME` IN (?)", creationTimes).Find(&results).Error

	return
}

// GetFromLastUpdateTime 通过LAST_UPDATE_TIME获取内容
func (obj *_AllVirtualFilesMgr) GetFromLastUpdateTime(lastUpdateTime time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LAST_UPDATE_TIME` = ?", lastUpdateTime).Find(&results).Error

	return
}

// GetBatchFromLastUpdateTime 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromLastUpdateTime(lastUpdateTimes []time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LAST_UPDATE_TIME` IN (?)", lastUpdateTimes).Find(&results).Error

	return
}

// GetFromLastAccessTime 通过LAST_ACCESS_TIME获取内容
func (obj *_AllVirtualFilesMgr) GetFromLastAccessTime(lastAccessTime time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LAST_ACCESS_TIME` = ?", lastAccessTime).Find(&results).Error

	return
}

// GetBatchFromLastAccessTime 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromLastAccessTime(lastAccessTimes []time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`LAST_ACCESS_TIME` IN (?)", lastAccessTimes).Find(&results).Error

	return
}

// GetFromRecoverTime 通过RECOVER_TIME获取内容
func (obj *_AllVirtualFilesMgr) GetFromRecoverTime(recoverTime int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`RECOVER_TIME` = ?", recoverTime).Find(&results).Error

	return
}

// GetBatchFromRecoverTime 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromRecoverTime(recoverTimes []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`RECOVER_TIME` IN (?)", recoverTimes).Find(&results).Error

	return
}

// GetFromTransactionCounter 通过TRANSACTION_COUNTER获取内容
func (obj *_AllVirtualFilesMgr) GetFromTransactionCounter(transactionCounter int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TRANSACTION_COUNTER` = ?", transactionCounter).Find(&results).Error

	return
}

// GetBatchFromTransactionCounter 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromTransactionCounter(transactionCounters []int64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TRANSACTION_COUNTER` IN (?)", transactionCounters).Find(&results).Error

	return
}

// GetFromVersion 通过VERSION获取内容
func (obj *_AllVirtualFilesMgr) GetFromVersion(version uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`VERSION` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromVersion(versions []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`VERSION` IN (?)", versions).Find(&results).Error

	return
}

// GetFromRowFormat 通过ROW_FORMAT获取内容
func (obj *_AllVirtualFilesMgr) GetFromRowFormat(rowFormat string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`ROW_FORMAT` = ?", rowFormat).Find(&results).Error

	return
}

// GetBatchFromRowFormat 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromRowFormat(rowFormats []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`ROW_FORMAT` IN (?)", rowFormats).Find(&results).Error

	return
}

// GetFromTableRows 通过TABLE_ROWS获取内容
func (obj *_AllVirtualFilesMgr) GetFromTableRows(tableRows uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_ROWS` = ?", tableRows).Find(&results).Error

	return
}

// GetBatchFromTableRows 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromTableRows(tableRowss []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`TABLE_ROWS` IN (?)", tableRowss).Find(&results).Error

	return
}

// GetFromAvgRowLength 通过AVG_ROW_LENGTH获取内容
func (obj *_AllVirtualFilesMgr) GetFromAvgRowLength(avgRowLength uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`AVG_ROW_LENGTH` = ?", avgRowLength).Find(&results).Error

	return
}

// GetBatchFromAvgRowLength 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromAvgRowLength(avgRowLengths []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`AVG_ROW_LENGTH` IN (?)", avgRowLengths).Find(&results).Error

	return
}

// GetFromDataLength 通过DATA_LENGTH获取内容
func (obj *_AllVirtualFilesMgr) GetFromDataLength(dataLength uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`DATA_LENGTH` = ?", dataLength).Find(&results).Error

	return
}

// GetBatchFromDataLength 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromDataLength(dataLengths []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`DATA_LENGTH` IN (?)", dataLengths).Find(&results).Error

	return
}

// GetFromMaxDataLength 通过MAX_DATA_LENGTH获取内容
func (obj *_AllVirtualFilesMgr) GetFromMaxDataLength(maxDataLength uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`MAX_DATA_LENGTH` = ?", maxDataLength).Find(&results).Error

	return
}

// GetBatchFromMaxDataLength 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromMaxDataLength(maxDataLengths []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`MAX_DATA_LENGTH` IN (?)", maxDataLengths).Find(&results).Error

	return
}

// GetFromIndexLength 通过INDEX_LENGTH获取内容
func (obj *_AllVirtualFilesMgr) GetFromIndexLength(indexLength uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`INDEX_LENGTH` = ?", indexLength).Find(&results).Error

	return
}

// GetBatchFromIndexLength 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromIndexLength(indexLengths []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`INDEX_LENGTH` IN (?)", indexLengths).Find(&results).Error

	return
}

// GetFromDataFree 通过DATA_FREE获取内容
func (obj *_AllVirtualFilesMgr) GetFromDataFree(dataFree uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`DATA_FREE` = ?", dataFree).Find(&results).Error

	return
}

// GetBatchFromDataFree 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromDataFree(dataFrees []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`DATA_FREE` IN (?)", dataFrees).Find(&results).Error

	return
}

// GetFromCreateTime 通过CREATE_TIME获取内容
func (obj *_AllVirtualFilesMgr) GetFromCreateTime(createTime time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CREATE_TIME` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CREATE_TIME` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过UPDATE_TIME获取内容
func (obj *_AllVirtualFilesMgr) GetFromUpdateTime(updateTime time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`UPDATE_TIME` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`UPDATE_TIME` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCheckTime 通过CHECK_TIME获取内容
func (obj *_AllVirtualFilesMgr) GetFromCheckTime(checkTime time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CHECK_TIME` = ?", checkTime).Find(&results).Error

	return
}

// GetBatchFromCheckTime 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromCheckTime(checkTimes []time.Time) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CHECK_TIME` IN (?)", checkTimes).Find(&results).Error

	return
}

// GetFromChecksum 通过CHECKSUM获取内容
func (obj *_AllVirtualFilesMgr) GetFromChecksum(checksum uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CHECKSUM` = ?", checksum).Find(&results).Error

	return
}

// GetBatchFromChecksum 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromChecksum(checksums []uint64) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`CHECKSUM` IN (?)", checksums).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_AllVirtualFilesMgr) GetFromStatus(status string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromExtra 通过EXTRA获取内容
func (obj *_AllVirtualFilesMgr) GetFromExtra(extra string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`EXTRA` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找
func (obj *_AllVirtualFilesMgr) GetBatchFromExtra(extras []string) (results []*AllVirtualFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFiles{}).Where("`EXTRA` IN (?)", extras).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
