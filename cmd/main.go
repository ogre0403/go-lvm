package main

//#cgo pkg-config: blockdev glib-2.0
/*
#include <blockdev/blockdev.h>
#include <blockdev/lvm.h>
//#include <stdio.h>
//#include <stdlib.h>


static GError* to_error(void* err) {
	return (GError*)err;
}

static inline char* to_charptr(const gchar* s) { return (char*)s; }


int init(){

    GError *error = malloc(sizeof(GError));
    error = NULL;
    gboolean ret = FALSE;


    BDPluginSpec lvm_plugin = {BD_PLUGIN_LVM, "libbd_lvm.so.2"};
    BDPluginSpec *plugins[] = {&lvm_plugin, NULL};

    ret = bd_switch_init_checks (FALSE, &error);
    if (!ret) {
        return 1;
    }

    ret = bd_ensure_init (plugins, NULL, &error);
    if (!ret) {
        return 1;
    }
}

int createLV(const gchar *vg, const gchar *lv, int mb_size){

    GError *error = malloc(sizeof(GError));
    error = NULL;


    gboolean ret = FALSE;

    ret = bd_lvm_lvcreate(vg, lv, mb_size ,"linear", NULL,NULL, &error);
    if (!ret) {
        return 1;
    }

}

int createLV2(const gchar *vg, const gchar *lv, int mb_size, GError **error){

    gboolean ret = FALSE;
    ret = bd_lvm_lvcreate(vg, lv, mb_size ,"linear", NULL,NULL, error);
    if (!ret) {
        return FALSE;
    }
}



*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Error struct {
	GError *C.GError
}

func (v *Error) Error() string {
	return v.Message()
}

func (v *Error) Message() string {
	if unsafe.Pointer(v.GError) == nil || unsafe.Pointer(v.GError.message) == nil {
		return ""
	}
	return C.GoString(C.to_charptr(v.GError.message))
}

func ErrorFromNative(err unsafe.Pointer) *Error {
	return &Error{
		C.to_error(err)}
}

func main() {
	fmt.Println("Call Init")
	C.init()

	fmt.Println("Call from go wrapper  w/o error code")
	C.createLV(C.CString("vg-0"), C.CString("aa"), C.int(10485760))

	var gerror1 *C.GError
	fmt.Println("Call from go wrapper  w/ error code")

	rr := C.createLV2(C.CString("vg-0"), C.CString("mmm"), C.int(10485760), &gerror1)
	if rr == C.FALSE {
		ee := ErrorFromNative(unsafe.Pointer(gerror1))
		fmt.Println(ee.Message())
	}

	fmt.Println("Call from go wrapper  w/ error code")
	var gerror2 *C.GError
	rr = C.createLV2(C.CString("vg-0"), C.CString("mmm"), C.int(10485760), &gerror2)

	if rr == C.FALSE {
		ee := ErrorFromNative(unsafe.Pointer(gerror2))
		fmt.Println(ee.Message())
	}

	fmt.Println("Call without  wrapper  w/ error code")
	var gerror3 *C.GError
	rr = C.bd_lvm_lvcreate(C.CString("vg-0"), C.CString("mm1m"), C.ulong(10485760),
		C.CString("linear"),
		nil, nil, &gerror3)

	if rr == C.FALSE {
		ee := ErrorFromNative(unsafe.Pointer(gerror3))
		fmt.Println(ee.Message())
	}

}
