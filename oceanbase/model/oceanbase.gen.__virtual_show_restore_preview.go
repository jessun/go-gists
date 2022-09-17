package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _VirtualShowRestorePreviewMgr struct {
	*_BaseMgr
}

// VirtualShowRestorePreviewMgr open func
func VirtualShowRestorePreviewMgr(db *gorm.DB) *_VirtualShowRestorePreviewMgr {
	if db == nil {
		panic(fmt.Errorf("VirtualShowRestorePreviewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VirtualShowRestorePreviewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__virtual_show_restore_preview"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VirtualShowRestorePreviewMgr) GetTableName() string {
	return "__virtual_show_restore_preview"
}

// Reset 重置gorm会话
func (obj *_VirtualShowRestorePreviewMgr) Reset() *_VirtualShowRestorePreviewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VirtualShowRestorePreviewMgr) Get() (result VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_VirtualShowRestorePreviewMgr) Gets() (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VirtualShowRestorePreviewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBackupType backup_type获取
func (obj *_VirtualShowRestorePreviewMgr) WithBackupType(backupType string) Option {
	return optionFunc(func(o *options) { o.query["backup_type"] = backupType })
}

// WithBackupID backup_id获取
func (obj *_VirtualShowRestorePreviewMgr) WithBackupID(backupID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_id"] = backupID })
}

// WithCopyID copy_id获取
func (obj *_VirtualShowRestorePreviewMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithBackupDest backup_dest获取
func (obj *_VirtualShowRestorePreviewMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithFileStatus file_status获取
func (obj *_VirtualShowRestorePreviewMgr) WithFileStatus(fileStatus string) Option {
	return optionFunc(func(o *options) { o.query["file_status"] = fileStatus })
}

// GetByOption 功能选项模式获取
func (obj *_VirtualShowRestorePreviewMgr) GetByOption(opts ...Option) (result VirtualShowRestorePreview, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VirtualShowRestorePreviewMgr) GetByOptions(opts ...Option) (results []*VirtualShowRestorePreview, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_VirtualShowRestorePreviewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]VirtualShowRestorePreview, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where(options.query)
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

// GetFromBackupType 通过backup_type获取内容
func (obj *_VirtualShowRestorePreviewMgr) GetFromBackupType(backupType string) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`backup_type` = ?", backupType).Find(&results).Error

	return
}

// GetBatchFromBackupType 批量查找
func (obj *_VirtualShowRestorePreviewMgr) GetBatchFromBackupType(backupTypes []string) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`backup_type` IN (?)", backupTypes).Find(&results).Error

	return
}

// GetFromBackupID 通过backup_id获取内容
func (obj *_VirtualShowRestorePreviewMgr) GetFromBackupID(backupID int64) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`backup_id` = ?", backupID).Find(&results).Error

	return
}

// GetBatchFromBackupID 批量查找
func (obj *_VirtualShowRestorePreviewMgr) GetBatchFromBackupID(backupIDs []int64) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`backup_id` IN (?)", backupIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_VirtualShowRestorePreviewMgr) GetFromCopyID(copyID int64) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_VirtualShowRestorePreviewMgr) GetBatchFromCopyID(copyIDs []int64) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_VirtualShowRestorePreviewMgr) GetFromBackupDest(backupDest string) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_VirtualShowRestorePreviewMgr) GetBatchFromBackupDest(backupDests []string) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromFileStatus 通过file_status获取内容
func (obj *_VirtualShowRestorePreviewMgr) GetFromFileStatus(fileStatus string) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`file_status` = ?", fileStatus).Find(&results).Error

	return
}

// GetBatchFromFileStatus 批量查找
func (obj *_VirtualShowRestorePreviewMgr) GetBatchFromFileStatus(fileStatuss []string) (results []*VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`file_status` IN (?)", fileStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_VirtualShowRestorePreviewMgr) FetchByPrimaryKey(backupType string, backupID int64, copyID int64) (result VirtualShowRestorePreview, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VirtualShowRestorePreview{}).Where("`backup_type` = ? AND `backup_id` = ? AND `copy_id` = ?", backupType, backupID, copyID).First(&result).Error

	return
}
