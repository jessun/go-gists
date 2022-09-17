package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllZoneMgr struct {
	*_BaseMgr
}

// AllZoneMgr open func
func AllZoneMgr(db *gorm.DB) *_AllZoneMgr {
	if db == nil {
		panic(fmt.Errorf("AllZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllZoneMgr) GetTableName() string {
	return "__all_zone"
}

// Reset 重置gorm会话
func (obj *_AllZoneMgr) Reset() *_AllZoneMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllZoneMgr) Get() (result AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllZoneMgr) Gets() (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllZoneMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllZone{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllZoneMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllZoneMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithZone zone获取
func (obj *_AllZoneMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithName name获取
func (obj *_AllZoneMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取
func (obj *_AllZoneMgr) WithValue(value int64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllZoneMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// GetByOption 功能选项模式获取
func (obj *_AllZoneMgr) GetByOption(opts ...Option) (result AllZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllZoneMgr) GetByOptions(opts ...Option) (results []*AllZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllZoneMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllZone, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where(options.query)
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
func (obj *_AllZoneMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllZoneMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllZoneMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllZoneMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllZoneMgr) GetFromZone(zone string) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllZoneMgr) GetBatchFromZone(zones []string) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllZoneMgr) GetFromName(name string) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllZoneMgr) GetBatchFromName(names []string) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllZoneMgr) GetFromValue(value int64) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllZoneMgr) GetBatchFromValue(values []int64) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllZoneMgr) GetFromInfo(info string) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllZoneMgr) GetBatchFromInfo(infos []string) (results []*AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllZoneMgr) FetchByPrimaryKey(zone string, name string) (result AllZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllZone{}).Where("`zone` = ? AND `name` = ?", zone, name).First(&result).Error

	return
}
