package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualUserHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualUserHistoryMgr open func
func AllVirtualUserHistoryMgr(db *gorm.DB) *_AllVirtualUserHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualUserHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualUserHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_user_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualUserHistoryMgr) GetTableName() string {
	return "__all_virtual_user_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualUserHistoryMgr) Reset() *_AllVirtualUserHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualUserHistoryMgr) Get() (result AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualUserHistoryMgr) Gets() (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualUserHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualUserHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllVirtualUserHistoryMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualUserHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualUserHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualUserHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualUserHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithUserName user_name获取
func (obj *_AllVirtualUserHistoryMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithHost host获取
func (obj *_AllVirtualUserHistoryMgr) WithHost(host string) Option {
	return optionFunc(func(o *options) { o.query["host"] = host })
}

// WithPasswd passwd获取
func (obj *_AllVirtualUserHistoryMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithInfo info获取
func (obj *_AllVirtualUserHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithPrivAlter priv_alter获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivAlter(privAlter int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter"] = privAlter })
}

// WithPrivCreate priv_create获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivCreate(privCreate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create"] = privCreate })
}

// WithPrivDelete priv_delete获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivDelete(privDelete int64) Option {
	return optionFunc(func(o *options) { o.query["priv_delete"] = privDelete })
}

// WithPrivDrop priv_drop获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivDrop(privDrop int64) Option {
	return optionFunc(func(o *options) { o.query["priv_drop"] = privDrop })
}

// WithPrivGrantOption priv_grant_option获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivGrantOption(privGrantOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_grant_option"] = privGrantOption })
}

// WithPrivInsert priv_insert获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivInsert(privInsert int64) Option {
	return optionFunc(func(o *options) { o.query["priv_insert"] = privInsert })
}

// WithPrivUpdate priv_update获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivUpdate(privUpdate int64) Option {
	return optionFunc(func(o *options) { o.query["priv_update"] = privUpdate })
}

// WithPrivSelect priv_select获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivSelect(privSelect int64) Option {
	return optionFunc(func(o *options) { o.query["priv_select"] = privSelect })
}

// WithPrivIndex priv_index获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivIndex(privIndex int64) Option {
	return optionFunc(func(o *options) { o.query["priv_index"] = privIndex })
}

// WithPrivCreateView priv_create_view获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivCreateView(privCreateView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_view"] = privCreateView })
}

// WithPrivShowView priv_show_view获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivShowView(privShowView int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_view"] = privShowView })
}

// WithPrivShowDb priv_show_db获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivShowDb(privShowDb int64) Option {
	return optionFunc(func(o *options) { o.query["priv_show_db"] = privShowDb })
}

// WithPrivCreateUser priv_create_user获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivCreateUser(privCreateUser int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_user"] = privCreateUser })
}

// WithPrivSuper priv_super获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivSuper(privSuper int64) Option {
	return optionFunc(func(o *options) { o.query["priv_super"] = privSuper })
}

// WithIsLocked is_locked获取
func (obj *_AllVirtualUserHistoryMgr) WithIsLocked(isLocked int64) Option {
	return optionFunc(func(o *options) { o.query["is_locked"] = isLocked })
}

// WithPrivProcess priv_process获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivProcess(privProcess int64) Option {
	return optionFunc(func(o *options) { o.query["priv_process"] = privProcess })
}

// WithPrivCreateSynonym priv_create_synonym获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivCreateSynonym(privCreateSynonym int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_synonym"] = privCreateSynonym })
}

// WithSslType ssl_type获取
func (obj *_AllVirtualUserHistoryMgr) WithSslType(sslType int64) Option {
	return optionFunc(func(o *options) { o.query["ssl_type"] = sslType })
}

// WithSslCipher ssl_cipher获取
func (obj *_AllVirtualUserHistoryMgr) WithSslCipher(sslCipher string) Option {
	return optionFunc(func(o *options) { o.query["ssl_cipher"] = sslCipher })
}

// WithX509Issuer x509_issuer获取
func (obj *_AllVirtualUserHistoryMgr) WithX509Issuer(x509Issuer string) Option {
	return optionFunc(func(o *options) { o.query["x509_issuer"] = x509Issuer })
}

