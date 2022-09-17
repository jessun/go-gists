package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllUserMgr struct {
	*_BaseMgr
}

// AllUserMgr open func
func AllUserMgr(db *gorm.DB) *_AllUserMgr {
	if db == nil {
		panic(fmt.Errorf("AllUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllUserMgr) GetTableName() string {
	return "__all_user"
}

// Reset 重置gorm会话
func (obj *_AllUserMgr) Reset() *_AllUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllUserMgr) Get() (result AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllUserMgr) Gets() (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllUserMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllUserMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllUserMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllUserMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserName user_name获取
func (obj *_AllUserMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithHost host获取
func (obj *_AllUserMgr) WithHost(host string) Option {
	return optionFunc(func(o *options) { o.query["host"] = host })
}

// WithPasswd passwd获取
func (obj *_AllUserMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithInfo info获取
func (obj *_AllUserMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithPrivAlter priv_alter获取
func (obj *_AllUserMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllUserMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllUserMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllUserMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllUserMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllUserMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllUserMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllUserMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllUserMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllUserMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllUserMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// WithPrivShowDb priv_show_db获取
func (obj *_AllUserMgr) WithPrivShowDb(privShowDb int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_db"] = privShowDb })
}

// WithPrivCreateUser priv_create_user获取
func (obj *_AllUserMgr) WithPrivCreateUser(privCreateUser int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_user"] = privCreateUser })
}

// WithPrivSuper priv_super获取
func (obj *_AllUserMgr) WithPrivSuper(privSuper int64) Option {
	return optionFunc(func(o *options) { o.query["priv_super"] = privSuper })
}

// WithIsLocked is_locked获取
func (obj *_AllUserMgr) WithIsLocked(isLocked int64) Option {
	return optionFunc(func(o *options) { o.query["is_locked"] = isLocked })
}

// WithPrivProcess priv_process获取
func (obj *_AllUserMgr) WithPrivProcess(privProcess int64) Option {
	return optionFunc(func(o *options) { o.query["priv_process"] = privProcess })
}

// WithPrivCreateSynonym priv_create_synonym获取
func (obj *_AllUserMgr) WithPrivCreateSynonym(privCreateSynonym int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_synonym"] = privCreateSynonym })
}

// WithSslType ssl_type获取
func (obj *_AllUserMgr) WithSslType(sslType int64) Option {
	return optionFunc(func(o *options) { o.query["ssl_type"] = sslType })
}

// WithSslCipher ssl_cipher获取
func (obj *_AllUserMgr) WithSslCipher(sslCipher string) Option {
	return optionFunc(func(o *options) { o.query["ssl_cipher"] = sslCipher })
}

// WithX509Issuer x509_issuer获取
func (obj *_AllUserMgr) WithX509Issuer(x509Issuer string) Option {
	return optionFunc(func(o *options) { o.query["x509_issuer"] = x509Issuer })
}

// WithX509Subject x509_subject获取
func (obj *_AllUserMgr) WithX509Subject(x509Subject string) Option {
	return optionFunc(func(o *options) { o.query["x509_subject"] = x509Subject })
}

// WithType type获取
func (obj *_AllUserMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithProfileID profile_id获取
func (obj *_AllUserMgr) WithProfileID(profileID int64) Option {
	return optionFunc(func(o *options) { o.query["profile_id"] = profileID })
}

// WithPasswordLastChanged password_last_changed获取
func (obj *_AllUserMgr) WithPasswordLastChanged(passwordLastChanged time.Time) Option {
	return optionFunc(func(o *options) { o.query["password_last_changed"] = passwordLastChanged })
}

// WithPrivFile priv_file获取
func (obj *_AllUserMgr) WithPrivFile(privFile int64) Option {
	return optionFunc(func(o *options) { o.query["priv_file"] = privFile })
}

// WithPrivAlterTenant priv_alter_tenant获取
func (obj *_AllUserMgr) WithPrivAlterTenant(privAlterTenant int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_tenant"] = privAlterTenant })
}

// WithPrivAlterSystem priv_alter_system获取
func (obj *_AllUserMgr) WithPrivAlterSystem(privAlterSystem int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_system"] = privAlterSystem })
}

// WithPrivCreateResourcePool priv_create_resource_pool获取
func (obj *_AllUserMgr) WithPrivCreateResourcePool(privCreateResourcePool int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_pool"] = privCreateResourcePool })
}

// WithPrivCreateResourceUnit priv_create_resource_unit获取
func (obj *_AllUserMgr) WithPrivCreateResourceUnit(privCreateResourceUnit int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_unit"] = privCreateResourceUnit })
}

// WithMaxConnections max_connections获取
func (obj *_AllUserMgr) WithMaxConnections(maxConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_connections"] = maxConnections })
}

// WithMaxUserConnections max_user_connections获取
func (obj *_AllUserMgr) WithMaxUserConnections(maxUserConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_user_connections"] = maxUserConnections })
}

// GetByOption 功能选项模式获取
func (obj *_AllUserMgr) GetByOption(opts ...Option) (result AllUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllUserMgr) GetByOptions(opts ...Option) (results []*AllUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where(options.query)
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
func (obj *_AllUserMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllUserMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllUserMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllUserMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllUserMgr) GetFromTenantID(tenantID int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllUserMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllUserMgr) GetFromUserID(userID int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllUserMgr) GetBatchFromUserID(userIDs []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllUserMgr) GetFromUserName(userName string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllUserMgr) GetBatchFromUserName(userNames []string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromHost 通过host获取内容
func (obj *_AllUserMgr) GetFromHost(host string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`host` = ?", host).Find(&results).Error

	return
}

// GetBatchFromHost 批量查找
func (obj *_AllUserMgr) GetBatchFromHost(hosts []string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`host` IN (?)", hosts).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllUserMgr) GetFromPasswd(passwd string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllUserMgr) GetBatchFromPasswd(passwds []string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllUserMgr) GetFromInfo(info string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllUserMgr) GetBatchFromInfo(infos []string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllUserMgr) GetFromPrivAlter(privAlter int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllUserMgr) GetFromPrivCreate(privCreate int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllUserMgr) GetFromPrivDelete(privDelete int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllUserMgr) GetFromPrivDrop(privDrop int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllUserMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllUserMgr) GetFromPrivInsert(privInsert int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllUserMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllUserMgr) GetFromPrivSelect(privSelect int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllUserMgr) GetFromPrivIndex(privIndex int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllUserMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllUserMgr) GetFromPrivShowView(privShowView int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

// GetFromPrivShowDb 通过priv_show_db获取内容
func (obj *_AllUserMgr) GetFromPrivShowDb(privShowDb int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_show_db` = ?", privShowDb).Find(&results).Error

	return
}

// GetBatchFromPrivShowDb 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivShowDb(privShowDbs []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_show_db` IN (?)", privShowDbs).Find(&results).Error

	return
}

// GetFromPrivCreateUser 通过priv_create_user获取内容
func (obj *_AllUserMgr) GetFromPrivCreateUser(privCreateUser int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_user` = ?", privCreateUser).Find(&results).Error

	return
}

// GetBatchFromPrivCreateUser 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivCreateUser(privCreateUsers []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_user` IN (?)", privCreateUsers).Find(&results).Error

	return
}

// GetFromPrivSuper 通过priv_super获取内容
func (obj *_AllUserMgr) GetFromPrivSuper(privSuper int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_super` = ?", privSuper).Find(&results).Error

	return
}

// GetBatchFromPrivSuper 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivSuper(privSupers []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_super` IN (?)", privSupers).Find(&results).Error

	return
}

// GetFromIsLocked 通过is_locked获取内容
func (obj *_AllUserMgr) GetFromIsLocked(isLocked int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`is_locked` = ?", isLocked).Find(&results).Error

	return
}

// GetBatchFromIsLocked 批量查找
func (obj *_AllUserMgr) GetBatchFromIsLocked(isLockeds []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`is_locked` IN (?)", isLockeds).Find(&results).Error

	return
}

// GetFromPrivProcess 通过priv_process获取内容
func (obj *_AllUserMgr) GetFromPrivProcess(privProcess int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_process` = ?", privProcess).Find(&results).Error

	return
}

// GetBatchFromPrivProcess 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivProcess(privProcesss []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_process` IN (?)", privProcesss).Find(&results).Error

	return
}

// GetFromPrivCreateSynonym 通过priv_create_synonym获取内容
func (obj *_AllUserMgr) GetFromPrivCreateSynonym(privCreateSynonym int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_synonym` = ?", privCreateSynonym).Find(&results).Error

	return
}

// GetBatchFromPrivCreateSynonym 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivCreateSynonym(privCreateSynonyms []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_synonym` IN (?)", privCreateSynonyms).Find(&results).Error

	return
}

// GetFromSslType 通过ssl_type获取内容
func (obj *_AllUserMgr) GetFromSslType(sslType int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`ssl_type` = ?", sslType).Find(&results).Error

	return
}

// GetBatchFromSslType 批量查找
func (obj *_AllUserMgr) GetBatchFromSslType(sslTypes []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`ssl_type` IN (?)", sslTypes).Find(&results).Error

	return
}

// GetFromSslCipher 通过ssl_cipher获取内容
func (obj *_AllUserMgr) GetFromSslCipher(sslCipher string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`ssl_cipher` = ?", sslCipher).Find(&results).Error

	return
}

// GetBatchFromSslCipher 批量查找
func (obj *_AllUserMgr) GetBatchFromSslCipher(sslCiphers []string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`ssl_cipher` IN (?)", sslCiphers).Find(&results).Error

	return
}

// GetFromX509Issuer 通过x509_issuer获取内容
func (obj *_AllUserMgr) GetFromX509Issuer(x509Issuer string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`x509_issuer` = ?", x509Issuer).Find(&results).Error

	return
}

// GetBatchFromX509Issuer 批量查找
func (obj *_AllUserMgr) GetBatchFromX509Issuer(x509Issuers []string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`x509_issuer` IN (?)", x509Issuers).Find(&results).Error

	return
}

// GetFromX509Subject 通过x509_subject获取内容
func (obj *_AllUserMgr) GetFromX509Subject(x509Subject string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`x509_subject` = ?", x509Subject).Find(&results).Error

	return
}

// GetBatchFromX509Subject 批量查找
func (obj *_AllUserMgr) GetBatchFromX509Subject(x509Subjects []string) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`x509_subject` IN (?)", x509Subjects).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllUserMgr) GetFromType(_type int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllUserMgr) GetBatchFromType(_types []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromProfileID 通过profile_id获取内容
func (obj *_AllUserMgr) GetFromProfileID(profileID int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`profile_id` = ?", profileID).Find(&results).Error

	return
}

// GetBatchFromProfileID 批量查找
func (obj *_AllUserMgr) GetBatchFromProfileID(profileIDs []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`profile_id` IN (?)", profileIDs).Find(&results).Error

	return
}

// GetFromPasswordLastChanged 通过password_last_changed获取内容
func (obj *_AllUserMgr) GetFromPasswordLastChanged(passwordLastChanged time.Time) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`password_last_changed` = ?", passwordLastChanged).Find(&results).Error

	return
}

// GetBatchFromPasswordLastChanged 批量查找
func (obj *_AllUserMgr) GetBatchFromPasswordLastChanged(passwordLastChangeds []time.Time) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`password_last_changed` IN (?)", passwordLastChangeds).Find(&results).Error

	return
}

// GetFromPrivFile 通过priv_file获取内容
func (obj *_AllUserMgr) GetFromPrivFile(privFile int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_file` = ?", privFile).Find(&results).Error

	return
}

// GetBatchFromPrivFile 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivFile(privFiles []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_file` IN (?)", privFiles).Find(&results).Error

	return
}

// GetFromPrivAlterTenant 通过priv_alter_tenant获取内容
func (obj *_AllUserMgr) GetFromPrivAlterTenant(privAlterTenant int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_alter_tenant` = ?", privAlterTenant).Find(&results).Error

	return
}

// GetBatchFromPrivAlterTenant 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivAlterTenant(privAlterTenants []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_alter_tenant` IN (?)", privAlterTenants).Find(&results).Error

	return
}

// GetFromPrivAlterSystem 通过priv_alter_system获取内容
func (obj *_AllUserMgr) GetFromPrivAlterSystem(privAlterSystem int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_alter_system` = ?", privAlterSystem).Find(&results).Error

	return
}

// GetBatchFromPrivAlterSystem 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivAlterSystem(privAlterSystems []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_alter_system` IN (?)", privAlterSystems).Find(&results).Error

	return
}

// GetFromPrivCreateResourcePool 通过priv_create_resource_pool获取内容
func (obj *_AllUserMgr) GetFromPrivCreateResourcePool(privCreateResourcePool int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_resource_pool` = ?", privCreateResourcePool).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourcePool 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivCreateResourcePool(privCreateResourcePools []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_resource_pool` IN (?)", privCreateResourcePools).Find(&results).Error

	return
}

// GetFromPrivCreateResourceUnit 通过priv_create_resource_unit获取内容
func (obj *_AllUserMgr) GetFromPrivCreateResourceUnit(privCreateResourceUnit int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_resource_unit` = ?", privCreateResourceUnit).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourceUnit 批量查找
func (obj *_AllUserMgr) GetBatchFromPrivCreateResourceUnit(privCreateResourceUnits []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`priv_create_resource_unit` IN (?)", privCreateResourceUnits).Find(&results).Error

	return
}

// GetFromMaxConnections 通过max_connections获取内容
func (obj *_AllUserMgr) GetFromMaxConnections(maxConnections int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`max_connections` = ?", maxConnections).Find(&results).Error

	return
}

// GetBatchFromMaxConnections 批量查找
func (obj *_AllUserMgr) GetBatchFromMaxConnections(maxConnectionss []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`max_connections` IN (?)", maxConnectionss).Find(&results).Error

	return
}

// GetFromMaxUserConnections 通过max_user_connections获取内容
func (obj *_AllUserMgr) GetFromMaxUserConnections(maxUserConnections int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`max_user_connections` = ?", maxUserConnections).Find(&results).Error

	return
}

// GetBatchFromMaxUserConnections 批量查找
func (obj *_AllUserMgr) GetBatchFromMaxUserConnections(maxUserConnectionss []int64) (results []*AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`max_user_connections` IN (?)", maxUserConnectionss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllUserMgr) FetchByPrimaryKey(tenantID int64, userID int64) (result AllUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUser{}).Where("`tenant_id` = ? AND `user_id` = ?", tenantID, userID).First(&result).Error

	return
}
