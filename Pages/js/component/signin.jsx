import React from 'react'
import '../../css/SignIn.css'
import RenderForm from './forms.jsx'

export default class SignIn extends React.Component{
    formSubmit(eid) {
        const target = document.getElementById(eid);
        target.method = 'get';
        target.submit();
    }

    render() {
        return(
            <div className="Container">
                <div className="BgOpacity">
                    <div className="Space" />
                    <div className="SigninForm">
                        <label htmlFor="signInForms">
                            <form id="checkForm" action="/kiyopon">
                                <p>ログイン</p>
                                <RenderForm
                                    title='ユーザ名'
                                    formId='userName'
                                    formName='_userName'
                                    formType='text'
                                />
                                <RenderForm
                                    title='パスワード'
                                    formId='userPass'
                                    formName='_userPass'
                                    formType='password'
                                />
                                <div className="SignInBtn" onClick={() => this.formSubmit('checkForm')}>
                                    <div className="BtnCentering">ログインする</div>
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