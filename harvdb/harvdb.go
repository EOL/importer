package importer

import (
  "strings"
  "time"
)

type Agent struct {
  ID         uint
  HarvestID  uint
  Harvest    Harvest
  ResourcePk string
  FullName   string
  Role       string
  Email      string
  Uri        string
  OtherInfo  string
}

type Article struct {
  ID                      uint
  Guid                    string
  ResourcePk              string
  LanguageCodeVerbatim    string
  ResourceID              uint
  Resource                Resource
  HarvestID               uint
  Harvest                 Harvest
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
  RemovedByHarvestID      uint
  CreatedAt               time.Time
  UpdatedAt               time.Time
  ArticlesReferences      []ArticlesReference
  ArticlesSections        []ArticlesSection
}

type ArticlesReference struct {
  ID                uint
  HarvestID         uint
  Harvest           Harvest
  ArticleID         uint
  Article           Article
  ReferenceID       uint
  Reference         Reference
  RefResourceFk     string
  ArticleResourceFk string
}

type ArticlesSection struct {
  ArticleID uint
  Article   Article
  SectionID uint
  Section   Section
}

type AssocTrait struct {
  ID                      uint
  ResourceID              uint
  Resource                Resource
  HarvestID               uint
  Harvest                 Harvest
  TraitID                 uint
  Trait                   Trait
  PredicateTermID         uint
  Term                    Term
  ObjectTermID            uint
  UnitsTermID             uint
  StatisticalMethodTermID uint
  RemovedByHarvestID      uint
  TraitResourcePk         string
  Measurement             string
  Literal                 string
  Source                  string
}

type Assoc struct {
  ID                         uint
  ResourceID                 uint
  Resource                   Resource
  HarvestID                  uint
  Harvest                    Harvest
  RemovedByHarvestID         uint
  PredicateTermID            uint
  Term                       Term
  NodeID                     uint
  Node                       Node
  TargetNodeID               uint
  SexTermID                  uint
  LifestageTermID            uint
  ResourcePk                 string
  OccurrenceResourceFk       string
  TargetOccurrenceResourceFk string
  Source                     string
  CreatedAt                  time.Time
  UpdatedAt                  time.Time
  OccurrenceID               uint
  Occurrence                 Occurrence
  TargetOccurrenceID         uint
  AssocsReferences           []AssocsReference
  MetaAssocs                 []MetaAssoc
}

type AssocsReference struct {
  ID              uint
  HarvestID       uint
  Harvest         Harvest
  AssocID         uint
  Assoc           Assoc
  ReferenceID     uint
  Reference       Reference
  RefResourceFk   string
  AssocResourceFk string
}

type Attribution struct {
  ID                   uint
  ResourceID           uint
  Resource             Resource
  HarvestID            uint
  Harvest              Harvest
  ResourcePk           string
  Name                 string
  Email                string
  Value                string
  RemovedByHarvestID   uint
  CreatedAt            time.Time
  UpdatedAt            time.Time
  AttributionsContents []AttributionsContent
}

type AttributionsContent struct {
  ID            uint
  AttributionID uint
  Attribution   Attribution
  RoleID        uint
  Role          Role
}

type BibliographicCitation struct {
  Articles  []Article
  ID        uint
  Body      string
  CreatedAt time.Time
  UpdatedAt time.Time
  Media     []Medium
}

type DataReference struct {
  ID          uint
  ReferenceID uint
  Reference   Reference
  DataID      uint
}

