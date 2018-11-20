import React from 'react';
import ReactDOM from 'react-dom';
import * as app from '../../application/application.js'
import * as user from '../../user/user.js'

class Messages extends React.Component {
  render() {
    console.log(app.SUCCESS)
    return (
      <div>
        <h2>{user.CONNECTION_REFUSED.header}</h2>
        <p>{user.CONNECTION_REFUSED.message}</p>
      </div>
    );
  }
}


ReactDOM.render(
  <Messages />,
  document.getElementById('root')
);
