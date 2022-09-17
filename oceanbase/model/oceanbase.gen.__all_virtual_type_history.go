package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTypeHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualTypeHistoryMgr open func
func AllVirtualTypeHistoryMgr(db *gorm.DB) *_AllVirtualTypeHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTypeHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTypeHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_type_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTypeHistoryMgr) GetTableName() string {
	return "__all_virtual_type_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTypeHistoryMgr) Reset() *_AllVirtualTypeHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTypeHistoryMgr) Get() (result AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTypeHistoryMgr) Gets() (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTypeHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTypeHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllVirtualTypeHistoryMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTypeHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTypeHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTypeHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualTypeHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualTypeHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithTypecode typecode获取
func (obj *_AllVirtualTypeHistoryMgr) WithTypecode(typecode int64) Option {
	return optionFunc(func(o *options) { o.query["typecode"] = typecode })
}

// WithProperties properties获取
func (obj *_AllVirtualTypeHistoryMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithAttributes attributes获取
func (obj *_AllVirtualTypeHistoryMgr) WithAttributes(attributes int64) Option {
	return optionFunc(func(o *options) { o.query["attributes"] = attributes })
}

// WithMethods methods获取
func (obj *_AllVirtualTypeHistoryMgr) WithMethods(methods int64) Option {
	return optionFunc(func(o *options) { o.query["methods"] = methods })
}

// WithHiddenmethods hiddenmethods获取
func (obj *_AllVirtualTypeHistoryMgr) WithHiddenmethods(hiddenmethods int64) Option {
	return optionFunc(func(o *options) { o.query["hiddenmethods"] = hiddenmethods })
}

// WithSupertypes supertypes获取
func (obj *_AllVirtualTypeHistoryMgr) WithSupertypes(supertypes int64) Option {
	return optionFunc(func(o *options) { o.query["supertypes"] = supertypes })
}

// WithSubtypes subtypes获取
func (obj *_AllVirtualTypeHistoryMgr) WithSubtypes(subtypes int64) Option {
	return optionFunc(func(o *options) { o.query["subtypes"] = subtypes })
}

// WithExterntype externtype获取
func (obj *_AllVirtualTypeHistoryMgr) WithExterntype(externtype int64) Option {
	return optionFunc(func(o *options) { o.query["externtype"] = externtype })
}

// WithExternname externname获取
func (obj *_AllVirtualTypeHistoryMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithHelperclassname helperclassname获取
func (obj *_AllVirtualTypeHistoryMgr) WithHelperclassname(helperclassname string) Option {
	return optionFunc(func(o *options) { o.query["helperclassname"] = helperclassname })
}

// WithLocalAttrs local_attrs获取
func (obj *_AllVirtualTypeHistoryMgr) WithLocalAttrs(localAttrs int64) Option {
	return optionFunc(func(o *options) { o.query["local_attrs"] = localAttrs })
}

// WithLocalMethods local_methods获取
func (obj *_AllVirtualTypeHistoryMgr) WithLocalMethods(localMethods int64) Option {
	return optionFunc(func(o *options) { o.query["local_methods"] = localMethods })
}

// WithSupertypeid supertypeid获取
func (obj *_AllVirtualTypeHistoryMgr) WithSupertypeid(supertypeid int64) Option {
	return optionFunc(func(o *options) { o.query["supertypeid"] = supertypeid })
}

// WithTypeName type_name获取
func (obj *_AllVirtualTypeHistoryMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithPackageID package_id获取
func (obj *_AllVirtualTypeHistoryMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTypeHistoryMgr) GetByOption(opts ...Option) (result AllVirtualTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTypeHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTypeHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTypeHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where(options.query)
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
func (obj *_AllVirtualTypeHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromTypeID(typeID int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromTypecode 通过typecode获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromTypecode(typecode int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`typecode` = ?", typecode).Find(&results).Error

	return
}

// GetBatchFromTypecode 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromTypecode(typecodes []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`typecode` IN (?)", typecodes).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromProperties(properties int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromProperties(propertiess []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromAttributes 通过attributes获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromAttributes(attributes int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`attributes` = ?", attributes).Find(&results).Error

	return
}

// GetBatchFromAttributes 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromAttributes(attributess []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`attributes` IN (?)", attributess).Find(&results).Error

	return
}

// GetFromMethods 通过methods获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromMethods(methods int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`methods` = ?", methods).Find(&results).Error

	return
}

// GetBatchFromMethods 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromMethods(methodss []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`methods` IN (?)", methodss).Find(&results).Error

	return
}

// GetFromHiddenmethods 通过hiddenmethods获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromHiddenmethods(hiddenmethods int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`hiddenmethods` = ?", hiddenmethods).Find(&results).Error

	return
}

// GetBatchFromHiddenmethods 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromHiddenmethods(hiddenmethodss []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`hiddenmethods` IN (?)", hiddenmethodss).Find(&results).Error

	return
}

// GetFromSupertypes 通过supertypes获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromSupertypes(supertypes int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`supertypes` = ?", supertypes).Find(&results).Error

	return
}

// GetBatchFromSupertypes 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromSupertypes(supertypess []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`supertypes` IN (?)", supertypess).Find(&results).Error

	return
}

// GetFromSubtypes 通过subtypes获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromSubtypes(subtypes int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`subtypes` = ?", subtypes).Find(&results).Error

	return
}

// GetBatchFromSubtypes 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromSubtypes(subtypess []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`subtypes` IN (?)", subtypess).Find(&results).Error

	return
}

// GetFromExterntype 通过externtype获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromExterntype(externtype int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`externtype` = ?", externtype).Find(&results).Error

	return
}

// GetBatchFromExterntype 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromExterntype(externtypes []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`externtype` IN (?)", externtypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromExternname(externname string) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromExternname(externnames []string) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromHelperclassname 通过helperclassname获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromHelperclassname(helperclassname string) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`helperclassname` = ?", helperclassname).Find(&results).Error

	return
}

// GetBatchFromHelperclassname 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromHelperclassname(helperclassnames []string) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`helperclassname` IN (?)", helperclassnames).Find(&results).Error

	return
}

// GetFromLocalAttrs 通过local_attrs获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromLocalAttrs(localAttrs int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`local_attrs` = ?", localAttrs).Find(&results).Error

	return
}

// GetBatchFromLocalAttrs 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromLocalAttrs(localAttrss []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`local_attrs` IN (?)", localAttrss).Find(&results).Error

	return
}

// GetFromLocalMethods 通过local_methods获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromLocalMethods(localMethods int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`local_methods` = ?", localMethods).Find(&results).Error

	return
}

// GetBatchFromLocalMethods 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromLocalMethods(localMethodss []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`local_methods` IN (?)", localMethodss).Find(&results).Error

	return
}

// GetFromSupertypeid 通过supertypeid获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromSupertypeid(supertypeid int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`supertypeid` = ?", supertypeid).Find(&results).Error

	return
}

// GetBatchFromSupertypeid 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromSupertypeid(supertypeids []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`supertypeid` IN (?)", supertypeids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromTypeName(typeName string) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromTypeName(typeNames []string) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllVirtualTypeHistoryMgr) GetFromPackageID(packageID int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllVirtualTypeHistoryMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTypeHistoryMgr) FetchByPrimaryKey(tenantID int64, typeID int64, schemaVersion int64) (result AllVirtualTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeHistory{}).Where("`tenant_id` = ? AND `type_id` = ? AND `schema_version` = ?", tenantID, typeID, schemaVersion).First(&result).Error

	return
}
