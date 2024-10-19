// ---------------------------------------------------------------------------------------------------------------------
// Part of Emby's Go SDK 4.8.9.0
// https://github.com/MediaBrowser/Emby.SDK
// ---------------------------------------------------------------------------------------------------------------------

package api

import "time"

type (
	ConnectUserLinkType      string
	DayOfWeek                string
	DrawingImageOrientation  string
	DynamicDayOfWeek         string
	ExtendedVideoTypes       string
	ExtendedVideoSubTypes    string
	ImageType                string
	LiveTvTimerType          string
	LocationType             string
	MarkerType               string
	MediaProtocol            string
	MediaSourceType          string
	MediaStreamType          string
	MetadataFields           string
	PersonType               string
	PlayMethod               string
	RepeatMode               string
	SegmentSkipMode          string
	SubtitleDeliveryMethod   string
	SubtitleLocationType     string
	SubtitlePlaybackMode     string
	SyncJobItemStatus        string
	TranscodeReason          string
	TransportStreamTimestamp string
	UnratedItem              string
	UserItemShareLevel       string
	Video3DFormat            string
	TranscodingVpStepTypes   string
)

type TupleDoubleDouble struct {
	Item1 float64 `json:"Item1,omitempty"`
	Item2 float64 `json:"Item2,omitempty"`
}

const (
	AUDIO_MediaStreamType MediaStreamType = "Audio"
	VIDEO_MediaStreamType MediaStreamType = "Video"
)

const (
	ACTOR_PersonType    PersonType = "Actor"
	DIRECTOR_PersonType PersonType = "Director"
)

const (
	PRIMARY_ImageType ImageType = "Primary"
)

// A

type AccessSchedule struct {
	DayOfWeek *DynamicDayOfWeek `json:"DayOfWeek,omitempty"`
	StartHour float64           `json:"StartHour,omitempty"`
	EndHour   float64           `json:"EndHour,omitempty"`
}

type AuthenticationResult struct {
	User        *UserDto     `json:"User,omitempty"`
	SessionInfo *SessionInfo `json:"SessionInfo,omitempty"`
	AccessToken string       `json:"AccessToken,omitempty"`
	ServerId    string       `json:"ServerId,omitempty"`
}

// B

