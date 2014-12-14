object = main.o stack.o link.o queue.o
test : $(object)
	gcc -o test $(object) 
stack.o : datastructure/stack.c datastructure/stack.h datastructure/node.h
	gcc -c datastructure/stack.c
link.o : datastructure/linked_list.c datastructure/linked_list.h datastructure/node.h
	gcc -c datastructure/linked_list.c -o link.o
queue.o : datastructure/queue.c datastructure/queue.h datastructure/node.h
	gcc -c datastructure/queue.c
main.o : main.c
	gcc -c main.c

clean : 
	rm test $(object)
