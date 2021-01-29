#include <blockdev/blockdev.h>
#include <blockdev/lvm.h>
#include <stdio.h>

#define MIB * 1024 * 1024
#define GIB * 1024 MIB


int init(){

    printf("Init\n");

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

    printf("Create LV\n");
    GError *error = malloc(sizeof(GError));
    error = NULL;


    gboolean ret = FALSE;

    ret = bd_lvm_lvcreate(vg, lv, mb_size ,"linear", NULL,NULL, &error);
    if (!ret) {
        return 1;
    }

}




int main (int argc, char *argv[]) {

    init();
    createLV("vg-0","ttt",20 MB);
    createLV("vg-0","aaa",30 MB);

    return 0;
}