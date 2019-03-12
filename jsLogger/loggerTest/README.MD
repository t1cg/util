logger.test.js has the jest test for all the helper functions
"npm test"

compare.js has the tests to make sure the output is correct
- when you run node compare.js, it updates the output file before comparing
- when you make a change you can run "node compare.js -u" to change the golden file