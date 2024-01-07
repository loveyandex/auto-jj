package model


import "time"

type JobsResp struct {
	TraceID    any    `json:"traceId,omitempty"`
	IsSuccess  bool   `json:"isSuccess,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       Data   `json:"data,omitempty"`
}
type UserJobPostInfo struct {
	MatchingScore          int  `json:"matchingScore,omitempty"`
	IsSuitableForCandidate bool `json:"isSuitableForCandidate,omitempty"`
	IsBookmarked           bool `json:"isBookmarked,omitempty"`
	IsApplied              bool `json:"isApplied,omitempty"`
	IsCanceledApply        bool `json:"isCanceledApply,omitempty"`
	IsAppliedAsVip         bool `json:"isAppliedAsVip,omitempty"`
}
type Properties struct {
	IsInternship                   bool `json:"isInternship,omitempty"`
	IsRemote                       bool `json:"isRemote,omitempty"`
	IsUrgent                       bool `json:"isUrgent,omitempty"`
	RequiredRelatedExperienceYears int  `json:"requiredRelatedExperienceYears,omitempty"`
	SuitableForDisabled            bool `json:"suitableForDisabled,omitempty"`
	SalaryCanBeShown               bool `json:"salaryCanBeShown,omitempty"`
	IsCheckingResumes              bool `json:"isCheckingResumes,omitempty"`
	TypeID                         int  `json:"typeId,omitempty"`
	LinkOutAddress                 any  `json:"linkOutAddress,omitempty"`
	IsBlueCollarJob                bool `json:"isBlueCollarJob,omitempty"`
}
type Company struct {
	ID                   int     `json:"id,omitempty"`
	NameFa               string  `json:"nameFa,omitempty"`
	NameEn               string  `json:"nameEn,omitempty"`
	HasPicture           bool    `json:"hasPicture,omitempty"`
	LogoFileID           int     `json:"logoFileId,omitempty"`
	LogoURL              string  `json:"logoUrl,omitempty"`
	PageURL              string  `json:"pageUrl,omitempty"`
	IsEmployerResponsive bool    `json:"isEmployerResponsive,omitempty"`
	IsFamous             bool    `json:"isFamous,omitempty"`
	CompanyScore         float64 `json:"companyScore,omitempty"`
}
type Country struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type Province struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type City struct {
	CitySize        int      `json:"citySize,omitempty"`
	CitySizeGroupID int      `json:"citySizeGroupId,omitempty"`
	Province        Province `json:"province,omitempty"`
	ID              int      `json:"id,omitempty"`
	Title           any      `json:"title,omitempty"`
	TitleFa         string   `json:"titleFa,omitempty"`
	TitleEn         string   `json:"titleEn,omitempty"`
}
type RegionGroup struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type Region struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type Location struct {
	Country     Country     `json:"country,omitempty"`
	Province    Province    `json:"province,omitempty"`
	City        City        `json:"city,omitempty"`
	RegionGroup RegionGroup `json:"regionGroup,omitempty"`
	Region      Region      `json:"region,omitempty"`
}
type JobCategories struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type Benefits struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type WorkType struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type SeniorityLevel struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type Industry struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type Gender struct {
	ID      int    `json:"id,omitempty"`
	Title   any    `json:"title,omitempty"`
	TitleFa string `json:"titleFa,omitempty"`
	TitleEn string `json:"titleEn,omitempty"`
}
type FirstActivationTime struct {
	BeautifyFa string    `json:"beautifyFa,omitempty"`
	BeautifyEn string    `json:"beautifyEn,omitempty"`
	Date       time.Time `json:"date,omitempty"`
}
type ActivationTime struct {
	PassedDays int       `json:"passedDays,omitempty"`
	BeautifyFa string    `json:"beautifyFa,omitempty"`
	BeautifyEn string    `json:"beautifyEn,omitempty"`
	Date       time.Time `json:"date,omitempty"`
}
type ExpireTime struct {
	DaysLeftUntil int       `json:"daysLeftUntil,omitempty"`
	IsExpired     bool      `json:"isExpired,omitempty"`
	BeautifyFa    string    `json:"beautifyFa,omitempty"`
	BeautifyEn    string    `json:"beautifyEn,omitempty"`
	Date          time.Time `json:"date,omitempty"`
}
type JobPosts struct {
	Score                     float64             `json:"score,omitempty"`
	ItemIndex                 int                 `json:"itemIndex,omitempty"`
	SearchPositionCoefficient float64             `json:"searchPositionCoefficient,omitempty"`
	ID                        int                 `json:"id,omitempty"`
	Title                     string              `json:"title,omitempty"`
	IsPersian                 bool                `json:"isPersian,omitempty"`
	UserJobPostInfo           UserJobPostInfo     `json:"userJobPostInfo,omitempty"`
	Properties                Properties          `json:"properties,omitempty"`
	Company                   Company             `json:"company,omitempty"`
	Location                  Location            `json:"location,omitempty"`
	JobCategories             []JobCategories     `json:"jobCategories,omitempty"`
	Benefits                  []Benefits          `json:"benefits,omitempty"`
	WorkType                  WorkType            `json:"workType,omitempty"`
	SeniorityLevel            SeniorityLevel      `json:"seniorityLevel,omitempty"`
	Salary                    any                 `json:"salary,omitempty"`
	Industry                  Industry            `json:"industry,omitempty"`
	Gender                    Gender              `json:"gender,omitempty"`
	Labels                    []any               `json:"labels,omitempty"`
	FirstActivationTime       FirstActivationTime `json:"firstActivationTime,omitempty"`
	ActivationTime            ActivationTime      `json:"activationTime,omitempty"`
	ExpireTime                ExpireTime          `json:"expireTime,omitempty"`
}
type Filters struct {
	Keyword            any  `json:"keyword,omitempty"`
	JobCategory        any  `json:"jobCategory,omitempty"`
	LocationWrapper    any  `json:"locationWrapper,omitempty"`
	WorkExperiences    any  `json:"workExperiences,omitempty"`
	Salaries           any  `json:"salaries,omitempty"`
	WorkTypes          any  `json:"workTypes,omitempty"`
	SeniorityLevels    any  `json:"seniorityLevels,omitempty"`
	JobPostPublishTime any  `json:"jobPostPublishTime,omitempty"`
	Industries         any  `json:"industries,omitempty"`
	Company            any  `json:"company,omitempty"`
	Benefits           any  `json:"benefits,omitempty"`
	Remote             bool `json:"remote,omitempty"`
}
type Data struct {
	CurrentPage  int        `json:"currentPage,omitempty"`
	PageSize     int        `json:"pageSize,omitempty"`
	JobPosts     []JobPosts `json:"jobPosts,omitempty"`
	Filters      Filters    `json:"filters,omitempty"`
	SearchID     string     `json:"searchId,omitempty"`
	JobPostCount int        `json:"jobPostCount,omitempty"`
}