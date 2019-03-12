"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var _this = this;
exports.__esModule = true;
var cp = require("child_process");
var fs = require("fs");
var TestOutput = function () { return __awaiter(_this, void 0, void 0, function () {
    var logFile, capture;
    return __generator(this, function (_a) {
        console.log("running test file");
        logFile = fs.createWriteStream('./output/TestOutput.txt', { flags: 'w' });
        capture = cp.spawn('node', ['output.js']);
        capture.stdout.on('data', function (chunk) {
            logFile.write(chunk.toString('utf8'));
        });
        return [2 /*return*/];
    });
}); };
var compareFiles = function (outPath, goldPath) {
    console.log("comparing files");
    var fail = 0;
    var out = fs.readFileSync(outPath, 'utf8').trim().split("\n");
    var gold = fs.readFileSync(goldPath, 'utf8').trim().split("\n");
    var index = (out.length > gold.length) ? out.length : gold.length;
    for (var i = 0; i < index; i++) {
        if (i > out.length - 1) {
            fail += 1;
            console.log('error line ', i);
            console.log("+++ ", stripDate(gold[i]));
        }
        else if (i > gold.length - 1) {
            fail += 1;
            console.log('error line ', i);
            console.log("--- ", stripDate(out[i]));
        }
        else {
            var o = stripDate(out[i]);
            var g = stripDate(gold[i]);
            if (o !== g) {
                fail += 1;
                console.log('error line ', i);
                console.log("--- ", g, "\n+++", o);
            }
        }
    }
    return fail;
};
var updateFiles = function (outPath, goldPath) {
    console.log("updating files");
    try {
        var out = fs.readFileSync(outPath, 'utf8');
        var logFile = fs.createWriteStream(goldPath, { flags: 'w' });
        logFile.write(out);
        logFile.close();
        return true;
    }
    catch (e) {
        console.log(e);
        return false;
    }
};
// remove date from logs for testing purposes
var stripDate = function (log) {
    if (log.trim() === "") {
        return "";
    }
    var arr = log.split(" ");
    if (Date.parse(arr[0])) {
        arr[0] = '[DATE]';
        arr[1] = '[TIME]';
    }
    return arr.join(" ");
};
var main = function () {
    if (process.argv.length > 1 && process.argv[2] === '-u') {
        console.log("UPDATING");
        setTimeout(TestOutput, 10);
        var updated = updateFiles(__dirname + '/output/TestOutput.txt', __dirname + '/golden/TestOutput.txt');
        if (updated) {
            console.log("update succeeded");
        }
        else {
            console.log("update failed");
        }
    }
    else {
        console.log("TESTING");
        TestOutput()
            .then(function () {
            var errors = compareFiles(__dirname + '/output/TestOutput.txt', __dirname + '/golden/TestOutput.txt');
            if (errors > 0) {
                console.log("failed with ", errors, "errors");
            }
            else {
                console.log("no errors");
            }
        });
    }
};
main();
