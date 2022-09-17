package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualDataTypeClassMgr struct {
	*_BaseMgr
}

// AllVirtualDataTypeClassMgr open func
func AllVirtualDataTypeClassMgr(db *gorm.DB) *_AllVirtualDataTypeClassMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDataTypeClassMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDataTypeClassMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_data_type_class"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDataTypeClassMgr) GetTableName() string {
	return "__all_virtual_data_type_class"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDataTypeClassMgr) Reset() *_AllVirtualDataTypeClassMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDataTypeClassMgr) Get() (result AllVirtualDataTypeClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDataTypeClassMgr) Gets() (results []*AllVirtualDataTypeClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDataTypeClassMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDataTypeClass data_type_class获取
func (obj *_AllVirtualDataTypeClassMgr) WithDataTypeClass(dataTypeClass int64) Option {
	return optionFunc(func(o *options) { o.query["data_type_class"] = dataTypeClass })
}

// WithDataTypeClassStr data_type_class_str获取
func (obj *_AllVirtualDataTypeClassMgr) WithDataTypeClassStr(dataTypeClassStr string) Option {
	return optionFunc(func(o *options) { o.query["data_type_class_str"] = dataTypeClassStr })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDataTypeClassMgr) GetByOption(opts ...Option) (result AllVirtualDataTypeClass, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDataTypeClassMgr) GetByOptions(opts ...Option) (results []*AllVirtualDataTypeClass, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDataTypeClassMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDataTypeClass, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where(options.query)
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

// GetFromDataTypeClass 通过data_type_class获取内容
func (obj *_AllVirtualDataTypeClassMgr) GetFromDataTypeClass(dataTypeClass int64) (result AllVirtualDataTypeClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where("`data_type_class` = ?", dataTypeClass).First(&result).Error

	return
}

// GetBatchFromDataTypeClass 批量查找
func (obj *_AllVirtualDataTypeClassMgr) GetBatchFromDataTypeClass(dataTypeClasss []int64) (results []*AllVirtualDataTypeClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where("`data_type_class` IN (?)", dataTypeClasss).Find(&results).Error

	return
}

// GetFromDataTypeClassStr 通过data_type_class_str获取内容
func (obj *_AllVirtualDataTypeClassMgr) GetFromDataTypeClassStr(dataTypeClassStr string) (results []*AllVirtualDataTypeClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where("`data_type_class_str` = ?", dataTypeClassStr).Find(&results).Error

	return
}

// GetBatchFromDataTypeClassStr 批量查找
func (obj *_AllVirtualDataTypeClassMgr) GetBatchFromDataTypeClassStr(dataTypeClassStrs []string) (results []*AllVirtualDataTypeClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where("`data_type_class_str` IN (?)", dataTypeClassStrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDataTypeClassMgr) FetchByPrimaryKey(dataTypeClass int64) (result AllVirtualDataTypeClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDataTypeClass{}).Where("`data_type_class` = ?", dataTypeClass).First(&result).Error

	return
}
