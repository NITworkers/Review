import React from 'react';
import '../../css/Forms.css';
import PropTypes from 'prop-types';

const propTypes = {
    formVal: PropTypes.func,
};

export default class RenderForm extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            title: '',
            formId: '',
            formName: '',
            formType: '',
            message: ''
        };
    }

    inputForm(value) {
        return this.props.formVal(value);
    }

    render() {
        return (
            <div className="FormLiner">
                <li className="FormLeft">{this.props.title}</li>
                <li className="FormRight">
                    <input
                        id={this.props.formId}
                        name={this.props.formName}
                        type={this.props.formType}
                        onChange={(event) => this.inputForm(event.target.value)}
                        size="20" />
                    <small>{ this.props.message }</small>
                </li>
            </div>
        )
    }
}