type Dataset struct {
  ID              string
  Name            string
  Link            string
  Publisher       string
  Supplier        string
  Metadata        string
  ScientificNames []ScientificName
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

type Field struct {
  ID              uint
  FormatID        uint
  Format          Format
  Position        uint
  Validation      uint
  Mapping         uint
  SpecialHandling uint
  Submapping      string
  ExpectedHeader  string
  UniqueInFormat  bool
  CanBeEmpty      bool
}

type Format struct {
  Fields           []Field
  ID               uint
  ResourceID       uint
  Resource         Resource
  HarvestID        uint
  Harvest          Harvest
  Sheet            uint
  HeaderLines      uint
  DataBeginsOnLine uint
  Represents       uint
  GetFrom          string
  File             string
  Diff             string
  FieldSep         string
  LineSep          string
  Utf8             bool
  Hlogs            []Hlog
}

type Harvest struct {
  Agents                    []Agent
  Articles                  []Article
  ArticlesReferences        []ArticlesReference
  AssocTraits               []AssocTrait
  Assocs                    []Assoc
  AssocsReferences          []AssocsReference
  Attributions              []Attribution
  Formats                   []Format
  ID                        uint
  ResourceID                uint
  Resource                  Resource
  TimeInMinutes             uint
  Hold                      bool
  FetchedAt                 time.Time
  ValidatedAt               time.Time
  DeltasCreatedAt           time.Time
  StoredAt                  time.Time
  ConsistencyCheckedAt      time.Time
  NamesParsedAt             time.Time
  NodesMatchedAt            time.Time
  AncestryBuiltAt           time.Time
  UnitsNormalizedAt         time.Time
  LinkedAt                  time.Time
  IndexedAt                 time.Time
  FailedAt                  time.Time
  CompletedAt               time.Time
  CreatedAt                 time.Time
  UpdatedAt                 time.Time
  Stage                     uint
  Hlogs                     []Hlog
  Identifiers               []Identifier
  Links                     []Link
  Media                     []Medium
  MediaReferences           []MediaReference
  MetaAssocs                []MetaAssoc
  MetaTraits                []MetaTrait
  Nodes                     []Node
  NodesReferences           []NodesReference
  OccurrenceMetadata        []OccurrenceMetadatum
  Occurrences               []Occurrence
  References                []Reference
  ScientificNames           []ScientificName
  ScientificNamesReferences []ScientificNamesReference
  Traits                    []Trait
  TraitsReferences          []TraitsReference
  Vernaculars               []Vernacular
}

type Hlog struct {
  ID        uint
  HarvestID uint
  Harvest   Harvest
  FormatID  uint
  Format    Format
  Category  uint
  Message   string
  Backtrace string
  Line      uint
  CreatedAt time.Time
}

type Identifier struct {
  ID             uint
  ResourceID     uint
  Resource       Resource
  HarvestID      uint
  Harvest        Harvest
  NodeID         uint
  Node           Node
  IDentifier     string
  NodeResourcePk string
}

type Language struct {
  Articles    []Article
  ID          uint
  Code        string
  GroupCode   string
  Links       []Link
  Media       []Medium
  Resources   []Resource
  Vernaculars []Vernacular
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
  ID                   uint
  Guid                 string
  ResourcePk           string
  LanguageCodeVerbatim string
  ResourceID           uint
  Resource             Resource
  HarvestID            uint
  Harvest              Harvest
  LanguageID           uint
  Language             Language
  Name                 string
  SourceUrl            string
  Description          string
  RemovedByHarvestID   uint
  CreatedAt            time.Time
  UpdatedAt            time.Time
  LinksSections        []LinksSection
}

type LinksSection struct {
  LinkID    uint
  Link      Link
  SectionID uint
  Section   Section
}

type Location struct {
  Articles    []Article
  ID          uint
  LatLiteral  string
  LongLiteral string
  AltLiteral  string
  Locality    string
  Created     string
  Lat         float32
  Long        float32
  Alt         float32
  CreatedAt   time.Time
  UpdatedAt   time.Time
  Media       []Medium
}

type Medium struct {
  ID                      uint
  Guid                    string
  ResourcePk              string
  NodeResourcePk          string
  UnmodifiedUrl           string
  NameVerbatim            string
  Name                    string
  SourcePageUrl           string
  SourceUrl               string
  BaseUrl                 string
  RightsStatement         string
  UsageStatement          string
  Sizes                   string
  BibliographicCitationFk string
  LanguageCodeVerbatim    string
  Subclass                uint
  Format                  uint
  ResourceID              uint
  Resource                Resource
  HarvestID               uint
  Harvest                 Harvest
  NodeID                  uint
  Node                    Node
  LicenseID               uint
  License                 License
  LanguageID              uint
  Language                Language
  LocationID              uint
  Location                Location
  W                       uint
  H                       uint
  CropXPct                uint
  CropYPct                uint
  CropWPct                uint
  CropHPct                uint
  BibliographicCitationID uint
  BibliographicCitation   BibliographicCitation
  Owner                   string
  DescriptionVerbatim     string
  Description             string
  DerivedFrom             string
  RemovedByHarvestID      uint
  DownloadedAt            time.Time
  CreatedAt               time.Time
  UpdatedAt               time.Time
  MediaReferences         []MediaReference
  MediaSections           []MediaSection
}

type MediaDownloadError struct {
  ID        uint
  Message   string
  CreatedAt time.Time
  UpdatedAt time.Time
}

type MediaReference struct {
  ID               uint
  HarvestID        uint
  Harvest          Harvest
  MediumID         uint
  Medium           Medium
  ReferenceID      uint
  Reference        Reference
  RefResourceFk    string
  MediumResourceFk string
}

type MediaSection struct {
  MediumID  uint
  Medium    Medium
  SectionID uint
  Section   Section
}

type MetaAssoc struct {
  ID                      uint
  ResourceID              uint
  Resource                Resource
  HarvestID               uint
  Harvest                 Harvest
  RemovedByHarvestID      uint
  AssocID                 uint
  Assoc                   Assoc
  PredicateTermID         uint
  Term                    Term
  ObjectTermID            uint
  UnitsTermID             uint
  StatisticalMethodTermID uint
  AssocResourceFk         string
  Measurement             string
  Literal                 string
  Source                  string
}

type MetaTrait struct {
  ID                      uint
  ResourceID              uint
  Resource                Resource
  HarvestID               uint
  Harvest                 Harvest
  RemovedByHarvestID      uint
  TraitID                 uint
  Trait                   Trait
  PredicateTermID         uint
  Term                    Term
  ObjectTermID            uint
  UnitsTermID             uint
  StatisticalMethodTermID uint
  TraitResourcePk         string
  Measurement             string
  Literal                 string
  Source                  string
}

type MetaXmlField struct {
  ID         uint
  Term       string
  ForFormat  string
  Represents string
  Submapping string
  IsUnique   bool
  IsRequired bool
}

type NodeAncestor struct {
  ID         uint
  ResourceID uint
  Resource   Resource
  NodeID     uint
  Node       Node
  AncestorID uint
  Ancestor   Node
  Depth      uint
  AncestorFk string
}

type Node struct {
  Assocs                  []Assoc
  Identifiers             []Identifier
  Media                   []Medium
  NodeAncestors           []NodeAncestor
  ID                      uint
  ResourceID              uint
  Resource                Resource
  HarvestID               uint
  Harvest                 Harvest
  PageID                  uint
  Page                    Page
  ParentID                uint
  ScientificNameID        uint
  ScientificName          ScientificName
  RemovedByHarvestID      uint
  Landmark                uint
  Canonical               string
  TaxonomicStatusVerbatim string
  ResourcePk              string
  ParentResourcePk        string
  FurtherInformationUrl   string
  Rank                    string
  RankVerbatim            string
  InUnmappedArea          bool
  CreatedAt               time.Time
  UpdatedAt               time.Time
  NodesReferences         []NodesReference
  Occurrences             []Occurrence
  Traits                  []Trait
  Vernaculars             []Vernacular
}

type NodesReference struct {
  ID             uint
  HarvestID      uint
  Harvest        Harvest
  NodeID         uint
  Node           Node
  ReferenceID    uint
  Reference      Reference
  RefResourceFk  string
  NodeResourceFk string
}

type OccurrenceMetadatum struct {
  ID                      uint
  HarvestID               uint
  Harvest                 Harvest
  OccurrenceID            uint
  Occurrence              Occurrence
  PredicateTermID         uint
  Term                    Term
  ObjectTermID            uint
  Literal                 string
  ResourceID              uint
  Resource                Resource
  UnitsTermID             uint
  StatisticalMethodTermID uint
  ResourcePk              string
  Measurement             string
  OccurrenceResourcePk    string
}

type Occurrence struct {
  Assocs             []Assoc
  OccurrenceMetadata []OccurrenceMetadatum
  ID                 uint
  HarvestID          uint
  Harvest            Harvest
  ResourcePk         string
  NodeID             uint
  Node               Node
  NodeResourcePk     string
  SexTermID          string
  Term               Term
  LifestageTermID    string
  Traits             []Trait
}

type Page struct {
  Nodes        []Node
  ID           uint
  NativeNodeID uint
}

type Partner struct {
  ID            uint
  Name          string
  Abbr          string
  ShortName     string
  HomepageUrl   string
  Description   string
  LinksJson     string
  AutoPublish   bool
  NotTrusted    bool
  CreatedAt     time.Time
  UpdatedAt     time.Time
  PartnersUsers []PartnersUser
  Resources     []Resource
}

type PartnersUser struct {
  UserID    uint
  User      User
  PartnerID uint
  Partner   Partner
}

type Reference struct {
  ArticlesReferences        []ArticlesReference
  AssocsReferences          []AssocsReference
  DataReferences            []DataReference
  MediaReferences           []MediaReference
  NodesReferences           []NodesReference
  ID                        uint
  Body                      string
  ResourceID                uint
  Resource                  Resource
  HarvestID                 uint
  Harvest                   Harvest
  ResourcePk                string
  Url                       string
  Doi                       string
  RemovedByHarvestID        uint
  CreatedAt                 time.Time
  UpdatedAt                 time.Time
  ScientificNamesReferences []ScientificNamesReference
  TraitsReferences          []TraitsReference
}

type Resource struct {
  Articles               []Article
  AssocTraits            []AssocTrait
  Assocs                 []Assoc
  Attributions           []Attribution
  Formats                []Format
  Harvests               []Harvest
  Identifiers            []Identifier
  Links                  []Link
  Media                  []Medium
  MetaAssocs             []MetaAssoc
  MetaTraits             []MetaTrait
  NodeAncestors          []NodeAncestor
  Nodes                  []Node
  OccurrenceMetadata     []OccurrenceMetadatum
  References             []Reference
  ID                     uint
  Position               uint
  MinDaysBetweenHarvests uint
  HarvestDayOfMonth      uint
  NodesCount             uint
  PartnerID              uint
  Partner                Partner
  HarvestMonthsJson      string
  Name                   string
  Abbr                   string
  PkUrl                  string
  AutoPublish            bool
  NotTrusted             bool
  HoldHarvesting         bool
  MightHaveDuplicateTaxa bool
  ForceHarvest           bool
  CreatedAt              time.Time
  UpdatedAt              time.Time
  Description            string
  Notes                  string
  IsBrowsable            bool
  DefaultLanguageID      uint
  Language               Language
  DefaultLicenseID       uint
  License                License
  DefaultRightsStatement string
  DefaultRightsHolder    string
  PublishStatus          uint
  DatasetLicenseID       uint
  DatasetRightsHolder    string
  DatasetRightsStatement string
  OpendataUrl            string
  ScientificNames        []ScientificName
  Traits                 []Trait
  Vernaculars            []Vernacular
}

type Role struct {
  AttributionsContents []AttributionsContent
  ID                   uint
  Name                 string
  CreatedAt            time.Time
  UpdatedAt            time.Time
}

type ScientificName struct {
  Nodes                     []Node
  ID                        uint
  ResourceID                uint
  Resource                  Resource
  HarvestID                 uint
  Harvest                   Harvest
  NodeID                    uint
  NormalizedNameID          uint
  ParseQuality              uint
  TaxonomicStatus           uint
  NodeResourcePk            string
  TaxonomicStatusVerbatim   string
  Warnings                  string
  Genus                     string
  SpecificEpithet           string
  InfraspecificEpithet      string
  InfragenericEpithet       string
  Normalized                string
  Canonical                 string
  Uninomial                 string
  Verbatim                  string
  Authorship                string
  Publication               string
  Remarks                   string
  Year                      uint
  IsPreferred               bool
  IsUsedForMerges           bool
  IsPublishable             bool
  Hybrid                    bool
  Surrogate                 bool
  Virus                     bool
  RemovedByHarvestID        uint
  DatasetID                 string
  Dataset                   Dataset
  CreatedAt                 time.Time
  UpdatedAt                 time.Time
  ResourcePk                string
  ScientificNamesReferences []ScientificNamesReference
}

func (name ScientificName) AttributionHTML (attribution string) {
  if harvNode.ResourceID == 1 {
    attribution = fmt.Sprintf("Reference taxon: %v",
      strings.Join([]string{scientificNameLink, accordingTo, viaStatement}, " "))
  }
  return attribution
}

func (name ScientificName) scientificNameLink (string) {
  if strings.TrimSpace(name.Node.FurtherInformationUrl) > 0 {
    return "<a href='"+name.Node.FurtherInformationUrl+"'>"+name.Normalized+"</a>"
  }
  return nil
}

func (name ScientificName) accordingTo (string) {
  if name.Dataset == nil { return nil }
  return "according to <a href='"+name.Dataset.Link+"'>"+name.Dataset.Name+"</a>"
}

func (name ScientificName) viaStatement (string) {
  if name.Dataset != nil && !(strings.TrimSpace(name.Dataset.Pulisher) == 0 && strings.TrimSpace(name.Dataset.supplier) == 0) {
    return "via #{dataset.publisher}#{'/' if !dataset.publisher.blank? && !dataset.supplier.blank?}#{dataset.supplier}"
  }
  if strings.TrimSpace(name.Publication) == 0 { return nil }
  return "via "+name.Publication
}

type ScientificNamesReference struct {
  ID               uint
  HarvestID        uint
  Harvest          Harvest
  ScientificNameID uint
  ScientificName   ScientificName
  ReferenceID      uint
  Reference        Reference
  RefResourceFk    string
  NameResourceFk   string
}

type Section struct {
  ArticlesSections []ArticlesSection
  LinksSections    []LinksSection
  MediaSections    []MediaSection
  ID               uint
  Name             string
  SectionsTerms    []SectionsTerm
}

type SectionsTerm struct {
  SectionID uint
  Section   Section
  TermID    uint
  Term      Term
}

type Term struct {
  AssocTraits            []AssocTrait
  Assocs                 []Assoc
  MetaAssocs             []MetaAssoc
  MetaTraits             []MetaTrait
  OccurrenceMetadata     []OccurrenceMetadatum
  Occurrences            []Occurrence
  SectionsTerms          []SectionsTerm
  ID                     uint
  Uri                    string
  Name                   string
  Definition             string
  Comment                string
  Attribution            string
  IsHiddenFromOverview   bool
  IsHiddenFromGlossary   bool
  CreatedAt              time.Time
  UpdatedAt              time.Time
  OntologyInformationUrl string
  OntologySourceUrl      string
  IsTextOnly             bool
  IsVerbatimOnly         bool
  Position               uint
  UsedFor                uint
  Traits                 []Trait
  UnitConversions        []UnitConversion
}

type Trait struct {
  AssocTraits             []AssocTrait
  MetaTraits              []MetaTrait
  ID                      uint
  ResourceID              uint
  Resource                Resource
  ParentID                uint
  Parent                  Node
  HarvestID               uint
  Harvest                 Harvest
  NodeID                  uint
  Node                    Node
  PredicateTermID         uint
  Term                    Term
  ObjectTermID            uint
  UnitsTermID             uint
  StatisticalMethodTermID uint
  SexTermID               uint
  LifestageTermID         uint
  RemovedByHarvestID      uint
  OfTaxon                 bool
  OccurrenceResourcePk    string
  AssocResourcePk         string
  ParentPk                string
  ResourcePk              string
  Measurement             string
  Literal                 string
  Source                  string
  CreatedAt               time.Time
  UpdatedAt               time.Time
  OccurrenceID            uint
  Occurrence              Occurrence
  TraitsReferences        []TraitsReference
}

type TraitsReference struct {
  ID              uint
  HarvestID       uint
  Harvest         Harvest
  TraitID         uint
  Trait           Trait
  ReferenceID     uint
  Reference       Reference
  RefResourceFk   string
  TraitResourceFk string
}

type UnitConversion struct {
  ID         uint
  FromTermID uint
  Term       Term
  ToTermID   uint
  Method     string
}

type User struct {
  PartnersUsers []PartnersUser
  ID            uint
  Name          string
  Description   string
}

type Vernacular struct {
  ID                   uint
  ResourceID           uint
  Resource             Resource
  HarvestID            uint
  Harvest              Harvest
  NodeID               uint
  Node                 Node
  LanguageID           uint
  Language             Language
  NodeResourcePk       string
  Verbatim             string
  LanguageCodeVerbatim string
  Locality             string
  Remarks              string
  Source               string
  IsPreferred          bool
  RemovedByHarvestID   uint
  CreatedAt            time.Time
  UpdatedAt            time.Time
}
