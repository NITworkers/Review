import React from 'react';
import '../../css/Forms.css';

export default class RenderForm extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            title: '',
            formId: '',
            formName: '',
            formType: ''
        };
    }

    render() {
        return (
            <div className="FormLiner">
                <li className="FormLeft">{this.props.title}</li>
                <li className="FormRight">
                    <input id={this.props.formId} name={this.props.formName} type={this.props.formType} size="20" />
                </li>
            </div>
        )
    }
}