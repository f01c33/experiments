"use strict";
const pretty = require('js-object-pretty-print').pretty;
// Object.prototype.toString = function() { return pretty(this); };

const esprima = require('esprima');
const fs = require("fs");
const toSource = require("tosource");
const olhada = (coisa) => {
    console.log(pretty(coisa));
    return coisa;
}
const restify = (name) => {
    let program;
    if(typeof name == "function"){
        program = name.toString();
    }else if (typeof name == "object"){
        program = toSource(name);
    }else{
        program = String(fs.readFileSync(name));
    }
    const parse = esprima.parseScript(program);
    
    //tree-walking
    const functions = parse.body.filter(block => block.type == "FunctionDeclaration" || (block.type == "VariableDeclaration" && block.declarations.filter(dec => dec.init.type == "ArrowFunctionExpression").length > 0));
    const variables = parse.body.filter(block => block.type == "VariableDeclaration");
    // console.log(pretty(functions));
    //code-generation
    const restified = [];
    variables.forEach(v => {
        const decs = v.declarations.filter(d => d.init.type != "ArrowFunctionExpression");
        restified.push(...decs.map(d => `
app.get('/${d.id.name}/:query', (req, res) => {
    Promise.try(()=>jp.query(${d.id.name},req.params.query))
        .then(ok)
        .catch(error)
        .finally(res.json);
})
`));
    });
    functions.forEach(f =>{
        if(f.type == "FunctionDeclaration"){
            restified.push(`
app.get('/${f.id.name}${f.params.map(p => "/:"+p.name).join("")}', (req, res) => {
    Promise.try(()=>${f.id.name}(${f.params.map(p => "req.params."+p.name)}))
        .then(ok)
        .catch(error)
        .finally(res.json);
})
`)
        }else if(f.type == "VariableDeclaration"){
            // console.log("wee");
            restified.push(...f.declarations.map(d => {
                return `
app.get('/${d.id.name}${d.init.params.map(p => "/:"+p.name).join("")}', (req, res) => {
    Promise.try(()=>${d.id.name}(${d.init.params.map(p => "req.params."+p.name)}))
        .then(ok)
        .catch(error)
        .finally(res.json);
})`}));
        }
    });
    
    return `${program}\n;\n 
const Promise = require("bluebird");
const express = require('express');
const bodyParser = require('body-parser');
const jp = require('json-path');
const port = 3000;

let error = (...opt) => {
    return { status: "erro", dados: [...opt] };
}

let ok = (...opt) => {
    return { status: "ok", dados: [...opt] };
}

const ifErr = (err, data, msg) => {
    if (!err) {
        return error(msg, err);
    }
}

const tem = (body, ...keys) => {
    for (c of keys) {
        if (!(c in body)) return false;
    }
    return true;
}

const temQueTerBody = (...keys) => {
    return (req, res, next) => {
        if (tem(req.body, ...keys)) {
            next();
        } else {
            res.json(error("Missing arguments: ", ...keys));
        }
    }
}

const temQueTerQuery = (...keys) => {
    return (req, res, next) => {
        if (tem(req.query, ...keys)) {
            next();
        } else {
            res.json(error("Missing arguments: ", ...keys));
        }
    }
}

const temQueTerParam = (...keys) => {
    return (req, res, next) => {
        if (tem(req.params, ...keys)) {
            next();
        } else {
            res.json(error("Missing arguments: ", ...keys));
        }
    }
}
const app = express();
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

${restified.join("\n")}

app.listen(port, () => {
    console.log("http://localhost:"+port);;
});`;
}

// eval(restify(process.argv[2]||"test.js"));

export const restify;