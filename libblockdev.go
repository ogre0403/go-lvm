package lvm

//#cgo pkg-config: blockdev glib-2.0
/*
#include <blockdev/blockdev.h>
#include <blockdev/lvm.h>


static GError* to_error(void* err) {
	return (GError*)err;
}

static BDLVMLVdata* to_bdlvmlvdata(void* lv){
	return (BDLVMLVdata*)lv;
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

//gboolean bd_lvm_pvcreate (const gchar *device,
//							guint64 data_alignment,
//							guint64 metadata_size,
//							const BDExtraArg **extra,
//							GError **error);
func BD_LVM_PvCreate(device string) error {
	var gerror *C.GError
	if C.bd_lvm_pvcreate(C.CString(device), C.ulong(0), C.ulong(0), nil, &gerror) == C.FALSE {
		return errors.New(strings.TrimSpace(gErrorFromNative(unsafe.Pointer(gerror)).message()))
	}
	return nil
}

//gboolean bd_lvm_pvremove (const gchar *device,
//					 	    const BDExtraArg **extra,
//				            GError **error);
func BD_LVM_PvRemove(device string) error {
	var gerror *C.GError
	if C.bd_lvm_pvremove(C.CString(device), nil, &gerror) == C.FALSE {
		return errors.New(strings.TrimSpace(gErrorFromNative(unsafe.Pointer(gerror)).message()))
	}
	return nil
}

//gboolean bd_lvm_vgcreate (const gchar *name,
//							const gchar **pv_list,
//							guint64 pe_size,
//							const BDExtraArg **extra,
//							GError **error);
func BD_LVM_VgCreate(vg string, pvs []string) error {
	// convert pv string slice to gchar**
	// https://developer.aliyun.com/article/481791
	pv_list := make([]*C.char, 0)
	for i, _ := range pvs {
		char := C.CString(pvs[i])
		defer C.free(unsafe.Pointer(char))
		strptr := (*C.char)(unsafe.Pointer(char))
		pv_list = append(pv_list, strptr)
	}


	var gerror *C.GError
	if C.bd_lvm_vgcreate(C.CString(vg), (**C.char)(unsafe.Pointer(&pv_list[0])), C.ulong(0), nil, &gerror) == C.FALSE {
		return errors.New(strings.TrimSpace(gErrorFromNative(unsafe.Pointer(gerror)).message()))
	}
	return nil
}

//gboolean bd_lvm_vgremove (const gchar *vg_name,
//							const BDExtraArg **extra,
//							GError **error);
func BD_LVM_VgRemove(vg string) error {
	var gerror *C.GError
	if C.bd_lvm_vgremove(C.CString(vg), nil, &gerror) == C.FALSE {
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
func BD_LVM_ThLvCreate(vg, pool, lv string, sizeByte uint64) error {
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

type BDLVData struct {
	Lv_name string
	Vg_name string
	Uuid    string
	Size    uint64
	Attr    string
	Segtype string
}

//BDLVMLVdata* bd_lvm_lvinfo(gchar *vg_name,
//							 gchar *lv_name,
//							 GError **error);
func BD_LVM_LvInfo(vg, lv string) (*BDLVData, error) {
	var gerror *C.GError

	lvdata := C.bd_lvm_lvinfo(C.CString(vg), C.CString(lv), &gerror)

	eee := gErrorFromNative(unsafe.Pointer(gerror))
	if unsafe.Pointer(eee.GError) != nil {
		return nil, errors.New(eee.message())
	}

	d := C.to_bdlvmlvdata(unsafe.Pointer(lvdata))
	return &BDLVData{
		Lv_name: C.GoString(C.to_charptr(d.lv_name)),
		Vg_name: C.GoString(C.to_charptr(d.vg_name)),
		Uuid:    C.GoString(C.to_charptr(d.uuid)),
		Attr:    C.GoString(C.to_charptr(d.attr)),
		Segtype: C.GoString(C.to_charptr(d.segtype)),
		Size:    uint64(C.ulonglong(d.size)),
	}, nil
}

// Verify first bit of attribute is V
//https://www.mankier.com/8/lvs
func (lv *BDLVData) IsThinVolume() bool {
	return string(lv.Attr[0]) == "V"
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
