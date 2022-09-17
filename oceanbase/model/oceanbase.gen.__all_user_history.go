package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllUserHistoryMgr struct {
	*_BaseMgr
}

// AllUserHistoryMgr open func
func AllUserHistoryMgr(db *gorm.DB) *_AllUserHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllUserHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllUserHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_user_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllUserHistoryMgr) GetTableName() string {
	return "__all_user_history"
}

// Reset 重置gorm会话
func (obj *_AllUserHistoryMgr) Reset() *_AllUserHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllUserHistoryMgr) Get() (result AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllUserHistoryMgr) Gets() (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllUserHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllUserHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllUserHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllUserHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllUserHistoryMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllUserHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllUserHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithUserName user_name获取
func (obj *_AllUserHistoryMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithHost host获取
func (obj *_AllUserHistoryMgr) WithHost(host string) Option {
	return optionFunc(func(o *options) { o.query["host"] = host })
}

// WithPasswd passwd获取
func (obj *_AllUserHistoryMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithInfo info获取
func (obj *_AllUserHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithPrivAlter priv_alter获取
func (obj *_AllUserHistoryMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllUserHistoryMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllUserHistoryMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllUserHistoryMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllUserHistoryMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllUserHistoryMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllUserHistoryMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllUserHistoryMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllUserHistoryMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllUserHistoryMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllUserHistoryMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// WithPrivShowDb priv_show_db获取
func (obj *_AllUserHistoryMgr) WithPrivShowDb(privShowDb int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_db"] = privShowDb })
}

// WithPrivCreateUser priv_create_user获取
func (obj *_AllUserHistoryMgr) WithPrivCreateUser(privCreateUser int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_user"] = privCreateUser })
}

// WithPrivSuper priv_super获取
func (obj *_AllUserHistoryMgr) WithPrivSuper(privSuper int64) Option {
	return optionFunc(func(o *options) { o.query["priv_super"] = privSuper })
}

// WithIsLocked is_locked获取
func (obj *_AllUserHistoryMgr) WithIsLocked(isLocked int64) Option {
	return optionFunc(func(o *options) { o.query["is_locked"] = isLocked })
}

// WithPrivProcess priv_process获取
func (obj *_AllUserHistoryMgr) WithPrivProcess(privProcess int64) Option {
	return optionFunc(func(o *options) { o.query["priv_process"] = privProcess })
}

// WithPrivCreateSynonym priv_create_synonym获取
func (obj *_AllUserHistoryMgr) WithPrivCreateSynonym(privCreateSynonym int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_synonym"] = privCreateSynonym })
}

// WithSslType ssl_type获取
func (obj *_AllUserHistoryMgr) WithSslType(sslType int64) Option {
	return optionFunc(func(o *options) { o.query["ssl_type"] = sslType })
}

// WithSslCipher ssl_cipher获取
func (obj *_AllUserHistoryMgr) WithSslCipher(sslCipher string) Option {
	return optionFunc(func(o *options) { o.query["ssl_cipher"] = sslCipher })
}

// WithX509Issuer x509_issuer获取
func (obj *_AllUserHistoryMgr) WithX509Issuer(x509Issuer string) Option {
	return optionFunc(func(o *options) { o.query["x509_issuer"] = x509Issuer })
}

// WithX509Subject x509_subject获取
func (obj *_AllUserHistoryMgr) WithX509Subject(x509Subject string) Option {
	return optionFunc(func(o *options) { o.query["x509_subject"] = x509Subject })
}

// WithType type获取
func (obj *_AllUserHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithProfileID profile_id获取
func (obj *_AllUserHistoryMgr) WithProfileID(profileID int64) Option {
	return optionFunc(func(o *options) { o.query["profile_id"] = profileID })
}

// WithPasswordLastChanged password_last_changed获取
func (obj *_AllUserHistoryMgr) WithPasswordLastChanged(passwordLastChanged time.Time) Option {
	return optionFunc(func(o *options) { o.query["password_last_changed"] = passwordLastChanged })
}

// WithPrivFile priv_file获取
func (obj *_AllUserHistoryMgr) WithPrivFile(privFile int64) Option {
	return optionFunc(func(o *options) { o.query["priv_file"] = privFile })
}

// WithPrivAlterTenant priv_alter_tenant获取
func (obj *_AllUserHistoryMgr) WithPrivAlterTenant(privAlterTenant int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_tenant"] = privAlterTenant })
}

// WithPrivAlterSystem priv_alter_system获取
func (obj *_AllUserHistoryMgr) WithPrivAlterSystem(privAlterSystem int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_system"] = privAlterSystem })
}

// WithPrivCreateResourcePool priv_create_resource_pool获取
func (obj *_AllUserHistoryMgr) WithPrivCreateResourcePool(privCreateResourcePool int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_pool"] = privCreateResourcePool })
}

// WithPrivCreateResourceUnit priv_create_resource_unit获取
func (obj *_AllUserHistoryMgr) WithPrivCreateResourceUnit(privCreateResourceUnit int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_unit"] = privCreateResourceUnit })
}

// WithMaxConnections max_connections获取
func (obj *_AllUserHistoryMgr) WithMaxConnections(maxConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_connections"] = maxConnections })
}

// WithMaxUserConnections max_user_connections获取
func (obj *_AllUserHistoryMgr) WithMaxUserConnections(maxUserConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_user_connections"] = maxUserConnections })
}

// GetByOption 功能选项模式获取
func (obj *_AllUserHistoryMgr) GetByOption(opts ...Option) (result AllUserHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllUserHistoryMgr) GetByOptions(opts ...Option) (results []*AllUserHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllUserHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllUserHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where(options.query)
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
func (obj *_AllUserHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllUserHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllUserHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllUserHistoryMgr) GetFromUserID(userID int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromUserID(userIDs []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllUserHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllUserHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllUserHistoryMgr) GetFromUserName(userName string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromUserName(userNames []string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromHost 通过host获取内容
func (obj *_AllUserHistoryMgr) GetFromHost(host string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`host` = ?", host).Find(&results).Error

	return
}

// GetBatchFromHost 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromHost(hosts []string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`host` IN (?)", hosts).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllUserHistoryMgr) GetFromPasswd(passwd string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPasswd(passwds []string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllUserHistoryMgr) GetFromInfo(info string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivAlter(privAlter int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivCreate(privCreate int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivDelete(privDelete int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivDrop(privDrop int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivInsert(privInsert int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivSelect(privSelect int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivIndex(privIndex int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivShowView(privShowView int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

// GetFromPrivShowDb 通过priv_show_db获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivShowDb(privShowDb int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_show_db` = ?", privShowDb).Find(&results).Error

	return
}

// GetBatchFromPrivShowDb 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivShowDb(privShowDbs []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_show_db` IN (?)", privShowDbs).Find(&results).Error

	return
}

// GetFromPrivCreateUser 通过priv_create_user获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivCreateUser(privCreateUser int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_user` = ?", privCreateUser).Find(&results).Error

	return
}

// GetBatchFromPrivCreateUser 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivCreateUser(privCreateUsers []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_user` IN (?)", privCreateUsers).Find(&results).Error

	return
}

// GetFromPrivSuper 通过priv_super获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivSuper(privSuper int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_super` = ?", privSuper).Find(&results).Error

	return
}

// GetBatchFromPrivSuper 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivSuper(privSupers []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_super` IN (?)", privSupers).Find(&results).Error

	return
}

// GetFromIsLocked 通过is_locked获取内容
func (obj *_AllUserHistoryMgr) GetFromIsLocked(isLocked int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`is_locked` = ?", isLocked).Find(&results).Error

	return
}

// GetBatchFromIsLocked 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromIsLocked(isLockeds []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`is_locked` IN (?)", isLockeds).Find(&results).Error

	return
}

// GetFromPrivProcess 通过priv_process获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivProcess(privProcess int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_process` = ?", privProcess).Find(&results).Error

	return
}

// GetBatchFromPrivProcess 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivProcess(privProcesss []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_process` IN (?)", privProcesss).Find(&results).Error

	return
}

// GetFromPrivCreateSynonym 通过priv_create_synonym获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivCreateSynonym(privCreateSynonym int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_synonym` = ?", privCreateSynonym).Find(&results).Error

	return
}

