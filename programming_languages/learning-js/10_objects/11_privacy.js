const robot = {
    _energyLevel: 100,
    recharge(){
        this._energyLevel += 35;
        console.log(`Recharged! Energy is currently at ${this._energyLevel} %.`)
    }
};

robot._energyLevel = 'high ';
robot.recharge();
  
  