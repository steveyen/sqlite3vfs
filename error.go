package sqlite3vfs

import "fmt"

type SqliteError struct {
	code int
	text string
}

func (e SqliteError) Error() string {
	return fmt.Sprintf("sqlite (%d) %s", e.code, e.text)
}

// https://www.sqlite.org/rescode.html

const (
	SqliteOK = 0
)

var (
	GenericError    = SqliteError{1, "Generic Error"}
	InternalError   = SqliteError{2, "Internal Error"}
	PermError       = SqliteError{3, "Perm Error"}
	AbortError      = SqliteError{4, "Abort Error"}
	BusyError       = SqliteError{5, "Busy Error"}
	LockedError     = SqliteError{6, "Locked Error"}
	NoMemError      = SqliteError{7, "No Mem Error"}
	ReadOnlyError   = SqliteError{8, "Read Only Error"}
	InterruptError  = SqliteError{9, "Interrupt Error"}
	IOError         = SqliteError{10, "IO Error"}
	CorruptError    = SqliteError{11, "Corrupt Error"}
	NotFoundError   = SqliteError{12, "Not Found Error"}
	FullError       = SqliteError{13, "Full Error"}
	CantOpenError   = SqliteError{14, "CantOpen Error"}
	ProtocolError   = SqliteError{15, "Protocol Error"}
	EmptyError      = SqliteError{16, "Empty Error"}
	SchemaError     = SqliteError{17, "Schema Error"}
	TooBigError     = SqliteError{18, "TooBig Error"}
	ConstraintError = SqliteError{19, "Constraint Error"}
	MismatchError   = SqliteError{20, "Mismatch Error"}
	MisuseError     = SqliteError{21, "Misuse Error"}
	NoLFSError      = SqliteError{22, "No Large File Support Error"}
	AuthError       = SqliteError{23, "Auth Error"}
	FormatError     = SqliteError{24, "Format Error"}
	RangeError      = SqliteError{25, "Range Error"}
	NotaDBError     = SqliteError{26, "Not a DB Error"}
	NoticeError     = SqliteError{27, "Notice Error"}
	WarningError    = SqliteError{28, "Warning Error"}

	IOErrorRead      = SqliteError{266, "IO Error Read"}
	IOErrorShortRead = SqliteError{522, "IO Error Short Read"}
	IOErrorWrite     = SqliteError{778, "IO Error Write"}
)

var ErrMap = map[int]SqliteError{
	1:  GenericError,
	2:  InternalError,
	3:  PermError,
	4:  AbortError,
	5:  BusyError,
	6:  LockedError,
	7:  NoMemError,
	8:  ReadOnlyError,
	9:  InterruptError,
	10: IOError,
	11: CorruptError,
	12: NotFoundError,
	13: FullError,
	14: CantOpenError,
	15: ProtocolError,
	16: EmptyError,
	17: SchemaError,
	18: TooBigError,
	19: ConstraintError,
	20: MismatchError,
	21: MisuseError,
	22: NoLFSError,
	23: AuthError,
	24: FormatError,
	25: RangeError,
	26: NotaDBError,
	27: NoticeError,
	28: WarningError,

	266: IOErrorRead,
	522: IOErrorShortRead,
	778: IOErrorWrite,
}

func ErrFromCode(code int) error {
	if code == 0 {
		return nil
	}

	err, ok := ErrMap[code]
	if ok {
		return err
	}

	return SqliteError{
		code: code,
		text: "unknown err code",
	}
}
