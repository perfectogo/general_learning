let nestedArray = [1, [2, 3], 4];
let [x, [y, z], w] = nestedArray;
console.log(x); // Output: 1
console.log(y); // Output: 2
console.log(z); // Output: 3
console.log(typeof w); // Output: 4