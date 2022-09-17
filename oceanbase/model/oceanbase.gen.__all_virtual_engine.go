package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualEngineMgr struct {
	*_BaseMgr
}

// AllVirtualEngineMgr open func
func AllVirtualEngineMgr(db *gorm.DB) *_AllVirtualEngineMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualEngineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualEngineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_engine"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualEngineMgr) GetTableName() string {
	return "__all_virtual_engine"
}

// Reset 重置gorm会话
func (obj *_AllVirtualEngineMgr) Reset() *_AllVirtualEngineMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualEngineMgr) Get() (result AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualEngineMgr) Gets() (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualEngineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithEngine Engine获取
func (obj *_AllVirtualEngineMgr) WithEngine(engine string) Option {
	return optionFunc(func(o *options) { o.query["Engine"] = engine })
}

// WithSupport Support获取
func (obj *_AllVirtualEngineMgr) WithSupport(support string) Option {
	return optionFunc(func(o *options) { o.query["Support"] = support })
}

// WithComment Comment获取
func (obj *_AllVirtualEngineMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["Comment"] = comment })
}

// WithTransactions Transactions获取
func (obj *_AllVirtualEngineMgr) WithTransactions(transactions string) Option {
	return optionFunc(func(o *options) { o.query["Transactions"] = transactions })
}

// WithXa XA获取
func (obj *_AllVirtualEngineMgr) WithXa(xa string) Option {
	return optionFunc(func(o *options) { o.query["XA"] = xa })
}

// WithSavepoints Savepoints获取
func (obj *_AllVirtualEngineMgr) WithSavepoints(savepoints string) Option {
	return optionFunc(func(o *options) { o.query["Savepoints"] = savepoints })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualEngineMgr) GetByOption(opts ...Option) (result AllVirtualEngine, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualEngineMgr) GetByOptions(opts ...Option) (results []*AllVirtualEngine, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualEngineMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualEngine, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where(options.query)
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

// GetFromEngine 通过Engine获取内容
func (obj *_AllVirtualEngineMgr) GetFromEngine(engine string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Engine` = ?", engine).Find(&results).Error

	return
}

// GetBatchFromEngine 批量查找
func (obj *_AllVirtualEngineMgr) GetBatchFromEngine(engines []string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Engine` IN (?)", engines).Find(&results).Error

	return
}

// GetFromSupport 通过Support获取内容
func (obj *_AllVirtualEngineMgr) GetFromSupport(support string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Support` = ?", support).Find(&results).Error

	return
}

// GetBatchFromSupport 批量查找
func (obj *_AllVirtualEngineMgr) GetBatchFromSupport(supports []string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Support` IN (?)", supports).Find(&results).Error

	return
}

// GetFromComment 通过Comment获取内容
func (obj *_AllVirtualEngineMgr) GetFromComment(comment string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualEngineMgr) GetBatchFromComment(comments []string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromTransactions 通过Transactions获取内容
func (obj *_AllVirtualEngineMgr) GetFromTransactions(transactions string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Transactions` = ?", transactions).Find(&results).Error

	return
}

// GetBatchFromTransactions 批量查找
func (obj *_AllVirtualEngineMgr) GetBatchFromTransactions(transactionss []string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Transactions` IN (?)", transactionss).Find(&results).Error

	return
}

// GetFromXa 通过XA获取内容
func (obj *_AllVirtualEngineMgr) GetFromXa(xa string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`XA` = ?", xa).Find(&results).Error

	return
}

// GetBatchFromXa 批量查找
func (obj *_AllVirtualEngineMgr) GetBatchFromXa(xas []string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`XA` IN (?)", xas).Find(&results).Error

	return
}

// GetFromSavepoints 通过Savepoints获取内容
func (obj *_AllVirtualEngineMgr) GetFromSavepoints(savepoints string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Savepoints` = ?", savepoints).Find(&results).Error

	return
}

// GetBatchFromSavepoints 批量查找
func (obj *_AllVirtualEngineMgr) GetBatchFromSavepoints(savepointss []string) (results []*AllVirtualEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualEngine{}).Where("`Savepoints` IN (?)", savepointss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