// GetBatchFromPrivCreateSynonym 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivCreateSynonym(privCreateSynonyms []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_synonym` IN (?)", privCreateSynonyms).Find(&results).Error

	return
}

// GetFromSslType 通过ssl_type获取内容
func (obj *_AllUserHistoryMgr) GetFromSslType(sslType int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`ssl_type` = ?", sslType).Find(&results).Error

	return
}

// GetBatchFromSslType 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromSslType(sslTypes []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`ssl_type` IN (?)", sslTypes).Find(&results).Error

	return
}

// GetFromSslCipher 通过ssl_cipher获取内容
func (obj *_AllUserHistoryMgr) GetFromSslCipher(sslCipher string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`ssl_cipher` = ?", sslCipher).Find(&results).Error

	return
}

// GetBatchFromSslCipher 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromSslCipher(sslCiphers []string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`ssl_cipher` IN (?)", sslCiphers).Find(&results).Error

	return
}

// GetFromX509Issuer 通过x509_issuer获取内容
func (obj *_AllUserHistoryMgr) GetFromX509Issuer(x509Issuer string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`x509_issuer` = ?", x509Issuer).Find(&results).Error

	return
}

// GetBatchFromX509Issuer 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromX509Issuer(x509Issuers []string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`x509_issuer` IN (?)", x509Issuers).Find(&results).Error

	return
}

