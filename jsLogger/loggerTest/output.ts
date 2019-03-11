import * as cp from 'child_process';
import * as fs from 'fs';


let TestOutput = () => {
  console.log("running test file")
  var logFile: fs.WriteStream = fs.createWriteStream('./output/TestOutput.txt', { flags: 'w' })
  const capture = cp.spawn('node', ['out.js']);

  capture.stdout.on('data', function(chunk) {
      logFile.write(chunk.toString('utf8'));
  });
}

let compareFiles = (outPath: string, goldPath: string): number => {
  console.log("comparing files")
  var fail: number = 0;
  var out: string[] = fs.readFileSync(outPath, 'utf8').trim().split("\n");
  var gold: string[] = fs.readFileSync(goldPath, 'utf8').trim().split("\n");
  var index: number = (out.length > gold.length) ? out.length : gold.length

  for (let i: number = 0; i < index; i++) {
      if (i > out.length - 1) {
          fail += 1;
          console.log('error line ', i)
          console.log("+++ ", stripDate(gold[i]))
      } else if (i > gold.length - 1) {
          fail += 1;
          console.log('error line ', i)
          console.log("--- ", stripDate(out[i]))
      } else {
          let o: string = stripDate(out[i])
          let g: string = stripDate(gold[i])

          console.log(o, g)
          
          if (o !== g) {
              fail += 1;
              console.log('error line ', i)
              console.log("--- ", g, "\n+++", o)
          }
      }
  }
  return fail
}

let updateFiles = (outPath: string, goldPath: string): boolean => {
  console.log("updating files")
  try {
      var out: string = fs.readFileSync(outPath, 'utf8')
      var logFile: fs.WriteStream = fs.createWriteStream(goldPath, { flags: 'w' })
      logFile.write(out)
      logFile.close()
      return true
  } catch (e) {
      console.log(e)
      return false
  }
}

// remove date from logs for testing purposes
let stripDate = (log: string): string => {
  if(log.trim() === "") {
      return "";
  }
  let arr: string[] = log.split(" ")
  if (Date.parse(arr[0])){
      arr[0] = '[DATE]'
  }
  return arr.join(" ")
}

let main = () => {
  if (process.argv.length > 1 && process.argv[2] === '-u') {
      console.log("UPDATING")
      setTimeout(TestOutput, 10)
      let updated: boolean = updateFiles(__dirname + '/output/TestOutput.txt', __dirname + '/golden/TestOutput.txt')
      if (updated) {
          console.log("update succeeded")
      } else {
          console.log("update failed")
      }
  } else {
      console.log("TESTING")
      setTimeout(TestOutput, 10)
      let errors: number = compareFiles(__dirname + '/output/TestOutput.txt', __dirname + '/golden/TestOutput.txt')
      if (errors > 0) {
          console.log("failed with ", errors, "errors")
      } else {
          console.log("no errors")
      }
  }
}

main()
