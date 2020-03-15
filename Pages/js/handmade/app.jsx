import React from 'react';
import ReactDOM from 'react-dom';
import Header from './header.jsx';
import SignIn from './signin.jsx';

class App extends React.Component{
    render(){
        return(
            <div>
                <div>
                    <Header />
                </div>
                <div>
                    <SignIn />
                </div>
            </div>
        );
    }
}

ReactDOM.render(<App />, document.getElementById('app'));