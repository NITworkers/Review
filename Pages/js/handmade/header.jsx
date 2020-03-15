import React from 'react'
import '../../css/Header.css'

export default class Header extends React.Component{
    render() {
        return(
            <div className="Header-wrapper">
                <div className="Header-logo">レビューサイト</div>
                <label htmlFor="menuOn">
                    <input
                        id="menuOn"
                        className="MenuImage"
                        type="checkbox"
                    />
                    <menu>
                      <ul>
                        <li><a href="#menu1">新規登録</a></li>
                        <li><a href="#menu2">ログイン</a></li>
                        <li><a href="#menu3">投稿者から探す</a></li>
                        <li><a href="#menu4">レビューから探す</a></li>
                        <li className="liSpace"><a href="#menu5">運営について</a></li>
                        <li><a href="#menu6">お問い合わせ</a></li>
                      </ul>
                    </menu>
                    <div className="overlay"></div>
                </label>
            </div>
        );
    }
}