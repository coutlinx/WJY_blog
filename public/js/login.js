$(document).ready(function () {
    $('#loginForm').submit(function (e) {
        e.preventDefault();
        var loginType = $('#loginType').val();
        var usernameOrEmail = $('#usernameOrEmail').val();
        var password = $('#password').val();

        // 根据登录类型构造要发送的数据
        var data;
        if (loginType === "email") {
            data = {
                email: usernameOrEmail,
                password: password
            }
        } else {
            data = {
                account: usernameOrEmail,
                password: password
            }
        }

        // 使用jQuery的ajax方法向接口发起POST请求
        $.ajax({
            url: '/api/v1/auth/login',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (response) {
                console.log('登录成功，后端返回信息:', response);
                let data = response.data;
                // 保存token到localStorage中，这里假设后端返回的数据中token所在的属性名为"data"，你需要根据实际情况调整
                localStorage.setItem('token', data.token);
                alert("登录成功！！！")
                // 这里可以添加登录成功后的页面跳转等其他逻辑，比如跳转到用户个人中心页面
                // window.location.href = 'user_profile.html';
            },
            error: function (xhr, status, error) {
                let errorMessage = xhr.responseJSON.msg;

                alert(errorMessage);
            }
        });
    });
});