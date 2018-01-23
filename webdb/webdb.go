package importer

import "time"

type Article struct {
  ID                      uint
  Guid                    string
  ResourcePk              string
  LicenseID               uint
  License                 License
  LanguageID              uint
  Language                Language
  LocationID              uint
  Location                Location
  StylesheetID            uint
  JavascriptID            uint
  BibliographicCitationID uint
  BibliographicCitation   BibliographicCitation
  Owner                   string
  Name                    string
  SourceUrl               string
  Body                    string
  CreatedAt               time.Time
  UpdatedAt               time.Time
  ResourceID              uint
  Resource                Resource
  RightsStatement         string
  PageID                  uint
  Page                    Page
  ArticlesCollectedPages  []ArticlesCollectedPage
}

type ArticlesCollectedPage struct {
  CollectedPageID uint
  CollectedPage   CollectedPage
  ArticleID       uint
  Article         Article
  Position        uint
}

type Attribution struct {
  ID         uint
  RoleID     uint
  Role       Role
  Value      string
  CreatedAt  time.Time
  UpdatedAt  time.Time
  Url        string
  ResourceID uint
  Resource   Resource
  ResourcePk string
}

type BibliographicCitation struct {
  Articles   []Article
  ID         uint
  ResourceID uint
  Resource   Resource
  Body       string
  CreatedAt  time.Time
  UpdatedAt  time.Time
  Media      []Medium
}

type Change struct {
  ID         uint
  UserID     uint
  User       User
  PageID     uint
  Page       Page
  ActivityID uint
  CreatedAt  time.Time
  UpdatedAt  time.Time
}

type CollectedPage struct {
  ArticlesCollectedPages []ArticlesCollectedPage
  ID                     uint
  CollectionID           uint
  Collection             Collection
  PageID                 uint
  Page                   Page
  Position               uint
  CreatedAt              time.Time
  UpdatedAt              time.Time
  Annotation             string
  CollectedPagesLinks    []CollectedPagesLink
  CollectedPagesMedia    []CollectedPagesMedium
}

type CollectedPagesLink struct {
  CollectedPageID uint
  CollectedPage   CollectedPage
  LinkID          uint
  Link            Link
  Position        uint
}

type CollectedPagesMedium struct {
  CollectedPageID uint
  CollectedPage   CollectedPage
  MediumID        uint
  Medium          Medium
  Position        uint
}

type Collecting struct {
  ID                     uint
  UserID                 uint
  User                   User
  CollectionID           uint
  Collection             Collection
  Action                 uint
  PageID                 uint
  Page                   Page
  AssociatedCollectionID uint
  AssociatedCollection   Collection
  ChangedField           string
  ChangedFrom            string
  ChangedTo              string
  CreatedAt              time.Time
  UpdatedAt              time.Time
}

type CollectionAssociation struct {
  ID           uint
  CollectionID uint
  Collection   Collection
  Position     uint
  CreatedAt    time.Time
  UpdatedAt    time.Time
  AssociatedID uint
  Associated   Collection
  Annotation   string
}

type Collection struct {
  CollectedPages              []CollectedPage
  Collectings                 []Collecting
  CollectionAssociations      []CollectionAssociation
  ID                          uint
  Name                        string
  Description                 string
  IconFileName                string
  IconFileSize                uint
  IconUpdatedAt               time.Time
  CreatedAt                   time.Time
  UpdatedAt                   time.Time
  CollectedPagesCount         uint
  CollectionAssociationsCount uint
  DefaultSort                 uint
  CollectionsUsers            []CollectionsUser
}

type CollectionsUser struct {
  UserID       uint
  User         User
  CollectionID uint
  Collection   Collection
  IsManager    bool
  CreatedAt    time.Time
  UpdatedAt    time.Time
}

type ContentEdit struct {
  ID            uint
  UserIDID      uint
  User          User
  PageContentID uint
  PageContent   PageContent
  ChangedField  string
  ChangedFrom   string
  ChangedTo     string
  Comment       string
  CreatedAt     time.Time
  UpdatedAt     time.Time
}

type ContentReposition struct {
  ID            uint
  UserIDID      uint
  User          User
  PageContentID uint
  PageContent   PageContent
  ChangedFrom   uint
  ChangedTo     uint
  CreatedAt     time.Time
  UpdatedAt     time.Time
}

type ContentSection struct {
  ID        uint
  SectionID uint
  Section   Section
}