// WithX509Subject x509_subject获取
func (obj *_AllVirtualUserHistoryMgr) WithX509Subject(x509Subject string) Option {
	return optionFunc(func(o *options) { o.query["x509_subject"] = x509Subject })
}

// WithType type获取
func (obj *_AllVirtualUserHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithProfileID profile_id获取
func (obj *_AllVirtualUserHistoryMgr) WithProfileID(profileID int64) Option {
	return optionFunc(func(o *options) { o.query["profile_id"] = profileID })
}

// WithPasswordLastChanged password_last_changed获取
func (obj *_AllVirtualUserHistoryMgr) WithPasswordLastChanged(passwordLastChanged time.Time) Option {
	return optionFunc(func(o *options) { o.query["password_last_changed"] = passwordLastChanged })
}

// WithPrivFile priv_file获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivFile(privFile int64) Option {
	return optionFunc(func(o *options) { o.query["priv_file"] = privFile })
}

// WithPrivAlterTenant priv_alter_tenant获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivAlterTenant(privAlterTenant int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_tenant"] = privAlterTenant })
}

// WithPrivAlterSystem priv_alter_system获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivAlterSystem(privAlterSystem int64) Option {
	return optionFunc(func(o *options) { o.query["priv_alter_system"] = privAlterSystem })
}

// WithPrivCreateResourcePool priv_create_resource_pool获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivCreateResourcePool(privCreateResourcePool int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_pool"] = privCreateResourcePool })
}

// WithPrivCreateResourceUnit priv_create_resource_unit获取
func (obj *_AllVirtualUserHistoryMgr) WithPrivCreateResourceUnit(privCreateResourceUnit int64) Option {
	return optionFunc(func(o *options) { o.query["priv_create_resource_unit"] = privCreateResourceUnit })
}

// WithMaxConnections max_connections获取
func (obj *_AllVirtualUserHistoryMgr) WithMaxConnections(maxConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_connections"] = maxConnections })
}

// WithMaxUserConnections max_user_connections获取
func (obj *_AllVirtualUserHistoryMgr) WithMaxUserConnections(maxUserConnections int64) Option {
	return optionFunc(func(o *options) { o.query["max_user_connections"] = maxUserConnections })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualUserHistoryMgr) GetByOption(opts ...Option) (result AllVirtualUserHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualUserHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualUserHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualUserHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualUserHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where(options.query)
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
func (obj *_AllVirtualUserHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromUserID(userID int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromUserName(userName string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromHost 通过host获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromHost(host string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`host` = ?", host).Find(&results).Error

	return
}

// GetBatchFromHost 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromHost(hosts []string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`host` IN (?)", hosts).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPasswd(passwd string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPasswd(passwds []string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromInfo(info string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromPrivAlter 通过priv_alter获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivAlter(privAlter int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_alter` = ?", privAlter).Find(&results).Error

	return
}

// GetBatchFromPrivAlter 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivAlter(privAlters []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_alter` IN (?)", privAlters).Find(&results).Error

	return
}

// GetFromPrivCreate 通过priv_create获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivCreate(privCreate int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create` = ?", privCreate).Find(&results).Error

	return
}

// GetBatchFromPrivCreate 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivCreate(privCreates []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create` IN (?)", privCreates).Find(&results).Error

	return
}

// GetFromPrivDelete 通过priv_delete获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivDelete(privDelete int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_delete` = ?", privDelete).Find(&results).Error

	return
}

// GetBatchFromPrivDelete 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivDelete(privDeletes []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_delete` IN (?)", privDeletes).Find(&results).Error

	return
}

// GetFromPrivDrop 通过priv_drop获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivDrop(privDrop int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_drop` = ?", privDrop).Find(&results).Error

	return
}

// GetBatchFromPrivDrop 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivDrop(privDrops []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_drop` IN (?)", privDrops).Find(&results).Error

	return
}

// GetFromPrivGrantOption 通过priv_grant_option获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivGrantOption(privGrantOption int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_grant_option` = ?", privGrantOption).Find(&results).Error

	return
}

// GetBatchFromPrivGrantOption 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivGrantOption(privGrantOptions []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_grant_option` IN (?)", privGrantOptions).Find(&results).Error

	return
}

// GetFromPrivInsert 通过priv_insert获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivInsert(privInsert int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_insert` = ?", privInsert).Find(&results).Error

	return
}

