Object.defineProperty(exports, "__esModule", { value: true });
const logger_1 = require("../logger");
let StandardOutput = () => {
    let L = new logger_1.default();
    console.log("standard output");
    // for standard output  
    L.Log("one");
    L.SetLogLevel("error");
    L.Log("two");
    L.SetLogLevel("info");
    L.Log("three");
    console.log("file output");
    // for files
    L.SetLogFile('./output/TestStandardOutput.txt', 'trace');
    L.Log("four");
    L.SetLogFile('./output/TestStandardOutput.txt', 'error');
    L.Log("five");
    L.Log("test");
};
StandardOutput();
