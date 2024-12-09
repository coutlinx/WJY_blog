$(document).ready(function () {
    $('#registerForm').submit(function (e) {
        e.preventDefault();
        var registerType = $('#registerType').val();
        var usernameOrEmail = $('#usernameOrEmail').val();
        var password = $('#password').val();

        // 根据注册类型构造要发送的数据
        var data;
        if (registerType === "email") {
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
            url: '/api/v1/auth/register',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (response) {
                // 请求成功后的处理逻辑，比如根据后端返回的信息提示用户注册成功等
                console.log('注册成功，后端返回信息:', response);
                (data) = response.data
                // 保存token到storage中
                localStorage.setItem('token', data.token);
                alert("注册成功")
                location.href = "/login"
            },
            error: function (xhr, status, error) {
                console.log('注册失败，错误信息:', error);
                console.log('状态码:', xhr.status);
                let errorMessage = '';
                try {
                    // 尝试解析后端返回的JSON格式错误信息（假设后端以JSON格式返回错误内容）
                    const responseData = JSON.parse(xhr.responseText);
                    if (responseData && responseData.message) {
                        // 如果后端返回的数据中有message属性，就将其作为错误提示信息
                        errorMessage = responseData.message;
                    } else {
                        errorMessage = '注册出现未知错误，请稍后再试';
                    }
                } catch (parseError) {
                    // 如果解析JSON出现错误，使用默认的提示信息
                    errorMessage = '注册出现未知错误，请稍后再试';
                }
                alert(errorMessage);
            }
        });
    });
});