// GetBatchFromPrivInsert 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivInsert(privInserts []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_insert` IN (?)", privInserts).Find(&results).Error

	return
}

// GetFromPrivUpdate 通过priv_update获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivUpdate(privUpdate int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_update` = ?", privUpdate).Find(&results).Error

	return
}

// GetBatchFromPrivUpdate 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivUpdate(privUpdates []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_update` IN (?)", privUpdates).Find(&results).Error

	return
}

// GetFromPrivSelect 通过priv_select获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivSelect(privSelect int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_select` = ?", privSelect).Find(&results).Error

	return
}

// GetBatchFromPrivSelect 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivSelect(privSelects []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_select` IN (?)", privSelects).Find(&results).Error

	return
}

// GetFromPrivIndex 通过priv_index获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivIndex(privIndex int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_index` = ?", privIndex).Find(&results).Error

	return
}

// GetBatchFromPrivIndex 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivIndex(privIndexs []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_index` IN (?)", privIndexs).Find(&results).Error

	return
}

// GetFromPrivCreateView 通过priv_create_view获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivCreateView(privCreateView int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_view` = ?", privCreateView).Find(&results).Error

	return
}

// GetBatchFromPrivCreateView 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivCreateView(privCreateViews []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_view` IN (?)", privCreateViews).Find(&results).Error

	return
}

// GetFromPrivShowView 通过priv_show_view获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivShowView(privShowView int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_show_view` = ?", privShowView).Find(&results).Error

	return
}

// GetBatchFromPrivShowView 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivShowView(privShowViews []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_show_view` IN (?)", privShowViews).Find(&results).Error

	return
}

// GetFromPrivShowDb 通过priv_show_db获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivShowDb(privShowDb int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_show_db` = ?", privShowDb).Find(&results).Error

	return
}

// GetBatchFromPrivShowDb 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivShowDb(privShowDbs []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_show_db` IN (?)", privShowDbs).Find(&results).Error

	return
}

// GetFromPrivCreateUser 通过priv_create_user获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivCreateUser(privCreateUser int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_user` = ?", privCreateUser).Find(&results).Error

	return
}

// GetBatchFromPrivCreateUser 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivCreateUser(privCreateUsers []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_user` IN (?)", privCreateUsers).Find(&results).Error

	return
}

// GetFromPrivSuper 通过priv_super获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivSuper(privSuper int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_super` = ?", privSuper).Find(&results).Error

	return
}

// GetBatchFromPrivSuper 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivSuper(privSupers []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_super` IN (?)", privSupers).Find(&results).Error

	return
}

// GetFromIsLocked 通过is_locked获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromIsLocked(isLocked int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`is_locked` = ?", isLocked).Find(&results).Error

	return
}

// GetBatchFromIsLocked 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromIsLocked(isLockeds []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`is_locked` IN (?)", isLockeds).Find(&results).Error

	return
}

// GetFromPrivProcess 通过priv_process获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivProcess(privProcess int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_process` = ?", privProcess).Find(&results).Error

	return
}

// GetBatchFromPrivProcess 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivProcess(privProcesss []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_process` IN (?)", privProcesss).Find(&results).Error

	return
}

// GetFromPrivCreateSynonym 通过priv_create_synonym获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivCreateSynonym(privCreateSynonym int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_synonym` = ?", privCreateSynonym).Find(&results).Error

	return
}

// GetBatchFromPrivCreateSynonym 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivCreateSynonym(privCreateSynonyms []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_synonym` IN (?)", privCreateSynonyms).Find(&results).Error

	return
}

// GetFromSslType 通过ssl_type获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromSslType(sslType int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`ssl_type` = ?", sslType).Find(&results).Error

	return
}

// GetBatchFromSslType 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromSslType(sslTypes []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`ssl_type` IN (?)", sslTypes).Find(&results).Error

	return
}

// GetFromSslCipher 通过ssl_cipher获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromSslCipher(sslCipher string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`ssl_cipher` = ?", sslCipher).Find(&results).Error

	return
}

// GetBatchFromSslCipher 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromSslCipher(sslCiphers []string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`ssl_cipher` IN (?)", sslCiphers).Find(&results).Error

	return
}

