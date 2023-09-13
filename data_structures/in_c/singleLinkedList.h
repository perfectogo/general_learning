#include <stdio.h>
#include <stdlib.h>

// A linked list node
struct Node {
    int data;
    struct Node* next;
};


void insertAtFront(struct Node** head_ref, int new_data) {
    
    struct Node* new_node;
    new_node = (struct Node*)malloc(sizeof(struct Node));

    new_node->data = new_data;

    new_node->next = (*head_ref);

    (*head_ref) = new_node;
}

// Function to insert element in LL
void push(struct Node** head_ref, int new_data) {
    
    struct Node* new_node;
    new_node = (struct Node*)malloc(sizeof(struct Node));
    
    new_node->data = new_data;
    
    new_node->next = (*head_ref);
    
    (*head_ref) = new_node;
}

void printList(struct Node* node){
    
    while (node != NULL) {
        printf(" %d", node->data);
        node = node->next;
    }

    printf("\n");
}

void work(void) {
 // Start with the empty list
    struct Node* head = NULL;

    push(&head, 6);

    printf("Created Linked list is: ");
    printList(head);

    // Insert 1 at the beginning.
    insertAtFront(&head, 1);

    printf("After inserting 1 at front: ");
    printList(head);
}