// GetFromX509Subject 通过x509_subject获取内容
func (obj *_AllUserHistoryMgr) GetFromX509Subject(x509Subject string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`x509_subject` = ?", x509Subject).Find(&results).Error

	return
}

// GetBatchFromX509Subject 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromX509Subject(x509Subjects []string) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`x509_subject` IN (?)", x509Subjects).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllUserHistoryMgr) GetFromType(_type int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromType(_types []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromProfileID 通过profile_id获取内容
func (obj *_AllUserHistoryMgr) GetFromProfileID(profileID int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`profile_id` = ?", profileID).Find(&results).Error

	return
}

// GetBatchFromProfileID 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromProfileID(profileIDs []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`profile_id` IN (?)", profileIDs).Find(&results).Error

	return
}

// GetFromPasswordLastChanged 通过password_last_changed获取内容
func (obj *_AllUserHistoryMgr) GetFromPasswordLastChanged(passwordLastChanged time.Time) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`password_last_changed` = ?", passwordLastChanged).Find(&results).Error

	return
}

// GetBatchFromPasswordLastChanged 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPasswordLastChanged(passwordLastChangeds []time.Time) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`password_last_changed` IN (?)", passwordLastChangeds).Find(&results).Error

	return
}

// GetFromPrivFile 通过priv_file获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivFile(privFile int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_file` = ?", privFile).Find(&results).Error

	return
}

// GetBatchFromPrivFile 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivFile(privFiles []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_file` IN (?)", privFiles).Find(&results).Error

	return
}

// GetFromPrivAlterTenant 通过priv_alter_tenant获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivAlterTenant(privAlterTenant int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_alter_tenant` = ?", privAlterTenant).Find(&results).Error

	return
}

// GetBatchFromPrivAlterTenant 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivAlterTenant(privAlterTenants []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_alter_tenant` IN (?)", privAlterTenants).Find(&results).Error

	return
}

// GetFromPrivAlterSystem 通过priv_alter_system获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivAlterSystem(privAlterSystem int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_alter_system` = ?", privAlterSystem).Find(&results).Error

	return
}

// GetBatchFromPrivAlterSystem 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivAlterSystem(privAlterSystems []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_alter_system` IN (?)", privAlterSystems).Find(&results).Error

	return
}

// GetFromPrivCreateResourcePool 通过priv_create_resource_pool获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivCreateResourcePool(privCreateResourcePool int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_resource_pool` = ?", privCreateResourcePool).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourcePool 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivCreateResourcePool(privCreateResourcePools []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_resource_pool` IN (?)", privCreateResourcePools).Find(&results).Error

	return
}

// GetFromPrivCreateResourceUnit 通过priv_create_resource_unit获取内容
func (obj *_AllUserHistoryMgr) GetFromPrivCreateResourceUnit(privCreateResourceUnit int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_resource_unit` = ?", privCreateResourceUnit).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourceUnit 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromPrivCreateResourceUnit(privCreateResourceUnits []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`priv_create_resource_unit` IN (?)", privCreateResourceUnits).Find(&results).Error

	return
}

// GetFromMaxConnections 通过max_connections获取内容
func (obj *_AllUserHistoryMgr) GetFromMaxConnections(maxConnections int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`max_connections` = ?", maxConnections).Find(&results).Error

	return
}

// GetBatchFromMaxConnections 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromMaxConnections(maxConnectionss []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`max_connections` IN (?)", maxConnectionss).Find(&results).Error

	return
}

// GetFromMaxUserConnections 通过max_user_connections获取内容
func (obj *_AllUserHistoryMgr) GetFromMaxUserConnections(maxUserConnections int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`max_user_connections` = ?", maxUserConnections).Find(&results).Error

	return
}

// GetBatchFromMaxUserConnections 批量查找
func (obj *_AllUserHistoryMgr) GetBatchFromMaxUserConnections(maxUserConnectionss []int64) (results []*AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`max_user_connections` IN (?)", maxUserConnectionss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllUserHistoryMgr) FetchByPrimaryKey(tenantID int64, userID int64, schemaVersion int64) (result AllUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUserHistory{}).Where("`tenant_id` = ? AND `user_id` = ? AND `schema_version` = ?", tenantID, userID, schemaVersion).First(&result).Error

	return
}
