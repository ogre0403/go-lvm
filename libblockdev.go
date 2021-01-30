package lvm

//#cgo pkg-config: blockdev glib-2.0
/*
#include <blockdev/blockdev.h>
#include <blockdev/lvm.h>


static GError* to_error(void* err) {
	return (GError*)err;
}

static inline char* to_charptr(const gchar* s) { return (char*)s; }


int init(GError **error){
	gboolean ret = FALSE;
    BDPluginSpec lvm_plugin = {BD_PLUGIN_LVM, "libbd_lvm.so.2"};
    BDPluginSpec *plugins[] = {&lvm_plugin, NULL};

    ret = bd_switch_init_checks (FALSE, error);
    if (!ret) {
        return FALSE;
    }

    ret = bd_ensure_init (plugins, NULL, error);
    if (!ret) {
        return FALSE;
    }
}
*/
import "C"
import (
	"errors"
	"strings"
	"unsafe"
)

func Initialize() error {
	var gerror *C.GError
	if C.init(&gerror) == C.FALSE {
		return errors.New(strings.TrimSpace(gErrorFromNative(unsafe.Pointer(gerror)).message()))
	}
	return nil
}

//gboolean bd_lvm_lvcreate(const gchar *vg_name,
//						   const gchar *lv_name,
//						   guint64 size,
//						   const gchar *type,
//						   const gchar **pv_list,
//						   const BDExtraArg **extra,
//						   GError **error);
func BD_LVM_LvCreate(vg, lv string, sizeByte uint64) error {
	var gerror *C.GError
	if C.bd_lvm_lvcreate(C.CString(vg), C.CString(lv), C.ulong(sizeByte), C.CString("linear"), nil, nil, &gerror) == C.FALSE {
		return errors.New(strings.TrimSpace(gErrorFromNative(unsafe.Pointer(gerror)).message()))
	}
	return nil
}

//gboolean bd_lvm_thlvcreate(const gchar *vg_name,
//							 const gchar *pool_name,
//							 const gchar *lv_name,
//							 guint64 size,
//							 const BDExtraArg **extra,
//							 GError **error);
func BD_LVM_ThLVCreate(vg, pool, lv string, sizeByte uint64) error {
	var gerror *C.GError
	if C.bd_lvm_thlvcreate(C.CString(vg), C.CString(pool), C.CString(lv), C.ulong(sizeByte), nil, &gerror) == C.FALSE {
		return errors.New(strings.TrimSpace(gErrorFromNative(unsafe.Pointer(gerror)).message()))
	}
	return nil
}

// gboolean bd_lvm_lvremove(const gchar *vg_name,
//                          const gchar *lv_name,
//                          gboolean force,
//                          const BDExtraArg **extra,
//                          GError **error);
func BD_LVM_LvRemove(vg, lv string) error {
	var gerror *C.GError
	if C.bd_lvm_lvremove(C.CString(vg), C.CString(lv), C.FALSE, nil, &gerror) == C.FALSE {
		return errors.New(strings.TrimSpace(gErrorFromNative(unsafe.Pointer(gerror)).message()))
	}
	return nil
}

type gError struct {
	GError *C.GError
}

func (v *gError) message() string {
	if unsafe.Pointer(v.GError) == nil || unsafe.Pointer(v.GError.message) == nil {
		return ""
	}
	return C.GoString(C.to_charptr(v.GError.message))
}

func gErrorFromNative(err unsafe.Pointer) *gError {
	return &gError{
		C.to_error(err)}
}
