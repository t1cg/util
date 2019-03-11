Object.defineProperty(exports, "__esModule", { value: true });
const cp = require("child_process");
const fs = require("fs");
let TestOutput = () => {
    console.log("running test file");
    var logFile = fs.createWriteStream('./output/TestOutput.txt', { flags: 'w' });
    const capture = cp.spawn('node', ['out.js']);
    capture.stdout.on('data', function (chunk) {
        logFile.write(chunk.toString('utf8'));
    });
};
let compareFiles = (outPath, goldPath) => {
    console.log("comparing files");
    var fail = 0;
    var out = fs.readFileSync(outPath, 'utf8').trim().split("\n");
    var gold = fs.readFileSync(goldPath, 'utf8').trim().split("\n");
    var index = (out.length > gold.length) ? out.length : gold.length;
    for (let i = 0; i < index; i++) {
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
            let o = stripDate(out[i]);
            let g = stripDate(gold[i]);
            console.log(o, g);
            if (o !== g) {
                fail += 1;
                console.log('error line ', i);
                console.log("--- ", g, "\n+++", o);
            }
        }
    }
    return fail;
};
let updateFiles = (outPath, goldPath) => {
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
let stripDate = (log) => {
    if (log.trim() === "") {
        return "";
    }
    let arr = log.split(" ");
    if (Date.parse(arr[0])) {
        arr[0] = '[DATE]';
    }
    return arr.join(" ");
};
let main = () => {
    if (process.argv.length > 1 && process.argv[2] === '-u') {
        console.log("UPDATING");
        setTimeout(TestOutput, 10);
        let updated = updateFiles(__dirname + '/output/TestOutput.txt', __dirname + '/golden/TestOutput.txt');
        if (updated) {
            console.log("update succeeded");
        }
        else {
            console.log("update failed");
        }
    }
    else {
        console.log("TESTING");
        setTimeout(TestOutput, 10);
        let errors = compareFiles(__dirname + '/output/TestOutput.txt', __dirname + '/golden/TestOutput.txt');
        if (errors > 0) {
            console.log("failed with ", errors, "errors");
        }
        else {
            console.log("no errors");
        }
    }
};
main();
