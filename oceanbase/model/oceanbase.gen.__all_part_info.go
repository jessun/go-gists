package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllPartInfoMgr struct {
	*_BaseMgr
}

// AllPartInfoMgr open func
func AllPartInfoMgr(db *gorm.DB) *_AllPartInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllPartInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllPartInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_part_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllPartInfoMgr) GetTableName() string {
	return "__all_part_info"
}

// Reset 重置gorm会话
func (obj *_AllPartInfoMgr) Reset() *_AllPartInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllPartInfoMgr) Get() (result AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllPartInfoMgr) Gets() (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllPartInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllPartInfoMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllPartInfoMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllPartInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllPartInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartType part_type获取
func (obj *_AllPartInfoMgr) WithPartType(partType int64) Option {
	return optionFunc(func(o *options) { o.query["part_type"] = partType })
}

// WithSchemaVersion schema_version获取
func (obj *_AllPartInfoMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithPartNum part_num获取
func (obj *_AllPartInfoMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithPartSpace part_space获取
func (obj *_AllPartInfoMgr) WithPartSpace(partSpace int64) Option {
	return optionFunc(func(o *options) { o.query["part_space"] = partSpace })
}

// WithNewPartNum new_part_num获取
func (obj *_AllPartInfoMgr) WithNewPartNum(newPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["new_part_num"] = newPartNum })
}

// WithNewPartSpace new_part_space获取
func (obj *_AllPartInfoMgr) WithNewPartSpace(newPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["new_part_space"] = newPartSpace })
}

// WithSubPartType sub_part_type获取
func (obj *_AllPartInfoMgr) WithSubPartType(subPartType int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_type"] = subPartType })
}

// WithDefSubPartNum def_sub_part_num获取
func (obj *_AllPartInfoMgr) WithDefSubPartNum(defSubPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["def_sub_part_num"] = defSubPartNum })
}

// WithPartExpr part_expr获取
func (obj *_AllPartInfoMgr) WithPartExpr(partExpr string) Option {
	return optionFunc(func(o *options) { o.query["part_expr"] = partExpr })
}

// WithSubPartExpr sub_part_expr获取
func (obj *_AllPartInfoMgr) WithSubPartExpr(subPartExpr string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_expr"] = subPartExpr })
}

// WithPartInterval part_interval获取
func (obj *_AllPartInfoMgr) WithPartInterval(partInterval string) Option {
	return optionFunc(func(o *options) { o.query["part_interval"] = partInterval })
}

// WithIntervalStart interval_start获取
func (obj *_AllPartInfoMgr) WithIntervalStart(intervalStart string) Option {
	return optionFunc(func(o *options) { o.query["interval_start"] = intervalStart })
}

// WithNewPartInterval new_part_interval获取
func (obj *_AllPartInfoMgr) WithNewPartInterval(newPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["new_part_interval"] = newPartInterval })
}

// WithNewIntervalStart new_interval_start获取
func (obj *_AllPartInfoMgr) WithNewIntervalStart(newIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["new_interval_start"] = newIntervalStart })
}

// WithDefSubPartInterval def_sub_part_interval获取
func (obj *_AllPartInfoMgr) WithDefSubPartInterval(defSubPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["def_sub_part_interval"] = defSubPartInterval })
}

// WithDefSubIntervalStart def_sub_interval_start获取
func (obj *_AllPartInfoMgr) WithDefSubIntervalStart(defSubIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["def_sub_interval_start"] = defSubIntervalStart })
}

// WithNewDefSubPartInterval new_def_sub_part_interval获取
func (obj *_AllPartInfoMgr) WithNewDefSubPartInterval(newDefSubPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["new_def_sub_part_interval"] = newDefSubPartInterval })
}

// WithNewDefSubIntervalStart new_def_sub_interval_start获取
func (obj *_AllPartInfoMgr) WithNewDefSubIntervalStart(newDefSubIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["new_def_sub_interval_start"] = newDefSubIntervalStart })
}

