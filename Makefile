object = main.o stack.o link.o queue.o random.o tree.o sort.o dynamic.o unix_test.o
test : $(object)
	gcc -o test $(object) 
stack.o : datastructure/stack.c datastructure/stack.h datastructure/node.h
	gcc -c datastructure/stack.c
link.o : datastructure/linked_list.c datastructure/linked_list.h datastructure/node.h
	gcc -c datastructure/linked_list.c -o link.o
queue.o : datastructure/queue.c datastructure/queue.h datastructure/node.h
	gcc -c datastructure/queue.c
tree.o : datastructure/tree.c datastructure/tree.h
	gcc -c datastructure/tree.c
random.o : algorithm/random.c algorithm/random.h
	gcc -c algorithm/random.c 
sort.o : algorithm/sort.c algorithm/sort.h
	gcc -c algorithm/sort.c 
dynamic.o : algorithm/dynamic.c algorithm/dynamic.h
	gcc -c algorithm/dynamic.c
unix_test.o : unix/unix.c unix/unix.h
	gcc -c unix/unix.c -o unix_test.o
main.o : main.c
	gcc -c main.c

clean : 
	rm $(object)