type Curation struct {
  ID                 uint
  UserID             uint
  User               User
  PageContentID      uint
  PageContent        PageContent
  Trust              uint
  IsIncorrect        bool
  IsMisidentified    bool
  IsHidden           bool
  IsDuplicate        bool
  IsLowQuality       bool
  CreatedAt          time.Time
  UpdatedAt          time.Time
  OldTrust           uint
  OldIsIncorrect     bool
  OldIsMisidentified bool
  OldIsHidden        bool
  OldIsDuplicate     bool
  OldIsLowQuality    bool
  Comment            string
}

type DelayedJob struct {
  ID        uint
  Priority  uint
  Attempts  uint
  Handler   string
  LastError string
  RunAt     time.Time
  LockedAt  time.Time
  FailedAt  time.Time
  LockedBy  string
  Queue     string
  CreatedAt time.Time
  UpdatedAt time.Time
}

type Identifier struct {
  ID             uint
  ResourceID     uint
  Resource       Resource
  NodeID         uint
  Node           Node
  NodeResourcePk string
  IDentifier     string
}

type ImageInfo struct {
  ID           uint
  ResourceID   uint
  Resource     Resource
  MediumID     uint
  Medium       Medium
  OriginalSize string
  LargeSize    string
  MediumSize   string
  SmallSize    string
  CropX        float32
  CropY        float32
  CropW        float32
  CreatedAt    time.Time
  UpdatedAt    time.Time
  ResourcePk   string
}