type BaseItemDto struct {
	Name                         string                   `json:"Name,omitempty"`
	OriginalTitle                string                   `json:"OriginalTitle,omitempty"`
	ServerId                     string                   `json:"ServerId,omitempty"`
	Id                           string                   `json:"Id,omitempty"`
	Guid                         string                   `json:"Guid,omitempty"`
	Etag                         string                   `json:"Etag,omitempty"`
	Prefix                       string                   `json:"Prefix,omitempty"`
	PlaylistItemId               string                   `json:"PlaylistItemId,omitempty"`
	DateCreated                  time.Time                `json:"DateCreated,omitempty"`
	ExtraType                    string                   `json:"ExtraType,omitempty"`
	SortIndexNumber              int32                    `json:"SortIndexNumber,omitempty"`
	SortParentIndexNumber        int32                    `json:"SortParentIndexNumber,omitempty"`
	CanDelete                    bool                     `json:"CanDelete,omitempty"`
	CanDownload                  bool                     `json:"CanDownload,omitempty"`
	CanEditItems                 bool                     `json:"CanEditItems,omitempty"`
	SupportsResume               bool                     `json:"SupportsResume,omitempty"`
	PresentationUniqueKey        string                   `json:"PresentationUniqueKey,omitempty"`
	PreferredMetadataLanguage    string                   `json:"PreferredMetadataLanguage,omitempty"`
	PreferredMetadataCountryCode string                   `json:"PreferredMetadataCountryCode,omitempty"`
	SupportsSync                 bool                     `json:"SupportsSync,omitempty"`
	SyncStatus                   *SyncJobItemStatus       `json:"SyncStatus,omitempty"`
	CanManageAccess              bool                     `json:"CanManageAccess,omitempty"`
	CanLeaveContent              bool                     `json:"CanLeaveContent,omitempty"`
	CanMakePublic                bool                     `json:"CanMakePublic,omitempty"`
	Container                    string                   `json:"Container,omitempty"`
	SortName                     string                   `json:"SortName,omitempty"`
	ForcedSortName               string                   `json:"ForcedSortName,omitempty"`
	Video3DFormat                *Video3DFormat           `json:"Video3DFormat,omitempty"`
	PremiereDate                 time.Time                `json:"PremiereDate,omitempty"`
	ExternalUrls                 []ExternalUrl            `json:"ExternalUrls,omitempty"`
	MediaSources                 []MediaSourceInfo        `json:"MediaSources,omitempty"`
	CriticRating                 float32                  `json:"CriticRating,omitempty"`
	GameSystemId                 int64                    `json:"GameSystemId,omitempty"`
	AsSeries                     bool                     `json:"AsSeries,omitempty"`
	GameSystem                   string                   `json:"GameSystem,omitempty"`
	ProductionLocations          []string                 `json:"ProductionLocations,omitempty"`
	Path                         string                   `json:"Path,omitempty"`
	OfficialRating               string                   `json:"OfficialRating,omitempty"`
	CustomRating                 string                   `json:"CustomRating,omitempty"`
	ChannelId                    string                   `json:"ChannelId,omitempty"`
	ChannelName                  string                   `json:"ChannelName,omitempty"`
	Overview                     string                   `json:"Overview,omitempty"`
	Taglines                     []string                 `json:"Taglines,omitempty"`
	Genres                       []string                 `json:"Genres,omitempty"`
	CommunityRating              float32                  `json:"CommunityRating,omitempty"`
	RunTimeTicks                 int64                    `json:"RunTimeTicks,omitempty"`
	Size                         int64                    `json:"Size,omitempty"`
	FileName                     string                   `json:"FileName,omitempty"`
	Bitrate                      int32                    `json:"Bitrate,omitempty"`
	ProductionYear               int32                    `json:"ProductionYear,omitempty"`
	Number                       string                   `json:"Number,omitempty"`
	ChannelNumber                string                   `json:"ChannelNumber,omitempty"`
	IndexNumber                  int32                    `json:"IndexNumber,omitempty"`
	IndexNumberEnd               int32                    `json:"IndexNumberEnd,omitempty"`
	ParentIndexNumber            int32                    `json:"ParentIndexNumber,omitempty"`
	RemoteTrailers               []MediaUrl               `json:"RemoteTrailers,omitempty"`
	ProviderIds                  *map[string]string       `json:"ProviderIds,omitempty"`
	IsFolder                     bool                     `json:"IsFolder,omitempty"`
	ParentId                     string                   `json:"ParentId,omitempty"`
	Type_                        string                   `json:"Type,omitempty"`
	People                       []BaseItemPerson         `json:"People,omitempty"`
	Studios                      []NameLongIdPair         `json:"Studios,omitempty"`
	GenreItems                   []NameLongIdPair         `json:"GenreItems,omitempty"`
	TagItems                     []NameLongIdPair         `json:"TagItems,omitempty"`
	ParentLogoItemId             string                   `json:"ParentLogoItemId,omitempty"`
	ParentBackdropItemId         string                   `json:"ParentBackdropItemId,omitempty"`
	ParentBackdropImageTags      []string                 `json:"ParentBackdropImageTags,omitempty"`
	LocalTrailerCount            int32                    `json:"LocalTrailerCount,omitempty"`
	UserData                     *UserItemDataDto         `json:"UserData,omitempty"`
	RecursiveItemCount           int32                    `json:"RecursiveItemCount,omitempty"`
	ChildCount                   int32                    `json:"ChildCount,omitempty"`
	SeriesName                   string                   `json:"SeriesName,omitempty"`
	SeriesId                     string                   `json:"SeriesId,omitempty"`
	SeasonId                     string                   `json:"SeasonId,omitempty"`
	SpecialFeatureCount          int32                    `json:"SpecialFeatureCount,omitempty"`
	DisplayPreferencesId         string                   `json:"DisplayPreferencesId,omitempty"`
	Status                       string                   `json:"Status,omitempty"`
	AirDays                      []DayOfWeek              `json:"AirDays,omitempty"`
	Tags                         []string                 `json:"Tags,omitempty"`
	PrimaryImageAspectRatio      float64                  `json:"PrimaryImageAspectRatio,omitempty"`
	Artists                      []string                 `json:"Artists,omitempty"`
	ArtistItems                  []NameIdPair             `json:"ArtistItems,omitempty"`
	Composers                    []NameIdPair             `json:"Composers,omitempty"`
	Album                        string                   `json:"Album,omitempty"`
	CollectionType               string                   `json:"CollectionType,omitempty"`
	DisplayOrder                 string                   `json:"DisplayOrder,omitempty"`
	AlbumId                      string                   `json:"AlbumId,omitempty"`
	AlbumPrimaryImageTag         string                   `json:"AlbumPrimaryImageTag,omitempty"`
	SeriesPrimaryImageTag        string                   `json:"SeriesPrimaryImageTag,omitempty"`
	AlbumArtist                  string                   `json:"AlbumArtist,omitempty"`
	AlbumArtists                 []NameIdPair             `json:"AlbumArtists,omitempty"`
	SeasonName                   string                   `json:"SeasonName,omitempty"`
	MediaStreams                 []MediaStream            `json:"MediaStreams,omitempty"`
	PartCount                    int32                    `json:"PartCount,omitempty"`
	ImageTags                    *map[string]string       `json:"ImageTags,omitempty"`
	BackdropImageTags            []string                 `json:"BackdropImageTags,omitempty"`
	ParentLogoImageTag           string                   `json:"ParentLogoImageTag,omitempty"`
	SeriesStudio                 string                   `json:"SeriesStudio,omitempty"`
	PrimaryImageItemId           string                   `json:"PrimaryImageItemId,omitempty"`
	PrimaryImageTag              string                   `json:"PrimaryImageTag,omitempty"`
	ParentThumbItemId            string                   `json:"ParentThumbItemId,omitempty"`
	ParentThumbImageTag          string                   `json:"ParentThumbImageTag,omitempty"`
	Chapters                     []ChapterInfo            `json:"Chapters,omitempty"`
	LocationType                 *LocationType            `json:"LocationType,omitempty"`
	MediaType                    string                   `json:"MediaType,omitempty"`
	EndDate                      time.Time                `json:"EndDate,omitempty"`
	LockedFields                 []MetadataFields         `json:"LockedFields,omitempty"`
	LockData                     bool                     `json:"LockData,omitempty"`
	Width                        int32                    `json:"Width,omitempty"`
	Height                       int32                    `json:"Height,omitempty"`
	CameraMake                   string                   `json:"CameraMake,omitempty"`
	CameraModel                  string                   `json:"CameraModel,omitempty"`
	Software                     string                   `json:"Software,omitempty"`
	ExposureTime                 float64                  `json:"ExposureTime,omitempty"`
	FocalLength                  float64                  `json:"FocalLength,omitempty"`
	ImageOrientation             *DrawingImageOrientation `json:"ImageOrientation,omitempty"`
	Aperture                     float64                  `json:"Aperture,omitempty"`
	ShutterSpeed                 float64                  `json:"ShutterSpeed,omitempty"`
	Latitude                     float64                  `json:"Latitude,omitempty"`
	Longitude                    float64                  `json:"Longitude,omitempty"`
	Altitude                     float64                  `json:"Altitude,omitempty"`
	IsoSpeedRating               int32                    `json:"IsoSpeedRating,omitempty"`
	SeriesTimerId                string                   `json:"SeriesTimerId,omitempty"`
	ChannelPrimaryImageTag       string                   `json:"ChannelPrimaryImageTag,omitempty"`
	StartDate                    time.Time                `json:"StartDate,omitempty"`
	CompletionPercentage         float64                  `json:"CompletionPercentage,omitempty"`
	IsRepeat                     bool                     `json:"IsRepeat,omitempty"`
	IsNew                        bool                     `json:"IsNew,omitempty"`
	EpisodeTitle                 string                   `json:"EpisodeTitle,omitempty"`
	IsMovie                      bool                     `json:"IsMovie,omitempty"`
	IsSports                     bool                     `json:"IsSports,omitempty"`
	IsSeries                     bool                     `json:"IsSeries,omitempty"`
	IsLive                       bool                     `json:"IsLive,omitempty"`
	IsNews                       bool                     `json:"IsNews,omitempty"`
	IsKids                       bool                     `json:"IsKids,omitempty"`
	IsPremiere                   bool                     `json:"IsPremiere,omitempty"`
	TimerType                    *LiveTvTimerType         `json:"TimerType,omitempty"`
	Disabled                     bool                     `json:"Disabled,omitempty"`
	ManagementId                 string                   `json:"ManagementId,omitempty"`
	TimerId                      string                   `json:"TimerId,omitempty"`
	CurrentProgram               *BaseItemDto             `json:"CurrentProgram,omitempty"`
	MovieCount                   int32                    `json:"MovieCount,omitempty"`
	SeriesCount                  int32                    `json:"SeriesCount,omitempty"`
	AlbumCount                   int32                    `json:"AlbumCount,omitempty"`
	SongCount                    int32                    `json:"SongCount,omitempty"`
	MusicVideoCount              int32                    `json:"MusicVideoCount,omitempty"`
	Subviews                     []string                 `json:"Subviews,omitempty"`
	ListingsProviderId           string                   `json:"ListingsProviderId,omitempty"`
	ListingsChannelId            string                   `json:"ListingsChannelId,omitempty"`
	ListingsPath                 string                   `json:"ListingsPath,omitempty"`
	ListingsId                   string                   `json:"ListingsId,omitempty"`
	ListingsChannelName          string                   `json:"ListingsChannelName,omitempty"`
	ListingsChannelNumber        string                   `json:"ListingsChannelNumber,omitempty"`
	AffiliateCallSign            string                   `json:"AffiliateCallSign,omitempty"`
}

