const nerdamer = require("nerdamer/all.min")

const polinomial = (degree) => {
    let pol = "y="
    for (let i = 0; i <= degree; i++) {
        if (i > 0) {
            pol += "+";
        }
        pol += String.fromCharCode(97 + i) + "*x^" + String(i);
    }
    return pol
};

console.log(polinomial(0))
console.log(polinomial(1))
console.log(polinomial(2))

for (let i = 0; i < 5; i++) {
    let x = nerdamer(polinomial(i), { y: 0 })
    console.log(x.solveFor('x').toString())
}
// var e = nerdamer.solveEquations(polinomial(1), 'x')
// console.log(e.toString())