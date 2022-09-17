package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTypeHistoryMgr struct {
	*_BaseMgr
}

// AllTypeHistoryMgr open func
func AllTypeHistoryMgr(db *gorm.DB) *_AllTypeHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTypeHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTypeHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_type_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTypeHistoryMgr) GetTableName() string {
	return "__all_type_history"
}

// Reset 重置gorm会话
func (obj *_AllTypeHistoryMgr) Reset() *_AllTypeHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTypeHistoryMgr) Get() (result AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTypeHistoryMgr) Gets() (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTypeHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTypeHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTypeHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTypeHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllTypeHistoryMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTypeHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTypeHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllTypeHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithTypecode typecode获取
func (obj *_AllTypeHistoryMgr) WithTypecode(typecode int64) Option {
	return optionFunc(func(o *options) { o.query["typecode"] = typecode })
}

// WithProperties properties获取
func (obj *_AllTypeHistoryMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithAttributes attributes获取
func (obj *_AllTypeHistoryMgr) WithAttributes(attributes int64) Option {
	return optionFunc(func(o *options) { o.query["attributes"] = attributes })
}

// WithMethods methods获取
func (obj *_AllTypeHistoryMgr) WithMethods(methods int64) Option {
	return optionFunc(func(o *options) { o.query["methods"] = methods })
}

// WithHiddenmethods hiddenmethods获取
func (obj *_AllTypeHistoryMgr) WithHiddenmethods(hiddenmethods int64) Option {
	return optionFunc(func(o *options) { o.query["hiddenmethods"] = hiddenmethods })
}

// WithSupertypes supertypes获取
func (obj *_AllTypeHistoryMgr) WithSupertypes(supertypes int64) Option {
	return optionFunc(func(o *options) { o.query["supertypes"] = supertypes })
}

// WithSubtypes subtypes获取
func (obj *_AllTypeHistoryMgr) WithSubtypes(subtypes int64) Option {
	return optionFunc(func(o *options) { o.query["subtypes"] = subtypes })
}

// WithExterntype externtype获取
func (obj *_AllTypeHistoryMgr) WithExterntype(externtype int64) Option {
	return optionFunc(func(o *options) { o.query["externtype"] = externtype })
}

// WithExternname externname获取
func (obj *_AllTypeHistoryMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithHelperclassname helperclassname获取
func (obj *_AllTypeHistoryMgr) WithHelperclassname(helperclassname string) Option {
	return optionFunc(func(o *options) { o.query["helperclassname"] = helperclassname })
}

// WithLocalAttrs local_attrs获取
func (obj *_AllTypeHistoryMgr) WithLocalAttrs(localAttrs int64) Option {
	return optionFunc(func(o *options) { o.query["local_attrs"] = localAttrs })
}

// WithLocalMethods local_methods获取
func (obj *_AllTypeHistoryMgr) WithLocalMethods(localMethods int64) Option {
	return optionFunc(func(o *options) { o.query["local_methods"] = localMethods })
}

// WithSupertypeid supertypeid获取
func (obj *_AllTypeHistoryMgr) WithSupertypeid(supertypeid int64) Option {
	return optionFunc(func(o *options) { o.query["supertypeid"] = supertypeid })
}

// WithTypeName type_name获取
func (obj *_AllTypeHistoryMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithPackageID package_id获取
func (obj *_AllTypeHistoryMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// GetByOption 功能选项模式获取
func (obj *_AllTypeHistoryMgr) GetByOption(opts ...Option) (result AllTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTypeHistoryMgr) GetByOptions(opts ...Option) (results []*AllTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTypeHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTypeHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where(options.query)
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
func (obj *_AllTypeHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTypeHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTypeHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllTypeHistoryMgr) GetFromTypeID(typeID int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTypeHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTypeHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTypeHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromTypecode 通过typecode获取内容
func (obj *_AllTypeHistoryMgr) GetFromTypecode(typecode int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`typecode` = ?", typecode).Find(&results).Error

	return
}

// GetBatchFromTypecode 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromTypecode(typecodes []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`typecode` IN (?)", typecodes).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllTypeHistoryMgr) GetFromProperties(properties int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromProperties(propertiess []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromAttributes 通过attributes获取内容
func (obj *_AllTypeHistoryMgr) GetFromAttributes(attributes int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`attributes` = ?", attributes).Find(&results).Error

	return
}

// GetBatchFromAttributes 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromAttributes(attributess []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`attributes` IN (?)", attributess).Find(&results).Error

	return
}

// GetFromMethods 通过methods获取内容
func (obj *_AllTypeHistoryMgr) GetFromMethods(methods int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`methods` = ?", methods).Find(&results).Error

	return
}

// GetBatchFromMethods 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromMethods(methodss []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`methods` IN (?)", methodss).Find(&results).Error

	return
}

// GetFromHiddenmethods 通过hiddenmethods获取内容
func (obj *_AllTypeHistoryMgr) GetFromHiddenmethods(hiddenmethods int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`hiddenmethods` = ?", hiddenmethods).Find(&results).Error

	return
}

// GetBatchFromHiddenmethods 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromHiddenmethods(hiddenmethodss []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`hiddenmethods` IN (?)", hiddenmethodss).Find(&results).Error

	return
}

// GetFromSupertypes 通过supertypes获取内容
func (obj *_AllTypeHistoryMgr) GetFromSupertypes(supertypes int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`supertypes` = ?", supertypes).Find(&results).Error

	return
}

// GetBatchFromSupertypes 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromSupertypes(supertypess []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`supertypes` IN (?)", supertypess).Find(&results).Error

	return
}

// GetFromSubtypes 通过subtypes获取内容
func (obj *_AllTypeHistoryMgr) GetFromSubtypes(subtypes int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`subtypes` = ?", subtypes).Find(&results).Error

	return
}

// GetBatchFromSubtypes 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromSubtypes(subtypess []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`subtypes` IN (?)", subtypess).Find(&results).Error

	return
}

// GetFromExterntype 通过externtype获取内容
func (obj *_AllTypeHistoryMgr) GetFromExterntype(externtype int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`externtype` = ?", externtype).Find(&results).Error

	return
}

// GetBatchFromExterntype 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromExterntype(externtypes []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`externtype` IN (?)", externtypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllTypeHistoryMgr) GetFromExternname(externname string) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromExternname(externnames []string) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromHelperclassname 通过helperclassname获取内容
func (obj *_AllTypeHistoryMgr) GetFromHelperclassname(helperclassname string) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`helperclassname` = ?", helperclassname).Find(&results).Error

	return
}

// GetBatchFromHelperclassname 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromHelperclassname(helperclassnames []string) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`helperclassname` IN (?)", helperclassnames).Find(&results).Error

	return
}

