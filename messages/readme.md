#Adding or updating a message
1. check throughly that the message you are adding does not already exist.
2. create new branch for util
3. find the relevant csv
4. append to the bottom or make changes. Please follow the style in the readme's to ensure consistency.
5. run ```python generate.py```
6. make pr
7. email phillip.shebel@t1cg.com to have him publish the change to npm


#Using messages in a project
##For javascript
1. run npm i t1cg-messages
2. make sure it is in your package.json dependencies
3. import and use as needed
```
// es5
var messages = require('t1cg-messages');

// es6
import messages from 't1cg-messages';


console.log(messages.ApplicationMessages.SUCCESS.message.
console.log(messages.UserMessages.CONNECTION_REFUSED.message.

```



##For golang
1. go into the Gopkg.toml and add
```
[[constraint]]
  branch = "master"
  name = "github.com/t1cg/util"
```

2. run ```dep ensure```
3. include ```import "github.com/t1cg/util/messages/go/<type>"```
4. call <type>.Messages.<message name>.<message property>