// WithBlockSize block_size获取
func (obj *_AllPartInfoMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllPartInfoMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllPartInfoMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithSpare1 spare1获取
func (obj *_AllPartInfoMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllPartInfoMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllPartInfoMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// GetByOption 功能选项模式获取
func (obj *_AllPartInfoMgr) GetByOption(opts ...Option) (result AllPartInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllPartInfoMgr) GetByOptions(opts ...Option) (results []*AllPartInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllPartInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllPartInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where(options.query)
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
func (obj *_AllPartInfoMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllPartInfoMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllPartInfoMgr) GetFromTenantID(tenantID int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllPartInfoMgr) GetFromTableID(tableID int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartType 通过part_type获取内容
func (obj *_AllPartInfoMgr) GetFromPartType(partType int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_type` = ?", partType).Find(&results).Error

	return
}

// GetBatchFromPartType 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromPartType(partTypes []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_type` IN (?)", partTypes).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllPartInfoMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromPartNum 通过part_num获取内容
func (obj *_AllPartInfoMgr) GetFromPartNum(partNum int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_num` = ?", partNum).Find(&results).Error

	return
}

// GetBatchFromPartNum 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromPartNum(partNums []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_num` IN (?)", partNums).Find(&results).Error

	return
}

// GetFromPartSpace 通过part_space获取内容
func (obj *_AllPartInfoMgr) GetFromPartSpace(partSpace int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_space` = ?", partSpace).Find(&results).Error

	return
}

// GetBatchFromPartSpace 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromPartSpace(partSpaces []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_space` IN (?)", partSpaces).Find(&results).Error

	return
}

// GetFromNewPartNum 通过new_part_num获取内容
func (obj *_AllPartInfoMgr) GetFromNewPartNum(newPartNum int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_part_num` = ?", newPartNum).Find(&results).Error

	return
}

// GetBatchFromNewPartNum 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromNewPartNum(newPartNums []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_part_num` IN (?)", newPartNums).Find(&results).Error

	return
}

// GetFromNewPartSpace 通过new_part_space获取内容
func (obj *_AllPartInfoMgr) GetFromNewPartSpace(newPartSpace int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_part_space` = ?", newPartSpace).Find(&results).Error

	return
}

// GetBatchFromNewPartSpace 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromNewPartSpace(newPartSpaces []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_part_space` IN (?)", newPartSpaces).Find(&results).Error

	return
}

// GetFromSubPartType 通过sub_part_type获取内容
func (obj *_AllPartInfoMgr) GetFromSubPartType(subPartType int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`sub_part_type` = ?", subPartType).Find(&results).Error

	return
}

// GetBatchFromSubPartType 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromSubPartType(subPartTypes []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`sub_part_type` IN (?)", subPartTypes).Find(&results).Error

	return
}

// GetFromDefSubPartNum 通过def_sub_part_num获取内容
func (obj *_AllPartInfoMgr) GetFromDefSubPartNum(defSubPartNum int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`def_sub_part_num` = ?", defSubPartNum).Find(&results).Error

	return
}

// GetBatchFromDefSubPartNum 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromDefSubPartNum(defSubPartNums []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`def_sub_part_num` IN (?)", defSubPartNums).Find(&results).Error

	return
}

// GetFromPartExpr 通过part_expr获取内容
func (obj *_AllPartInfoMgr) GetFromPartExpr(partExpr string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_expr` = ?", partExpr).Find(&results).Error

	return
}

// GetBatchFromPartExpr 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromPartExpr(partExprs []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_expr` IN (?)", partExprs).Find(&results).Error

	return
}

// GetFromSubPartExpr 通过sub_part_expr获取内容
func (obj *_AllPartInfoMgr) GetFromSubPartExpr(subPartExpr string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`sub_part_expr` = ?", subPartExpr).Find(&results).Error

	return
}

// GetBatchFromSubPartExpr 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromSubPartExpr(subPartExprs []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`sub_part_expr` IN (?)", subPartExprs).Find(&results).Error

	return
}

// GetFromPartInterval 通过part_interval获取内容
func (obj *_AllPartInfoMgr) GetFromPartInterval(partInterval string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_interval` = ?", partInterval).Find(&results).Error

	return
}

// GetBatchFromPartInterval 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromPartInterval(partIntervals []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`part_interval` IN (?)", partIntervals).Find(&results).Error

	return
}

// GetFromIntervalStart 通过interval_start获取内容
func (obj *_AllPartInfoMgr) GetFromIntervalStart(intervalStart string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`interval_start` = ?", intervalStart).Find(&results).Error

	return
}

// GetBatchFromIntervalStart 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromIntervalStart(intervalStarts []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`interval_start` IN (?)", intervalStarts).Find(&results).Error

	return
}

// GetFromNewPartInterval 通过new_part_interval获取内容
func (obj *_AllPartInfoMgr) GetFromNewPartInterval(newPartInterval string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_part_interval` = ?", newPartInterval).Find(&results).Error

	return
}

// GetBatchFromNewPartInterval 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromNewPartInterval(newPartIntervals []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_part_interval` IN (?)", newPartIntervals).Find(&results).Error

	return
}

// GetFromNewIntervalStart 通过new_interval_start获取内容
func (obj *_AllPartInfoMgr) GetFromNewIntervalStart(newIntervalStart string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_interval_start` = ?", newIntervalStart).Find(&results).Error

	return
}

// GetBatchFromNewIntervalStart 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromNewIntervalStart(newIntervalStarts []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_interval_start` IN (?)", newIntervalStarts).Find(&results).Error

	return
}

// GetFromDefSubPartInterval 通过def_sub_part_interval获取内容
func (obj *_AllPartInfoMgr) GetFromDefSubPartInterval(defSubPartInterval string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`def_sub_part_interval` = ?", defSubPartInterval).Find(&results).Error

	return
}

// GetBatchFromDefSubPartInterval 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromDefSubPartInterval(defSubPartIntervals []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`def_sub_part_interval` IN (?)", defSubPartIntervals).Find(&results).Error

	return
}

