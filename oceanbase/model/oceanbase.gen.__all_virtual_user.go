package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualUserMgr struct {
	*_BaseMgr
}

// AllVirtualUserMgr open func
func AllVirtualUserMgr(db *gorm.DB) *_AllVirtualUserMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualUserMgr) GetTableName() string {
	return "__all_virtual_user"
}

// Reset 重置gorm会话
func (obj *_AllVirtualUserMgr) Reset() *_AllVirtualUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualUserMgr) Get() (result AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualUserMgr) Gets() (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualUserMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllVirtualUserMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualUserMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualUserMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithUserName user_name获取
func (obj *_AllVirtualUserMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithHost host获取
func (obj *_AllVirtualUserMgr) WithHost(host string) Option {
	return optionFunc(func(o *options) { o.query["host"] = host })
}

// WithPasswd passwd获取
func (obj *_AllVirtualUserMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithInfo info获取
func (obj *_AllVirtualUserMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithPrivAlter priv_alter获取
func (obj *_AllVirtualUserMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllVirtualUserMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllVirtualUserMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllVirtualUserMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllVirtualUserMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllVirtualUserMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllVirtualUserMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllVirtualUserMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllVirtualUserMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllVirtualUserMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllVirtualUserMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// WithPrivShowDb priv_show_db获取
func (obj *_AllVirtualUserMgr) WithPrivShowDb(privShowDb int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_db"] = privShowDb })
}

// WithPrivCreateUser priv_create_user获取
func (obj *_AllVirtualUserMgr) WithPrivCreateUser(privCreateUser int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_user"] = privCreateUser })
}

// WithPrivSuper priv_super获取
func (obj *_AllVirtualUserMgr) WithPrivSuper(privSuper int64) Option {
	return optionFunc(func(o *options) { o.query["priv_super"] = privSuper })
}

// WithIsLocked is_locked获取
func (obj *_AllVirtualUserMgr) WithIsLocked(isLocked int64) Option {
	return optionFunc(func(o *options) { o.query["is_locked"] = isLocked })
}

// WithPrivProcess priv_process获取
func (obj *_AllVirtualUserMgr) WithPrivProcess(privProcess int64) Option {
	return optionFunc(func(o *options) { o.query["priv_process"] = privProcess })
}

// WithPrivCreateSynonym priv_create_synonym获取
func (obj *_AllVirtualUserMgr) WithPrivCreateSynonym(privCreateSynonym int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_synonym"] = privCreateSynonym })
}

// WithSslType ssl_type获取
func (obj *_AllVirtualUserMgr) WithSslType(sslType int64) Option {
	return optionFunc(func(o *options) { o.query["ssl_type"] = sslType })
}

// WithSslCipher ssl_cipher获取
func (obj *_AllVirtualUserMgr) WithSslCipher(sslCipher string) Option {
	return optionFunc(func(o *options) { o.query["ssl_cipher"] = sslCipher })
}

// WithX509Issuer x509_issuer获取
func (obj *_AllVirtualUserMgr) WithX509Issuer(x509Issuer string) Option {
	return optionFunc(func(o *options) { o.query["x509_issuer"] = x509Issuer })
}

// WithX509Subject x509_subject获取
func (obj *_AllVirtualUserMgr) WithX509Subject(x509Subject string) Option {
	return optionFunc(func(o *options) { o.query["x509_subject"] = x509Subject })
}

// WithType type获取
func (obj *_AllVirtualUserMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithProfileID profile_id获取
func (obj *_AllVirtualUserMgr) WithProfileID(profileID int64) Option {
	return optionFunc(func(o *options) { o.query["profile_id"] = profileID })
}

// WithPasswordLastChanged password_last_changed获取
func (obj *_AllVirtualUserMgr) WithPasswordLastChanged(passwordLastChanged time.Time) Option {
	return optionFunc(func(o *options) { o.query["password_last_changed"] = passwordLastChanged })
}

// WithPrivFile priv_file获取
func (obj *_AllVirtualUserMgr) WithPrivFile(privFile int64) Option {
	return optionFunc(func(o *options) { o.query["priv_file"] = privFile })
}

// WithPrivAlterTenant priv_alter_tenant获取
func (obj *_AllVirtualUserMgr) WithPrivAlterTenant(privAlterTenant int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_tenant"] = privAlterTenant })
}

// WithPrivAlterSystem priv_alter_system获取
func (obj *_AllVirtualUserMgr) WithPrivAlterSystem(privAlterSystem int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_system"] = privAlterSystem })
}

// WithPrivCreateResourcePool priv_create_resource_pool获取
func (obj *_AllVirtualUserMgr) WithPrivCreateResourcePool(privCreateResourcePool int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_pool"] = privCreateResourcePool })
}

// WithPrivCreateResourceUnit priv_create_resource_unit获取
func (obj *_AllVirtualUserMgr) WithPrivCreateResourceUnit(privCreateResourceUnit int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_unit"] = privCreateResourceUnit })
}

// WithMaxConnections max_connections获取
func (obj *_AllVirtualUserMgr) WithMaxConnections(maxConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_connections"] = maxConnections })
}

// WithMaxUserConnections max_user_connections获取
func (obj *_AllVirtualUserMgr) WithMaxUserConnections(maxUserConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_user_connections"] = maxUserConnections })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualUserMgr) GetByOption(opts ...Option) (result AllVirtualUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualUserMgr) GetByOptions(opts ...Option) (results []*AllVirtualUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualUserMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualUserMgr) GetFromUserID(userID int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualUserMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualUserMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllVirtualUserMgr) GetFromUserName(userName string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromHost 通过host获取内容
func (obj *_AllVirtualUserMgr) GetFromHost(host string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`host` = ?", host).Find(&results).Error

	return
}

// GetBatchFromHost 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromHost(hosts []string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`host` IN (?)", hosts).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllVirtualUserMgr) GetFromPasswd(passwd string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPasswd(passwds []string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualUserMgr) GetFromInfo(info string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivAlter(privAlter int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivCreate(privCreate int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivDelete(privDelete int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivDrop(privDrop int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivInsert(privInsert int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivSelect(privSelect int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivIndex(privIndex int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivShowView(privShowView int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

// GetFromPrivShowDb 通过priv_show_db获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivShowDb(privShowDb int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_show_db` = ?", privShowDb).Find(&results).Error

	return
}

// GetBatchFromPrivShowDb 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivShowDb(privShowDbs []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_show_db` IN (?)", privShowDbs).Find(&results).Error

	return
}

// GetFromPrivCreateUser 通过priv_create_user获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivCreateUser(privCreateUser int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_user` = ?", privCreateUser).Find(&results).Error

	return
}

// GetBatchFromPrivCreateUser 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivCreateUser(privCreateUsers []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_user` IN (?)", privCreateUsers).Find(&results).Error

	return
}

// GetFromPrivSuper 通过priv_super获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivSuper(privSuper int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_super` = ?", privSuper).Find(&results).Error

	return
}

// GetBatchFromPrivSuper 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivSuper(privSupers []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_super` IN (?)", privSupers).Find(&results).Error

	return
}

// GetFromIsLocked 通过is_locked获取内容
func (obj *_AllVirtualUserMgr) GetFromIsLocked(isLocked int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`is_locked` = ?", isLocked).Find(&results).Error

	return
}

// GetBatchFromIsLocked 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromIsLocked(isLockeds []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`is_locked` IN (?)", isLockeds).Find(&results).Error

	return
}

