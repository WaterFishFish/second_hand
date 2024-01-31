$('.reg-form').on('submit', function (e) {
    e.preventDefault()

    const data = $('.reg-form').serialize()
     
    $.post('http://localhost:8080/register', data, res => {
        console.log(res)
        
        if (res.code === 200)
        {
            alert("注册成功")
            window.location.href='/login'
        }
        
        if (res.code === 4004)
        {
            alert("该用户名已被使用")
            }
        if (res.code === 4005)
        {
            alert("重复密码与第一次输入的密码不一致,请重新输入")
        }
    })
})