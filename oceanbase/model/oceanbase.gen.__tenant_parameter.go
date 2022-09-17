package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TenantParameterMgr struct {
	*_BaseMgr
}

// TenantParameterMgr open func
func TenantParameterMgr(db *gorm.DB) *_TenantParameterMgr {
	if db == nil {
		panic(fmt.Errorf("TenantParameterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantParameterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_parameter"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantParameterMgr) GetTableName() string {
	return "__tenant_parameter"
}

// Reset 重置gorm会话
func (obj *_TenantParameterMgr) Reset() *_TenantParameterMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantParameterMgr) Get() (result TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantParameterMgr) Gets() (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantParameterMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_TenantParameterMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_TenantParameterMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithZone zone获取
func (obj *_TenantParameterMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrType svr_type获取
func (obj *_TenantParameterMgr) WithSvrType(svrType string) Option {
	return optionFunc(func(o *options) { o.query["svr_type"] = svrType })
}

// WithSvrIP svr_ip获取
func (obj *_TenantParameterMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_TenantParameterMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithName name获取
func (obj *_TenantParameterMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_TenantParameterMgr) WithDataType(dataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_TenantParameterMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_TenantParameterMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithSection section获取
func (obj *_TenantParameterMgr) WithSection(section string) Option {
	return optionFunc(func(o *options) { o.query["section"] = section })
}

// WithScope scope获取
func (obj *_TenantParameterMgr) WithScope(scope string) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithSource source获取
func (obj *_TenantParameterMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithEditLevel edit_level获取
func (obj *_TenantParameterMgr) WithEditLevel(editLevel string) Option {
	return optionFunc(func(o *options) { o.query["edit_level"] = editLevel })
}

// WithConfigVersion config_version获取
func (obj *_TenantParameterMgr) WithConfigVersion(configVersion int64) Option {
	return optionFunc(func(o *options) { o.query["config_version"] = configVersion })
}

// GetByOption 功能选项模式获取
func (obj *_TenantParameterMgr) GetByOption(opts ...Option) (result TenantParameter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantParameterMgr) GetByOptions(opts ...Option) (results []*TenantParameter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantParameterMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantParameter, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where(options.query)
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
func (obj *_TenantParameterMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_TenantParameterMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_TenantParameterMgr) GetFromGmtModified(gmtModified time.Time) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_TenantParameterMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_TenantParameterMgr) GetFromZone(zone string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_TenantParameterMgr) GetBatchFromZone(zones []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrType 通过svr_type获取内容
func (obj *_TenantParameterMgr) GetFromSvrType(svrType string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`svr_type` = ?", svrType).Find(&results).Error

	return
}

// GetBatchFromSvrType 批量查找
func (obj *_TenantParameterMgr) GetBatchFromSvrType(svrTypes []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`svr_type` IN (?)", svrTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_TenantParameterMgr) GetFromSvrIP(svrIP string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_TenantParameterMgr) GetBatchFromSvrIP(svrIPs []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_TenantParameterMgr) GetFromSvrPort(svrPort int64) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_TenantParameterMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TenantParameterMgr) GetFromName(name string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_TenantParameterMgr) GetBatchFromName(names []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_TenantParameterMgr) GetFromDataType(dataType string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_TenantParameterMgr) GetBatchFromDataType(dataTypes []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_TenantParameterMgr) GetFromValue(value string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_TenantParameterMgr) GetBatchFromValue(values []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_TenantParameterMgr) GetFromInfo(info string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_TenantParameterMgr) GetBatchFromInfo(infos []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromSection 通过section获取内容
func (obj *_TenantParameterMgr) GetFromSection(section string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`section` = ?", section).Find(&results).Error

	return
}

// GetBatchFromSection 批量查找
func (obj *_TenantParameterMgr) GetBatchFromSection(sections []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`section` IN (?)", sections).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容
func (obj *_TenantParameterMgr) GetFromScope(scope string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`scope` = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量查找
func (obj *_TenantParameterMgr) GetBatchFromScope(scopes []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`scope` IN (?)", scopes).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_TenantParameterMgr) GetFromSource(source string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_TenantParameterMgr) GetBatchFromSource(sources []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromEditLevel 通过edit_level获取内容
func (obj *_TenantParameterMgr) GetFromEditLevel(editLevel string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`edit_level` = ?", editLevel).Find(&results).Error

	return
}

// GetBatchFromEditLevel 批量查找
func (obj *_TenantParameterMgr) GetBatchFromEditLevel(editLevels []string) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`edit_level` IN (?)", editLevels).Find(&results).Error

	return
}

// GetFromConfigVersion 通过config_version获取内容
func (obj *_TenantParameterMgr) GetFromConfigVersion(configVersion int64) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`config_version` = ?", configVersion).Find(&results).Error

	return
}

// GetBatchFromConfigVersion 批量查找
func (obj *_TenantParameterMgr) GetBatchFromConfigVersion(configVersions []int64) (results []*TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`config_version` IN (?)", configVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantParameterMgr) FetchByPrimaryKey(zone string, svrType string, svrIP string, svrPort int64, name string) (result TenantParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantParameter{}).Where("`zone` = ? AND `svr_type` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `name` = ?", zone, svrType, svrIP, svrPort, name).First(&result).Error

	return
}