// GetFromPrivProcess 通过priv_process获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivProcess(privProcess int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_process` = ?", privProcess).Find(&results).Error

	return
}

// GetBatchFromPrivProcess 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivProcess(privProcesss []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_process` IN (?)", privProcesss).Find(&results).Error

	return
}

// GetFromPrivCreateSynonym 通过priv_create_synonym获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivCreateSynonym(privCreateSynonym int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_synonym` = ?", privCreateSynonym).Find(&results).Error

	return
}

// GetBatchFromPrivCreateSynonym 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivCreateSynonym(privCreateSynonyms []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_synonym` IN (?)", privCreateSynonyms).Find(&results).Error

	return
}

// GetFromSslType 通过ssl_type获取内容
func (obj *_AllVirtualUserMgr) GetFromSslType(sslType int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`ssl_type` = ?", sslType).Find(&results).Error

	return
}

// GetBatchFromSslType 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromSslType(sslTypes []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`ssl_type` IN (?)", sslTypes).Find(&results).Error

	return
}

// GetFromSslCipher 通过ssl_cipher获取内容
func (obj *_AllVirtualUserMgr) GetFromSslCipher(sslCipher string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`ssl_cipher` = ?", sslCipher).Find(&results).Error

	return
}

// GetBatchFromSslCipher 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromSslCipher(sslCiphers []string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`ssl_cipher` IN (?)", sslCiphers).Find(&results).Error

	return
}

