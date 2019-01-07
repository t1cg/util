"use strict";
exports.__esModule = true;
var logger_1 = require("./logger");
function main() {
    var L = new logger_1["default"]();
    // L.SetEncoding('json');
    // L.SetLogLevel('info')
    L.Log("hello");
    L.SetEncoding('json');
    L.SetLogLevel('warn');
    L.Log("world");
}
main();
