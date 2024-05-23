const R = require("ramda");

//church encoding
let _0 = fn => x => x;
let _1 = fn => x => fn(x);
let _2 = fn => x => fn(fn(x));
let _3 = fn => x => R.pipe(...R.repeat(fn,3))(x);
let _256 = [_0,_1,_2,_3,...R.range(4,256).map(z => fn => x => R.pipe(...R.repeat(fn,z))(x))];

//R.range(0,256).map(x => {
//    console.log(_256[x](n => n+1)(0));//boom, mind blown
//});

succ = x => x+1
prev = x => x-1
mul = x => x*x
exp = x => mul(x)*mul(x)


