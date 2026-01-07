package setft

import (
	"syscall"
	"time"
	"unsafe"
)

const (
	_ATTR_BIT_MAP_COUNT = 5
	_ATTR_CMN_CRTIME    = 0x00000200
	_ATTR_CMN_MODTIME   = 0x00000400
	_ATTR_CMN_ACCTIME   = 0x00001000
	_SYS_SETATTRLIST    = 221
)

type _AttrList struct {
	bitmapCount uint16
	reserved    uint16
	commonAttr  uint32
	volAttr     uint32
	dirAttr     uint32
	fileAttr    uint32
	forkattr    uint32
}

func SetFileTime(path string, atime, ctime, mtime time.Time) (err error) {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	attrList := _AttrList{
		bitmapCount: _ATTR_BIT_MAP_COUNT,
		commonAttr:  _ATTR_CMN_CRTIME | _ATTR_CMN_MODTIME | _ATTR_CMN_ACCTIME,
	}
	times := []syscall.Timespec{
		syscall.NsecToTimespec(ctime.UnixNano()),
		syscall.NsecToTimespec(mtime.UnixNano()),
		syscall.NsecToTimespec(atime.UnixNano()),
	}
	_, _, errno := syscall.Syscall6(
		_SYS_SETATTRLIST,
		uintptr(unsafe.Pointer(p)),
		uintptr(unsafe.Pointer(&attrList)),
		uintptr(unsafe.Pointer(&times[0])),
		uintptr(len(times)*int(unsafe.Sizeof(times[0]))),
		0,
		0,
	)
	if errno != 0 {
		return errno
	}
	return nil
}
