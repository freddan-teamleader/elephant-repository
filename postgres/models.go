// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package postgres

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Acl struct {
	UUID        uuid.UUID
	URI         string
	Permissions []string
}

type AclAudit struct {
	ID         int64
	UUID       uuid.UUID
	Updated    pgtype.Timestamptz
	UpdaterUri string
	State      []byte
	Archived   bool
	Type       pgtype.Text
	Language   pgtype.Text
}

type ActiveSchema struct {
	Name    string
	Version string
}

type DeleteRecord struct {
	ID         int64
	UUID       uuid.UUID
	URI        string
	Type       string
	Version    int64
	Created    pgtype.Timestamptz
	CreatorUri string
	Meta       []byte
	MainDoc    pgtype.UUID
	Language   pgtype.Text
}

type Document struct {
	UUID           uuid.UUID
	URI            string
	Type           string
	Created        pgtype.Timestamptz
	CreatorUri     string
	Updated        pgtype.Timestamptz
	UpdaterUri     string
	CurrentVersion int64
	Deleting       bool
	MainDoc        pgtype.UUID
	Language       pgtype.Text
}

type DocumentLink struct {
	FromDocument uuid.UUID
	Version      int64
	ToDocument   uuid.UUID
	Rel          pgtype.Text
	Type         pgtype.Text
}

type DocumentLock struct {
	UUID    uuid.UUID
	Token   string
	Created pgtype.Timestamptz
	Expires pgtype.Timestamptz
	URI     pgtype.Text
	App     pgtype.Text
	Comment pgtype.Text
}

type DocumentSchema struct {
	Name    string
	Version string
	Spec    []byte
}

type DocumentStatus struct {
	UUID           uuid.UUID
	Name           string
	ID             int64
	Version        int64
	Created        pgtype.Timestamptz
	CreatorUri     string
	Meta           []byte
	Archived       bool
	Signature      pgtype.Text
	MetaDocVersion pgtype.Int8
}

type DocumentVersion struct {
	UUID         uuid.UUID
	Version      int64
	Created      pgtype.Timestamptz
	CreatorUri   string
	Meta         []byte
	DocumentData []byte
	Archived     bool
	Signature    pgtype.Text
}

type Eventlog struct {
	ID          int64
	Event       string
	UUID        uuid.UUID
	Timestamp   pgtype.Timestamptz
	Type        pgtype.Text
	Version     pgtype.Int8
	Status      pgtype.Text
	StatusID    pgtype.Int8
	Acl         []byte
	Updater     pgtype.Text
	MainDoc     pgtype.UUID
	Language    pgtype.Text
	OldLanguage pgtype.Text
}

type Eventsink struct {
	Name          string
	Position      int64
	Configuration []byte
}

type JobLock struct {
	Name      string
	Holder    string
	Touched   pgtype.Timestamptz
	Iteration int64
}

type MetaType struct {
	MetaType         string
	ExclusiveForMeta bool
}

type MetaTypeUse struct {
	MainType string
	MetaType string
}

type Metric struct {
	UUID  uuid.UUID
	Kind  string
	Label string
	Value int64
}

type MetricKind struct {
	Name        string
	Aggregation int16
}

type PlanningAssignee struct {
	Assignment uuid.UUID
	Assignee   uuid.UUID
	Version    int64
	Role       string
}

type PlanningAssignment struct {
	UUID         uuid.UUID
	Version      int64
	PlanningItem uuid.UUID
	Status       pgtype.Text
	Publish      pgtype.Timestamptz
	PublishSlot  pgtype.Int2
	Starts       pgtype.Timestamptz
	Ends         pgtype.Timestamptz
	StartDate    pgtype.Date
	EndDate      pgtype.Date
	FullDay      bool
	Public       bool
	Kind         []string
	Description  string
}

type PlanningDeliverable struct {
	Assignment uuid.UUID
	Document   uuid.UUID
	Version    int64
}

type PlanningItem struct {
	UUID        uuid.UUID
	Version     int64
	Title       string
	Description string
	Public      bool
	Tentative   bool
	StartDate   pgtype.Date
	EndDate     pgtype.Date
	Priority    pgtype.Int2
	Event       pgtype.UUID
}

type Report struct {
	Name          string
	Enabled       bool
	NextExecution pgtype.Timestamptz
	Spec          []byte
}

type SchemaVersion struct {
	Version int32
}

type SigningKey struct {
	Kid  string
	Spec []byte
}

type Status struct {
	Name     string
	Disabled bool
}

type StatusHead struct {
	UUID       uuid.UUID
	Name       string
	CurrentID  int64
	Updated    pgtype.Timestamptz
	UpdaterUri string
	Type       pgtype.Text
	Version    pgtype.Int8
	Language   pgtype.Text
}

type StatusRule struct {
	Name        string
	Description string
	AccessRule  bool
	AppliesTo   []string
	ForTypes    []string
	Expression  string
}
