"use strict";
exports.__esModule = true;
var logger_1 = require("./logger");
function main() {
    var L = new logger_1["default"]();
    L.Log("hello");
    L.SetEncoding('json');
    L.SetLogStream('warn');
    L.Log("world");
    L.SetLogFile('../test.txt', 'info');
    L.Log("hello");
}
main();
