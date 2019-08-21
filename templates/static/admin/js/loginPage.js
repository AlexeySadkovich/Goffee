let loginButton = document.querySelector("#loginButton");

loginButton.addEventListener('click', async () => {

    let login = document.querySelector("#loginField").value;
    let passwd = document.querySelector("#passwordField").value;
    let msg = document.querySelector("#messageField");

    if (login.length == 0) {
        msg.innerHTML = 'Login field is empty'; 
        return
    };
    if (passwd.length == 0) {
        msg.innerHTML = 'Password field is empty'; 
        return
    };

    let data = {
        Login: login,
        Password: passwd
    };

    let response = await fetch(adress + '/auth', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(data)
    });

    if (response.ok) {
        let res = await response.json();
        let code = res[0];

        if (code == 200) {
            location.reload();
        }
    }

});