function formSubmit(eid) {
    let target = document.getElementById(eid);
    target.method = 'Post';
    target.submit();
}