object = main.o stack.o link.o queue.o tree.o sort.o acm.o unix_test.o
test : $(object)
	gcc -o run $(object) -g
stack.o : datastructure/stack.c datastructure/stack.h datastructure/node.h
	gcc -c datastructure/stack.c -g
link.o : datastructure/linked_list.c datastructure/linked_list.h datastructure/node.h
	gcc -c datastructure/linked_list.c -o link.o -g
queue.o : datastructure/queue.c datastructure/queue.h datastructure/node.h
	gcc -c datastructure/queue.c -g
tree.o : datastructure/tree.c datastructure/tree.h
	gcc -c datastructure/tree.c -g
sort.o : algorithm/sort.c algorithm/sort.h
	gcc -c algorithm/sort.c -g
acm.o : algorithm/acm.c algorithm/acm.h
	gcc -c algorithm/acm.c -g
unix_test.o : unix/unix.c unix/unix.h
	gcc -c unix/unix.c -o unix_test.o -g
main.o : main.c
	gcc -c main.c -g

clean : 
	rm -f *.o run
