class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

const a = new Node("A");
const b = new Node("B");
const c = new Node("C");
const d = new Node("D");

a.Next = b;
b.Next = c;
c.Next = d;

 
const PrintList = (head) =>   {
    let current = head;

    while (current) {
        console.log(current.val);
        current = current.Next;
    }
    
}

PrintList(a);
