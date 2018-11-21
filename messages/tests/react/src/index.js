import React from 'react';
import ReactDOM from 'react-dom';
import messages from 't1cg-messages';

class Messages extends React.Component {
  render() {
    console.log(messages.ApplicationMessages.SUCCESS.message)
    return (
      <div>
        <h2>{messages.UserMessages.CONNECTION_REFUSED.header}</h2>
        <p>{messages.UserMessages.CONNECTION_REFUSED.message}</p>
      </div>
    );
  }
}


ReactDOM.render(
  <Messages />,
  document.getElementById('root')
);