// GetFromLocalAttrs 通过local_attrs获取内容
func (obj *_AllTypeHistoryMgr) GetFromLocalAttrs(localAttrs int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`local_attrs` = ?", localAttrs).Find(&results).Error

	return
}

// GetBatchFromLocalAttrs 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromLocalAttrs(localAttrss []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`local_attrs` IN (?)", localAttrss).Find(&results).Error

	return
}

// GetFromLocalMethods 通过local_methods获取内容
func (obj *_AllTypeHistoryMgr) GetFromLocalMethods(localMethods int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`local_methods` = ?", localMethods).Find(&results).Error

	return
}

// GetBatchFromLocalMethods 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromLocalMethods(localMethodss []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`local_methods` IN (?)", localMethodss).Find(&results).Error

	return
}

// GetFromSupertypeid 通过supertypeid获取内容
func (obj *_AllTypeHistoryMgr) GetFromSupertypeid(supertypeid int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`supertypeid` = ?", supertypeid).Find(&results).Error

	return
}

// GetBatchFromSupertypeid 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromSupertypeid(supertypeids []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`supertypeid` IN (?)", supertypeids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容
func (obj *_AllTypeHistoryMgr) GetFromTypeName(typeName string) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromTypeName(typeNames []string) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllTypeHistoryMgr) GetFromPackageID(packageID int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllTypeHistoryMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTypeHistoryMgr) FetchByPrimaryKey(tenantID int64, typeID int64, schemaVersion int64) (result AllTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeHistory{}).Where("`tenant_id` = ? AND `type_id` = ? AND `schema_version` = ?", tenantID, typeID, schemaVersion).First(&result).Error

	return
}
