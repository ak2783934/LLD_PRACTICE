package filesystem

import "time"

type User struct {
	ID     string
	Name   string
	Email  string
	Files  *[]File
	Shared *[]SharedFile
}

type File struct {
	Id          string
	Name        string
	OwnerId     string
	Size        int64
	Path        string
	Permissions *[]Permission
	Versions    *[]Version
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDirectory bool
}

type SharedFile struct {
	FileID     string
	SharedBy   string
	SharedWith string
	AccessType AccessType
	SharedAt   time.Time
}

type Permission struct {
	UserID     string
	FileID     string
	AccessType AccessType
}

type Version struct {
	VersionID string
	FileID    string
	CreatedAt time.Time
	Path      string
	Size      int64
}

type AccessType int

const (
	Read AccessType = iota
	Write
	Owner
)

/*


File system


root folder.

all other folders.
Each folder has owners. And shared with whom.


Root folder should list all the owner folders.
Shared folder should have list of all things shared with that person.

What other entity you think we should have?



*/
