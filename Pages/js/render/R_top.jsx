import React from 'react';
import ReactDOM from 'react-dom';
import Header from '../component/header.jsx';
import Top from '../component/top.jsx';

class R_Top extends React.Component{
    render(){
        return(
            <div>
                <div>
                    <Header />
                </div>
                <div>
                    <Top />
                </div>
            </div>
        );
    }
}

ReactDOM.render(<R_Top />, document.getElementById('r_top'));