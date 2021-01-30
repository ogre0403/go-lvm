package lvm

//#cgo pkg-config: blockdev glib-2.0
/*
#include <blockdev/blockdev.h>
#include <blockdev/lvm.h>
int init(GError **error){

    //GError *error = malloc(sizeof(GError));
    //error = NULL;
    //gboolean ret = FALSE;


    BDPluginSpec lvm_plugin = {BD_PLUGIN_LVM, "libbd_lvm.so.2"};
    BDPluginSpec *plugins[] = {&lvm_plugin, NULL};

    ret = bd_switch_init_checks (FALSE, error);
    if (!ret) {
        return 1;
    }

    ret = bd_ensure_init (plugins, NULL, error);
    if (!ret) {
        return 1;
    }
}
 */
import "C"
import "unsafe"

func Initialize() {
	var gerror *C.GError

	C.init(&gerror)


	//if unsafe.Pointer(gerror) != nil {
	//	err = ErrorFromNative(unsafe.Pointer(gerror))
	//} else {
	//	err = nil
	//}
}

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

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

// converts a C string array to a Go string slice
func toSlice(ar **C.gchar) []string {
	result := make([]string, 0)
	for i := 0; ; i++ {
		str := C.GoString(C.to_charptr(*ar))
		if str == "" {
			break
		}
		result = append(result, str)
		*ar = C.next_string(*ar)
	}
	return result
}