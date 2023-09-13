// library.js

const inventory = {
    sunglasses: 5,
    bags: 10,
    shoes: 8,
    // ... other items and quantities
  };
  
  const checkInventory = (order) => {
    return new Promise((resolve, reject) => {
      const itemsInStock = Object.keys(inventory);
  
      const itemsToOrder = order.map(item => item[0]);
      const quantitiesToOrder = order.map(item => item[1]);
  
      for (let i = 0; i < itemsToOrder.length; i++) {
        const itemToCheck = itemsToOrder[i];
        if (!itemsInStock.includes(itemToCheck)) {
          reject(`Item "${itemToCheck}" is not in stock.`);
          return;
        }
  
        const quantityInStock = inventory[itemToCheck];
        const quantityToOrder = quantitiesToOrder[i];
  
        if (quantityInStock < quantityToOrder) {
          reject(`Not enough "${itemToCheck}" in stock. Available: ${quantityInStock}`);
          return;
        }
      }
  
      const remainingInventory = {};
      itemsToOrder.forEach((item, index) => {
        remainingInventory[item] = inventory[item] - quantitiesToOrder[index];
      });
  
      resolve(remainingInventory);
    });
  };
  
  module.exports = { checkInventory };
  