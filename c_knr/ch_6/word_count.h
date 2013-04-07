struct tnode {
  char *word;
  int count;
  struct tnode *left;
  struct tnode *right;
};

struct tnode *addtree(struct tnode *p, char *word);
void treeprint(struct tnode *p);
int get_word(char *word, int lim);
struct tnode *tallc(void);
char *str_dup(char *s);
struct tnode *talloc(void);
