
const btnTmpl = require('../templates/button.pug')
const buttonsForMainPage = {
  login: {
    name: 'Вход',
    link: 'login',
    classes: [
        'mainLink',
        ]
  },
  registration: {
    name: 'Регистрация',
    link: 'registration',
    classes: [
        'mainLink',
        ]
  },
  settings: {
    name: 'Настройки',
    link: 'settings',
    classes: [
        'mainLink',
    ]
  },
  about: {
    name: 'О проекте',
    link: 'about',
    classes: [
        'mainLink',
    ]
  },
};


export function createMainPage(parent) {
  parent.innerHTML = '';
  for (let button in buttonsForMainPage) {
      parent.innerHTML += btnTmpl(buttonsForMainPage[button]);
  }
}
