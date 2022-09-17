package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualDataTypeMgr struct {
	*_BaseMgr
}

// AllVirtualDataTypeMgr open func
func AllVirtualDataTypeMgr(db *gorm.DB) *_AllVirtualDataTypeMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDataTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDataTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_data_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDataTypeMgr) GetTableName() string {
	return "__all_virtual_data_type"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDataTypeMgr) Reset() *_AllVirtualDataTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDataTypeMgr) Get() (result AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDataTypeMgr) Gets() (results []*AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDataTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDataType data_type获取
func (obj *_AllVirtualDataTypeMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithDataTypeStr data_type_str获取
func (obj *_AllVirtualDataTypeMgr) WithDataTypeStr(dataTypeStr string) Option {
	return optionFunc(func(o *options) { o.query["data_type_str"] = dataTypeStr })
}

// WithDataTypeClass data_type_class获取
func (obj *_AllVirtualDataTypeMgr) WithDataTypeClass(dataTypeClass int64) Option {
	return optionFunc(func(o *options) { o.query["data_type_class"] = dataTypeClass })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDataTypeMgr) GetByOption(opts ...Option) (result AllVirtualDataType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDataTypeMgr) GetByOptions(opts ...Option) (results []*AllVirtualDataType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDataTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDataType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where(options.query)
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

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualDataTypeMgr) GetFromDataType(dataType int64) (result AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where("`data_type` = ?", dataType).First(&result).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualDataTypeMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromDataTypeStr 通过data_type_str获取内容
func (obj *_AllVirtualDataTypeMgr) GetFromDataTypeStr(dataTypeStr string) (results []*AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where("`data_type_str` = ?", dataTypeStr).Find(&results).Error

	return
}

// GetBatchFromDataTypeStr 批量查找
func (obj *_AllVirtualDataTypeMgr) GetBatchFromDataTypeStr(dataTypeStrs []string) (results []*AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where("`data_type_str` IN (?)", dataTypeStrs).Find(&results).Error

	return
}

// GetFromDataTypeClass 通过data_type_class获取内容
func (obj *_AllVirtualDataTypeMgr) GetFromDataTypeClass(dataTypeClass int64) (results []*AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where("`data_type_class` = ?", dataTypeClass).Find(&results).Error

	return
}

// GetBatchFromDataTypeClass 批量查找
func (obj *_AllVirtualDataTypeMgr) GetBatchFromDataTypeClass(dataTypeClasss []int64) (results []*AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where("`data_type_class` IN (?)", dataTypeClasss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDataTypeMgr) FetchByPrimaryKey(dataType int64) (result AllVirtualDataType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataType{}).Where("`data_type` = ?", dataType).First(&result).Error

	return
}
