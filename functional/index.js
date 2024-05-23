"use strict";
const beautify = require('js-beautify').js;
var toSource = require("tosource-polyfill");
// function builder, that does arbitrary out of order functionally composable api for metaprogramming
const optsJSbeautify = {
    "indent_size": 4,
    "indent_char": " ",
    "indent_with_tabs": false,
    "editorconfig": false,
    "eol": "\n",
    "end_with_newline": false,
    "indent_level": 0,
    "preserve_newlines": true,
    "max_preserve_newlines": 0,
    "space_in_paren": false,
    "space_in_empty_paren": false,
    "jslint_happy": false,
    "space_after_anon_function": false,
    "space_after_named_function": false,
    "brace_style": "collapse",
    "unindent_chained_methods": false,
    "break_chained_methods": false,
    "keep_array_indentation": false,
    "unescape_strings": false,
    "wrap_line_length": 0,
    "e4x": false,
    "comma_first": false,
    "operator_position": "before-newline",
    "indent_empty_lines": false,
    "templating": ["auto"]
};

const fx = () => {
    const fxb = {
        params:[],
        body:[],
        normalize:() => {
            const tmp = beautify(`(\n${fxb.params.join(",")}\n)=>{${fxb.body.join(";")}}`,optsJSbeautify).split("\n");
            fxb.body = tmp.slice(3,tmp.length-1);
            fxb.params= tmp[1].split(",");
            return fxb;
        },
        addParam:(...params) => {
            fxb.params.push(...params);
            return fxb;
        },
        addBody:(...params) => {
            fxb.body.push(...params);
            return fxb;
        },
        build:() => {
            return beautify(`(${fxb.params.join(",")})=>{${fxb.body.join("\n")}}`,optsJSbeautify);
        },
        load:(f) => {
            const tmp = beautify(toSource(f),optsJSbeautify).split("\n");
            fxb.addBody(...tmp.slice(1,tmp.length-1));
            fxb.addParam(...(tmp[0].match(/(?!\().+(?=\))/g)||[""])[0].split(","));
            return fxb;
        },
        eval:() => {
            return eval(fxb.build());
        },
        evalW:(...params) => {
            return eval(`(${fxb.build()})(${params.join(",")})`);
        },
    };
    return fxb;
};

const a = ()=>{
    const a = 1;
    const b = 1;
    throw "";
    // const b2 = ;
    const a2 = a+b;
    return a+b+b2+a2;
}


// const macroSum = (...params) => {
//     const rtf = fx().addParam(...params.map(x => `_${x}`));
//     return rtf.addBody(`return ${rtf.params.join("+")}`).evalW(...params);
// };

// console.log(macroSum(1,2,3,4,5,6,7,8,9,10));

const fLogger = (f) => {
    const func =  fx().load(f);
    func.body = [`let _______cont=0;`,...func.body.map(line => `console.log(_______cont++,decodeURI("${encodeURI(line)}").trim());\n${line}`)];
    // console.log(func.body);
    return func.eval();
}
console.log(fLogger(a)());