package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSeedParameterMgr struct {
	*_BaseMgr
}

// AllSeedParameterMgr open func
func AllSeedParameterMgr(db *gorm.DB) *_AllSeedParameterMgr {
	if db == nil {
		panic(fmt.Errorf("AllSeedParameterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSeedParameterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_seed_parameter"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSeedParameterMgr) GetTableName() string {
	return "__all_seed_parameter"
}

// Reset 重置gorm会话
func (obj *_AllSeedParameterMgr) Reset() *_AllSeedParameterMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSeedParameterMgr) Get() (result AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSeedParameterMgr) Gets() (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSeedParameterMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSeedParameterMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSeedParameterMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithZone zone获取
func (obj *_AllSeedParameterMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrType svr_type获取
func (obj *_AllSeedParameterMgr) WithSvrType(svrType string) Option {
	return optionFunc(func(o *options) { o.query["svr_type"] = svrType })
}

// WithSvrIP svr_ip获取
func (obj *_AllSeedParameterMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllSeedParameterMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithName name获取
func (obj *_AllSeedParameterMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_AllSeedParameterMgr) WithDataType(dataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllSeedParameterMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllSeedParameterMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithSection section获取
func (obj *_AllSeedParameterMgr) WithSection(section string) Option {
	return optionFunc(func(o *options) { o.query["section"] = section })
}

// WithScope scope获取
func (obj *_AllSeedParameterMgr) WithScope(scope string) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithSource source获取
func (obj *_AllSeedParameterMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithEditLevel edit_level获取
func (obj *_AllSeedParameterMgr) WithEditLevel(editLevel string) Option {
	return optionFunc(func(o *options) { o.query["edit_level"] = editLevel })
}

// WithConfigVersion config_version获取
func (obj *_AllSeedParameterMgr) WithConfigVersion(configVersion int64) Option {
	return optionFunc(func(o *options) { o.query["config_version"] = configVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllSeedParameterMgr) GetByOption(opts ...Option) (result AllSeedParameter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSeedParameterMgr) GetByOptions(opts ...Option) (results []*AllSeedParameter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSeedParameterMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSeedParameter, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where(options.query)
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
func (obj *_AllSeedParameterMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSeedParameterMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllSeedParameterMgr) GetFromZone(zone string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromZone(zones []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrType 通过svr_type获取内容
func (obj *_AllSeedParameterMgr) GetFromSvrType(svrType string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`svr_type` = ?", svrType).Find(&results).Error

	return
}

// GetBatchFromSvrType 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromSvrType(svrTypes []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`svr_type` IN (?)", svrTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllSeedParameterMgr) GetFromSvrIP(svrIP string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllSeedParameterMgr) GetFromSvrPort(svrPort int64) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllSeedParameterMgr) GetFromName(name string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromName(names []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllSeedParameterMgr) GetFromDataType(dataType string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromDataType(dataTypes []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllSeedParameterMgr) GetFromValue(value string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromValue(values []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllSeedParameterMgr) GetFromInfo(info string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromInfo(infos []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromSection 通过section获取内容
func (obj *_AllSeedParameterMgr) GetFromSection(section string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`section` = ?", section).Find(&results).Error

	return
}

// GetBatchFromSection 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromSection(sections []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`section` IN (?)", sections).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容
func (obj *_AllSeedParameterMgr) GetFromScope(scope string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`scope` = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromScope(scopes []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`scope` IN (?)", scopes).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllSeedParameterMgr) GetFromSource(source string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromSource(sources []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromEditLevel 通过edit_level获取内容
func (obj *_AllSeedParameterMgr) GetFromEditLevel(editLevel string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`edit_level` = ?", editLevel).Find(&results).Error

	return
}

// GetBatchFromEditLevel 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromEditLevel(editLevels []string) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`edit_level` IN (?)", editLevels).Find(&results).Error

	return
}

// GetFromConfigVersion 通过config_version获取内容
func (obj *_AllSeedParameterMgr) GetFromConfigVersion(configVersion int64) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`config_version` = ?", configVersion).Find(&results).Error

	return
}

// GetBatchFromConfigVersion 批量查找
func (obj *_AllSeedParameterMgr) GetBatchFromConfigVersion(configVersions []int64) (results []*AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`config_version` IN (?)", configVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSeedParameterMgr) FetchByPrimaryKey(zone string, svrType string, svrIP string, svrPort int64, name string) (result AllSeedParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSeedParameter{}).Where("`zone` = ? AND `svr_type` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `name` = ?", zone, svrType, svrIP, svrPort, name).First(&result).Error

	return
}
