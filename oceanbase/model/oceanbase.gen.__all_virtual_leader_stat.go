package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualLeaderStatMgr struct {
	*_BaseMgr
}

// AllVirtualLeaderStatMgr open func
func AllVirtualLeaderStatMgr(db *gorm.DB) *_AllVirtualLeaderStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualLeaderStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualLeaderStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_leader_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualLeaderStatMgr) GetTableName() string {
	return "__all_virtual_leader_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualLeaderStatMgr) Reset() *_AllVirtualLeaderStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualLeaderStatMgr) Get() (result AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualLeaderStatMgr) Gets() (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualLeaderStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualLeaderStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllVirtualLeaderStatMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualLeaderStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualLeaderStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualLeaderStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllVirtualLeaderStatMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithRegion region获取
func (obj *_AllVirtualLeaderStatMgr) WithRegion(region string) Option {
	return optionFunc(func(o *options) { o.query["region"] = region })
}

// WithRegionScore region_score获取
func (obj *_AllVirtualLeaderStatMgr) WithRegionScore(regionScore int64) Option {
	return optionFunc(func(o *options) { o.query["region_score"] = regionScore })
}

// WithNotMerging not_merging获取
func (obj *_AllVirtualLeaderStatMgr) WithNotMerging(notMerging int64) Option {
	return optionFunc(func(o *options) { o.query["not_merging"] = notMerging })
}

// WithCandidateCount candidate_count获取
func (obj *_AllVirtualLeaderStatMgr) WithCandidateCount(candidateCount int64) Option {
	return optionFunc(func(o *options) { o.query["candidate_count"] = candidateCount })
}

// WithIsCandidate is_candidate获取
func (obj *_AllVirtualLeaderStatMgr) WithIsCandidate(isCandidate int64) Option {
	return optionFunc(func(o *options) { o.query["is_candidate"] = isCandidate })
}

// WithMigrateOutOrTransformCount migrate_out_or_transform_count获取
func (obj *_AllVirtualLeaderStatMgr) WithMigrateOutOrTransformCount(migrateOutOrTransformCount int64) Option {
	return optionFunc(func(o *options) { o.query["migrate_out_or_transform_count"] = migrateOutOrTransformCount })
}

// WithInNormalUnitCount in_normal_unit_count获取
func (obj *_AllVirtualLeaderStatMgr) WithInNormalUnitCount(inNormalUnitCount int64) Option {
	return optionFunc(func(o *options) { o.query["in_normal_unit_count"] = inNormalUnitCount })
}

// WithZone zone获取
func (obj *_AllVirtualLeaderStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithZoneScore zone_score获取
func (obj *_AllVirtualLeaderStatMgr) WithZoneScore(zoneScore int64) Option {
	return optionFunc(func(o *options) { o.query["zone_score"] = zoneScore })
}

// WithOriginalLeaderCount original_leader_count获取
func (obj *_AllVirtualLeaderStatMgr) WithOriginalLeaderCount(originalLeaderCount int64) Option {
	return optionFunc(func(o *options) { o.query["original_leader_count"] = originalLeaderCount })
}

// WithRandomScore random_score获取
func (obj *_AllVirtualLeaderStatMgr) WithRandomScore(randomScore int64) Option {
	return optionFunc(func(o *options) { o.query["random_score"] = randomScore })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualLeaderStatMgr) GetByOption(opts ...Option) (result AllVirtualLeaderStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualLeaderStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualLeaderStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualLeaderStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualLeaderStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where(options.query)
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
func (obj *_AllVirtualLeaderStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromPrimaryZone(primaryZone string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromRegion 通过region获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromRegion(region string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`region` = ?", region).Find(&results).Error

	return
}

// GetBatchFromRegion 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromRegion(regions []string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`region` IN (?)", regions).Find(&results).Error

	return
}

// GetFromRegionScore 通过region_score获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromRegionScore(regionScore int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`region_score` = ?", regionScore).Find(&results).Error

	return
}

// GetBatchFromRegionScore 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromRegionScore(regionScores []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`region_score` IN (?)", regionScores).Find(&results).Error

	return
}

// GetFromNotMerging 通过not_merging获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromNotMerging(notMerging int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`not_merging` = ?", notMerging).Find(&results).Error

	return
}

// GetBatchFromNotMerging 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromNotMerging(notMergings []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`not_merging` IN (?)", notMergings).Find(&results).Error

	return
}

// GetFromCandidateCount 通过candidate_count获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromCandidateCount(candidateCount int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`candidate_count` = ?", candidateCount).Find(&results).Error

	return
}

// GetBatchFromCandidateCount 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromCandidateCount(candidateCounts []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`candidate_count` IN (?)", candidateCounts).Find(&results).Error

	return
}

// GetFromIsCandidate 通过is_candidate获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromIsCandidate(isCandidate int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`is_candidate` = ?", isCandidate).Find(&results).Error

	return
}

// GetBatchFromIsCandidate 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromIsCandidate(isCandidates []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`is_candidate` IN (?)", isCandidates).Find(&results).Error

	return
}

// GetFromMigrateOutOrTransformCount 通过migrate_out_or_transform_count获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromMigrateOutOrTransformCount(migrateOutOrTransformCount int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`migrate_out_or_transform_count` = ?", migrateOutOrTransformCount).Find(&results).Error

	return
}

// GetBatchFromMigrateOutOrTransformCount 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromMigrateOutOrTransformCount(migrateOutOrTransformCounts []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`migrate_out_or_transform_count` IN (?)", migrateOutOrTransformCounts).Find(&results).Error

	return
}

// GetFromInNormalUnitCount 通过in_normal_unit_count获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromInNormalUnitCount(inNormalUnitCount int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`in_normal_unit_count` = ?", inNormalUnitCount).Find(&results).Error

	return
}

// GetBatchFromInNormalUnitCount 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromInNormalUnitCount(inNormalUnitCounts []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`in_normal_unit_count` IN (?)", inNormalUnitCounts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromZone(zone string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromZoneScore 通过zone_score获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromZoneScore(zoneScore int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`zone_score` = ?", zoneScore).Find(&results).Error

	return
}

// GetBatchFromZoneScore 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromZoneScore(zoneScores []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`zone_score` IN (?)", zoneScores).Find(&results).Error

	return
}

// GetFromOriginalLeaderCount 通过original_leader_count获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromOriginalLeaderCount(originalLeaderCount int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`original_leader_count` = ?", originalLeaderCount).Find(&results).Error

	return
}

// GetBatchFromOriginalLeaderCount 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromOriginalLeaderCount(originalLeaderCounts []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`original_leader_count` IN (?)", originalLeaderCounts).Find(&results).Error

	return
}

// GetFromRandomScore 通过random_score获取内容
func (obj *_AllVirtualLeaderStatMgr) GetFromRandomScore(randomScore int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`random_score` = ?", randomScore).Find(&results).Error

	return
}

// GetBatchFromRandomScore 批量查找
func (obj *_AllVirtualLeaderStatMgr) GetBatchFromRandomScore(randomScores []int64) (results []*AllVirtualLeaderStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLeaderStat{}).Where("`random_score` IN (?)", randomScores).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
