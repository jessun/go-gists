package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllPrivilegeMgr struct {
	*_BaseMgr
}

// AllPrivilegeMgr open func
func AllPrivilegeMgr(db *gorm.DB) *_AllPrivilegeMgr {
	if db == nil {
		panic(fmt.Errorf("AllPrivilegeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllPrivilegeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_privilege"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllPrivilegeMgr) GetTableName() string {
	return "__all_privilege"
}

// Reset 重置gorm会话
func (obj *_AllPrivilegeMgr) Reset() *_AllPrivilegeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllPrivilegeMgr) Get() (result AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllPrivilegeMgr) Gets() (results []*AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllPrivilegeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPrivilege Privilege获取
func (obj *_AllPrivilegeMgr) WithPrivilege(privilege string) Option {
	return optionFunc(func(o *options) { o.query["Privilege"] = privilege })
}

// WithContext Context获取
func (obj *_AllPrivilegeMgr) WithContext(context string) Option {
	return optionFunc(func(o *options) { o.query["Context"] = context })
}

// WithComment Comment获取
func (obj *_AllPrivilegeMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["Comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllPrivilegeMgr) GetByOption(opts ...Option) (result AllPrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllPrivilegeMgr) GetByOptions(opts ...Option) (results []*AllPrivilege, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllPrivilegeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllPrivilege, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where(options.query)
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

// GetFromPrivilege 通过Privilege获取内容
func (obj *_AllPrivilegeMgr) GetFromPrivilege(privilege string) (result AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where("`Privilege` = ?", privilege).First(&result).Error

	return
}

// GetBatchFromPrivilege 批量查找
func (obj *_AllPrivilegeMgr) GetBatchFromPrivilege(privileges []string) (results []*AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where("`Privilege` IN (?)", privileges).Find(&results).Error

	return
}

// GetFromContext 通过Context获取内容
func (obj *_AllPrivilegeMgr) GetFromContext(context string) (results []*AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where("`Context` = ?", context).Find(&results).Error

	return
}

// GetBatchFromContext 批量查找
func (obj *_AllPrivilegeMgr) GetBatchFromContext(contexts []string) (results []*AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where("`Context` IN (?)", contexts).Find(&results).Error

	return
}

// GetFromComment 通过Comment获取内容
func (obj *_AllPrivilegeMgr) GetFromComment(comment string) (results []*AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where("`Comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllPrivilegeMgr) GetBatchFromComment(comments []string) (results []*AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where("`Comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllPrivilegeMgr) FetchByPrimaryKey(privilege string) (result AllPrivilege, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPrivilege{}).Where("`Privilege` = ?", privilege).First(&result).Error

	return
}
