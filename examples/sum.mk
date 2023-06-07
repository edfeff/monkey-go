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

let sum = fn(arr){
  reduce(arr,0,fn(initial,el){initial+el});
};

let max =fn(arr){
  reduce(arr,0,fn(a,b){
    if(a>b){
      a;
    }else{
      b;
    }
  });
}

puts(sum([1,2,3,4]));
puts(max([1,2,3,4]));
