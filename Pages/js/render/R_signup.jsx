import React from 'react';
import ReactDOM from 'react-dom';
import Header from '../component/header.jsx';
import SignUp from '../component/signup.jsx';

class R_SignUp extends React.Component{
    render(){
        return(
            <div>
                <div>
                    <Header />
                </div>
                <div>
                    <SignUp />
                    kiuopon

                </div>
            </div>
        );
    }
}

ReactDOM.render(<R_SignUp />, document.getElementById('r_signup'));