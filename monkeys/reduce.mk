let reduce = fn(arr,initial,f){
  let iter = fn(arr,result){
    if( len(arr) == 0 ) {
      result
    } else {
      iter( rest(arr) , f(result, first(arr)));
    }
  };
  iter(arr,initial);
};