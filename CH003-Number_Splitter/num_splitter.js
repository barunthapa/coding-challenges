function split(num, div){
  // num - Integer that is to be split
  // div - Number of specified parts 
  var arr = [];
    for(var i = div; i > 1; i--) {
      var limit = num - (i- 1);
      var random = parseInt(Math.random()*limit) + 1;
      arr.push(random);
      num= num - random;
    }; 
  arr.push(num);
  return arr;
}
