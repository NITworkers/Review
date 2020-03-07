const $form = document.getElementById("checkForm");
const $error = document.getElementById("notMatch");

// 入力と変更を検知
$form.addEventListener('change', update, false);
$form.addEventListener('input', update, false);

function update() {

  // フォームの子要素の取得
  const $password = $form.elements['_password'];
  const $passwordConfirm = $form.elements['_checkpass'];

  // カスタムエラーを初期化
  $passwordConfirm.setCustomValidity('');
  if ($password.value !== $passwordConfirm.value) {

    if ($passwordConfirm.value != "") {
      // パスワードと新しいパスワードが一致しない場合、カスタムエラーをセットする
      $passwordConfirm.setCustomValidity('パスワードが一致しません');
    }
  }
  // エラーメッセージを更新
  $error.innerHTML = $passwordConfirm.validationMessage;
}