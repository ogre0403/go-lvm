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

int init2(GError **error){

    printf("Init\n");
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



int createLV2(const gchar *vg, const gchar *lv, int mb_size, GError **error){

    printf("Create LV\n");


    gboolean ret = FALSE;

    ret = bd_lvm_lvcreate(vg, lv, mb_size ,"linear", NULL,NULL, error);
    if (!ret) {
        return FALSE;
    }

}


int createPV(const gchar *pv){

    printf("Create PV\n");
    GError *error = malloc(sizeof(GError));
    error = NULL;
    gboolean ret = FALSE;

    ret = bd_lvm_pvcreate(pv,
                             0,    /* data alignment (first PE), 0 for default */
                             0,    /* size reserved for metadata, 0 for default */
                             NULL, /* extra options passed to the lvm tool */
                             &error);
      if (!ret) {
        printf("%s",error->message);
        return 1;
      }

}


int createPV2(const gchar *pv, GError **error){

    printf("Create PV\n");
    gboolean ret = FALSE;

    ret = bd_lvm_pvcreate(pv,
                             0,    /* data alignment (first PE), 0 for default */
                             0,    /* size reserved for metadata, 0 for default */
                             NULL, /* extra options passed to the lvm tool */
                             error);
      if (!ret) {
        printf("%s",(*error)->message);
        return FALSE;
      }

}

int lvInfo(const gchar *vg, const gchar *lv, GError **error){

    BDLVMLVdata*  lv_data = malloc(sizeof(BDLVMLVdata));
    lv_data = NULL;

    lv_data = bd_lvm_lvinfo("vg-0","ttt",error);

    printf("%s\n",lv_data->attr);
    printf("%s\n",lv_data->segtype);
    printf("%d\n",lv_data->size);
    printf("%s\n",lv_data->uuid);
}

int main (int argc, char *argv[]) {

    //init();
    //createPV("/dev/loop2")
    //createLV("vg-0","ttt",20 MB);
    //createLV("vg-0","aaa",30 MB);


    GError *error = malloc(sizeof(GError));
    error = NULL;
    gboolean r = FALSE;

    r = init2(&error);
    if(!r){
        printf("main: %s",error->message);
    }

    r = createPV2("/dev/loop2", &error);
    if(!r){
        printf("main: %s",error->message);
    }

    r = createLV2("vg-0","ttt",20 MB, &error);
    if(!r){
        printf("main: %s",error->message);
    }

    lvInfo("vg-0","ttt",&error);

    return 0;
}
