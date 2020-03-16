import React from 'react'
import '../../css/SignUp.css'
import RenderForm from '../component/forms.jsx'
import Header from '../component/header.jsx';

export default class SignUp extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            username:   '',
            email:      '',
            userpass:   '',
            checkpass:  '',
            message:    '',
            isActive:   false,
        }
    }

    passValidation(key, val) {
        this.setState((state) => {
            return { [key]: val }
        }, () => {
            if(this.state.userpass === this.state.checkpass) {
                this.setState({ message:"" });
                this.setState({ isActive:true });
            } else {
                this.setState({ message:"パスワードが一致しません" });
                this.setState({ isActive:false });
            }
        });
    }

    formSubmit(eid) {
        const target = document.getElementById(eid);
        if (this.state.loading === true) {
            target.method = 'get';
            target.submit();
        }
    }

    render() {
        return(
            <div className="Container">
                <div className="BgOpacity">
                    <div className="Space" />
                    <div className="SignupForm">
                        <label htmlFor="signUpForms">
                            <form id="checkForm" action="/kiyopon">
                                <p>新規登録</p>
                                <RenderForm
                                    title='ユーザ名'
                                    formId='userName'
                                    formName='_userName'
                                    formType='text'
                                />
                                <RenderForm
                                    title='メールアドレス'
                                    formId='mailAddress'
                                    formName='_mailAddress'
                                    formType='text'
                                />
                                <RenderForm
                                    title='パスワード'
                                    formId='userPass'
                                    formName='_userPass'
                                    formType='password'
                                    formVal={(val) => this.passValidation('userpass', val)}
                                />
                                <RenderForm
                                    title='確認用パスワード'
                                    formId='checkPass'
                                    formName='_checkPass'
                                    formType='password'
                                    formVal={(val) => this.passValidation('checkpass', val)}
                                    message={this.state.message}
                                />
                                <div className={[
                                    'post',
                                    this.state.isActive ? 'SignUpBtn' : 'CantSignUp'
                                    ].join(' ')}
                                    onClick={() => this.formSubmit('checkForm')}>
                                    <div className="BtnCentering">登録する</div>
                                </div>
                            </form>
                        </label>
                    </div>
                    <div className="Space" />
                </div>
            </div>
        );
    }
}