// GetFromDefSubIntervalStart 通过def_sub_interval_start获取内容
func (obj *_AllPartInfoMgr) GetFromDefSubIntervalStart(defSubIntervalStart string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`def_sub_interval_start` = ?", defSubIntervalStart).Find(&results).Error

	return
}

// GetBatchFromDefSubIntervalStart 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromDefSubIntervalStart(defSubIntervalStarts []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`def_sub_interval_start` IN (?)", defSubIntervalStarts).Find(&results).Error

	return
}

// GetFromNewDefSubPartInterval 通过new_def_sub_part_interval获取内容
func (obj *_AllPartInfoMgr) GetFromNewDefSubPartInterval(newDefSubPartInterval string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_def_sub_part_interval` = ?", newDefSubPartInterval).Find(&results).Error

	return
}

// GetBatchFromNewDefSubPartInterval 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromNewDefSubPartInterval(newDefSubPartIntervals []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_def_sub_part_interval` IN (?)", newDefSubPartIntervals).Find(&results).Error

	return
}

// GetFromNewDefSubIntervalStart 通过new_def_sub_interval_start获取内容
func (obj *_AllPartInfoMgr) GetFromNewDefSubIntervalStart(newDefSubIntervalStart string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_def_sub_interval_start` = ?", newDefSubIntervalStart).Find(&results).Error

	return
}

// GetBatchFromNewDefSubIntervalStart 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromNewDefSubIntervalStart(newDefSubIntervalStarts []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`new_def_sub_interval_start` IN (?)", newDefSubIntervalStarts).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllPartInfoMgr) GetFromBlockSize(blockSize int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllPartInfoMgr) GetFromReplicaNum(replicaNum int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllPartInfoMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllPartInfoMgr) GetFromSpare1(spare1 int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllPartInfoMgr) GetFromSpare2(spare2 int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllPartInfoMgr) GetFromSpare3(spare3 string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllPartInfoMgr) GetBatchFromSpare3(spare3s []string) (results []*AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllPartInfoMgr) FetchByPrimaryKey(tenantID int64, tableID int64) (result AllPartInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartInfo{}).Where("`tenant_id` = ? AND `table_id` = ?", tenantID, tableID).First(&result).Error

	return
}