// GetFromX509Issuer 通过x509_issuer获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromX509Issuer(x509Issuer string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`x509_issuer` = ?", x509Issuer).Find(&results).Error

	return
}

// GetBatchFromX509Issuer 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromX509Issuer(x509Issuers []string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`x509_issuer` IN (?)", x509Issuers).Find(&results).Error

	return
}

// GetFromX509Subject 通过x509_subject获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromX509Subject(x509Subject string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`x509_subject` = ?", x509Subject).Find(&results).Error

	return
}

// GetBatchFromX509Subject 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromX509Subject(x509Subjects []string) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`x509_subject` IN (?)", x509Subjects).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromType(_type int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromType(_types []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromProfileID 通过profile_id获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromProfileID(profileID int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`profile_id` = ?", profileID).Find(&results).Error

	return
}

// GetBatchFromProfileID 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromProfileID(profileIDs []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`profile_id` IN (?)", profileIDs).Find(&results).Error

	return
}

// GetFromPasswordLastChanged 通过password_last_changed获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPasswordLastChanged(passwordLastChanged time.Time) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`password_last_changed` = ?", passwordLastChanged).Find(&results).Error

	return
}

// GetBatchFromPasswordLastChanged 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPasswordLastChanged(passwordLastChangeds []time.Time) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`password_last_changed` IN (?)", passwordLastChangeds).Find(&results).Error

	return
}

// GetFromPrivFile 通过priv_file获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivFile(privFile int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_file` = ?", privFile).Find(&results).Error

	return
}

// GetBatchFromPrivFile 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivFile(privFiles []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_file` IN (?)", privFiles).Find(&results).Error

	return
}

// GetFromPrivAlterTenant 通过priv_alter_tenant获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivAlterTenant(privAlterTenant int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_alter_tenant` = ?", privAlterTenant).Find(&results).Error

	return
}

// GetBatchFromPrivAlterTenant 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivAlterTenant(privAlterTenants []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_alter_tenant` IN (?)", privAlterTenants).Find(&results).Error

	return
}

// GetFromPrivAlterSystem 通过priv_alter_system获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivAlterSystem(privAlterSystem int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_alter_system` = ?", privAlterSystem).Find(&results).Error

	return
}

// GetBatchFromPrivAlterSystem 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivAlterSystem(privAlterSystems []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_alter_system` IN (?)", privAlterSystems).Find(&results).Error

	return
}

// GetFromPrivCreateResourcePool 通过priv_create_resource_pool获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivCreateResourcePool(privCreateResourcePool int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_resource_pool` = ?", privCreateResourcePool).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourcePool 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivCreateResourcePool(privCreateResourcePools []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_resource_pool` IN (?)", privCreateResourcePools).Find(&results).Error

	return
}

// GetFromPrivCreateResourceUnit 通过priv_create_resource_unit获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromPrivCreateResourceUnit(privCreateResourceUnit int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_resource_unit` = ?", privCreateResourceUnit).Find(&results).Error

	return
}

// GetBatchFromPrivCreateResourceUnit 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromPrivCreateResourceUnit(privCreateResourceUnits []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`priv_create_resource_unit` IN (?)", privCreateResourceUnits).Find(&results).Error

	return
}

// GetFromMaxConnections 通过max_connections获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromMaxConnections(maxConnections int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`max_connections` = ?", maxConnections).Find(&results).Error

	return
}

// GetBatchFromMaxConnections 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromMaxConnections(maxConnectionss []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`max_connections` IN (?)", maxConnectionss).Find(&results).Error

	return
}

// GetFromMaxUserConnections 通过max_user_connections获取内容
func (obj *_AllVirtualUserHistoryMgr) GetFromMaxUserConnections(maxUserConnections int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`max_user_connections` = ?", maxUserConnections).Find(&results).Error

	return
}

// GetBatchFromMaxUserConnections 批量查找
func (obj *_AllVirtualUserHistoryMgr) GetBatchFromMaxUserConnections(maxUserConnectionss []int64) (results []*AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`max_user_connections` IN (?)", maxUserConnectionss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualUserHistoryMgr) FetchByPrimaryKey(tenantID int64, userID int64, schemaVersion int64) (result AllVirtualUserHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualUserHistory{}).Where("`tenant_id` = ? AND `user_id` = ? AND `schema_version` = ?", tenantID, userID, schemaVersion).First(&result).Error

	return
}
