const groceryList = ['orange juice', 'bananas', 'coffee beans', 'brown rice', 'pasta', 'coconut oil', 'plantains'];

groceryList.shift(); // Use the .shift() method to remove the first item from the array groceryList.

console.log(groceryList);

groceryList.unshift('popcorn'); // Under the code added in step 1, use the .unshift() method to add 'popcorn' to the beginning of your grocery list.

console.log(groceryList);

console.log(groceryList.slice(1, 4));

console.log(groceryList);

const pastaIndex = groceryList.indexOf('pasta');

console.log(pastaIndex);