type BaseItemPerson struct {
	Name            string      `json:"Name,omitempty"`
	Id              string      `json:"Id,omitempty"`
	Role            string      `json:"Role,omitempty"`
	Type_           *PersonType `json:"Type,omitempty"`
	PrimaryImageTag string      `json:"PrimaryImageTag,omitempty"`
}

// C

type ChapterInfo struct {
	StartPositionTicks int64       `json:"StartPositionTicks,omitempty"`
	Name               string      `json:"Name,omitempty"`
	ImageTag           string      `json:"ImageTag,omitempty"`
	MarkerType         *MarkerType `json:"MarkerType,omitempty"`
	ChapterIndex       int32       `json:"ChapterIndex,omitempty"`
}

// E

type ExternalUrl struct {
	Name string `json:"Name,omitempty"`
	Url  string `json:"Url,omitempty"`
}

// M

type MediaSourceInfo struct {
	Protocol                   *MediaProtocol            `json:"Protocol,omitempty"`
	Id                         string                    `json:"Id,omitempty"`
	Path                       string                    `json:"Path,omitempty"`
	EncoderPath                string                    `json:"EncoderPath,omitempty"`
	EncoderProtocol            *MediaProtocol            `json:"EncoderProtocol,omitempty"`
	Type_                      *MediaSourceType          `json:"Type,omitempty"`
	ProbePath                  string                    `json:"ProbePath,omitempty"`
	ProbeProtocol              *MediaProtocol            `json:"ProbeProtocol,omitempty"`
	Container                  string                    `json:"Container,omitempty"`
	Size                       int64                     `json:"Size,omitempty"`
	Name                       string                    `json:"Name,omitempty"`
	SortName                   string                    `json:"SortName,omitempty"`
	IsRemote                   bool                      `json:"IsRemote,omitempty"`
	HasMixedProtocols          bool                      `json:"HasMixedProtocols,omitempty"`
	RunTimeTicks               int64                     `json:"RunTimeTicks,omitempty"`
	ContainerStartTimeTicks    int64                     `json:"ContainerStartTimeTicks,omitempty"`
	SupportsTranscoding        bool                      `json:"SupportsTranscoding,omitempty"`
	TrancodeLiveStartIndex     int32                     `json:"TrancodeLiveStartIndex,omitempty"`
	WallClockStart             time.Time                 `json:"WallClockStart,omitempty"`
	SupportsDirectStream       bool                      `json:"SupportsDirectStream,omitempty"`
	SupportsDirectPlay         bool                      `json:"SupportsDirectPlay,omitempty"`
	IsInfiniteStream           bool                      `json:"IsInfiniteStream,omitempty"`
	RequiresOpening            bool                      `json:"RequiresOpening,omitempty"`
	OpenToken                  string                    `json:"OpenToken,omitempty"`
	RequiresClosing            bool                      `json:"RequiresClosing,omitempty"`
	LiveStreamId               string                    `json:"LiveStreamId,omitempty"`
	BufferMs                   int32                     `json:"BufferMs,omitempty"`
	RequiresLooping            bool                      `json:"RequiresLooping,omitempty"`
	SupportsProbing            bool                      `json:"SupportsProbing,omitempty"`
	Video3DFormat              *Video3DFormat            `json:"Video3DFormat,omitempty"`
	MediaStreams               []MediaStream             `json:"MediaStreams,omitempty"`
	Formats                    []string                  `json:"Formats,omitempty"`
	Bitrate                    int32                     `json:"Bitrate,omitempty"`
	Timestamp                  *TransportStreamTimestamp `json:"Timestamp,omitempty"`
	RequiredHttpHeaders        *map[string]string        `json:"RequiredHttpHeaders,omitempty"`
	DirectStreamUrl            string                    `json:"DirectStreamUrl,omitempty"`
	AddApiKeyToDirectStreamUrl bool                      `json:"AddApiKeyToDirectStreamUrl,omitempty"`
	TranscodingUrl             string                    `json:"TranscodingUrl,omitempty"`
	TranscodingSubProtocol     string                    `json:"TranscodingSubProtocol,omitempty"`
	TranscodingContainer       string                    `json:"TranscodingContainer,omitempty"`
	AnalyzeDurationMs          int32                     `json:"AnalyzeDurationMs,omitempty"`
	ReadAtNativeFramerate      bool                      `json:"ReadAtNativeFramerate,omitempty"`
	DefaultAudioStreamIndex    int32                     `json:"DefaultAudioStreamIndex,omitempty"`
	DefaultSubtitleStreamIndex int32                     `json:"DefaultSubtitleStreamIndex,omitempty"`
	ItemId                     string                    `json:"ItemId,omitempty"`
	ServerId                   string                    `json:"ServerId,omitempty"`
}