type ImportEvent struct {
  ID          uint
  ImportLogID uint
  ImportLog   ImportLog
  Cat         uint
  Body        string
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type ImportLog struct {
  ImportEvents []ImportEvent
  ID           uint
  ResourceID   uint
  Resource     Resource
  CompletedAt  time.Time
  FailedAt     time.Time
  Status       string
  CreatedAt    time.Time
  UpdatedAt    time.Time
}

type ImportRun struct {
  ID          uint
  CompletedAt time.Time
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type Javascript struct {
  ID         uint
  ResourceID uint
  Resource   Resource
  Filename   string
}

type Language struct {
  Articles      []Article
  ID            uint
  Code          string
  Group         string
  CanBrowseSite bool
  Links         []Link
  Media         []Medium
  Vernaculars   []Vernacular
}

type License struct {
  Articles              []Article
  ID                    uint
  Name                  string
  SourceUrl             string
  IconUrl               string
  CanBeChosenByPartners bool
  CreatedAt             time.Time
  UpdatedAt             time.Time
  Media                 []Medium
  Resources             []Resource
}

type Link struct {
  CollectedPagesLinks []CollectedPagesLink
  ID                  uint
  Guid                string
  ResourcePk          string
  LanguageID          uint
  Language            Language
  Name                string
  SourceUrl           string
  Description         string
  IconUrl             string
  CreatedAt           time.Time
  UpdatedAt           time.Time
  ResourceID          uint
  Resource            Resource
  RightsStatement     string
  PageID              uint
  Page                Page
}

type Location struct {
  Articles        []Article
  ID              uint
  ResourceID      uint
  Resource        Resource
  Location        string
  Longitude       float32
  Latitude        float32
  Altitude        float32
  SpatialLocation string
  Media           []Medium
}

type Medium struct {
  CollectedPagesMedia     []CollectedPagesMedium
  ImageInfos              []ImageInfo
  ID                      uint
  Guid                    string
  ResourcePk              string
  Subclass                uint
  Format                  uint
  LicenseID               uint
  License                 License
  LanguageID              uint
  Language                Language
  LocationID              uint
  Location                Location
  StylesheetID            uint
  JavascriptID            uint
  BibliographicCitationID uint
  BibliographicCitation   BibliographicCitation
  Owner                   string
  Name                    string
  SourceUrl               string
  Description             string
  BaseUrl                 string
  CreatedAt               time.Time
  UpdatedAt               time.Time
  UnmodifiedUrl           string
  SourcePageUrl           string
  ResourceID              uint
  Resource                Resource
  RightsStatement         string
  PageID                  uint
  Page                    *Page    // TODO: Why a pointer? Without it, it causes "invalid recursive type Page" (?!)
  UsageStatement          string
  PageIcons               []PageIcon
  Pages                   []Page
}

type NodeAncestor struct {
  ID                 uint
  ResourceID         uint
  Resource           Resource
  NodeID             uint
  Node               Node
  AncestorID         uint
  Ancestor           Node
  NodeResourcePk     string
  AncestorResourcePk string
  Depth              uint
}

type Node struct {
  Identifiers      []Identifier
  NodeAncestors    []NodeAncestor
  ScientificNames  []ScientificName
  ID               uint
  ResourceID       uint
  Resource         Resource
  PageID           uint
  Page             Page
  RankID           uint
  Rank             Rank
  ParentID         uint
  Parent           *Node
  ScientificName   string
  CanonicalForm    string
  ResourcePk       string
  SourceUrl        string
  IsHidden         bool
  InUnmappedArea   bool
  ChildrenCount    uint
  CreatedAt        time.Time
  UpdatedAt        time.Time
  HasBreadcrumb    bool
  ParentResourcePk string
  Landmark         uint
  // TMP removing because it screws up bulk insert?
  // References       []Reference
  // RefineryPages    []RefineryPage
  // Sections         []Section
  // TaxonRemarks     []TaxonRemark
  // Vernaculars      []Vernacular
}

func (Node) TableName() (string) {
  return "nodes"
}

type OccurrenceMap struct {
  ID         uint
  ResourceID uint
  Resource   Resource
  PageID     uint
  Page       Page
  Url        string
}

type OpenAuthentication struct {
  ID        uint
  Provider  string
  Uid       string
  CreatedAt time.Time
  UpdatedAt time.Time
  UserID    uint
  User      User
}

type PageContent struct {
  ContentEdits             []ContentEdit
  ContentRepositions       []ContentReposition
  Curations                []Curation
  ID                       uint
  PageID                   uint
  Page                     Page
  ResourceID               uint
  Resource                 Resource
  SourcePageID             uint
  SourcePage               Page
  Position                 uint
  AssociationAddedByUserID uint
  AssociationAddedByUser   User
  Trust                    uint
  IsIncorrect              bool
  IsMisidentified          bool
  IsHidden                 bool
  IsDuplicate              bool
  IsLowQuality             bool
  CreatedAt                time.Time
  UpdatedAt                time.Time
}

type PageIcon struct {
  ID        uint
  PageID    uint
  Page      Page
  UserID    uint
  User      User
  MediumID  uint
  Medium    Medium
  CreatedAt time.Time
  UpdatedAt time.Time
}

type Page struct {
  Articles             []Article
  Changes              []Change
  CollectedPages       []CollectedPage
  Collectings          []Collecting
  Links                []Link
  Media                []Medium
  Nodes                []Node
  OccurrenceMaps       []OccurrenceMap
  PageContents         []PageContent
  PageIcons            []PageIcon
  ID                   uint
  NativeNodeID         uint
  MovedToPageID        uint
  MovedToPage          *Page
  CreatedAt            time.Time
  UpdatedAt            time.Time
  PageContentsCount    uint
  MediaCount           uint
  ArticlesCount        uint
  LinksCount           uint
  MapsCount            uint
  DataCount            uint
  NodesCount           uint
  VernacularsCount     uint
  ScientificNamesCount uint
  ReferentsCount       uint
  SpeciesCount         uint
  IsExtinct            bool
  IsMarine             bool
  HasCheckedExtinct    bool
  HasCheckedMarine     bool
  IucnStatus           string
  TrophicStrategy      string
  GeographicContext    string
  Habitat              string
  PageRichness         uint
  MediumID             uint
  Medium               Medium
  PagesReferents       []PagesReferent
  ScientificNames      []ScientificName
  SearchSuggestions    []SearchSuggestion
  Vernaculars          []Vernacular
}

type PagesReferent struct {
  PageID     uint
  Page       Page
  ReferentID uint
  Referent   Referent
  Position   uint
}

type Partner struct {
  ID            uint
  Name          string
  Abbr          string
  ShortName     string
  HomepageUrl   string
  Description   string
  Notes         string
  LinksJson     string
  CreatedAt     time.Time
  UpdatedAt     time.Time
  IconFileName  string
  IconFileSize  uint
  IconUpdatedAt time.Time
  RepositoryID  uint
  PartnersUsers []PartnersUser
  Resources     []Resource
}

type PartnersUser struct {
  PartnerID uint
  Partner   Partner
  UserID    uint
  User      User
}

type Rank struct {
  Nodes   []Node
  ID      uint
  Name    string
  TreatAs uint
}

type Reference struct {
  ParentID   uint
  ReferentID uint
  Referent   Referent
  ResourceID uint
  Resource   Resource
}

type Referent struct {
  PagesReferents []PagesReferent
  References     []Reference
  ID             uint
  Body           string
  CreatedAt      time.Time
  UpdatedAt      time.Time
  ResourceID     uint
  Resource       Resource
}

type RefineryImageTranslation struct {
  ID              uint
  RefineryImageID uint
  RefineryImage   RefineryImage
  Locale          string
  CreatedAt       time.Time
  UpdatedAt       time.Time
  ImageAlt        string
  ImageTitle      string
}

type RefineryImage struct {
  RefineryImageTranslations []RefineryImageTranslation
  ID                        uint
  ImageName                 string
  ImageSize                 uint
  ImageWidth                uint
  ImageHeight               uint
  ImageUid                  string
  CreatedAt                 time.Time
  UpdatedAt                 time.Time
  ImageTitle                string
  ImageAlt                  string
}

type RefineryPagePartTranslation struct {
  ID                 uint
  RefineryPagePartID uint
  RefineryPagePart   RefineryPagePart
  Locale             string
  CreatedAt          time.Time
  UpdatedAt          time.Time
  Body               string
}

type RefineryPagePart struct {
  RefineryPagePartTranslations []RefineryPagePartTranslation
  ID                           uint
  RefineryPageID               uint
  RefineryPage                 RefineryPage
  Slug                         string
  Body                         string
  Position                     uint
  CreatedAt                    time.Time
  UpdatedAt                    time.Time
  Title                        string
}

type RefineryPageTranslation struct {
  ID             uint
  RefineryPageID uint
  RefineryPage   RefineryPage
  Locale         string
  CreatedAt      time.Time
  UpdatedAt      time.Time
  Title          string
  CustomSlug     string
  MenuTitle      string
  Slug           string
}

type RefineryPage struct {
  RefineryPageParts        []RefineryPagePart
  RefineryPageTranslations []RefineryPageTranslation
  ID                       uint
  ParentID                 uint
  Parent                   Node
  Path                     string
  Slug                     string
  CustomSlug               string
  ShowInMenu               bool
  LinkUrl                  string
  MenuMatch                string
  Deletable                bool
  Draft                    bool
  SkipToFirstChild         bool
  Lft                      uint
  Rgt                      uint
  Depth                    uint
  ViewTemplate             string
  LayoutTemplate           string
  CreatedAt                time.Time
  UpdatedAt                time.Time
  ShowDate                 bool
}

type RefineryResourceTranslation struct {
  ID                 uint
  RefineryResourceID uint
  RefineryResource   RefineryResource
  Locale             string
  CreatedAt          time.Time
  UpdatedAt          time.Time
  ResourceTitle      string
}

type RefineryResource struct {
  RefineryResourceTranslations []RefineryResourceTranslation
  ID                           uint
  FileName                     string
  FileSize                     uint
  FileUid                      string
  FileExt                      string
  CreatedAt                    time.Time
  UpdatedAt                    time.Time
}

type Resource struct {
  Articles               []Article
  Attributions           []Attribution
  BibliographicCitations []BibliographicCitation
  Identifiers            []Identifier
  ImageInfos             []ImageInfo
  ImportLogs             []ImportLog
  Javascripts            []Javascript
  Links                  []Link
  Locations              []Location
  Media                  []Medium
  NodeAncestors          []NodeAncestor
  Nodes                  []Node
  OccurrenceMaps         []OccurrenceMap
  PageContents           []PageContent
  References             []Reference
  Referents              []Referent
  ID                     uint
  PartnerID              uint
  Partner                Partner
  Name                   string
  Url                    string
  Description            string
  Notes                  string
  NodesCount             uint
  IsBrowsable            bool
  HasDuplicateNodes      bool
  NodeSourceUrlTemplate  string
  LastPublishedAt        time.Time
  LastPublishSeconds     uint
  DatasetLicenseID       uint
  DatasetLicense         License
  DatasetRightsHolder    string
  DatasetRightsStatement string
  CreatedAt              time.Time
  UpdatedAt              time.Time
  IconFileName           string
  IconFileSize           uint
  IconUpdatedAt          time.Time
  Abbr                   string
  RepositoryID           uint
  ScientificNames        []ScientificName
  Stylesheets            []Stylesheet
  Vernaculars            []Vernacular
}

type Role struct {
  Attributions []Attribution
  ID           uint
  Name         string
  CreatedAt    time.Time
  UpdatedAt    time.Time
}

type ScientificName struct {
  ID                   uint
  NodeID               uint
  PageID               uint
  Page                 Page
  Italicized           string
  CanonicalForm        string
  TaxonomicStatusID    uint
  TaxonomicStatus      TaxonomicStatus
  IsPreferred          bool
  CreatedAt            time.Time
  UpdatedAt            time.Time
  ResourceID           uint
  Resource             Resource
  NodeResourcePk       string
  SourceReference      string
  Genus                string
  SpecificEpithet      string
  InfraspecificEpithet string
  InfragenericEpithet  string
  Uninomial            string
  Verbatim             string
  Authorship           string
  Publication          string
  Remarks              string
  ParseQuality         uint
  Year                 uint
  Hybrid               bool
  Surrogate            bool
  Virus                bool
  Attribution          string
}

type SearchSuggestion struct {
  ID          uint
  PageID      uint
  Page        Page
  SynonymOfID uint
  SynonymOf   *SearchSuggestion
  Match       string
  ObjectTerm  string
  Path        string
  WktString   string
}

type Section struct {
  ContentSections []ContentSection
  ID              uint
  ParentID        uint
  Parent          Node
  Position        uint
  Name            string
}

type SeoMetum struct {
  ID              uint
  SeoMetaID       uint
  BrowserTitle    string
  MetaDescription string
  CreatedAt       time.Time
  UpdatedAt       time.Time
}

type Stylesheet struct {
  ID         uint
  ResourceID uint
  Resource   Resource
  Filename   string
}

type TaxonRemark struct {
  ID     uint
  NodeID uint
  Node   Node
  Body   string
}

type TaxonomicStatus struct {
  ScientificNames        []ScientificName
  ID                     uint
  Name                   string
  IsPreferred            bool
  IsProblematic          bool
  IsAlternativePreferred bool
  CanMerge               bool
}

type TermQuery struct {
  ID             uint
  Pairs          string
  Clade          uint
  CreatedAt      time.Time
  UpdatedAt      time.Time
  TermQueryPairs []TermQueryPair
  UserDownloads  []UserDownload
}

type TermQueryPair struct {
  ID          uint
  Predicate   string
  Object      string
  TermQueryID uint
  TermQuery   TermQuery
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type UserDownload struct {
  ID          uint
  UserID      uint
  User        User
  Count       uint
  Filename    string
  CompletedAt time.Time
  ExpiredAt   time.Time
  CreatedAt   time.Time
  UpdatedAt   time.Time
  TermQueryID uint
  TermQuery   TermQuery
}

type User struct {
  Changes             []Change
  Collectings         []Collecting
  CollectionsUsers    []CollectionsUser
  ContentEdits        []ContentEdit
  ContentRepositions  []ContentReposition
  Curations           []Curation
  OpenAuthentications []OpenAuthentication
  PageContents        []PageContent
  PageIcons           []PageIcon
  PartnersUsers       []PartnersUser
  UserDownloads       []UserDownload
  ID                  uint
  Email               string
  EncryptedPassword   string
  ResetPasswordToken  string
  ResetPasswordSentAt time.Time
  RememberCreatedAt   time.Time
  SignInCount         uint
  CurrentSignInAt     time.Time
  LastSignInAt        time.Time
  CurrentSignInIp     string
  LastSignInIp        string
  ConfirmationToken   string
  ConfirmedAt         time.Time
  ConfirmationSentAt  time.Time
  UnconfirmedEmail    string
  CreatedAt           time.Time
  UpdatedAt           time.Time
  Username            string
  Name                string
  Active              bool
  ApiKey              string
  TagLine             string
  Bio                 string
  Provider            string
  Uid                 string
  DeletedAt           time.Time
  Admin               bool
  FailedAttempts      uint
  UnlockToken         string
  LockedAt            time.Time
}

type Vernacular struct {
  ID                    uint
  String                string
  LanguageID            uint
  Language              Language
  NodeID                uint
  Node                  Node
  PageID                uint
  Page                  Page
  IsPreferred           bool
  IsPreferredByResource bool
  CreatedAt             time.Time
  UpdatedAt             time.Time
  Trust                 uint
  NodeResourcePk        string
  Locality              string
  Remarks               string
  Source                string
  ResourceID            uint
  Resource              Resource
}
