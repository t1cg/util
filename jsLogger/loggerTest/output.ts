import Logger from '../logger';

let StandardOutput = () => {
  let L = new Logger()

  console.log("standard output")
  // for standard output  
  L.Log("one")
  L.SetLogStream("error")
  L.Log("two")
  L.SetLogStream("info")
  L.Log("three")

  console.log("file output")
  // for files
  L.SetLogFile('./output/TestStandardOutput.txt', 'trace')
  L.Log("four")
  L.SetLogFile('./output/TestStandardOutput.txt', 'error')
  L.Log("five")
  L.Log("test")
}

StandardOutput();