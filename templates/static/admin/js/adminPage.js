let logoutButton = document.querySelector('#logoutButton');
let deleteTableButton = document.querySelector('#deleteTableButton');
let changePassButton = document.querySelector('#changePassButton');

logoutButton.addEventListener('click', async () => {
    let response = await fetch(adress + '/deauth');

    if (response.ok) {
        let res = await response.json();
        let code = res[0];

        if (code == 200) {
            location.reload();
        }
    }
})

deleteTableButton.addEventListener('click', async () => {

    let tableName = document.querySelector('#tableName').value;
    let msg = document.querySelector('#messageField__deleteTable');
    let checkbox = document.querySelector('#confirmDeleting');

    if (tableName.length == 0) {
        msg.innerHTML = 'Fill table name'; 
        return
    }

    msg.innerHTML = ''; 

    if (!checkbox.checked) {
        return
    }

    let data = {
        Table: tableName
    };

    let response = await fetch(adress + '/table/delete', {
        method: "POST",
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(data)
    });

    if (response.ok) {
        let res = await response.json();
        let code = res[0];

        if (code == 200) {
            alert('Table ' + tableName + ' deleted');
        }
    };
})

changePassButton.addEventListener('click', async () => {

    let oldPassword = document.querySelector('#oldPassword').value;
    let newPassword = document.querySelector('#newPassword').value;
    let newPasswordConfirm = document.querySelector('#newPasswordConfirm').value;

    let msg = document.querySelector('#messageField__changePassword');

    if (oldPassword.length == 0) {
        msg.innerHTML = 'Fill old password'; 
        return
    };
    if (newPassword.length == 0) {
        msg.innerHTML = 'Fill new password'; 
        return
    };
    if (newPasswordConfirm.length  == 0) {
        msg.innerHTML = 'Confirm new password'; 
        return
    };
    if (newPasswordConfirm != newPassword) {
        msg.innerHTML = 'New password not confirmed'; 
        return
    };

    msg.innerHTML = '';

    let data = {
        OldPassword: oldPassword,
        NewPassword: newPassword
    };

    let response = await fetch(adress + '/change-password', {
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
            alert('Password changed');
        }
    };
})