$('.login-form').on('submit', function (e) {
    e.preventDefault()

    const data = $('.login-form').serialize()
    $.post('http://localhost:8080/login', data, res => {
        console.log(res)

        if (res.code === 4001)
        {
            $('.none').css('display', 'block')
            $('.wrong').css('display', 'none')
            return
        }
        else if (res.code === 4002)
        {
            $('.wrong').css('display', 'block')
            $('.none').css('display', 'none')
            return
        }
        if (res.code === 2000)
        {
            window.localStorage.setItem('token', res.token);
            console.log(res.token);
            window.location.href = '/home'
            }
    })
    
})
    