#include <stdio.h>
#include <sys/types.h>
#include <dirent.h>
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

char **paths=0;
int npaths=0;

void addToPaths(char *path){
  npaths++;
  paths=(char **)realloc(paths,npaths*sizeof(char *));
  paths[npaths-1]=path;
}

char contains(char *full,char *p){
  for(;*full && *p;full++) if(tolower(*full)==tolower(*p) && !*++p) return 1;
  return 0;
}

void find(char *path,int n,char  *sub[]){
  if(!n){
    addToPaths(path);
    return;
  }

  size_t pl=strlen(path);

  struct dirent *de;
  DIR *d=opendir(path);

  while(de=readdir(d)){
    if(de->d_type==DT_DIR && contains(de->d_name,sub[0])){
      char *dir=(char *)malloc(pl+strlen(de->d_name)+2);
      strcpy(dir,path);
      dir[pl]='/';
      strcpy(dir+pl+1,de->d_name);

      find(dir,n-1,sub+1);
    }
  }
  closedir(d);

  free(path);
}


int main(int argc, char *argv[]){
  int i=0;

  find(strdup(argv[1]),argc-2,argv+2);

  for(i=0;i<npaths;i++){
    printf("%s\n",paths[i]);
    free(paths[i]);
  }

  free(paths);

  return 0;
}