// GetFromX509Issuer 通过x509_issuer获取内容
func (obj *_AllVirtualUserMgr) GetFromX509Issuer(x509Issuer string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`x509_issuer` = ?", x509Issuer).Find(&results).Error

	return
}

// GetBatchFromX509Issuer 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromX509Issuer(x509Issuers []string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`x509_issuer` IN (?)", x509Issuers).Find(&results).Error

	return
}

// GetFromX509Subject 通过x509_subject获取内容
func (obj *_AllVirtualUserMgr) GetFromX509Subject(x509Subject string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`x509_subject` = ?", x509Subject).Find(&results).Error

	return
}

// GetBatchFromX509Subject 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromX509Subject(x509Subjects []string) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`x509_subject` IN (?)", x509Subjects).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualUserMgr) GetFromType(_type int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromType(_types []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromProfileID 通过profile_id获取内容
func (obj *_AllVirtualUserMgr) GetFromProfileID(profileID int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`profile_id` = ?", profileID).Find(&results).Error

	return
}

// GetBatchFromProfileID 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromProfileID(profileIDs []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`profile_id` IN (?)", profileIDs).Find(&results).Error

	return
}

// GetFromPasswordLastChanged 通过password_last_changed获取内容
func (obj *_AllVirtualUserMgr) GetFromPasswordLastChanged(passwordLastChanged time.Time) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`password_last_changed` = ?", passwordLastChanged).Find(&results).Error

	return
}

// GetBatchFromPasswordLastChanged 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPasswordLastChanged(passwordLastChangeds []time.Time) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`password_last_changed` IN (?)", passwordLastChangeds).Find(&results).Error

	return
}

// GetFromPrivFile 通过priv_file获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivFile(privFile int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_file` = ?", privFile).Find(&results).Error

	return
}

// GetBatchFromPrivFile 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivFile(privFiles []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_file` IN (?)", privFiles).Find(&results).Error

	return
}

// GetFromPrivAlterTenant 通过priv_alter_tenant获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivAlterTenant(privAlterTenant int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_alter_tenant` = ?", privAlterTenant).Find(&results).Error

	return
}

// GetBatchFromPrivAlterTenant 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivAlterTenant(privAlterTenants []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_alter_tenant` IN (?)", privAlterTenants).Find(&results).Error

	return
}

// GetFromPrivAlterSystem 通过priv_alter_system获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivAlterSystem(privAlterSystem int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_alter_system` = ?", privAlterSystem).Find(&results).Error

	return
}

// GetBatchFromPrivAlterSystem 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivAlterSystem(privAlterSystems []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_alter_system` IN (?)", privAlterSystems).Find(&results).Error

	return
}

// GetFromPrivCreateResourcePool 通过priv_create_resource_pool获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivCreateResourcePool(privCreateResourcePool int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_resource_pool` = ?", privCreateResourcePool).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourcePool 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivCreateResourcePool(privCreateResourcePools []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_resource_pool` IN (?)", privCreateResourcePools).Find(&results).Error

	return
}

// GetFromPrivCreateResourceUnit 通过priv_create_resource_unit获取内容
func (obj *_AllVirtualUserMgr) GetFromPrivCreateResourceUnit(privCreateResourceUnit int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_resource_unit` = ?", privCreateResourceUnit).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourceUnit 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromPrivCreateResourceUnit(privCreateResourceUnits []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`priv_create_resource_unit` IN (?)", privCreateResourceUnits).Find(&results).Error

	return
}

// GetFromMaxConnections 通过max_connections获取内容
func (obj *_AllVirtualUserMgr) GetFromMaxConnections(maxConnections int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`max_connections` = ?", maxConnections).Find(&results).Error

	return
}

// GetBatchFromMaxConnections 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromMaxConnections(maxConnectionss []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`max_connections` IN (?)", maxConnectionss).Find(&results).Error

	return
}

// GetFromMaxUserConnections 通过max_user_connections获取内容
func (obj *_AllVirtualUserMgr) GetFromMaxUserConnections(maxUserConnections int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`max_user_connections` = ?", maxUserConnections).Find(&results).Error

	return
}

// GetBatchFromMaxUserConnections 批量查找
func (obj *_AllVirtualUserMgr) GetBatchFromMaxUserConnections(maxUserConnectionss []int64) (results []*AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`max_user_connections` IN (?)", maxUserConnectionss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualUserMgr) FetchByPrimaryKey(tenantID int64, userID int64) (result AllVirtualUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUser{}).Where("`tenant_id` = ? AND `user_id` = ?", tenantID, userID).First(&result).Error

	return
}
