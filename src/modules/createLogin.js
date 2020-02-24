import ajax from "./ajax.js";
import createBackButton from "./createBackButton";

const formTemp = require('../templates/form.pug')

const loginItems = {
    classes: [
        'loginForm',
    ],
    id: 'loginForm',
    formItems: {
        email: {
            title: 'Логин',
            name: 'email',
            placeholder: 'ivan.ivanov@mail.ru',
            type: 'email',
        },
        password: {
            title: 'Пароль',
            name: 'password',
            placeholder: '',
            type: 'password'
        }
    },
    buttonName: 'Войти'
};

export function createLogin(parent = document.body) {
    parent.innerHTML = '';
    parent.innerHTML += formTemp(loginItems);
    parent.innerHTML += createBackButton();
    const loginForm = document.getElementById("loginForm");
    console.log(loginForm);
    loginForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const email = loginForm.elements['email'].value;
        const password = loginForm.elements['password'].value;



       let response = fetch('http://localhost:3001/login', {
		    method: 'POST',
            headers: {
            'Access-Control-Allow-Origin' : 'http://localhost:3001',
        }

        
    })
    let txt = response.txt;
    console.log(txt);

    // if (response.ok) { 
    //      let json = await response.json();
    // } else {
    //      alert("Ошибка HTTP: " + response.status);
    // }

        // ajax('POST', 'localhost:3001/login' ,{email,password}, (status, response) => {
        //     if (status === 200) {
        //         // createProfile(parent);
        //         // return;
        //     } else {
        //         console.log(JSON.parse(response))
        //         const {error} = JSON.parse(response);
        //         alert(error);
        //     }

        // });

    });
}