type MediaStream struct {
	Codec                           string                  `json:"Codec,omitempty"`
	CodecTag                        string                  `json:"CodecTag,omitempty"`
	Language                        string                  `json:"Language,omitempty"`
	ColorTransfer                   string                  `json:"ColorTransfer,omitempty"`
	ColorPrimaries                  string                  `json:"ColorPrimaries,omitempty"`
	ColorSpace                      string                  `json:"ColorSpace,omitempty"`
	Comment                         string                  `json:"Comment,omitempty"`
	StreamStartTimeTicks            int64                   `json:"StreamStartTimeTicks,omitempty"`
	TimeBase                        string                  `json:"TimeBase,omitempty"`
	Title                           string                  `json:"Title,omitempty"`
	Extradata                       string                  `json:"Extradata,omitempty"`
	VideoRange                      string                  `json:"VideoRange,omitempty"`
	DisplayTitle                    string                  `json:"DisplayTitle,omitempty"`
	DisplayLanguage                 string                  `json:"DisplayLanguage,omitempty"`
	NalLengthSize                   string                  `json:"NalLengthSize,omitempty"`
	IsInterlaced                    bool                    `json:"IsInterlaced,omitempty"`
	IsAVC                           bool                    `json:"IsAVC,omitempty"`
	ChannelLayout                   string                  `json:"ChannelLayout,omitempty"`
	BitRate                         int32                   `json:"BitRate,omitempty"`
	BitDepth                        int32                   `json:"BitDepth,omitempty"`
	RefFrames                       int32                   `json:"RefFrames,omitempty"`
	Rotation                        int32                   `json:"Rotation,omitempty"`
	Channels                        int32                   `json:"Channels,omitempty"`
	SampleRate                      int32                   `json:"SampleRate,omitempty"`
	IsDefault                       bool                    `json:"IsDefault,omitempty"`
	IsForced                        bool                    `json:"IsForced,omitempty"`
	IsHearingImpaired               bool                    `json:"IsHearingImpaired,omitempty"`
	Height                          int32                   `json:"Height,omitempty"`
	Width                           int32                   `json:"Width,omitempty"`
	AverageFrameRate                float32                 `json:"AverageFrameRate,omitempty"`
	RealFrameRate                   float32                 `json:"RealFrameRate,omitempty"`
	Profile                         string                  `json:"Profile,omitempty"`
	Type_                           *MediaStreamType        `json:"Type,omitempty"`
	AspectRatio                     string                  `json:"AspectRatio,omitempty"`
	Index                           int32                   `json:"Index,omitempty"`
	IsExternal                      bool                    `json:"IsExternal,omitempty"`
	DeliveryMethod                  *SubtitleDeliveryMethod `json:"DeliveryMethod,omitempty"`
	DeliveryUrl                     string                  `json:"DeliveryUrl,omitempty"`
	IsExternalUrl                   bool                    `json:"IsExternalUrl,omitempty"`
	IsTextSubtitleStream            bool                    `json:"IsTextSubtitleStream,omitempty"`
	SupportsExternalStream          bool                    `json:"SupportsExternalStream,omitempty"`
	Path                            string                  `json:"Path,omitempty"`
	Protocol                        *MediaProtocol          `json:"Protocol,omitempty"`
	PixelFormat                     string                  `json:"PixelFormat,omitempty"`
	Level                           float64                 `json:"Level,omitempty"`
	IsAnamorphic                    bool                    `json:"IsAnamorphic,omitempty"`
	ExtendedVideoType               *ExtendedVideoTypes     `json:"ExtendedVideoType,omitempty"`
	ExtendedVideoSubType            *ExtendedVideoSubTypes  `json:"ExtendedVideoSubType,omitempty"`
	ExtendedVideoSubTypeDescription string                  `json:"ExtendedVideoSubTypeDescription,omitempty"`
	ItemId                          string                  `json:"ItemId,omitempty"`
	ServerId                        string                  `json:"ServerId,omitempty"`
	AttachmentSize                  int32                   `json:"AttachmentSize,omitempty"`
	MimeType                        string                  `json:"MimeType,omitempty"`
	SubtitleLocationType            *SubtitleLocationType   `json:"SubtitleLocationType,omitempty"`
}

