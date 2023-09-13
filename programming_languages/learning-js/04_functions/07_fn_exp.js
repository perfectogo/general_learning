const plantNeedsWater = function(day) {
    if (day === 'Wednesday') {
        return true;
    }
  
    return false;
};
  
  
const waterCheck = plantNeedsWater('Tuesday');
console.log(waterCheck); 