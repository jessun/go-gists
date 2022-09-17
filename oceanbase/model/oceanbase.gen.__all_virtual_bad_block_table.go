package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualBadBlockTableMgr struct {
	*_BaseMgr
}

// AllVirtualBadBlockTableMgr open func
func AllVirtualBadBlockTableMgr(db *gorm.DB) *_AllVirtualBadBlockTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualBadBlockTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualBadBlockTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_bad_block_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualBadBlockTableMgr) GetTableName() string {
	return "__all_virtual_bad_block_table"
}

// Reset 重置gorm会话
func (obj *_AllVirtualBadBlockTableMgr) Reset() *_AllVirtualBadBlockTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualBadBlockTableMgr) Get() (result AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualBadBlockTableMgr) Gets() (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualBadBlockTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualBadBlockTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualBadBlockTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithDiskID disk_id获取
func (obj *_AllVirtualBadBlockTableMgr) WithDiskID(diskID int64) Option {
	return optionFunc(func(o *options) { o.query["disk_id"] = diskID })
}

// WithStoreFilePath store_file_path获取
func (obj *_AllVirtualBadBlockTableMgr) WithStoreFilePath(storeFilePath string) Option {
	return optionFunc(func(o *options) { o.query["store_file_path"] = storeFilePath })
}

// WithMacroBlockIndex macro_block_index获取
func (obj *_AllVirtualBadBlockTableMgr) WithMacroBlockIndex(macroBlockIndex int64) Option {
	return optionFunc(func(o *options) { o.query["macro_block_index"] = macroBlockIndex })
}

// WithErrorType error_type获取
func (obj *_AllVirtualBadBlockTableMgr) WithErrorType(errorType int64) Option {
	return optionFunc(func(o *options) { o.query["error_type"] = errorType })
}

// WithErrorMsg error_msg获取
func (obj *_AllVirtualBadBlockTableMgr) WithErrorMsg(errorMsg string) Option {
	return optionFunc(func(o *options) { o.query["error_msg"] = errorMsg })
}

// WithCheckTime check_time获取
func (obj *_AllVirtualBadBlockTableMgr) WithCheckTime(checkTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["check_time"] = checkTime })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualBadBlockTableMgr) GetByOption(opts ...Option) (result AllVirtualBadBlockTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualBadBlockTableMgr) GetByOptions(opts ...Option) (results []*AllVirtualBadBlockTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualBadBlockTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualBadBlockTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromDiskID 通过disk_id获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromDiskID(diskID int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`disk_id` = ?", diskID).Find(&results).Error

	return
}

// GetBatchFromDiskID 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromDiskID(diskIDs []int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`disk_id` IN (?)", diskIDs).Find(&results).Error

	return
}

// GetFromStoreFilePath 通过store_file_path获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromStoreFilePath(storeFilePath string) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`store_file_path` = ?", storeFilePath).Find(&results).Error

	return
}

// GetBatchFromStoreFilePath 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromStoreFilePath(storeFilePaths []string) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`store_file_path` IN (?)", storeFilePaths).Find(&results).Error

	return
}

// GetFromMacroBlockIndex 通过macro_block_index获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromMacroBlockIndex(macroBlockIndex int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`macro_block_index` = ?", macroBlockIndex).Find(&results).Error

	return
}

// GetBatchFromMacroBlockIndex 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromMacroBlockIndex(macroBlockIndexs []int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`macro_block_index` IN (?)", macroBlockIndexs).Find(&results).Error

	return
}

// GetFromErrorType 通过error_type获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromErrorType(errorType int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`error_type` = ?", errorType).Find(&results).Error

	return
}

// GetBatchFromErrorType 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromErrorType(errorTypes []int64) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`error_type` IN (?)", errorTypes).Find(&results).Error

	return
}

// GetFromErrorMsg 通过error_msg获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromErrorMsg(errorMsg string) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`error_msg` = ?", errorMsg).Find(&results).Error

	return
}

// GetBatchFromErrorMsg 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromErrorMsg(errorMsgs []string) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`error_msg` IN (?)", errorMsgs).Find(&results).Error

	return
}

// GetFromCheckTime 通过check_time获取内容
func (obj *_AllVirtualBadBlockTableMgr) GetFromCheckTime(checkTime time.Time) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`check_time` = ?", checkTime).Find(&results).Error

	return
}

// GetBatchFromCheckTime 批量查找
func (obj *_AllVirtualBadBlockTableMgr) GetBatchFromCheckTime(checkTimes []time.Time) (results []*AllVirtualBadBlockTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualBadBlockTable{}).Where("`check_time` IN (?)", checkTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
