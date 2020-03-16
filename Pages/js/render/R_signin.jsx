import React from 'react';
import ReactDOM from 'react-dom';
import Header from '../component/header.jsx';
import SignIn from '../component/signin.jsx';

class R_SignIn extends React.Component{
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

ReactDOM.render(<R_SignIn />, document.getElementById('r_signin'));