type MediaUrl struct {
	Url  string `json:"Url,omitempty"`
	Name string `json:"Name,omitempty"`
}

// N

type NameIdPair struct {
	Name string `json:"Name,omitempty"`
	Id   string `json:"Id,omitempty"`
}

type NameLongIdPair struct {
	Name string `json:"Name,omitempty"`
	Id   int64  `json:"Id,omitempty"`
}

// P

type PlayerStateInfo struct {
	PositionTicks       int64       `json:"PositionTicks,omitempty"`
	CanSeek             bool        `json:"CanSeek,omitempty"`
	IsPaused            bool        `json:"IsPaused,omitempty"`
	IsMuted             bool        `json:"IsMuted,omitempty"`
	VolumeLevel         int32       `json:"VolumeLevel,omitempty"`
	AudioStreamIndex    int32       `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex int32       `json:"SubtitleStreamIndex,omitempty"`
	MediaSourceId       string      `json:"MediaSourceId,omitempty"`
	PlayMethod          *PlayMethod `json:"PlayMethod,omitempty"`
	RepeatMode          *RepeatMode `json:"RepeatMode,omitempty"`
	SubtitleOffset      int32       `json:"SubtitleOffset,omitempty"`
	Shuffle             bool        `json:"Shuffle,omitempty"`
	PlaybackRate        float64     `json:"PlaybackRate,omitempty"`
}

type ProcessMetricPoint struct {
	Time          string  `json:"Time,omitempty"`
	CpuPercent    float64 `json:"CpuPercent,omitempty"`
	VirtualMemory float64 `json:"VirtualMemory,omitempty"`
	WorkingSet    float64 `json:"WorkingSet,omitempty"`
}

type ProcessStatistics struct {
	CurrentCpu           float64              `json:"CurrentCpu,omitempty"`
	AverageCpu           float64              `json:"AverageCpu,omitempty"`
	CurrentVirtualMemory float64              `json:"CurrentVirtualMemory,omitempty"`
	CurrentWorkingSet    float64              `json:"CurrentWorkingSet,omitempty"`
	Metrics              []ProcessMetricPoint `json:"Metrics,omitempty"`
}

// Q

type QueryResultBaseItemDto struct {
	Items            []BaseItemDto `json:"Items,omitempty"`
	TotalRecordCount int32         `json:"TotalRecordCount,omitempty"`
}

// S

type SessionInfo struct {
	PlayState             *PlayerStateInfo  `json:"PlayState,omitempty"`
	AdditionalUsers       []SessionUserInfo `json:"AdditionalUsers,omitempty"`
	RemoteEndPoint        string            `json:"RemoteEndPoint,omitempty"`
	Protocol              string            `json:"Protocol,omitempty"`
	PlayableMediaTypes    []string          `json:"PlayableMediaTypes,omitempty"`
	PlaylistItemId        string            `json:"PlaylistItemId,omitempty"`
	PlaylistIndex         int32             `json:"PlaylistIndex,omitempty"`
	PlaylistLength        int32             `json:"PlaylistLength,omitempty"`
	Id                    string            `json:"Id,omitempty"`
	ServerId              string            `json:"ServerId,omitempty"`
	UserId                string            `json:"UserId,omitempty"`
	UserName              string            `json:"UserName,omitempty"`
	UserPrimaryImageTag   string            `json:"UserPrimaryImageTag,omitempty"`
	Client                string            `json:"Client,omitempty"`
	LastActivityDate      time.Time         `json:"LastActivityDate,omitempty"`
	DeviceName            string            `json:"DeviceName,omitempty"`
	DeviceType            string            `json:"DeviceType,omitempty"`
	NowPlayingItem        *BaseItemDto      `json:"NowPlayingItem,omitempty"`
	InternalDeviceId      int64             `json:"InternalDeviceId,omitempty"`
	DeviceId              string            `json:"DeviceId,omitempty"`
	ApplicationVersion    string            `json:"ApplicationVersion,omitempty"`
	AppIconUrl            string            `json:"AppIconUrl,omitempty"`
	SupportedCommands     []string          `json:"SupportedCommands,omitempty"`
	TranscodingInfo       *TranscodingInfo  `json:"TranscodingInfo,omitempty"`
	SupportsRemoteControl bool              `json:"SupportsRemoteControl,omitempty"`
}

type SessionUserInfo struct {
	UserId         string `json:"UserId,omitempty"`
	UserName       string `json:"UserName,omitempty"`
	UserInternalId int64  `json:"UserInternalId,omitempty"`
}

// T

type TranscodingInfo struct {
	AudioCodec                    string                    `json:"AudioCodec,omitempty"`
	VideoCodec                    string                    `json:"VideoCodec,omitempty"`
	SubProtocol                   string                    `json:"SubProtocol,omitempty"`
	Container                     string                    `json:"Container,omitempty"`
	IsVideoDirect                 bool                      `json:"IsVideoDirect,omitempty"`
	IsAudioDirect                 bool                      `json:"IsAudioDirect,omitempty"`
	Bitrate                       int32                     `json:"Bitrate,omitempty"`
	AudioBitrate                  int32                     `json:"AudioBitrate,omitempty"`
	VideoBitrate                  int32                     `json:"VideoBitrate,omitempty"`
	Framerate                     float32                   `json:"Framerate,omitempty"`
	CompletionPercentage          float64                   `json:"CompletionPercentage,omitempty"`
	TranscodingPositionTicks      float64                   `json:"TranscodingPositionTicks,omitempty"`
	TranscodingStartPositionTicks float64                   `json:"TranscodingStartPositionTicks,omitempty"`
	Width                         int32                     `json:"Width,omitempty"`
	Height                        int32                     `json:"Height,omitempty"`
	AudioChannels                 int32                     `json:"AudioChannels,omitempty"`
	TranscodeReasons              []TranscodeReason         `json:"TranscodeReasons,omitempty"`
	CurrentCpuUsage               float64                   `json:"CurrentCpuUsage,omitempty"`
	AverageCpuUsage               float64                   `json:"AverageCpuUsage,omitempty"`
	CpuHistory                    []TupleDoubleDouble       `json:"CpuHistory,omitempty"`
	ProcessStatistics             *ProcessStatistics        `json:"ProcessStatistics,omitempty"`
	CurrentThrottle               int32                     `json:"CurrentThrottle,omitempty"`
	VideoDecoder                  string                    `json:"VideoDecoder,omitempty"`
	VideoDecoderIsHardware        bool                      `json:"VideoDecoderIsHardware,omitempty"`
	VideoDecoderMediaType         string                    `json:"VideoDecoderMediaType,omitempty"`
	VideoDecoderHwAccel           string                    `json:"VideoDecoderHwAccel,omitempty"`
	VideoEncoder                  string                    `json:"VideoEncoder,omitempty"`
	VideoEncoderIsHardware        bool                      `json:"VideoEncoderIsHardware,omitempty"`
	VideoEncoderMediaType         string                    `json:"VideoEncoderMediaType,omitempty"`
	VideoEncoderHwAccel           string                    `json:"VideoEncoderHwAccel,omitempty"`
	VideoPipelineInfo             []TranscodingVpStepInfo   `json:"VideoPipelineInfo,omitempty"`
	SubtitlePipelineInfos         [][]TranscodingVpStepInfo `json:"SubtitlePipelineInfos,omitempty"`
}

type TranscodingVpStepInfo struct {
	StepType            *TranscodingVpStepTypes `json:"StepType,omitempty"`
	StepTypeName        string                  `json:"StepTypeName,omitempty"`
	HardwareContextName string                  `json:"HardwareContextName,omitempty"`
	IsHardwareContext   bool                    `json:"IsHardwareContext,omitempty"`
	Name                string                  `json:"Name,omitempty"`
	Short               string                  `json:"Short,omitempty"`
	FfmpegName          string                  `json:"FfmpegName,omitempty"`
	FfmpegDescription   string                  `json:"FfmpegDescription,omitempty"`
	FfmpegOptions       string                  `json:"FfmpegOptions,omitempty"`
	Param               string                  `json:"Param,omitempty"`
	ParamShort          string                  `json:"ParamShort,omitempty"`
}

// U

type UserConfiguration struct {
	AudioLanguagePreference    string                `json:"AudioLanguagePreference,omitempty"`
	PlayDefaultAudioTrack      bool                  `json:"PlayDefaultAudioTrack,omitempty"`
	SubtitleLanguagePreference string                `json:"SubtitleLanguagePreference,omitempty"`
	ProfilePin                 string                `json:"ProfilePin,omitempty"`
	DisplayMissingEpisodes     bool                  `json:"DisplayMissingEpisodes,omitempty"`
	SubtitleMode               *SubtitlePlaybackMode `json:"SubtitleMode,omitempty"`
	OrderedViews               []string              `json:"OrderedViews,omitempty"`
	LatestItemsExcludes        []string              `json:"LatestItemsExcludes,omitempty"`
	MyMediaExcludes            []string              `json:"MyMediaExcludes,omitempty"`
	HidePlayedInLatest         bool                  `json:"HidePlayedInLatest,omitempty"`
	HidePlayedInMoreLikeThis   bool                  `json:"HidePlayedInMoreLikeThis,omitempty"`
	HidePlayedInSuggestions    bool                  `json:"HidePlayedInSuggestions,omitempty"`
	RememberAudioSelections    bool                  `json:"RememberAudioSelections,omitempty"`
	RememberSubtitleSelections bool                  `json:"RememberSubtitleSelections,omitempty"`
	EnableNextEpisodeAutoPlay  bool                  `json:"EnableNextEpisodeAutoPlay,omitempty"`
	ResumeRewindSeconds        int32                 `json:"ResumeRewindSeconds,omitempty"`
	IntroSkipMode              *SegmentSkipMode      `json:"IntroSkipMode,omitempty"`
	EnableLocalPassword        bool                  `json:"EnableLocalPassword,omitempty"`
}

type UserDto struct {
	Name                      string               `json:"Name,omitempty"`
	ServerId                  string               `json:"ServerId,omitempty"`
	ServerName                string               `json:"ServerName,omitempty"`
	Prefix                    string               `json:"Prefix,omitempty"`
	ConnectUserName           string               `json:"ConnectUserName,omitempty"`
	DateCreated               time.Time            `json:"DateCreated,omitempty"`
	ConnectLinkType           *ConnectUserLinkType `json:"ConnectLinkType,omitempty"`
	Id                        string               `json:"Id,omitempty"`
	PrimaryImageTag           string               `json:"PrimaryImageTag,omitempty"`
	HasPassword               bool                 `json:"HasPassword,omitempty"`
	HasConfiguredPassword     bool                 `json:"HasConfiguredPassword,omitempty"`
	EnableAutoLogin           bool                 `json:"EnableAutoLogin,omitempty"`
	LastLoginDate             time.Time            `json:"LastLoginDate,omitempty"`
	LastActivityDate          time.Time            `json:"LastActivityDate,omitempty"`
	Configuration             *UserConfiguration   `json:"Configuration,omitempty"`
	Policy                    *UserPolicy          `json:"Policy,omitempty"`
	PrimaryImageAspectRatio   float64              `json:"PrimaryImageAspectRatio,omitempty"`
	HasConfiguredEasyPassword bool                 `json:"HasConfiguredEasyPassword,omitempty"`
	UserItemShareLevel        *UserItemShareLevel  `json:"UserItemShareLevel,omitempty"`
}

type UserItemDataDto struct {
	Rating                float64   `json:"Rating,omitempty"`
	PlayedPercentage      float64   `json:"PlayedPercentage,omitempty"`
	UnplayedItemCount     int32     `json:"UnplayedItemCount,omitempty"`
	PlaybackPositionTicks int64     `json:"PlaybackPositionTicks,omitempty"`
	PlayCount             int32     `json:"PlayCount,omitempty"`
	IsFavorite            bool      `json:"IsFavorite,omitempty"`
	LastPlayedDate        time.Time `json:"LastPlayedDate,omitempty"`
	Played                bool      `json:"Played,omitempty"`
	Key                   string    `json:"Key,omitempty"`
	ItemId                string    `json:"ItemId,omitempty"`
	ServerId              string    `json:"ServerId,omitempty"`
}

type UserPolicy struct {
	IsAdministrator                  bool             `json:"IsAdministrator,omitempty"`
	IsHidden                         bool             `json:"IsHidden,omitempty"`
	IsHiddenRemotely                 bool             `json:"IsHiddenRemotely,omitempty"`
	IsHiddenFromUnusedDevices        bool             `json:"IsHiddenFromUnusedDevices,omitempty"`
	IsDisabled                       bool             `json:"IsDisabled,omitempty"`
	LockedOutDate                    int64            `json:"LockedOutDate,omitempty"`
	MaxParentalRating                int32            `json:"MaxParentalRating,omitempty"`
	AllowTagOrRating                 bool             `json:"AllowTagOrRating,omitempty"`
	BlockedTags                      []string         `json:"BlockedTags,omitempty"`
	IsTagBlockingModeInclusive       bool             `json:"IsTagBlockingModeInclusive,omitempty"`
	IncludeTags                      []string         `json:"IncludeTags,omitempty"`
	EnableUserPreferenceAccess       bool             `json:"EnableUserPreferenceAccess,omitempty"`
	AccessSchedules                  []AccessSchedule `json:"AccessSchedules,omitempty"`
	BlockUnratedItems                []UnratedItem    `json:"BlockUnratedItems,omitempty"`
	EnableRemoteControlOfOtherUsers  bool             `json:"EnableRemoteControlOfOtherUsers,omitempty"`
	EnableSharedDeviceControl        bool             `json:"EnableSharedDeviceControl,omitempty"`
	EnableRemoteAccess               bool             `json:"EnableRemoteAccess,omitempty"`
	EnableLiveTvManagement           bool             `json:"EnableLiveTvManagement,omitempty"`
	EnableLiveTvAccess               bool             `json:"EnableLiveTvAccess,omitempty"`
	EnableMediaPlayback              bool             `json:"EnableMediaPlayback,omitempty"`
	EnableAudioPlaybackTranscoding   bool             `json:"EnableAudioPlaybackTranscoding,omitempty"`
	EnableVideoPlaybackTranscoding   bool             `json:"EnableVideoPlaybackTranscoding,omitempty"`
	EnablePlaybackRemuxing           bool             `json:"EnablePlaybackRemuxing,omitempty"`
	EnableContentDeletion            bool             `json:"EnableContentDeletion,omitempty"`
	RestrictedFeatures               []string         `json:"RestrictedFeatures,omitempty"`
	EnableContentDeletionFromFolders []string         `json:"EnableContentDeletionFromFolders,omitempty"`
	EnableContentDownloading         bool             `json:"EnableContentDownloading,omitempty"`
	EnableSubtitleDownloading        bool             `json:"EnableSubtitleDownloading,omitempty"`
	EnableSubtitleManagement         bool             `json:"EnableSubtitleManagement,omitempty"`
	EnableSyncTranscoding            bool             `json:"EnableSyncTranscoding,omitempty"`
	EnableMediaConversion            bool             `json:"EnableMediaConversion,omitempty"`
	EnabledChannels                  []string         `json:"EnabledChannels,omitempty"`
	EnableAllChannels                bool             `json:"EnableAllChannels,omitempty"`
	EnabledFolders                   []string         `json:"EnabledFolders,omitempty"`
	EnableAllFolders                 bool             `json:"EnableAllFolders,omitempty"`
	InvalidLoginAttemptCount         int32            `json:"InvalidLoginAttemptCount,omitempty"`
	EnablePublicSharing              bool             `json:"EnablePublicSharing,omitempty"`
	BlockedMediaFolders              []string         `json:"BlockedMediaFolders,omitempty"`
	RemoteClientBitrateLimit         int32            `json:"RemoteClientBitrateLimit,omitempty"`
	AuthenticationProviderId         string           `json:"AuthenticationProviderId,omitempty"`
	ExcludedSubFolders               []string         `json:"ExcludedSubFolders,omitempty"`
	SimultaneousStreamLimit          int32            `json:"SimultaneousStreamLimit,omitempty"`
	EnabledDevices                   []string         `json:"EnabledDevices,omitempty"`
	EnableAllDevices                 bool             `json:"EnableAllDevices,omitempty"`
	AllowCameraUpload                bool             `json:"AllowCameraUpload,omitempty"`
	AllowSharingPersonalItems        bool             `json:"AllowSharingPersonalItems,omitempty"`
}
